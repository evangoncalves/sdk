# ApplicationPAMCredentialProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Profile name registered on senhasegura PAM module | 
**Target** | **string** | Target asset to execute this profile | 

## Methods

### NewApplicationPAMCredentialProfileRequest

`func NewApplicationPAMCredentialProfileRequest(identifier string, target string, ) *ApplicationPAMCredentialProfileRequest`

NewApplicationPAMCredentialProfileRequest instantiates a new ApplicationPAMCredentialProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPAMCredentialProfileRequestWithDefaults

`func NewApplicationPAMCredentialProfileRequestWithDefaults() *ApplicationPAMCredentialProfileRequest`

NewApplicationPAMCredentialProfileRequestWithDefaults instantiates a new ApplicationPAMCredentialProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ApplicationPAMCredentialProfileRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ApplicationPAMCredentialProfileRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ApplicationPAMCredentialProfileRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetTarget

`func (o *ApplicationPAMCredentialProfileRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ApplicationPAMCredentialProfileRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ApplicationPAMCredentialProfileRequest) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


