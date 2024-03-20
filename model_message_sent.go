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

// checks if the MessageSent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSent{}

// MessageSent struct for MessageSent
type MessageSent struct {
	MessageId *MessageSentMessageId `json:"messageId,omitempty"`
	// The status will be either queued, sent, delivered, expired or undeliverable.
	Status *string        `json:"status,omitempty"`
	To     *MessageSentTo `json:"to,omitempty"`
	// This will be either one of your Virtual Numbers or your senderName.
	From *string `json:"from,omitempty"`
	// The content of the message.
	MessageContent *string `json:"messageContent,omitempty"`
	// The multimedia content of the message (MMS only). It will include:
	Multimedia []MultimediaGet `json:"multimedia,omitempty"`
	// How many minutes you asked the server to keep trying to send the message.
	RetryTimeout *int32 `json:"retryTimeout,omitempty"`
	// The time (in Central Standard Time) the message is scheduled to send.
	ScheduleSend *time.Time `json:"scheduleSend,omitempty"`
	// If set to **true**, you will receive a notification to the statusCallbackUrl when your SMS is delivered (paid feature).
	DeliveryNotification *bool `json:"deliveryNotification,omitempty"`
	// The URL the API will call when the status of the message changes.
	StatusCallbackUrl *string `json:"statusCallbackUrl,omitempty"`
	// Any customisable tags assigned to the message.
	Tags []string `json:"tags,omitempty"`
}

// NewMessageSent instantiates a new MessageSent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSent() *MessageSent {
	this := MessageSent{}
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	return &this
}

// NewMessageSentWithDefaults instantiates a new MessageSent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSentWithDefaults() *MessageSent {
	this := MessageSent{}
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *MessageSent) GetMessageId() MessageSentMessageId {
	if o == nil || IsNil(o.MessageId) {
		var ret MessageSentMessageId
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetMessageIdOk() (*MessageSentMessageId, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *MessageSent) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given MessageSentMessageId and assigns it to the MessageId field.
func (o *MessageSent) SetMessageId(v MessageSentMessageId) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessageSent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessageSent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MessageSent) SetStatus(v string) {
	o.Status = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MessageSent) GetTo() MessageSentTo {
	if o == nil || IsNil(o.To) {
		var ret MessageSentTo
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetToOk() (*MessageSentTo, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MessageSent) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given MessageSentTo and assigns it to the To field.
func (o *MessageSent) SetTo(v MessageSentTo) {
	o.To = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MessageSent) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MessageSent) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *MessageSent) SetFrom(v string) {
	o.From = &v
}

// GetMessageContent returns the MessageContent field value if set, zero value otherwise.
func (o *MessageSent) GetMessageContent() string {
	if o == nil || IsNil(o.MessageContent) {
		var ret string
		return ret
	}
	return *o.MessageContent
}

// GetMessageContentOk returns a tuple with the MessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetMessageContentOk() (*string, bool) {
	if o == nil || IsNil(o.MessageContent) {
		return nil, false
	}
	return o.MessageContent, true
}

// HasMessageContent returns a boolean if a field has been set.
func (o *MessageSent) HasMessageContent() bool {
	if o != nil && !IsNil(o.MessageContent) {
		return true
	}

	return false
}

// SetMessageContent gets a reference to the given string and assigns it to the MessageContent field.
func (o *MessageSent) SetMessageContent(v string) {
	o.MessageContent = &v
}

// GetMultimedia returns the Multimedia field value if set, zero value otherwise.
func (o *MessageSent) GetMultimedia() []MultimediaGet {
	if o == nil || IsNil(o.Multimedia) {
		var ret []MultimediaGet
		return ret
	}
	return o.Multimedia
}

// GetMultimediaOk returns a tuple with the Multimedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetMultimediaOk() ([]MultimediaGet, bool) {
	if o == nil || IsNil(o.Multimedia) {
		return nil, false
	}
	return o.Multimedia, true
}

// HasMultimedia returns a boolean if a field has been set.
func (o *MessageSent) HasMultimedia() bool {
	if o != nil && !IsNil(o.Multimedia) {
		return true
	}

	return false
}

// SetMultimedia gets a reference to the given []MultimediaGet and assigns it to the Multimedia field.
func (o *MessageSent) SetMultimedia(v []MultimediaGet) {
	o.Multimedia = v
}

// GetRetryTimeout returns the RetryTimeout field value if set, zero value otherwise.
func (o *MessageSent) GetRetryTimeout() int32 {
	if o == nil || IsNil(o.RetryTimeout) {
		var ret int32
		return ret
	}
	return *o.RetryTimeout
}

// GetRetryTimeoutOk returns a tuple with the RetryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetRetryTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryTimeout) {
		return nil, false
	}
	return o.RetryTimeout, true
}

// HasRetryTimeout returns a boolean if a field has been set.
func (o *MessageSent) HasRetryTimeout() bool {
	if o != nil && !IsNil(o.RetryTimeout) {
		return true
	}

	return false
}

// SetRetryTimeout gets a reference to the given int32 and assigns it to the RetryTimeout field.
func (o *MessageSent) SetRetryTimeout(v int32) {
	o.RetryTimeout = &v
}

// GetScheduleSend returns the ScheduleSend field value if set, zero value otherwise.
func (o *MessageSent) GetScheduleSend() time.Time {
	if o == nil || IsNil(o.ScheduleSend) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleSend
}

// GetScheduleSendOk returns a tuple with the ScheduleSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetScheduleSendOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleSend) {
		return nil, false
	}
	return o.ScheduleSend, true
}

// HasScheduleSend returns a boolean if a field has been set.
func (o *MessageSent) HasScheduleSend() bool {
	if o != nil && !IsNil(o.ScheduleSend) {
		return true
	}

	return false
}

// SetScheduleSend gets a reference to the given time.Time and assigns it to the ScheduleSend field.
func (o *MessageSent) SetScheduleSend(v time.Time) {
	o.ScheduleSend = &v
}

// GetDeliveryNotification returns the DeliveryNotification field value if set, zero value otherwise.
func (o *MessageSent) GetDeliveryNotification() bool {
	if o == nil || IsNil(o.DeliveryNotification) {
		var ret bool
		return ret
	}
	return *o.DeliveryNotification
}

// GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetDeliveryNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveryNotification) {
		return nil, false
	}
	return o.DeliveryNotification, true
}

// HasDeliveryNotification returns a boolean if a field has been set.
func (o *MessageSent) HasDeliveryNotification() bool {
	if o != nil && !IsNil(o.DeliveryNotification) {
		return true
	}

	return false
}

// SetDeliveryNotification gets a reference to the given bool and assigns it to the DeliveryNotification field.
func (o *MessageSent) SetDeliveryNotification(v bool) {
	o.DeliveryNotification = &v
}

// GetStatusCallbackUrl returns the StatusCallbackUrl field value if set, zero value otherwise.
func (o *MessageSent) GetStatusCallbackUrl() string {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		var ret string
		return ret
	}
	return *o.StatusCallbackUrl
}

// GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetStatusCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		return nil, false
	}
	return o.StatusCallbackUrl, true
}

// HasStatusCallbackUrl returns a boolean if a field has been set.
func (o *MessageSent) HasStatusCallbackUrl() bool {
	if o != nil && !IsNil(o.StatusCallbackUrl) {
		return true
	}

	return false
}

// SetStatusCallbackUrl gets a reference to the given string and assigns it to the StatusCallbackUrl field.
func (o *MessageSent) SetStatusCallbackUrl(v string) {
	o.StatusCallbackUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MessageSent) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSent) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MessageSent) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MessageSent) SetTags(v []string) {
	o.Tags = v
}

func (o MessageSent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
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

type NullableMessageSent struct {
	value *MessageSent
	isSet bool
}

func (v NullableMessageSent) Get() *MessageSent {
	return v.value
}

func (v *NullableMessageSent) Set(val *MessageSent) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSent) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSent(val *MessageSent) *NullableMessageSent {
	return &NullableMessageSent{value: val, isSet: true}
}

func (v NullableMessageSent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
