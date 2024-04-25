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

// ListServiceAccountsResponse struct for ListServiceAccountsResponse.
type ListServiceAccountsResponse struct {
	Pagination      *KeysetPaginationResponse `json:"pagination,omitempty"`
	ServiceAccounts []ServiceAccount          `json:"service_accounts"`
}

// NewListServiceAccountsResponse instantiates a new ListServiceAccountsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceAccountsResponse(serviceAccounts []ServiceAccount) *ListServiceAccountsResponse {
	p := ListServiceAccountsResponse{}
	p.ServiceAccounts = serviceAccounts
	return &p
}

// NewListServiceAccountsResponseWithDefaults instantiates a new ListServiceAccountsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceAccountsResponseWithDefaults() *ListServiceAccountsResponse {
	p := ListServiceAccountsResponse{}
	return &p
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListServiceAccountsResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ListServiceAccountsResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}

// GetServiceAccounts returns the ServiceAccounts field value.
func (o *ListServiceAccountsResponse) GetServiceAccounts() []ServiceAccount {
	if o == nil {
		var ret []ServiceAccount
		return ret
	}

	return o.ServiceAccounts
}

// SetServiceAccounts sets field value.
func (o *ListServiceAccountsResponse) SetServiceAccounts(v []ServiceAccount) {
	o.ServiceAccounts = v
}
