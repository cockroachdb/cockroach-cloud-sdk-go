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

// ReplicationStreamList struct for ReplicationStreamList.
type ReplicationStreamList struct {
	Pagination *KeysetPaginationResponse `json:"pagination,omitempty"`
	// replication_streams is a list of ReplicationStreams.
	ReplicationStreams []ReplicationStream `json:"replication_streams"`
}

// NewReplicationStreamList instantiates a new ReplicationStreamList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationStreamList(replicationStreams []ReplicationStream) *ReplicationStreamList {
	p := ReplicationStreamList{}
	p.ReplicationStreams = replicationStreams
	return &p
}

// NewReplicationStreamListWithDefaults instantiates a new ReplicationStreamList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationStreamListWithDefaults() *ReplicationStreamList {
	p := ReplicationStreamList{}
	return &p
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ReplicationStreamList) GetPagination() KeysetPaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret KeysetPaginationResponse
		return ret
	}
	return *o.Pagination
}

// SetPagination gets a reference to the given KeysetPaginationResponse and assigns it to the Pagination field.
func (o *ReplicationStreamList) SetPagination(v KeysetPaginationResponse) {
	o.Pagination = &v
}

// GetReplicationStreams returns the ReplicationStreams field value.
func (o *ReplicationStreamList) GetReplicationStreams() []ReplicationStream {
	if o == nil {
		var ret []ReplicationStream
		return ret
	}

	return o.ReplicationStreams
}

// SetReplicationStreams sets field value.
func (o *ReplicationStreamList) SetReplicationStreams(v []ReplicationStream) {
	o.ReplicationStreams = v
}
