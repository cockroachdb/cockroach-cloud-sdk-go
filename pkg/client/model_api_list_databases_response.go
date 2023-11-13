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
// API version: development

package client

// ApiListDatabasesResponse struct for ApiListDatabasesResponse.
type ApiListDatabasesResponse struct {
	Databases  []ApiDatabase             `json:"databases"`
	Pagination *KeysetPaginationResponse `json:"pagination,omitempty"`
}

// NewApiListDatabasesResponse instantiates a new ApiListDatabasesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiListDatabasesResponse(databases []ApiDatabase) *ApiListDatabasesResponse {
	p := ApiListDatabasesResponse{}
	p.Databases = databases
	return &p
}

// NewApiListDatabasesResponseWithDefaults instantiates a new ApiListDatabasesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiListDatabasesResponseWithDefaults() *ApiListDatabasesResponse {
	p := ApiListDatabasesResponse{}
	return &p
}

// GetDatabases returns the Databases field value.
func (o *ApiListDatabasesResponse) GetDatabases() []ApiDatabase {
	if o == nil {
		var ret []ApiDatabase
		return ret
	}

	return o.Databases
}

// SetDatabases sets field value.
func (o *ApiListDatabasesResponse) SetDatabases(v []ApiDatabase) {
	o.Databases = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ApiListDatabasesResponse) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ApiListDatabasesResponse) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}
