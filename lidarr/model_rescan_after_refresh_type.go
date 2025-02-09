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

// RescanAfterRefreshType the model 'RescanAfterRefreshType'
type RescanAfterRefreshType string

// List of RescanAfterRefreshType
const (
	RESCANAFTERREFRESHTYPE_ALWAYS RescanAfterRefreshType = "always"
	RESCANAFTERREFRESHTYPE_AFTER_MANUAL RescanAfterRefreshType = "afterManual"
	RESCANAFTERREFRESHTYPE_NEVER RescanAfterRefreshType = "never"
)

// All allowed values of RescanAfterRefreshType enum
var AllowedRescanAfterRefreshTypeEnumValues = []RescanAfterRefreshType{
	"always",
	"afterManual",
	"never",
}

func (v *RescanAfterRefreshType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RescanAfterRefreshType(value)
	for _, existing := range AllowedRescanAfterRefreshTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RescanAfterRefreshType", value)
}

// NewRescanAfterRefreshTypeFromValue returns a pointer to a valid RescanAfterRefreshType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRescanAfterRefreshTypeFromValue(v string) (*RescanAfterRefreshType, error) {
	ev := RescanAfterRefreshType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RescanAfterRefreshType: valid values are %v", v, AllowedRescanAfterRefreshTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RescanAfterRefreshType) IsValid() bool {
	for _, existing := range AllowedRescanAfterRefreshTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RescanAfterRefreshType value
func (v RescanAfterRefreshType) Ptr() *RescanAfterRefreshType {
	return &v
}

type NullableRescanAfterRefreshType struct {
	value *RescanAfterRefreshType
	isSet bool
}

func (v NullableRescanAfterRefreshType) Get() *RescanAfterRefreshType {
	return v.value
}

func (v *NullableRescanAfterRefreshType) Set(val *RescanAfterRefreshType) {
	v.value = val
	v.isSet = true
}

func (v NullableRescanAfterRefreshType) IsSet() bool {
	return v.isSet
}

func (v *NullableRescanAfterRefreshType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRescanAfterRefreshType(val *RescanAfterRefreshType) *NullableRescanAfterRefreshType {
	return &NullableRescanAfterRefreshType{value: val, isSet: true}
}

func (v NullableRescanAfterRefreshType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRescanAfterRefreshType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

