# SecretPAMDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Hostname** | **string** |  | 
**Ip** | **string** |  | 
**Type** | **string** |  | 
**Vendor** | **string** |  | 
**Product** | **string** |  | 
**Tenant** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Connectivities** | Pointer to [**[]SecretPAMDeviceConnectivitiesInner**](SecretPAMDeviceConnectivitiesInner.md) |  | [optional] 

## Methods

### NewSecretPAMDevice

`func NewSecretPAMDevice(id int32, hostname string, ip string, type_ string, vendor string, product string, ) *SecretPAMDevice`

NewSecretPAMDevice instantiates a new SecretPAMDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretPAMDeviceWithDefaults

`func NewSecretPAMDeviceWithDefaults() *SecretPAMDevice`

NewSecretPAMDeviceWithDefaults instantiates a new SecretPAMDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretPAMDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretPAMDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretPAMDevice) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *SecretPAMDevice) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SecretPAMDevice) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SecretPAMDevice) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIp

`func (o *SecretPAMDevice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SecretPAMDevice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SecretPAMDevice) SetIp(v string)`

SetIp sets Ip field to given value.


### GetType

`func (o *SecretPAMDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretPAMDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretPAMDevice) SetType(v string)`

SetType sets Type field to given value.


### GetVendor

`func (o *SecretPAMDevice) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SecretPAMDevice) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SecretPAMDevice) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetProduct

`func (o *SecretPAMDevice) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SecretPAMDevice) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SecretPAMDevice) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetTenant

`func (o *SecretPAMDevice) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SecretPAMDevice) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SecretPAMDevice) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *SecretPAMDevice) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetSite

`func (o *SecretPAMDevice) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SecretPAMDevice) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SecretPAMDevice) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *SecretPAMDevice) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetTags

`func (o *SecretPAMDevice) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecretPAMDevice) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecretPAMDevice) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecretPAMDevice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDomains

`func (o *SecretPAMDevice) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SecretPAMDevice) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SecretPAMDevice) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SecretPAMDevice) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetConnectivities

`func (o *SecretPAMDevice) GetConnectivities() []SecretPAMDeviceConnectivitiesInner`

GetConnectivities returns the Connectivities field if non-nil, zero value otherwise.

### GetConnectivitiesOk

`func (o *SecretPAMDevice) GetConnectivitiesOk() (*[]SecretPAMDeviceConnectivitiesInner, bool)`

GetConnectivitiesOk returns a tuple with the Connectivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivities

`func (o *SecretPAMDevice) SetConnectivities(v []SecretPAMDeviceConnectivitiesInner)`

SetConnectivities sets Connectivities field to given value.

### HasConnectivities

`func (o *SecretPAMDevice) HasConnectivities() bool`

HasConnectivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


