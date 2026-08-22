# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [blanketops/networks/v1alpha1/domain.proto](#blanketops_networks_v1alpha1_domain-proto)
    - [Domain](#blanketops-networks-v1alpha1-Domain)
    - [DomainCondition](#blanketops-networks-v1alpha1-DomainCondition)
    - [DomainMTLS](#blanketops-networks-v1alpha1-DomainMTLS)
    - [DomainObjectRef](#blanketops-networks-v1alpha1-DomainObjectRef)
    - [DomainRouteRef](#blanketops-networks-v1alpha1-DomainRouteRef)
    - [DomainSpec](#blanketops-networks-v1alpha1-DomainSpec)
    - [DomainStatus](#blanketops-networks-v1alpha1-DomainStatus)
  
- [blanketops/networks/v1alpha1/route.proto](#blanketops_networks_v1alpha1_route-proto)
    - [CreateRouteRequest](#blanketops-networks-v1alpha1-CreateRouteRequest)
    - [CreateRouteResponse](#blanketops-networks-v1alpha1-CreateRouteResponse)
    - [DeleteRouteRequest](#blanketops-networks-v1alpha1-DeleteRouteRequest)
    - [DeleteRouteResponse](#blanketops-networks-v1alpha1-DeleteRouteResponse)
    - [GetRouteRequest](#blanketops-networks-v1alpha1-GetRouteRequest)
    - [GetRouteResponse](#blanketops-networks-v1alpha1-GetRouteResponse)
    - [ListRoutesRequest](#blanketops-networks-v1alpha1-ListRoutesRequest)
    - [ListRoutesResponse](#blanketops-networks-v1alpha1-ListRoutesResponse)
    - [PatchRouteRequest](#blanketops-networks-v1alpha1-PatchRouteRequest)
    - [PatchRouteResponse](#blanketops-networks-v1alpha1-PatchRouteResponse)
    - [Route](#blanketops-networks-v1alpha1-Route)
    - [RouteCondition](#blanketops-networks-v1alpha1-RouteCondition)
    - [RouteSpec](#blanketops-networks-v1alpha1-RouteSpec)
    - [RouteStatus](#blanketops-networks-v1alpha1-RouteStatus)
    - [ServiceUnitRef](#blanketops-networks-v1alpha1-ServiceUnitRef)
    - [UpdateRouteRequest](#blanketops-networks-v1alpha1-UpdateRouteRequest)
    - [UpdateRouteResponse](#blanketops-networks-v1alpha1-UpdateRouteResponse)
    - [WatchRouteRequest](#blanketops-networks-v1alpha1-WatchRouteRequest)
    - [WatchRouteResponse](#blanketops-networks-v1alpha1-WatchRouteResponse)
  
    - [RouteService](#blanketops-networks-v1alpha1-RouteService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="blanketops_networks_v1alpha1_domain-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/networks/v1alpha1/domain.proto



<a name="blanketops-networks-v1alpha1-Domain"></a>

### Domain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [DomainSpec](#blanketops-networks-v1alpha1-DomainSpec) |  | Desired state declared by the operator. |
| status | [DomainStatus](#blanketops-networks-v1alpha1-DomainStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-networks-v1alpha1-DomainCondition"></a>

### DomainCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Condition type. e.g. DomainResolved, CertIssued, MTLSReady, DomainMappingReady |
| status | [string](#string) |  | Condition status: True, False, or Unknown. |
| reason | [string](#string) |  | Machine-readable reason for the last transition. |






<a name="blanketops-networks-v1alpha1-DomainMTLS"></a>

### DomainMTLS
DomainMTLS configures inter-service mTLS for this domain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enforced | [bool](#bool) |  | Enforced controls whether blanketops-proxy &#43; blanketK mTLS is active. When true, the platform wires sidecar identity automatically. |






<a name="blanketops-networks-v1alpha1-DomainObjectRef"></a>

### DomainObjectRef
DomainObjectRef is a namespaced reference to a Kubernetes resource
materialized by the Domain controller (Certificate, DomainMapping).
Set in status once the resource is created — never in spec.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the referenced resource. |
| namespace | [string](#string) |  | Namespace of the referenced resource. For DomainMapping: same namespace as this Domain. For Certificate: same namespace as this Domain. |






<a name="blanketops-networks-v1alpha1-DomainRouteRef"></a>

### DomainRouteRef
DomainRouteRef identifies the Route CR that owns this Domain.
Domain and Route must be in the same namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name is the Route CR name in the same namespace as this Domain. |






<a name="blanketops-networks-v1alpha1-DomainSpec"></a>

### DomainSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | Fully qualified domain name this assignment covers. platform strategy: must match the platform wildcard pattern (e.g. *.dev.domain.co.za) custom strategy: client-owned zone, any FQDN. |
| route_ref | [DomainRouteRef](#blanketops-networks-v1alpha1-DomainRouteRef) |  | RouteRef is the Route in the same namespace that owns this Domain. Route owns the workload binding (runtime, path, enabled). Domain is cascade-deleted when the owning Route is deleted. |
| tls_strategy | [blanketops.common.v1.DomainTLSStrategy](#blanketops-common-v1-DomainTLSStrategy) |  | TlsStrategy controls which cert provisioning path the controller takes. platform: platform wildcard cert already covers this host — DomainClaim &#43; DomainMapping only. custom: client-owned zone — Issuer &#43; DomainClaim &#43; DomainMapping &#43; Certificate. |
| mtls | [DomainMTLS](#blanketops-networks-v1alpha1-DomainMTLS) |  | Mtls configures inter-service mTLS for this domain. When enforced, blanketops-proxy sidecars are injected and blanketK issues identities. Absent or not enforced means plain TLS only. |
| renew_before | [string](#string) |  | RenewBefore is the cert renewal window (e.g. &#34;720h&#34;). Applies to custom strategy only — platform wildcard certs use the platform schedule. Empty string means cert-manager default applies. |






<a name="blanketops-networks-v1alpha1-DomainStatus"></a>

### DomainStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.DomainPhase](#blanketops-common-v1-DomainPhase) |  | Current lifecycle phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status message. Populated on failure or provisioning. |
| cert_issued | [bool](#bool) |  | CertIssued is true once a valid TLS certificate has been issued for this host. Always true for platform strategy once DomainMapping is active. Tracks cert-manager issuance for custom strategy. |
| tls_status | [blanketops.common.v1.DomainTLSStatus](#blanketops-common-v1-DomainTLSStatus) |  | TlsStatus is the current TLS provisioning state. |
| conditions | [DomainCondition](#blanketops-networks-v1alpha1-DomainCondition) | repeated | Conditions for UI and automation consumers. |
| last_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | LastUpdatedAt is when the controller last wrote to this status. |
| domain_ready | [bool](#bool) |  | DomainReady is true when both DomainClaim and DomainMapping are reconciled and active. Mirrors the DomainMappingReady condition as a scalar for fast consumer reads. |
| certificate_ref | [DomainObjectRef](#blanketops-networks-v1alpha1-DomainObjectRef) |  | CertificateRef is the Certificate resource emitted by the controller. Set only when tls_strategy is custom — empty for platform strategy. The Route mediator reads this ref to gate DomainMapping dispatch on TLS secret existence without deriving the secret name externally. |
| domain_mapping_ref | [DomainObjectRef](#blanketops-networks-v1alpha1-DomainObjectRef) |  | DomainMappingRef is the Knative DomainMapping resource emitted by the controller. Set once materialized, cleared on deletion. |





 

 

 

 



<a name="blanketops_networks_v1alpha1_route-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/networks/v1alpha1/route.proto



<a name="blanketops-networks-v1alpha1-CreateRouteRequest"></a>

### CreateRouteRequest
CreateRoute — declare a new Route intent.
Controller materializes the runtime resource on reconciliation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [RouteSpec](#blanketops-networks-v1alpha1-RouteSpec) |  | Desired spec for the new Route. |






<a name="blanketops-networks-v1alpha1-CreateRouteResponse"></a>

### CreateRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1alpha1-Route) |  | The created Route including generated metadata and initial status. |






<a name="blanketops-networks-v1alpha1-DeleteRouteRequest"></a>

### DeleteRouteRequest
DeleteRoute — delete a Route CR.
The controller garbage-collects the materialized runtime resource
and any TLS certificates owned by this route.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to delete. |






<a name="blanketops-networks-v1alpha1-DeleteRouteResponse"></a>

### DeleteRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the Route was successfully deleted. |






<a name="blanketops-networks-v1alpha1-GetRouteRequest"></a>

### GetRouteRequest
GetRoute — fetch a Route by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to fetch. |






<a name="blanketops-networks-v1alpha1-GetRouteResponse"></a>

### GetRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1alpha1-Route) |  |  |






<a name="blanketops-networks-v1alpha1-ListRoutesRequest"></a>

### ListRoutesRequest
ListRoutes — list Route CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.RoutePhase](#blanketops-common-v1-RoutePhase) | optional | Filter by current lifecycle phase. |
| runtime | [blanketops.common.v1.RouteRuntime](#blanketops-common-v1-RouteRuntime) | optional | Filter by materializing runtime. |
| host | [string](#string) | optional | Filter by host FQDN. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListRoutesResponse.next_page_token. Omit for the first request. |






<a name="blanketops-networks-v1alpha1-ListRoutesResponse"></a>

### ListRoutesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routes | [Route](#blanketops-networks-v1alpha1-Route) | repeated | The list of Route CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-networks-v1alpha1-PatchRouteRequest"></a>

### PatchRouteRequest
PatchRoute — partial update using JSON merge patch RFC 7396.
Only specified fields are updated — others left unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;enabled&#34;:false}} |






<a name="blanketops-networks-v1alpha1-PatchRouteResponse"></a>

### PatchRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1alpha1-Route) |  | The patched Route as observed by the controller. |






<a name="blanketops-networks-v1alpha1-Route"></a>

### Route
Resource


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [RouteSpec](#blanketops-networks-v1alpha1-RouteSpec) |  | Desired state declared by the operator. |
| status | [RouteStatus](#blanketops-networks-v1alpha1-RouteStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-networks-v1alpha1-RouteCondition"></a>

### RouteCondition
Conditions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Condition type. e.g. RouteResolved, TLSReady, RuntimeReady |
| status | [string](#string) |  | Condition status: True, False, or Unknown. |
| reason | [string](#string) |  | Machine-readable reason for the last transition. |
| message | [string](#string) |  | Human-readable detail for the last transition. |
| last_transition_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the condition last changed status. |






<a name="blanketops-networks-v1alpha1-RouteSpec"></a>

### RouteSpec
Spec (intent)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | Fully qualified domain name to serve. e.g. app.dev.blanketops.online |
| enabled | [bool](#bool) |  | Whether the route is active. Disabled routes are removed from the runtime but the CR is retained. |
| path | [string](#string) |  | HTTP path prefix to match. e.g. /, /v1, /api |
| tls_enabled | [bool](#bool) |  | Whether TLS is required for this route. Certificate provisioning is delegated to cert-manager via the platform issuer (DNS01 wildcard or namespaced HTTP01). |
| runtime | [blanketops.common.v1.RouteRuntime](#blanketops-common-v1-RouteRuntime) |  | Runtime responsible for materializing this route. |
| service_unit_ref | [ServiceUnitRef](#blanketops-networks-v1alpha1-ServiceUnitRef) |  | Required. The ServiceUnit this route exposes. Controller derives the ksvc/K8s service name by convention: ksvc name == service_unit_ref.name No ServiceUnit status lookup required — convention is enforced by the ServiceUnit controller at materialisation time. |






<a name="blanketops-networks-v1alpha1-RouteStatus"></a>

### RouteStatus
Status (observed state)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.RoutePhase](#blanketops-common-v1-RoutePhase) |  | Current lifecycle phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status message. Populated on failure. |
| resolved_address | [string](#string) |  | Resolved address or endpoint once materialized. e.g. load balancer hostname or serving URL. |
| tls_status | [blanketops.common.v1.RouteTLSStatus](#blanketops-common-v1-RouteTLSStatus) |  | TLS provisioning status. Set by the controller. |
| conditions | [RouteCondition](#blanketops-networks-v1alpha1-RouteCondition) | repeated | Conditions for UI and automation consumers. |
| last_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the controller last updated this status. |






<a name="blanketops-networks-v1alpha1-ServiceUnitRef"></a>

### ServiceUnitRef
ServiceUnitRef

Reference to a ServiceUnit CR in the same namespace as this Route.
The controller resolves the materialised ksvc or K8s service by convention:
  ksvc name == service_unit_ref.name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the ServiceUnit CR in this namespace. |






<a name="blanketops-networks-v1alpha1-UpdateRouteRequest"></a>

### UpdateRouteRequest
UpdateRoute — full replace of the Route spec.
Equivalent to kubectl apply — all fields replaced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1alpha1-Route) |  | Full Route object including metadata and desired spec. |






<a name="blanketops-networks-v1alpha1-UpdateRouteResponse"></a>

### UpdateRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1alpha1-Route) |  | The updated Route as observed by the controller. |






<a name="blanketops-networks-v1alpha1-WatchRouteRequest"></a>

### WatchRouteRequest
WatchRoute — stream phase transitions for a Route CR.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to watch. |






<a name="blanketops-networks-v1alpha1-WatchRouteResponse"></a>

### WatchRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1alpha1-Route) |  | Current state of the Route at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |





 

 

 


<a name="blanketops-networks-v1alpha1-RouteService"></a>

### RouteService
Service

── CRUD — aligned with BuildService pattern ─────────────────────────────

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRoute | [CreateRouteRequest](#blanketops-networks-v1alpha1-CreateRouteRequest) | [CreateRouteResponse](#blanketops-networks-v1alpha1-CreateRouteResponse) | Declare a new Route intent. Controller materializes the runtime resource on reconciliation. |
| GetRoute | [GetRouteRequest](#blanketops-networks-v1alpha1-GetRouteRequest) | [GetRouteResponse](#blanketops-networks-v1alpha1-GetRouteResponse) | Fetch a Route CR by name. |
| UpdateRoute | [UpdateRouteRequest](#blanketops-networks-v1alpha1-UpdateRouteRequest) | [UpdateRouteResponse](#blanketops-networks-v1alpha1-UpdateRouteResponse) | Full replace of a Route spec. |
| PatchRoute | [PatchRouteRequest](#blanketops-networks-v1alpha1-PatchRouteRequest) | [PatchRouteResponse](#blanketops-networks-v1alpha1-PatchRouteResponse) | Partial update via JSON merge patch RFC 7396. |
| ListRoutes | [ListRoutesRequest](#blanketops-networks-v1alpha1-ListRoutesRequest) | [ListRoutesResponse](#blanketops-networks-v1alpha1-ListRoutesResponse) | List Route CRs with optional phase/runtime/host filter and paging. |
| DeleteRoute | [DeleteRouteRequest](#blanketops-networks-v1alpha1-DeleteRouteRequest) | [DeleteRouteResponse](#blanketops-networks-v1alpha1-DeleteRouteResponse) | Delete a Route CR and its materialized runtime resources. |
| WatchRoute | [WatchRouteRequest](#blanketops-networks-v1alpha1-WatchRouteRequest) | [WatchRouteResponse](#blanketops-networks-v1alpha1-WatchRouteResponse) stream | Stream phase transitions for a Route CR. Streams indefinitely — routes have no terminal phase; close client-side. |

 



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

