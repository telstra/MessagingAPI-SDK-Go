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

// checks if the MessageUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageUpdate{}

// MessageUpdate struct for MessageUpdate
type MessageUpdate struct {
	// Use this UUID with our other endpoints to fetch, update or delete the message.
	MessageId *string `json:"messageId,omitempty"`
	// The status will be either queued, sent, delivered, expired or undeliverable.
	Status *string `json:"status,omitempty"`
	// The recipient's mobile number.
	To *string `json:"to,omitempty"`
	// This will be either one of your Virtual Numbers or your senderName.
	From *string `json:"from,omitempty"`
	// The content of the message.
	MessageContent *string `json:"messageContent,omitempty"`
	// The multimedia content of the message (MMS only). It will include:
	Multimedia []MultimediaGet `json:"multimedia,omitempty"`
	// How many minutes you asked the server to keep trying to send the message.
	RetryTimeout *int32 `json:"retryTimeout,omitempty"`
	// The time the message is scheduled to send.
	ScheduleSend *time.Time `json:"scheduleSend,omitempty"`
	// If set to **true**, you will receive a notification to the statusCallbackUrl when your message is delivered (paid feature).
	DeliveryNotification *bool `json:"deliveryNotification,omitempty"`
	// The URL the API will call when the status of the message changes.
	StatusCallbackUrl *string `json:"statusCallbackUrl,omitempty"`
	// The priority assigned to the message.
	QueuePriority *int32 `json:"queuePriority,omitempty"`
	// Any customisable tags assigned to the message.
	Tags []string `json:"tags,omitempty"`
}

// NewMessageUpdate instantiates a new MessageUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageUpdate() *MessageUpdate {
	this := MessageUpdate{}
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	var queuePriority int32 = 2
	this.QueuePriority = &queuePriority
	return &this
}

// NewMessageUpdateWithDefaults instantiates a new MessageUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageUpdateWithDefaults() *MessageUpdate {
	this := MessageUpdate{}
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	var queuePriority int32 = 2
	this.QueuePriority = &queuePriority
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *MessageUpdate) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *MessageUpdate) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *MessageUpdate) SetMessageId(v string) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessageUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessageUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MessageUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MessageUpdate) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MessageUpdate) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *MessageUpdate) SetTo(v string) {
	o.To = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MessageUpdate) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MessageUpdate) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *MessageUpdate) SetFrom(v string) {
	o.From = &v
}

// GetMessageContent returns the MessageContent field value if set, zero value otherwise.
func (o *MessageUpdate) GetMessageContent() string {
	if o == nil || IsNil(o.MessageContent) {
		var ret string
		return ret
	}
	return *o.MessageContent
}

// GetMessageContentOk returns a tuple with the MessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetMessageContentOk() (*string, bool) {
	if o == nil || IsNil(o.MessageContent) {
		return nil, false
	}
	return o.MessageContent, true
}

// HasMessageContent returns a boolean if a field has been set.
func (o *MessageUpdate) HasMessageContent() bool {
	if o != nil && !IsNil(o.MessageContent) {
		return true
	}

	return false
}

// SetMessageContent gets a reference to the given string and assigns it to the MessageContent field.
func (o *MessageUpdate) SetMessageContent(v string) {
	o.MessageContent = &v
}

// GetMultimedia returns the Multimedia field value if set, zero value otherwise.
func (o *MessageUpdate) GetMultimedia() []MultimediaGet {
	if o == nil || IsNil(o.Multimedia) {
		var ret []MultimediaGet
		return ret
	}
	return o.Multimedia
}

// GetMultimediaOk returns a tuple with the Multimedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetMultimediaOk() ([]MultimediaGet, bool) {
	if o == nil || IsNil(o.Multimedia) {
		return nil, false
	}
	return o.Multimedia, true
}

// HasMultimedia returns a boolean if a field has been set.
func (o *MessageUpdate) HasMultimedia() bool {
	if o != nil && !IsNil(o.Multimedia) {
		return true
	}

	return false
}

// SetMultimedia gets a reference to the given []MultimediaGet and assigns it to the Multimedia field.
func (o *MessageUpdate) SetMultimedia(v []MultimediaGet) {
	o.Multimedia = v
}

// GetRetryTimeout returns the RetryTimeout field value if set, zero value otherwise.
func (o *MessageUpdate) GetRetryTimeout() int32 {
	if o == nil || IsNil(o.RetryTimeout) {
		var ret int32
		return ret
	}
	return *o.RetryTimeout
}

// GetRetryTimeoutOk returns a tuple with the RetryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetRetryTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryTimeout) {
		return nil, false
	}
	return o.RetryTimeout, true
}

// HasRetryTimeout returns a boolean if a field has been set.
func (o *MessageUpdate) HasRetryTimeout() bool {
	if o != nil && !IsNil(o.RetryTimeout) {
		return true
	}

	return false
}

// SetRetryTimeout gets a reference to the given int32 and assigns it to the RetryTimeout field.
func (o *MessageUpdate) SetRetryTimeout(v int32) {
	o.RetryTimeout = &v
}

// GetScheduleSend returns the ScheduleSend field value if set, zero value otherwise.
func (o *MessageUpdate) GetScheduleSend() time.Time {
	if o == nil || IsNil(o.ScheduleSend) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleSend
}

// GetScheduleSendOk returns a tuple with the ScheduleSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetScheduleSendOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleSend) {
		return nil, false
	}
	return o.ScheduleSend, true
}

// HasScheduleSend returns a boolean if a field has been set.
func (o *MessageUpdate) HasScheduleSend() bool {
	if o != nil && !IsNil(o.ScheduleSend) {
		return true
	}

	return false
}

// SetScheduleSend gets a reference to the given time.Time and assigns it to the ScheduleSend field.
func (o *MessageUpdate) SetScheduleSend(v time.Time) {
	o.ScheduleSend = &v
}

// GetDeliveryNotification returns the DeliveryNotification field value if set, zero value otherwise.
func (o *MessageUpdate) GetDeliveryNotification() bool {
	if o == nil || IsNil(o.DeliveryNotification) {
		var ret bool
		return ret
	}
	return *o.DeliveryNotification
}

// GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetDeliveryNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveryNotification) {
		return nil, false
	}
	return o.DeliveryNotification, true
}

// HasDeliveryNotification returns a boolean if a field has been set.
func (o *MessageUpdate) HasDeliveryNotification() bool {
	if o != nil && !IsNil(o.DeliveryNotification) {
		return true
	}

	return false
}

// SetDeliveryNotification gets a reference to the given bool and assigns it to the DeliveryNotification field.
func (o *MessageUpdate) SetDeliveryNotification(v bool) {
	o.DeliveryNotification = &v
}

// GetStatusCallbackUrl returns the StatusCallbackUrl field value if set, zero value otherwise.
func (o *MessageUpdate) GetStatusCallbackUrl() string {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		var ret string
		return ret
	}
	return *o.StatusCallbackUrl
}

// GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetStatusCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		return nil, false
	}
	return o.StatusCallbackUrl, true
}

// HasStatusCallbackUrl returns a boolean if a field has been set.
func (o *MessageUpdate) HasStatusCallbackUrl() bool {
	if o != nil && !IsNil(o.StatusCallbackUrl) {
		return true
	}

	return false
}

// SetStatusCallbackUrl gets a reference to the given string and assigns it to the StatusCallbackUrl field.
func (o *MessageUpdate) SetStatusCallbackUrl(v string) {
	o.StatusCallbackUrl = &v
}

// GetQueuePriority returns the QueuePriority field value if set, zero value otherwise.
func (o *MessageUpdate) GetQueuePriority() int32 {
	if o == nil || IsNil(o.QueuePriority) {
		var ret int32
		return ret
	}
	return *o.QueuePriority
}

// GetQueuePriorityOk returns a tuple with the QueuePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetQueuePriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.QueuePriority) {
		return nil, false
	}
	return o.QueuePriority, true
}

// HasQueuePriority returns a boolean if a field has been set.
func (o *MessageUpdate) HasQueuePriority() bool {
	if o != nil && !IsNil(o.QueuePriority) {
		return true
	}

	return false
}

// SetQueuePriority gets a reference to the given int32 and assigns it to the QueuePriority field.
func (o *MessageUpdate) SetQueuePriority(v int32) {
	o.QueuePriority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MessageUpdate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MessageUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MessageUpdate) SetTags(v []string) {
	o.Tags = v
}

func (o MessageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageUpdate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.QueuePriority) {
		toSerialize["queuePriority"] = o.QueuePriority
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableMessageUpdate struct {
	value *MessageUpdate
	isSet bool
}

func (v NullableMessageUpdate) Get() *MessageUpdate {
	return v.value
}

func (v *NullableMessageUpdate) Set(val *MessageUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageUpdate(val *MessageUpdate) *NullableMessageUpdate {
	return &NullableMessageUpdate{value: val, isSet: true}
}

func (v NullableMessageUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
