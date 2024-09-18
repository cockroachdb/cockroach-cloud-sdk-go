// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2024-09-16

package client

import (
	"fmt"
)

// AuditLogAction the model 'AuditLogAction'.
type AuditLogAction string

// List of AuditLogAction.
const (
	AUDITLOGACTION_CREATE_CLUSTER                          AuditLogAction = "AUDIT_LOG_ACTION_CREATE_CLUSTER"
	AUDITLOGACTION_DELETE_CLUSTER                          AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLUSTER"
	AUDITLOGACTION_INVITE_USER_TO_ORGANIZATION             AuditLogAction = "AUDIT_LOG_ACTION_INVITE_USER_TO_ORGANIZATION"
	AUDITLOGACTION_EDIT_USER_INVITE                        AuditLogAction = "AUDIT_LOG_ACTION_EDIT_USER_INVITE"
	AUDITLOGACTION_REVOKE_USER_INVITE                      AuditLogAction = "AUDIT_LOG_ACTION_REVOKE_USER_INVITE"
	AUDITLOGACTION_ACCEPT_USER_INVITE                      AuditLogAction = "AUDIT_LOG_ACTION_ACCEPT_USER_INVITE"
	AUDITLOGACTION_ASSIGN_USER_ROLE                        AuditLogAction = "AUDIT_LOG_ACTION_ASSIGN_USER_ROLE"
	AUDITLOGACTION_DELETE_USER_FROM_ORGANIZATION           AuditLogAction = "AUDIT_LOG_ACTION_DELETE_USER_FROM_ORGANIZATION"
	AUDITLOGACTION_CREATE_SERVICE_ACCOUNT                  AuditLogAction = "AUDIT_LOG_ACTION_CREATE_SERVICE_ACCOUNT"
	AUDITLOGACTION_UPDATE_SERVICE_ACCOUNT                  AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_SERVICE_ACCOUNT"
	AUDITLOGACTION_DELETE_SERVICE_ACCOUNT                  AuditLogAction = "AUDIT_LOG_ACTION_DELETE_SERVICE_ACCOUNT"
	AUDITLOGACTION_CREATE_API_KEY                          AuditLogAction = "AUDIT_LOG_ACTION_CREATE_API_KEY"
	AUDITLOGACTION_UPDATE_API_KEY                          AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_API_KEY"
	AUDITLOGACTION_DELETE_API_KEY                          AuditLogAction = "AUDIT_LOG_ACTION_DELETE_API_KEY"
	AUDITLOGACTION_UPDATE_CLUSTER                          AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLUSTER"
	AUDITLOGACTION_CREATE_SQL_USER                         AuditLogAction = "AUDIT_LOG_ACTION_CREATE_SQL_USER"
	AUDITLOGACTION_CHANGE_SQL_USER_PASSWORD                AuditLogAction = "AUDIT_LOG_ACTION_CHANGE_SQL_USER_PASSWORD"
	AUDITLOGACTION_DELETE_SQL_USER                         AuditLogAction = "AUDIT_LOG_ACTION_DELETE_SQL_USER"
	AUDITLOGACTION_ADD_IP_ALLOWLIST                        AuditLogAction = "AUDIT_LOG_ACTION_ADD_IP_ALLOWLIST"
	AUDITLOGACTION_EDIT_IP_ALLOWLIST                       AuditLogAction = "AUDIT_LOG_ACTION_EDIT_IP_ALLOWLIST"
	AUDITLOGACTION_DELETE_IP_ALLOWLIST                     AuditLogAction = "AUDIT_LOG_ACTION_DELETE_IP_ALLOWLIST"
	AUDITLOGACTION_CREATE_VPC_PEERING                      AuditLogAction = "AUDIT_LOG_ACTION_CREATE_VPC_PEERING"
	AUDITLOGACTION_DELETE_VPC_PEERING                      AuditLogAction = "AUDIT_LOG_ACTION_DELETE_VPC_PEERING"
	AUDITLOGACTION_CREATE_PRIVATE_LINK                     AuditLogAction = "AUDIT_LOG_ACTION_CREATE_PRIVATE_LINK"
	AUDITLOGACTION_ACCEPT_PRIVATE_LINK                     AuditLogAction = "AUDIT_LOG_ACTION_ACCEPT_PRIVATE_LINK"
	AUDITLOGACTION_REJECT_PRIVATE_LINK                     AuditLogAction = "AUDIT_LOG_ACTION_REJECT_PRIVATE_LINK"
	AUDITLOGACTION_USER_LOGIN                              AuditLogAction = "AUDIT_LOG_ACTION_USER_LOGIN"
	AUDITLOGACTION_ADD_USER_TO_ROLE                        AuditLogAction = "AUDIT_LOG_ACTION_ADD_USER_TO_ROLE"
	AUDITLOGACTION_REMOVE_USER_FROM_ROLE                   AuditLogAction = "AUDIT_LOG_ACTION_REMOVE_USER_FROM_ROLE"
	AUDITLOGACTION_CREATE_USER                             AuditLogAction = "AUDIT_LOG_ACTION_CREATE_USER"
	AUDITLOGACTION_DELETE_USER                             AuditLogAction = "AUDIT_LOG_ACTION_DELETE_USER"
	AUDITLOGACTION_UPDATE_USER                             AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_USER"
	AUDITLOGACTION_CREATE_GROUP                            AuditLogAction = "AUDIT_LOG_ACTION_CREATE_GROUP"
	AUDITLOGACTION_DELETE_GROUP                            AuditLogAction = "AUDIT_LOG_ACTION_DELETE_GROUP"
	AUDITLOGACTION_UPDATE_GROUP                            AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_GROUP"
	AUDITLOGACTION_SET_CLIENT_CA_CERT                      AuditLogAction = "AUDIT_LOG_ACTION_SET_CLIENT_CA_CERT"
	AUDITLOGACTION_UPDATE_CLIENT_CA_CERT                   AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLIENT_CA_CERT"
	AUDITLOGACTION_DELETE_CLIENT_CA_CERT                   AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLIENT_CA_CERT"
	AUDITLOGACTION_CREATE_API_OIDC_CONFIG                  AuditLogAction = "AUDIT_LOG_ACTION_CREATE_API_OIDC_CONFIG"
	AUDITLOGACTION_DELETE_API_OIDC_CONFIG                  AuditLogAction = "AUDIT_LOG_ACTION_DELETE_API_OIDC_CONFIG"
	AUDITLOGACTION_UPDATE_API_OIDC_CONFIG                  AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_API_OIDC_CONFIG"
	AUDITLOGACTION_CREATE_FOLDER                           AuditLogAction = "AUDIT_LOG_ACTION_CREATE_FOLDER"
	AUDITLOGACTION_DELETE_FOLDER                           AuditLogAction = "AUDIT_LOG_ACTION_DELETE_FOLDER"
	AUDITLOGACTION_UPDATE_FOLDER                           AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_FOLDER"
	AUDITLOGACTION_ADD_PRIVATE_ENDPOINT_TRUSTED_OWNER      AuditLogAction = "AUDIT_LOG_ACTION_ADD_PRIVATE_ENDPOINT_TRUSTED_OWNER"
	AUDITLOGACTION_REMOVE_PRIVATE_ENDPOINT_TRUSTED_OWNER   AuditLogAction = "AUDIT_LOG_ACTION_REMOVE_PRIVATE_ENDPOINT_TRUSTED_OWNER"
	AUDITLOGACTION_ADD_ALERT_RECIPIENT                     AuditLogAction = "AUDIT_LOG_ACTION_ADD_ALERT_RECIPIENT"
	AUDITLOGACTION_REMOVE_ALERT_RECIPIENT                  AuditLogAction = "AUDIT_LOG_ACTION_REMOVE_ALERT_RECIPIENT"
	AUDITLOGACTION_TOGGLE_ALERTS                           AuditLogAction = "AUDIT_LOG_ACTION_TOGGLE_ALERTS"
	AUDITLOGACTION_TEST_ALERT_EMAIL                        AuditLogAction = "AUDIT_LOG_ACTION_TEST_ALERT_EMAIL"
	AUDITLOGACTION_UPDATE_CMEK                             AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CMEK"
	AUDITLOGACTION_REVOKE_CMEK                             AuditLogAction = "AUDIT_LOG_ACTION_REVOKE_CMEK"
	AUDITLOGACTION_UPDATE_CLUSTER_LOG_EXPORT               AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLUSTER_LOG_EXPORT"
	AUDITLOGACTION_DELETE_CLUSTER_LOG_EXPORT               AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLUSTER_LOG_EXPORT"
	AUDITLOGACTION_UPDATE_CLUSTER_METRIC_EXPORT            AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLUSTER_METRIC_EXPORT"
	AUDITLOGACTION_DELETE_CLUSTER_METRIC_EXPORT            AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLUSTER_METRIC_EXPORT"
	AUDITLOGACTION_RESTORE_CLUSTER                         AuditLogAction = "AUDIT_LOG_ACTION_RESTORE_CLUSTER"
	AUDITLOGACTION_UPDATE_CLUSTER_MAJOR_VERSION            AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLUSTER_MAJOR_VERSION"
	AUDITLOGACTION_ROLLBACK_CLUSTER_MAJOR_VERSION_UPDATE   AuditLogAction = "AUDIT_LOG_ACTION_ROLLBACK_CLUSTER_MAJOR_VERSION_UPDATE"
	AUDITLOGACTION_FINALIZE_CLUSTER_MAJOR_VERSION_UPDATE   AuditLogAction = "AUDIT_LOG_ACTION_FINALIZE_CLUSTER_MAJOR_VERSION_UPDATE"
	AUDITLOGACTION_UPDATE_CLUSTER_VERSION_UPGRADE_DEFERRAL AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLUSTER_VERSION_UPGRADE_DEFERRAL"
	AUDITLOGACTION_SET_CLUSTER_MAINTENANCE_WINDOW          AuditLogAction = "AUDIT_LOG_ACTION_SET_CLUSTER_MAINTENANCE_WINDOW"
	AUDITLOGACTION_DELETE_CLUSTER_MAINTENANCE_WINDOW       AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLUSTER_MAINTENANCE_WINDOW"
	AUDITLOGACTION_SET_EGRESS_TRAFFIC_POLICY               AuditLogAction = "AUDIT_LOG_ACTION_SET_EGRESS_TRAFFIC_POLICY"
	AUDITLOGACTION_ADD_EGRESS_RULE                         AuditLogAction = "AUDIT_LOG_ACTION_ADD_EGRESS_RULE"
	AUDITLOGACTION_EDIT_EGRESS_RULE                        AuditLogAction = "AUDIT_LOG_ACTION_EDIT_EGRESS_RULE"
	AUDITLOGACTION_DELETE_EGRESS_RULE                      AuditLogAction = "AUDIT_LOG_ACTION_DELETE_EGRESS_RULE"
	AUDITLOGACTION_ENABLE_CLOUD_ORG_SSO                    AuditLogAction = "AUDIT_LOG_ACTION_ENABLE_CLOUD_ORG_SSO"
	AUDITLOGACTION_ADD_AUTHENTICATION_METHOD               AuditLogAction = "AUDIT_LOG_ACTION_ADD_AUTHENTICATION_METHOD"
	AUDITLOGACTION_UPDATE_AUTHENTICATION_METHOD            AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_AUTHENTICATION_METHOD"
	AUDITLOGACTION_DELETE_AUTHENTICATION_METHOD            AuditLogAction = "AUDIT_LOG_ACTION_DELETE_AUTHENTICATION_METHOD"
	AUDITLOGACTION_SET_DELETE_PROTECTION                   AuditLogAction = "AUDIT_LOG_ACTION_SET_DELETE_PROTECTION"
	AUDITLOGACTION_MARKETPLACE_CREATE_SUBSCRIPTION         AuditLogAction = "AUDIT_LOG_ACTION_MARKETPLACE_CREATE_SUBSCRIPTION"
	AUDITLOGACTION_MARKETPLACE_CANCEL_SUBSCRIPTION         AuditLogAction = "AUDIT_LOG_ACTION_MARKETPLACE_CANCEL_SUBSCRIPTION"
	AUDITLOGACTION_ADD_JWT_ISSUER                          AuditLogAction = "AUDIT_LOG_ACTION_ADD_JWT_ISSUER"
	AUDITLOGACTION_DELETE_JWT_ISSUER                       AuditLogAction = "AUDIT_LOG_ACTION_DELETE_JWT_ISSUER"
	AUDITLOGACTION_UPDATE_JWT_ISSUER                       AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_JWT_ISSUER"
)

// All allowed values of AuditLogAction enum.
var AllowedAuditLogActionEnumValues = []AuditLogAction{
	"AUDIT_LOG_ACTION_CREATE_CLUSTER",
	"AUDIT_LOG_ACTION_DELETE_CLUSTER",
	"AUDIT_LOG_ACTION_INVITE_USER_TO_ORGANIZATION",
	"AUDIT_LOG_ACTION_EDIT_USER_INVITE",
	"AUDIT_LOG_ACTION_REVOKE_USER_INVITE",
	"AUDIT_LOG_ACTION_ACCEPT_USER_INVITE",
	"AUDIT_LOG_ACTION_ASSIGN_USER_ROLE",
	"AUDIT_LOG_ACTION_DELETE_USER_FROM_ORGANIZATION",
	"AUDIT_LOG_ACTION_CREATE_SERVICE_ACCOUNT",
	"AUDIT_LOG_ACTION_UPDATE_SERVICE_ACCOUNT",
	"AUDIT_LOG_ACTION_DELETE_SERVICE_ACCOUNT",
	"AUDIT_LOG_ACTION_CREATE_API_KEY",
	"AUDIT_LOG_ACTION_UPDATE_API_KEY",
	"AUDIT_LOG_ACTION_DELETE_API_KEY",
	"AUDIT_LOG_ACTION_UPDATE_CLUSTER",
	"AUDIT_LOG_ACTION_CREATE_SQL_USER",
	"AUDIT_LOG_ACTION_CHANGE_SQL_USER_PASSWORD",
	"AUDIT_LOG_ACTION_DELETE_SQL_USER",
	"AUDIT_LOG_ACTION_ADD_IP_ALLOWLIST",
	"AUDIT_LOG_ACTION_EDIT_IP_ALLOWLIST",
	"AUDIT_LOG_ACTION_DELETE_IP_ALLOWLIST",
	"AUDIT_LOG_ACTION_CREATE_VPC_PEERING",
	"AUDIT_LOG_ACTION_DELETE_VPC_PEERING",
	"AUDIT_LOG_ACTION_CREATE_PRIVATE_LINK",
	"AUDIT_LOG_ACTION_ACCEPT_PRIVATE_LINK",
	"AUDIT_LOG_ACTION_REJECT_PRIVATE_LINK",
	"AUDIT_LOG_ACTION_USER_LOGIN",
	"AUDIT_LOG_ACTION_ADD_USER_TO_ROLE",
	"AUDIT_LOG_ACTION_REMOVE_USER_FROM_ROLE",
	"AUDIT_LOG_ACTION_CREATE_USER",
	"AUDIT_LOG_ACTION_DELETE_USER",
	"AUDIT_LOG_ACTION_UPDATE_USER",
	"AUDIT_LOG_ACTION_CREATE_GROUP",
	"AUDIT_LOG_ACTION_DELETE_GROUP",
	"AUDIT_LOG_ACTION_UPDATE_GROUP",
	"AUDIT_LOG_ACTION_SET_CLIENT_CA_CERT",
	"AUDIT_LOG_ACTION_UPDATE_CLIENT_CA_CERT",
	"AUDIT_LOG_ACTION_DELETE_CLIENT_CA_CERT",
	"AUDIT_LOG_ACTION_CREATE_API_OIDC_CONFIG",
	"AUDIT_LOG_ACTION_DELETE_API_OIDC_CONFIG",
	"AUDIT_LOG_ACTION_UPDATE_API_OIDC_CONFIG",
	"AUDIT_LOG_ACTION_CREATE_FOLDER",
	"AUDIT_LOG_ACTION_DELETE_FOLDER",
	"AUDIT_LOG_ACTION_UPDATE_FOLDER",
	"AUDIT_LOG_ACTION_ADD_PRIVATE_ENDPOINT_TRUSTED_OWNER",
	"AUDIT_LOG_ACTION_REMOVE_PRIVATE_ENDPOINT_TRUSTED_OWNER",
	"AUDIT_LOG_ACTION_ADD_ALERT_RECIPIENT",
	"AUDIT_LOG_ACTION_REMOVE_ALERT_RECIPIENT",
	"AUDIT_LOG_ACTION_TOGGLE_ALERTS",
	"AUDIT_LOG_ACTION_TEST_ALERT_EMAIL",
	"AUDIT_LOG_ACTION_UPDATE_CMEK",
	"AUDIT_LOG_ACTION_REVOKE_CMEK",
	"AUDIT_LOG_ACTION_UPDATE_CLUSTER_LOG_EXPORT",
	"AUDIT_LOG_ACTION_DELETE_CLUSTER_LOG_EXPORT",
	"AUDIT_LOG_ACTION_UPDATE_CLUSTER_METRIC_EXPORT",
	"AUDIT_LOG_ACTION_DELETE_CLUSTER_METRIC_EXPORT",
	"AUDIT_LOG_ACTION_RESTORE_CLUSTER",
	"AUDIT_LOG_ACTION_UPDATE_CLUSTER_MAJOR_VERSION",
	"AUDIT_LOG_ACTION_ROLLBACK_CLUSTER_MAJOR_VERSION_UPDATE",
	"AUDIT_LOG_ACTION_FINALIZE_CLUSTER_MAJOR_VERSION_UPDATE",
	"AUDIT_LOG_ACTION_UPDATE_CLUSTER_VERSION_UPGRADE_DEFERRAL",
	"AUDIT_LOG_ACTION_SET_CLUSTER_MAINTENANCE_WINDOW",
	"AUDIT_LOG_ACTION_DELETE_CLUSTER_MAINTENANCE_WINDOW",
	"AUDIT_LOG_ACTION_SET_EGRESS_TRAFFIC_POLICY",
	"AUDIT_LOG_ACTION_ADD_EGRESS_RULE",
	"AUDIT_LOG_ACTION_EDIT_EGRESS_RULE",
	"AUDIT_LOG_ACTION_DELETE_EGRESS_RULE",
	"AUDIT_LOG_ACTION_ENABLE_CLOUD_ORG_SSO",
	"AUDIT_LOG_ACTION_ADD_AUTHENTICATION_METHOD",
	"AUDIT_LOG_ACTION_UPDATE_AUTHENTICATION_METHOD",
	"AUDIT_LOG_ACTION_DELETE_AUTHENTICATION_METHOD",
	"AUDIT_LOG_ACTION_SET_DELETE_PROTECTION",
	"AUDIT_LOG_ACTION_MARKETPLACE_CREATE_SUBSCRIPTION",
	"AUDIT_LOG_ACTION_MARKETPLACE_CANCEL_SUBSCRIPTION",
	"AUDIT_LOG_ACTION_ADD_JWT_ISSUER",
	"AUDIT_LOG_ACTION_DELETE_JWT_ISSUER",
	"AUDIT_LOG_ACTION_UPDATE_JWT_ISSUER",
}

// NewAuditLogActionFromValue returns a pointer to a valid AuditLogAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuditLogActionFromValue(v string) (*AuditLogAction, error) {
	ev := AuditLogAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuditLogAction: valid values are %v", v, AllowedAuditLogActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuditLogAction) IsValid() bool {
	for _, existing := range AllowedAuditLogActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditLogAction value.
func (v AuditLogAction) Ptr() *AuditLogAction {
	return &v
}
