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

// LineItem struct for LineItem.
type LineItem struct {
	// Description contains the details of the line item (i.e t3 micro).
	Description string `json:"description"`
	// Quantity is the number of the specific line items used.
	Quantity float64 `json:"quantity"`
	// UnitCost is the cost per unit of line item.
	UnitCost             float64          `json:"unit_cost"`
	Total                CurrencyAmount   `json:"total"`
	QuantityUnit         QuantityUnitType `json:"quantity_unit"`
	AdditionalProperties map[string]interface{}
}

type lineItem LineItem

// NewLineItem instantiates a new LineItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItem(description string, quantity float64, unitCost float64, total CurrencyAmount, quantityUnit QuantityUnitType) *LineItem {
	p := LineItem{}
	p.Description = description
	p.Quantity = quantity
	p.UnitCost = unitCost
	p.Total = total
	p.QuantityUnit = quantityUnit
	return &p
}

// NewLineItemWithDefaults instantiates a new LineItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemWithDefaults() *LineItem {
	p := LineItem{}
	return &p
}

// GetDescription returns the Description field value.
func (o *LineItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value.
func (o *LineItem) SetDescription(v string) {
	o.Description = v
}

// GetQuantity returns the Quantity field value.
func (o *LineItem) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// SetQuantity sets field value.
func (o *LineItem) SetQuantity(v float64) {
	o.Quantity = v
}

// GetUnitCost returns the UnitCost field value.
func (o *LineItem) GetUnitCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UnitCost
}

// SetUnitCost sets field value.
func (o *LineItem) SetUnitCost(v float64) {
	o.UnitCost = v
}

// GetTotal returns the Total field value.
func (o *LineItem) GetTotal() CurrencyAmount {
	if o == nil {
		var ret CurrencyAmount
		return ret
	}

	return o.Total
}

// SetTotal sets field value.
func (o *LineItem) SetTotal(v CurrencyAmount) {
	o.Total = v
}

// GetQuantityUnit returns the QuantityUnit field value.
func (o *LineItem) GetQuantityUnit() QuantityUnitType {
	if o == nil {
		var ret QuantityUnitType
		return ret
	}

	return o.QuantityUnit
}

// SetQuantityUnit sets field value.
func (o *LineItem) SetQuantityUnit(v QuantityUnitType) {
	o.QuantityUnit = v
}

func (o LineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["unit_cost"] = o.UnitCost
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["quantity_unit"] = o.QuantityUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LineItem) UnmarshalJSON(bytes []byte) (err error) {
	varLineItem := lineItem{}

	if err = json.Unmarshal(bytes, &varLineItem); err == nil {
		*o = LineItem(varLineItem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unit_cost")
		delete(additionalProperties, "total")
		delete(additionalProperties, "quantity_unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}