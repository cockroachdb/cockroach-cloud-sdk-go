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

// GetSchemasResponse struct for GetSchemasResponse.
type GetSchemasResponse struct {
	Resources    *[]ScimSchema `json:"Resources,omitempty"`
	ItemsPerPage *int32        `json:"itemsPerPage,omitempty"`
	Schemas      []string      `json:"schemas"`
	StartIndex   *int32        `json:"startIndex,omitempty"`
	TotalResults int32         `json:"totalResults"`
}

// NewGetSchemasResponse instantiates a new GetSchemasResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSchemasResponse(schemas []string, totalResults int32) *GetSchemasResponse {
	p := GetSchemasResponse{}
	p.Schemas = schemas
	p.TotalResults = totalResults
	return &p
}

// NewGetSchemasResponseWithDefaults instantiates a new GetSchemasResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSchemasResponseWithDefaults() *GetSchemasResponse {
	p := GetSchemasResponse{}
	return &p
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *GetSchemasResponse) GetResources() []ScimSchema {
	if o == nil || o.Resources == nil {
		var ret []ScimSchema
		return ret
	}
	return *o.Resources
}

// SetResources gets a reference to the given []ScimSchema and assigns it to the Resources field.
func (o *GetSchemasResponse) SetResources(v []ScimSchema) {
	o.Resources = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *GetSchemasResponse) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *GetSchemasResponse) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetSchemas returns the Schemas field value.
func (o *GetSchemasResponse) GetSchemas() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Schemas
}

// SetSchemas sets field value.
func (o *GetSchemasResponse) SetSchemas(v []string) {
	o.Schemas = v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *GetSchemasResponse) GetStartIndex() int32 {
	if o == nil || o.StartIndex == nil {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *GetSchemasResponse) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetTotalResults returns the TotalResults field value.
func (o *GetSchemasResponse) GetTotalResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResults
}

// SetTotalResults sets field value.
func (o *GetSchemasResponse) SetTotalResults(v int32) {
	o.TotalResults = v
}
