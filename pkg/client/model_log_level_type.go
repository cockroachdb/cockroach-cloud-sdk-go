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

import (
	"fmt"
)

// LogLevelType  - UNSPECIFIED: The unspecified log level includes all logs.  - WARNING: The WARNING severity is used for situations which may require special handling, where normal operation is expected to resume automatically.  - ERROR: The ERROR severity is used for situations that require special handling, where normal operation could not proceed as expected. Other operations can continue mostly unaffected.  - FATAL: The FATAL severity is used for situations that require an immediate, hard server shutdown. A report is also sent to telemetry if telemetry is enabled.
type LogLevelType string

// List of LogLevel.Type.
const (
	LOGLEVELTYPE_UNSPECIFIED LogLevelType = "UNSPECIFIED"
	LOGLEVELTYPE_WARNING     LogLevelType = "WARNING"
	LOGLEVELTYPE_ERROR       LogLevelType = "ERROR"
	LOGLEVELTYPE_FATAL       LogLevelType = "FATAL"
)

// All allowed values of LogLevelType enum.
var AllowedLogLevelTypeEnumValues = []LogLevelType{
	"UNSPECIFIED",
	"WARNING",
	"ERROR",
	"FATAL",
}

// NewLogLevelTypeFromValue returns a pointer to a valid LogLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogLevelTypeFromValue(v string) (*LogLevelType, error) {
	ev := LogLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogLevelType: valid values are %v", v, AllowedLogLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogLevelType) IsValid() bool {
	for _, existing := range AllowedLogLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogLevel.Type value.
func (v LogLevelType) Ptr() *LogLevelType {
	return &v
}
