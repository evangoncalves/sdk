# SecretCloudCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AccessKeyIdLabel** | Pointer to **string** |  | [optional] 
**SecretAccessKeyLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretCloudCredentialRequest

`func NewSecretCloudCredentialRequest(id int32, ) *SecretCloudCredentialRequest`

NewSecretCloudCredentialRequest instantiates a new SecretCloudCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretCloudCredentialRequestWithDefaults

`func NewSecretCloudCredentialRequestWithDefaults() *SecretCloudCredentialRequest`

NewSecretCloudCredentialRequestWithDefaults instantiates a new SecretCloudCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretCloudCredentialRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretCloudCredentialRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretCloudCredentialRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetAccessKeyIdLabel

`func (o *SecretCloudCredentialRequest) GetAccessKeyIdLabel() string`

GetAccessKeyIdLabel returns the AccessKeyIdLabel field if non-nil, zero value otherwise.

### GetAccessKeyIdLabelOk

`func (o *SecretCloudCredentialRequest) GetAccessKeyIdLabelOk() (*string, bool)`

GetAccessKeyIdLabelOk returns a tuple with the AccessKeyIdLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyIdLabel

`func (o *SecretCloudCredentialRequest) SetAccessKeyIdLabel(v string)`

SetAccessKeyIdLabel sets AccessKeyIdLabel field to given value.

### HasAccessKeyIdLabel

`func (o *SecretCloudCredentialRequest) HasAccessKeyIdLabel() bool`

HasAccessKeyIdLabel returns a boolean if a field has been set.

### GetSecretAccessKeyLabel

`func (o *SecretCloudCredentialRequest) GetSecretAccessKeyLabel() string`

GetSecretAccessKeyLabel returns the SecretAccessKeyLabel field if non-nil, zero value otherwise.

### GetSecretAccessKeyLabelOk

`func (o *SecretCloudCredentialRequest) GetSecretAccessKeyLabelOk() (*string, bool)`

GetSecretAccessKeyLabelOk returns a tuple with the SecretAccessKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKeyLabel

`func (o *SecretCloudCredentialRequest) SetSecretAccessKeyLabel(v string)`

SetSecretAccessKeyLabel sets SecretAccessKeyLabel field to given value.

### HasSecretAccessKeyLabel

`func (o *SecretCloudCredentialRequest) HasSecretAccessKeyLabel() bool`

HasSecretAccessKeyLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


