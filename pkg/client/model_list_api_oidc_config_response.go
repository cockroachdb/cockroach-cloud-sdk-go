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

// ListApiOidcConfigResponse struct for ListApiOidcConfigResponse.
type ListApiOidcConfigResponse struct {
	ApiOidcConfigs *[]ApiOidcConfig          `json:"api_oidc_configs,omitempty"`
	Pagination     *KeysetPaginationResponse `json:"pagination,omitempty"`
}

// NewListApiOidcConfigResponse instantiates a new ListApiOidcConfigResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApiOidcConfigResponse() *ListApiOidcConfigResponse {
	p := ListApiOidcConfigResponse{}
	return &p
}

// GetApiOidcConfigs returns the ApiOidcConfigs field value if set, zero value otherwise.
func (o *ListApiOidcConfigResponse) GetApiOidcConfigs() []ApiOidcConfig {
	if o == nil || o.ApiOidcConfigs == nil {
		var ret []ApiOidcConfig
		return ret
	}
	return *o.ApiOidcConfigs
}

// SetApiOidcConfigs gets a reference to the given []ApiOidcConfig and assigns it to the ApiOidcConfigs field.
func (o *ListApiOidcConfigResponse) SetApiOidcConfigs(v []ApiOidcConfig) {
	o.ApiOidcConfigs = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListApiOidcConfigResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ListApiOidcConfigResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}
