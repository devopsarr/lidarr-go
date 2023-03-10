/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// Medium struct for Medium
type Medium struct {
	Number *int32 `json:"number,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Format NullableString `json:"format,omitempty"`
}

// NewMedium instantiates a new Medium object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedium() *Medium {
	this := Medium{}
	return &this
}

// NewMediumWithDefaults instantiates a new Medium object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediumWithDefaults() *Medium {
	this := Medium{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Medium) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Medium) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Medium) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *Medium) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Medium) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Medium) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Medium) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Medium) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Medium) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Medium) UnsetName() {
	o.Name.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Medium) GetFormat() string {
	if o == nil || isNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Medium) GetFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *Medium) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *Medium) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *Medium) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *Medium) UnsetFormat() {
	o.Format.Unset()
}

func (o Medium) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Format.IsSet() {
		toSerialize["format"] = o.Format.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMedium struct {
	value *Medium
	isSet bool
}

func (v NullableMedium) Get() *Medium {
	return v.value
}

func (v *NullableMedium) Set(val *Medium) {
	v.value = val
	v.isSet = true
}

func (v NullableMedium) IsSet() bool {
	return v.isSet
}

func (v *NullableMedium) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedium(val *Medium) *NullableMedium {
	return &NullableMedium{value: val, isSet: true}
}

func (v NullableMedium) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedium) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


