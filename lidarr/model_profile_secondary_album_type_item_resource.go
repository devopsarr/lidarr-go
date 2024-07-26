/*
Lidarr

Lidarr API docs

API version: v2.4.3.4248
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the ProfileSecondaryAlbumTypeItemResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileSecondaryAlbumTypeItemResource{}

// ProfileSecondaryAlbumTypeItemResource struct for ProfileSecondaryAlbumTypeItemResource
type ProfileSecondaryAlbumTypeItemResource struct {
	Id *int32 `json:"id,omitempty"`
	AlbumType *SecondaryAlbumType `json:"albumType,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewProfileSecondaryAlbumTypeItemResource instantiates a new ProfileSecondaryAlbumTypeItemResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSecondaryAlbumTypeItemResource() *ProfileSecondaryAlbumTypeItemResource {
	this := ProfileSecondaryAlbumTypeItemResource{}
	return &this
}

// NewProfileSecondaryAlbumTypeItemResourceWithDefaults instantiates a new ProfileSecondaryAlbumTypeItemResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSecondaryAlbumTypeItemResourceWithDefaults() *ProfileSecondaryAlbumTypeItemResource {
	this := ProfileSecondaryAlbumTypeItemResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProfileSecondaryAlbumTypeItemResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSecondaryAlbumTypeItemResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfileSecondaryAlbumTypeItemResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProfileSecondaryAlbumTypeItemResource) SetId(v int32) {
	o.Id = &v
}

// GetAlbumType returns the AlbumType field value if set, zero value otherwise.
func (o *ProfileSecondaryAlbumTypeItemResource) GetAlbumType() SecondaryAlbumType {
	if o == nil || IsNil(o.AlbumType) {
		var ret SecondaryAlbumType
		return ret
	}
	return *o.AlbumType
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSecondaryAlbumTypeItemResource) GetAlbumTypeOk() (*SecondaryAlbumType, bool) {
	if o == nil || IsNil(o.AlbumType) {
		return nil, false
	}
	return o.AlbumType, true
}

// HasAlbumType returns a boolean if a field has been set.
func (o *ProfileSecondaryAlbumTypeItemResource) HasAlbumType() bool {
	if o != nil && !IsNil(o.AlbumType) {
		return true
	}

	return false
}

// SetAlbumType gets a reference to the given SecondaryAlbumType and assigns it to the AlbumType field.
func (o *ProfileSecondaryAlbumTypeItemResource) SetAlbumType(v SecondaryAlbumType) {
	o.AlbumType = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *ProfileSecondaryAlbumTypeItemResource) GetAllowed() bool {
	if o == nil || IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSecondaryAlbumTypeItemResource) GetAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *ProfileSecondaryAlbumTypeItemResource) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *ProfileSecondaryAlbumTypeItemResource) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o ProfileSecondaryAlbumTypeItemResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileSecondaryAlbumTypeItemResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AlbumType) {
		toSerialize["albumType"] = o.AlbumType
	}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableProfileSecondaryAlbumTypeItemResource struct {
	value *ProfileSecondaryAlbumTypeItemResource
	isSet bool
}

func (v NullableProfileSecondaryAlbumTypeItemResource) Get() *ProfileSecondaryAlbumTypeItemResource {
	return v.value
}

func (v *NullableProfileSecondaryAlbumTypeItemResource) Set(val *ProfileSecondaryAlbumTypeItemResource) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSecondaryAlbumTypeItemResource) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSecondaryAlbumTypeItemResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSecondaryAlbumTypeItemResource(val *ProfileSecondaryAlbumTypeItemResource) *NullableProfileSecondaryAlbumTypeItemResource {
	return &NullableProfileSecondaryAlbumTypeItemResource{value: val, isSet: true}
}

func (v NullableProfileSecondaryAlbumTypeItemResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSecondaryAlbumTypeItemResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


