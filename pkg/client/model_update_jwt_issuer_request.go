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

// UpdateJWTIssuerRequest struct for UpdateJWTIssuerRequest.
type UpdateJWTIssuerRequest struct {
	Audience    *string                      `json:"audience,omitempty"`
	Claim       *string                      `json:"claim,omitempty"`
	IdentityMap *[]JWTIssuerIdentityMapEntry `json:"identity_map,omitempty"`
	IssuerUrl   *string                      `json:"issuer_url,omitempty"`
	Jwks        *string                      `json:"jwks,omitempty"`
}

// NewUpdateJWTIssuerRequest instantiates a new UpdateJWTIssuerRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateJWTIssuerRequest() *UpdateJWTIssuerRequest {
	p := UpdateJWTIssuerRequest{}
	return &p
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *UpdateJWTIssuerRequest) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *UpdateJWTIssuerRequest) SetAudience(v string) {
	o.Audience = &v
}

// GetClaim returns the Claim field value if set, zero value otherwise.
func (o *UpdateJWTIssuerRequest) GetClaim() string {
	if o == nil || o.Claim == nil {
		var ret string
		return ret
	}
	return *o.Claim
}

// SetClaim gets a reference to the given string and assigns it to the Claim field.
func (o *UpdateJWTIssuerRequest) SetClaim(v string) {
	o.Claim = &v
}

// GetIdentityMap returns the IdentityMap field value if set, zero value otherwise.
func (o *UpdateJWTIssuerRequest) GetIdentityMap() []JWTIssuerIdentityMapEntry {
	if o == nil || o.IdentityMap == nil {
		var ret []JWTIssuerIdentityMapEntry
		return ret
	}
	return *o.IdentityMap
}

// SetIdentityMap gets a reference to the given []JWTIssuerIdentityMapEntry and assigns it to the IdentityMap field.
func (o *UpdateJWTIssuerRequest) SetIdentityMap(v []JWTIssuerIdentityMapEntry) {
	o.IdentityMap = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *UpdateJWTIssuerRequest) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *UpdateJWTIssuerRequest) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *UpdateJWTIssuerRequest) GetJwks() string {
	if o == nil || o.Jwks == nil {
		var ret string
		return ret
	}
	return *o.Jwks
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *UpdateJWTIssuerRequest) SetJwks(v string) {
	o.Jwks = &v
}
