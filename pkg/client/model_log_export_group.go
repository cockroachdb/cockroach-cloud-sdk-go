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

// LogExportGroup LogExportGroup contains an export configuration for a single log group which can route logs for a subset of CRDB channels..
type LogExportGroup struct {
	// channels is a list of CRDB log channels to include in this group.
	Channels []string `json:"channels"`
	// log_name is the name of the group, reflected in the log sink.
	LogName  string        `json:"log_name"`
	MinLevel *LogLevelType `json:"min_level,omitempty"`
	// redact is a boolean that governs whether this log group should aggregate redacted logs. Redaction settings will inherit from the cluster log export defaults if unset.
	Redact *bool `json:"redact,omitempty"`
}

// NewLogExportGroup instantiates a new LogExportGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExportGroup(channels []string, logName string) *LogExportGroup {
	p := LogExportGroup{}
	p.Channels = channels
	p.LogName = logName
	return &p
}

// NewLogExportGroupWithDefaults instantiates a new LogExportGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExportGroupWithDefaults() *LogExportGroup {
	p := LogExportGroup{}
	return &p
}

// GetChannels returns the Channels field value.
func (o *LogExportGroup) GetChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Channels
}

// SetChannels sets field value.
func (o *LogExportGroup) SetChannels(v []string) {
	o.Channels = v
}

// GetLogName returns the LogName field value.
func (o *LogExportGroup) GetLogName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogName
}

// SetLogName sets field value.
func (o *LogExportGroup) SetLogName(v string) {
	o.LogName = v
}

// GetMinLevel returns the MinLevel field value if set, zero value otherwise.
func (o *LogExportGroup) GetMinLevel() LogLevelType {
	if o == nil || o.MinLevel == nil {
		var ret LogLevelType
		return ret
	}
	return *o.MinLevel
}

// SetMinLevel gets a reference to the given LogLevelType and assigns it to the MinLevel field.
func (o *LogExportGroup) SetMinLevel(v LogLevelType) {
	o.MinLevel = &v
}

// GetRedact returns the Redact field value if set, zero value otherwise.
func (o *LogExportGroup) GetRedact() bool {
	if o == nil || o.Redact == nil {
		var ret bool
		return ret
	}
	return *o.Redact
}

// SetRedact gets a reference to the given bool and assigns it to the Redact field.
func (o *LogExportGroup) SetRedact(v bool) {
	o.Redact = &v
}
