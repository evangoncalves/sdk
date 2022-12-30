# Secrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Secret ID generated, by senhasegura | [readonly] 
**Name** | **string** | Secret name | 
**Identifier** | **string** | Secret identifier (must be unique on the platform) | 
**Engine** | **string** | Engine where this secret will be used | 
**Environment** | **string** |  | 
**Expiration** | Pointer to **time.Time** | Secret expiration date in UTC pattern | [optional] 
**Description** | Pointer to **string** | A description about the usage of this Secret | [optional] 
**Tags** | Pointer to **[]string** | Secret tags used for access segregation | [optional] 

## Methods

### NewSecrets

`func NewSecrets(id int32, name string, identifier string, engine string, environment string, ) *Secrets`

NewSecrets instantiates a new Secrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsWithDefaults

`func NewSecretsWithDefaults() *Secrets`

NewSecretsWithDefaults instantiates a new Secrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Secrets) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Secrets) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Secrets) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Secrets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secrets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secrets) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *Secrets) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Secrets) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Secrets) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetEngine

`func (o *Secrets) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *Secrets) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *Secrets) SetEngine(v string)`

SetEngine sets Engine field to given value.


### GetEnvironment

`func (o *Secrets) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Secrets) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Secrets) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetExpiration

`func (o *Secrets) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Secrets) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Secrets) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Secrets) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetDescription

`func (o *Secrets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Secrets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Secrets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Secrets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *Secrets) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Secrets) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Secrets) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Secrets) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


