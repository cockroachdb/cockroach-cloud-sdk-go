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

// GetAllRolesForUserResponse struct for GetAllRolesForUserResponse.
type GetAllRolesForUserResponse struct {
	GroupRoles *[]BuiltInFromGroups `json:"group_roles,omitempty"`
	Roles      *[]BuiltInRole       `json:"roles,omitempty"`
}

// NewGetAllRolesForUserResponse instantiates a new GetAllRolesForUserResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllRolesForUserResponse() *GetAllRolesForUserResponse {
	p := GetAllRolesForUserResponse{}
	return &p
}

// GetGroupRoles returns the GroupRoles field value if set, zero value otherwise.
func (o *GetAllRolesForUserResponse) GetGroupRoles() []BuiltInFromGroups {
	if o == nil || o.GroupRoles == nil {
		var ret []BuiltInFromGroups
		return ret
	}
	return *o.GroupRoles
}

// SetGroupRoles gets a reference to the given []BuiltInFromGroups and assigns it to the GroupRoles field.
func (o *GetAllRolesForUserResponse) SetGroupRoles(v []BuiltInFromGroups) {
	o.GroupRoles = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GetAllRolesForUserResponse) GetRoles() []BuiltInRole {
	if o == nil || o.Roles == nil {
		var ret []BuiltInRole
		return ret
	}
	return *o.Roles
}

// SetRoles gets a reference to the given []BuiltInRole and assigns it to the Roles field.
func (o *GetAllRolesForUserResponse) SetRoles(v []BuiltInRole) {
	o.Roles = &v
}
