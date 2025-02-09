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

// CommandStatus the model 'CommandStatus'
type CommandStatus string

// List of CommandStatus
const (
	COMMANDSTATUS_QUEUED CommandStatus = "queued"
	COMMANDSTATUS_STARTED CommandStatus = "started"
	COMMANDSTATUS_COMPLETED CommandStatus = "completed"
	COMMANDSTATUS_FAILED CommandStatus = "failed"
	COMMANDSTATUS_ABORTED CommandStatus = "aborted"
	COMMANDSTATUS_CANCELLED CommandStatus = "cancelled"
	COMMANDSTATUS_ORPHANED CommandStatus = "orphaned"
)

// All allowed values of CommandStatus enum
var AllowedCommandStatusEnumValues = []CommandStatus{
	"queued",
	"started",
	"completed",
	"failed",
	"aborted",
	"cancelled",
	"orphaned",
}

func (v *CommandStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommandStatus(value)
	for _, existing := range AllowedCommandStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommandStatus", value)
}

// NewCommandStatusFromValue returns a pointer to a valid CommandStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommandStatusFromValue(v string) (*CommandStatus, error) {
	ev := CommandStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommandStatus: valid values are %v", v, AllowedCommandStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommandStatus) IsValid() bool {
	for _, existing := range AllowedCommandStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommandStatus value
func (v CommandStatus) Ptr() *CommandStatus {
	return &v
}

type NullableCommandStatus struct {
	value *CommandStatus
	isSet bool
}

func (v NullableCommandStatus) Get() *CommandStatus {
	return v.value
}

func (v *NullableCommandStatus) Set(val *CommandStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandStatus(val *CommandStatus) *NullableCommandStatus {
	return &NullableCommandStatus{value: val, isSet: true}
}

func (v NullableCommandStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

