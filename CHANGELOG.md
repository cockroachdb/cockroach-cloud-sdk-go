## 6.0.0

BREAKING CHANGES:

* CreateUser endpoint returns CreateUserResponse instead of ScimUser. The fields within the response remained the same.

* Marked certain fields in SCIM endpoints as required.

NEW FEATURES:

* A Labels field was added to the Cluster and FolderResource models.

* CreateCluster, UpdateCluster, CreateFolder, UpdateFolder endpoints now accept an optional labels field to manage labels assigned to the specified resource.

* Cluster Disruption API is now in available in LIMITED ACCESS.

* Added PatchUser SCIM endpoint.

* The following updates were made to the GetUsers endpoint:
  * Added Count and StartIndex optional parameters to GetUsersRequest.
  * Added ItemsPerPage and StartIndex fields to GetUsersResponse.

## 5.1.1

BREAKING CHANGES:

* The SupportPhysicalClusterReplication field has been removed from
  cluster creation. We are not bumping the major version since this flag
  was never moved out of preview.

NEW FEATURES:

* Added the list endpoint for PCR, which was previously accidentally
  omitted.

## 5.1.0

BREAKING CHANGES:

* The OIDC config API endpoints have been removed. Seeing as these
  endpoints were never moved out of private preview, we are not bumping
  the major version.

NEW FEATURES:

* Added endpoints for Physical Cluster Replication.

## 5.0.0

BREAKING CHANGES:

* NewCluster constructor now includes cidr_range

NEW FEATURES:

* cidr_range for GCP Advanced tier clusters is added.

* support for API authentication via JWT is added.

## 4.1.0

NEW FEATURES:

* audit log actions `CREATE_LICENSE` and `UPDATE_ORGANIZATION_NAME` were added.

* endpoints to manage cluster backup configurations.

## 4.0.0

BREAKING CHANGES:

* Removed IsRegex from JWTIssuerIdentityMapEntry. Regex token identities
  must start with a forward slash ('/') to be detected as regexes.

NEW FEATURES:

* Added CockroachVersion to UpdateClusterSpecification. It is now possible
  to upgrade or roll back a CockroachDB Cloud cluster by specifying its
  version.

* Added VCPU_HOURS to QuantityUnitType enum, used by billing and invoice
  endpoints.

## 3.0.0

BREAKING CHANGES:

* RequestUnitLimit and StorageMibLimit inside UsageLimits became optional and
  have hence been updated to use pointers in the UsageLimits struct. This also
  changed the NewUsageLimits constructor which now has no required arguments.

* Added UpgradeType to Add ServerlessClusterConfig. Technically a breaking
  change to the SDK because new required fields are added to the object
  constructor.

* SpendLimit was removed as an optional value in ServerlessClusterConfig.
  UsageLimits should be used instead.

* PlanType enum now consists of the new plan types: "BASIC", "STANDARD" and
  "ADVANCED".

* Add new fields to the ClusterMajorVersion struct. Technically a breaking
  change to the SDK because new required fields are added to the object
  constructor.

BUG FIXES

* fixed install path examples in README

NEW FEATURES:

* ProvisionedVirtualCpus was added to UsageLimits

* Add GIB and GIB_HOURS to QuantityUnit enum used in billing line items.

## 2.0.1

BUG FIXES:

* Update module name for v2 rev

## 2.0.0

BREAKING CHANGES:

* EnablePrometheusMetricExport no longer accepts a map as the 3rd parameter.
  From now on, an attempt will be made to rev the major version whenever
  known breaking changes to the client are introduced.

NEW FEATURES:

* The private endpoint DNS url has been added to the Cluster resource.

* The LogExport endpoints now support Azure.

* Folders endpoints have moved from LIMITED ACCESS to PREVIEW.

* New audit log actions were added.

* Added new JWT Issuers resource.

DEPRECATIONS:

* ApiOidcConfig resource is deprecated and replaced with JWTIssuers resource.

BUG FIXES:

* various docs and example fixes were brought in from the openapi spec.

* various api documentation template improvements including always showing ctx as a
  path param, fixing the name of the options struct, not showing an options
  struct when it doesn't exist.

## 1.11.0

NEW FEATURES:

* New apis for managing service accounts and api keys.
* Setting DeleteProtection is now supported on Cluster Creation.

## 1.10.0

NEW FEATURES:

* A ServiceName field was added to the PrivateEndpointConnection model.
* A DeleteProtection field was added to the Cluster model. This field can be
  set during cluster creation and updated with the UpdateCluster method.
  Requests to delete a cluster with DeleteProtection enabled will fail.
  Disablingthis field to false will allow the cluster to be deleted.

DEPRECATIONS:

* The Paths field on EgressRule has been deprecated. It will be removed in a
  future version.

## 1.9.0

BREAKING CHANGES:

* DEVELOPER and ADMIN were removed from the list of allowed roles. These roles
  have not been allowed since the release of fine grained access controls and
  now they are formally removed from the generated client.

NEW FEATURES:

* The ListFolders endpoint was added.  ListFolders can be used to fetch all
  folders or folders filtered by the path for an organization.

BUG FIXES:

* SCIM search endpoints are now available using POST.  The previous version
  erroneously exposed them via PUT. See
  (Querying Resources Using HTTP POST)[https://www.rfc-editor.org/rfc/rfc7644.html#section-3.4.3]
  for more info. The PUT based methods are kept for backwards compatibility but
  marked deprecated and will be removed in a future version.

* Updated templates to include the deprecation notice in the service interface
  docs.

## 1.8.0

NEW FEATURES:

* The private endpoint connections API is now cloud neutral.

* Various Audit Log Actions have been added.  See [AuditLogActions](docs/AuditLogAction)
  for more details.

* various documentation fixes since the last version have been included.

BREAKING CHANGES:

* various schema model names have been changed to remove unintentional
  prefixing. For example NewCockroachCloudAddPrivateEndpointTrustedOwnerRequest
  has become NewAddPrivateEndpointTrustedOwnerRequest. In general,
  CockroachCloud has been removed from method names where it was unnecessary.
  Similar changes were made to methods and models prefixed with the word Api.
  For example, ApiDatabase was changed to Database.

## 1.7.0 (September 6, 2023)

NEW FEATURES:

* Enhanced support for API OIDC Config Identity maps

## 1.6.0 (August 28, 2023)

NEW FEATURES:

* Support for Private Endpoint Trusted Owners (Limited access)

## 1.5.0 (August 16, 2023)

NEW FEATURES:

* API OIDC Config operations added (Limited access)

## 1.4.0 (August 9, 2023)

NEW FEATURES:

* Folder operations added (Limited access)
* Log Export now supports omitted channels

BUG FIXES:

* Unmarshalling enum values in server responses no longer results in an error if the value is unrecognized.

BREAKING CHANGES:

* GetConnectionStringResponse now represents Params as a ConnectionStringParameters object instead of a string map.
  Both ConnectionString and Params are now required fields and no longer use pointers.

## 1.3.0 (July 12, 2023)

NEW FEATURES:

* Regions can be changed for Multi-Region Serverless clusters.

## 1.2.0 (May 23, 2023)

NEW FEATURES:

* Service docs are now segmented by endpoint category.

BREAKING CHANGES:

* The following preview SCIM endpoints have been removed:
  * PatchUser
  * PatchGroup

## 1.1.0 (May 11, 2023)

NEW FEATURES:

* Enum Docs now include the description field from the api spec
* Limited Access features now included in the SDK. Access must be requested for
  these features. Currently these include:
  * ListAuditLogs
  * Azure Cloud Provider
* SCIM

## 1.0.1 (April 21, 2023)

BUG FIXES:

* 64-bit integers were incorrectly represented as strings. Changed to int64.

## 1.0.0 (April 10, 2023)

NOTES:

* The Go SDK is now generally available with semantic versioning compatibility promises.
* Now using CC API version 2023-04-10.

BREAKING CHANGES:

* The following enums have been renamed, adding a "Type" suffix:
  * ApiCloudProvider (renamed to CloudProviderType)
  * ApiDatadogSite (renamed to DatadogSiteType)
  * AwsEndpointConnectionStatus
  * ClusterMajorVersionSupportStatus
  * ClusterUpgradeStatus
  * Currency
  * EgressTrafficPolicy
  * LogLevel
  * MetricExportStatus
  * NetworkVisibility
  * NodeStatus
  * OrganizationUserRole
  * OperatingSystem
  * Plan
  * PrivateEndpointServiceStatus
* CockroachCloudSetAwsEndpointConnectionStateRequest was renamed to SetAwsEndpointConnectionStateRequest, and
now takes a SetAwsEndpointConnectionStateType enum, which contains the allowed subset of AwsEndpointConnectionStatusType
values.

NEW FEATURES:

* GetPersonUsersByEmail method added.
* UsageLimits added to create and update Serverless cluster operations.
