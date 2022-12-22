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

// DedicatedMachineTypeSpecification struct for DedicatedMachineTypeSpecification.
type DedicatedMachineTypeSpecification struct {
	// machine_type is the machine type identifier within the given cloud provider, ex. m5.xlarge, n2-standard-4.
	MachineType *string `json:"machine_type,omitempty"`
	// num_virtual_cpus may be used to automatically select a machine type according to the desired number of vCPUs.
	NumVirtualCpus *int32 `json:"num_virtual_cpus,omitempty"`
}

// NewDedicatedMachineTypeSpecification instantiates a new DedicatedMachineTypeSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedMachineTypeSpecification() *DedicatedMachineTypeSpecification {
	p := DedicatedMachineTypeSpecification{}
	return &p
}

// GetMachineType returns the MachineType field value if set, zero value otherwise.
func (o *DedicatedMachineTypeSpecification) GetMachineType() string {
	if o == nil || o.MachineType == nil {
		var ret string
		return ret
	}
	return *o.MachineType
}

// SetMachineType gets a reference to the given string and assigns it to the MachineType field.
func (o *DedicatedMachineTypeSpecification) SetMachineType(v string) {
	o.MachineType = &v
}

// GetNumVirtualCpus returns the NumVirtualCpus field value if set, zero value otherwise.
func (o *DedicatedMachineTypeSpecification) GetNumVirtualCpus() int32 {
	if o == nil || o.NumVirtualCpus == nil {
		var ret int32
		return ret
	}
	return *o.NumVirtualCpus
}

// SetNumVirtualCpus gets a reference to the given int32 and assigns it to the NumVirtualCpus field.
func (o *DedicatedMachineTypeSpecification) SetNumVirtualCpus(v int32) {
	o.NumVirtualCpus = &v
}

func (o DedicatedMachineTypeSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MachineType != nil {
		toSerialize["machine_type"] = o.MachineType
	}
	if o.NumVirtualCpus != nil {
		toSerialize["num_virtual_cpus"] = o.NumVirtualCpus
	}
	return json.Marshal(toSerialize)
}
