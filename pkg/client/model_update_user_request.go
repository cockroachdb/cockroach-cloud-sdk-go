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
// API version: development

package client

// UpdateUserRequest struct for UpdateUserRequest.
type UpdateUserRequest struct {
	Active      bool         `json:"active"`
	DisplayName *string      `json:"displayName,omitempty"`
	Emails      *[]ScimEmail `json:"emails,omitempty"`
	ExternalId  *string      `json:"externalId,omitempty"`
	Name        *ScimName    `json:"name,omitempty"`
	Schemas     *[]string    `json:"schemas,omitempty"`
	UserName    *string      `json:"userName,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest(active bool) *UpdateUserRequest {
	p := UpdateUserRequest{}
	p.Active = active
	return &p
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	p := UpdateUserRequest{}
	return &p
}

// GetActive returns the Active field value.
func (o *UpdateUserRequest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// SetActive sets field value.
func (o *UpdateUserRequest) SetActive(v bool) {
	o.Active = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetEmails() []ScimEmail {
	if o == nil || o.Emails == nil {
		var ret []ScimEmail
		return ret
	}
	return *o.Emails
}

// SetEmails gets a reference to the given []ScimEmail and assigns it to the Emails field.
func (o *UpdateUserRequest) SetEmails(v []ScimEmail) {
	o.Emails = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UpdateUserRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetName() ScimName {
	if o == nil || o.Name == nil {
		var ret ScimName
		return ret
	}
	return *o.Name
}

// SetName gets a reference to the given ScimName and assigns it to the Name field.
func (o *UpdateUserRequest) SetName(v ScimName) {
	o.Name = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return *o.Schemas
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *UpdateUserRequest) SetSchemas(v []string) {
	o.Schemas = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UpdateUserRequest) SetUserName(v string) {
	o.UserName = &v
}
