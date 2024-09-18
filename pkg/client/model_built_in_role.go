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

// BuiltInRole struct for BuiltInRole.
type BuiltInRole struct {
	Name     OrganizationUserRoleType `json:"name"`
	Resource Resource                 `json:"resource"`
}

// NewBuiltInRole instantiates a new BuiltInRole object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuiltInRole(name OrganizationUserRoleType, resource Resource) *BuiltInRole {
	p := BuiltInRole{}
	p.Name = name
	p.Resource = resource
	return &p
}

// NewBuiltInRoleWithDefaults instantiates a new BuiltInRole object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuiltInRoleWithDefaults() *BuiltInRole {
	p := BuiltInRole{}
	return &p
}

// GetName returns the Name field value.
func (o *BuiltInRole) GetName() OrganizationUserRoleType {
	if o == nil {
		var ret OrganizationUserRoleType
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *BuiltInRole) SetName(v OrganizationUserRoleType) {
	o.Name = v
}

// GetResource returns the Resource field value.
func (o *BuiltInRole) GetResource() Resource {
	if o == nil {
		var ret Resource
		return ret
	}

	return o.Resource
}

// SetResource sets field value.
func (o *BuiltInRole) SetResource(v Resource) {
	o.Resource = v
}
