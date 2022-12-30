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

// checks if the SecretCloudCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretCloudCredential{}

// SecretCloudCredential This model describes all Cloud Credential properties
type SecretCloudCredential struct {
	// Cloud credential ID generated by senhasegura
	Id int32 `json:"id"`
	// Cloud credential provideder
	Provider string `json:"provider"`
	// Cloud credential Access Key ID
	AccessKeyId string `json:"access_key_id"`
	// Cloud credential Secret Access Key
	SecretAccessKey string `json:"secret_access_key"`
	AdditionalProperties map[string]interface{}
}

type _SecretCloudCredential SecretCloudCredential

// NewSecretCloudCredential instantiates a new SecretCloudCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretCloudCredential(id int32, provider string, accessKeyId string, secretAccessKey string) *SecretCloudCredential {
	this := SecretCloudCredential{}
	this.Id = id
	this.Provider = provider
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewSecretCloudCredentialWithDefaults instantiates a new SecretCloudCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretCloudCredentialWithDefaults() *SecretCloudCredential {
	this := SecretCloudCredential{}
	return &this
}

// GetId returns the Id field value
func (o *SecretCloudCredential) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretCloudCredential) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretCloudCredential) SetId(v int32) {
	o.Id = v
}

// GetProvider returns the Provider field value
func (o *SecretCloudCredential) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SecretCloudCredential) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SecretCloudCredential) SetProvider(v string) {
	o.Provider = v
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *SecretCloudCredential) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *SecretCloudCredential) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *SecretCloudCredential) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *SecretCloudCredential) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *SecretCloudCredential) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *SecretCloudCredential) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

func (o SecretCloudCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretCloudCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["provider"] = o.Provider
	toSerialize["access_key_id"] = o.AccessKeyId
	toSerialize["secret_access_key"] = o.SecretAccessKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecretCloudCredential) UnmarshalJSON(bytes []byte) (err error) {
	varSecretCloudCredential := _SecretCloudCredential{}

	if err = json.Unmarshal(bytes, &varSecretCloudCredential); err == nil {
		*o = SecretCloudCredential(varSecretCloudCredential)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "access_key_id")
		delete(additionalProperties, "secret_access_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecretCloudCredential struct {
	value *SecretCloudCredential
	isSet bool
}

func (v NullableSecretCloudCredential) Get() *SecretCloudCredential {
	return v.value
}

func (v *NullableSecretCloudCredential) Set(val *SecretCloudCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretCloudCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretCloudCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretCloudCredential(val *SecretCloudCredential) *NullableSecretCloudCredential {
	return &NullableSecretCloudCredential{value: val, isSet: true}
}

func (v NullableSecretCloudCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretCloudCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

