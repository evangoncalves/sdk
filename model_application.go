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

// checks if the Application type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Application{}

// Application This model describes all application properties
type Application struct {
	// Application ID, generated by senhasegura
	Id int32 `json:"id"`
	// Name of the application
	Name string `json:"name"`
	// A descriptive text for the application
	Description *string `json:"description,omitempty"`
	// Authentication method used by this application's authorizations
	AuthMethod string `json:"auth_method"`
	// Line of business that best describe this application usage
	LineOfBusiness *string `json:"line_of_business,omitempty"`
	// Type of application
	Type *string `json:"type,omitempty"`
	// List of tags used for access segregation
	Tags []string `json:"tags,omitempty"`
	// Cloud profiles used for Dynamic Provisioning on Secrets
	CloudProfiles []ApplicationCloudCredentialProfile `json:"cloud_profiles,omitempty"`
	// PAM profiles used for Dynamic Provisioning on Secrets
	PamProfiles []ApplicationPAMCredentialProfile `json:"pam_profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Application Application

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(id int32, name string, authMethod string) *Application {
	this := Application{}
	this.Id = id
	this.Name = name
	this.AuthMethod = authMethod
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetId returns the Id field value
func (o *Application) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Application) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Application) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description = &v
}

// GetAuthMethod returns the AuthMethod field value
func (o *Application) GetAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value
// and a boolean to check if the value has been set.
func (o *Application) GetAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMethod, true
}

// SetAuthMethod sets field value
func (o *Application) SetAuthMethod(v string) {
	o.AuthMethod = v
}

// GetLineOfBusiness returns the LineOfBusiness field value if set, zero value otherwise.
func (o *Application) GetLineOfBusiness() string {
	if o == nil || isNil(o.LineOfBusiness) {
		var ret string
		return ret
	}
	return *o.LineOfBusiness
}

// GetLineOfBusinessOk returns a tuple with the LineOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLineOfBusinessOk() (*string, bool) {
	if o == nil || isNil(o.LineOfBusiness) {
		return nil, false
	}
	return o.LineOfBusiness, true
}

// HasLineOfBusiness returns a boolean if a field has been set.
func (o *Application) HasLineOfBusiness() bool {
	if o != nil && !isNil(o.LineOfBusiness) {
		return true
	}

	return false
}

// SetLineOfBusiness gets a reference to the given string and assigns it to the LineOfBusiness field.
func (o *Application) SetLineOfBusiness(v string) {
	o.LineOfBusiness = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Application) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Application) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Application) SetType(v string) {
	o.Type = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Application) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Application) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Application) SetTags(v []string) {
	o.Tags = v
}

// GetCloudProfiles returns the CloudProfiles field value if set, zero value otherwise.
func (o *Application) GetCloudProfiles() []ApplicationCloudCredentialProfile {
	if o == nil || isNil(o.CloudProfiles) {
		var ret []ApplicationCloudCredentialProfile
		return ret
	}
	return o.CloudProfiles
}

// GetCloudProfilesOk returns a tuple with the CloudProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCloudProfilesOk() ([]ApplicationCloudCredentialProfile, bool) {
	if o == nil || isNil(o.CloudProfiles) {
		return nil, false
	}
	return o.CloudProfiles, true
}

// HasCloudProfiles returns a boolean if a field has been set.
func (o *Application) HasCloudProfiles() bool {
	if o != nil && !isNil(o.CloudProfiles) {
		return true
	}

	return false
}

// SetCloudProfiles gets a reference to the given []ApplicationCloudCredentialProfile and assigns it to the CloudProfiles field.
func (o *Application) SetCloudProfiles(v []ApplicationCloudCredentialProfile) {
	o.CloudProfiles = v
}

// GetPamProfiles returns the PamProfiles field value if set, zero value otherwise.
func (o *Application) GetPamProfiles() []ApplicationPAMCredentialProfile {
	if o == nil || isNil(o.PamProfiles) {
		var ret []ApplicationPAMCredentialProfile
		return ret
	}
	return o.PamProfiles
}

// GetPamProfilesOk returns a tuple with the PamProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetPamProfilesOk() ([]ApplicationPAMCredentialProfile, bool) {
	if o == nil || isNil(o.PamProfiles) {
		return nil, false
	}
	return o.PamProfiles, true
}

// HasPamProfiles returns a boolean if a field has been set.
func (o *Application) HasPamProfiles() bool {
	if o != nil && !isNil(o.PamProfiles) {
		return true
	}

	return false
}

// SetPamProfiles gets a reference to the given []ApplicationPAMCredentialProfile and assigns it to the PamProfiles field.
func (o *Application) SetPamProfiles(v []ApplicationPAMCredentialProfile) {
	o.PamProfiles = v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Application) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["auth_method"] = o.AuthMethod
	if !isNil(o.LineOfBusiness) {
		toSerialize["line_of_business"] = o.LineOfBusiness
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.CloudProfiles) {
		toSerialize["cloud_profiles"] = o.CloudProfiles
	}
	if !isNil(o.PamProfiles) {
		toSerialize["pam_profiles"] = o.PamProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Application) UnmarshalJSON(bytes []byte) (err error) {
	varApplication := _Application{}

	if err = json.Unmarshal(bytes, &varApplication); err == nil {
		*o = Application(varApplication)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "auth_method")
		delete(additionalProperties, "line_of_business")
		delete(additionalProperties, "type")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "cloud_profiles")
		delete(additionalProperties, "pam_profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

