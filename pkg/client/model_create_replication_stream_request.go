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

// CreateReplicationStreamRequest struct for CreateReplicationStreamRequest.
type CreateReplicationStreamRequest struct {
	// source_cluster_id is the ID of the cluster that is being replicated.
	SourceClusterId string `json:"source_cluster_id"`
	// target_cluster_id is the ID of the cluster that data is being replicated to.
	TargetClusterId string `json:"target_cluster_id"`
}

// NewCreateReplicationStreamRequest instantiates a new CreateReplicationStreamRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReplicationStreamRequest(sourceClusterId string, targetClusterId string) *CreateReplicationStreamRequest {
	p := CreateReplicationStreamRequest{}
	p.SourceClusterId = sourceClusterId
	p.TargetClusterId = targetClusterId
	return &p
}

// NewCreateReplicationStreamRequestWithDefaults instantiates a new CreateReplicationStreamRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReplicationStreamRequestWithDefaults() *CreateReplicationStreamRequest {
	p := CreateReplicationStreamRequest{}
	return &p
}

// GetSourceClusterId returns the SourceClusterId field value.
func (o *CreateReplicationStreamRequest) GetSourceClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceClusterId
}

// SetSourceClusterId sets field value.
func (o *CreateReplicationStreamRequest) SetSourceClusterId(v string) {
	o.SourceClusterId = v
}

// GetTargetClusterId returns the TargetClusterId field value.
func (o *CreateReplicationStreamRequest) GetTargetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetClusterId
}

// SetTargetClusterId sets field value.
func (o *CreateReplicationStreamRequest) SetTargetClusterId(v string) {
	o.TargetClusterId = v
}
