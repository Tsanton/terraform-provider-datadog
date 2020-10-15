/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// TeamsResponse Response with a list of team payloads.
type TeamsResponse struct {
	// An array of teams.
	Data []TeamResponseData `json:"data"`
	// Included related resources which the user requested.
	Included *[]TeamIncludedItems  `json:"included,omitempty"`
	Meta     *ServicesResponseMeta `json:"meta,omitempty"`
}

// NewTeamsResponse instantiates a new TeamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsResponse(data []TeamResponseData) *TeamsResponse {
	this := TeamsResponse{}
	this.Data = data
	return &this
}

// NewTeamsResponseWithDefaults instantiates a new TeamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsResponseWithDefaults() *TeamsResponse {
	this := TeamsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TeamsResponse) GetData() []TeamResponseData {
	if o == nil {
		var ret []TeamResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetDataOk() (*[]TeamResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TeamsResponse) SetData(v []TeamResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TeamsResponse) GetIncluded() []TeamIncludedItems {
	if o == nil || o.Included == nil {
		var ret []TeamIncludedItems
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetIncludedOk() (*[]TeamIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TeamsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []TeamIncludedItems and assigns it to the Included field.
func (o *TeamsResponse) SetIncluded(v []TeamIncludedItems) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TeamsResponse) GetMeta() ServicesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ServicesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetMetaOk() (*ServicesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TeamsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ServicesResponseMeta and assigns it to the Meta field.
func (o *TeamsResponse) SetMeta(v ServicesResponseMeta) {
	o.Meta = &v
}

func (o TeamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableTeamsResponse struct {
	value *TeamsResponse
	isSet bool
}

func (v NullableTeamsResponse) Get() *TeamsResponse {
	return v.value
}

func (v *NullableTeamsResponse) Set(val *TeamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsResponse(val *TeamsResponse) *NullableTeamsResponse {
	return &NullableTeamsResponse{value: val, isSet: true}
}

func (v NullableTeamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}