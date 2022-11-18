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

// DedicatedClusterCreateSpecification struct for DedicatedClusterCreateSpecification.
type DedicatedClusterCreateSpecification struct {
	// Region keys should match the cloud provider's zone code. For example, for Oregon, set region_name to \"us-west2\" for GCP and \"us-west-2\" for AWS. Values represent the node count.
	RegionNodes map[string]int32                     `json:"region_nodes"`
	Hardware    DedicatedHardwareCreateSpecification `json:"hardware"`
	// The CockroachDB version for the cluster. The current version is used if omitted.
	CockroachVersion *string `json:"cockroach_version,omitempty"`
}

// NewDedicatedClusterCreateSpecification instantiates a new DedicatedClusterCreateSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedClusterCreateSpecification(regionNodes map[string]int32, hardware DedicatedHardwareCreateSpecification) *DedicatedClusterCreateSpecification {
	p := DedicatedClusterCreateSpecification{}
	p.RegionNodes = regionNodes
	p.Hardware = hardware
	return &p
}

// NewDedicatedClusterCreateSpecificationWithDefaults instantiates a new DedicatedClusterCreateSpecification object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedClusterCreateSpecificationWithDefaults() *DedicatedClusterCreateSpecification {
	p := DedicatedClusterCreateSpecification{}
	return &p
}

// GetRegionNodes returns the RegionNodes field value.
func (o *DedicatedClusterCreateSpecification) GetRegionNodes() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.RegionNodes
}

// SetRegionNodes sets field value.
func (o *DedicatedClusterCreateSpecification) SetRegionNodes(v map[string]int32) {
	o.RegionNodes = v
}

// GetHardware returns the Hardware field value.
func (o *DedicatedClusterCreateSpecification) GetHardware() DedicatedHardwareCreateSpecification {
	if o == nil {
		var ret DedicatedHardwareCreateSpecification
		return ret
	}

	return o.Hardware
}

// SetHardware sets field value.
func (o *DedicatedClusterCreateSpecification) SetHardware(v DedicatedHardwareCreateSpecification) {
	o.Hardware = v
}

// GetCockroachVersion returns the CockroachVersion field value if set, zero value otherwise.
func (o *DedicatedClusterCreateSpecification) GetCockroachVersion() string {
	if o == nil || o.CockroachVersion == nil {
		var ret string
		return ret
	}
	return *o.CockroachVersion
}

// SetCockroachVersion gets a reference to the given string and assigns it to the CockroachVersion field.
func (o *DedicatedClusterCreateSpecification) SetCockroachVersion(v string) {
	o.CockroachVersion = &v
}

func (o DedicatedClusterCreateSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["region_nodes"] = o.RegionNodes
	}
	if true {
		toSerialize["hardware"] = o.Hardware
	}
	if o.CockroachVersion != nil {
		toSerialize["cockroach_version"] = o.CockroachVersion
	}
	return json.Marshal(toSerialize)
}
