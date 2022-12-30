# AuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**Response**](Response.md) |  | 
**Authorization** | [**Authorization**](Authorization.md) |  | 

## Methods

### NewAuthorizationResponse

`func NewAuthorizationResponse(response Response, authorization Authorization, ) *AuthorizationResponse`

NewAuthorizationResponse instantiates a new AuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResponseWithDefaults

`func NewAuthorizationResponseWithDefaults() *AuthorizationResponse`

NewAuthorizationResponseWithDefaults instantiates a new AuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *AuthorizationResponse) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AuthorizationResponse) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AuthorizationResponse) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetAuthorization

`func (o *AuthorizationResponse) GetAuthorization() Authorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *AuthorizationResponse) GetAuthorizationOk() (*Authorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *AuthorizationResponse) SetAuthorization(v Authorization)`

SetAuthorization sets Authorization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


