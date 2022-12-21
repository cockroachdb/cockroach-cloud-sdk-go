// Copyright 2022 The Cockroach Authors
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
// API version: 2022-09-20

package client

import (
	"encoding/json"
)

// ListAllowlistEntriesResponse struct for ListAllowlistEntriesResponse.
type ListAllowlistEntriesResponse struct {
	Allowlist   []AllowlistEntry          `json:"allowlist"`
	Pagination  *KeysetPaginationResponse `json:"pagination,omitempty"`
	Propagating bool                      `json:"propagating"`
}

// NewListAllowlistEntriesResponse instantiates a new ListAllowlistEntriesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllowlistEntriesResponse(allowlist []AllowlistEntry, propagating bool) *ListAllowlistEntriesResponse {
	p := ListAllowlistEntriesResponse{}
	p.Allowlist = allowlist
	p.Propagating = propagating
	return &p
}

// NewListAllowlistEntriesResponseWithDefaults instantiates a new ListAllowlistEntriesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllowlistEntriesResponseWithDefaults() *ListAllowlistEntriesResponse {
	p := ListAllowlistEntriesResponse{}
	return &p
}

// GetAllowlist returns the Allowlist field value.
func (o *ListAllowlistEntriesResponse) GetAllowlist() []AllowlistEntry {
	if o == nil {
		var ret []AllowlistEntry
		return ret
	}

	return o.Allowlist
}

// SetAllowlist sets field value.
func (o *ListAllowlistEntriesResponse) SetAllowlist(v []AllowlistEntry) {
	o.Allowlist = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListAllowlistEntriesResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ListAllowlistEntriesResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}

// GetPropagating returns the Propagating field value.
func (o *ListAllowlistEntriesResponse) GetPropagating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Propagating
}

// SetPropagating sets field value.
func (o *ListAllowlistEntriesResponse) SetPropagating(v bool) {
	o.Propagating = v
}

func (o ListAllowlistEntriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowlist"] = o.Allowlist
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["propagating"] = o.Propagating
	}
	return json.Marshal(toSerialize)
}
