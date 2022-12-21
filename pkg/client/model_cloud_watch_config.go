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

// CloudWatchConfig struct for CloudWatchConfig.
type CloudWatchConfig struct {
	// log_group_name is the customized log group name.
	LogGroupName *string `json:"log_group_name,omitempty"`
	// role_arn is the IAM role used to upload metric segments to the target AWS account.
	RoleArn *string `json:"role_arn,omitempty"`
	// target_region specifies the specific AWS region that the metrics will be exported to.
	TargetRegion *string `json:"target_region,omitempty"`
}

// NewCloudWatchConfig instantiates a new CloudWatchConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudWatchConfig() *CloudWatchConfig {
	p := CloudWatchConfig{}
	return &p
}

// GetLogGroupName returns the LogGroupName field value if set, zero value otherwise.
func (o *CloudWatchConfig) GetLogGroupName() string {
	if o == nil || o.LogGroupName == nil {
		var ret string
		return ret
	}
	return *o.LogGroupName
}

// SetLogGroupName gets a reference to the given string and assigns it to the LogGroupName field.
func (o *CloudWatchConfig) SetLogGroupName(v string) {
	o.LogGroupName = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise.
func (o *CloudWatchConfig) GetRoleArn() string {
	if o == nil || o.RoleArn == nil {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *CloudWatchConfig) SetRoleArn(v string) {
	o.RoleArn = &v
}

// GetTargetRegion returns the TargetRegion field value if set, zero value otherwise.
func (o *CloudWatchConfig) GetTargetRegion() string {
	if o == nil || o.TargetRegion == nil {
		var ret string
		return ret
	}
	return *o.TargetRegion
}

// SetTargetRegion gets a reference to the given string and assigns it to the TargetRegion field.
func (o *CloudWatchConfig) SetTargetRegion(v string) {
	o.TargetRegion = &v
}

func (o CloudWatchConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogGroupName != nil {
		toSerialize["log_group_name"] = o.LogGroupName
	}
	if o.RoleArn != nil {
		toSerialize["role_arn"] = o.RoleArn
	}
	if o.TargetRegion != nil {
		toSerialize["target_region"] = o.TargetRegion
	}
	return json.Marshal(toSerialize)
}
