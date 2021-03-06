/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InitramfsConfig struct for InitramfsConfig
type InitramfsConfig struct {
	Path string `json:"path"`
}

// NewInitramfsConfig instantiates a new InitramfsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitramfsConfig(path string) *InitramfsConfig {
	this := InitramfsConfig{}
	this.Path = path
	return &this
}

// NewInitramfsConfigWithDefaults instantiates a new InitramfsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitramfsConfigWithDefaults() *InitramfsConfig {
	this := InitramfsConfig{}
	return &this
}

// GetPath returns the Path field value
func (o *InitramfsConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *InitramfsConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *InitramfsConfig) SetPath(v string) {
	o.Path = v
}

func (o InitramfsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableInitramfsConfig struct {
	value *InitramfsConfig
	isSet bool
}

func (v NullableInitramfsConfig) Get() *InitramfsConfig {
	return v.value
}

func (v *NullableInitramfsConfig) Set(val *InitramfsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInitramfsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInitramfsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitramfsConfig(val *InitramfsConfig) *NullableInitramfsConfig {
	return &NullableInitramfsConfig{value: val, isSet: true}
}

func (v NullableInitramfsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitramfsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
