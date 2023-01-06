# Medium

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMedium

`func NewMedium() *Medium`

NewMedium instantiates a new Medium object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediumWithDefaults

`func NewMediumWithDefaults() *Medium`

NewMediumWithDefaults instantiates a new Medium object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Medium) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Medium) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Medium) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Medium) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *Medium) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Medium) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Medium) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Medium) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Medium) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Medium) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFormat

`func (o *Medium) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Medium) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Medium) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Medium) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *Medium) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *Medium) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


