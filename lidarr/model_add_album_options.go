/*
Lidarr

Lidarr API docs

API version: v2.1.7.4030
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// AddAlbumOptions struct for AddAlbumOptions
type AddAlbumOptions struct {
	AddType *AlbumAddType `json:"addType,omitempty"`
	SearchForNewAlbum *bool `json:"searchForNewAlbum,omitempty"`
}

// NewAddAlbumOptions instantiates a new AddAlbumOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAlbumOptions() *AddAlbumOptions {
	this := AddAlbumOptions{}
	return &this
}

// NewAddAlbumOptionsWithDefaults instantiates a new AddAlbumOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAlbumOptionsWithDefaults() *AddAlbumOptions {
	this := AddAlbumOptions{}
	return &this
}

// GetAddType returns the AddType field value if set, zero value otherwise.
func (o *AddAlbumOptions) GetAddType() AlbumAddType {
	if o == nil || isNil(o.AddType) {
		var ret AlbumAddType
		return ret
	}
	return *o.AddType
}

// GetAddTypeOk returns a tuple with the AddType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAlbumOptions) GetAddTypeOk() (*AlbumAddType, bool) {
	if o == nil || isNil(o.AddType) {
    return nil, false
	}
	return o.AddType, true
}

// HasAddType returns a boolean if a field has been set.
func (o *AddAlbumOptions) HasAddType() bool {
	if o != nil && !isNil(o.AddType) {
		return true
	}

	return false
}

// SetAddType gets a reference to the given AlbumAddType and assigns it to the AddType field.
func (o *AddAlbumOptions) SetAddType(v AlbumAddType) {
	o.AddType = &v
}

// GetSearchForNewAlbum returns the SearchForNewAlbum field value if set, zero value otherwise.
func (o *AddAlbumOptions) GetSearchForNewAlbum() bool {
	if o == nil || isNil(o.SearchForNewAlbum) {
		var ret bool
		return ret
	}
	return *o.SearchForNewAlbum
}

// GetSearchForNewAlbumOk returns a tuple with the SearchForNewAlbum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAlbumOptions) GetSearchForNewAlbumOk() (*bool, bool) {
	if o == nil || isNil(o.SearchForNewAlbum) {
    return nil, false
	}
	return o.SearchForNewAlbum, true
}

// HasSearchForNewAlbum returns a boolean if a field has been set.
func (o *AddAlbumOptions) HasSearchForNewAlbum() bool {
	if o != nil && !isNil(o.SearchForNewAlbum) {
		return true
	}

	return false
}

// SetSearchForNewAlbum gets a reference to the given bool and assigns it to the SearchForNewAlbum field.
func (o *AddAlbumOptions) SetSearchForNewAlbum(v bool) {
	o.SearchForNewAlbum = &v
}

func (o AddAlbumOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AddType) {
		toSerialize["addType"] = o.AddType
	}
	if !isNil(o.SearchForNewAlbum) {
		toSerialize["searchForNewAlbum"] = o.SearchForNewAlbum
	}
	return json.Marshal(toSerialize)
}

type NullableAddAlbumOptions struct {
	value *AddAlbumOptions
	isSet bool
}

func (v NullableAddAlbumOptions) Get() *AddAlbumOptions {
	return v.value
}

func (v *NullableAddAlbumOptions) Set(val *AddAlbumOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAlbumOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAlbumOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAlbumOptions(val *AddAlbumOptions) *NullableAddAlbumOptions {
	return &NullableAddAlbumOptions{value: val, isSet: true}
}

func (v NullableAddAlbumOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAlbumOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


