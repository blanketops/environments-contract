# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [blanketops/events/v1alpha1/githubevent.proto](#blanketops_events_v1alpha1_githubevent-proto)
    - [CreateGitHubEventRequest](#blanketops-events-v1alpha1-CreateGitHubEventRequest)
    - [CreateGitHubEventResponse](#blanketops-events-v1alpha1-CreateGitHubEventResponse)
    - [DeleteGitHubEventRequest](#blanketops-events-v1alpha1-DeleteGitHubEventRequest)
    - [DeleteGitHubEventResponse](#blanketops-events-v1alpha1-DeleteGitHubEventResponse)
    - [GetGitHubEventRequest](#blanketops-events-v1alpha1-GetGitHubEventRequest)
    - [GetGitHubEventResponse](#blanketops-events-v1alpha1-GetGitHubEventResponse)
    - [GitHubEvent](#blanketops-events-v1alpha1-GitHubEvent)
    - [GitHubEventSpec](#blanketops-events-v1alpha1-GitHubEventSpec)
    - [GitHubEventStatus](#blanketops-events-v1alpha1-GitHubEventStatus)
    - [ListGitHubEventsRequest](#blanketops-events-v1alpha1-ListGitHubEventsRequest)
    - [ListGitHubEventsResponse](#blanketops-events-v1alpha1-ListGitHubEventsResponse)
    - [WatchGitHubEventRequest](#blanketops-events-v1alpha1-WatchGitHubEventRequest)
    - [WatchGitHubEventResponse](#blanketops-events-v1alpha1-WatchGitHubEventResponse)
  
    - [GitHubEventService](#blanketops-events-v1alpha1-GitHubEventService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="blanketops_events_v1alpha1_githubevent-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## blanketops/events/v1alpha1/githubevent.proto
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


<a name="blanketops-events-v1alpha1-CreateGitHubEventRequest"></a>

### CreateGitHubEventRequest
Requests / Responses
CreateGitHubEvent — record a new event.
Primary path for MANUAL dispatch; webhook deliveries arrive via
platform-generated GitHubPayload, not this RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [GitHubEventSpec](#blanketops-events-v1alpha1-GitHubEventSpec) |  | The event to record. |






<a name="blanketops-events-v1alpha1-CreateGitHubEventResponse"></a>

### CreateGitHubEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| github_event | [GitHubEvent](#blanketops-events-v1alpha1-GitHubEvent) |  | The created GitHubEvent including generated metadata and initial status. |






<a name="blanketops-events-v1alpha1-DeleteGitHubEventRequest"></a>

### DeleteGitHubEventRequest
DeleteGitHubEvent — delete a GitHubEvent CR.
Events are normally garbage-collected by TTL; manual delete is for cleanup.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the GitHubEvent CR to delete. |






<a name="blanketops-events-v1alpha1-DeleteGitHubEventResponse"></a>

### DeleteGitHubEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | True if the GitHubEvent was successfully deleted. |






<a name="blanketops-events-v1alpha1-GetGitHubEventRequest"></a>

### GetGitHubEventRequest
GetGitHubEvent — fetch a GitHubEvent by name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the GitHubEvent CR to fetch. |






<a name="blanketops-events-v1alpha1-GetGitHubEventResponse"></a>

### GetGitHubEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| github_event | [GitHubEvent](#blanketops-events-v1alpha1-GitHubEvent) |  |  |






<a name="blanketops-events-v1alpha1-GitHubEvent"></a>

### GitHubEvent
GitHubEvent

User-facing record of a Git provider event observed by the platform.
Immutable once accepted — the controller writes status, never the spec.
GitHubPayload is the platform-generated counterpart (auto, ephemeral,
auditable); GitHubEvent is the user-facing surface.
Pipeline chain:
  GitRepository → GitHubEvent → Build → SupplyChain → Package → Deployment

Resource


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [blanketops.common.v1.Metadata](#blanketops-common-v1-Metadata) |  | Standard BlanketOps metadata (name, namespace, labels, ownerRef). |
| spec | [GitHubEventSpec](#blanketops-events-v1alpha1-GitHubEventSpec) |  | The observed event. Immutable once accepted. |
| status | [GitHubEventStatus](#blanketops-events-v1alpha1-GitHubEventStatus) |  | Observed state set by the controller. Never manually edited. |






<a name="blanketops-events-v1alpha1-GitHubEventSpec"></a>

### GitHubEventSpec
Spec (observed event)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| repository | [string](#string) |  | Repository the event originated from — owner/name. e.g. ntlaletsi70/blanketops-environments-contract |
| event_type | [blanketops.common.v1.GitHubEventType](#blanketops-common-v1-GitHubEventType) |  | The provider event type. |
| ref | [string](#string) |  | Git ref the event applies to. e.g. refs/heads/main, refs/tags/v0.1.0 |
| commit_sha | [string](#string) |  | Commit SHA at the head of the event. |
| actor | [string](#string) |  | Provider login of the actor who caused the event. |
| event_id | [string](#string) |  | Provider-assigned delivery ID. Used for idempotency and audit. |
| occurred_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the event occurred on the provider. |
| webhook_secret_ref | [string](#string) |  | Name of the Kubernetes Secret containing the webhook HMAC secret used to verify the delivery signature. Sourced from the environment SecretStore via ExternalSecret. |






<a name="blanketops-events-v1alpha1-GitHubEventStatus"></a>

### GitHubEventStatus
Status (observed state)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accepted | [bool](#bool) |  | Whether the event passed signature verification and contract match. |
| triggered | [bool](#bool) |  | Whether the event resulted in a BuildTrigger firing. |
| triggered_ref | [string](#string) |  | Name of the resource the event triggered, if any. e.g. the Build CR patched by the BuildTrigger. |
| reason | [string](#string) |  | Human-readable reason if not accepted or not triggered. |
| processed_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | When the controller finished processing the event. |






<a name="blanketops-events-v1alpha1-ListGitHubEventsRequest"></a>

### ListGitHubEventsRequest
ListGitHubEvents — list GitHubEvent CRs with optional filtering and paging.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| repository | [string](#string) | optional | Filter by source repository — owner/name. |
| event_type | [blanketops.common.v1.GitHubEventType](#blanketops-common-v1-GitHubEventType) | optional | Filter by provider event type. |
| triggered | [bool](#bool) | optional | Filter by whether the event fired a trigger. |
| page_size | [int32](#int32) | optional | Maximum number of results to return. Server may return fewer. |
| page_token | [string](#string) | optional | Token from a previous ListGitHubEventsResponse.next_page_token. Omit for the first request. |






<a name="blanketops-events-v1alpha1-ListGitHubEventsResponse"></a>

### ListGitHubEventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| github_events | [GitHubEvent](#blanketops-events-v1alpha1-GitHubEvent) | repeated | The list of GitHubEvent CRs matching the request filters. |
| next_page_token | [string](#string) | optional | Token to retrieve the next page. Empty if no more results. |






<a name="blanketops-events-v1alpha1-WatchGitHubEventRequest"></a>

### WatchGitHubEventRequest
WatchGitHubEvent — stream state transitions for GitHubEvent CRs.
Delivers an event for every controller reconciliation loop.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of a specific GitHubEvent CR to watch. Mutually exclusive with label_selector. |
| label_selector | [string](#string) |  | Kubernetes label selector to watch a set of events. e.g. environments.blanketops.dev/type=push |






<a name="blanketops-events-v1alpha1-WatchGitHubEventResponse"></a>

### WatchGitHubEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| github_event | [GitHubEvent](#blanketops-events-v1alpha1-GitHubEvent) |  | Current state of the GitHubEvent at this point in the stream. |
| type | [blanketops.common.v1.EventType](#blanketops-common-v1-EventType) |  | The type of change that triggered this stream event. |





 

 

 


<a name="blanketops-events-v1alpha1-GitHubEventService"></a>

### GitHubEventService
Service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateGitHubEvent | [CreateGitHubEventRequest](#blanketops-events-v1alpha1-CreateGitHubEventRequest) | [CreateGitHubEventResponse](#blanketops-events-v1alpha1-CreateGitHubEventResponse) | ── CRD surface — events are immutable, no Update/Patch ───────────────── Record a new event. Primary path for MANUAL dispatch. |
| GetGitHubEvent | [GetGitHubEventRequest](#blanketops-events-v1alpha1-GetGitHubEventRequest) | [GetGitHubEventResponse](#blanketops-events-v1alpha1-GetGitHubEventResponse) | Fetch a GitHubEvent CR by name. |
| ListGitHubEvents | [ListGitHubEventsRequest](#blanketops-events-v1alpha1-ListGitHubEventsRequest) | [ListGitHubEventsResponse](#blanketops-events-v1alpha1-ListGitHubEventsResponse) | List GitHubEvent CRs with optional repository/type/triggered filter and paging. |
| DeleteGitHubEvent | [DeleteGitHubEventRequest](#blanketops-events-v1alpha1-DeleteGitHubEventRequest) | [DeleteGitHubEventResponse](#blanketops-events-v1alpha1-DeleteGitHubEventResponse) | Delete a GitHubEvent CR. Normally handled by TTL garbage collection. |
| WatchGitHubEvent | [WatchGitHubEventRequest](#blanketops-events-v1alpha1-WatchGitHubEventRequest) | [WatchGitHubEventResponse](#blanketops-events-v1alpha1-WatchGitHubEventResponse) stream | ── Streaming ──────────────────────────────────────────────────────────── Stream state transitions for GitHubEvent CRs by name or label selector. |

 



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

