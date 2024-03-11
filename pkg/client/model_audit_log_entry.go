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
	"time"
)

// AuditLogEntry AuditLogEntry represents an entry in the cloud event log. Note that this message definition should always match exactly with the corresponding `AuditLogEntry` message in `console/consolepb/console.proto`..
type AuditLogEntry struct {
	Action *AuditLogAction `json:"action,omitempty"`
	// ClusterId is the ID of the cluster to which this log entry applies, if it applies to a single cluster.
	ClusterId *string `json:"cluster_id,omitempty"`
	// ClusterName is the name of the cluster to which this log entry applies, if it applies to a single cluster.
	ClusterName *string `json:"cluster_name,omitempty"`
	// CreatedAt is the time that this log entry was recorded.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Error is the error that applies to this entry if it represents a failure.
	Error *string `json:"error,omitempty"`
	// Id uniquely identifies this entry.
	Id       *string           `json:"id,omitempty"`
	Metadata *AuditLogMetadata `json:"metadata,omitempty"`
	// Payload is a representation of the essential details relating to this log entry.
	Payload *map[string]interface{} `json:"payload,omitempty"`
	// ServiceAccountName is the name of the service account that triggered this log entry. If it was not a service account, it will be empty.
	ServiceAccountName *string `json:"service_account_name,omitempty"`
	// SessionId is an ID that can be used to correlate this log entry with others that are emitted as part of the same user session, typically for users interacting through the UI. It should be treated as an opaque string with no guaranteed structure.
	SessionId *string         `json:"session_id,omitempty"`
	Source    *AuditLogSource `json:"source,omitempty"`
	// TraceId is an ID that can be used to correlate this log entry with others that are emitted as part of the same process. It should be treated as an opaque string with no guaranteed structure.
	TraceId *string `json:"trace_id,omitempty"`
	// UserEmail is the email address of the user that triggered this log entry. If it was not a human user, it will be empty.
	UserEmail *string `json:"user_email,omitempty"`
}

// NewAuditLogEntry instantiates a new AuditLogEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogEntry() *AuditLogEntry {
	p := AuditLogEntry{}
	return &p
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AuditLogEntry) GetAction() AuditLogAction {
	if o == nil || o.Action == nil {
		var ret AuditLogAction
		return ret
	}
	return *o.Action
}

// SetAction gets a reference to the given AuditLogAction and assigns it to the Action field.
func (o *AuditLogEntry) SetAction(v AuditLogAction) {
	o.Action = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *AuditLogEntry) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *AuditLogEntry) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AuditLogEntry) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AuditLogEntry) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuditLogEntry) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuditLogEntry) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AuditLogEntry) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AuditLogEntry) SetError(v string) {
	o.Error = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditLogEntry) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuditLogEntry) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AuditLogEntry) GetMetadata() AuditLogMetadata {
	if o == nil || o.Metadata == nil {
		var ret AuditLogMetadata
		return ret
	}
	return *o.Metadata
}

// SetMetadata gets a reference to the given AuditLogMetadata and assigns it to the Metadata field.
func (o *AuditLogEntry) SetMetadata(v AuditLogMetadata) {
	o.Metadata = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AuditLogEntry) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Payload
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *AuditLogEntry) SetPayload(v map[string]interface{}) {
	o.Payload = &v
}

// GetServiceAccountName returns the ServiceAccountName field value if set, zero value otherwise.
func (o *AuditLogEntry) GetServiceAccountName() string {
	if o == nil || o.ServiceAccountName == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountName
}

// SetServiceAccountName gets a reference to the given string and assigns it to the ServiceAccountName field.
func (o *AuditLogEntry) SetServiceAccountName(v string) {
	o.ServiceAccountName = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *AuditLogEntry) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *AuditLogEntry) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AuditLogEntry) GetSource() AuditLogSource {
	if o == nil || o.Source == nil {
		var ret AuditLogSource
		return ret
	}
	return *o.Source
}

// SetSource gets a reference to the given AuditLogSource and assigns it to the Source field.
func (o *AuditLogEntry) SetSource(v AuditLogSource) {
	o.Source = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *AuditLogEntry) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *AuditLogEntry) SetTraceId(v string) {
	o.TraceId = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *AuditLogEntry) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *AuditLogEntry) SetUserEmail(v string) {
	o.UserEmail = &v
}
