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

// AllowlistEntry struct for AllowlistEntry.
type AllowlistEntry struct {
	CidrIp   string  `json:"cidr_ip"`
	CidrMask int32   `json:"cidr_mask"`
	Name     *string `json:"name,omitempty"`
	Sql      bool    `json:"sql"`
	Ui       bool    `json:"ui"`
}

// NewAllowlistEntry instantiates a new AllowlistEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowlistEntry(cidrIp string, cidrMask int32, sql bool, ui bool) *AllowlistEntry {
	p := AllowlistEntry{}
	p.CidrIp = cidrIp
	p.CidrMask = cidrMask
	p.Sql = sql
	p.Ui = ui
	return &p
}

// NewAllowlistEntryWithDefaults instantiates a new AllowlistEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowlistEntryWithDefaults() *AllowlistEntry {
	p := AllowlistEntry{}
	return &p
}

// GetCidrIp returns the CidrIp field value.
func (o *AllowlistEntry) GetCidrIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrIp
}

// SetCidrIp sets field value.
func (o *AllowlistEntry) SetCidrIp(v string) {
	o.CidrIp = v
}

// GetCidrMask returns the CidrMask field value.
func (o *AllowlistEntry) GetCidrMask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CidrMask
}

// SetCidrMask sets field value.
func (o *AllowlistEntry) SetCidrMask(v int32) {
	o.CidrMask = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllowlistEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllowlistEntry) SetName(v string) {
	o.Name = &v
}

// GetSql returns the Sql field value.
func (o *AllowlistEntry) GetSql() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sql
}

// SetSql sets field value.
func (o *AllowlistEntry) SetSql(v bool) {
	o.Sql = v
}

// GetUi returns the Ui field value.
func (o *AllowlistEntry) GetUi() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ui
}

// SetUi sets field value.
func (o *AllowlistEntry) SetUi(v bool) {
	o.Ui = v
}
