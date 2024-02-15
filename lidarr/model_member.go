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

// checks if the Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Member{}

// Member struct for Member
type Member struct {
	Name NullableString `json:"name,omitempty"`
	Instrument NullableString `json:"instrument,omitempty"`
	Images []MediaCover `json:"images,omitempty"`
}

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember() *Member {
	this := Member{}
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Member) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Member) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Member) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Member) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Member) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Member) UnsetName() {
	o.Name.Unset()
}

// GetInstrument returns the Instrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Member) GetInstrument() string {
	if o == nil || IsNil(o.Instrument.Get()) {
		var ret string
		return ret
	}
	return *o.Instrument.Get()
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Member) GetInstrumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instrument.Get(), o.Instrument.IsSet()
}

// HasInstrument returns a boolean if a field has been set.
func (o *Member) HasInstrument() bool {
	if o != nil && o.Instrument.IsSet() {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given NullableString and assigns it to the Instrument field.
func (o *Member) SetInstrument(v string) {
	o.Instrument.Set(&v)
}
// SetInstrumentNil sets the value for Instrument to be an explicit nil
func (o *Member) SetInstrumentNil() {
	o.Instrument.Set(nil)
}

// UnsetInstrument ensures that no value is present for Instrument, not even an explicit nil
func (o *Member) UnsetInstrument() {
	o.Instrument.Unset()
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Member) GetImages() []MediaCover {
	if o == nil {
		var ret []MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Member) GetImagesOk() ([]MediaCover, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Member) HasImages() bool {
	if o != nil && IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *Member) SetImages(v []MediaCover) {
	o.Images = v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Instrument.IsSet() {
		toSerialize["instrument"] = o.Instrument.Get()
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


