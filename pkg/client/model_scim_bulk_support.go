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

// ScimBulkSupport struct for ScimBulkSupport.
type ScimBulkSupport struct {
	MaxOperations  *int32 `json:"maxOperations,omitempty"`
	MaxPayloadSize *int32 `json:"maxPayloadSize,omitempty"`
	Supported      *bool  `json:"supported,omitempty"`
}

// NewScimBulkSupport instantiates a new ScimBulkSupport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimBulkSupport() *ScimBulkSupport {
	p := ScimBulkSupport{}
	return &p
}

// GetMaxOperations returns the MaxOperations field value if set, zero value otherwise.
func (o *ScimBulkSupport) GetMaxOperations() int32 {
	if o == nil || o.MaxOperations == nil {
		var ret int32
		return ret
	}
	return *o.MaxOperations
}

// SetMaxOperations gets a reference to the given int32 and assigns it to the MaxOperations field.
func (o *ScimBulkSupport) SetMaxOperations(v int32) {
	o.MaxOperations = &v
}

// GetMaxPayloadSize returns the MaxPayloadSize field value if set, zero value otherwise.
func (o *ScimBulkSupport) GetMaxPayloadSize() int32 {
	if o == nil || o.MaxPayloadSize == nil {
		var ret int32
		return ret
	}
	return *o.MaxPayloadSize
}

// SetMaxPayloadSize gets a reference to the given int32 and assigns it to the MaxPayloadSize field.
func (o *ScimBulkSupport) SetMaxPayloadSize(v int32) {
	o.MaxPayloadSize = &v
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *ScimBulkSupport) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *ScimBulkSupport) SetSupported(v bool) {
	o.Supported = &v
}
