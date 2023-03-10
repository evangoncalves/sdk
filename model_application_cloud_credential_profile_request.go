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

// checks if the ApplicationCloudCredentialProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCloudCredentialProfileRequest{}

// ApplicationCloudCredentialProfileRequest This model describes all Cloud Credential Profile properties
type ApplicationCloudCredentialProfileRequest struct {
	// Profile name registered on senhasegura Cloud IAM module
	Identifier string `json:"identifier"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationCloudCredentialProfileRequest ApplicationCloudCredentialProfileRequest

// NewApplicationCloudCredentialProfileRequest instantiates a new ApplicationCloudCredentialProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCloudCredentialProfileRequest(identifier string) *ApplicationCloudCredentialProfileRequest {
	this := ApplicationCloudCredentialProfileRequest{}
	this.Identifier = identifier
	return &this
}

// NewApplicationCloudCredentialProfileRequestWithDefaults instantiates a new ApplicationCloudCredentialProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCloudCredentialProfileRequestWithDefaults() *ApplicationCloudCredentialProfileRequest {
	this := ApplicationCloudCredentialProfileRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *ApplicationCloudCredentialProfileRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *ApplicationCloudCredentialProfileRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *ApplicationCloudCredentialProfileRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o ApplicationCloudCredentialProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCloudCredentialProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationCloudCredentialProfileRequest) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationCloudCredentialProfileRequest := _ApplicationCloudCredentialProfileRequest{}

	if err = json.Unmarshal(bytes, &varApplicationCloudCredentialProfileRequest); err == nil {
		*o = ApplicationCloudCredentialProfileRequest(varApplicationCloudCredentialProfileRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationCloudCredentialProfileRequest struct {
	value *ApplicationCloudCredentialProfileRequest
	isSet bool
}

func (v NullableApplicationCloudCredentialProfileRequest) Get() *ApplicationCloudCredentialProfileRequest {
	return v.value
}

func (v *NullableApplicationCloudCredentialProfileRequest) Set(val *ApplicationCloudCredentialProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCloudCredentialProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCloudCredentialProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCloudCredentialProfileRequest(val *ApplicationCloudCredentialProfileRequest) *NullableApplicationCloudCredentialProfileRequest {
	return &NullableApplicationCloudCredentialProfileRequest{value: val, isSet: true}
}

func (v NullableApplicationCloudCredentialProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCloudCredentialProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


