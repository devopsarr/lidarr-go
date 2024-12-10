/*
Lidarr

Lidarr API docs

API version: v2.7.1.4417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"fmt"
)

// ApplyTags the model 'ApplyTags'
type ApplyTags string

// List of ApplyTags
const (
	APPLYTAGS_ADD ApplyTags = "add"
	APPLYTAGS_REMOVE ApplyTags = "remove"
	APPLYTAGS_REPLACE ApplyTags = "replace"
)

// All allowed values of ApplyTags enum
var AllowedApplyTagsEnumValues = []ApplyTags{
	"add",
	"remove",
	"replace",
}

func (v *ApplyTags) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplyTags(value)
	for _, existing := range AllowedApplyTagsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplyTags", value)
}

// NewApplyTagsFromValue returns a pointer to a valid ApplyTags
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplyTagsFromValue(v string) (*ApplyTags, error) {
	ev := ApplyTags(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplyTags: valid values are %v", v, AllowedApplyTagsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplyTags) IsValid() bool {
	for _, existing := range AllowedApplyTagsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplyTags value
func (v ApplyTags) Ptr() *ApplyTags {
	return &v
}

type NullableApplyTags struct {
	value *ApplyTags
	isSet bool
}

func (v NullableApplyTags) Get() *ApplyTags {
	return v.value
}

func (v *NullableApplyTags) Set(val *ApplyTags) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyTags) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyTags(val *ApplyTags) *NullableApplyTags {
	return &NullableApplyTags{value: val, isSet: true}
}

func (v NullableApplyTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

