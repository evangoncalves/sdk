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

// checks if the AuthorizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationsResponse{}

// AuthorizationsResponse This model describes all Authorization properties
type AuthorizationsResponse struct {
	Response Response `json:"response"`
	Authorizations []Authorizations `json:"authorizations"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationsResponse AuthorizationsResponse

// NewAuthorizationsResponse instantiates a new AuthorizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationsResponse(response Response, authorizations []Authorizations) *AuthorizationsResponse {
	this := AuthorizationsResponse{}
	this.Response = response
	this.Authorizations = authorizations
	return &this
}

// NewAuthorizationsResponseWithDefaults instantiates a new AuthorizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationsResponseWithDefaults() *AuthorizationsResponse {
	this := AuthorizationsResponse{}
	return &this
}

// GetResponse returns the Response field value
func (o *AuthorizationsResponse) GetResponse() Response {
	if o == nil {
		var ret Response
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *AuthorizationsResponse) GetResponseOk() (*Response, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *AuthorizationsResponse) SetResponse(v Response) {
	o.Response = v
}

// GetAuthorizations returns the Authorizations field value
func (o *AuthorizationsResponse) GetAuthorizations() []Authorizations {
	if o == nil {
		var ret []Authorizations
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *AuthorizationsResponse) GetAuthorizationsOk() ([]Authorizations, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *AuthorizationsResponse) SetAuthorizations(v []Authorizations) {
	o.Authorizations = v
}

func (o AuthorizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response"] = o.Response
	toSerialize["authorizations"] = o.Authorizations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationsResponse := _AuthorizationsResponse{}

	if err = json.Unmarshal(bytes, &varAuthorizationsResponse); err == nil {
		*o = AuthorizationsResponse(varAuthorizationsResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "response")
		delete(additionalProperties, "authorizations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationsResponse struct {
	value *AuthorizationsResponse
	isSet bool
}

func (v NullableAuthorizationsResponse) Get() *AuthorizationsResponse {
	return v.value
}

func (v *NullableAuthorizationsResponse) Set(val *AuthorizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationsResponse(val *AuthorizationsResponse) *NullableAuthorizationsResponse {
	return &NullableAuthorizationsResponse{value: val, isSet: true}
}

func (v NullableAuthorizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


