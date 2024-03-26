/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelsAddRegKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsAddRegKey{}

// ModelsAddRegKey struct for ModelsAddRegKey
type ModelsAddRegKey struct {
	// Description of the registration key.
	Description *string `json:"description,omitempty"`
	// ExpiresAt is optional, if set the registration key is only valid until the ExpiresAt time.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// SecurityGroupId is the ID of the security group to assign to the device.
	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// ServiceNetworkID is the ID of the Service Network the device can join.
	ServiceNetworkId *string `json:"service_network_id,omitempty"`
	// Settings contains general settings for the device.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// SingleUse only allows the registration key to be used once.
	SingleUse *bool `json:"single_use,omitempty"`
	// VpcID is the ID of the VPC the device will join.
	VpcId *string `json:"vpc_id,omitempty"`
}

// NewModelsAddRegKey instantiates a new ModelsAddRegKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAddRegKey() *ModelsAddRegKey {
	this := ModelsAddRegKey{}
	return &this
}

// NewModelsAddRegKeyWithDefaults instantiates a new ModelsAddRegKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAddRegKeyWithDefaults() *ModelsAddRegKey {
	this := ModelsAddRegKey{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsAddRegKey) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *ModelsAddRegKey) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetSecurityGroupId() string {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *ModelsAddRegKey) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetServiceNetworkId returns the ServiceNetworkId field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetServiceNetworkId() string {
	if o == nil || IsNil(o.ServiceNetworkId) {
		var ret string
		return ret
	}
	return *o.ServiceNetworkId
}

// GetServiceNetworkIdOk returns a tuple with the ServiceNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetServiceNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceNetworkId) {
		return nil, false
	}
	return o.ServiceNetworkId, true
}

// HasServiceNetworkId returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasServiceNetworkId() bool {
	if o != nil && !IsNil(o.ServiceNetworkId) {
		return true
	}

	return false
}

// SetServiceNetworkId gets a reference to the given string and assigns it to the ServiceNetworkId field.
func (o *ModelsAddRegKey) SetServiceNetworkId(v string) {
	o.ServiceNetworkId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *ModelsAddRegKey) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetSingleUse() bool {
	if o == nil || IsNil(o.SingleUse) {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetSingleUseOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleUse) {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasSingleUse() bool {
	if o != nil && !IsNil(o.SingleUse) {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *ModelsAddRegKey) SetSingleUse(v bool) {
	o.SingleUse = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *ModelsAddRegKey) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAddRegKey) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *ModelsAddRegKey) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *ModelsAddRegKey) SetVpcId(v string) {
	o.VpcId = &v
}

func (o ModelsAddRegKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsAddRegKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.SecurityGroupId) {
		toSerialize["security_group_id"] = o.SecurityGroupId
	}
	if !IsNil(o.ServiceNetworkId) {
		toSerialize["service_network_id"] = o.ServiceNetworkId
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.SingleUse) {
		toSerialize["single_use"] = o.SingleUse
	}
	if !IsNil(o.VpcId) {
		toSerialize["vpc_id"] = o.VpcId
	}
	return toSerialize, nil
}

type NullableModelsAddRegKey struct {
	value *ModelsAddRegKey
	isSet bool
}

func (v NullableModelsAddRegKey) Get() *ModelsAddRegKey {
	return v.value
}

func (v *NullableModelsAddRegKey) Set(val *ModelsAddRegKey) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAddRegKey) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAddRegKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAddRegKey(val *ModelsAddRegKey) *NullableModelsAddRegKey {
	return &NullableModelsAddRegKey{value: val, isSet: true}
}

func (v NullableModelsAddRegKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAddRegKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}