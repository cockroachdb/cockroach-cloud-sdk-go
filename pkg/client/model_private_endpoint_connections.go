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

// PrivateEndpointConnections struct for PrivateEndpointConnections.
type PrivateEndpointConnections struct {
	// Connections is a list of private endpoints.
	Connections []PrivateEndpointConnection `json:"connections"`
}

// NewPrivateEndpointConnections instantiates a new PrivateEndpointConnections object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateEndpointConnections(connections []PrivateEndpointConnection) *PrivateEndpointConnections {
	p := PrivateEndpointConnections{}
	p.Connections = connections
	return &p
}

// NewPrivateEndpointConnectionsWithDefaults instantiates a new PrivateEndpointConnections object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateEndpointConnectionsWithDefaults() *PrivateEndpointConnections {
	p := PrivateEndpointConnections{}
	return &p
}

// GetConnections returns the Connections field value.
func (o *PrivateEndpointConnections) GetConnections() []PrivateEndpointConnection {
	if o == nil {
		var ret []PrivateEndpointConnection
		return ret
	}

	return o.Connections
}

// SetConnections sets field value.
func (o *PrivateEndpointConnections) SetConnections(v []PrivateEndpointConnection) {
	o.Connections = v
}
