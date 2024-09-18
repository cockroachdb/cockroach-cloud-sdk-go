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

// ScimSchema struct for ScimSchema.
type ScimSchema struct {
	Attributes  *[]ScimSchemaAttribute `json:"attributes,omitempty"`
	Description *string                `json:"description,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Meta        *ScimMetadata          `json:"meta,omitempty"`
	Name        *string                `json:"name,omitempty"`
}

// NewScimSchema instantiates a new ScimSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimSchema() *ScimSchema {
	p := ScimSchema{}
	return &p
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScimSchema) GetAttributes() []ScimSchemaAttribute {
	if o == nil || o.Attributes == nil {
		var ret []ScimSchemaAttribute
		return ret
	}
	return *o.Attributes
}

// SetAttributes gets a reference to the given []ScimSchemaAttribute and assigns it to the Attributes field.
func (o *ScimSchema) SetAttributes(v []ScimSchemaAttribute) {
	o.Attributes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScimSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScimSchema) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScimSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScimSchema) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimSchema) GetMeta() ScimMetadata {
	if o == nil || o.Meta == nil {
		var ret ScimMetadata
		return ret
	}
	return *o.Meta
}

// SetMeta gets a reference to the given ScimMetadata and assigns it to the Meta field.
func (o *ScimSchema) SetMeta(v ScimMetadata) {
	o.Meta = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScimSchema) SetName(v string) {
	o.Name = &v
}
