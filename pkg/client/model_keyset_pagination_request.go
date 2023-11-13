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

import (
	"time"
)

// KeysetPaginationRequest struct for KeysetPaginationRequest.
type KeysetPaginationRequest struct {
	AsOfTime  *time.Time `json:"as_of_time,omitempty"`
	Limit     *int32     `json:"limit,omitempty"`
	Page      *string    `json:"page,omitempty"`
	SortOrder *SortOrder `json:"sort_order,omitempty"`
}

// NewKeysetPaginationRequest instantiates a new KeysetPaginationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysetPaginationRequest() *KeysetPaginationRequest {
	p := KeysetPaginationRequest{}
	return &p
}

// GetAsOfTime returns the AsOfTime field value if set, zero value otherwise.
func (o *KeysetPaginationRequest) GetAsOfTime() time.Time {
	if o == nil || o.AsOfTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AsOfTime
}

// SetAsOfTime gets a reference to the given time.Time and assigns it to the AsOfTime field.
func (o *KeysetPaginationRequest) SetAsOfTime(v time.Time) {
	o.AsOfTime = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *KeysetPaginationRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *KeysetPaginationRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *KeysetPaginationRequest) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *KeysetPaginationRequest) SetPage(v string) {
	o.Page = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *KeysetPaginationRequest) GetSortOrder() SortOrder {
	if o == nil || o.SortOrder == nil {
		var ret SortOrder
		return ret
	}
	return *o.SortOrder
}

// SetSortOrder gets a reference to the given SortOrder and assigns it to the SortOrder field.
func (o *KeysetPaginationRequest) SetSortOrder(v SortOrder) {
	o.SortOrder = &v
}
