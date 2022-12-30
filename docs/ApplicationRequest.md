# ApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the application | 
**Description** | Pointer to **string** | A descriptive text for the application | [optional] 
**AuthMethod** | **string** | Authentication method used by this application&#39;s authorizations | 
**LineOfBusiness** | Pointer to **string** | Line of business that best describe this application usage | [optional] 
**Type** | Pointer to **string** | Type of application | [optional] 
**Tags** | Pointer to **[]string** | List of tags used for access segregation | [optional] 

## Methods

### NewApplicationRequest

`func NewApplicationRequest(name string, authMethod string, ) *ApplicationRequest`

NewApplicationRequest instantiates a new ApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestWithDefaults

`func NewApplicationRequestWithDefaults() *ApplicationRequest`

NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApplicationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthMethod

`func (o *ApplicationRequest) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *ApplicationRequest) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *ApplicationRequest) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetLineOfBusiness

`func (o *ApplicationRequest) GetLineOfBusiness() string`

GetLineOfBusiness returns the LineOfBusiness field if non-nil, zero value otherwise.

### GetLineOfBusinessOk

`func (o *ApplicationRequest) GetLineOfBusinessOk() (*string, bool)`

GetLineOfBusinessOk returns a tuple with the LineOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOfBusiness

`func (o *ApplicationRequest) SetLineOfBusiness(v string)`

SetLineOfBusiness sets LineOfBusiness field to given value.

### HasLineOfBusiness

`func (o *ApplicationRequest) HasLineOfBusiness() bool`

HasLineOfBusiness returns a boolean if a field has been set.

### GetType

`func (o *ApplicationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *ApplicationRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


