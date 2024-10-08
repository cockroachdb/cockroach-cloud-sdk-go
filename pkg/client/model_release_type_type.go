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

import (
	"fmt"
)

// ReleaseTypeType the model 'ReleaseTypeType'.
type ReleaseTypeType string

// List of ReleaseType.Type.
const (
	RELEASETYPETYPE_REGULAR    ReleaseTypeType = "REGULAR"
	RELEASETYPETYPE_INNOVATION ReleaseTypeType = "INNOVATION"
)

// All allowed values of ReleaseTypeType enum.
var AllowedReleaseTypeTypeEnumValues = []ReleaseTypeType{
	"REGULAR",
	"INNOVATION",
}

// NewReleaseTypeTypeFromValue returns a pointer to a valid ReleaseTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReleaseTypeTypeFromValue(v string) (*ReleaseTypeType, error) {
	ev := ReleaseTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReleaseTypeType: valid values are %v", v, AllowedReleaseTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReleaseTypeType) IsValid() bool {
	for _, existing := range AllowedReleaseTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReleaseType.Type value.
func (v ReleaseTypeType) Ptr() *ReleaseTypeType {
	return &v
}
