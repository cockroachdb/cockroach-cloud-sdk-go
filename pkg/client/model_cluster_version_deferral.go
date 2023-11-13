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
// API version: development

package client

// ClusterVersionDeferral ClusterVersionDeferral specifies whether automatic patch version upgrades are applied immediately or deferred. If upgrades are deferred, the cluster will be automatically upgraded to each patch version 60 days after the version is released to CockroachDB Cloud..
type ClusterVersionDeferral struct {
	DeferralPolicy ClusterVersionDeferralPolicyType `json:"deferral_policy"`
}

// NewClusterVersionDeferral instantiates a new ClusterVersionDeferral object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterVersionDeferral(deferralPolicy ClusterVersionDeferralPolicyType) *ClusterVersionDeferral {
	p := ClusterVersionDeferral{}
	p.DeferralPolicy = deferralPolicy
	return &p
}

// NewClusterVersionDeferralWithDefaults instantiates a new ClusterVersionDeferral object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterVersionDeferralWithDefaults() *ClusterVersionDeferral {
	p := ClusterVersionDeferral{}
	return &p
}

// GetDeferralPolicy returns the DeferralPolicy field value.
func (o *ClusterVersionDeferral) GetDeferralPolicy() ClusterVersionDeferralPolicyType {
	if o == nil {
		var ret ClusterVersionDeferralPolicyType
		return ret
	}

	return o.DeferralPolicy
}

// SetDeferralPolicy sets field value.
func (o *ClusterVersionDeferral) SetDeferralPolicy(v ClusterVersionDeferralPolicyType) {
	o.DeferralPolicy = v
}
