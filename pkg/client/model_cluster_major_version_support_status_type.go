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
	"encoding/json"
	"fmt"
)

// ClusterMajorVersionSupportStatusType the model 'ClusterMajorVersionSupportStatusType'.
type ClusterMajorVersionSupportStatusType string

// List of ClusterMajorVersionSupportStatus.Type.
const (
	CLUSTERMAJORVERSIONSUPPORTSTATUSTYPE_UNSUPPORTED ClusterMajorVersionSupportStatusType = "UNSUPPORTED"
	CLUSTERMAJORVERSIONSUPPORTSTATUSTYPE_SUPPORTED   ClusterMajorVersionSupportStatusType = "SUPPORTED"
	CLUSTERMAJORVERSIONSUPPORTSTATUSTYPE_PREVIEW     ClusterMajorVersionSupportStatusType = "PREVIEW"
)

// All allowed values of ClusterMajorVersionSupportStatusType enum.
var AllowedClusterMajorVersionSupportStatusTypeEnumValues = []ClusterMajorVersionSupportStatusType{
	"UNSUPPORTED",
	"SUPPORTED",
	"PREVIEW",
}

func (v *ClusterMajorVersionSupportStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterMajorVersionSupportStatusType(value)
	for _, existing := range AllowedClusterMajorVersionSupportStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterMajorVersionSupportStatusType", value)
}

// NewClusterMajorVersionSupportStatusTypeFromValue returns a pointer to a valid ClusterMajorVersionSupportStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterMajorVersionSupportStatusTypeFromValue(v string) (*ClusterMajorVersionSupportStatusType, error) {
	ev := ClusterMajorVersionSupportStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterMajorVersionSupportStatusType: valid values are %v", v, AllowedClusterMajorVersionSupportStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterMajorVersionSupportStatusType) IsValid() bool {
	for _, existing := range AllowedClusterMajorVersionSupportStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterMajorVersionSupportStatus.Type value.
func (v ClusterMajorVersionSupportStatusType) Ptr() *ClusterMajorVersionSupportStatusType {
	return &v
}