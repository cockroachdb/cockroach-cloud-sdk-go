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

// UpdateClusterSpecification struct for UpdateClusterSpecification.
type UpdateClusterSpecification struct {
	Dedicated        *DedicatedClusterUpdateSpecification `json:"dedicated,omitempty"`
	DeleteProtection *DeleteProtectionStateType           `json:"delete_protection,omitempty"`
	// Preview: The parent ID is a folder ID. An empty string or \"root\" represents the root level.
	ParentId      *string                               `json:"parent_id,omitempty"`
	Plan          *PlanType                             `json:"plan,omitempty"`
	Serverless    *ServerlessClusterUpdateSpecification `json:"serverless,omitempty"`
	UpgradeStatus *ClusterUpgradeStatusType             `json:"upgrade_status,omitempty"`
}

// NewUpdateClusterSpecification instantiates a new UpdateClusterSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClusterSpecification() *UpdateClusterSpecification {
	p := UpdateClusterSpecification{}
	return &p
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetDedicated() DedicatedClusterUpdateSpecification {
	if o == nil || o.Dedicated == nil {
		var ret DedicatedClusterUpdateSpecification
		return ret
	}
	return *o.Dedicated
}

// SetDedicated gets a reference to the given DedicatedClusterUpdateSpecification and assigns it to the Dedicated field.
func (o *UpdateClusterSpecification) SetDedicated(v DedicatedClusterUpdateSpecification) {
	o.Dedicated = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetDeleteProtection() DeleteProtectionStateType {
	if o == nil || o.DeleteProtection == nil {
		var ret DeleteProtectionStateType
		return ret
	}
	return *o.DeleteProtection
}

// SetDeleteProtection gets a reference to the given DeleteProtectionStateType and assigns it to the DeleteProtection field.
func (o *UpdateClusterSpecification) SetDeleteProtection(v DeleteProtectionStateType) {
	o.DeleteProtection = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *UpdateClusterSpecification) SetParentId(v string) {
	o.ParentId = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetPlan() PlanType {
	if o == nil || o.Plan == nil {
		var ret PlanType
		return ret
	}
	return *o.Plan
}

// SetPlan gets a reference to the given PlanType and assigns it to the Plan field.
func (o *UpdateClusterSpecification) SetPlan(v PlanType) {
	o.Plan = &v
}

// GetServerless returns the Serverless field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetServerless() ServerlessClusterUpdateSpecification {
	if o == nil || o.Serverless == nil {
		var ret ServerlessClusterUpdateSpecification
		return ret
	}
	return *o.Serverless
}

// SetServerless gets a reference to the given ServerlessClusterUpdateSpecification and assigns it to the Serverless field.
func (o *UpdateClusterSpecification) SetServerless(v ServerlessClusterUpdateSpecification) {
	o.Serverless = &v
}

// GetUpgradeStatus returns the UpgradeStatus field value if set, zero value otherwise.
func (o *UpdateClusterSpecification) GetUpgradeStatus() ClusterUpgradeStatusType {
	if o == nil || o.UpgradeStatus == nil {
		var ret ClusterUpgradeStatusType
		return ret
	}
	return *o.UpgradeStatus
}

// SetUpgradeStatus gets a reference to the given ClusterUpgradeStatusType and assigns it to the UpgradeStatus field.
func (o *UpdateClusterSpecification) SetUpgradeStatus(v ClusterUpgradeStatusType) {
	o.UpgradeStatus = &v
}
