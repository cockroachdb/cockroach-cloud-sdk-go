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

// PlanType  - DEDICATED: A paid plan that offers dedicated hardware in any location.  - CUSTOM: A plan option that is used for clusters whose machine configs are not supported in self-service. All INVOICE clusters are under this plan option.  - SERVERLESS: A paid plan that runs on shared hardware and caps the users' maximum monthly spending to a user-specified (possibly 0) amount.
type PlanType string

// List of Plan.Type.
const (
	PLANTYPE_DEDICATED  PlanType = "DEDICATED"
	PLANTYPE_CUSTOM     PlanType = "CUSTOM"
	PLANTYPE_SERVERLESS PlanType = "SERVERLESS"
)

// All allowed values of PlanType enum.
var AllowedPlanTypeEnumValues = []PlanType{
	"DEDICATED",
	"CUSTOM",
	"SERVERLESS",
}

// NewPlanTypeFromValue returns a pointer to a valid PlanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlanTypeFromValue(v string) (*PlanType, error) {
	ev := PlanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlanType: valid values are %v", v, AllowedPlanTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PlanType) IsValid() bool {
	for _, existing := range AllowedPlanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Plan.Type value.
func (v PlanType) Ptr() *PlanType {
	return &v
}
