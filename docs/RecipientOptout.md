# RecipientOptout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptoutNumber** | Pointer to **string** | The mobile number that sent the optout request. | [optional] 
**CreateTimestamp** | Pointer to **time.Time** | The date and time we received the optout request. | [optional] 

## Methods

### NewRecipientOptout

`func NewRecipientOptout() *RecipientOptout`

NewRecipientOptout instantiates a new RecipientOptout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientOptoutWithDefaults

`func NewRecipientOptoutWithDefaults() *RecipientOptout`

NewRecipientOptoutWithDefaults instantiates a new RecipientOptout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptoutNumber

`func (o *RecipientOptout) GetOptoutNumber() string`

GetOptoutNumber returns the OptoutNumber field if non-nil, zero value otherwise.

### GetOptoutNumberOk

`func (o *RecipientOptout) GetOptoutNumberOk() (*string, bool)`

GetOptoutNumberOk returns a tuple with the OptoutNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptoutNumber

`func (o *RecipientOptout) SetOptoutNumber(v string)`

SetOptoutNumber sets OptoutNumber field to given value.

### HasOptoutNumber

`func (o *RecipientOptout) HasOptoutNumber() bool`

HasOptoutNumber returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *RecipientOptout) GetCreateTimestamp() time.Time`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *RecipientOptout) GetCreateTimestampOk() (*time.Time, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *RecipientOptout) SetCreateTimestamp(v time.Time)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *RecipientOptout) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


