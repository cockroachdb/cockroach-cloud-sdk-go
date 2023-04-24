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
)

// AWSPrivateLinkServiceDetail struct for AWSPrivateLinkServiceDetail.
type AWSPrivateLinkServiceDetail struct {
	// availability_zone_ids are the identifiers for the availability zones that the service is available in.
	AvailabilityZoneIds []string `json:"availability_zone_ids"`
	// service_id is the server side of the PrivateLink connection. This is the same as AWSPrivateLinkEndpoint.service_id.
	ServiceId string `json:"service_id,string"`
	// service_name is the AWS service name customers use to create endpoints on their end.
	ServiceName string `json:"service_name,string"`
}

// NewAWSPrivateLinkServiceDetail instantiates a new AWSPrivateLinkServiceDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSPrivateLinkServiceDetail(availabilityZoneIds []string, serviceId string, serviceName string) *AWSPrivateLinkServiceDetail {
	p := AWSPrivateLinkServiceDetail{}
	p.AvailabilityZoneIds = availabilityZoneIds
	p.ServiceId = serviceId
	p.ServiceName = serviceName
	return &p
}

// NewAWSPrivateLinkServiceDetailWithDefaults instantiates a new AWSPrivateLinkServiceDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSPrivateLinkServiceDetailWithDefaults() *AWSPrivateLinkServiceDetail {
	p := AWSPrivateLinkServiceDetail{}
	return &p
}

// GetAvailabilityZoneIds returns the AvailabilityZoneIds field value.
func (o *AWSPrivateLinkServiceDetail) GetAvailabilityZoneIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailabilityZoneIds
}

// SetAvailabilityZoneIds sets field value.
func (o *AWSPrivateLinkServiceDetail) SetAvailabilityZoneIds(v []string) {
	o.AvailabilityZoneIds = v
}

// GetServiceId returns the ServiceId field value.
func (o *AWSPrivateLinkServiceDetail) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// SetServiceId sets field value.
func (o *AWSPrivateLinkServiceDetail) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceName returns the ServiceName field value.
func (o *AWSPrivateLinkServiceDetail) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// SetServiceName sets field value.
func (o *AWSPrivateLinkServiceDetail) SetServiceName(v string) {
	o.ServiceName = v
}

func (o AWSPrivateLinkServiceDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["availability_zone_ids"] = o.AvailabilityZoneIds
	}
	if true {
		toSerialize["service_id"] = o.ServiceId
	}
	if true {
		toSerialize["service_name"] = o.ServiceName
	}
	return json.Marshal(toSerialize)
}
