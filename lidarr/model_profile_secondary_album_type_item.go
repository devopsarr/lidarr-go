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

// ProfileSecondaryAlbumTypeItem struct for ProfileSecondaryAlbumTypeItem
type ProfileSecondaryAlbumTypeItem struct {
	SecondaryAlbumType *SecondaryAlbumType `json:"secondaryAlbumType,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewProfileSecondaryAlbumTypeItem instantiates a new ProfileSecondaryAlbumTypeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSecondaryAlbumTypeItem() *ProfileSecondaryAlbumTypeItem {
	this := ProfileSecondaryAlbumTypeItem{}
	return &this
}

// NewProfileSecondaryAlbumTypeItemWithDefaults instantiates a new ProfileSecondaryAlbumTypeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSecondaryAlbumTypeItemWithDefaults() *ProfileSecondaryAlbumTypeItem {
	this := ProfileSecondaryAlbumTypeItem{}
	return &this
}

// GetSecondaryAlbumType returns the SecondaryAlbumType field value if set, zero value otherwise.
func (o *ProfileSecondaryAlbumTypeItem) GetSecondaryAlbumType() SecondaryAlbumType {
	if o == nil || isNil(o.SecondaryAlbumType) {
		var ret SecondaryAlbumType
		return ret
	}
	return *o.SecondaryAlbumType
}

// GetSecondaryAlbumTypeOk returns a tuple with the SecondaryAlbumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSecondaryAlbumTypeItem) GetSecondaryAlbumTypeOk() (*SecondaryAlbumType, bool) {
	if o == nil || isNil(o.SecondaryAlbumType) {
    return nil, false
	}
	return o.SecondaryAlbumType, true
}

// HasSecondaryAlbumType returns a boolean if a field has been set.
func (o *ProfileSecondaryAlbumTypeItem) HasSecondaryAlbumType() bool {
	if o != nil && !isNil(o.SecondaryAlbumType) {
		return true
	}

	return false
}

// SetSecondaryAlbumType gets a reference to the given SecondaryAlbumType and assigns it to the SecondaryAlbumType field.
func (o *ProfileSecondaryAlbumTypeItem) SetSecondaryAlbumType(v SecondaryAlbumType) {
	o.SecondaryAlbumType = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *ProfileSecondaryAlbumTypeItem) GetAllowed() bool {
	if o == nil || isNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSecondaryAlbumTypeItem) GetAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.Allowed) {
    return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *ProfileSecondaryAlbumTypeItem) HasAllowed() bool {
	if o != nil && !isNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *ProfileSecondaryAlbumTypeItem) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o ProfileSecondaryAlbumTypeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SecondaryAlbumType) {
		toSerialize["secondaryAlbumType"] = o.SecondaryAlbumType
	}
	if !isNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return json.Marshal(toSerialize)
}

type NullableProfileSecondaryAlbumTypeItem struct {
	value *ProfileSecondaryAlbumTypeItem
	isSet bool
}

func (v NullableProfileSecondaryAlbumTypeItem) Get() *ProfileSecondaryAlbumTypeItem {
	return v.value
}

func (v *NullableProfileSecondaryAlbumTypeItem) Set(val *ProfileSecondaryAlbumTypeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSecondaryAlbumTypeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSecondaryAlbumTypeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSecondaryAlbumTypeItem(val *ProfileSecondaryAlbumTypeItem) *NullableProfileSecondaryAlbumTypeItem {
	return &NullableProfileSecondaryAlbumTypeItem{value: val, isSet: true}
}

func (v NullableProfileSecondaryAlbumTypeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSecondaryAlbumTypeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


