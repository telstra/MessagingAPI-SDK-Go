# MessageUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Use this UUID with our other endpoints to fetch, update or delete the message. | [optional] 
**Status** | Pointer to **string** | The status will be either queued, sent, delivered, expired or undeliverable. | [optional] 
**To** | Pointer to **string** | The recipient&#39;s mobile number. | [optional] 
**From** | Pointer to **string** | This will be either one of your Virtual Numbers or your senderName. | [optional] 
**MessageContent** | Pointer to **string** | The content of the message. | [optional] 
**Multimedia** | Pointer to [**[]MultimediaGet**](MultimediaGet.md) | The multimedia content of the message (MMS only). It will include:  | [optional] 
**RetryTimeout** | Pointer to **int32** | How many minutes you asked the server to keep trying to send the message. | [optional] 
**ScheduleSend** | Pointer to **time.Time** | The time the message is scheduled to send. | [optional] 
**DeliveryNotification** | Pointer to **bool** | If set to **true**, you will receive a notification to the statusCallbackUrl when your message is delivered (paid feature). | [optional] [default to false]
**StatusCallbackUrl** | Pointer to **string** | The URL the API will call when the status of the message changes. | [optional] 
**QueuePriority** | Pointer to **int32** | The priority assigned to the message. | [optional] [default to 2]
**Tags** | Pointer to **[]string** | Any customisable tags assigned to the message. | [optional] 

## Methods

### NewMessageUpdate

`func NewMessageUpdate() *MessageUpdate`

NewMessageUpdate instantiates a new MessageUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUpdateWithDefaults

`func NewMessageUpdateWithDefaults() *MessageUpdate`

NewMessageUpdateWithDefaults instantiates a new MessageUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MessageUpdate) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageUpdate) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageUpdate) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageUpdate) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *MessageUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *MessageUpdate) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageUpdate) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageUpdate) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageUpdate) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *MessageUpdate) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageUpdate) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageUpdate) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageUpdate) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMessageContent

`func (o *MessageUpdate) GetMessageContent() string`

GetMessageContent returns the MessageContent field if non-nil, zero value otherwise.

### GetMessageContentOk

`func (o *MessageUpdate) GetMessageContentOk() (*string, bool)`

GetMessageContentOk returns a tuple with the MessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContent

`func (o *MessageUpdate) SetMessageContent(v string)`

SetMessageContent sets MessageContent field to given value.

### HasMessageContent

`func (o *MessageUpdate) HasMessageContent() bool`

HasMessageContent returns a boolean if a field has been set.

### GetMultimedia

`func (o *MessageUpdate) GetMultimedia() []MultimediaGet`

GetMultimedia returns the Multimedia field if non-nil, zero value otherwise.

### GetMultimediaOk

`func (o *MessageUpdate) GetMultimediaOk() (*[]MultimediaGet, bool)`

GetMultimediaOk returns a tuple with the Multimedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultimedia

`func (o *MessageUpdate) SetMultimedia(v []MultimediaGet)`

SetMultimedia sets Multimedia field to given value.

### HasMultimedia

`func (o *MessageUpdate) HasMultimedia() bool`

HasMultimedia returns a boolean if a field has been set.

### GetRetryTimeout

`func (o *MessageUpdate) GetRetryTimeout() int32`

GetRetryTimeout returns the RetryTimeout field if non-nil, zero value otherwise.

### GetRetryTimeoutOk

`func (o *MessageUpdate) GetRetryTimeoutOk() (*int32, bool)`

GetRetryTimeoutOk returns a tuple with the RetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimeout

`func (o *MessageUpdate) SetRetryTimeout(v int32)`

SetRetryTimeout sets RetryTimeout field to given value.

### HasRetryTimeout

`func (o *MessageUpdate) HasRetryTimeout() bool`

HasRetryTimeout returns a boolean if a field has been set.

### GetScheduleSend

`func (o *MessageUpdate) GetScheduleSend() time.Time`

GetScheduleSend returns the ScheduleSend field if non-nil, zero value otherwise.

### GetScheduleSendOk

`func (o *MessageUpdate) GetScheduleSendOk() (*time.Time, bool)`

GetScheduleSendOk returns a tuple with the ScheduleSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSend

`func (o *MessageUpdate) SetScheduleSend(v time.Time)`

SetScheduleSend sets ScheduleSend field to given value.

### HasScheduleSend

`func (o *MessageUpdate) HasScheduleSend() bool`

HasScheduleSend returns a boolean if a field has been set.

### GetDeliveryNotification

`func (o *MessageUpdate) GetDeliveryNotification() bool`

GetDeliveryNotification returns the DeliveryNotification field if non-nil, zero value otherwise.

### GetDeliveryNotificationOk

`func (o *MessageUpdate) GetDeliveryNotificationOk() (*bool, bool)`

GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNotification

`func (o *MessageUpdate) SetDeliveryNotification(v bool)`

SetDeliveryNotification sets DeliveryNotification field to given value.

### HasDeliveryNotification

`func (o *MessageUpdate) HasDeliveryNotification() bool`

HasDeliveryNotification returns a boolean if a field has been set.

### GetStatusCallbackUrl

`func (o *MessageUpdate) GetStatusCallbackUrl() string`

GetStatusCallbackUrl returns the StatusCallbackUrl field if non-nil, zero value otherwise.

### GetStatusCallbackUrlOk

`func (o *MessageUpdate) GetStatusCallbackUrlOk() (*string, bool)`

GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackUrl

`func (o *MessageUpdate) SetStatusCallbackUrl(v string)`

SetStatusCallbackUrl sets StatusCallbackUrl field to given value.

### HasStatusCallbackUrl

`func (o *MessageUpdate) HasStatusCallbackUrl() bool`

HasStatusCallbackUrl returns a boolean if a field has been set.

### GetQueuePriority

`func (o *MessageUpdate) GetQueuePriority() int32`

GetQueuePriority returns the QueuePriority field if non-nil, zero value otherwise.

### GetQueuePriorityOk

`func (o *MessageUpdate) GetQueuePriorityOk() (*int32, bool)`

GetQueuePriorityOk returns a tuple with the QueuePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePriority

`func (o *MessageUpdate) SetQueuePriority(v int32)`

SetQueuePriority sets QueuePriority field to given value.

### HasQueuePriority

`func (o *MessageUpdate) HasQueuePriority() bool`

HasQueuePriority returns a boolean if a field has been set.

### GetTags

`func (o *MessageUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MessageUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MessageUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MessageUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


