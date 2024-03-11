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

// ClientCACertInfo struct for ClientCACertInfo.
type ClientCACertInfo struct {
	Status      *ClientCACertStatus `json:"status,omitempty"`
	X509PemCert *string             `json:"x509_pem_cert,omitempty"`
}

// NewClientCACertInfo instantiates a new ClientCACertInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientCACertInfo() *ClientCACertInfo {
	p := ClientCACertInfo{}
	return &p
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClientCACertInfo) GetStatus() ClientCACertStatus {
	if o == nil || o.Status == nil {
		var ret ClientCACertStatus
		return ret
	}
	return *o.Status
}

// SetStatus gets a reference to the given ClientCACertStatus and assigns it to the Status field.
func (o *ClientCACertInfo) SetStatus(v ClientCACertStatus) {
	o.Status = &v
}

// GetX509PemCert returns the X509PemCert field value if set, zero value otherwise.
func (o *ClientCACertInfo) GetX509PemCert() string {
	if o == nil || o.X509PemCert == nil {
		var ret string
		return ret
	}
	return *o.X509PemCert
}

// SetX509PemCert gets a reference to the given string and assigns it to the X509PemCert field.
func (o *ClientCACertInfo) SetX509PemCert(v string) {
	o.X509PemCert = &v
}
