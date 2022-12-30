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
	"time"
)

// checks if the Secret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Secret{}

// Secret This model describes all Secret properties
type Secret struct {
	// Secret ID generated, by senhasegura
	Id int32 `json:"id"`
	// Secret name
	Name string `json:"name"`
	// Secret identifier (must be unique on the platform)
	Identifier string `json:"identifier"`
	// Engine where this secret will be used
	Engine string `json:"engine"`
	Environment string `json:"environment"`
	// Secret expiration date in UTC pattern
	Expiration *time.Time `json:"expiration,omitempty"`
	// A description about the usage of this Secret
	Description *string `json:"description,omitempty"`
	// Secret tags used for access segregation
	Tags []string `json:"tags,omitempty"`
	// Cloud IAM credentials which can be accessed through this Secret
	CloudCredentials []SecretCloudCredential `json:"cloud_credentials,omitempty"`
	// PAM Credentials which can be accessed through this Secret
	PamCredentials []SecretPAMCredential `json:"pam_credentials,omitempty"`
	// Ephemeral credentials generated by Dynamic Provisioning
	EphemeralCredentials []SecretPAMEphemeralCredential `json:"ephemeral_credentials,omitempty"`
	SshKeys []SecretPAMSSHKey `json:"ssh_keys,omitempty"`
	// This model describes Key/Value pairs properties
	KeyValues *map[string]string `json:"key_values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Secret Secret

// NewSecret instantiates a new Secret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecret(id int32, name string, identifier string, engine string, environment string) *Secret {
	this := Secret{}
	this.Id = id
	this.Name = name
	this.Identifier = identifier
	this.Engine = engine
	this.Environment = environment
	return &this
}

// NewSecretWithDefaults instantiates a new Secret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretWithDefaults() *Secret {
	this := Secret{}
	return &this
}

// GetId returns the Id field value
func (o *Secret) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Secret) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Secret) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Secret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Secret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Secret) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value
func (o *Secret) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Secret) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Secret) SetIdentifier(v string) {
	o.Identifier = v
}

// GetEngine returns the Engine field value
func (o *Secret) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *Secret) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value
func (o *Secret) SetEngine(v string) {
	o.Engine = v
}

// GetEnvironment returns the Environment field value
func (o *Secret) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Secret) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *Secret) SetEnvironment(v string) {
	o.Environment = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Secret) GetExpiration() time.Time {
	if o == nil || isNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetExpirationOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Secret) HasExpiration() bool {
	if o != nil && !isNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *Secret) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Secret) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Secret) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Secret) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Secret) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Secret) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Secret) SetTags(v []string) {
	o.Tags = v
}

// GetCloudCredentials returns the CloudCredentials field value if set, zero value otherwise.
func (o *Secret) GetCloudCredentials() []SecretCloudCredential {
	if o == nil || isNil(o.CloudCredentials) {
		var ret []SecretCloudCredential
		return ret
	}
	return o.CloudCredentials
}

// GetCloudCredentialsOk returns a tuple with the CloudCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetCloudCredentialsOk() ([]SecretCloudCredential, bool) {
	if o == nil || isNil(o.CloudCredentials) {
		return nil, false
	}
	return o.CloudCredentials, true
}

// HasCloudCredentials returns a boolean if a field has been set.
func (o *Secret) HasCloudCredentials() bool {
	if o != nil && !isNil(o.CloudCredentials) {
		return true
	}

	return false
}

// SetCloudCredentials gets a reference to the given []SecretCloudCredential and assigns it to the CloudCredentials field.
func (o *Secret) SetCloudCredentials(v []SecretCloudCredential) {
	o.CloudCredentials = v
}

// GetPamCredentials returns the PamCredentials field value if set, zero value otherwise.
func (o *Secret) GetPamCredentials() []SecretPAMCredential {
	if o == nil || isNil(o.PamCredentials) {
		var ret []SecretPAMCredential
		return ret
	}
	return o.PamCredentials
}

// GetPamCredentialsOk returns a tuple with the PamCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetPamCredentialsOk() ([]SecretPAMCredential, bool) {
	if o == nil || isNil(o.PamCredentials) {
		return nil, false
	}
	return o.PamCredentials, true
}

// HasPamCredentials returns a boolean if a field has been set.
func (o *Secret) HasPamCredentials() bool {
	if o != nil && !isNil(o.PamCredentials) {
		return true
	}

	return false
}

// SetPamCredentials gets a reference to the given []SecretPAMCredential and assigns it to the PamCredentials field.
func (o *Secret) SetPamCredentials(v []SecretPAMCredential) {
	o.PamCredentials = v
}

// GetEphemeralCredentials returns the EphemeralCredentials field value if set, zero value otherwise.
func (o *Secret) GetEphemeralCredentials() []SecretPAMEphemeralCredential {
	if o == nil || isNil(o.EphemeralCredentials) {
		var ret []SecretPAMEphemeralCredential
		return ret
	}
	return o.EphemeralCredentials
}

// GetEphemeralCredentialsOk returns a tuple with the EphemeralCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetEphemeralCredentialsOk() ([]SecretPAMEphemeralCredential, bool) {
	if o == nil || isNil(o.EphemeralCredentials) {
		return nil, false
	}
	return o.EphemeralCredentials, true
}

// HasEphemeralCredentials returns a boolean if a field has been set.
func (o *Secret) HasEphemeralCredentials() bool {
	if o != nil && !isNil(o.EphemeralCredentials) {
		return true
	}

	return false
}

// SetEphemeralCredentials gets a reference to the given []SecretPAMEphemeralCredential and assigns it to the EphemeralCredentials field.
func (o *Secret) SetEphemeralCredentials(v []SecretPAMEphemeralCredential) {
	o.EphemeralCredentials = v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Secret) GetSshKeys() []SecretPAMSSHKey {
	if o == nil || isNil(o.SshKeys) {
		var ret []SecretPAMSSHKey
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetSshKeysOk() ([]SecretPAMSSHKey, bool) {
	if o == nil || isNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Secret) HasSshKeys() bool {
	if o != nil && !isNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []SecretPAMSSHKey and assigns it to the SshKeys field.
func (o *Secret) SetSshKeys(v []SecretPAMSSHKey) {
	o.SshKeys = v
}

// GetKeyValues returns the KeyValues field value if set, zero value otherwise.
func (o *Secret) GetKeyValues() map[string]string {
	if o == nil || isNil(o.KeyValues) {
		var ret map[string]string
		return ret
	}
	return *o.KeyValues
}

// GetKeyValuesOk returns a tuple with the KeyValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetKeyValuesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.KeyValues) {
		return nil, false
	}
	return o.KeyValues, true
}

// HasKeyValues returns a boolean if a field has been set.
func (o *Secret) HasKeyValues() bool {
	if o != nil && !isNil(o.KeyValues) {
		return true
	}

	return false
}

// SetKeyValues gets a reference to the given map[string]string and assigns it to the KeyValues field.
func (o *Secret) SetKeyValues(v map[string]string) {
	o.KeyValues = &v
}

func (o Secret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Secret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["identifier"] = o.Identifier
	toSerialize["engine"] = o.Engine
	toSerialize["environment"] = o.Environment
	if !isNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.CloudCredentials) {
		toSerialize["cloud_credentials"] = o.CloudCredentials
	}
	if !isNil(o.PamCredentials) {
		toSerialize["pam_credentials"] = o.PamCredentials
	}
	if !isNil(o.EphemeralCredentials) {
		toSerialize["ephemeral_credentials"] = o.EphemeralCredentials
	}
	if !isNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !isNil(o.KeyValues) {
		toSerialize["key_values"] = o.KeyValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Secret) UnmarshalJSON(bytes []byte) (err error) {
	varSecret := _Secret{}

	if err = json.Unmarshal(bytes, &varSecret); err == nil {
		*o = Secret(varSecret)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "engine")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "cloud_credentials")
		delete(additionalProperties, "pam_credentials")
		delete(additionalProperties, "ephemeral_credentials")
		delete(additionalProperties, "ssh_keys")
		delete(additionalProperties, "key_values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecret struct {
	value *Secret
	isSet bool
}

func (v NullableSecret) Get() *Secret {
	return v.value
}

func (v *NullableSecret) Set(val *Secret) {
	v.value = val
	v.isSet = true
}

func (v NullableSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecret(val *Secret) *NullableSecret {
	return &NullableSecret{value: val, isSet: true}
}

func (v NullableSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


