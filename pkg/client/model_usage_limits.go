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

// UsageLimits struct for UsageLimits.
type UsageLimits struct {
	// request_unit_limit is the maximum number of request units that the cluster can consume during the month. If this limit is exceeded, then the cluster is disabled until the limit is increased, or until the beginning of the next month when more free request units are granted. It is an error for this to be zero.
	RequestUnitLimit int64 `json:"request_unit_limit"`
	// storage_mib_limit is the maximum number of Mebibytes of storage that the cluster can have at any time during the month. If this limit is exceeded, then the cluster is throttled; only one SQL connection is allowed at a time, with the expectation that it is used to delete data to reduce storage usage. It is an error for this to be zero.
	StorageMibLimit int64 `json:"storage_mib_limit"`
}

// NewUsageLimits instantiates a new UsageLimits object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimits(requestUnitLimit int64, storageMibLimit int64) *UsageLimits {
	p := UsageLimits{}
	p.RequestUnitLimit = requestUnitLimit
	p.StorageMibLimit = storageMibLimit
	return &p
}

// NewUsageLimitsWithDefaults instantiates a new UsageLimits object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLimitsWithDefaults() *UsageLimits {
	p := UsageLimits{}
	return &p
}

// GetRequestUnitLimit returns the RequestUnitLimit field value.
func (o *UsageLimits) GetRequestUnitLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RequestUnitLimit
}

// SetRequestUnitLimit sets field value.
func (o *UsageLimits) SetRequestUnitLimit(v int64) {
	o.RequestUnitLimit = v
}

// GetStorageMibLimit returns the StorageMibLimit field value.
func (o *UsageLimits) GetStorageMibLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageMibLimit
}

// SetStorageMibLimit sets field value.
func (o *UsageLimits) SetStorageMibLimit(v int64) {
	o.StorageMibLimit = v
}

func (o UsageLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_unit_limit"] = o.RequestUnitLimit
	}
	if true {
		toSerialize["storage_mib_limit"] = o.StorageMibLimit
	}
	return json.Marshal(toSerialize)
}
