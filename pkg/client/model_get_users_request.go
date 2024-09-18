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

// GetUsersRequest struct for GetUsersRequest.
type GetUsersRequest struct {
	Attributes         *string `json:"attributes,omitempty"`
	ExcludedAttributes *string `json:"excludedAttributes,omitempty"`
	Filter             *string `json:"filter,omitempty"`
}

// NewGetUsersRequest instantiates a new GetUsersRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsersRequest() *GetUsersRequest {
	p := GetUsersRequest{}
	return &p
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GetUsersRequest) GetAttributes() string {
	if o == nil || o.Attributes == nil {
		var ret string
		return ret
	}
	return *o.Attributes
}

// SetAttributes gets a reference to the given string and assigns it to the Attributes field.
func (o *GetUsersRequest) SetAttributes(v string) {
	o.Attributes = &v
}

// GetExcludedAttributes returns the ExcludedAttributes field value if set, zero value otherwise.
func (o *GetUsersRequest) GetExcludedAttributes() string {
	if o == nil || o.ExcludedAttributes == nil {
		var ret string
		return ret
	}
	return *o.ExcludedAttributes
}

// SetExcludedAttributes gets a reference to the given string and assigns it to the ExcludedAttributes field.
func (o *GetUsersRequest) SetExcludedAttributes(v string) {
	o.ExcludedAttributes = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GetUsersRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GetUsersRequest) SetFilter(v string) {
	o.Filter = &v
}
