/*
Lidarr

Lidarr API docs

API version: v2.9.6.4552
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"fmt"
)

// NewItemMonitorTypes the model 'NewItemMonitorTypes'
type NewItemMonitorTypes string

// List of NewItemMonitorTypes
const (
	NEWITEMMONITORTYPES_ALL NewItemMonitorTypes = "all"
	NEWITEMMONITORTYPES_NONE NewItemMonitorTypes = "none"
	NEWITEMMONITORTYPES_NEW NewItemMonitorTypes = "new"
)

// All allowed values of NewItemMonitorTypes enum
var AllowedNewItemMonitorTypesEnumValues = []NewItemMonitorTypes{
	"all",
	"none",
	"new",
}

func (v *NewItemMonitorTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewItemMonitorTypes(value)
	for _, existing := range AllowedNewItemMonitorTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewItemMonitorTypes", value)
}

// NewNewItemMonitorTypesFromValue returns a pointer to a valid NewItemMonitorTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewItemMonitorTypesFromValue(v string) (*NewItemMonitorTypes, error) {
	ev := NewItemMonitorTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewItemMonitorTypes: valid values are %v", v, AllowedNewItemMonitorTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewItemMonitorTypes) IsValid() bool {
	for _, existing := range AllowedNewItemMonitorTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NewItemMonitorTypes value
func (v NewItemMonitorTypes) Ptr() *NewItemMonitorTypes {
	return &v
}

type NullableNewItemMonitorTypes struct {
	value *NewItemMonitorTypes
	isSet bool
}

func (v NullableNewItemMonitorTypes) Get() *NewItemMonitorTypes {
	return v.value
}

func (v *NullableNewItemMonitorTypes) Set(val *NewItemMonitorTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableNewItemMonitorTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableNewItemMonitorTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewItemMonitorTypes(val *NewItemMonitorTypes) *NullableNewItemMonitorTypes {
	return &NullableNewItemMonitorTypes{value: val, isSet: true}
}

func (v NullableNewItemMonitorTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewItemMonitorTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

