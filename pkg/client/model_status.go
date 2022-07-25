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

// Status struct for Status.
type Status struct {
	Code                 *int32  `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Details              *[]Any  `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type status Status

// NewStatus instantiates a new Status object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus() *Status {
	p := Status{}
	return &p
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Status) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Status) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Status) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Status) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Status) GetDetails() []Any {
	if o == nil || o.Details == nil {
		var ret []Any
		return ret
	}
	return *o.Details
}

// SetDetails gets a reference to the given []Any and assigns it to the Details field.
func (o *Status) SetDetails(v []Any) {
	o.Details = &v
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Status) UnmarshalJSON(bytes []byte) (err error) {
	varStatus := status{}

	if err = json.Unmarshal(bytes, &varStatus); err == nil {
		*o = Status(varStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
