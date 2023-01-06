/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"fmt"
)

// ProxyType the model 'ProxyType'
type ProxyType string

// List of ProxyType
const (
	PROXYTYPE_HTTP ProxyType = "http"
	PROXYTYPE_SOCKS4 ProxyType = "socks4"
	PROXYTYPE_SOCKS5 ProxyType = "socks5"
)

// All allowed values of ProxyType enum
var AllowedProxyTypeEnumValues = []ProxyType{
	"http",
	"socks4",
	"socks5",
}

func (v *ProxyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProxyType(value)
	for _, existing := range AllowedProxyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProxyType", value)
}

// NewProxyTypeFromValue returns a pointer to a valid ProxyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxyTypeFromValue(v string) (*ProxyType, error) {
	ev := ProxyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProxyType: valid values are %v", v, AllowedProxyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxyType) IsValid() bool {
	for _, existing := range AllowedProxyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxyType value
func (v ProxyType) Ptr() *ProxyType {
	return &v
}

type NullableProxyType struct {
	value *ProxyType
	isSet bool
}

func (v NullableProxyType) Get() *ProxyType {
	return v.value
}

func (v *NullableProxyType) Set(val *ProxyType) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyType) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyType(val *ProxyType) *NullableProxyType {
	return &NullableProxyType{value: val, isSet: true}
}

func (v NullableProxyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

