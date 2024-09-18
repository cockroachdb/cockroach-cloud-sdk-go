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

// GetServiceProviderConfigResponse struct for GetServiceProviderConfigResponse.
type GetServiceProviderConfigResponse struct {
	AuthenticationSchemes *[]ScimAuthenticationScheme `json:"authenticationSchemes,omitempty"`
	Bulk                  *ScimBulkSupport            `json:"bulk,omitempty"`
	ChangePassword        *ScimChangePasswordSupport  `json:"changePassword,omitempty"`
	Etag                  *ScimEtagSupport            `json:"etag,omitempty"`
	Filter                *ScimFilterSupport          `json:"filter,omitempty"`
	Meta                  *ScimMetadata               `json:"meta,omitempty"`
	Schemas               *[]string                   `json:"schemas,omitempty"`
	Sort                  *ScimSortSupport            `json:"sort,omitempty"`
}

// NewGetServiceProviderConfigResponse instantiates a new GetServiceProviderConfigResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceProviderConfigResponse() *GetServiceProviderConfigResponse {
	p := GetServiceProviderConfigResponse{}
	return &p
}

// GetAuthenticationSchemes returns the AuthenticationSchemes field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetAuthenticationSchemes() []ScimAuthenticationScheme {
	if o == nil || o.AuthenticationSchemes == nil {
		var ret []ScimAuthenticationScheme
		return ret
	}
	return *o.AuthenticationSchemes
}

// SetAuthenticationSchemes gets a reference to the given []ScimAuthenticationScheme and assigns it to the AuthenticationSchemes field.
func (o *GetServiceProviderConfigResponse) SetAuthenticationSchemes(v []ScimAuthenticationScheme) {
	o.AuthenticationSchemes = &v
}

// GetBulk returns the Bulk field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetBulk() ScimBulkSupport {
	if o == nil || o.Bulk == nil {
		var ret ScimBulkSupport
		return ret
	}
	return *o.Bulk
}

// SetBulk gets a reference to the given ScimBulkSupport and assigns it to the Bulk field.
func (o *GetServiceProviderConfigResponse) SetBulk(v ScimBulkSupport) {
	o.Bulk = &v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetChangePassword() ScimChangePasswordSupport {
	if o == nil || o.ChangePassword == nil {
		var ret ScimChangePasswordSupport
		return ret
	}
	return *o.ChangePassword
}

// SetChangePassword gets a reference to the given ScimChangePasswordSupport and assigns it to the ChangePassword field.
func (o *GetServiceProviderConfigResponse) SetChangePassword(v ScimChangePasswordSupport) {
	o.ChangePassword = &v
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetEtag() ScimEtagSupport {
	if o == nil || o.Etag == nil {
		var ret ScimEtagSupport
		return ret
	}
	return *o.Etag
}

// SetEtag gets a reference to the given ScimEtagSupport and assigns it to the Etag field.
func (o *GetServiceProviderConfigResponse) SetEtag(v ScimEtagSupport) {
	o.Etag = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetFilter() ScimFilterSupport {
	if o == nil || o.Filter == nil {
		var ret ScimFilterSupport
		return ret
	}
	return *o.Filter
}

// SetFilter gets a reference to the given ScimFilterSupport and assigns it to the Filter field.
func (o *GetServiceProviderConfigResponse) SetFilter(v ScimFilterSupport) {
	o.Filter = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetMeta() ScimMetadata {
	if o == nil || o.Meta == nil {
		var ret ScimMetadata
		return ret
	}
	return *o.Meta
}

// SetMeta gets a reference to the given ScimMetadata and assigns it to the Meta field.
func (o *GetServiceProviderConfigResponse) SetMeta(v ScimMetadata) {
	o.Meta = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return *o.Schemas
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *GetServiceProviderConfigResponse) SetSchemas(v []string) {
	o.Schemas = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *GetServiceProviderConfigResponse) GetSort() ScimSortSupport {
	if o == nil || o.Sort == nil {
		var ret ScimSortSupport
		return ret
	}
	return *o.Sort
}

// SetSort gets a reference to the given ScimSortSupport and assigns it to the Sort field.
func (o *GetServiceProviderConfigResponse) SetSort(v ScimSortSupport) {
	o.Sort = &v
}
