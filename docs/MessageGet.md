# MessageGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Use this UUID with our other endpoints to fetch, update or delete the message. | [optional] 
**Status** | Pointer to **string** | The status will be either queued, sent, delivered, expired or undeliverable. | [optional] 
**CreateTimestamp** | Pointer to **time.Time** | The time you submitted the message to the queue for sending. | [optional] 
**SentTimestamp** | Pointer to **time.Time** | The time the message was sent from the server. | [optional] 
**ReceivedTimestamp** | Pointer to **time.Time** | The time the message was received by the recipient&#39;s device. | [optional] 
**To** | Pointer to **string** | The recipient&#39;s mobile number. | [optional] 
**From** | Pointer to **string** | This will be either  one of your Virtual Numbers or your senderName. | [optional] 
**MessageContent** | Pointer to **string** | The content of the message. | [optional] 
**Multimedia** | Pointer to [**[]MultimediaGet**](MultimediaGet.md) | The multimedia content of the message (MMS only). It will include:  | [optional] 
**Direction** | Pointer to **string** | * **outgoing** – messages sent from your account. * **incoming** – messages sent to your account.  | [optional] 
**RetryTimeout** | Pointer to **int32** | How many minutes you asked the server to keep trying to send the message. | [optional] [default to 10]
**ScheduleSend** | Pointer to **time.Time** | The time the message is scheduled to send. | [optional] 
**DeliveryNotification** | Pointer to **bool** | If set to **true**, you will receive a notification to the statusCallbackUrl when your message is delivered (paid feature). | [optional] [default to false]
**StatusCallbackUrl** | Pointer to **string** | The URL the API will call when the status of the message changes. | [optional] 
**QueuePriority** | Pointer to **int32** | The priority assigned to the message. | [optional] [default to 2]
**Tags** | Pointer to **[]string** | Any customisable tags assigned to the message. | [optional] 

## Methods

### NewMessageGet

`func NewMessageGet() *MessageGet`

NewMessageGet instantiates a new MessageGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageGetWithDefaults

`func NewMessageGetWithDefaults() *MessageGet`

NewMessageGetWithDefaults instantiates a new MessageGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MessageGet) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageGet) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageGet) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageGet) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *MessageGet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageGet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageGet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageGet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *MessageGet) GetCreateTimestamp() time.Time`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *MessageGet) GetCreateTimestampOk() (*time.Time, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *MessageGet) SetCreateTimestamp(v time.Time)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *MessageGet) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetSentTimestamp

`func (o *MessageGet) GetSentTimestamp() time.Time`

GetSentTimestamp returns the SentTimestamp field if non-nil, zero value otherwise.

### GetSentTimestampOk

`func (o *MessageGet) GetSentTimestampOk() (*time.Time, bool)`

GetSentTimestampOk returns a tuple with the SentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTimestamp

`func (o *MessageGet) SetSentTimestamp(v time.Time)`

SetSentTimestamp sets SentTimestamp field to given value.

### HasSentTimestamp

`func (o *MessageGet) HasSentTimestamp() bool`

HasSentTimestamp returns a boolean if a field has been set.

### GetReceivedTimestamp

`func (o *MessageGet) GetReceivedTimestamp() time.Time`

GetReceivedTimestamp returns the ReceivedTimestamp field if non-nil, zero value otherwise.

### GetReceivedTimestampOk

`func (o *MessageGet) GetReceivedTimestampOk() (*time.Time, bool)`

GetReceivedTimestampOk returns a tuple with the ReceivedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimestamp

`func (o *MessageGet) SetReceivedTimestamp(v time.Time)`

SetReceivedTimestamp sets ReceivedTimestamp field to given value.

### HasReceivedTimestamp

`func (o *MessageGet) HasReceivedTimestamp() bool`

HasReceivedTimestamp returns a boolean if a field has been set.

### GetTo

`func (o *MessageGet) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageGet) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageGet) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageGet) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *MessageGet) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageGet) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageGet) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageGet) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMessageContent

`func (o *MessageGet) GetMessageContent() string`

GetMessageContent returns the MessageContent field if non-nil, zero value otherwise.

### GetMessageContentOk

`func (o *MessageGet) GetMessageContentOk() (*string, bool)`

GetMessageContentOk returns a tuple with the MessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContent

`func (o *MessageGet) SetMessageContent(v string)`

SetMessageContent sets MessageContent field to given value.

### HasMessageContent

`func (o *MessageGet) HasMessageContent() bool`

HasMessageContent returns a boolean if a field has been set.

### GetMultimedia

`func (o *MessageGet) GetMultimedia() []MultimediaGet`

GetMultimedia returns the Multimedia field if non-nil, zero value otherwise.

### GetMultimediaOk

`func (o *MessageGet) GetMultimediaOk() (*[]MultimediaGet, bool)`

GetMultimediaOk returns a tuple with the Multimedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultimedia

`func (o *MessageGet) SetMultimedia(v []MultimediaGet)`

SetMultimedia sets Multimedia field to given value.

### HasMultimedia

`func (o *MessageGet) HasMultimedia() bool`

HasMultimedia returns a boolean if a field has been set.

### GetDirection

`func (o *MessageGet) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MessageGet) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MessageGet) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MessageGet) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRetryTimeout

`func (o *MessageGet) GetRetryTimeout() int32`

GetRetryTimeout returns the RetryTimeout field if non-nil, zero value otherwise.

### GetRetryTimeoutOk

`func (o *MessageGet) GetRetryTimeoutOk() (*int32, bool)`

GetRetryTimeoutOk returns a tuple with the RetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimeout

`func (o *MessageGet) SetRetryTimeout(v int32)`

SetRetryTimeout sets RetryTimeout field to given value.

### HasRetryTimeout

`func (o *MessageGet) HasRetryTimeout() bool`

HasRetryTimeout returns a boolean if a field has been set.

### GetScheduleSend

`func (o *MessageGet) GetScheduleSend() time.Time`

GetScheduleSend returns the ScheduleSend field if non-nil, zero value otherwise.

### GetScheduleSendOk

`func (o *MessageGet) GetScheduleSendOk() (*time.Time, bool)`

GetScheduleSendOk returns a tuple with the ScheduleSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSend

`func (o *MessageGet) SetScheduleSend(v time.Time)`

SetScheduleSend sets ScheduleSend field to given value.

### HasScheduleSend

`func (o *MessageGet) HasScheduleSend() bool`

HasScheduleSend returns a boolean if a field has been set.

### GetDeliveryNotification

`func (o *MessageGet) GetDeliveryNotification() bool`

GetDeliveryNotification returns the DeliveryNotification field if non-nil, zero value otherwise.

### GetDeliveryNotificationOk

`func (o *MessageGet) GetDeliveryNotificationOk() (*bool, bool)`

GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNotification

`func (o *MessageGet) SetDeliveryNotification(v bool)`

SetDeliveryNotification sets DeliveryNotification field to given value.

### HasDeliveryNotification

`func (o *MessageGet) HasDeliveryNotification() bool`

HasDeliveryNotification returns a boolean if a field has been set.

### GetStatusCallbackUrl

`func (o *MessageGet) GetStatusCallbackUrl() string`

GetStatusCallbackUrl returns the StatusCallbackUrl field if non-nil, zero value otherwise.

### GetStatusCallbackUrlOk

`func (o *MessageGet) GetStatusCallbackUrlOk() (*string, bool)`

GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackUrl

`func (o *MessageGet) SetStatusCallbackUrl(v string)`

SetStatusCallbackUrl sets StatusCallbackUrl field to given value.

### HasStatusCallbackUrl

`func (o *MessageGet) HasStatusCallbackUrl() bool`

HasStatusCallbackUrl returns a boolean if a field has been set.

### GetQueuePriority

`func (o *MessageGet) GetQueuePriority() int32`

GetQueuePriority returns the QueuePriority field if non-nil, zero value otherwise.

### GetQueuePriorityOk

`func (o *MessageGet) GetQueuePriorityOk() (*int32, bool)`

GetQueuePriorityOk returns a tuple with the QueuePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePriority

`func (o *MessageGet) SetQueuePriority(v int32)`

SetQueuePriority sets QueuePriority field to given value.

### HasQueuePriority

`func (o *MessageGet) HasQueuePriority() bool`

HasQueuePriority returns a boolean if a field has been set.

### GetTags

`func (o *MessageGet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MessageGet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MessageGet) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MessageGet) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


