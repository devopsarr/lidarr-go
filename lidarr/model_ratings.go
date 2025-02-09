/*
Lidarr

Lidarr API docs

API version: v2.9.6.4552
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the Ratings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ratings{}

// Ratings struct for Ratings
type Ratings struct {
	Votes *int32 `json:"votes,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

// NewRatings instantiates a new Ratings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatings() *Ratings {
	this := Ratings{}
	return &this
}

// NewRatingsWithDefaults instantiates a new Ratings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingsWithDefaults() *Ratings {
	this := Ratings{}
	return &this
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *Ratings) GetVotes() int32 {
	if o == nil || IsNil(o.Votes) {
		var ret int32
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetVotesOk() (*int32, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *Ratings) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given int32 and assigns it to the Votes field.
func (o *Ratings) SetVotes(v int32) {
	o.Votes = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Ratings) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Ratings) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *Ratings) SetValue(v float64) {
	o.Value = &v
}

func (o Ratings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ratings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRatings struct {
	value *Ratings
	isSet bool
}

func (v NullableRatings) Get() *Ratings {
	return v.value
}

func (v *NullableRatings) Set(val *Ratings) {
	v.value = val
	v.isSet = true
}

func (v NullableRatings) IsSet() bool {
	return v.isSet
}

func (v *NullableRatings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatings(val *Ratings) *NullableRatings {
	return &NullableRatings{value: val, isSet: true}
}

func (v NullableRatings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


