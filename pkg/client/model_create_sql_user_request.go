// Copyright 2022 The Cockroach Authors
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
// API version: 2022-03-31

package client

import (
	"encoding/json"
)

// CreateSQLUserRequest struct for CreateSQLUserRequest.
type CreateSQLUserRequest struct {
	Name                 string `json:"name"`
	Password             string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type createSQLUserRequest CreateSQLUserRequest

// NewCreateSQLUserRequest instantiates a new CreateSQLUserRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSQLUserRequest(name string, password string) *CreateSQLUserRequest {
	p := CreateSQLUserRequest{}
	p.Name = name
	p.Password = password
	return &p
}

// NewCreateSQLUserRequestWithDefaults instantiates a new CreateSQLUserRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSQLUserRequestWithDefaults() *CreateSQLUserRequest {
	p := CreateSQLUserRequest{}
	return &p
}

// GetName returns the Name field value.
func (o *CreateSQLUserRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *CreateSQLUserRequest) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value.
func (o *CreateSQLUserRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value.
func (o *CreateSQLUserRequest) SetPassword(v string) {
	o.Password = v
}

func (o CreateSQLUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateSQLUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateSQLUserRequest := createSQLUserRequest{}

	if err = json.Unmarshal(bytes, &varCreateSQLUserRequest); err == nil {
		*o = CreateSQLUserRequest(varCreateSQLUserRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
