# SendMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**SendMessagesRequestTo**](SendMessagesRequestTo.md) |  | 
**From** | **string** | When the recipient receives your message, you can choose whether they&#39;ll see a virtualNumber or senderName (paid plans only) in the **from** field.   04xxxxxxxx: Use one of the Virtual Numbers associated with your account. You&#39;ll also be able to receive SMS replies to this number.  * senderName: Choose a unique alphanumeric string of up to 11 characters (paid feature).  | 
**MessageContent** | Pointer to **string** | Use this field to send an SMS. Your text message goes here.     Note: either messageContent or multimedia are required, or you can use both field if you want to send multimedia with text.  | [optional] 
**Multimedia** | Pointer to [**[]Multimedia**](Multimedia.md) | Use this field to send an MMS. Add your image, video or audio content here.   Note: either messageContent or multimedia are required, or you can use both field if you want to send multimedia with text.  Include a JSON payload with:  type: the type of multimedia content file you&#39;re sending (image, audio or video) followed by the file type. Use the format \&quot;multimedia type/file type\&quot;, e.g. \&quot;image/PNG\&quot; or \&quot;audio/MP3\&quot;. Supported file types: JPEG, BMP, GIF87a, GIF89a, PNG, MP3, WAV, MPEG, MPG, MP4, 3GP and US-ASCII.  fileName: the name of your multimedia file.  payload: the base64 encoded content. You can use [this online tool](https://elmah.io/tools/base64-image-encoder/) to encode an image, or [Base64 Guru](https://base64.guru/) to encode a video or audio file.  | [optional] 
**RetryTimeout** | Pointer to **int32** | If the message is queued or unable to reach the recipient&#39;s device, tell us how many minutes the network should keep trying. Use an integer between 1 and 1440. If you don&#39;t set a value, we&#39;ll try for 10 minutes.  | [optional] [default to 10]
**ScheduleSend** | Pointer to **time.Time** | Don&#39;t want to send the message right away? Tell us what time you want to add it to the queue for sending instead.  Set the time in London Greenwich Mean Time (adjusting for any time difference) and use ISO format, e.g. \&quot;2019-08-24T15:39:00Z\&quot;.  You can schedule a message up to 10 days into the future. If you specify a timestamp outside of this limit, the API will return a FIELD_INVALID error.  | [optional] 
**DeliveryNotification** | Pointer to **bool** | To receive a notification when your SMS has been delivered, set this parameter to **true** and make sure you provide a **statusCallbackUrl** (paid feature).  | [optional] [default to false]
**StatusCallbackUrl** | Pointer to **string** | Tell us the URL you want the API to call when the status of your SMS updates.   To receive a status update, this field must be provided and deliveryNotification must be set to **true**.   The status will be either:   * **queued** – the message is in the queue for sending (default). * **sent** – your message has been sent from the server. * **expired** – we weren&#39;t able to send the message within the **retryTimeout** timeframe. * **delivered** – the message has successfully reached the recipient&#39;s device. Note that we will only be able to return this status if you set **deliveryNotification** to **true** (paid feature). * **undeliverable** – the delivery of your message failed (paid feature).  Sample callback response:  &lt;pre&gt;&lt;code class&#x3D;\&quot;language-sh\&quot;&gt;{   \&quot;to\&quot;:\&quot;0476543210\&quot;,    \&quot;from\&quot;:\&quot;0401234567\&quot;,     \&quot;timestamp\&quot;:\&quot;2022-11-10T05:06:42.823Z\&quot;,    \&quot;messageId\&quot;:\&quot;1520b774-46b0-4415-a05f-7bcb1c032c59\&quot;,    \&quot;status\&quot;:\&quot;delivered\&quot;  }&lt;/code&gt;&lt;/pre&gt;  | [optional] 
**Tags** | Pointer to **[]string** | Create your own tags and use them to fetch and sort your messages through our other endpoints. You can assign up to 10 tags per message.  | [optional] 

## Methods

### NewSendMessagesRequest

`func NewSendMessagesRequest(to SendMessagesRequestTo, from string, ) *SendMessagesRequest`

NewSendMessagesRequest instantiates a new SendMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessagesRequestWithDefaults

`func NewSendMessagesRequestWithDefaults() *SendMessagesRequest`

NewSendMessagesRequestWithDefaults instantiates a new SendMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SendMessagesRequest) GetTo() SendMessagesRequestTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendMessagesRequest) GetToOk() (*SendMessagesRequestTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendMessagesRequest) SetTo(v SendMessagesRequestTo)`

SetTo sets To field to given value.


### GetFrom

`func (o *SendMessagesRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendMessagesRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendMessagesRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetMessageContent

`func (o *SendMessagesRequest) GetMessageContent() string`

GetMessageContent returns the MessageContent field if non-nil, zero value otherwise.

### GetMessageContentOk

`func (o *SendMessagesRequest) GetMessageContentOk() (*string, bool)`

GetMessageContentOk returns a tuple with the MessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContent

`func (o *SendMessagesRequest) SetMessageContent(v string)`

SetMessageContent sets MessageContent field to given value.

### HasMessageContent

`func (o *SendMessagesRequest) HasMessageContent() bool`

HasMessageContent returns a boolean if a field has been set.

### GetMultimedia

`func (o *SendMessagesRequest) GetMultimedia() []Multimedia`

GetMultimedia returns the Multimedia field if non-nil, zero value otherwise.

### GetMultimediaOk

`func (o *SendMessagesRequest) GetMultimediaOk() (*[]Multimedia, bool)`

GetMultimediaOk returns a tuple with the Multimedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultimedia

`func (o *SendMessagesRequest) SetMultimedia(v []Multimedia)`

SetMultimedia sets Multimedia field to given value.

### HasMultimedia

`func (o *SendMessagesRequest) HasMultimedia() bool`

HasMultimedia returns a boolean if a field has been set.

### GetRetryTimeout

`func (o *SendMessagesRequest) GetRetryTimeout() int32`

GetRetryTimeout returns the RetryTimeout field if non-nil, zero value otherwise.

### GetRetryTimeoutOk

`func (o *SendMessagesRequest) GetRetryTimeoutOk() (*int32, bool)`

GetRetryTimeoutOk returns a tuple with the RetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimeout

`func (o *SendMessagesRequest) SetRetryTimeout(v int32)`

SetRetryTimeout sets RetryTimeout field to given value.

### HasRetryTimeout

`func (o *SendMessagesRequest) HasRetryTimeout() bool`

HasRetryTimeout returns a boolean if a field has been set.

### GetScheduleSend

`func (o *SendMessagesRequest) GetScheduleSend() time.Time`

GetScheduleSend returns the ScheduleSend field if non-nil, zero value otherwise.

### GetScheduleSendOk

`func (o *SendMessagesRequest) GetScheduleSendOk() (*time.Time, bool)`

GetScheduleSendOk returns a tuple with the ScheduleSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSend

`func (o *SendMessagesRequest) SetScheduleSend(v time.Time)`

SetScheduleSend sets ScheduleSend field to given value.

### HasScheduleSend

`func (o *SendMessagesRequest) HasScheduleSend() bool`

HasScheduleSend returns a boolean if a field has been set.

### GetDeliveryNotification

`func (o *SendMessagesRequest) GetDeliveryNotification() bool`

GetDeliveryNotification returns the DeliveryNotification field if non-nil, zero value otherwise.

### GetDeliveryNotificationOk

`func (o *SendMessagesRequest) GetDeliveryNotificationOk() (*bool, bool)`

GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNotification

`func (o *SendMessagesRequest) SetDeliveryNotification(v bool)`

SetDeliveryNotification sets DeliveryNotification field to given value.

### HasDeliveryNotification

`func (o *SendMessagesRequest) HasDeliveryNotification() bool`

HasDeliveryNotification returns a boolean if a field has been set.

### GetStatusCallbackUrl

`func (o *SendMessagesRequest) GetStatusCallbackUrl() string`

GetStatusCallbackUrl returns the StatusCallbackUrl field if non-nil, zero value otherwise.

### GetStatusCallbackUrlOk

`func (o *SendMessagesRequest) GetStatusCallbackUrlOk() (*string, bool)`

GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackUrl

`func (o *SendMessagesRequest) SetStatusCallbackUrl(v string)`

SetStatusCallbackUrl sets StatusCallbackUrl field to given value.

### HasStatusCallbackUrl

`func (o *SendMessagesRequest) HasStatusCallbackUrl() bool`

HasStatusCallbackUrl returns a boolean if a field has been set.

### GetTags

`func (o *SendMessagesRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SendMessagesRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SendMessagesRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SendMessagesRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


