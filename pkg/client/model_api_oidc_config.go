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

// ApiOidcConfig api_oidc_config contains information about an OIDC provider that can generate JWT tokens that can be used for authentication with the Cockroach Cloud API..
type ApiOidcConfig struct {
	Audience    string                     `json:"audience"`
	Claim       *string                    `json:"claim,omitempty"`
	Id          string                     `json:"id"`
	IdentityMap *[]ApiOidcIdentityMapEntry `json:"identity_map,omitempty"`
	Issuer      string                     `json:"issuer"`
	Jwks        string                     `json:"jwks"`
}

// NewApiOidcConfig instantiates a new ApiOidcConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOidcConfig(audience string, id string, issuer string, jwks string) *ApiOidcConfig {
	p := ApiOidcConfig{}
	p.Audience = audience
	p.Id = id
	p.Issuer = issuer
	p.Jwks = jwks
	return &p
}

// NewApiOidcConfigWithDefaults instantiates a new ApiOidcConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOidcConfigWithDefaults() *ApiOidcConfig {
	p := ApiOidcConfig{}
	return &p
}

// GetAudience returns the Audience field value.
func (o *ApiOidcConfig) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// SetAudience sets field value.
func (o *ApiOidcConfig) SetAudience(v string) {
	o.Audience = v
}

// GetClaim returns the Claim field value if set, zero value otherwise.
func (o *ApiOidcConfig) GetClaim() string {
	if o == nil || o.Claim == nil {
		var ret string
		return ret
	}
	return *o.Claim
}

// SetClaim gets a reference to the given string and assigns it to the Claim field.
func (o *ApiOidcConfig) SetClaim(v string) {
	o.Claim = &v
}

// GetId returns the Id field value.
func (o *ApiOidcConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// SetId sets field value.
func (o *ApiOidcConfig) SetId(v string) {
	o.Id = v
}

// GetIdentityMap returns the IdentityMap field value if set, zero value otherwise.
func (o *ApiOidcConfig) GetIdentityMap() []ApiOidcIdentityMapEntry {
	if o == nil || o.IdentityMap == nil {
		var ret []ApiOidcIdentityMapEntry
		return ret
	}
	return *o.IdentityMap
}

// SetIdentityMap gets a reference to the given []ApiOidcIdentityMapEntry and assigns it to the IdentityMap field.
func (o *ApiOidcConfig) SetIdentityMap(v []ApiOidcIdentityMapEntry) {
	o.IdentityMap = &v
}

// GetIssuer returns the Issuer field value.
func (o *ApiOidcConfig) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// SetIssuer sets field value.
func (o *ApiOidcConfig) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwks returns the Jwks field value.
func (o *ApiOidcConfig) GetJwks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jwks
}

// SetJwks sets field value.
func (o *ApiOidcConfig) SetJwks(v string) {
	o.Jwks = v
}
