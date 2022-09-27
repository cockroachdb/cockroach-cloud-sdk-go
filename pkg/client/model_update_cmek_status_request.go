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
// API version: 2022-09-20

package client

import (
	"encoding/json"
)

// UpdateCMEKStatusRequest struct for UpdateCMEKStatusRequest.
type UpdateCMEKStatusRequest struct {
	Action               CMEKCustomerAction `json:"action"`
	AdditionalProperties map[string]interface{}
}

type updateCMEKStatusRequest UpdateCMEKStatusRequest

// NewUpdateCMEKStatusRequest instantiates a new UpdateCMEKStatusRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCMEKStatusRequest(action CMEKCustomerAction) *UpdateCMEKStatusRequest {
	p := UpdateCMEKStatusRequest{}
	p.Action = action
	return &p
}

// NewUpdateCMEKStatusRequestWithDefaults instantiates a new UpdateCMEKStatusRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCMEKStatusRequestWithDefaults() *UpdateCMEKStatusRequest {
	p := UpdateCMEKStatusRequest{}
	return &p
}

// GetAction returns the Action field value.
func (o *UpdateCMEKStatusRequest) GetAction() CMEKCustomerAction {
	if o == nil {
		var ret CMEKCustomerAction
		return ret
	}

	return o.Action
}

// SetAction sets field value.
func (o *UpdateCMEKStatusRequest) SetAction(v CMEKCustomerAction) {
	o.Action = v
}

func (o UpdateCMEKStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateCMEKStatusRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateCMEKStatusRequest := updateCMEKStatusRequest{}

	if err = json.Unmarshal(bytes, &varUpdateCMEKStatusRequest); err == nil {
		*o = UpdateCMEKStatusRequest(varUpdateCMEKStatusRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
