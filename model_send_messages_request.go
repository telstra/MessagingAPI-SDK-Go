/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
	"time"
)

// checks if the SendMessagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendMessagesRequest{}

// SendMessagesRequest struct for SendMessagesRequest
type SendMessagesRequest struct {
	To SendMessagesRequestTo `json:"to"`
	// When the recipient receives your message, you can choose whether they'll see a virtualNumber or senderName (paid plans only) in the **from** field.   * 04xxxxxxxx: Use one of the Virtual Numbers associated with your account. You'll also be able to receive SMS replies to this number.  * senderName: Choose a unique alphanumeric string of up to 11 characters (paid feature).
	From string `json:"from"`
	// Use this field to send an SMS. Your text message goes here.     Note: either messageContent or multimedia are required, or you can use both field if you want to send multimedia with text.
	MessageContent *string `json:"messageContent,omitempty"`
	// Use this field to send an MMS. Add your image, video or audio content here.   Note: either messageContent or multimedia are required, or you can use both field if you want to send multimedia with text.  Include a JSON payload with:  type: the type of multimedia content file you're sending (image, audio or video) followed by the file type. Use the format \"multimedia type/file type\", e.g. \"image/PNG\" or \"audio/MP3\". Supported file types: JPEG, BMP, GIF87a, GIF89a, PNG, MP3, WAV, MPEG, MPG, MP4, 3GP and US-ASCII.  fileName: the name of your multimedia file.  payload: the base64 encoded content. You can use [this online tool](https://elmah.io/tools/base64-image-encoder/) to encode an image, or [Base64 Guru](https://base64.guru/) to encode a video or audio file.
	Multimedia []Multimedia `json:"multimedia,omitempty"`
	// If the message is queued or unable to reach the recipient's device, tell us how many minutes the network should keep trying. Use an integer between 1 and 1440. If you don't set a value, we'll try for 10 minutes.
	RetryTimeout *int32 `json:"retryTimeout,omitempty"`
	// Don't want to send the message right away? Tell us what time you want to add it to the queue for sending instead.  Set the time in London Greenwich Mean Time (adjusting for any time difference) and use ISO format, e.g. \"2019-08-24T15:39:00Z\".  You can schedule a message up to 10 days into the future. If you specify a timestamp outside of this limit, the API will return a FIELD_INVALID error.
	ScheduleSend *time.Time `json:"scheduleSend,omitempty"`
	// To receive a notification when your SMS has been delivered, set this parameter to **true** and make sure you provide a **statusCallbackUrl** (paid feature).
	DeliveryNotification *bool `json:"deliveryNotification,omitempty"`
	// Tell us the URL you want the API to call when the status of your SMS updates.   To receive a status update, this field must be provided and deliveryNotification must be set to **true**.   The status will be either:   * **queued** – the message is in the queue for sending (default). * **sent** – your message has been sent from the server. * **expired** – we weren't able to send the message within the **retryTimeout** timeframe. * **delivered** – the message has successfully reached the recipient's device. Note that we will only be able to return this status if you set **deliveryNotification** to **true** (paid feature). * **undeliverable** – the delivery of your message failed (paid feature).  Sample callback response:  <pre><code class=\"language-sh\">{   \"to\":\"0476543210\",    \"from\":\"0401234567\",     \"timestamp\":\"2022-11-10T05:06:42.823Z\",    \"messageId\":\"1520b774-46b0-4415-a05f-7bcb1c032c59\",    \"status\":\"delivered\"  }</code></pre>
	StatusCallbackUrl *string `json:"statusCallbackUrl,omitempty"`
	// Create your own tags and use them to fetch and sort your messages through our other endpoints. You can assign up to 10 tags per message.
	Tags []string `json:"tags,omitempty"`
}

// NewSendMessagesRequest instantiates a new SendMessagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMessagesRequest(to SendMessagesRequestTo, from string) *SendMessagesRequest {
	this := SendMessagesRequest{}
	this.To = to
	this.From = from
	var retryTimeout int32 = 10
	this.RetryTimeout = &retryTimeout
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	return &this
}

// NewSendMessagesRequestWithDefaults instantiates a new SendMessagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMessagesRequestWithDefaults() *SendMessagesRequest {
	this := SendMessagesRequest{}
	var retryTimeout int32 = 10
	this.RetryTimeout = &retryTimeout
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	return &this
}

// GetTo returns the To field value
func (o *SendMessagesRequest) GetTo() SendMessagesRequestTo {
	if o == nil {
		var ret SendMessagesRequestTo
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetToOk() (*SendMessagesRequestTo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *SendMessagesRequest) SetTo(v SendMessagesRequestTo) {
	o.To = v
}

// GetFrom returns the From field value
func (o *SendMessagesRequest) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *SendMessagesRequest) SetFrom(v string) {
	o.From = v
}

// GetMessageContent returns the MessageContent field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetMessageContent() string {
	if o == nil || IsNil(o.MessageContent) {
		var ret string
		return ret
	}
	return *o.MessageContent
}

// GetMessageContentOk returns a tuple with the MessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetMessageContentOk() (*string, bool) {
	if o == nil || IsNil(o.MessageContent) {
		return nil, false
	}
	return o.MessageContent, true
}

// HasMessageContent returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasMessageContent() bool {
	if o != nil && !IsNil(o.MessageContent) {
		return true
	}

	return false
}

// SetMessageContent gets a reference to the given string and assigns it to the MessageContent field.
func (o *SendMessagesRequest) SetMessageContent(v string) {
	o.MessageContent = &v
}

// GetMultimedia returns the Multimedia field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetMultimedia() []Multimedia {
	if o == nil || IsNil(o.Multimedia) {
		var ret []Multimedia
		return ret
	}
	return o.Multimedia
}

// GetMultimediaOk returns a tuple with the Multimedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetMultimediaOk() ([]Multimedia, bool) {
	if o == nil || IsNil(o.Multimedia) {
		return nil, false
	}
	return o.Multimedia, true
}

// HasMultimedia returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasMultimedia() bool {
	if o != nil && !IsNil(o.Multimedia) {
		return true
	}

	return false
}

// SetMultimedia gets a reference to the given []Multimedia and assigns it to the Multimedia field.
func (o *SendMessagesRequest) SetMultimedia(v []Multimedia) {
	o.Multimedia = v
}

// GetRetryTimeout returns the RetryTimeout field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetRetryTimeout() int32 {
	if o == nil || IsNil(o.RetryTimeout) {
		var ret int32
		return ret
	}
	return *o.RetryTimeout
}

// GetRetryTimeoutOk returns a tuple with the RetryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetRetryTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryTimeout) {
		return nil, false
	}
	return o.RetryTimeout, true
}

// HasRetryTimeout returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasRetryTimeout() bool {
	if o != nil && !IsNil(o.RetryTimeout) {
		return true
	}

	return false
}

// SetRetryTimeout gets a reference to the given int32 and assigns it to the RetryTimeout field.
func (o *SendMessagesRequest) SetRetryTimeout(v int32) {
	o.RetryTimeout = &v
}

// GetScheduleSend returns the ScheduleSend field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetScheduleSend() time.Time {
	if o == nil || IsNil(o.ScheduleSend) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleSend
}

// GetScheduleSendOk returns a tuple with the ScheduleSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetScheduleSendOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleSend) {
		return nil, false
	}
	return o.ScheduleSend, true
}

// HasScheduleSend returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasScheduleSend() bool {
	if o != nil && !IsNil(o.ScheduleSend) {
		return true
	}

	return false
}

// SetScheduleSend gets a reference to the given time.Time and assigns it to the ScheduleSend field.
func (o *SendMessagesRequest) SetScheduleSend(v time.Time) {
	o.ScheduleSend = &v
}

// GetDeliveryNotification returns the DeliveryNotification field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetDeliveryNotification() bool {
	if o == nil || IsNil(o.DeliveryNotification) {
		var ret bool
		return ret
	}
	return *o.DeliveryNotification
}

// GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetDeliveryNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveryNotification) {
		return nil, false
	}
	return o.DeliveryNotification, true
}

// HasDeliveryNotification returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasDeliveryNotification() bool {
	if o != nil && !IsNil(o.DeliveryNotification) {
		return true
	}

	return false
}

// SetDeliveryNotification gets a reference to the given bool and assigns it to the DeliveryNotification field.
func (o *SendMessagesRequest) SetDeliveryNotification(v bool) {
	o.DeliveryNotification = &v
}

// GetStatusCallbackUrl returns the StatusCallbackUrl field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetStatusCallbackUrl() string {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		var ret string
		return ret
	}
	return *o.StatusCallbackUrl
}

// GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetStatusCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		return nil, false
	}
	return o.StatusCallbackUrl, true
}

// HasStatusCallbackUrl returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasStatusCallbackUrl() bool {
	if o != nil && !IsNil(o.StatusCallbackUrl) {
		return true
	}

	return false
}

// SetStatusCallbackUrl gets a reference to the given string and assigns it to the StatusCallbackUrl field.
func (o *SendMessagesRequest) SetStatusCallbackUrl(v string) {
	o.StatusCallbackUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SendMessagesRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMessagesRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SendMessagesRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SendMessagesRequest) SetTags(v []string) {
	o.Tags = v
}

func (o SendMessagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendMessagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["to"] = o.To
	toSerialize["from"] = o.From
	if !IsNil(o.MessageContent) {
		toSerialize["messageContent"] = o.MessageContent
	}
	if !IsNil(o.Multimedia) {
		toSerialize["multimedia"] = o.Multimedia
	}
	if !IsNil(o.RetryTimeout) {
		toSerialize["retryTimeout"] = o.RetryTimeout
	}
	if !IsNil(o.ScheduleSend) {
		toSerialize["scheduleSend"] = o.ScheduleSend
	}
	if !IsNil(o.DeliveryNotification) {
		toSerialize["deliveryNotification"] = o.DeliveryNotification
	}
	if !IsNil(o.StatusCallbackUrl) {
		toSerialize["statusCallbackUrl"] = o.StatusCallbackUrl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableSendMessagesRequest struct {
	value *SendMessagesRequest
	isSet bool
}

func (v NullableSendMessagesRequest) Get() *SendMessagesRequest {
	return v.value
}

func (v *NullableSendMessagesRequest) Set(val *SendMessagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessagesRequest(val *SendMessagesRequest) *NullableSendMessagesRequest {
	return &NullableSendMessagesRequest{value: val, isSet: true}
}

func (v NullableSendMessagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
