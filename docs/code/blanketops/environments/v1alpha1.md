# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [blanketops/environments/v1alpha1/build.proto](#blanketops_environments_v1alpha1_build-proto)
    - [Build](#blanketops-environments-v1alpha1-Build)
    - [BuildPolicy](#blanketops-environments-v1alpha1-BuildPolicy)
    - [BuildSpec](#blanketops-environments-v1alpha1-BuildSpec)
    - [BuildStatus](#blanketops-environments-v1alpha1-BuildStatus)
    - [BuildStrategy](#blanketops-environments-v1alpha1-BuildStrategy)
    - [BuildTriggerPolicy](#blanketops-environments-v1alpha1-BuildTriggerPolicy)
    - [CreateBuildRequest](#blanketops-environments-v1alpha1-CreateBuildRequest)
    - [CreateBuildResponse](#blanketops-environments-v1alpha1-CreateBuildResponse)
    - [DeleteBuildRequest](#blanketops-environments-v1alpha1-DeleteBuildRequest)
    - [DeleteBuildResponse](#blanketops-environments-v1alpha1-DeleteBuildResponse)
    - [GetBuildRequest](#blanketops-environments-v1alpha1-GetBuildRequest)
    - [GetBuildResponse](#blanketops-environments-v1alpha1-GetBuildResponse)
    - [GitSource](#blanketops-environments-v1alpha1-GitSource)
    - [ListBuildsRequest](#blanketops-environments-v1alpha1-ListBuildsRequest)
    - [ListBuildsResponse](#blanketops-environments-v1alpha1-ListBuildsResponse)
    - [Param](#blanketops-environments-v1alpha1-Param)
    - [PatchBuildRequest](#blanketops-environments-v1alpha1-PatchBuildRequest)
    - [PatchBuildResponse](#blanketops-environments-v1alpha1-PatchBuildResponse)
    - [RetryPolicy](#blanketops-environments-v1alpha1-RetryPolicy)
    - [ServiceAccount](#blanketops-environments-v1alpha1-ServiceAccount)
    - [TriggerBuildRequest](#blanketops-environments-v1alpha1-TriggerBuildRequest)
    - [TriggerBuildResponse](#blanketops-environments-v1alpha1-TriggerBuildResponse)
    - [UpdateBuildRequest](#blanketops-environments-v1alpha1-UpdateBuildRequest)
    - [UpdateBuildResponse](#blanketops-environments-v1alpha1-UpdateBuildResponse)
    - [WatchBuildRequest](#blanketops-environments-v1alpha1-WatchBuildRequest)
    - [WatchBuildResponse](#blanketops-environments-v1alpha1-WatchBuildResponse)
  
    - [BuildService](#blanketops-environments-v1alpha1-BuildService)
  
- [blanketops/environments/v1alpha1/deployment.proto](#blanketops_environments_v1alpha1_deployment-proto)
    - [CreateDeploymentRequest](#blanketops-environments-v1alpha1-CreateDeploymentRequest)
    - [CreateDeploymentResponse](#blanketops-environments-v1alpha1-CreateDeploymentResponse)
    - [DeleteDeploymentRequest](#blanketops-environments-v1alpha1-DeleteDeploymentRequest)
    - [DeleteDeploymentResponse](#blanketops-environments-v1alpha1-DeleteDeploymentResponse)
    - [Deployment](#blanketops-environments-v1alpha1-Deployment)
    - [DeploymentSpec](#blanketops-environments-v1alpha1-DeploymentSpec)
    - [DeploymentStatus](#blanketops-environments-v1alpha1-DeploymentStatus)
    - [GetDeploymentRequest](#blanketops-environments-v1alpha1-GetDeploymentRequest)
    - [GetDeploymentResponse](#blanketops-environments-v1alpha1-GetDeploymentResponse)
    - [ListDeploymentsRequest](#blanketops-environments-v1alpha1-ListDeploymentsRequest)
    - [ListDeploymentsResponse](#blanketops-environments-v1alpha1-ListDeploymentsResponse)
    - [ManifestsRepo](#blanketops-environments-v1alpha1-ManifestsRepo)
    - [PatchDeploymentRequest](#blanketops-environments-v1alpha1-PatchDeploymentRequest)
    - [PatchDeploymentResponse](#blanketops-environments-v1alpha1-PatchDeploymentResponse)
    - [ServiceUnitDeploymentStatus](#blanketops-environments-v1alpha1-ServiceUnitDeploymentStatus)
    - [UpdateDeploymentRequest](#blanketops-environments-v1alpha1-UpdateDeploymentRequest)
    - [UpdateDeploymentResponse](#blanketops-environments-v1alpha1-UpdateDeploymentResponse)
    - [WatchDeploymentRequest](#blanketops-environments-v1alpha1-WatchDeploymentRequest)
    - [WatchDeploymentResponse](#blanketops-environments-v1alpha1-WatchDeploymentResponse)
  
    - [DeploymentService](#blanketops-environments-v1alpha1-DeploymentService)
  
- [blanketops/environments/v1alpha1/environment.proto](#blanketops_environments_v1alpha1_environment-proto)
    - [CreateEnvironmentRequest](#blanketops-environments-v1alpha1-CreateEnvironmentRequest)
    - [CreateEnvironmentResponse](#blanketops-environments-v1alpha1-CreateEnvironmentResponse)
    - [DeleteEnvironmentRequest](#blanketops-environments-v1alpha1-DeleteEnvironmentRequest)
    - [DeleteEnvironmentResponse](#blanketops-environments-v1alpha1-DeleteEnvironmentResponse)
    - [Environment](#blanketops-environments-v1alpha1-Environment)
    - [EnvironmentCondition](#blanketops-environments-v1alpha1-EnvironmentCondition)
    - [EnvironmentContract](#blanketops-environments-v1alpha1-EnvironmentContract)
    - [EnvironmentSpec](#blanketops-environments-v1alpha1-EnvironmentSpec)
    - [EnvironmentStatus](#blanketops-environments-v1alpha1-EnvironmentStatus)
    - [GetEnvironmentRequest](#blanketops-environments-v1alpha1-GetEnvironmentRequest)
    - [GetEnvironmentResponse](#blanketops-environments-v1alpha1-GetEnvironmentResponse)
    - [ListEnvironmentsRequest](#blanketops-environments-v1alpha1-ListEnvironmentsRequest)
    - [ListEnvironmentsResponse](#blanketops-environments-v1alpha1-ListEnvironmentsResponse)
    - [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef)
    - [PatchEnvironmentRequest](#blanketops-environments-v1alpha1-PatchEnvironmentRequest)
    - [PatchEnvironmentResponse](#blanketops-environments-v1alpha1-PatchEnvironmentResponse)
    - [UpdateEnvironmentRequest](#blanketops-environments-v1alpha1-UpdateEnvironmentRequest)
    - [UpdateEnvironmentResponse](#blanketops-environments-v1alpha1-UpdateEnvironmentResponse)
    - [WatchEnvironmentRequest](#blanketops-environments-v1alpha1-WatchEnvironmentRequest)
    - [WatchEnvironmentResponse](#blanketops-environments-v1alpha1-WatchEnvironmentResponse)
  
    - [EnvironmentService](#blanketops-environments-v1alpha1-EnvironmentService)
  
- [blanketops/environments/v1alpha1/package.proto](#blanketops_environments_v1alpha1_package-proto)
    - [CreatePackageRequest](#blanketops-environments-v1alpha1-CreatePackageRequest)
    - [CreatePackageResponse](#blanketops-environments-v1alpha1-CreatePackageResponse)
    - [DeletePackageRequest](#blanketops-environments-v1alpha1-DeletePackageRequest)
    - [DeletePackageResponse](#blanketops-environments-v1alpha1-DeletePackageResponse)
    - [GetPackageRequest](#blanketops-environments-v1alpha1-GetPackageRequest)
    - [GetPackageResponse](#blanketops-environments-v1alpha1-GetPackageResponse)
    - [ListPackagesRequest](#blanketops-environments-v1alpha1-ListPackagesRequest)
    - [ListPackagesResponse](#blanketops-environments-v1alpha1-ListPackagesResponse)
    - [Maintainer](#blanketops-environments-v1alpha1-Maintainer)
    - [Package](#blanketops-environments-v1alpha1-Package)
    - [PackageRepository](#blanketops-environments-v1alpha1-PackageRepository)
    - [PackageSpec](#blanketops-environments-v1alpha1-PackageSpec)
    - [PackageState](#blanketops-environments-v1alpha1-PackageState)
    - [PackageStatus](#blanketops-environments-v1alpha1-PackageStatus)
    - [PatchPackageRequest](#blanketops-environments-v1alpha1-PatchPackageRequest)
    - [PatchPackageResponse](#blanketops-environments-v1alpha1-PatchPackageResponse)
    - [StateRepository](#blanketops-environments-v1alpha1-StateRepository)
    - [UpdatePackageRequest](#blanketops-environments-v1alpha1-UpdatePackageRequest)
    - [UpdatePackageResponse](#blanketops-environments-v1alpha1-UpdatePackageResponse)
    - [WatchPackageRequest](#blanketops-environments-v1alpha1-WatchPackageRequest)
    - [WatchPackageResponse](#blanketops-environments-v1alpha1-WatchPackageResponse)
  
    - [PackageService](#blanketops-environments-v1alpha1-PackageService)
  
- [blanketops/environments/v1alpha1/serviceunit.proto](#blanketops_environments_v1alpha1_serviceunit-proto)
    - [BuildReference](#blanketops-environments-v1alpha1-BuildReference)
    - [BuildResolutionStatus](#blanketops-environments-v1alpha1-BuildResolutionStatus)
    - [CreateServiceUnitRequest](#blanketops-environments-v1alpha1-CreateServiceUnitRequest)
    - [CreateServiceUnitResponse](#blanketops-environments-v1alpha1-CreateServiceUnitResponse)
    - [DeleteServiceUnitRequest](#blanketops-environments-v1alpha1-DeleteServiceUnitRequest)
    - [DeleteServiceUnitResponse](#blanketops-environments-v1alpha1-DeleteServiceUnitResponse)
    - [GetServiceUnitRequest](#blanketops-environments-v1alpha1-GetServiceUnitRequest)
    - [GetServiceUnitResponse](#blanketops-environments-v1alpha1-GetServiceUnitResponse)
    - [ListServiceUnitsRequest](#blanketops-environments-v1alpha1-ListServiceUnitsRequest)
    - [ListServiceUnitsResponse](#blanketops-environments-v1alpha1-ListServiceUnitsResponse)
    - [PatchServiceUnitRequest](#blanketops-environments-v1alpha1-PatchServiceUnitRequest)
    - [PatchServiceUnitResponse](#blanketops-environments-v1alpha1-PatchServiceUnitResponse)
    - [RouteReference](#blanketops-environments-v1alpha1-RouteReference)
    - [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit)
    - [ServiceUnitCondition](#blanketops-environments-v1alpha1-ServiceUnitCondition)
    - [ServiceUnitSpec](#blanketops-environments-v1alpha1-ServiceUnitSpec)
    - [ServiceUnitStatus](#blanketops-environments-v1alpha1-ServiceUnitStatus)
    - [UpdateServiceUnitRequest](#blanketops-environments-v1alpha1-UpdateServiceUnitRequest)
    - [UpdateServiceUnitResponse](#blanketops-environments-v1alpha1-UpdateServiceUnitResponse)
    - [WatchServiceUnitRequest](#blanketops-environments-v1alpha1-WatchServiceUnitRequest)
    - [WatchServiceUnitResponse](#blanketops-environments-v1alpha1-WatchServiceUnitResponse)
    - [WorkloadReference](#blanketops-environments-v1alpha1-WorkloadReference)
  
    - [ServiceUnitService](#blanketops-environments-v1alpha1-ServiceUnitService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="blanketops_environments_v1alpha1_build-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/environments/v1alpha1/build.proto



<a name="blanketops-environments-v1alpha1-Build"></a>

### Build



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [BuildSpec](#blanketops-environments-v1alpha1-BuildSpec) |  | Desired state declared by the operator. |
| status | [BuildStatus](#blanketops-environments-v1alpha1-BuildStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-environments-v1alpha1-BuildPolicy"></a>

### BuildPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowed_triggers | [BuildTriggerPolicy](#blanketops-environments-v1alpha1-BuildTriggerPolicy) | repeated | Events that are permitted to trigger a BuildRun. At least one trigger must be specified. |
| retry | [RetryPolicy](#blanketops-environments-v1alpha1-RetryPolicy) |  | Retry behaviour on BuildRun failure. |






<a name="blanketops-environments-v1alpha1-BuildSpec"></a>

### BuildSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image | [string](#string) |  | Fully qualified target image — registry/repo/app:tag. Tag strategy: git-sha, semver, or literal. e.g. docker.io/nkanyezisolutions/for-kaniko-app:main |
| strategy | [BuildStrategy](#blanketops-environments-v1alpha1-BuildStrategy) |  | Shipwright ClusterBuildStrategy or namespaced BuildStrategy to use. |
| source | [GitSource](#blanketops-environments-v1alpha1-GitSource) |  | Git source configuration — where to clone from. |
| policy | [BuildPolicy](#blanketops-environments-v1alpha1-BuildPolicy) |  | Trigger and retry policy for this build. |
| service_account | [ServiceAccount](#blanketops-environments-v1alpha1-ServiceAccount) |  | Kubernetes ServiceAccount used by the build pod. Must have push access to the image registry. |
| params | [Param](#blanketops-environments-v1alpha1-Param) | repeated | Additional parameters passed to the Shipwright BuildStrategy. Merged with strategy defaults — these take precedence. |
| timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | Maximum duration a BuildRun may run before being cancelled. Defaults to the strategy default if unset. |
| time_to_live | [google.protobuf.Duration](#google-protobuf-Duration) |  | How long a completed BuildRun is retained before garbage collection. |
| github_event | [string](#string) |  | GitHubEvent reference — (Optional) The canonical link to the subscription that initiated this build. If provided, observers use this to propagate event metadata (e.g., Short SHA, Actor) as build tags and provenance annotations. If omitted, the build is treated as a manual execution. |






<a name="blanketops-environments-v1alpha1-BuildStatus"></a>

### BuildStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.BuildStatusPhase](#blanketops-common-v1-BuildStatusPhase) |  | Current lifecycle phase. Set by the controller. |
| image | [string](#string) |  | Fully qualified image reference including digest, once pushed. e.g. docker.io/nkanyezisolutions/for-kaniko-app@sha256:abc... |
| execution_ref | [string](#string) |  | Name of the Shipwright BuildRun that last executed this Build. Use this to fetch full build logs and step details. |
| message | [string](#string) |  | Human-readable status message. Populated on failure. |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the most recent BuildRun started. |
| completion_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the most recent BuildRun completed (success or failure). |






<a name="blanketops-environments-v1alpha1-BuildStrategy"></a>

### BuildStrategy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Shipwright BuildStrategy. Known values: kaniko | buildah-shipwright-managed-push | buildpacks-v3 |
| kind | [blanketops.common.v1.BuildStrategyKind](#blanketops-common-v1-BuildStrategyKind) |  | Scope of the strategy resource. |






<a name="blanketops-environments-v1alpha1-BuildTriggerPolicy"></a>

### BuildTriggerPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [blanketops.common.v1.BuildTriggerType](#blanketops-common-v1-BuildTriggerType) |  | The event type that permits a BuildRun to be triggered. |






<a name="blanketops-environments-v1alpha1-CreateBuildRequest"></a>

### CreateBuildRequest
CreateBuild — declare a new Build intent.
Controller emits a Shipwright Build on reconciliation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [BuildSpec](#blanketops-environments-v1alpha1-BuildSpec) |  | Desired spec for the new Build. |






<a name="blanketops-environments-v1alpha1-CreateBuildResponse"></a>

### CreateBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  | The created Build including generated metadata and initial status. |






<a name="blanketops-environments-v1alpha1-DeleteBuildRequest"></a>

### DeleteBuildRequest
DeleteBuild — delete a Build CR.
The controller garbage-collects the Shipwright Build and all BuildRuns.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Build CR to delete. |






<a name="blanketops-environments-v1alpha1-DeleteBuildResponse"></a>

### DeleteBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the Build was successfully deleted. |






<a name="blanketops-environments-v1alpha1-GetBuildRequest"></a>

### GetBuildRequest
GetBuild — fetch a Build by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Build CR to fetch. |






<a name="blanketops-environments-v1alpha1-GetBuildResponse"></a>

### GetBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  |  |






<a name="blanketops-environments-v1alpha1-GitSource"></a>

### GitSource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | SSH or HTTPS URL of the Git repository. SSH format: git@github.com:&lt;owner&gt;/&lt;repo&gt;.git |
| revision | [string](#string) |  | Branch, tag, or commit SHA to build from. e.g. main, refs/heads/main, abc1234 |
| context_dir | [string](#string) |  | Path within the repository to use as the build context. Required for kaniko strategy. Defaults to repository root. |
| depth | [uint32](#uint32) |  | Clone depth — 0 means full history. Set to 1 for faster CI builds when history is not needed. |
| ssl_verify | [bool](#bool) |  | Whether to verify the Git server TLS certificate. Set to false only for internal Git servers with self-signed certs. |
| clone_secret | [string](#string) |  | Name of the Kubernetes Secret containing SSH private key. Sourced from the environment SecretStore via ExternalSecret. Key: ssh-privatekey |






<a name="blanketops-environments-v1alpha1-ListBuildsRequest"></a>

### ListBuildsRequest
ListBuilds — list Build CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.BuildStatusPhase](#blanketops-common-v1-BuildStatusPhase) | optional | Filter by current lifecycle phase. |
| strategy | [string](#string) | optional | Filter by strategy name. e.g. kaniko | buildah-shipwright-managed-push | buildpacks-v3 |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListBuildsResponse.next_page_token. Omit for the first request. |






<a name="blanketops-environments-v1alpha1-ListBuildsResponse"></a>

### ListBuildsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| builds | [Build](#blanketops-environments-v1alpha1-Build) | repeated | The list of Build CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-environments-v1alpha1-Param"></a>

### Param



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Parameter name — must match a param declared in the BuildStrategy. |
| value | [string](#string) |  | Parameter value — overrides strategy defaults. |






<a name="blanketops-environments-v1alpha1-PatchBuildRequest"></a>

### PatchBuildRequest
PatchBuild — partial update using JSON merge patch RFC 7396.
Only specified fields are updated — others left unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Build CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;policy&#34;:{&#34;retry&#34;:{&#34;max_attempts&#34;:5}}}} |






<a name="blanketops-environments-v1alpha1-PatchBuildResponse"></a>

### PatchBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  | The patched Build as observed by the controller. |






<a name="blanketops-environments-v1alpha1-RetryPolicy"></a>

### RetryPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| on_failure | [bool](#bool) |  | Whether to automatically retry a failed BuildRun. |
| max_attempts | [uint32](#uint32) |  | Maximum number of retry attempts before marking the Build failed. |






<a name="blanketops-environments-v1alpha1-ServiceAccount"></a>

### ServiceAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Kubernetes ServiceAccount. auto-generated as build-bot if not specified. |
| secret | [string](#string) |  | Name of the Kubernetes Secret containing registry credentials. Sourced from the environment SecretStore via ExternalSecret. Must be of type kubernetes.io/dockerconfigjson. |






<a name="blanketops-environments-v1alpha1-TriggerBuildRequest"></a>

### TriggerBuildRequest
TriggerBuild — fire a BuildRun against an existing Build CR.
Build-specific RPC — no equivalent in OrganizationService.
Useful for manual reruns or CI webhook dispatch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Build CR to trigger a run against. |
| params | [Param](#blanketops-environments-v1alpha1-Param) | repeated | Per-run param overrides. Applied on top of spec.params. Useful for SHA overrides on manual triggers. |






<a name="blanketops-environments-v1alpha1-TriggerBuildResponse"></a>

### TriggerBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| execution_ref | [string](#string) |  | Name of the emitted Shipwright BuildRun. Use GetBuild or WatchBuild to track progress. |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  | The Build CR after the trigger was applied. |






<a name="blanketops-environments-v1alpha1-UpdateBuildRequest"></a>

### UpdateBuildRequest
UpdateBuild — full replace of the Build spec.
Equivalent to kubectl apply — all fields replaced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  | Full Build object including metadata and desired spec. |






<a name="blanketops-environments-v1alpha1-UpdateBuildResponse"></a>

### UpdateBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  | The updated Build as observed by the controller. |






<a name="blanketops-environments-v1alpha1-WatchBuildRequest"></a>

### WatchBuildRequest
WatchBuild — stream phase transitions for a Build CR.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Build CR to watch. |






<a name="blanketops-environments-v1alpha1-WatchBuildResponse"></a>

### WatchBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [Build](#blanketops-environments-v1alpha1-Build) |  | Current state of the Build at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |





 

 

 


<a name="blanketops-environments-v1alpha1-BuildService"></a>

### BuildService
── CRUD — aligned with OrganizationService pattern ─────────────────────

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBuild | [CreateBuildRequest](#blanketops-environments-v1alpha1-CreateBuildRequest) | [CreateBuildResponse](#blanketops-environments-v1alpha1-CreateBuildResponse) | Declare a new Build intent. Controller emits Shipwright Build &#43; BuildRun on first trigger. |
| GetBuild | [GetBuildRequest](#blanketops-environments-v1alpha1-GetBuildRequest) | [GetBuildResponse](#blanketops-environments-v1alpha1-GetBuildResponse) | Fetch a Build CR by name. |
| UpdateBuild | [UpdateBuildRequest](#blanketops-environments-v1alpha1-UpdateBuildRequest) | [UpdateBuildResponse](#blanketops-environments-v1alpha1-UpdateBuildResponse) | Full replace of a Build spec. |
| PatchBuild | [PatchBuildRequest](#blanketops-environments-v1alpha1-PatchBuildRequest) | [PatchBuildResponse](#blanketops-environments-v1alpha1-PatchBuildResponse) | Partial update via JSON merge patch RFC 7396. |
| ListBuilds | [ListBuildsRequest](#blanketops-environments-v1alpha1-ListBuildsRequest) | [ListBuildsResponse](#blanketops-environments-v1alpha1-ListBuildsResponse) | List Build CRs with optional phase/strategy filter and paging. |
| DeleteBuild | [DeleteBuildRequest](#blanketops-environments-v1alpha1-DeleteBuildRequest) | [DeleteBuildResponse](#blanketops-environments-v1alpha1-DeleteBuildResponse) | Delete a Build CR and its owned Shipwright resources. |
| TriggerBuild | [TriggerBuildRequest](#blanketops-environments-v1alpha1-TriggerBuildRequest) | [TriggerBuildResponse](#blanketops-environments-v1alpha1-TriggerBuildResponse) | Fire a BuildRun against an existing Build CR. Useful for manual reruns and CI webhook dispatch. |
| WatchBuild | [WatchBuildRequest](#blanketops-environments-v1alpha1-WatchBuildRequest) | [WatchBuildResponse](#blanketops-environments-v1alpha1-WatchBuildResponse) stream | Stream phase transitions for a Build CR. Closes when the build reaches a terminal phase (SUCCEEDED, FAILED, CANCELLED). |

 



<a name="blanketops_environments_v1alpha1_deployment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/environments/v1alpha1/deployment.proto



<a name="blanketops-environments-v1alpha1-CreateDeploymentRequest"></a>

### CreateDeploymentRequest
CreateDeployment — declare a new Deployment intent.
Controller materializes workloads on reconciliation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [DeploymentSpec](#blanketops-environments-v1alpha1-DeploymentSpec) |  | Desired spec for the new Deployment. |






<a name="blanketops-environments-v1alpha1-CreateDeploymentResponse"></a>

### CreateDeploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#blanketops-environments-v1alpha1-Deployment) |  | The created Deployment including generated metadata and initial status. |






<a name="blanketops-environments-v1alpha1-DeleteDeploymentRequest"></a>

### DeleteDeploymentRequest
DeleteDeployment — delete a Deployment CR.
The controller garbage-collects materialized workloads and any
GitOps sources owned by this deployment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Deployment CR to delete. |






<a name="blanketops-environments-v1alpha1-DeleteDeploymentResponse"></a>

### DeleteDeploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the Deployment was successfully deleted. |






<a name="blanketops-environments-v1alpha1-Deployment"></a>

### Deployment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [DeploymentSpec](#blanketops-environments-v1alpha1-DeploymentSpec) |  | Desired state declared by the operator. |
| status | [DeploymentStatus](#blanketops-environments-v1alpha1-DeploymentStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-environments-v1alpha1-DeploymentSpec"></a>

### DeploymentSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_units | [string](#string) | repeated | Names of the ServiceUnit CRs that make up this deployment. Resolved in the same namespace as this Deployment. |
| runtime | [blanketops.common.v1.DeploymentRuntime](#blanketops-common-v1-DeploymentRuntime) |  | Target runtime backend for all service units. |
| image_automation | [bool](#bool) |  | Enable automatic image updates from Build outputs. Materialized as FluxCD ImageUpdateAutomation when GitOps is active. |
| manifests_repo | [ManifestsRepo](#blanketops-environments-v1alpha1-ManifestsRepo) |  | GitOps repository containing rendered manifests. Optional — when unset, the controller applies workloads directly. |
| reconciliation_strategy | [blanketops.common.v1.ReconciliationStrategy](#blanketops-common-v1-ReconciliationStrategy) |  | Manifest application strategy. Required when manifests_repo is set. |






<a name="blanketops-environments-v1alpha1-DeploymentStatus"></a>

### DeploymentStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.DeploymentPhase](#blanketops-common-v1-DeploymentPhase) |  | Current lifecycle phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status summary. Populated on failure. |
| runtime | [blanketops.common.v1.DeploymentRuntime](#blanketops-common-v1-DeploymentRuntime) |  | Runtime actually used to materialize workloads. |
| strategy | [blanketops.common.v1.DeploymentStrategy](#blanketops-common-v1-DeploymentStrategy) |  | Rollout strategy applied by the runtime. |
| service_unit_statuses | [ServiceUnitDeploymentStatus](#blanketops-environments-v1alpha1-ServiceUnitDeploymentStatus) | repeated | Per-ServiceUnit deployment status. |
| last_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the controller last updated this status. |






<a name="blanketops-environments-v1alpha1-GetDeploymentRequest"></a>

### GetDeploymentRequest
GetDeployment — fetch a Deployment by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Deployment CR to fetch. |






<a name="blanketops-environments-v1alpha1-GetDeploymentResponse"></a>

### GetDeploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#blanketops-environments-v1alpha1-Deployment) |  |  |






<a name="blanketops-environments-v1alpha1-ListDeploymentsRequest"></a>

### ListDeploymentsRequest
ListDeployments — list Deployment CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.DeploymentPhase](#blanketops-common-v1-DeploymentPhase) | optional | Filter by current lifecycle phase. |
| runtime | [blanketops.common.v1.DeploymentRuntime](#blanketops-common-v1-DeploymentRuntime) | optional | Filter by target runtime. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListDeploymentsResponse.next_page_token. Omit for the first request. |






<a name="blanketops-environments-v1alpha1-ListDeploymentsResponse"></a>

### ListDeploymentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployments | [Deployment](#blanketops-environments-v1alpha1-Deployment) | repeated | The list of Deployment CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-environments-v1alpha1-ManifestsRepo"></a>

### ManifestsRepo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | SSH or HTTPS URL of the manifests repository. SSH format: git@github.com:&lt;owner&gt;/&lt;repo&gt;.git |
| ref | [string](#string) |  | Git ref to reconcile from — branch, tag, or commit SHA. |
| clone_secret | [string](#string) |  | Name of the Kubernetes Secret used to authenticate the clone. Sourced from the environment SecretStore via ExternalSecret. |
| strategy | [string](#string) |  | Application strategy hint for the repository contents. e.g. kustomize, helm |
| path | [string](#string) |  | Path inside the repository to reconcile. e.g. ./clusters/dev |






<a name="blanketops-environments-v1alpha1-PatchDeploymentRequest"></a>

### PatchDeploymentRequest
PatchDeployment — partial update using JSON merge patch RFC 7396.
Only specified fields are updated — others left unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Deployment CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;image_automation&#34;:true}} |






<a name="blanketops-environments-v1alpha1-PatchDeploymentResponse"></a>

### PatchDeploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#blanketops-environments-v1alpha1-Deployment) |  | The patched Deployment as observed by the controller. |






<a name="blanketops-environments-v1alpha1-ServiceUnitDeploymentStatus"></a>

### ServiceUnitDeploymentStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the ServiceUnit CR. |
| phase | [blanketops.common.v1.DeploymentPhase](#blanketops-common-v1-DeploymentPhase) |  | Current deployment phase for this unit. |
| image | [string](#string) |  | Fully qualified image that was deployed, including digest when known. |
| runtime | [blanketops.common.v1.DeploymentRuntime](#blanketops-common-v1-DeploymentRuntime) |  | Runtime backend used for this unit. |
| message | [string](#string) |  | Human-readable status message. |
| error | [string](#string) |  | Error detail when phase is FAILED. |
| last_transition_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When this unit last changed phase. |






<a name="blanketops-environments-v1alpha1-UpdateDeploymentRequest"></a>

### UpdateDeploymentRequest
UpdateDeployment — full replace of the Deployment spec.
Equivalent to kubectl apply — all fields replaced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#blanketops-environments-v1alpha1-Deployment) |  | Full Deployment object including metadata and desired spec. |






<a name="blanketops-environments-v1alpha1-UpdateDeploymentResponse"></a>

### UpdateDeploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#blanketops-environments-v1alpha1-Deployment) |  | The updated Deployment as observed by the controller. |






<a name="blanketops-environments-v1alpha1-WatchDeploymentRequest"></a>

### WatchDeploymentRequest
WatchDeployment — stream phase transitions for a Deployment CR.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Deployment CR to watch. |






<a name="blanketops-environments-v1alpha1-WatchDeploymentResponse"></a>

### WatchDeploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#blanketops-environments-v1alpha1-Deployment) |  | Current state of the Deployment at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |





 

 

 


<a name="blanketops-environments-v1alpha1-DeploymentService"></a>

### DeploymentService
── CRUD — aligned with BuildService pattern ─────────────────────────────

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateDeployment | [CreateDeploymentRequest](#blanketops-environments-v1alpha1-CreateDeploymentRequest) | [CreateDeploymentResponse](#blanketops-environments-v1alpha1-CreateDeploymentResponse) | Declare a new Deployment intent. |
| GetDeployment | [GetDeploymentRequest](#blanketops-environments-v1alpha1-GetDeploymentRequest) | [GetDeploymentResponse](#blanketops-environments-v1alpha1-GetDeploymentResponse) | Fetch a Deployment CR by name. |
| UpdateDeployment | [UpdateDeploymentRequest](#blanketops-environments-v1alpha1-UpdateDeploymentRequest) | [UpdateDeploymentResponse](#blanketops-environments-v1alpha1-UpdateDeploymentResponse) | Full replace of a Deployment spec. |
| PatchDeployment | [PatchDeploymentRequest](#blanketops-environments-v1alpha1-PatchDeploymentRequest) | [PatchDeploymentResponse](#blanketops-environments-v1alpha1-PatchDeploymentResponse) | Partial update via JSON merge patch RFC 7396. |
| ListDeployments | [ListDeploymentsRequest](#blanketops-environments-v1alpha1-ListDeploymentsRequest) | [ListDeploymentsResponse](#blanketops-environments-v1alpha1-ListDeploymentsResponse) | List Deployment CRs with optional phase/runtime filter and paging. |
| DeleteDeployment | [DeleteDeploymentRequest](#blanketops-environments-v1alpha1-DeleteDeploymentRequest) | [DeleteDeploymentResponse](#blanketops-environments-v1alpha1-DeleteDeploymentResponse) | Delete a Deployment CR and its materialized workloads. |
| WatchDeployment | [WatchDeploymentRequest](#blanketops-environments-v1alpha1-WatchDeploymentRequest) | [WatchDeploymentResponse](#blanketops-environments-v1alpha1-WatchDeploymentResponse) stream | Stream phase transitions for a Deployment CR. Streams indefinitely — deployments have no terminal phase; close client-side. |

 



<a name="blanketops_environments_v1alpha1_environment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/environments/v1alpha1/environment.proto
Copyright 2026 The BlanketOps Authors.
Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


<a name="blanketops-environments-v1alpha1-CreateEnvironmentRequest"></a>

### CreateEnvironmentRequest
Requests / Responses


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [EnvironmentSpec](#blanketops-environments-v1alpha1-EnvironmentSpec) |  | Desired spec for the new Environment. |






<a name="blanketops-environments-v1alpha1-CreateEnvironmentResponse"></a>

### CreateEnvironmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [Environment](#blanketops-environments-v1alpha1-Environment) |  | The created Environment including generated metadata and initial status. |






<a name="blanketops-environments-v1alpha1-DeleteEnvironmentRequest"></a>

### DeleteEnvironmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Environment CR to delete. |






<a name="blanketops-environments-v1alpha1-DeleteEnvironmentResponse"></a>

### DeleteEnvironmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the Environment was successfully deleted. |






<a name="blanketops-environments-v1alpha1-Environment"></a>

### Environment
Environment

The envelope of the delivery chain: a versioned, isolated execution
context where applications run (ESP-0001). Composes the CRs that make
up an application&#39;s delivery — builds, packages, service units,
deployment, and routing — by reference, and owns them via ownerReference
(cascade delete).

Resource


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [EnvironmentSpec](#blanketops-environments-v1alpha1-EnvironmentSpec) |  | Desired state declared by the operator. |
| status | [EnvironmentStatus](#blanketops-environments-v1alpha1-EnvironmentStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-environments-v1alpha1-EnvironmentCondition"></a>

### EnvironmentCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Condition type. e.g. BuildReady, DeploymentReady, RouteReady |
| status | [string](#string) |  | Condition status: True, False, or Unknown. |
| reason | [string](#string) |  | Machine-readable reason for the last transition. |
| message | [string](#string) |  | Human-readable detail for the last transition. |
| last_transition_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the condition last changed status. |






<a name="blanketops-environments-v1alpha1-EnvironmentContract"></a>

### EnvironmentContract
Contract
EnvironmentContract — platform-level bindings declared per environment.
Drives ESO ClusterSecretStore selection at reconciliation time.
Each composed CR owns its own credential secrets — this declares only
which provider backs the environment so the platform can select the
correct ClusterSecretStore. Users may reference their own store by
setting provider accordingly.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_store | [blanketops.common.v1.SecretStore](#blanketops-common-v1-SecretStore) |  | External secrets backend for this environment. Drives ESO ClusterSecretStore selection. Each CR manages its own credential secrets independently. |






<a name="blanketops-environments-v1alpha1-EnvironmentSpec"></a>

### EnvironmentSpec
Spec (intent)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_name | [string](#string) |  | Human-readable application identifier. e.g. for-kaniko-app |
| branch | [string](#string) |  | Git branch associated with this environment. e.g. main, feature/checkout |
| git_owner | [string](#string) |  | Owning Git organisation or user. e.g. ntlaletsi70 |
| environment_type | [blanketops.common.v1.EnvironmentType](#blanketops-common-v1-EnvironmentType) |  | Environment classification. |
| version | [string](#string) |  | Application version — semantic or otherwise. e.g. v1.2.3 |
| description | [string](#string) |  | Human-readable description. Optional. |
| service_units | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) | repeated | ServiceUnits composed into this environment. |
| deployment | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) |  | Deployment intent — how service units are materialized. |
| route | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) |  | External routing. Optional. |
| package | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) |  | Packaging / release artifact. Optional. |
| build | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) |  | Build configuration. |
| git_repository | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) |  | Git repository associated with this environment. The GitRepository CR owns branch, owner, and clone config. |
| contract | [EnvironmentContract](#blanketops-environments-v1alpha1-EnvironmentContract) |  | Platform-level bindings — secret store provider selection. |
| git_hub_event | [ObjectRef](#blanketops-environments-v1alpha1-ObjectRef) |  | GitHub event source for this environment. The GitHubEvent CR owns webhook event routing and Argo Events config. |






<a name="blanketops-environments-v1alpha1-EnvironmentStatus"></a>

### EnvironmentStatus
Status (observed state)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.EnvironmentPhase](#blanketops-common-v1-EnvironmentPhase) |  | Current aggregate phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status summary. Populated on failure. |
| conditions | [EnvironmentCondition](#blanketops-environments-v1alpha1-EnvironmentCondition) | repeated | Per-resource readiness conditions for UI and automation consumers. e.g. BuildReady, DeploymentReady, RouteReady |
| last_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the controller last updated this status. |






<a name="blanketops-environments-v1alpha1-GetEnvironmentRequest"></a>

### GetEnvironmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Environment CR to fetch. |






<a name="blanketops-environments-v1alpha1-GetEnvironmentResponse"></a>

### GetEnvironmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [Environment](#blanketops-environments-v1alpha1-Environment) |  |  |






<a name="blanketops-environments-v1alpha1-ListEnvironmentsRequest"></a>

### ListEnvironmentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.EnvironmentPhase](#blanketops-common-v1-EnvironmentPhase) | optional | Filter by current aggregate phase. |
| environment_type | [blanketops.common.v1.EnvironmentType](#blanketops-common-v1-EnvironmentType) | optional | Filter by environment classification. |
| application_name | [string](#string) | optional | Filter by application name. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListEnvironmentsResponse.next_page_token. Omit for the first request. |






<a name="blanketops-environments-v1alpha1-ListEnvironmentsResponse"></a>

### ListEnvironmentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environments | [Environment](#blanketops-environments-v1alpha1-Environment) | repeated | The list of Environment CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-environments-v1alpha1-ObjectRef"></a>

### ObjectRef
References (composition, not ownership)
Reference to a CR composed into this environment.
Resolved in the same namespace as the Environment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the referenced CR. |






<a name="blanketops-environments-v1alpha1-PatchEnvironmentRequest"></a>

### PatchEnvironmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Environment CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;version&#34;:&#34;v1.3.0&#34;}} |






<a name="blanketops-environments-v1alpha1-PatchEnvironmentResponse"></a>

### PatchEnvironmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [Environment](#blanketops-environments-v1alpha1-Environment) |  | The patched Environment as observed by the controller. |






<a name="blanketops-environments-v1alpha1-UpdateEnvironmentRequest"></a>

### UpdateEnvironmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [Environment](#blanketops-environments-v1alpha1-Environment) |  | Full Environment object including metadata and desired spec. |






<a name="blanketops-environments-v1alpha1-UpdateEnvironmentResponse"></a>

### UpdateEnvironmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [Environment](#blanketops-environments-v1alpha1-Environment) |  | The updated Environment as observed by the controller. |






<a name="blanketops-environments-v1alpha1-WatchEnvironmentRequest"></a>

### WatchEnvironmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Environment CR to watch. |






<a name="blanketops-environments-v1alpha1-WatchEnvironmentResponse"></a>

### WatchEnvironmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [Environment](#blanketops-environments-v1alpha1-Environment) |  | Current state of the Environment at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |





 

 

 


<a name="blanketops-environments-v1alpha1-EnvironmentService"></a>

### EnvironmentService
Service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateEnvironment | [CreateEnvironmentRequest](#blanketops-environments-v1alpha1-CreateEnvironmentRequest) | [CreateEnvironmentResponse](#blanketops-environments-v1alpha1-CreateEnvironmentResponse) | ── CRUD — aligned with BuildService pattern ───────────────────────────── Declare a new Environment. |
| GetEnvironment | [GetEnvironmentRequest](#blanketops-environments-v1alpha1-GetEnvironmentRequest) | [GetEnvironmentResponse](#blanketops-environments-v1alpha1-GetEnvironmentResponse) | Fetch an Environment CR by name. |
| UpdateEnvironment | [UpdateEnvironmentRequest](#blanketops-environments-v1alpha1-UpdateEnvironmentRequest) | [UpdateEnvironmentResponse](#blanketops-environments-v1alpha1-UpdateEnvironmentResponse) | Full replace of an Environment spec. |
| PatchEnvironment | [PatchEnvironmentRequest](#blanketops-environments-v1alpha1-PatchEnvironmentRequest) | [PatchEnvironmentResponse](#blanketops-environments-v1alpha1-PatchEnvironmentResponse) | Partial update via JSON merge patch RFC 7396. |
| ListEnvironments | [ListEnvironmentsRequest](#blanketops-environments-v1alpha1-ListEnvironmentsRequest) | [ListEnvironmentsResponse](#blanketops-environments-v1alpha1-ListEnvironmentsResponse) | List Environment CRs with optional phase/type/application filter and paging. |
| DeleteEnvironment | [DeleteEnvironmentRequest](#blanketops-environments-v1alpha1-DeleteEnvironmentRequest) | [DeleteEnvironmentResponse](#blanketops-environments-v1alpha1-DeleteEnvironmentResponse) | Delete an Environment CR and cascade-delete owned resources. |
| WatchEnvironment | [WatchEnvironmentRequest](#blanketops-environments-v1alpha1-WatchEnvironmentRequest) | [WatchEnvironmentResponse](#blanketops-environments-v1alpha1-WatchEnvironmentResponse) stream | ── Streaming ──────────────────────────────────────────────────────────── Stream phase transitions for an Environment CR. Streams indefinitely — environments have no terminal phase; close client-side. |

 



<a name="blanketops_environments_v1alpha1_package-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/environments/v1alpha1/package.proto



<a name="blanketops-environments-v1alpha1-CreatePackageRequest"></a>

### CreatePackageRequest
CreatePackage — declare a new Package.
Controller reconciles the package sources on creation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [PackageSpec](#blanketops-environments-v1alpha1-PackageSpec) |  | Desired spec for the new Package. |






<a name="blanketops-environments-v1alpha1-CreatePackageResponse"></a>

### CreatePackageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| package | [Package](#blanketops-environments-v1alpha1-Package) |  | The created Package including generated metadata and initial status. |






<a name="blanketops-environments-v1alpha1-DeletePackageRequest"></a>

### DeletePackageRequest
DeletePackage — delete a Package CR.
The controller garbage-collects reconciliation sources owned by
this package. State repository contents are not modified.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Package CR to delete. |






<a name="blanketops-environments-v1alpha1-DeletePackageResponse"></a>

### DeletePackageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the Package was successfully deleted. |






<a name="blanketops-environments-v1alpha1-GetPackageRequest"></a>

### GetPackageRequest
GetPackage — fetch a Package by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Package CR to fetch. |






<a name="blanketops-environments-v1alpha1-GetPackageResponse"></a>

### GetPackageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| package | [Package](#blanketops-environments-v1alpha1-Package) |  |  |






<a name="blanketops-environments-v1alpha1-ListPackagesRequest"></a>

### ListPackagesRequest
ListPackages — list Package CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.PackagePhase](#blanketops-common-v1-PackagePhase) | optional | Filter by current lifecycle phase. |
| enabled | [bool](#bool) | optional | Filter by enabled state. |
| name | [string](#string) | optional | Filter by logical package name. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListPackagesResponse.next_page_token. Omit for the first request. |






<a name="blanketops-environments-v1alpha1-ListPackagesResponse"></a>

### ListPackagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| packages | [Package](#blanketops-environments-v1alpha1-Package) | repeated | The list of Package CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-environments-v1alpha1-Maintainer"></a>

### Maintainer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Maintainer display name. |
| email | [string](#string) |  | Maintainer contact email. |






<a name="blanketops-environments-v1alpha1-Package"></a>

### Package



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [PackageSpec](#blanketops-environments-v1alpha1-PackageSpec) |  | Desired state declared by the operator. |
| status | [PackageStatus](#blanketops-environments-v1alpha1-PackageStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-environments-v1alpha1-PackageRepository"></a>

### PackageRepository



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | SSH or HTTPS URL of the package repository. SSH format: git@github.com:&lt;owner&gt;/&lt;repo&gt;.git |
| credentials_secret | [string](#string) |  | Name of the Kubernetes Secret containing repository credentials. Sourced from the environment SecretStore via ExternalSecret. |






<a name="blanketops-environments-v1alpha1-PackageSpec"></a>

### PackageSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether this package is enabled. Disabled packages are not reconciled but the CR is retained. |
| name | [string](#string) |  | Logical name of the package — may differ from metadata.name. e.g. for-kaniko-app |
| version | [string](#string) |  | Version identifier of the package contents. e.g. v1.2.3 |
| maintainers | [Maintainer](#blanketops-environments-v1alpha1-Maintainer) | repeated | Package maintainers, for provenance and contact. |
| description | [string](#string) |  | Human-readable description of the package contents. |
| repository | [PackageRepository](#blanketops-environments-v1alpha1-PackageRepository) |  | Source repository containing the package definitions. |
| diff_enabled | [bool](#bool) |  | Enable diff preview (kapp) during reconciliation. Diff output is recorded in status for audit. |
| state_repository | [StateRepository](#blanketops-environments-v1alpha1-StateRepository) |  | Repository where rendered/applied state is stored. The FluxCD reconciliation source of truth. |






<a name="blanketops-environments-v1alpha1-PackageState"></a>

### PackageState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| repository | [string](#string) |  | Repository the applied state was sourced from. |
| revision | [string](#string) |  | Revision (commit SHA) of the applied state. |






<a name="blanketops-environments-v1alpha1-PackageStatus"></a>

### PackageStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.PackagePhase](#blanketops-common-v1-PackagePhase) |  | Current lifecycle phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status message. Populated on failure. |
| last_reconciled_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When reconciliation was last attempted. |
| diff_applied | [bool](#bool) |  | Whether the last diff resulted in applied changes. |
| diff_output | [string](#string) |  | Raw diff output from the last reconciliation, for audit and debug. |
| state | [PackageState](#blanketops-environments-v1alpha1-PackageState) |  | Applied source-of-truth state. |






<a name="blanketops-environments-v1alpha1-PatchPackageRequest"></a>

### PatchPackageRequest
PatchPackage — partial update using JSON merge patch RFC 7396.
Only specified fields are updated — others left unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Package CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;version&#34;:&#34;v1.3.0&#34;}} |






<a name="blanketops-environments-v1alpha1-PatchPackageResponse"></a>

### PatchPackageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| package | [Package](#blanketops-environments-v1alpha1-Package) |  | The patched Package as observed by the controller. |






<a name="blanketops-environments-v1alpha1-StateRepository"></a>

### StateRepository



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | SSH or HTTPS URL of the state repository. |
| ref | [string](#string) |  | Git ref to reconcile against — branch, tag, or commit SHA. |
| clone_secret | [string](#string) |  | Name of the Kubernetes Secret used to authenticate the clone. |
| strategy | [string](#string) |  | Application strategy for the state contents. e.g. kustomization |
| path | [string](#string) |  | Path inside the repository to reconcile. e.g. ./clusters/dev |






<a name="blanketops-environments-v1alpha1-UpdatePackageRequest"></a>

### UpdatePackageRequest
UpdatePackage — full replace of the Package spec.
Equivalent to kubectl apply — all fields replaced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| package | [Package](#blanketops-environments-v1alpha1-Package) |  | Full Package object including metadata and desired spec. |






<a name="blanketops-environments-v1alpha1-UpdatePackageResponse"></a>

### UpdatePackageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| package | [Package](#blanketops-environments-v1alpha1-Package) |  | The updated Package as observed by the controller. |






<a name="blanketops-environments-v1alpha1-WatchPackageRequest"></a>

### WatchPackageRequest
WatchPackage — stream phase transitions for a Package CR.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Package CR to watch. |






<a name="blanketops-environments-v1alpha1-WatchPackageResponse"></a>

### WatchPackageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| package | [Package](#blanketops-environments-v1alpha1-Package) |  | Current state of the Package at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |





 

 

 


<a name="blanketops-environments-v1alpha1-PackageService"></a>

### PackageService
── CRUD — aligned with BuildService pattern ─────────────────────────────

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreatePackage | [CreatePackageRequest](#blanketops-environments-v1alpha1-CreatePackageRequest) | [CreatePackageResponse](#blanketops-environments-v1alpha1-CreatePackageResponse) | Declare a new Package. |
| GetPackage | [GetPackageRequest](#blanketops-environments-v1alpha1-GetPackageRequest) | [GetPackageResponse](#blanketops-environments-v1alpha1-GetPackageResponse) | Fetch a Package CR by name. |
| UpdatePackage | [UpdatePackageRequest](#blanketops-environments-v1alpha1-UpdatePackageRequest) | [UpdatePackageResponse](#blanketops-environments-v1alpha1-UpdatePackageResponse) | Full replace of a Package spec. |
| PatchPackage | [PatchPackageRequest](#blanketops-environments-v1alpha1-PatchPackageRequest) | [PatchPackageResponse](#blanketops-environments-v1alpha1-PatchPackageResponse) | Partial update via JSON merge patch RFC 7396. |
| ListPackages | [ListPackagesRequest](#blanketops-environments-v1alpha1-ListPackagesRequest) | [ListPackagesResponse](#blanketops-environments-v1alpha1-ListPackagesResponse) | List Package CRs with optional phase/enabled/name filter and paging. |
| DeletePackage | [DeletePackageRequest](#blanketops-environments-v1alpha1-DeletePackageRequest) | [DeletePackageResponse](#blanketops-environments-v1alpha1-DeletePackageResponse) | Delete a Package CR and its owned reconciliation sources. |
| WatchPackage | [WatchPackageRequest](#blanketops-environments-v1alpha1-WatchPackageRequest) | [WatchPackageResponse](#blanketops-environments-v1alpha1-WatchPackageResponse) stream | Stream phase transitions for a Package CR. Streams indefinitely — packages have no terminal phase; close client-side. |

 



<a name="blanketops_environments_v1alpha1_serviceunit-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/environments/v1alpha1/serviceunit.proto



<a name="blanketops-environments-v1alpha1-BuildReference"></a>

### BuildReference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the source Build CR. |
| namespace | [string](#string) |  | Namespace of the source Build CR. Defaults to the ServiceUnit&#39;s namespace if empty. |






<a name="blanketops-environments-v1alpha1-BuildResolutionStatus"></a>

### BuildResolutionStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the resolved Build CR. |
| phase | [blanketops.common.v1.ServiceUnitDeploymentPhase](#blanketops-common-v1-ServiceUnitDeploymentPhase) |  | Observed phase of the source Build. |
| image | [string](#string) |  | Image produced by the source Build, including digest. |
| error | [string](#string) |  | Error detail when resolution failed. |
| last_transition_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the resolution last changed. |






<a name="blanketops-environments-v1alpha1-CreateServiceUnitRequest"></a>

### CreateServiceUnitRequest
CreateServiceUnit — declare a new ServiceUnit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ServiceUnitSpec](#blanketops-environments-v1alpha1-ServiceUnitSpec) |  | Desired spec for the new ServiceUnit. |






<a name="blanketops-environments-v1alpha1-CreateServiceUnitResponse"></a>

### CreateServiceUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_unit | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) |  | The created ServiceUnit including generated metadata and initial status. |






<a name="blanketops-environments-v1alpha1-DeleteServiceUnitRequest"></a>

### DeleteServiceUnitRequest
DeleteServiceUnit — delete a ServiceUnit CR.
The controller garbage-collects the materialized workload.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the ServiceUnit CR to delete. |






<a name="blanketops-environments-v1alpha1-DeleteServiceUnitResponse"></a>

### DeleteServiceUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the ServiceUnit was successfully deleted. |






<a name="blanketops-environments-v1alpha1-GetServiceUnitRequest"></a>

### GetServiceUnitRequest
GetServiceUnit — fetch a ServiceUnit by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the ServiceUnit CR to fetch. |






<a name="blanketops-environments-v1alpha1-GetServiceUnitResponse"></a>

### GetServiceUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_unit | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) |  |  |






<a name="blanketops-environments-v1alpha1-ListServiceUnitsRequest"></a>

### ListServiceUnitsRequest
ListServiceUnits — list ServiceUnit CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.ServiceUnitDeploymentPhase](#blanketops-common-v1-ServiceUnitDeploymentPhase) | optional | Filter by current lifecycle phase. |
| type | [blanketops.common.v1.ServiceUnitType](#blanketops-common-v1-ServiceUnitType) | optional | Filter by source type. |
| app_type | [string](#string) | optional | Filter by application category. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListServiceUnitsResponse.next_page_token. Omit for the first request. |






<a name="blanketops-environments-v1alpha1-ListServiceUnitsResponse"></a>

### ListServiceUnitsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_units | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) | repeated | The list of ServiceUnit CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-environments-v1alpha1-PatchServiceUnitRequest"></a>

### PatchServiceUnitRequest
PatchServiceUnit — partial update using JSON merge patch RFC 7396.
Only specified fields are updated — others left unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the ServiceUnit CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;size&#34;:3}} |






<a name="blanketops-environments-v1alpha1-PatchServiceUnitResponse"></a>

### PatchServiceUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_unit | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) |  | The patched ServiceUnit as observed by the controller. |






<a name="blanketops-environments-v1alpha1-RouteReference"></a>

### RouteReference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR. Resolved in the same namespace. |






<a name="blanketops-environments-v1alpha1-ServiceUnit"></a>

### ServiceUnit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [ServiceUnitSpec](#blanketops-environments-v1alpha1-ServiceUnitSpec) |  | Desired state declared by the operator. |
| status | [ServiceUnitStatus](#blanketops-environments-v1alpha1-ServiceUnitStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-environments-v1alpha1-ServiceUnitCondition"></a>

### ServiceUnitCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Condition type. e.g. ImageResolved, WorkloadReady, RouteBound |
| status | [string](#string) |  | Condition status: True, False, or Unknown. |
| reason | [string](#string) |  | Machine-readable reason for the last transition. |
| message | [string](#string) |  | Human-readable detail for the last transition. |
| last_transition_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the condition last changed status. |






<a name="blanketops-environments-v1alpha1-ServiceUnitSpec"></a>

### ServiceUnitSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [blanketops.common.v1.ServiceUnitType](#blanketops-common-v1-ServiceUnitType) |  | How this service unit&#39;s image is sourced. |
| build_ref | [BuildReference](#blanketops-environments-v1alpha1-BuildReference) |  | Reference to the source Build. Required when type is BUILD. |
| image | [string](#string) |  | Fully qualified static image reference. Required when type is STATIC. e.g. docker.io/nkanyezisolutions/api:v1.2.3 |
| container_port | [int32](#int32) |  | Port the workload listens on. |
| size | [int32](#int32) |  | Desired replica count. |
| route | [RouteReference](#blanketops-environments-v1alpha1-RouteReference) |  | External route for this unit. Optional. |
| app_type | [string](#string) |  | Application category hint. e.g. web, worker, job |
| stack_type | [string](#string) |  | Language or stack hint. e.g. nodejs, python, wasm |






<a name="blanketops-environments-v1alpha1-ServiceUnitStatus"></a>

### ServiceUnitStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.ServiceUnitDeploymentPhase](#blanketops-common-v1-ServiceUnitDeploymentPhase) |  | Current lifecycle phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status message. Populated on failure. |
| observed_generation | [int64](#int64) |  | Generation of the spec this status reflects. |
| resolved_image | [string](#string) |  | Final resolved image used at runtime, including digest when known. |
| build_status | [BuildResolutionStatus](#blanketops-environments-v1alpha1-BuildResolutionStatus) |  | Build resolution status. Populated when spec.type is BUILD. |
| runtime | [blanketops.common.v1.ServiceUnitRuntime](#blanketops-common-v1-ServiceUnitRuntime) |  | Runtime backend used to materialize the workload. |
| workload_ref | [WorkloadReference](#blanketops-environments-v1alpha1-WorkloadReference) |  | Reference to the concrete workload created. |
| conditions | [ServiceUnitCondition](#blanketops-environments-v1alpha1-ServiceUnitCondition) | repeated | Lifecycle conditions for UI and automation consumers. |
| last_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the controller last updated this status. |






<a name="blanketops-environments-v1alpha1-UpdateServiceUnitRequest"></a>

### UpdateServiceUnitRequest
UpdateServiceUnit — full replace of the ServiceUnit spec.
Equivalent to kubectl apply — all fields replaced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_unit | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) |  | Full ServiceUnit object including metadata and desired spec. |






<a name="blanketops-environments-v1alpha1-UpdateServiceUnitResponse"></a>

### UpdateServiceUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_unit | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) |  | The updated ServiceUnit as observed by the controller. |






<a name="blanketops-environments-v1alpha1-WatchServiceUnitRequest"></a>

### WatchServiceUnitRequest
WatchServiceUnit — stream phase transitions for a ServiceUnit CR.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the ServiceUnit CR to watch. |






<a name="blanketops-environments-v1alpha1-WatchServiceUnitResponse"></a>

### WatchServiceUnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_unit | [ServiceUnit](#blanketops-environments-v1alpha1-ServiceUnit) |  | Current state of the ServiceUnit at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |






<a name="blanketops-environments-v1alpha1-WorkloadReference"></a>

### WorkloadReference
Reference to the concrete workload materialized at runtime.
Opaque to the contract — identifies any runtime object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | API version of the workload object. e.g. apps/v1, serving.knative.dev/v1 |
| kind | [string](#string) |  | Kind of the workload object. e.g. Deployment, Service |
| name | [string](#string) |  | Name of the workload object. |
| namespace | [string](#string) |  | Namespace of the workload object. |





 

 

 


<a name="blanketops-environments-v1alpha1-ServiceUnitService"></a>

### ServiceUnitService
── CRUD — aligned with BuildService pattern ─────────────────────────────

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateServiceUnit | [CreateServiceUnitRequest](#blanketops-environments-v1alpha1-CreateServiceUnitRequest) | [CreateServiceUnitResponse](#blanketops-environments-v1alpha1-CreateServiceUnitResponse) | Declare a new ServiceUnit. |
| GetServiceUnit | [GetServiceUnitRequest](#blanketops-environments-v1alpha1-GetServiceUnitRequest) | [GetServiceUnitResponse](#blanketops-environments-v1alpha1-GetServiceUnitResponse) | Fetch a ServiceUnit CR by name. |
| UpdateServiceUnit | [UpdateServiceUnitRequest](#blanketops-environments-v1alpha1-UpdateServiceUnitRequest) | [UpdateServiceUnitResponse](#blanketops-environments-v1alpha1-UpdateServiceUnitResponse) | Full replace of a ServiceUnit spec. |
| PatchServiceUnit | [PatchServiceUnitRequest](#blanketops-environments-v1alpha1-PatchServiceUnitRequest) | [PatchServiceUnitResponse](#blanketops-environments-v1alpha1-PatchServiceUnitResponse) | Partial update via JSON merge patch RFC 7396. |
| ListServiceUnits | [ListServiceUnitsRequest](#blanketops-environments-v1alpha1-ListServiceUnitsRequest) | [ListServiceUnitsResponse](#blanketops-environments-v1alpha1-ListServiceUnitsResponse) | List ServiceUnit CRs with optional phase/type/app filter and paging. |
| DeleteServiceUnit | [DeleteServiceUnitRequest](#blanketops-environments-v1alpha1-DeleteServiceUnitRequest) | [DeleteServiceUnitResponse](#blanketops-environments-v1alpha1-DeleteServiceUnitResponse) | Delete a ServiceUnit CR and its materialized workload. |
| WatchServiceUnit | [WatchServiceUnitRequest](#blanketops-environments-v1alpha1-WatchServiceUnitRequest) | [WatchServiceUnitResponse](#blanketops-environments-v1alpha1-WatchServiceUnitResponse) stream | Stream phase transitions for a ServiceUnit CR. Streams indefinitely — service units have no terminal phase; close client-side. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

