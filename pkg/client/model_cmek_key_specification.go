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

import (
	"encoding/json"
)

// CMEKKeySpecification CMEKKeySpecification contains all the details necessary to use a customer-provided encryption key.  This involves the type/location of the key and the principal to authenticate as  when accessing it..
type CMEKKeySpecification struct {
	AuthPrincipal *string      `json:"auth_principal,omitempty,string"`
	Type          *CMEKKeyType `json:"type,omitempty"`
	Uri           *string      `json:"uri,omitempty,string"`
}

// NewCMEKKeySpecification instantiates a new CMEKKeySpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKKeySpecification() *CMEKKeySpecification {
	p := CMEKKeySpecification{}
	return &p
}

// GetAuthPrincipal returns the AuthPrincipal field value if set, zero value otherwise.
func (o *CMEKKeySpecification) GetAuthPrincipal() string {
	if o == nil || o.AuthPrincipal == nil {
		var ret string
		return ret
	}
	return *o.AuthPrincipal
}

// SetAuthPrincipal gets a reference to the given string and assigns it to the AuthPrincipal field.
func (o *CMEKKeySpecification) SetAuthPrincipal(v string) {
	o.AuthPrincipal = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CMEKKeySpecification) GetType() CMEKKeyType {
	if o == nil || o.Type == nil {
		var ret CMEKKeyType
		return ret
	}
	return *o.Type
}

// SetType gets a reference to the given CMEKKeyType and assigns it to the Type field.
func (o *CMEKKeySpecification) SetType(v CMEKKeyType) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *CMEKKeySpecification) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *CMEKKeySpecification) SetUri(v string) {
	o.Uri = &v
}

func (o CMEKKeySpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthPrincipal != nil {
		toSerialize["auth_principal"] = o.AuthPrincipal
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}
