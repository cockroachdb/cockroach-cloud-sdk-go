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

// CurrencyType the model 'CurrencyType'.
type CurrencyType string

// List of Currency.Type.
const (
	CURRENCYTYPE_USD                CurrencyType = "USD"
	CURRENCYTYPE_CRDB_CLOUD_CREDITS CurrencyType = "CRDB_CLOUD_CREDITS"
)

// All allowed values of CurrencyType enum.
var AllowedCurrencyTypeEnumValues = []CurrencyType{
	"USD",
	"CRDB_CLOUD_CREDITS",
}

// NewCurrencyTypeFromValue returns a pointer to a valid CurrencyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCurrencyTypeFromValue(v string) (*CurrencyType, error) {
	ev := CurrencyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CurrencyType: valid values are %v", v, AllowedCurrencyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CurrencyType) IsValid() bool {
	for _, existing := range AllowedCurrencyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Currency.Type value.
func (v CurrencyType) Ptr() *CurrencyType {
	return &v
}
