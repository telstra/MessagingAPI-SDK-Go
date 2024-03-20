# MessageSent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to [**MessageSentMessageId**](MessageSentMessageId.md) |  | [optional] 
**Status** | Pointer to **string** | The status will be either queued, sent, delivered, expired or undeliverable. | [optional] 
**To** | Pointer to [**MessageSentTo**](MessageSentTo.md) |  | [optional] 
**From** | Pointer to **string** | This will be either  one of your Virtual Numbers or your senderName. | [optional] 
**MessageContent** | Pointer to **string** | The content of the message. | [optional] 
**Multimedia** | Pointer to [**[]MultimediaGet**](MultimediaGet.md) | The multimedia content of the message (MMS only). It will include:  | [optional] 
**RetryTimeout** | Pointer to **int32** | How many minutes you asked the server to keep trying to send the message. | [optional] 
**ScheduleSend** | Pointer to **time.Time** | The time (in Central Standard Time) the message is scheduled to send. | [optional] 
**DeliveryNotification** | Pointer to **bool** | If set to **true**, you will receive a notification to the statusCallbackUrl when your SMS is delivered (paid feature). | [optional] [default to false]
**StatusCallbackUrl** | Pointer to **string** | The URL the API will call when the status of the message changes. | [optional] 
**Tags** | Pointer to **[]string** | Any customisable tags assigned to the message. | [optional] 

## Methods

### NewMessageSent

`func NewMessageSent() *MessageSent`

NewMessageSent instantiates a new MessageSent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSentWithDefaults

`func NewMessageSentWithDefaults() *MessageSent`

NewMessageSentWithDefaults instantiates a new MessageSent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MessageSent) GetMessageId() MessageSentMessageId`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageSent) GetMessageIdOk() (*MessageSentMessageId, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageSent) SetMessageId(v MessageSentMessageId)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageSent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *MessageSent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageSent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageSent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageSent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *MessageSent) GetTo() MessageSentTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageSent) GetToOk() (*MessageSentTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageSent) SetTo(v MessageSentTo)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageSent) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *MessageSent) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageSent) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageSent) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageSent) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMessageContent

`func (o *MessageSent) GetMessageContent() string`

GetMessageContent returns the MessageContent field if non-nil, zero value otherwise.

### GetMessageContentOk

`func (o *MessageSent) GetMessageContentOk() (*string, bool)`

GetMessageContentOk returns a tuple with the MessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContent

`func (o *MessageSent) SetMessageContent(v string)`

SetMessageContent sets MessageContent field to given value.

### HasMessageContent

`func (o *MessageSent) HasMessageContent() bool`

HasMessageContent returns a boolean if a field has been set.

### GetMultimedia

`func (o *MessageSent) GetMultimedia() []MultimediaGet`

GetMultimedia returns the Multimedia field if non-nil, zero value otherwise.

### GetMultimediaOk

`func (o *MessageSent) GetMultimediaOk() (*[]MultimediaGet, bool)`

GetMultimediaOk returns a tuple with the Multimedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultimedia

`func (o *MessageSent) SetMultimedia(v []MultimediaGet)`

SetMultimedia sets Multimedia field to given value.

### HasMultimedia

`func (o *MessageSent) HasMultimedia() bool`

HasMultimedia returns a boolean if a field has been set.

### GetRetryTimeout

`func (o *MessageSent) GetRetryTimeout() int32`

GetRetryTimeout returns the RetryTimeout field if non-nil, zero value otherwise.

### GetRetryTimeoutOk

`func (o *MessageSent) GetRetryTimeoutOk() (*int32, bool)`

GetRetryTimeoutOk returns a tuple with the RetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimeout

`func (o *MessageSent) SetRetryTimeout(v int32)`

SetRetryTimeout sets RetryTimeout field to given value.

### HasRetryTimeout

`func (o *MessageSent) HasRetryTimeout() bool`

HasRetryTimeout returns a boolean if a field has been set.

### GetScheduleSend

`func (o *MessageSent) GetScheduleSend() time.Time`

GetScheduleSend returns the ScheduleSend field if non-nil, zero value otherwise.

### GetScheduleSendOk

`func (o *MessageSent) GetScheduleSendOk() (*time.Time, bool)`

GetScheduleSendOk returns a tuple with the ScheduleSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSend

`func (o *MessageSent) SetScheduleSend(v time.Time)`

SetScheduleSend sets ScheduleSend field to given value.

### HasScheduleSend

`func (o *MessageSent) HasScheduleSend() bool`

HasScheduleSend returns a boolean if a field has been set.

### GetDeliveryNotification

`func (o *MessageSent) GetDeliveryNotification() bool`

GetDeliveryNotification returns the DeliveryNotification field if non-nil, zero value otherwise.

### GetDeliveryNotificationOk

`func (o *MessageSent) GetDeliveryNotificationOk() (*bool, bool)`

GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNotification

`func (o *MessageSent) SetDeliveryNotification(v bool)`

SetDeliveryNotification sets DeliveryNotification field to given value.

### HasDeliveryNotification

`func (o *MessageSent) HasDeliveryNotification() bool`

HasDeliveryNotification returns a boolean if a field has been set.

### GetStatusCallbackUrl

`func (o *MessageSent) GetStatusCallbackUrl() string`

GetStatusCallbackUrl returns the StatusCallbackUrl field if non-nil, zero value otherwise.

### GetStatusCallbackUrlOk

`func (o *MessageSent) GetStatusCallbackUrlOk() (*string, bool)`

GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackUrl

`func (o *MessageSent) SetStatusCallbackUrl(v string)`

SetStatusCallbackUrl sets StatusCallbackUrl field to given value.

### HasStatusCallbackUrl

`func (o *MessageSent) HasStatusCallbackUrl() bool`

HasStatusCallbackUrl returns a boolean if a field has been set.

### GetTags

`func (o *MessageSent) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MessageSent) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MessageSent) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MessageSent) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


