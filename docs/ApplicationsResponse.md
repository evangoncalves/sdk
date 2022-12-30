# ApplicationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**Response**](Response.md) |  | 
**Applications** | [**[]Applications**](Applications.md) |  | 

## Methods

### NewApplicationsResponse

`func NewApplicationsResponse(response Response, applications []Applications, ) *ApplicationsResponse`

NewApplicationsResponse instantiates a new ApplicationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsResponseWithDefaults

`func NewApplicationsResponseWithDefaults() *ApplicationsResponse`

NewApplicationsResponseWithDefaults instantiates a new ApplicationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ApplicationsResponse) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApplicationsResponse) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApplicationsResponse) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetApplications

`func (o *ApplicationsResponse) GetApplications() []Applications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationsResponse) GetApplicationsOk() (*[]Applications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationsResponse) SetApplications(v []Applications)`

SetApplications sets Applications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


