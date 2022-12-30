# SecretPAMCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UsernameLabel** | Pointer to **string** |  | [optional] 
**PasswordLabel** | Pointer to **string** |  | [optional] 
**AdditionalInformationLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretPAMCredentialRequest

`func NewSecretPAMCredentialRequest(id int32, ) *SecretPAMCredentialRequest`

NewSecretPAMCredentialRequest instantiates a new SecretPAMCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretPAMCredentialRequestWithDefaults

`func NewSecretPAMCredentialRequestWithDefaults() *SecretPAMCredentialRequest`

NewSecretPAMCredentialRequestWithDefaults instantiates a new SecretPAMCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretPAMCredentialRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretPAMCredentialRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretPAMCredentialRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUsernameLabel

`func (o *SecretPAMCredentialRequest) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *SecretPAMCredentialRequest) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *SecretPAMCredentialRequest) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.

### HasUsernameLabel

`func (o *SecretPAMCredentialRequest) HasUsernameLabel() bool`

HasUsernameLabel returns a boolean if a field has been set.

### GetPasswordLabel

`func (o *SecretPAMCredentialRequest) GetPasswordLabel() string`

GetPasswordLabel returns the PasswordLabel field if non-nil, zero value otherwise.

### GetPasswordLabelOk

`func (o *SecretPAMCredentialRequest) GetPasswordLabelOk() (*string, bool)`

GetPasswordLabelOk returns a tuple with the PasswordLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLabel

`func (o *SecretPAMCredentialRequest) SetPasswordLabel(v string)`

SetPasswordLabel sets PasswordLabel field to given value.

### HasPasswordLabel

`func (o *SecretPAMCredentialRequest) HasPasswordLabel() bool`

HasPasswordLabel returns a boolean if a field has been set.

### GetAdditionalInformationLabel

`func (o *SecretPAMCredentialRequest) GetAdditionalInformationLabel() string`

GetAdditionalInformationLabel returns the AdditionalInformationLabel field if non-nil, zero value otherwise.

### GetAdditionalInformationLabelOk

`func (o *SecretPAMCredentialRequest) GetAdditionalInformationLabelOk() (*string, bool)`

GetAdditionalInformationLabelOk returns a tuple with the AdditionalInformationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformationLabel

`func (o *SecretPAMCredentialRequest) SetAdditionalInformationLabel(v string)`

SetAdditionalInformationLabel sets AdditionalInformationLabel field to given value.

### HasAdditionalInformationLabel

`func (o *SecretPAMCredentialRequest) HasAdditionalInformationLabel() bool`

HasAdditionalInformationLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


