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
