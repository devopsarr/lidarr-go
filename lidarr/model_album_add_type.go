/*
Lidarr

Lidarr API docs

API version: v2.4.3.4248
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"fmt"
)

// AlbumAddType the model 'AlbumAddType'
type AlbumAddType string

// List of AlbumAddType
const (
	ALBUMADDTYPE_AUTOMATIC AlbumAddType = "automatic"
	ALBUMADDTYPE_MANUAL AlbumAddType = "manual"
)

// All allowed values of AlbumAddType enum
var AllowedAlbumAddTypeEnumValues = []AlbumAddType{
	"automatic",
	"manual",
}

func (v *AlbumAddType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlbumAddType(value)
	for _, existing := range AllowedAlbumAddTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlbumAddType", value)
}

// NewAlbumAddTypeFromValue returns a pointer to a valid AlbumAddType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlbumAddTypeFromValue(v string) (*AlbumAddType, error) {
	ev := AlbumAddType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlbumAddType: valid values are %v", v, AllowedAlbumAddTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlbumAddType) IsValid() bool {
	for _, existing := range AllowedAlbumAddTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlbumAddType value
func (v AlbumAddType) Ptr() *AlbumAddType {
	return &v
}

type NullableAlbumAddType struct {
	value *AlbumAddType
	isSet bool
}

func (v NullableAlbumAddType) Get() *AlbumAddType {
	return v.value
}

func (v *NullableAlbumAddType) Set(val *AlbumAddType) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumAddType) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumAddType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumAddType(val *AlbumAddType) *NullableAlbumAddType {
	return &NullableAlbumAddType{value: val, isSet: true}
}

func (v NullableAlbumAddType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumAddType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

