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
	"time"
)

// UpdateReplicationStreamSpec struct for UpdateReplicationStreamSpec.
type UpdateReplicationStreamSpec struct {
	// cutover describes whether or not to trigger cutover.
	Cutover bool `json:"cutover"`
	// cutover_at is the timestamp at which cutover occurs. If the user sets `cutover` to `true` but omits cutover_at, the cutover time will default to the latest consistent replicated time. Otherwise, the user can pick a time in the future to schedule a cutover in the future, or a time in the past to restore the cluster to a recent state.
	CutoverAt *time.Time `json:"cutover_at,omitempty"`
}

// NewUpdateReplicationStreamSpec instantiates a new UpdateReplicationStreamSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReplicationStreamSpec(cutover bool) *UpdateReplicationStreamSpec {
	p := UpdateReplicationStreamSpec{}
	p.Cutover = cutover
	return &p
}

// NewUpdateReplicationStreamSpecWithDefaults instantiates a new UpdateReplicationStreamSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReplicationStreamSpecWithDefaults() *UpdateReplicationStreamSpec {
	p := UpdateReplicationStreamSpec{}
	return &p
}

// GetCutover returns the Cutover field value.
func (o *UpdateReplicationStreamSpec) GetCutover() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cutover
}

// SetCutover sets field value.
func (o *UpdateReplicationStreamSpec) SetCutover(v bool) {
	o.Cutover = v
}

// GetCutoverAt returns the CutoverAt field value if set, zero value otherwise.
func (o *UpdateReplicationStreamSpec) GetCutoverAt() time.Time {
	if o == nil || o.CutoverAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CutoverAt
}

// SetCutoverAt gets a reference to the given time.Time and assigns it to the CutoverAt field.
func (o *UpdateReplicationStreamSpec) SetCutoverAt(v time.Time) {
	o.CutoverAt = &v
}