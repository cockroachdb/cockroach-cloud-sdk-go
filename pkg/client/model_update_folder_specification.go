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

// UpdateFolderSpecification Set `parent_id` to empty string ” or 'root' to move a folder to the root level..
type UpdateFolderSpecification struct {
	// labels are key-value pairs used to organize and categorize resources. If the labels field is included in the request: Any existing labels on the folder that are not included will be removed, and any new labels specified will be added. If the labels field is omitted from the request entirely, all existing labels will remain unchanged.
	Labels   *map[string]string `json:"labels,omitempty"`
	Name     *string            `json:"name,omitempty"`
	ParentId *string            `json:"parent_id,omitempty"`
}

// NewUpdateFolderSpecification instantiates a new UpdateFolderSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFolderSpecification() *UpdateFolderSpecification {
	p := UpdateFolderSpecification{}
	return &p
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateFolderSpecification) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *UpdateFolderSpecification) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFolderSpecification) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFolderSpecification) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *UpdateFolderSpecification) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *UpdateFolderSpecification) SetParentId(v string) {
	o.ParentId = &v
}
