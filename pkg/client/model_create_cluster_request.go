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

// CreateClusterRequest struct for CreateClusterRequest.
type CreateClusterRequest struct {
	// Name must be 6-20 characters in length and can include numbers, lowercase letters, and dashes (but no leading or trailing dashes).
	Name     string                     `json:"name"`
	Provider ApiCloudProvider           `json:"provider"`
	Spec     CreateClusterSpecification `json:"spec"`
}

// NewCreateClusterRequest instantiates a new CreateClusterRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequest(name string, provider ApiCloudProvider, spec CreateClusterSpecification) *CreateClusterRequest {
	p := CreateClusterRequest{}
	p.Name = name
	p.Provider = provider
	p.Spec = spec
	return &p
}

// NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestWithDefaults() *CreateClusterRequest {
	p := CreateClusterRequest{}
	return &p
}

// GetName returns the Name field value.
func (o *CreateClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *CreateClusterRequest) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value.
func (o *CreateClusterRequest) GetProvider() ApiCloudProvider {
	if o == nil {
		var ret ApiCloudProvider
		return ret
	}

	return o.Provider
}

// SetProvider sets field value.
func (o *CreateClusterRequest) SetProvider(v ApiCloudProvider) {
	o.Provider = v
}

// GetSpec returns the Spec field value.
func (o *CreateClusterRequest) GetSpec() CreateClusterSpecification {
	if o == nil {
		var ret CreateClusterSpecification
		return ret
	}

	return o.Spec
}

// SetSpec sets field value.
func (o *CreateClusterRequest) SetSpec(v CreateClusterSpecification) {
	o.Spec = v
}

func (o CreateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}
