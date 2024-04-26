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

// CreateApiKeyRequest struct for CreateApiKeyRequest.
type CreateApiKeyRequest struct {
	// The name of the api key.
	Name string `json:"name"`
	// The ID of the service account to create the api key for.
	ServiceAccountId string `json:"service_account_id"`
}

// NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyRequest(name string, serviceAccountId string) *CreateApiKeyRequest {
	p := CreateApiKeyRequest{}
	p.Name = name
	p.ServiceAccountId = serviceAccountId
	return &p
}

// NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest {
	p := CreateApiKeyRequest{}
	return &p
}

// GetName returns the Name field value.
func (o *CreateApiKeyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *CreateApiKeyRequest) SetName(v string) {
	o.Name = v
}

// GetServiceAccountId returns the ServiceAccountId field value.
func (o *CreateApiKeyRequest) GetServiceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAccountId
}

// SetServiceAccountId sets field value.
func (o *CreateApiKeyRequest) SetServiceAccountId(v string) {
	o.ServiceAccountId = v
}