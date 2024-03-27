# VirtualNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualNumber** | Pointer to **string** | The Virtual Number assigned to your account.  | [optional] 
**ReplyCallbackUrl** | Pointer to **string** | The URL that replies to the Virtual Number will be posted to. | [optional] 
**Tags** | Pointer to **[]string** | Any customisable tags assigned to the Virtual Number.  | [optional] 
**LastUse** | Pointer to **time.Time** | The last time the Virtual Number was used to send a message. | [optional] 

## Methods

### NewVirtualNumber

`func NewVirtualNumber() *VirtualNumber`

NewVirtualNumber instantiates a new VirtualNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNumberWithDefaults

`func NewVirtualNumberWithDefaults() *VirtualNumber`

NewVirtualNumberWithDefaults instantiates a new VirtualNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualNumber

`func (o *VirtualNumber) GetVirtualNumber() string`

GetVirtualNumber returns the VirtualNumber field if non-nil, zero value otherwise.

### GetVirtualNumberOk

`func (o *VirtualNumber) GetVirtualNumberOk() (*string, bool)`

GetVirtualNumberOk returns a tuple with the VirtualNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNumber

`func (o *VirtualNumber) SetVirtualNumber(v string)`

SetVirtualNumber sets VirtualNumber field to given value.

### HasVirtualNumber

`func (o *VirtualNumber) HasVirtualNumber() bool`

HasVirtualNumber returns a boolean if a field has been set.

### GetReplyCallbackUrl

`func (o *VirtualNumber) GetReplyCallbackUrl() string`

GetReplyCallbackUrl returns the ReplyCallbackUrl field if non-nil, zero value otherwise.

### GetReplyCallbackUrlOk

`func (o *VirtualNumber) GetReplyCallbackUrlOk() (*string, bool)`

GetReplyCallbackUrlOk returns a tuple with the ReplyCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCallbackUrl

`func (o *VirtualNumber) SetReplyCallbackUrl(v string)`

SetReplyCallbackUrl sets ReplyCallbackUrl field to given value.

### HasReplyCallbackUrl

`func (o *VirtualNumber) HasReplyCallbackUrl() bool`

HasReplyCallbackUrl returns a boolean if a field has been set.

### GetTags

`func (o *VirtualNumber) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualNumber) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualNumber) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualNumber) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastUse

`func (o *VirtualNumber) GetLastUse() time.Time`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *VirtualNumber) GetLastUseOk() (*time.Time, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *VirtualNumber) SetLastUse(v time.Time)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *VirtualNumber) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


