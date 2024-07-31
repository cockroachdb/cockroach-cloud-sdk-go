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

// UsageLimits struct for UsageLimits.
type UsageLimits struct {
	// provisioned_capacity is the maximum number of request units that the cluster can consume per second. Once this limit is reached, operation latency may increase due to throttling. It is an error for this to be zero.
	ProvisionedCapacity *int64 `json:"provisioned_capacity,omitempty,string"`
	// provisioned_vcpus is the maximum number of vCPUs that the cluster can use. Once this limit is reached, operation latency may increase due to throttling. It is an error for this to be zero.
	ProvisionedVcpus *int64 `json:"provisioned_vcpus,omitempty,string"`
	// request_unit_limit is the maximum number of request units that the cluster can consume during the month. If this limit is exceeded, then the cluster is disabled until the limit is increased, or until the beginning of the next month when more free request units are granted. It is an error for this to be zero.
	RequestUnitLimit *int64 `json:"request_unit_limit,omitempty,string"`
	// storage_mib_limit is the maximum number of Mebibytes of storage that the cluster can have at any time during the month. If this limit is exceeded, then the cluster is throttled; only one SQL connection is allowed at a time, with the expectation that it is used to delete data to reduce storage usage. It is an error for this to be zero.
	StorageMibLimit *int64 `json:"storage_mib_limit,omitempty,string"`
}

// NewUsageLimits instantiates a new UsageLimits object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimits() *UsageLimits {
	p := UsageLimits{}
	return &p
}

// GetProvisionedCapacity returns the ProvisionedCapacity field value if set, zero value otherwise.
func (o *UsageLimits) GetProvisionedCapacity() int64 {
	if o == nil || o.ProvisionedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.ProvisionedCapacity
}

// SetProvisionedCapacity gets a reference to the given int64 and assigns it to the ProvisionedCapacity field.
func (o *UsageLimits) SetProvisionedCapacity(v int64) {
	o.ProvisionedCapacity = &v
}

// GetProvisionedVcpus returns the ProvisionedVcpus field value if set, zero value otherwise.
func (o *UsageLimits) GetProvisionedVcpus() int64 {
	if o == nil || o.ProvisionedVcpus == nil {
		var ret int64
		return ret
	}
	return *o.ProvisionedVcpus
}

// SetProvisionedVcpus gets a reference to the given int64 and assigns it to the ProvisionedVcpus field.
func (o *UsageLimits) SetProvisionedVcpus(v int64) {
	o.ProvisionedVcpus = &v
}

// GetRequestUnitLimit returns the RequestUnitLimit field value if set, zero value otherwise.
func (o *UsageLimits) GetRequestUnitLimit() int64 {
	if o == nil || o.RequestUnitLimit == nil {
		var ret int64
		return ret
	}
	return *o.RequestUnitLimit
}

// SetRequestUnitLimit gets a reference to the given int64 and assigns it to the RequestUnitLimit field.
func (o *UsageLimits) SetRequestUnitLimit(v int64) {
	o.RequestUnitLimit = &v
}

// GetStorageMibLimit returns the StorageMibLimit field value if set, zero value otherwise.
func (o *UsageLimits) GetStorageMibLimit() int64 {
	if o == nil || o.StorageMibLimit == nil {
		var ret int64
		return ret
	}
	return *o.StorageMibLimit
}

// SetStorageMibLimit gets a reference to the given int64 and assigns it to the StorageMibLimit field.
func (o *UsageLimits) SetStorageMibLimit(v int64) {
	o.StorageMibLimit = &v
}
