# UpdateMessageTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** | Write the updated list of tag(s) here. You can assign up to 10 tags per message.  Note that if you provide an empty array, any pre-existing tags will be wiped.  | 

## Methods

### NewUpdateMessageTagsRequest

`func NewUpdateMessageTagsRequest(tags []string, ) *UpdateMessageTagsRequest`

NewUpdateMessageTagsRequest instantiates a new UpdateMessageTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageTagsRequestWithDefaults

`func NewUpdateMessageTagsRequestWithDefaults() *UpdateMessageTagsRequest`

NewUpdateMessageTagsRequestWithDefaults instantiates a new UpdateMessageTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpdateMessageTagsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateMessageTagsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateMessageTagsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


