//go:build mage

/*
Copyright 2026 The BlanketOps Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Magefile provides development workflow targets for blanketops-environments-contract.
//
// Usage:
//
//	mage lint     — lint proto files with buf
//	mage gen      — generate Go and TypeScript from protos
//	mage docs     — generate per-package Markdown docs from proto comments
//	mage verify   — lint + gen
//	mage clean    — remove generated .pb.go and _pb.ts files from blanketops/
//	mage bundle   — verify + build OCI contract bundle
//	mage publish  — bundle + push to OCI registry via ORAS
package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magefile/mage/sh"
)

const (
	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorCyan  = "\033[36m"
	colorRed   = "\033[31m"
)

func step(msg string)    { fmt.Printf("%s▶ %s%s\n", colorCyan, msg, colorReset) }
func success(msg string) { fmt.Printf("%s✔ %s%s\n", colorGreen, msg, colorReset) }
func fail(msg string)    { fmt.Printf("%s✘ %s%s\n", colorRed, msg, colorReset) }

// Lint runs buf lint across the proto tree.
func Lint() error {
	step("Linting protos...")
	if err := sh.RunV("buf", "lint"); err != nil {
		fail("Lint failed")
		return err
	}
	success("Lint complete")
	return nil
}

// Gen runs buf generate for Go and TypeScript targets. Generated files land
// in blanketops/ alongside the proto sources so import paths stay stable
// across releases without a gen/ prefix.
func Gen() error {
	step("Generating contracts (Go + TypeScript)...")
	if err := sh.RunV("buf", "generate"); err != nil {
		fail("Gen failed")
		return err
	}
	success("Gen complete")
	return nil
}

// Docs generates per-package Markdown documentation from proto comments via
// protoc-gen-doc, mirroring the docs/code/ convention gomarkdoc produces in
// the sibling environments and environments-controller repos: one file per
// proto package directory (e.g. docs/code/blanketops/networks/v1alpha1.md),
// not one combined file for the whole tree. Requires protoc-gen-doc on PATH
// (go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc).
func Docs() error {
	step("Generating contract docs...")

	dirs, err := protoPackageDirs("blanketops")
	if err != nil {
		fail("Docs failed")
		return err
	}

	for _, dir := range dirs {
		outDir := filepath.Join("docs", "code", filepath.Dir(dir))
		outFile := filepath.Base(dir) + ".md"
		if err := os.MkdirAll(outDir, 0755); err != nil {
			fail("Docs failed")
			return err
		}
		tmpl := fmt.Sprintf(
			`{"version":"v2","plugins":[{"local":"protoc-gen-doc","out":%q,"opt":["markdown,%s"]}]}`,
			outDir, outFile,
		)
		if err := sh.RunV("buf", "generate", "--template", tmpl, "--path", dir); err != nil {
			fail("Docs failed")
			return err
		}
	}

	success("Docs complete")
	return nil
}

// protoPackageDirs returns every directory under root that directly
// contains at least one .proto file, sorted for deterministic generation
// order (filepath.WalkDir already walks lexically).
func protoPackageDirs(root string) ([]string, error) {
	var dirs []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || !d.IsDir() {
			return err
		}
		matches, globErr := filepath.Glob(filepath.Join(path, "*.proto"))
		if globErr != nil {
			return globErr
		}
		if len(matches) > 0 {
			dirs = append(dirs, path)
		}
		return nil
	})
	return dirs, err
}

// Verify runs lint then gen — the full contract validation and generation gate.
func Verify() error {
	step("Verifying contracts...")
	if err := Lint(); err != nil {
		return err
	}
	return Gen()
}

// Clean removes all generated Go (.pb.go) and TypeScript (_pb.ts) files from
// the blanketops/ tree. Proto source files (.proto) are preserved.
// Also removes the bundle/ directory if present.
func Clean() error {
	step("Cleaning generated files...")

	// Walk blanketops/ and remove generated Go and TypeScript files.
	// Proto files are left untouched.
	err := filepath.Walk("blanketops", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".pb.go") ||
			strings.HasSuffix(path, "_grpc.pb.go") ||
			strings.HasSuffix(path, "_pb.ts") {
			if removeErr := os.Remove(path); removeErr != nil {
				return fmt.Errorf("remove %s: %w", path, removeErr)
			}
		}
		return nil
	})
	if err != nil {
		fail("Clean failed")
		return err
	}

	// Remove bundle artifacts.
	if err := os.RemoveAll("bundle"); err != nil {
		fail("Clean failed")
		return err
	}

	success("Clean complete")
	return nil
}

// Bundle assembles the OCI contract bundle: buf image, proto sources,
// generated code, and metadata. Runs Verify first to ensure the bundle
// always reflects the latest generated output.
func Bundle() error {
	if err := Verify(); err != nil {
		return err
	}

	step("Building contract bundle...")

	if err := os.MkdirAll("bundle", 0755); err != nil {
		return err
	}

	// Canonical compiled contract — language-neutral binary descriptor.
	if err := sh.RunV("buf", "build", "-o", "bundle/image.binpb"); err != nil {
		fail("Bundle failed")
		return err
	}

	if err := tarGz("blanketops", "bundle/protos.tar.gz"); err != nil {
		fail("Bundle failed")
		return err
	}

	if err := writeMetadata("bundle/metadata.yaml"); err != nil {
		fail("Bundle failed")
		return err
	}

	success("Bundle complete")
	return nil
}

// Publish builds the bundle and pushes it to the OCI registry via ORAS.
// Requires ORAS to be installed and the registry to be authenticated.
//
// Registry and package are configurable via environment variables:
//
//	CONTRACT_REGISTRY       — default: ghcr.io
//	GITHUB_REPOSITORY_OWNER — default: ntlaletsi70
//	CONTRACT_PACKAGE        — default: environments-contract
//	VERSION or GITHUB_REF_NAME — required; falls back to latest git tag
func Publish() error {
	if err := Bundle(); err != nil {
		return err
	}

	ref, err := contractRef()
	if err != nil {
		fail("Publish failed")
		return err
	}

	step(fmt.Sprintf("Publishing %s...", ref))

	err = sh.RunV("oras", "push", ref,
		"bundle/image.binpb:application/vnd.bufbuild.buf.image.v1+binary",
		"bundle/protos.tar.gz:application/gzip",
		"bundle/metadata.yaml:application/yaml",
	)
	if err != nil {
		fail("Publish failed")
		return err
	}

	success(fmt.Sprintf("Published %s", ref))
	return nil
}

// -----------------------------------------------------------------------------
// Internal helpers
// -----------------------------------------------------------------------------

func tarGz(srcDir, dst string) error {
	if _, err := os.Stat(srcDir); err != nil {
		return fmt.Errorf("stat %s: %w", srcDir, err)
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	gz := gzip.NewWriter(out)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hdr.Name = filepath.ToSlash(path)
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(tw, f)
		return err
	})
}

func writeMetadata(dst string) error {
	meta := fmt.Sprintf(
		"version: %s\nrepository: %s\npublishedAt: %s\n",
		version(),
		envOr("GITHUB_REPOSITORY", "ntlaletsi70/blanketops-environments-contract"),
		time.Now().UTC().Format(time.RFC3339),
	)
	return os.WriteFile(dst, []byte(meta), 0644)
}

func contractRef() (string, error) {
	v := version()
	if v == "" {
		return "", fmt.Errorf("no version: set VERSION or GITHUB_REF_NAME, or tag the commit")
	}
	registry := envOr("CONTRACT_REGISTRY", "ghcr.io")
	owner := envOr("GITHUB_REPOSITORY_OWNER", "ntlaletsi70")
	name := envOr("CONTRACT_PACKAGE", "environments-contract")
	return fmt.Sprintf("%s/%s/%s:%s", registry, owner, name, v), nil
}

func version() string {
	if v := os.Getenv("VERSION"); v != "" {
		return v
	}
	if v := os.Getenv("GITHUB_REF_NAME"); v != "" {
		return v
	}
	if v, err := sh.Output("git", "describe", "--tags", "--abbrev=0"); err == nil {
		return strings.TrimSpace(v)
	}
	return ""
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
