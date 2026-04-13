# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added pending deploy branch management to release workflow to ensure automated
  PRs that remain open after an SDK release are retargeted to the latest pending
  deploy branch
- Added release workflow that auto-tags from CHANGELOG and dispatches
  sdk-release events to ccloud-private. Requires a `CCLOUD_PRIVATE_DISPATCH_PAT`
  repository secret (fine-grained PAT with contents:write on
  cockroachdb/ccloud-private) and a corresponding `repository_dispatch`
  workflow in ccloud-private to handle the `sdk-release` event

## [6.10.0] - 2025-11-19

### Added

- Added new fields to `Restore` model: `BackupEndTime`, `ClientErrorCode`,
  `ClientErrorMessage`, `CompletedAt`, `CrdbJobId`, `SourceClusterName`,
  `DestinationClusterName`, `Objects`, `RestoreOpts`
- Added new audit log actions: `START_FAULT_TOLERANCE_DEMO`, `END_FAULT_TOLERANCE_DEMO`

## [6.9.0] - 2025-10-24

### Added

- Added `deferred_until` field to `ClusterVersionDeferral` endpoint.

## [6.8.0] - 2025-10-16

### Added

- Added `CreateBlackoutWindow`, `GetBlackoutWindow`, `UpdateBlackoutWindow`,
  `DeleteBlackoutWindow`, and `ListBlackoutWindows` endpoints.
- Added new options for patch deferral: `DEFERRAL_30_DAYS`, `DEFERRAL_60_DAYS`, `DEFERRAL_90_DAYS`.
- Added new RestoreOpts field, `SkipMissingViews`, for CreateRestore requests.

### Fixed

- Entries in the `objects` field in `CreateRestoreRequest` now serialize with
  lowercase JSON keys (`database`, `schema`, `table`), matching API
  requirements. Older SDK versions that send `Database`, `Schema`, `Table`
  remain incompatible - upgrade to this version to run TABLE and DATABASE
  restores successfully.

## [6.7.0] - 2025-10-01

### Added

- `UpdateEgressPrivateEndpoint` now returns a strongly-typed `EgressPrivateEndpoint` in LIMITED ACCESS.
- `EgressPrivateEndpoint` now contains a `DomainNamesState` type in LIMITED ACCESS.
- `ListInvoicesOptions` now contains an optional `status` filter.

### Removed

- Breaking Change: Removed `GetEgressPrivateEndpointResponse` from
  `GetEgressPrivateEndpoint` in LIMITED ACCESS. The private endpoint is now
  returned as a top-level object in the response.

## [6.6.0] - 2025-09-25

### Added

- Added new audit log actions: `CREATE_PHYSICAL_REPLICATION_STREAM`,
  `FAILOVER_PHYSICAL_REPLICATION_STREAM`, and `CANCEL_PHYSICAL_REPLICATION_STREAM`.
- Added new endpoint: `UpdateEgressPrivateEndpoint` in LIMITED ACCESS.
- Added new fields related to the Bring Your Own Cloud feature in LIMITED
  ACCESS: `CustomerCloudAccount` in `CreateClusterRequest` and `Cluster`.

### Deprecated

- Deprecated the `UpdateEgressPrivateEndpointDomainNames` endpoint. Prefer `UpdateEgressPrivateEndpoint`.

## [6.5.0] - 2025-09-08

### Added

- Added new endpoint: `GetEgressPrivateEndpoint` in LIMITED ACCESS.

## [6.4.1] - 2025-08-13

### Fixed

- Add in a few files that were accidentally omitted

## [6.4.0] - 2025-08-12

### Added

- Added new endpoints: `ListBackups`, `CreateRestore`, `GetRestore`,
  `ListRestores`, `CreateEgressPrivateEndpoint`,
  `DeleteEgressPrivateEndpoint`, `ListEgressPrivateEndpoints`,
  `UpdateEgressPrivateEndpointDomainNames` in LIMITED ACCESS.
- Added new CMEK Key Type: `AZURE_KEY_VAULT`
- Added audit log actions: `CREATE_EGRESS_PRIVATE_ENDPOINT`,
  `DELETE_EGRESS_PRIVATE_ENDPOINT`, and `UPDATE_EGRESS_PRIVATE_ENDPOINT`.
- Added `Count` and `StartIndex` fields to GetGroupsRequest and GetUsersRequest.

### Changed

- Breaking Change: The SupportPhysicalClusterReplication field in
  DedicatedClusterCreateSpecification has been renamed to SupportsClusterVirtualization.

## [6.3.0] - 2025-07-10

### Added

- The Cluster object now contains the AzureClusterIdentityClientId field for Azure clusters.
- The SupportPhysicalClusterReplication field has been added to DedicatedClusterCreateSpecification.
- PhysicalReplicationStreams now have a CanceledAt field, and have a CANCELED status.
- The Invoice object now contains the Status field.
- AuditLogActions now include some actions related to plan migrations.

### Changed

- Breaking Change: The data type of the Payload field within AuditLogEntry has
  been changed from a map to a string.

### Removed

- Breaking Change: The ReplicationStream endpoints are now fully removed, after
  having previously been marked as deprecated. The PhysicalReplicationStream endpoints still exist.

## [6.2.0] - 2025-06-11

### Added

- Add CreatePhysicalReplicationStream, GetPhysicalReplicationStream,
  ListPhysicalReplicationStreams, UpdatePhysicalReplicationStream endpoints
- Add PatchGroup SCIM endpoint

### Deprecated

- Deprecate CreateReplicationStream, GetReplicationStream,
  ListReplicationStreams, UpdateReplicationStream endpoints

## [6.1.0] - 2025-04-25

### Added

- Add external_id support for log and metric export
  - Add `external_id` field to `CloudWatchMetricExportInfo` and `EnableCloudWatchMetricExportRequest`
  - Add `aws_external_id` field to `EnableLogExportRequest` and `LogExportClusterSpecification`
- Add `CREATE_LICENSES` to `AuditLogAction` enum

## [6.0.0] - 2025-04-18

### Added

- A Labels field was added to the Cluster and FolderResource models.
- CreateCluster, UpdateCluster, CreateFolder, UpdateFolder endpoints now accept
  an optional labels field to manage labels assigned to the specified resource.
- Cluster Disruption API is now in available in LIMITED ACCESS.
- Added PatchUser SCIM endpoint.
- The following updates were made to the GetUsers endpoint:
  - Added Count and StartIndex optional parameters to GetUsersRequest.
  - Added ItemsPerPage and StartIndex fields to GetUsersResponse.

### Changed

- Breaking Change: CreateUser endpoint returns CreateUserResponse instead of
  ScimUser. The fields within the response remained the same.
- Breaking Change: Marked certain fields in SCIM endpoints as required.

## [5.1.1] - 2025-02-18

### Added

- Added the list endpoint for PCR, which was previously accidentally omitted.

### Removed

- Breaking Change: The SupportPhysicalClusterReplication field has been removed
  from cluster creation. We are not bumping the major version since this flag
  was never moved out of preview.

## [5.1.0] - 2025-01-27

### Added

- Added endpoints for Physical Cluster Replication.

### Removed

- Breaking Change: The OIDC config API endpoints have been removed. Seeing as
  these endpoints were never moved out of private preview, we are not bumping the major version.

## [5.0.0] - 2024-11-14

### Added

- cidr_range for GCP Advanced tier clusters is added.
- support for API authentication via JWT is added.

### Changed

- Breaking Change: NewCluster constructor now includes cidr_range

## [4.1.0] - 2024-10-25

### Added

- audit log actions `CREATE_LICENSE` and `UPDATE_ORGANIZATION_NAME` were added.
- endpoints to manage cluster backup configurations.

## [4.0.0] - 2024-10-03

### Added

- Added CockroachVersion to UpdateClusterSpecification. It is now possible to
  upgrade or roll back a CockroachDB Cloud cluster by specifying its version.
- Added VCPU_HOURS to QuantityUnitType enum, used by billing and invoice endpoints.

### Removed

- Breaking Change: Removed IsRegex from JWTIssuerIdentityMapEntry. Regex token
  identities must start with a forward slash ('/') to be detected as regexes.

## [3.0.0] - 2024-09-18

### Added

- ProvisionedVirtualCpus was added to UsageLimits
- Add GIB and GIB_HOURS to QuantityUnit enum used in billing line items.

### Changed

- Breaking Change: RequestUnitLimit and StorageMibLimit inside UsageLimits
  became optional and have hence been updated to use pointers in the UsageLimits
  struct. This also changed the NewUsageLimits constructor which now has no required arguments.
- Breaking Change: Added UpgradeType to Add ServerlessClusterConfig.
  Technically a breaking change to the SDK because new required fields are added to the object constructor.
- Breaking Change: PlanType enum now consists of the new plan types: "BASIC", "STANDARD" and "ADVANCED".
- Breaking Change: Add new fields to the ClusterMajorVersion struct.
  Technically a breaking change to the SDK because new required fields are added to the object constructor.

### Removed

- Breaking Change: SpendLimit was removed as an optional value in
  ServerlessClusterConfig. UsageLimits should be used instead.

### Fixed

- fixed install path examples in README

## [2.0.1] - 2024-08-29

### Fixed

- Update module name for v2 rev

## [2.0.0] - 2024-08-29

### Added

- The private endpoint DNS url has been added to the Cluster resource.
- The LogExport endpoints now support Azure.
- Folders endpoints have moved from LIMITED ACCESS to PREVIEW.
- New audit log actions were added.
- Added new JWT Issuers resource.

### Changed

- Breaking Change: EnablePrometheusMetricExport no longer accepts a map as the
  3rd parameter. From now on, an attempt will be made to rev the major version
  whenever known breaking changes to the client are introduced.

### Deprecated

- ApiOidcConfig resource is deprecated and replaced with JWTIssuers resource.

### Fixed

- various docs and example fixes were brought in from the openapi spec.
- various api documentation template improvements including always showing ctx
  as a path param, fixing the name of the options struct, not showing an options
  struct when it doesn't exist.

## [1.11.0] - 2024-04-25

### Added

- New apis for managing service accounts and api keys.
- Setting DeleteProtection is now supported on Cluster Creation.

## [1.10.0] - 2024-04-19

### Added

- A ServiceName field was added to the PrivateEndpointConnection model.
- A DeleteProtection field was added to the Cluster model. This field can be
  set during cluster creation and updated with the UpdateCluster method.
  Requests to delete a cluster with DeleteProtection enabled will fail.
  Disablingthis field to false will allow the cluster to be deleted.

### Deprecated

- The Paths field on EgressRule has been deprecated. It will be removed in a future version.

## [1.9.0] - 2024-03-20

### Added

- The ListFolders endpoint was added. ListFolders can be used to fetch all
  folders or folders filtered by the path for an organization.

### Removed

- Breaking Change: DEVELOPER and ADMIN were removed from the list of allowed
  roles. These roles have not been allowed since the release of fine grained
  access controls and now they are formally removed from the generated client.

### Fixed

- SCIM search endpoints are now available using POST. The previous version
  erroneously exposed them via PUT. See (Querying Resources Using HTTP
  POST)[https://www.rfc-editor.org/rfc/rfc7644.html#section-3.4.3] for more
  info. The PUT based methods are kept for backwards compatibility but marked
  deprecated and will be removed in a future version.
- Updated templates to include the deprecation notice in the service interface docs.

## [1.8.0] - 2024-03-08

### Added

- The private endpoint connections API is now cloud neutral.
- Various Audit Log Actions have been added. See [AuditLogActions](docs/AuditLogAction) for more details.
- various documentation fixes since the last version have been included.

### Changed

- Breaking Change: various schema model names have been changed to remove
  unintentional prefixing. For example NewCockroachCloudAddPrivateEndpointTrustedOwnerRequest has become
  NewAddPrivateEndpointTrustedOwnerRequest. In general, CockroachCloud has been
  removed from method names where it was unnecessary. Similar changes were made
  to methods and models prefixed with the word Api. For example, ApiDatabase was changed to Database.

## [1.7.0] - 2023-09-06

### Added

- Enhanced support for API OIDC Config Identity maps

## [1.6.0] - 2023-08-28

### Added

- Support for Private Endpoint Trusted Owners (Limited access)

## [1.5.0] - 2023-08-16

### Added

- API OIDC Config operations added (Limited access)

## [1.4.0] - 2023-08-09

### Added

- Folder operations added (Limited access)
- Log Export now supports omitted channels

### Changed

- Breaking Change: GetConnectionStringResponse now represents Params as a
  ConnectionStringParameters object instead of a string map. Both
  ConnectionString and Params are now required fields and no longer use pointers.

### Fixed

- Unmarshalling enum values in server responses no longer results in an error if the value is unrecognized.

## [1.3.0] - 2023-07-12

### Added

- Regions can be changed for Multi-Region Serverless clusters.

## [1.2.0] - 2023-05-23

### Added

- Service docs are now segmented by endpoint category.

### Removed

- Breaking Change: The following preview SCIM endpoints have been removed:
  - PatchUser
  - PatchGroup

## [1.1.0] - 2023-05-11

### Added

- Enum Docs now include the description field from the api spec
- Limited Access features now included in the SDK. Access must be requested for
  these features. Currently these include:
  - ListAuditLogs
  - Azure Cloud Provider
- SCIM

## [1.0.1] - 2023-04-21

### Fixed

- 64-bit integers were incorrectly represented as strings. Changed to int64.

## [1.0.0] - 2023-04-10

### Added

- GetPersonUsersByEmail method added.
- UsageLimits added to create and update Serverless cluster operations.

### Changed

- Breaking Change: The following enums have been renamed, adding a "Type" suffix:
  - ApiCloudProvider (renamed to CloudProviderType)
  - ApiDatadogSite (renamed to DatadogSiteType)
  - AwsEndpointConnectionStatus
  - ClusterMajorVersionSupportStatus
  - ClusterUpgradeStatus
  - Currency
  - EgressTrafficPolicy
  - LogLevel
  - MetricExportStatus
  - NetworkVisibility
  - NodeStatus
  - OrganizationUserRole
  - OperatingSystem
  - Plan
  - PrivateEndpointServiceStatus
- Breaking Change: CockroachCloudSetAwsEndpointConnectionStateRequest was
  renamed to SetAwsEndpointConnectionStateRequest, and now takes a
  SetAwsEndpointConnectionStateType enum, which contains the allowed subset of
  AwsEndpointConnectionStatusType values.

### Notes

- The Go SDK is now generally available with semantic versioning compatibility promises.
- Now using CC API version 2023-04-10.
