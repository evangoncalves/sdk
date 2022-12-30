# AuthorizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**Response**](Response.md) |  | 
**Authorizations** | [**[]Authorizations**](Authorizations.md) |  | 

## Methods

### NewAuthorizationsResponse

`func NewAuthorizationsResponse(response Response, authorizations []Authorizations, ) *AuthorizationsResponse`

NewAuthorizationsResponse instantiates a new AuthorizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationsResponseWithDefaults

`func NewAuthorizationsResponseWithDefaults() *AuthorizationsResponse`

NewAuthorizationsResponseWithDefaults instantiates a new AuthorizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *AuthorizationsResponse) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AuthorizationsResponse) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AuthorizationsResponse) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetAuthorizations

`func (o *AuthorizationsResponse) GetAuthorizations() []Authorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *AuthorizationsResponse) GetAuthorizationsOk() (*[]Authorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *AuthorizationsResponse) SetAuthorizations(v []Authorizations)`

SetAuthorizations sets Authorizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


