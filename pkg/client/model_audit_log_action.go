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
// API version: 2023-04-10

package client

import (
	"fmt"
)

// AuditLogAction the model 'AuditLogAction'.
type AuditLogAction string

// List of AuditLogAction.
const (
	AUDITLOGACTION_CREATE_CLUSTER                AuditLogAction = "AUDIT_LOG_ACTION_CREATE_CLUSTER"
	AUDITLOGACTION_DELETE_CLUSTER                AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLUSTER"
	AUDITLOGACTION_INVITE_USER_TO_ORGANIZATION   AuditLogAction = "AUDIT_LOG_ACTION_INVITE_USER_TO_ORGANIZATION"
	AUDITLOGACTION_EDIT_USER_INVITE              AuditLogAction = "AUDIT_LOG_ACTION_EDIT_USER_INVITE"
	AUDITLOGACTION_REVOKE_USER_INVITE            AuditLogAction = "AUDIT_LOG_ACTION_REVOKE_USER_INVITE"
	AUDITLOGACTION_ACCEPT_USER_INVITE            AuditLogAction = "AUDIT_LOG_ACTION_ACCEPT_USER_INVITE"
	AUDITLOGACTION_ASSIGN_USER_ROLE              AuditLogAction = "AUDIT_LOG_ACTION_ASSIGN_USER_ROLE"
	AUDITLOGACTION_DELETE_USER_FROM_ORGANIZATION AuditLogAction = "AUDIT_LOG_ACTION_DELETE_USER_FROM_ORGANIZATION"
	AUDITLOGACTION_CREATE_SERVICE_ACCOUNT        AuditLogAction = "AUDIT_LOG_ACTION_CREATE_SERVICE_ACCOUNT"
	AUDITLOGACTION_UPDATE_SERVICE_ACCOUNT        AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_SERVICE_ACCOUNT"
	AUDITLOGACTION_DELETE_SERVICE_ACCOUNT        AuditLogAction = "AUDIT_LOG_ACTION_DELETE_SERVICE_ACCOUNT"
	AUDITLOGACTION_CREATE_API_KEY                AuditLogAction = "AUDIT_LOG_ACTION_CREATE_API_KEY"
	AUDITLOGACTION_UPDATE_API_KEY                AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_API_KEY"
	AUDITLOGACTION_DELETE_API_KEY                AuditLogAction = "AUDIT_LOG_ACTION_DELETE_API_KEY"
	AUDITLOGACTION_UPDATE_CLUSTER                AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLUSTER"
	AUDITLOGACTION_CREATE_SQL_USER               AuditLogAction = "AUDIT_LOG_ACTION_CREATE_SQL_USER"
	AUDITLOGACTION_CHANGE_SQL_USER_PASSWORD      AuditLogAction = "AUDIT_LOG_ACTION_CHANGE_SQL_USER_PASSWORD"
	AUDITLOGACTION_DELETE_SQL_USER               AuditLogAction = "AUDIT_LOG_ACTION_DELETE_SQL_USER"
	AUDITLOGACTION_ADD_IP_ALLOWLIST              AuditLogAction = "AUDIT_LOG_ACTION_ADD_IP_ALLOWLIST"
	AUDITLOGACTION_EDIT_IP_ALLOWLIST             AuditLogAction = "AUDIT_LOG_ACTION_EDIT_IP_ALLOWLIST"
	AUDITLOGACTION_DELETE_IP_ALLOWLIST           AuditLogAction = "AUDIT_LOG_ACTION_DELETE_IP_ALLOWLIST"
	AUDITLOGACTION_CREATE_VPC_PEERING            AuditLogAction = "AUDIT_LOG_ACTION_CREATE_VPC_PEERING"
	AUDITLOGACTION_DELETE_VPC_PEERING            AuditLogAction = "AUDIT_LOG_ACTION_DELETE_VPC_PEERING"
	AUDITLOGACTION_CREATE_PRIVATE_LINK           AuditLogAction = "AUDIT_LOG_ACTION_CREATE_PRIVATE_LINK"
	AUDITLOGACTION_ACCEPT_PRIVATE_LINK           AuditLogAction = "AUDIT_LOG_ACTION_ACCEPT_PRIVATE_LINK"
	AUDITLOGACTION_REJECT_PRIVATE_LINK           AuditLogAction = "AUDIT_LOG_ACTION_REJECT_PRIVATE_LINK"
	AUDITLOGACTION_USER_LOGIN                    AuditLogAction = "AUDIT_LOG_ACTION_USER_LOGIN"
	AUDITLOGACTION_ADD_USER_TO_ROLE              AuditLogAction = "AUDIT_LOG_ACTION_ADD_USER_TO_ROLE"
	AUDITLOGACTION_REMOVE_USER_FROM_ROLE         AuditLogAction = "AUDIT_LOG_ACTION_REMOVE_USER_FROM_ROLE"
	AUDITLOGACTION_CREATE_USER                   AuditLogAction = "AUDIT_LOG_ACTION_CREATE_USER"
	AUDITLOGACTION_DELETE_USER                   AuditLogAction = "AUDIT_LOG_ACTION_DELETE_USER"
	AUDITLOGACTION_UPDATE_USER                   AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_USER"
	AUDITLOGACTION_CREATE_GROUP                  AuditLogAction = "AUDIT_LOG_ACTION_CREATE_GROUP"
	AUDITLOGACTION_DELETE_GROUP                  AuditLogAction = "AUDIT_LOG_ACTION_DELETE_GROUP"
	AUDITLOGACTION_UPDATE_GROUP                  AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_GROUP"
	AUDITLOGACTION_SET_CLIENT_CA_CERT            AuditLogAction = "AUDIT_LOG_ACTION_SET_CLIENT_CA_CERT"
	AUDITLOGACTION_UPDATE_CLIENT_CA_CERT         AuditLogAction = "AUDIT_LOG_ACTION_UPDATE_CLIENT_CA_CERT"
	AUDITLOGACTION_DELETE_CLIENT_CA_CERT         AuditLogAction = "AUDIT_LOG_ACTION_DELETE_CLIENT_CA_CERT"
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
