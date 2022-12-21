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
	"time"
)

// Cluster struct for Cluster.
type Cluster struct {
	AccountId         *string           `json:"account_id,omitempty"`
	CloudProvider     ApiCloudProvider  `json:"cloud_provider"`
	CockroachVersion  string            `json:"cockroach_version"`
	Config            ClusterConfig     `json:"config"`
	CreatedAt         *time.Time        `json:"created_at,omitempty"`
	CreatorId         string            `json:"creator_id"`
	DeletedAt         *time.Time        `json:"deleted_at,omitempty"`
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	NetworkVisibility *NetworkVisiblity `json:"network_visibility,omitempty"`
	OperationStatus   ClusterStatusType `json:"operation_status"`
	Plan              Plan              `json:"plan"`
	Regions           []Region          `json:"regions"`
	// sql_dns is the DNS name of SQL interface of the cluster.
	SqlDns    *string          `json:"sql_dns,omitempty"`
	State     ClusterStateType `json:"state"`
	UpdatedAt *time.Time       `json:"updated_at,omitempty"`
}

// NewCluster instantiates a new Cluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(cloudProvider ApiCloudProvider, cockroachVersion string, config ClusterConfig, creatorId string, id string, name string, operationStatus ClusterStatusType, plan Plan, regions []Region, state ClusterStateType) *Cluster {
	p := Cluster{}
	p.CloudProvider = cloudProvider
	p.CockroachVersion = cockroachVersion
	p.Config = config
	p.CreatorId = creatorId
	p.Id = id
	p.Name = name
	p.OperationStatus = operationStatus
	p.Plan = plan
	p.Regions = regions
	p.State = state
	return &p
}

// NewClusterWithDefaults instantiates a new Cluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	p := Cluster{}
	return &p
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Cluster) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Cluster) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCloudProvider returns the CloudProvider field value.
func (o *Cluster) GetCloudProvider() ApiCloudProvider {
	if o == nil {
		var ret ApiCloudProvider
		return ret
	}

	return o.CloudProvider
}

// SetCloudProvider sets field value.
func (o *Cluster) SetCloudProvider(v ApiCloudProvider) {
	o.CloudProvider = v
}

// GetCockroachVersion returns the CockroachVersion field value.
func (o *Cluster) GetCockroachVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CockroachVersion
}

// SetCockroachVersion sets field value.
func (o *Cluster) SetCockroachVersion(v string) {
	o.CockroachVersion = v
}

// GetConfig returns the Config field value.
func (o *Cluster) GetConfig() ClusterConfig {
	if o == nil {
		var ret ClusterConfig
		return ret
	}

	return o.Config
}

// SetConfig sets field value.
func (o *Cluster) SetConfig(v ClusterConfig) {
	o.Config = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Cluster) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Cluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatorId returns the CreatorId field value.
func (o *Cluster) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// SetCreatorId sets field value.
func (o *Cluster) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Cluster) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Cluster) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value.
func (o *Cluster) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// SetId sets field value.
func (o *Cluster) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetNetworkVisibility returns the NetworkVisibility field value if set, zero value otherwise.
func (o *Cluster) GetNetworkVisibility() NetworkVisiblity {
	if o == nil || o.NetworkVisibility == nil {
		var ret NetworkVisiblity
		return ret
	}
	return *o.NetworkVisibility
}

// SetNetworkVisibility gets a reference to the given NetworkVisiblity and assigns it to the NetworkVisibility field.
func (o *Cluster) SetNetworkVisibility(v NetworkVisiblity) {
	o.NetworkVisibility = &v
}

// GetOperationStatus returns the OperationStatus field value.
func (o *Cluster) GetOperationStatus() ClusterStatusType {
	if o == nil {
		var ret ClusterStatusType
		return ret
	}

	return o.OperationStatus
}

// SetOperationStatus sets field value.
func (o *Cluster) SetOperationStatus(v ClusterStatusType) {
	o.OperationStatus = v
}

// GetPlan returns the Plan field value.
func (o *Cluster) GetPlan() Plan {
	if o == nil {
		var ret Plan
		return ret
	}

	return o.Plan
}

// SetPlan sets field value.
func (o *Cluster) SetPlan(v Plan) {
	o.Plan = v
}

// GetRegions returns the Regions field value.
func (o *Cluster) GetRegions() []Region {
	if o == nil {
		var ret []Region
		return ret
	}

	return o.Regions
}

// SetRegions sets field value.
func (o *Cluster) SetRegions(v []Region) {
	o.Regions = v
}

// GetSqlDns returns the SqlDns field value if set, zero value otherwise.
func (o *Cluster) GetSqlDns() string {
	if o == nil || o.SqlDns == nil {
		var ret string
		return ret
	}
	return *o.SqlDns
}

// SetSqlDns gets a reference to the given string and assigns it to the SqlDns field.
func (o *Cluster) SetSqlDns(v string) {
	o.SqlDns = &v
}

// GetState returns the State field value.
func (o *Cluster) GetState() ClusterStateType {
	if o == nil {
		var ret ClusterStateType
		return ret
	}

	return o.State
}

// SetState sets field value.
func (o *Cluster) SetState(v ClusterStateType) {
	o.State = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Cluster) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Cluster) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if true {
		toSerialize["cockroach_version"] = o.CockroachVersion
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["creator_id"] = o.CreatorId
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NetworkVisibility != nil {
		toSerialize["network_visibility"] = o.NetworkVisibility
	}
	if true {
		toSerialize["operation_status"] = o.OperationStatus
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if true {
		toSerialize["regions"] = o.Regions
	}
	if o.SqlDns != nil {
		toSerialize["sql_dns"] = o.SqlDns
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}
