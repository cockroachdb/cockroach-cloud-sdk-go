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

// DeleteEgressRuleResponse struct for DeleteEgressRuleResponse.
type DeleteEgressRuleResponse struct {
	Rule *EgressRule `json:"Rule,omitempty"`
}

// NewDeleteEgressRuleResponse instantiates a new DeleteEgressRuleResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteEgressRuleResponse() *DeleteEgressRuleResponse {
	p := DeleteEgressRuleResponse{}
	return &p
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *DeleteEgressRuleResponse) GetRule() EgressRule {
	if o == nil || o.Rule == nil {
		var ret EgressRule
		return ret
	}
	return *o.Rule
}

// SetRule gets a reference to the given EgressRule and assigns it to the Rule field.
func (o *DeleteEgressRuleResponse) SetRule(v EgressRule) {
	o.Rule = &v
}
