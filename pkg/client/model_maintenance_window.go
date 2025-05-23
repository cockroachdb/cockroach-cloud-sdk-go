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

// MaintenanceWindow MaintenanceWindowSpec specifies a weekly recurring maintenance window for an ADVANCED cluster..
type MaintenanceWindow struct {
	// OffsetDuration is the duration from the start of a week (Monday 00:00 UTC) that this maintenance window will start after.  Must be less than 1 week.
	OffsetDuration string `json:"offset_duration"`
	// WindowDuration is the duration of the maintenance window.  Must be at least 6 hours and less than 1 week.
	WindowDuration string `json:"window_duration"`
}

// NewMaintenanceWindow instantiates a new MaintenanceWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindow(offsetDuration string, windowDuration string) *MaintenanceWindow {
	p := MaintenanceWindow{}
	p.OffsetDuration = offsetDuration
	p.WindowDuration = windowDuration
	return &p
}

// NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowWithDefaults() *MaintenanceWindow {
	p := MaintenanceWindow{}
	return &p
}

// GetOffsetDuration returns the OffsetDuration field value.
func (o *MaintenanceWindow) GetOffsetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OffsetDuration
}

// SetOffsetDuration sets field value.
func (o *MaintenanceWindow) SetOffsetDuration(v string) {
	o.OffsetDuration = v
}

// GetWindowDuration returns the WindowDuration field value.
func (o *MaintenanceWindow) GetWindowDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WindowDuration
}

// SetWindowDuration sets field value.
func (o *MaintenanceWindow) SetWindowDuration(v string) {
	o.WindowDuration = v
}
