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

// TrackListLazyLoaded struct for TrackListLazyLoaded
type TrackListLazyLoaded struct {
	Value []*Track `json:"value,omitempty"`
	IsLoaded *bool `json:"isLoaded,omitempty"`
}

// NewTrackListLazyLoaded instantiates a new TrackListLazyLoaded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackListLazyLoaded() *TrackListLazyLoaded {
	this := TrackListLazyLoaded{}
	return &this
}

// NewTrackListLazyLoadedWithDefaults instantiates a new TrackListLazyLoaded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackListLazyLoadedWithDefaults() *TrackListLazyLoaded {
	this := TrackListLazyLoaded{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackListLazyLoaded) GetValue() []*Track {
	if o == nil {
		var ret []*Track
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackListLazyLoaded) GetValueOk() ([]*Track, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TrackListLazyLoaded) HasValue() bool {
	if o != nil && isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Track and assigns it to the Value field.
func (o *TrackListLazyLoaded) SetValue(v []*Track) {
	o.Value = v
}

// GetIsLoaded returns the IsLoaded field value if set, zero value otherwise.
func (o *TrackListLazyLoaded) GetIsLoaded() bool {
	if o == nil || isNil(o.IsLoaded) {
		var ret bool
		return ret
	}
	return *o.IsLoaded
}

// GetIsLoadedOk returns a tuple with the IsLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackListLazyLoaded) GetIsLoadedOk() (*bool, bool) {
	if o == nil || isNil(o.IsLoaded) {
    return nil, false
	}
	return o.IsLoaded, true
}

// HasIsLoaded returns a boolean if a field has been set.
func (o *TrackListLazyLoaded) HasIsLoaded() bool {
	if o != nil && !isNil(o.IsLoaded) {
		return true
	}

	return false
}

// SetIsLoaded gets a reference to the given bool and assigns it to the IsLoaded field.
func (o *TrackListLazyLoaded) SetIsLoaded(v bool) {
	o.IsLoaded = &v
}

func (o TrackListLazyLoaded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.IsLoaded) {
		toSerialize["isLoaded"] = o.IsLoaded
	}
	return json.Marshal(toSerialize)
}

type NullableTrackListLazyLoaded struct {
	value *TrackListLazyLoaded
	isSet bool
}

func (v NullableTrackListLazyLoaded) Get() *TrackListLazyLoaded {
	return v.value
}

func (v *NullableTrackListLazyLoaded) Set(val *TrackListLazyLoaded) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackListLazyLoaded) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackListLazyLoaded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackListLazyLoaded(val *TrackListLazyLoaded) *NullableTrackListLazyLoaded {
	return &NullableTrackListLazyLoaded{value: val, isSet: true}
}

func (v NullableTrackListLazyLoaded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackListLazyLoaded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


