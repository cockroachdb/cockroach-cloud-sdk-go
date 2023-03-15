// Copyright 2022 The Cockroach Authors
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
// API version: 2022-09-20

package client

import (
	"encoding/json"
)

// InvoiceAdjustment struct for InvoiceAdjustment.
type InvoiceAdjustment struct {
	Amount CurrencyAmount `json:"amount"`
	// name identifies the adjustment.
	Name string `json:"name"`
}

// NewInvoiceAdjustment instantiates a new InvoiceAdjustment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceAdjustment(amount CurrencyAmount, name string) *InvoiceAdjustment {
	p := InvoiceAdjustment{}
	p.Amount = amount
	p.Name = name
	return &p
}

// NewInvoiceAdjustmentWithDefaults instantiates a new InvoiceAdjustment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceAdjustmentWithDefaults() *InvoiceAdjustment {
	p := InvoiceAdjustment{}
	return &p
}

// GetAmount returns the Amount field value.
func (o *InvoiceAdjustment) GetAmount() CurrencyAmount {
	if o == nil {
		var ret CurrencyAmount
		return ret
	}

	return o.Amount
}

// SetAmount sets field value.
func (o *InvoiceAdjustment) SetAmount(v CurrencyAmount) {
	o.Amount = v
}

// GetName returns the Name field value.
func (o *InvoiceAdjustment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *InvoiceAdjustment) SetName(v string) {
	o.Name = v
}

func (o InvoiceAdjustment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}
