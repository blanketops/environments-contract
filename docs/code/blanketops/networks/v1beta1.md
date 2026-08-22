# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [blanketops/networks/v1beta1/route.proto](#blanketops_networks_v1beta1_route-proto)
    - [CreateRouteRequest](#blanketops-networks-v1beta1-CreateRouteRequest)
    - [CreateRouteResponse](#blanketops-networks-v1beta1-CreateRouteResponse)
    - [DeleteRouteRequest](#blanketops-networks-v1beta1-DeleteRouteRequest)
    - [DeleteRouteResponse](#blanketops-networks-v1beta1-DeleteRouteResponse)
    - [GetRouteRequest](#blanketops-networks-v1beta1-GetRouteRequest)
    - [GetRouteResponse](#blanketops-networks-v1beta1-GetRouteResponse)
    - [ListRoutesRequest](#blanketops-networks-v1beta1-ListRoutesRequest)
    - [ListRoutesResponse](#blanketops-networks-v1beta1-ListRoutesResponse)
    - [PatchRouteRequest](#blanketops-networks-v1beta1-PatchRouteRequest)
    - [PatchRouteResponse](#blanketops-networks-v1beta1-PatchRouteResponse)
    - [Route](#blanketops-networks-v1beta1-Route)
    - [RouteCondition](#blanketops-networks-v1beta1-RouteCondition)
    - [RouteSpec](#blanketops-networks-v1beta1-RouteSpec)
    - [RouteStatus](#blanketops-networks-v1beta1-RouteStatus)
    - [UpdateRouteRequest](#blanketops-networks-v1beta1-UpdateRouteRequest)
    - [UpdateRouteResponse](#blanketops-networks-v1beta1-UpdateRouteResponse)
    - [WatchRouteRequest](#blanketops-networks-v1beta1-WatchRouteRequest)
    - [WatchRouteResponse](#blanketops-networks-v1beta1-WatchRouteResponse)
  
    - [RouteService](#blanketops-networks-v1beta1-RouteService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="blanketops_networks_v1beta1_route-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/networks/v1beta1/route.proto



<a name="blanketops-networks-v1beta1-CreateRouteRequest"></a>

### CreateRouteRequest
CreateRoute — declare a new Route intent.
Controller materializes the runtime resource on reconciliation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [RouteSpec](#blanketops-networks-v1beta1-RouteSpec) |  | Desired spec for the new Route. |






<a name="blanketops-networks-v1beta1-CreateRouteResponse"></a>

### CreateRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1beta1-Route) |  | The created Route including generated metadata and initial status. |






<a name="blanketops-networks-v1beta1-DeleteRouteRequest"></a>

### DeleteRouteRequest
DeleteRoute — delete a Route CR.
The controller garbage-collects the materialized runtime resource
and any TLS certificates owned by this route.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to delete. |






<a name="blanketops-networks-v1beta1-DeleteRouteResponse"></a>

### DeleteRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the Route was successfully deleted. |






<a name="blanketops-networks-v1beta1-GetRouteRequest"></a>

### GetRouteRequest
GetRoute — fetch a Route by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to fetch. |






<a name="blanketops-networks-v1beta1-GetRouteResponse"></a>

### GetRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1beta1-Route) |  |  |






<a name="blanketops-networks-v1beta1-ListRoutesRequest"></a>

### ListRoutesRequest
ListRoutes — list Route CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.RoutePhase](#blanketops-common-v1-RoutePhase) | optional | Filter by current lifecycle phase. |
| runtime | [blanketops.common.v1.RouteRuntime](#blanketops-common-v1-RouteRuntime) | optional | Filter by materializing runtime. |
| host | [string](#string) | optional | Filter by host FQDN. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListRoutesResponse.next_page_token. Omit for the first request. |






<a name="blanketops-networks-v1beta1-ListRoutesResponse"></a>

### ListRoutesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routes | [Route](#blanketops-networks-v1beta1-Route) | repeated | The list of Route CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-networks-v1beta1-PatchRouteRequest"></a>

### PatchRouteRequest
PatchRoute — partial update using JSON merge patch RFC 7396.
Only specified fields are updated — others left unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to patch. |
| patch | [string](#string) |  | JSON merge patch document — RFC 7396. e.g. {&#34;spec&#34;:{&#34;enabled&#34;:false}} |






<a name="blanketops-networks-v1beta1-PatchRouteResponse"></a>

### PatchRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1beta1-Route) |  | The patched Route as observed by the controller. |






<a name="blanketops-networks-v1beta1-Route"></a>

### Route



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [RouteSpec](#blanketops-networks-v1beta1-RouteSpec) |  | Desired state declared by the operator. |
| status | [RouteStatus](#blanketops-networks-v1beta1-RouteStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-networks-v1beta1-RouteCondition"></a>

### RouteCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Condition type. e.g. RouteResolved, TLSReady, RuntimeReady |
| status | [string](#string) |  | Condition status: True, False, or Unknown. |
| reason | [string](#string) |  | Machine-readable reason for the last transition. |
| message | [string](#string) |  | Human-readable detail for the last transition. |
| last_transition_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the condition last changed status. |






<a name="blanketops-networks-v1beta1-RouteSpec"></a>

### RouteSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | Fully qualified domain name to serve. e.g. app.dev.blanketops.online |
| enabled | [bool](#bool) |  | Whether the route is active. Disabled routes are removed from the runtime but the CR is retained. |
| path | [string](#string) |  | HTTP path prefix to match. e.g. /, /v1, /api |
| tls_enabled | [bool](#bool) |  | Whether TLS is required for this route. Certificate provisioning is delegated to cert-manager via the platform issuer (DNS01 wildcard or namespaced HTTP01). |
| runtime | [blanketops.common.v1.RouteRuntime](#blanketops-common-v1-RouteRuntime) |  | Runtime responsible for materializing this route. |






<a name="blanketops-networks-v1beta1-RouteStatus"></a>

### RouteStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [blanketops.common.v1.RoutePhase](#blanketops-common-v1-RoutePhase) |  | Current lifecycle phase. Set by the controller. |
| message | [string](#string) |  | Human-readable status message. Populated on failure. |
| resolved_address | [string](#string) |  | Resolved address or endpoint once materialized. e.g. load balancer hostname or serving URL. |
| tls_status | [blanketops.common.v1.RouteTLSStatus](#blanketops-common-v1-RouteTLSStatus) |  | TLS provisioning status. Set by the controller. |
| conditions | [RouteCondition](#blanketops-networks-v1beta1-RouteCondition) | repeated | Conditions for UI and automation consumers. |
| last_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the controller last updated this status. |






<a name="blanketops-networks-v1beta1-UpdateRouteRequest"></a>

### UpdateRouteRequest
UpdateRoute — full replace of the Route spec.
Equivalent to kubectl apply — all fields replaced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1beta1-Route) |  | Full Route object including metadata and desired spec. |






<a name="blanketops-networks-v1beta1-UpdateRouteResponse"></a>

### UpdateRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1beta1-Route) |  | The updated Route as observed by the controller. |






<a name="blanketops-networks-v1beta1-WatchRouteRequest"></a>

### WatchRouteRequest
WatchRoute — stream phase transitions for a Route CR.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the Route CR to watch. |






<a name="blanketops-networks-v1beta1-WatchRouteResponse"></a>

### WatchRouteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [Route](#blanketops-networks-v1beta1-Route) |  | Current state of the Route at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this event. |





 

 

 


<a name="blanketops-networks-v1beta1-RouteService"></a>

### RouteService
── CRUD — aligned with BuildService pattern ─────────────────────────────

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRoute | [CreateRouteRequest](#blanketops-networks-v1beta1-CreateRouteRequest) | [CreateRouteResponse](#blanketops-networks-v1beta1-CreateRouteResponse) | Declare a new Route intent. Controller materializes the runtime resource on reconciliation. |
| GetRoute | [GetRouteRequest](#blanketops-networks-v1beta1-GetRouteRequest) | [GetRouteResponse](#blanketops-networks-v1beta1-GetRouteResponse) | Fetch a Route CR by name. |
| UpdateRoute | [UpdateRouteRequest](#blanketops-networks-v1beta1-UpdateRouteRequest) | [UpdateRouteResponse](#blanketops-networks-v1beta1-UpdateRouteResponse) | Full replace of a Route spec. |
| PatchRoute | [PatchRouteRequest](#blanketops-networks-v1beta1-PatchRouteRequest) | [PatchRouteResponse](#blanketops-networks-v1beta1-PatchRouteResponse) | Partial update via JSON merge patch RFC 7396. |
| ListRoutes | [ListRoutesRequest](#blanketops-networks-v1beta1-ListRoutesRequest) | [ListRoutesResponse](#blanketops-networks-v1beta1-ListRoutesResponse) | List Route CRs with optional phase/runtime/host filter and paging. |
| DeleteRoute | [DeleteRouteRequest](#blanketops-networks-v1beta1-DeleteRouteRequest) | [DeleteRouteResponse](#blanketops-networks-v1beta1-DeleteRouteResponse) | Delete a Route CR and its materialized runtime resources. |
| WatchRoute | [WatchRouteRequest](#blanketops-networks-v1beta1-WatchRouteRequest) | [WatchRouteResponse](#blanketops-networks-v1beta1-WatchRouteResponse) stream | Stream phase transitions for a Route CR. Streams indefinitely — routes have no terminal phase; close client-side. |

 



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

