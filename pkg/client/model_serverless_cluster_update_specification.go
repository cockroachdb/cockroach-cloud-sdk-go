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

// ServerlessClusterUpdateSpecification struct for ServerlessClusterUpdateSpecification.
type ServerlessClusterUpdateSpecification struct {
	// spend_limit is the maximum monthly charge for a cluster, in US cents. We recommend using usage_limits instead, since spend_limit will be deprecated in the future.
	SpendLimit  Optional[int32] `json:"spend_limit,omitempty"`
	UsageLimits *UsageLimits    `json:"usage_limits,omitempty"`
}

// NewServerlessClusterUpdateSpecification instantiates a new ServerlessClusterUpdateSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessClusterUpdateSpecification() *ServerlessClusterUpdateSpecification {
	p := ServerlessClusterUpdateSpecification{}
	return &p
}

// GetSpendLimit returns the SpendLimit field value if set, zero value otherwise.
func (o *ServerlessClusterUpdateSpecification) GetSpendLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.SpendLimit.GetValue()
}

// SetSpendLimit gets a reference to the given int32 and assigns it to the SpendLimit field.
func (o *ServerlessClusterUpdateSpecification) SetSpendLimit(v *int32, marshalNull bool) {
	o.SpendLimit = Optional[int32]{inner: v, marshalNull: marshalNull}
}

// GetUsageLimits returns the UsageLimits field value if set, zero value otherwise.
func (o *ServerlessClusterUpdateSpecification) GetUsageLimits() UsageLimits {
	if o == nil || o.UsageLimits == nil {
		var ret UsageLimits
		return ret
	}
	return *o.UsageLimits
}

// SetUsageLimits gets a reference to the given UsageLimits and assigns it to the UsageLimits field.
func (o *ServerlessClusterUpdateSpecification) SetUsageLimits(v UsageLimits) {
	o.UsageLimits = &v
}

func (o ServerlessClusterUpdateSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpendLimit.marshalNull || o.SpendLimit.inner != nil {
		toSerialize["spend_limit"] = o.SpendLimit
	}
	if o.UsageLimits != nil {
		toSerialize["usage_limits"] = o.UsageLimits
	}
	return json.Marshal(toSerialize)
}
