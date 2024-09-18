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

// ScimAuthenticationScheme struct for ScimAuthenticationScheme.
type ScimAuthenticationScheme struct {
	Description      *string `json:"description,omitempty"`
	DocumentationUri *string `json:"documentationUri,omitempty"`
	Name             *string `json:"name,omitempty"`
	Primary          *bool   `json:"primary,omitempty"`
	SpecUri          *string `json:"specUri,omitempty"`
	Type             *string `json:"type,omitempty"`
}

// NewScimAuthenticationScheme instantiates a new ScimAuthenticationScheme object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimAuthenticationScheme() *ScimAuthenticationScheme {
	p := ScimAuthenticationScheme{}
	return &p
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScimAuthenticationScheme) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScimAuthenticationScheme) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUri returns the DocumentationUri field value if set, zero value otherwise.
func (o *ScimAuthenticationScheme) GetDocumentationUri() string {
	if o == nil || o.DocumentationUri == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUri
}

// SetDocumentationUri gets a reference to the given string and assigns it to the DocumentationUri field.
func (o *ScimAuthenticationScheme) SetDocumentationUri(v string) {
	o.DocumentationUri = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimAuthenticationScheme) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScimAuthenticationScheme) SetName(v string) {
	o.Name = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *ScimAuthenticationScheme) GetPrimary() bool {
	if o == nil || o.Primary == nil {
		var ret bool
		return ret
	}
	return *o.Primary
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *ScimAuthenticationScheme) SetPrimary(v bool) {
	o.Primary = &v
}

// GetSpecUri returns the SpecUri field value if set, zero value otherwise.
func (o *ScimAuthenticationScheme) GetSpecUri() string {
	if o == nil || o.SpecUri == nil {
		var ret string
		return ret
	}
	return *o.SpecUri
}

// SetSpecUri gets a reference to the given string and assigns it to the SpecUri field.
func (o *ScimAuthenticationScheme) SetSpecUri(v string) {
	o.SpecUri = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScimAuthenticationScheme) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScimAuthenticationScheme) SetType(v string) {
	o.Type = &v
}
