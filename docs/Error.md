# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Unique code of the error | 
**Issue** | **string** | The reason for the error | 
**SuggestedAction** | **string** | Suggest practical actions for this particular issue. | 
**Field** | Pointer to **string** | The field that caused the error | [optional] 
**Value** | Pointer to **string** | The value of the field that caused the error | [optional] 
**Location** | Pointer to **string** | The location of the field that caused the error. | [optional] 
**Link** | Pointer to **string** | URI for detailed information related to this error for the developer. | [optional] 

## Methods

### NewError

`func NewError(code string, issue string, suggestedAction string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.


### GetIssue

`func (o *Error) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *Error) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *Error) SetIssue(v string)`

SetIssue sets Issue field to given value.


### GetSuggestedAction

`func (o *Error) GetSuggestedAction() string`

GetSuggestedAction returns the SuggestedAction field if non-nil, zero value otherwise.

### GetSuggestedActionOk

`func (o *Error) GetSuggestedActionOk() (*string, bool)`

GetSuggestedActionOk returns a tuple with the SuggestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAction

`func (o *Error) SetSuggestedAction(v string)`

SetSuggestedAction sets SuggestedAction field to given value.


### GetField

`func (o *Error) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Error) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Error) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Error) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValue

`func (o *Error) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Error) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Error) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Error) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLocation

`func (o *Error) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Error) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Error) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Error) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLink

`func (o *Error) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Error) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Error) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Error) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


