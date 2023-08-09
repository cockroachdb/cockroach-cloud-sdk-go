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

// OrganizationUserRoleType  - DEVELOPER: To be deprecated  - ADMIN: To be deprecated  - FOLDER_ADMIN: Limited Access: A folder admin role.  - FOLDER_MOVER: Limited Access: A folder mover role.
type OrganizationUserRoleType string

// List of OrganizationUserRole.Type.
const (
	ORGANIZATIONUSERROLETYPE_DEVELOPER               OrganizationUserRoleType = "DEVELOPER"
	ORGANIZATIONUSERROLETYPE_ADMIN                   OrganizationUserRoleType = "ADMIN"
	ORGANIZATIONUSERROLETYPE_BILLING_COORDINATOR     OrganizationUserRoleType = "BILLING_COORDINATOR"
	ORGANIZATIONUSERROLETYPE_ORG_ADMIN               OrganizationUserRoleType = "ORG_ADMIN"
	ORGANIZATIONUSERROLETYPE_ORG_MEMBER              OrganizationUserRoleType = "ORG_MEMBER"
	ORGANIZATIONUSERROLETYPE_CLUSTER_ADMIN           OrganizationUserRoleType = "CLUSTER_ADMIN"
	ORGANIZATIONUSERROLETYPE_CLUSTER_OPERATOR_WRITER OrganizationUserRoleType = "CLUSTER_OPERATOR_WRITER"
	ORGANIZATIONUSERROLETYPE_CLUSTER_DEVELOPER       OrganizationUserRoleType = "CLUSTER_DEVELOPER"
	ORGANIZATIONUSERROLETYPE_CLUSTER_CREATOR         OrganizationUserRoleType = "CLUSTER_CREATOR"
	ORGANIZATIONUSERROLETYPE_FOLDER_ADMIN            OrganizationUserRoleType = "FOLDER_ADMIN"
	ORGANIZATIONUSERROLETYPE_FOLDER_MOVER            OrganizationUserRoleType = "FOLDER_MOVER"
)

// All allowed values of OrganizationUserRoleType enum.
var AllowedOrganizationUserRoleTypeEnumValues = []OrganizationUserRoleType{
	"DEVELOPER",
	"ADMIN",
	"BILLING_COORDINATOR",
	"ORG_ADMIN",
	"ORG_MEMBER",
	"CLUSTER_ADMIN",
	"CLUSTER_OPERATOR_WRITER",
	"CLUSTER_DEVELOPER",
	"CLUSTER_CREATOR",
	"FOLDER_ADMIN",
	"FOLDER_MOVER",
}

// NewOrganizationUserRoleTypeFromValue returns a pointer to a valid OrganizationUserRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrganizationUserRoleTypeFromValue(v string) (*OrganizationUserRoleType, error) {
	ev := OrganizationUserRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationUserRoleType: valid values are %v", v, AllowedOrganizationUserRoleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrganizationUserRoleType) IsValid() bool {
	for _, existing := range AllowedOrganizationUserRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationUserRole.Type value.
func (v OrganizationUserRoleType) Ptr() *OrganizationUserRoleType {
	return &v
}
