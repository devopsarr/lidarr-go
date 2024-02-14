/*
Lidarr

Lidarr API docs

API version: v2.1.7.4030
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// TagDifference struct for TagDifference
type TagDifference struct {
	Field NullableString `json:"field,omitempty"`
	OldValue NullableString `json:"oldValue,omitempty"`
	NewValue NullableString `json:"newValue,omitempty"`
}

// NewTagDifference instantiates a new TagDifference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagDifference() *TagDifference {
	this := TagDifference{}
	return &this
}

// NewTagDifferenceWithDefaults instantiates a new TagDifference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagDifferenceWithDefaults() *TagDifference {
	this := TagDifference{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDifference) GetField() string {
	if o == nil || isNil(o.Field.Get()) {
		var ret string
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDifference) GetFieldOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *TagDifference) HasField() bool {
	if o != nil && o.Field.IsSet() {
		return true
	}

	return false
}

// SetField gets a reference to the given NullableString and assigns it to the Field field.
func (o *TagDifference) SetField(v string) {
	o.Field.Set(&v)
}
// SetFieldNil sets the value for Field to be an explicit nil
func (o *TagDifference) SetFieldNil() {
	o.Field.Set(nil)
}

// UnsetField ensures that no value is present for Field, not even an explicit nil
func (o *TagDifference) UnsetField() {
	o.Field.Unset()
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDifference) GetOldValue() string {
	if o == nil || isNil(o.OldValue.Get()) {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDifference) GetOldValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *TagDifference) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *TagDifference) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *TagDifference) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *TagDifference) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDifference) GetNewValue() string {
	if o == nil || isNil(o.NewValue.Get()) {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDifference) GetNewValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *TagDifference) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *TagDifference) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *TagDifference) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *TagDifference) UnsetNewValue() {
	o.NewValue.Unset()
}

func (o TagDifference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field.IsSet() {
		toSerialize["field"] = o.Field.Get()
	}
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["newValue"] = o.NewValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTagDifference struct {
	value *TagDifference
	isSet bool
}

func (v NullableTagDifference) Get() *TagDifference {
	return v.value
}

func (v *NullableTagDifference) Set(val *TagDifference) {
	v.value = val
	v.isSet = true
}

func (v NullableTagDifference) IsSet() bool {
	return v.isSet
}

func (v *NullableTagDifference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagDifference(val *TagDifference) *NullableTagDifference {
	return &NullableTagDifference{value: val, isSet: true}
}

func (v NullableTagDifference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagDifference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


