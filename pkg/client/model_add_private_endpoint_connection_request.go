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

// AddPrivateEndpointConnectionRequest struct for AddPrivateEndpointConnectionRequest.
type AddPrivateEndpointConnectionRequest struct {
	// endpoint_id is the id of the private endpoint associated with a cluster's private endpoint service. The private endpoint is customer-created and its id is generated by the cloud provider at endpoint creation time.
	EndpointId string `json:"endpoint_id"`
}

// NewAddPrivateEndpointConnectionRequest instantiates a new AddPrivateEndpointConnectionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPrivateEndpointConnectionRequest(endpointId string) *AddPrivateEndpointConnectionRequest {
	p := AddPrivateEndpointConnectionRequest{}
	p.EndpointId = endpointId
	return &p
}

// NewAddPrivateEndpointConnectionRequestWithDefaults instantiates a new AddPrivateEndpointConnectionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPrivateEndpointConnectionRequestWithDefaults() *AddPrivateEndpointConnectionRequest {
	p := AddPrivateEndpointConnectionRequest{}
	return &p
}

// GetEndpointId returns the EndpointId field value.
func (o *AddPrivateEndpointConnectionRequest) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// SetEndpointId sets field value.
func (o *AddPrivateEndpointConnectionRequest) SetEndpointId(v string) {
	o.EndpointId = v
}
