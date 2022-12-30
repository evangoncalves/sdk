/*
senhasegura DSM API Reference

This is a senhasegura new API design

API version: 2.0.22
Contact: egoncalves@senhasegura.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the SecretPAMDeviceConnectivitiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretPAMDeviceConnectivitiesInner{}

// SecretPAMDeviceConnectivitiesInner struct for SecretPAMDeviceConnectivitiesInner
type SecretPAMDeviceConnectivitiesInner struct {
	Protocol *string `json:"protocol,omitempty"`
	Port *string `json:"port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecretPAMDeviceConnectivitiesInner SecretPAMDeviceConnectivitiesInner

// NewSecretPAMDeviceConnectivitiesInner instantiates a new SecretPAMDeviceConnectivitiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretPAMDeviceConnectivitiesInner() *SecretPAMDeviceConnectivitiesInner {
	this := SecretPAMDeviceConnectivitiesInner{}
	return &this
}

// NewSecretPAMDeviceConnectivitiesInnerWithDefaults instantiates a new SecretPAMDeviceConnectivitiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretPAMDeviceConnectivitiesInnerWithDefaults() *SecretPAMDeviceConnectivitiesInner {
	this := SecretPAMDeviceConnectivitiesInner{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SecretPAMDeviceConnectivitiesInner) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretPAMDeviceConnectivitiesInner) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SecretPAMDeviceConnectivitiesInner) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SecretPAMDeviceConnectivitiesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SecretPAMDeviceConnectivitiesInner) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretPAMDeviceConnectivitiesInner) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SecretPAMDeviceConnectivitiesInner) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *SecretPAMDeviceConnectivitiesInner) SetPort(v string) {
	o.Port = &v
}

func (o SecretPAMDeviceConnectivitiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretPAMDeviceConnectivitiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecretPAMDeviceConnectivitiesInner) UnmarshalJSON(bytes []byte) (err error) {
	varSecretPAMDeviceConnectivitiesInner := _SecretPAMDeviceConnectivitiesInner{}

	if err = json.Unmarshal(bytes, &varSecretPAMDeviceConnectivitiesInner); err == nil {
		*o = SecretPAMDeviceConnectivitiesInner(varSecretPAMDeviceConnectivitiesInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecretPAMDeviceConnectivitiesInner struct {
	value *SecretPAMDeviceConnectivitiesInner
	isSet bool
}

func (v NullableSecretPAMDeviceConnectivitiesInner) Get() *SecretPAMDeviceConnectivitiesInner {
	return v.value
}

func (v *NullableSecretPAMDeviceConnectivitiesInner) Set(val *SecretPAMDeviceConnectivitiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretPAMDeviceConnectivitiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretPAMDeviceConnectivitiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretPAMDeviceConnectivitiesInner(val *SecretPAMDeviceConnectivitiesInner) *NullableSecretPAMDeviceConnectivitiesInner {
	return &NullableSecretPAMDeviceConnectivitiesInner{value: val, isSet: true}
}

func (v NullableSecretPAMDeviceConnectivitiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretPAMDeviceConnectivitiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

