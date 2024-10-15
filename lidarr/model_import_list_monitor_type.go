/*
Lidarr

Lidarr API docs

API version: v2.6.4.4402
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"fmt"
)

// ImportListMonitorType the model 'ImportListMonitorType'
type ImportListMonitorType string

// List of ImportListMonitorType
const (
	IMPORTLISTMONITORTYPE_NONE ImportListMonitorType = "none"
	IMPORTLISTMONITORTYPE_SPECIFIC_ALBUM ImportListMonitorType = "specificAlbum"
	IMPORTLISTMONITORTYPE_ENTIRE_ARTIST ImportListMonitorType = "entireArtist"
)

// All allowed values of ImportListMonitorType enum
var AllowedImportListMonitorTypeEnumValues = []ImportListMonitorType{
	"none",
	"specificAlbum",
	"entireArtist",
}

func (v *ImportListMonitorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportListMonitorType(value)
	for _, existing := range AllowedImportListMonitorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportListMonitorType", value)
}

// NewImportListMonitorTypeFromValue returns a pointer to a valid ImportListMonitorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportListMonitorTypeFromValue(v string) (*ImportListMonitorType, error) {
	ev := ImportListMonitorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportListMonitorType: valid values are %v", v, AllowedImportListMonitorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportListMonitorType) IsValid() bool {
	for _, existing := range AllowedImportListMonitorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportListMonitorType value
func (v ImportListMonitorType) Ptr() *ImportListMonitorType {
	return &v
}

type NullableImportListMonitorType struct {
	value *ImportListMonitorType
	isSet bool
}

func (v NullableImportListMonitorType) Get() *ImportListMonitorType {
	return v.value
}

func (v *NullableImportListMonitorType) Set(val *ImportListMonitorType) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListMonitorType) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListMonitorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListMonitorType(val *ImportListMonitorType) *NullableImportListMonitorType {
	return &NullableImportListMonitorType{value: val, isSet: true}
}

func (v NullableImportListMonitorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListMonitorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

