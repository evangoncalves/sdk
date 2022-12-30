# AuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**System** | **string** | System which this authorization belongs to | 
**Environment** | **string** | Environment which this authorization belongs to | 
**Expiration** | Pointer to **time.Time** | Expiration date of an authorization (must be in YYYY-MM-DD HH:MM format) | [optional] 
**Encryption** | **bool** | Whether secret sensitive data will be encrypted on responses | 
**Permissions** | **string** | Permission level for reading and writing application data | 
**AllowedIps** | Pointer to **[]string** | List of IPs allowed to make requests using this authorization | [optional] 
**AllowedReferers** | Pointer to **[]string** | List of Referes allowed to make requests using this authorization | [optional] 
**Arns** | Pointer to **[]string** | List of AWS ARNS allowed to request data through this authorization | [optional] 
**Ceritificate** | Pointer to **string** | Certificate fingerprint used to validate access on secrets | [optional] 

## Methods

### NewAuthorizationRequest

`func NewAuthorizationRequest(system string, environment string, encryption bool, permissions string, ) *AuthorizationRequest`

NewAuthorizationRequest instantiates a new AuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationRequestWithDefaults

`func NewAuthorizationRequestWithDefaults() *AuthorizationRequest`

NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystem

`func (o *AuthorizationRequest) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AuthorizationRequest) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AuthorizationRequest) SetSystem(v string)`

SetSystem sets System field to given value.


### GetEnvironment

`func (o *AuthorizationRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizationRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizationRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetExpiration

`func (o *AuthorizationRequest) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AuthorizationRequest) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AuthorizationRequest) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AuthorizationRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetEncryption

`func (o *AuthorizationRequest) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *AuthorizationRequest) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *AuthorizationRequest) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.


### GetPermissions

`func (o *AuthorizationRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AuthorizationRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AuthorizationRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetAllowedIps

`func (o *AuthorizationRequest) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *AuthorizationRequest) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *AuthorizationRequest) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *AuthorizationRequest) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### GetAllowedReferers

`func (o *AuthorizationRequest) GetAllowedReferers() []string`

GetAllowedReferers returns the AllowedReferers field if non-nil, zero value otherwise.

### GetAllowedReferersOk

`func (o *AuthorizationRequest) GetAllowedReferersOk() (*[]string, bool)`

GetAllowedReferersOk returns a tuple with the AllowedReferers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedReferers

`func (o *AuthorizationRequest) SetAllowedReferers(v []string)`

SetAllowedReferers sets AllowedReferers field to given value.

### HasAllowedReferers

`func (o *AuthorizationRequest) HasAllowedReferers() bool`

HasAllowedReferers returns a boolean if a field has been set.

### GetArns

`func (o *AuthorizationRequest) GetArns() []string`

GetArns returns the Arns field if non-nil, zero value otherwise.

### GetArnsOk

`func (o *AuthorizationRequest) GetArnsOk() (*[]string, bool)`

GetArnsOk returns a tuple with the Arns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArns

`func (o *AuthorizationRequest) SetArns(v []string)`

SetArns sets Arns field to given value.

### HasArns

`func (o *AuthorizationRequest) HasArns() bool`

HasArns returns a boolean if a field has been set.

### GetCeritificate

`func (o *AuthorizationRequest) GetCeritificate() string`

GetCeritificate returns the Ceritificate field if non-nil, zero value otherwise.

### GetCeritificateOk

`func (o *AuthorizationRequest) GetCeritificateOk() (*string, bool)`

GetCeritificateOk returns a tuple with the Ceritificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeritificate

`func (o *AuthorizationRequest) SetCeritificate(v string)`

SetCeritificate sets Ceritificate field to given value.

### HasCeritificate

`func (o *AuthorizationRequest) HasCeritificate() bool`

HasCeritificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


