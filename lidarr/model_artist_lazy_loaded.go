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

// ArtistLazyLoaded struct for ArtistLazyLoaded
type ArtistLazyLoaded struct {
	Value *Artist `json:"value,omitempty"`
	IsLoaded *bool `json:"isLoaded,omitempty"`
}

// NewArtistLazyLoaded instantiates a new ArtistLazyLoaded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtistLazyLoaded() *ArtistLazyLoaded {
	this := ArtistLazyLoaded{}
	return &this
}

// NewArtistLazyLoadedWithDefaults instantiates a new ArtistLazyLoaded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtistLazyLoadedWithDefaults() *ArtistLazyLoaded {
	this := ArtistLazyLoaded{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ArtistLazyLoaded) GetValue() Artist {
	if o == nil || isNil(o.Value) {
		var ret Artist
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistLazyLoaded) GetValueOk() (*Artist, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ArtistLazyLoaded) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Artist and assigns it to the Value field.
func (o *ArtistLazyLoaded) SetValue(v Artist) {
	o.Value = &v
}

// GetIsLoaded returns the IsLoaded field value if set, zero value otherwise.
func (o *ArtistLazyLoaded) GetIsLoaded() bool {
	if o == nil || isNil(o.IsLoaded) {
		var ret bool
		return ret
	}
	return *o.IsLoaded
}

// GetIsLoadedOk returns a tuple with the IsLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistLazyLoaded) GetIsLoadedOk() (*bool, bool) {
	if o == nil || isNil(o.IsLoaded) {
    return nil, false
	}
	return o.IsLoaded, true
}

// HasIsLoaded returns a boolean if a field has been set.
func (o *ArtistLazyLoaded) HasIsLoaded() bool {
	if o != nil && !isNil(o.IsLoaded) {
		return true
	}

	return false
}

// SetIsLoaded gets a reference to the given bool and assigns it to the IsLoaded field.
func (o *ArtistLazyLoaded) SetIsLoaded(v bool) {
	o.IsLoaded = &v
}

func (o ArtistLazyLoaded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.IsLoaded) {
		toSerialize["isLoaded"] = o.IsLoaded
	}
	return json.Marshal(toSerialize)
}

type NullableArtistLazyLoaded struct {
	value *ArtistLazyLoaded
	isSet bool
}

func (v NullableArtistLazyLoaded) Get() *ArtistLazyLoaded {
	return v.value
}

func (v *NullableArtistLazyLoaded) Set(val *ArtistLazyLoaded) {
	v.value = val
	v.isSet = true
}

func (v NullableArtistLazyLoaded) IsSet() bool {
	return v.isSet
}

func (v *NullableArtistLazyLoaded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtistLazyLoaded(val *ArtistLazyLoaded) *NullableArtistLazyLoaded {
	return &NullableArtistLazyLoaded{value: val, isSet: true}
}

func (v NullableArtistLazyLoaded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtistLazyLoaded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


