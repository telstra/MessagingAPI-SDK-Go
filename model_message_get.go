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

// checks if the MessageGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageGet{}

// MessageGet struct for MessageGet
type MessageGet struct {
	// Use this UUID with our other endpoints to fetch, update or delete the message.
	MessageId *string `json:"messageId,omitempty"`
	// The status will be either queued, sent, delivered, expired or undeliverable.
	Status *string `json:"status,omitempty"`
	// The time you submitted the message to the queue for sending.
	CreateTimestamp *time.Time `json:"createTimestamp,omitempty"`
	// The time the message was sent from the server.
	SentTimestamp *time.Time `json:"sentTimestamp,omitempty"`
	// The time the message was received by the recipient's device.
	ReceivedTimestamp *time.Time `json:"receivedTimestamp,omitempty"`
	// The recipient's mobile number.
	To *string `json:"to,omitempty"`
	// This will be either one of your Virtual Numbers or your senderName.
	From *string `json:"from,omitempty"`
	// The content of the message.
	MessageContent *string `json:"messageContent,omitempty"`
	// The multimedia content of the message (MMS only). It will include:
	Multimedia []MultimediaGet `json:"multimedia,omitempty"`
	// * **outgoing** – messages sent from your account. * **incoming** – messages sent to your account.
	Direction *string `json:"direction,omitempty"`
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

// NewMessageGet instantiates a new MessageGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageGet() *MessageGet {
	this := MessageGet{}
	var retryTimeout int32 = 10
	this.RetryTimeout = &retryTimeout
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	var queuePriority int32 = 2
	this.QueuePriority = &queuePriority
	return &this
}

// NewMessageGetWithDefaults instantiates a new MessageGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageGetWithDefaults() *MessageGet {
	this := MessageGet{}
	var retryTimeout int32 = 10
	this.RetryTimeout = &retryTimeout
	var deliveryNotification bool = false
	this.DeliveryNotification = &deliveryNotification
	var queuePriority int32 = 2
	this.QueuePriority = &queuePriority
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *MessageGet) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *MessageGet) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *MessageGet) SetMessageId(v string) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessageGet) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessageGet) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MessageGet) SetStatus(v string) {
	o.Status = &v
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *MessageGet) GetCreateTimestamp() time.Time {
	if o == nil || IsNil(o.CreateTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetCreateTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTimestamp) {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *MessageGet) HasCreateTimestamp() bool {
	if o != nil && !IsNil(o.CreateTimestamp) {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given time.Time and assigns it to the CreateTimestamp field.
func (o *MessageGet) SetCreateTimestamp(v time.Time) {
	o.CreateTimestamp = &v
}

// GetSentTimestamp returns the SentTimestamp field value if set, zero value otherwise.
func (o *MessageGet) GetSentTimestamp() time.Time {
	if o == nil || IsNil(o.SentTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.SentTimestamp
}

// GetSentTimestampOk returns a tuple with the SentTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetSentTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentTimestamp) {
		return nil, false
	}
	return o.SentTimestamp, true
}

// HasSentTimestamp returns a boolean if a field has been set.
func (o *MessageGet) HasSentTimestamp() bool {
	if o != nil && !IsNil(o.SentTimestamp) {
		return true
	}

	return false
}

// SetSentTimestamp gets a reference to the given time.Time and assigns it to the SentTimestamp field.
func (o *MessageGet) SetSentTimestamp(v time.Time) {
	o.SentTimestamp = &v
}

// GetReceivedTimestamp returns the ReceivedTimestamp field value if set, zero value otherwise.
func (o *MessageGet) GetReceivedTimestamp() time.Time {
	if o == nil || IsNil(o.ReceivedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ReceivedTimestamp
}

// GetReceivedTimestampOk returns a tuple with the ReceivedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetReceivedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReceivedTimestamp) {
		return nil, false
	}
	return o.ReceivedTimestamp, true
}

// HasReceivedTimestamp returns a boolean if a field has been set.
func (o *MessageGet) HasReceivedTimestamp() bool {
	if o != nil && !IsNil(o.ReceivedTimestamp) {
		return true
	}

	return false
}

// SetReceivedTimestamp gets a reference to the given time.Time and assigns it to the ReceivedTimestamp field.
func (o *MessageGet) SetReceivedTimestamp(v time.Time) {
	o.ReceivedTimestamp = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MessageGet) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MessageGet) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *MessageGet) SetTo(v string) {
	o.To = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MessageGet) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MessageGet) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *MessageGet) SetFrom(v string) {
	o.From = &v
}

// GetMessageContent returns the MessageContent field value if set, zero value otherwise.
func (o *MessageGet) GetMessageContent() string {
	if o == nil || IsNil(o.MessageContent) {
		var ret string
		return ret
	}
	return *o.MessageContent
}

// GetMessageContentOk returns a tuple with the MessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetMessageContentOk() (*string, bool) {
	if o == nil || IsNil(o.MessageContent) {
		return nil, false
	}
	return o.MessageContent, true
}

// HasMessageContent returns a boolean if a field has been set.
func (o *MessageGet) HasMessageContent() bool {
	if o != nil && !IsNil(o.MessageContent) {
		return true
	}

	return false
}

// SetMessageContent gets a reference to the given string and assigns it to the MessageContent field.
func (o *MessageGet) SetMessageContent(v string) {
	o.MessageContent = &v
}

// GetMultimedia returns the Multimedia field value if set, zero value otherwise.
func (o *MessageGet) GetMultimedia() []MultimediaGet {
	if o == nil || IsNil(o.Multimedia) {
		var ret []MultimediaGet
		return ret
	}
	return o.Multimedia
}

// GetMultimediaOk returns a tuple with the Multimedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetMultimediaOk() ([]MultimediaGet, bool) {
	if o == nil || IsNil(o.Multimedia) {
		return nil, false
	}
	return o.Multimedia, true
}

// HasMultimedia returns a boolean if a field has been set.
func (o *MessageGet) HasMultimedia() bool {
	if o != nil && !IsNil(o.Multimedia) {
		return true
	}

	return false
}

// SetMultimedia gets a reference to the given []MultimediaGet and assigns it to the Multimedia field.
func (o *MessageGet) SetMultimedia(v []MultimediaGet) {
	o.Multimedia = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MessageGet) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MessageGet) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MessageGet) SetDirection(v string) {
	o.Direction = &v
}

// GetRetryTimeout returns the RetryTimeout field value if set, zero value otherwise.
func (o *MessageGet) GetRetryTimeout() int32 {
	if o == nil || IsNil(o.RetryTimeout) {
		var ret int32
		return ret
	}
	return *o.RetryTimeout
}

// GetRetryTimeoutOk returns a tuple with the RetryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetRetryTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryTimeout) {
		return nil, false
	}
	return o.RetryTimeout, true
}

// HasRetryTimeout returns a boolean if a field has been set.
func (o *MessageGet) HasRetryTimeout() bool {
	if o != nil && !IsNil(o.RetryTimeout) {
		return true
	}

	return false
}

// SetRetryTimeout gets a reference to the given int32 and assigns it to the RetryTimeout field.
func (o *MessageGet) SetRetryTimeout(v int32) {
	o.RetryTimeout = &v
}

// GetScheduleSend returns the ScheduleSend field value if set, zero value otherwise.
func (o *MessageGet) GetScheduleSend() time.Time {
	if o == nil || IsNil(o.ScheduleSend) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleSend
}

// GetScheduleSendOk returns a tuple with the ScheduleSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetScheduleSendOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduleSend) {
		return nil, false
	}
	return o.ScheduleSend, true
}

// HasScheduleSend returns a boolean if a field has been set.
func (o *MessageGet) HasScheduleSend() bool {
	if o != nil && !IsNil(o.ScheduleSend) {
		return true
	}

	return false
}

// SetScheduleSend gets a reference to the given time.Time and assigns it to the ScheduleSend field.
func (o *MessageGet) SetScheduleSend(v time.Time) {
	o.ScheduleSend = &v
}

// GetDeliveryNotification returns the DeliveryNotification field value if set, zero value otherwise.
func (o *MessageGet) GetDeliveryNotification() bool {
	if o == nil || IsNil(o.DeliveryNotification) {
		var ret bool
		return ret
	}
	return *o.DeliveryNotification
}

// GetDeliveryNotificationOk returns a tuple with the DeliveryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetDeliveryNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveryNotification) {
		return nil, false
	}
	return o.DeliveryNotification, true
}

// HasDeliveryNotification returns a boolean if a field has been set.
func (o *MessageGet) HasDeliveryNotification() bool {
	if o != nil && !IsNil(o.DeliveryNotification) {
		return true
	}

	return false
}

// SetDeliveryNotification gets a reference to the given bool and assigns it to the DeliveryNotification field.
func (o *MessageGet) SetDeliveryNotification(v bool) {
	o.DeliveryNotification = &v
}

// GetStatusCallbackUrl returns the StatusCallbackUrl field value if set, zero value otherwise.
func (o *MessageGet) GetStatusCallbackUrl() string {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		var ret string
		return ret
	}
	return *o.StatusCallbackUrl
}

// GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetStatusCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCallbackUrl) {
		return nil, false
	}
	return o.StatusCallbackUrl, true
}

// HasStatusCallbackUrl returns a boolean if a field has been set.
func (o *MessageGet) HasStatusCallbackUrl() bool {
	if o != nil && !IsNil(o.StatusCallbackUrl) {
		return true
	}

	return false
}

// SetStatusCallbackUrl gets a reference to the given string and assigns it to the StatusCallbackUrl field.
func (o *MessageGet) SetStatusCallbackUrl(v string) {
	o.StatusCallbackUrl = &v
}

// GetQueuePriority returns the QueuePriority field value if set, zero value otherwise.
func (o *MessageGet) GetQueuePriority() int32 {
	if o == nil || IsNil(o.QueuePriority) {
		var ret int32
		return ret
	}
	return *o.QueuePriority
}

// GetQueuePriorityOk returns a tuple with the QueuePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetQueuePriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.QueuePriority) {
		return nil, false
	}
	return o.QueuePriority, true
}

// HasQueuePriority returns a boolean if a field has been set.
func (o *MessageGet) HasQueuePriority() bool {
	if o != nil && !IsNil(o.QueuePriority) {
		return true
	}

	return false
}

// SetQueuePriority gets a reference to the given int32 and assigns it to the QueuePriority field.
func (o *MessageGet) SetQueuePriority(v int32) {
	o.QueuePriority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MessageGet) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageGet) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MessageGet) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MessageGet) SetTags(v []string) {
	o.Tags = v
}

func (o MessageGet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreateTimestamp) {
		toSerialize["createTimestamp"] = o.CreateTimestamp
	}
	if !IsNil(o.SentTimestamp) {
		toSerialize["sentTimestamp"] = o.SentTimestamp
	}
	if !IsNil(o.ReceivedTimestamp) {
		toSerialize["receivedTimestamp"] = o.ReceivedTimestamp
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
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
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

type NullableMessageGet struct {
	value *MessageGet
	isSet bool
}

func (v NullableMessageGet) Get() *MessageGet {
	return v.value
}

func (v *NullableMessageGet) Set(val *MessageGet) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageGet) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageGet(val *MessageGet) *NullableMessageGet {
	return &NullableMessageGet{value: val, isSet: true}
}

func (v NullableMessageGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
