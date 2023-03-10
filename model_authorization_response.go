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

// checks if the AuthorizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationResponse{}

// AuthorizationResponse This model describes all Authorization properties
type AuthorizationResponse struct {
	Response Response `json:"response"`
	Authorization Authorization `json:"authorization"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationResponse AuthorizationResponse

// NewAuthorizationResponse instantiates a new AuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationResponse(response Response, authorization Authorization) *AuthorizationResponse {
	this := AuthorizationResponse{}
	this.Response = response
	this.Authorization = authorization
	return &this
}

// NewAuthorizationResponseWithDefaults instantiates a new AuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationResponseWithDefaults() *AuthorizationResponse {
	this := AuthorizationResponse{}
	return &this
}

// GetResponse returns the Response field value
func (o *AuthorizationResponse) GetResponse() Response {
	if o == nil {
		var ret Response
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetResponseOk() (*Response, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *AuthorizationResponse) SetResponse(v Response) {
	o.Response = v
}

// GetAuthorization returns the Authorization field value
func (o *AuthorizationResponse) GetAuthorization() Authorization {
	if o == nil {
		var ret Authorization
		return ret
	}

	return o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value
// and a boolean to check if the value has been set.
func (o *AuthorizationResponse) GetAuthorizationOk() (*Authorization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorization, true
}

// SetAuthorization sets field value
func (o *AuthorizationResponse) SetAuthorization(v Authorization) {
	o.Authorization = v
}

func (o AuthorizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response"] = o.Response
	toSerialize["authorization"] = o.Authorization

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationResponse := _AuthorizationResponse{}

	if err = json.Unmarshal(bytes, &varAuthorizationResponse); err == nil {
		*o = AuthorizationResponse(varAuthorizationResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "response")
		delete(additionalProperties, "authorization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationResponse struct {
	value *AuthorizationResponse
	isSet bool
}

func (v NullableAuthorizationResponse) Get() *AuthorizationResponse {
	return v.value
}

func (v *NullableAuthorizationResponse) Set(val *AuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationResponse(val *AuthorizationResponse) *NullableAuthorizationResponse {
	return &NullableAuthorizationResponse{value: val, isSet: true}
}

func (v NullableAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


