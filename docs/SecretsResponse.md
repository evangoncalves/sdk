# SecretsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**Response**](Response.md) |  | 
**Secrets** | [**[]Secrets**](Secrets.md) |  | 

## Methods

### NewSecretsResponse

`func NewSecretsResponse(response Response, secrets []Secrets, ) *SecretsResponse`

NewSecretsResponse instantiates a new SecretsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsResponseWithDefaults

`func NewSecretsResponseWithDefaults() *SecretsResponse`

NewSecretsResponseWithDefaults instantiates a new SecretsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *SecretsResponse) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SecretsResponse) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SecretsResponse) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetSecrets

`func (o *SecretsResponse) GetSecrets() []Secrets`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SecretsResponse) GetSecretsOk() (*[]Secrets, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SecretsResponse) SetSecrets(v []Secrets)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


