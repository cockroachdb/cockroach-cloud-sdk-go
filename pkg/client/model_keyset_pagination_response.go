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
	"encoding/json"
)

// KeysetPaginationResponse struct for KeysetPaginationResponse.
type KeysetPaginationResponse struct {
	NextPage     *string `json:"next_page,omitempty,string"`
	PreviousPage *string `json:"previous_page,omitempty,string"`
}

// NewKeysetPaginationResponse instantiates a new KeysetPaginationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysetPaginationResponse() *KeysetPaginationResponse {
	p := KeysetPaginationResponse{}
	return &p
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetNextPage() string {
	if o == nil || o.NextPage == nil {
		var ret string
		return ret
	}
	return *o.NextPage
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *KeysetPaginationResponse) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *KeysetPaginationResponse) GetPreviousPage() string {
	if o == nil || o.PreviousPage == nil {
		var ret string
		return ret
	}
	return *o.PreviousPage
}

// SetPreviousPage gets a reference to the given string and assigns it to the PreviousPage field.
func (o *KeysetPaginationResponse) SetPreviousPage(v string) {
	o.PreviousPage = &v
}

func (o KeysetPaginationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPage != nil {
		toSerialize["next_page"] = o.NextPage
	}
	if o.PreviousPage != nil {
		toSerialize["previous_page"] = o.PreviousPage
	}
	return json.Marshal(toSerialize)
}
