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

// CMEKKeyType the model 'CMEKKeyType'.
type CMEKKeyType string

// List of CMEKKeyType.
const (
	CMEKKEYTYPE_AWS_KMS       CMEKKeyType = "AWS_KMS"
	CMEKKEYTYPE_GCP_CLOUD_KMS CMEKKeyType = "GCP_CLOUD_KMS"
	CMEKKEYTYPE_NULL_KMS      CMEKKeyType = "NULL_KMS"
)

// All allowed values of CMEKKeyType enum.
var AllowedCMEKKeyTypeEnumValues = []CMEKKeyType{
	"AWS_KMS",
	"GCP_CLOUD_KMS",
	"NULL_KMS",
}

// NewCMEKKeyTypeFromValue returns a pointer to a valid CMEKKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCMEKKeyTypeFromValue(v string) (*CMEKKeyType, error) {
	ev := CMEKKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CMEKKeyType: valid values are %v", v, AllowedCMEKKeyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CMEKKeyType) IsValid() bool {
	for _, existing := range AllowedCMEKKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CMEKKeyType value.
func (v CMEKKeyType) Ptr() *CMEKKeyType {
	return &v
}
