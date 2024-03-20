/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
	"fmt"
)

// CreateTrialNumbersRequestFreeTrialNumbers - These are the mobile numbers you want to message during your Free Trial. Write Australian numbers in national format (e.g. \"0412345678\"). Use a string for a single recipient, and an array of strings for multiple recipients, (e.g. [\"0412345678\", \"0487654321\"]).
type CreateTrialNumbersRequestFreeTrialNumbers struct {
	ArrayOfString *[]string
	String        *string
}

// []stringAsCreateTrialNumbersRequestFreeTrialNumbers is a convenience function that returns []string wrapped in CreateTrialNumbersRequestFreeTrialNumbers
func ArrayOfStringAsCreateTrialNumbersRequestFreeTrialNumbers(v *[]string) CreateTrialNumbersRequestFreeTrialNumbers {
	return CreateTrialNumbersRequestFreeTrialNumbers{
		ArrayOfString: v,
	}
}

// stringAsCreateTrialNumbersRequestFreeTrialNumbers is a convenience function that returns string wrapped in CreateTrialNumbersRequestFreeTrialNumbers
func StringAsCreateTrialNumbersRequestFreeTrialNumbers(v *string) CreateTrialNumbersRequestFreeTrialNumbers {
	return CreateTrialNumbersRequestFreeTrialNumbers{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateTrialNumbersRequestFreeTrialNumbers) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(CreateTrialNumbersRequestFreeTrialNumbers)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateTrialNumbersRequestFreeTrialNumbers)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateTrialNumbersRequestFreeTrialNumbers) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateTrialNumbersRequestFreeTrialNumbers) GetActualInstance() interface{} {
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

type NullableCreateTrialNumbersRequestFreeTrialNumbers struct {
	value *CreateTrialNumbersRequestFreeTrialNumbers
	isSet bool
}

func (v NullableCreateTrialNumbersRequestFreeTrialNumbers) Get() *CreateTrialNumbersRequestFreeTrialNumbers {
	return v.value
}

func (v *NullableCreateTrialNumbersRequestFreeTrialNumbers) Set(val *CreateTrialNumbersRequestFreeTrialNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTrialNumbersRequestFreeTrialNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTrialNumbersRequestFreeTrialNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTrialNumbersRequestFreeTrialNumbers(val *CreateTrialNumbersRequestFreeTrialNumbers) *NullableCreateTrialNumbersRequestFreeTrialNumbers {
	return &NullableCreateTrialNumbersRequestFreeTrialNumbers{value: val, isSet: true}
}

func (v NullableCreateTrialNumbersRequestFreeTrialNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTrialNumbersRequestFreeTrialNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
