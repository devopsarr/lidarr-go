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

// ImportListType the model 'ImportListType'
type ImportListType string

// List of ImportListType
const (
	IMPORTLISTTYPE_PROGRAM ImportListType = "program"
	IMPORTLISTTYPE_SPOTIFY ImportListType = "spotify"
	IMPORTLISTTYPE_LAST_FM ImportListType = "lastFm"
	IMPORTLISTTYPE_OTHER ImportListType = "other"
	IMPORTLISTTYPE_ADVANCED ImportListType = "advanced"
)

// All allowed values of ImportListType enum
var AllowedImportListTypeEnumValues = []ImportListType{
	"program",
	"spotify",
	"lastFm",
	"other",
	"advanced",
}

func (v *ImportListType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportListType(value)
	for _, existing := range AllowedImportListTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportListType", value)
}

// NewImportListTypeFromValue returns a pointer to a valid ImportListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportListTypeFromValue(v string) (*ImportListType, error) {
	ev := ImportListType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportListType: valid values are %v", v, AllowedImportListTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportListType) IsValid() bool {
	for _, existing := range AllowedImportListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportListType value
func (v ImportListType) Ptr() *ImportListType {
	return &v
}

type NullableImportListType struct {
	value *ImportListType
	isSet bool
}

func (v NullableImportListType) Get() *ImportListType {
	return v.value
}

func (v *NullableImportListType) Set(val *ImportListType) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListType) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListType(val *ImportListType) *NullableImportListType {
	return &NullableImportListType{value: val, isSet: true}
}

func (v NullableImportListType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

