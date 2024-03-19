
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
