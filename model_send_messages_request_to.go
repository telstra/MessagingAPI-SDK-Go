/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
	"fmt"
)

// SendMessagesRequestTo - This is the mobile number you want to send your message to. Write Australian numbers in national format (e.g. 0412345678) and international numbers (paid plans only) in E.164 format (e.g. +441234567890).   Use a string for a single recipient, and an array of string of multiple recipients, e.g. \"to\": [\"0412345678\", \"+441234567890\"]. If you're using the Free Trial, you can include up to 5 recipient numbers in the array. If you're using a paid plan, you can bulk send up to 500 messages at once.
type SendMessagesRequestTo struct {
	ArrayOfString *[]string
	String        *string
}

// []stringAsSendMessagesRequestTo is a convenience function that returns []string wrapped in SendMessagesRequestTo
func ArrayOfStringAsSendMessagesRequestTo(v *[]string) SendMessagesRequestTo {
	return SendMessagesRequestTo{
		ArrayOfString: v,
	}
}

// stringAsSendMessagesRequestTo is a convenience function that returns string wrapped in SendMessagesRequestTo
func StringAsSendMessagesRequestTo(v *string) SendMessagesRequestTo {
	return SendMessagesRequestTo{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SendMessagesRequestTo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SendMessagesRequestTo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SendMessagesRequestTo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SendMessagesRequestTo) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SendMessagesRequestTo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSendMessagesRequestTo struct {
	value *SendMessagesRequestTo
	isSet bool
}

func (v NullableSendMessagesRequestTo) Get() *SendMessagesRequestTo {
	return v.value
}

func (v *NullableSendMessagesRequestTo) Set(val *SendMessagesRequestTo) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMessagesRequestTo) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMessagesRequestTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMessagesRequestTo(val *SendMessagesRequestTo) *NullableSendMessagesRequestTo {
	return &NullableSendMessagesRequestTo{value: val, isSet: true}
}

func (v NullableSendMessagesRequestTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMessagesRequestTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
