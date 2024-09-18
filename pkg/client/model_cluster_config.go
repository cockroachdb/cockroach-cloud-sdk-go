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

// ClusterConfig struct for ClusterConfig.
type ClusterConfig struct {
	Dedicated  *DedicatedHardwareConfig `json:"dedicated,omitempty"`
	Serverless *ServerlessClusterConfig `json:"serverless,omitempty"`
}

// NewClusterConfig instantiates a new ClusterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfig() *ClusterConfig {
	p := ClusterConfig{}
	return &p
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *ClusterConfig) GetDedicated() DedicatedHardwareConfig {
	if o == nil || o.Dedicated == nil {
		var ret DedicatedHardwareConfig
		return ret
	}
	return *o.Dedicated
}

// SetDedicated gets a reference to the given DedicatedHardwareConfig and assigns it to the Dedicated field.
func (o *ClusterConfig) SetDedicated(v DedicatedHardwareConfig) {
	o.Dedicated = &v
}

// GetServerless returns the Serverless field value if set, zero value otherwise.
func (o *ClusterConfig) GetServerless() ServerlessClusterConfig {
	if o == nil || o.Serverless == nil {
		var ret ServerlessClusterConfig
		return ret
	}
	return *o.Serverless
}

// SetServerless gets a reference to the given ServerlessClusterConfig and assigns it to the Serverless field.
func (o *ClusterConfig) SetServerless(v ServerlessClusterConfig) {
	o.Serverless = &v
}
