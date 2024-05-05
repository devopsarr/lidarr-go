/*
Lidarr

Lidarr API docs

API version: v2.2.5.4141
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"fmt"
)

// UpdateMechanism the model 'UpdateMechanism'
type UpdateMechanism string

// List of UpdateMechanism
const (
	UPDATEMECHANISM_BUILT_IN UpdateMechanism = "builtIn"
	UPDATEMECHANISM_SCRIPT UpdateMechanism = "script"
	UPDATEMECHANISM_EXTERNAL UpdateMechanism = "external"
	UPDATEMECHANISM_APT UpdateMechanism = "apt"
	UPDATEMECHANISM_DOCKER UpdateMechanism = "docker"
)

// All allowed values of UpdateMechanism enum
var AllowedUpdateMechanismEnumValues = []UpdateMechanism{
	"builtIn",
	"script",
	"external",
	"apt",
	"docker",
}

func (v *UpdateMechanism) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateMechanism(value)
	for _, existing := range AllowedUpdateMechanismEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateMechanism", value)
}

// NewUpdateMechanismFromValue returns a pointer to a valid UpdateMechanism
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateMechanismFromValue(v string) (*UpdateMechanism, error) {
	ev := UpdateMechanism(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateMechanism: valid values are %v", v, AllowedUpdateMechanismEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateMechanism) IsValid() bool {
	for _, existing := range AllowedUpdateMechanismEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateMechanism value
func (v UpdateMechanism) Ptr() *UpdateMechanism {
	return &v
}

type NullableUpdateMechanism struct {
	value *UpdateMechanism
	isSet bool
}

func (v NullableUpdateMechanism) Get() *UpdateMechanism {
	return v.value
}

func (v *NullableUpdateMechanism) Set(val *UpdateMechanism) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMechanism) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMechanism) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMechanism(val *UpdateMechanism) *NullableUpdateMechanism {
	return &NullableUpdateMechanism{value: val, isSet: true}
}

func (v NullableUpdateMechanism) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMechanism) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

