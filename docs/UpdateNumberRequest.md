# UpdateNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyCallbackUrl** | Pointer to **string** | Tell us the URL that replies to the Virtual Number should be sent to.  Note that if you don&#39;t include this field, any pre-existing replyCallbackUrl will be wiped.  Sample callback response:  &lt;pre&gt;&lt;code class&#x3D;\&quot;language-sh\&quot;&gt;{   \&quot;to\&quot;:\&quot;0476543210\&quot;,    \&quot;from\&quot;:\&quot;0401234567\&quot;,    \&quot;timestamp\&quot;:\&quot;2022-11-10T05:06:42.823Z\&quot;,   \&quot;messageId\&quot;:\&quot;75f263c0-60b5-11ed-8456-71ae4c63550d\&quot;,    \&quot;messageContent\&quot;:\&quot;Hi, example message\&quot;,    \&quot;multimedia\&quot;: {      \&quot;fileName:\&quot;image.jpeg\&quot;      \&quot;type:\&quot;image/jpeg\&quot;      \&quot;payload\&quot;:\&quot;base64 payload\&quot;    } }&lt;/code&gt;&lt;/pre&gt;  | [optional] 
**Tags** | Pointer to **[]string** | Create your own tags and use them to fetch, sort and report on your Virtual Numbers through our other endpoints. You can assign up to 10 tags per number.   Note that if you don&#39;t include this field, any pre-existing tags will be wiped.  | [optional] 

## Methods

### NewUpdateNumberRequest

`func NewUpdateNumberRequest() *UpdateNumberRequest`

NewUpdateNumberRequest instantiates a new UpdateNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNumberRequestWithDefaults

`func NewUpdateNumberRequestWithDefaults() *UpdateNumberRequest`

NewUpdateNumberRequestWithDefaults instantiates a new UpdateNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyCallbackUrl

`func (o *UpdateNumberRequest) GetReplyCallbackUrl() string`

GetReplyCallbackUrl returns the ReplyCallbackUrl field if non-nil, zero value otherwise.

### GetReplyCallbackUrlOk

`func (o *UpdateNumberRequest) GetReplyCallbackUrlOk() (*string, bool)`

GetReplyCallbackUrlOk returns a tuple with the ReplyCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCallbackUrl

`func (o *UpdateNumberRequest) SetReplyCallbackUrl(v string)`

SetReplyCallbackUrl sets ReplyCallbackUrl field to given value.

### HasReplyCallbackUrl

`func (o *UpdateNumberRequest) HasReplyCallbackUrl() bool`

HasReplyCallbackUrl returns a boolean if a field has been set.

### GetTags

`func (o *UpdateNumberRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateNumberRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateNumberRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateNumberRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


