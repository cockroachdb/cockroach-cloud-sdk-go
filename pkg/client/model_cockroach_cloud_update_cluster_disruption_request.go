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

// CockroachCloudUpdateClusterDisruptionRequest struct for CockroachCloudUpdateClusterDisruptionRequest.
type CockroachCloudUpdateClusterDisruptionRequest struct {
	// regional_disruptor_specifications specify how regions are to be disrupted. Any Cluster region that is not specified here will not be disrupted. A cluster region that was previously disrupted but is not listed here will be removed from disruption. To stop all disruptions, set this to an empty list or omit it from the request.
	RegionalDisruptorSpecifications *[]RegionalDisruptorSpecification `json:"regional_disruptor_specifications,omitempty"`
}

// NewCockroachCloudUpdateClusterDisruptionRequest instantiates a new CockroachCloudUpdateClusterDisruptionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCockroachCloudUpdateClusterDisruptionRequest() *CockroachCloudUpdateClusterDisruptionRequest {
	p := CockroachCloudUpdateClusterDisruptionRequest{}
	return &p
}

// GetRegionalDisruptorSpecifications returns the RegionalDisruptorSpecifications field value if set, zero value otherwise.
func (o *CockroachCloudUpdateClusterDisruptionRequest) GetRegionalDisruptorSpecifications() []RegionalDisruptorSpecification {
	if o == nil || o.RegionalDisruptorSpecifications == nil {
		var ret []RegionalDisruptorSpecification
		return ret
	}
	return *o.RegionalDisruptorSpecifications
}

// SetRegionalDisruptorSpecifications gets a reference to the given []RegionalDisruptorSpecification and assigns it to the RegionalDisruptorSpecifications field.
func (o *CockroachCloudUpdateClusterDisruptionRequest) SetRegionalDisruptorSpecifications(v []RegionalDisruptorSpecification) {
	o.RegionalDisruptorSpecifications = &v
}
