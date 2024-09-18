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

// OperatingSystemType the model 'OperatingSystemType'.
type OperatingSystemType string

// List of OperatingSystemType.
const (
	OPERATINGSYSTEMTYPE_MAC     OperatingSystemType = "MAC"
	OPERATINGSYSTEMTYPE_LINUX   OperatingSystemType = "LINUX"
	OPERATINGSYSTEMTYPE_WINDOWS OperatingSystemType = "WINDOWS"
)

// All allowed values of OperatingSystemType enum.
var AllowedOperatingSystemTypeEnumValues = []OperatingSystemType{
	"MAC",
	"LINUX",
	"WINDOWS",
}

// NewOperatingSystemTypeFromValue returns a pointer to a valid OperatingSystemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOperatingSystemTypeFromValue(v string) (*OperatingSystemType, error) {
	ev := OperatingSystemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperatingSystemType: valid values are %v", v, AllowedOperatingSystemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OperatingSystemType) IsValid() bool {
	for _, existing := range AllowedOperatingSystemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperatingSystemType value.
func (v OperatingSystemType) Ptr() *OperatingSystemType {
	return &v
}
