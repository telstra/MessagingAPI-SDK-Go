/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
	"fmt"
)

// MessageSentTo - The recipient's mobile number(s).
type MessageSentTo struct {
	ArrayOfString *[]string
	String        *string
}

// []stringAsMessageSentTo is a convenience function that returns []string wrapped in MessageSentTo
func ArrayOfStringAsMessageSentTo(v *[]string) MessageSentTo {
	return MessageSentTo{
		ArrayOfString: v,
	}
}

// stringAsMessageSentTo is a convenience function that returns string wrapped in MessageSentTo
func StringAsMessageSentTo(v *string) MessageSentTo {
	return MessageSentTo{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageSentTo) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(MessageSentTo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MessageSentTo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageSentTo) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MessageSentTo) GetActualInstance() interface{} {
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

type NullableMessageSentTo struct {
	value *MessageSentTo
	isSet bool
}

func (v NullableMessageSentTo) Get() *MessageSentTo {
	return v.value
}

func (v *NullableMessageSentTo) Set(val *MessageSentTo) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSentTo) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSentTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSentTo(val *MessageSentTo) *NullableMessageSentTo {
	return &NullableMessageSentTo{value: val, isSet: true}
}

func (v NullableMessageSentTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSentTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
