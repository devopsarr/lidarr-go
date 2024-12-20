/*
Lidarr

Lidarr API docs

API version: v2.8.2.4493
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the ProfilePrimaryAlbumTypeItemResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfilePrimaryAlbumTypeItemResource{}

// ProfilePrimaryAlbumTypeItemResource struct for ProfilePrimaryAlbumTypeItemResource
type ProfilePrimaryAlbumTypeItemResource struct {
	Id *int32 `json:"id,omitempty"`
	AlbumType *PrimaryAlbumType `json:"albumType,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewProfilePrimaryAlbumTypeItemResource instantiates a new ProfilePrimaryAlbumTypeItemResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfilePrimaryAlbumTypeItemResource() *ProfilePrimaryAlbumTypeItemResource {
	this := ProfilePrimaryAlbumTypeItemResource{}
	return &this
}

// NewProfilePrimaryAlbumTypeItemResourceWithDefaults instantiates a new ProfilePrimaryAlbumTypeItemResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfilePrimaryAlbumTypeItemResourceWithDefaults() *ProfilePrimaryAlbumTypeItemResource {
	this := ProfilePrimaryAlbumTypeItemResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProfilePrimaryAlbumTypeItemResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePrimaryAlbumTypeItemResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfilePrimaryAlbumTypeItemResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProfilePrimaryAlbumTypeItemResource) SetId(v int32) {
	o.Id = &v
}

// GetAlbumType returns the AlbumType field value if set, zero value otherwise.
func (o *ProfilePrimaryAlbumTypeItemResource) GetAlbumType() PrimaryAlbumType {
	if o == nil || IsNil(o.AlbumType) {
		var ret PrimaryAlbumType
		return ret
	}
	return *o.AlbumType
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePrimaryAlbumTypeItemResource) GetAlbumTypeOk() (*PrimaryAlbumType, bool) {
	if o == nil || IsNil(o.AlbumType) {
		return nil, false
	}
	return o.AlbumType, true
}

// HasAlbumType returns a boolean if a field has been set.
func (o *ProfilePrimaryAlbumTypeItemResource) HasAlbumType() bool {
	if o != nil && !IsNil(o.AlbumType) {
		return true
	}

	return false
}

// SetAlbumType gets a reference to the given PrimaryAlbumType and assigns it to the AlbumType field.
func (o *ProfilePrimaryAlbumTypeItemResource) SetAlbumType(v PrimaryAlbumType) {
	o.AlbumType = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *ProfilePrimaryAlbumTypeItemResource) GetAllowed() bool {
	if o == nil || IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePrimaryAlbumTypeItemResource) GetAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *ProfilePrimaryAlbumTypeItemResource) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *ProfilePrimaryAlbumTypeItemResource) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o ProfilePrimaryAlbumTypeItemResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfilePrimaryAlbumTypeItemResource) ToMap() (map[string]interface{}, error) {
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

type NullableProfilePrimaryAlbumTypeItemResource struct {
	value *ProfilePrimaryAlbumTypeItemResource
	isSet bool
}

func (v NullableProfilePrimaryAlbumTypeItemResource) Get() *ProfilePrimaryAlbumTypeItemResource {
	return v.value
}

func (v *NullableProfilePrimaryAlbumTypeItemResource) Set(val *ProfilePrimaryAlbumTypeItemResource) {
	v.value = val
	v.isSet = true
}

func (v NullableProfilePrimaryAlbumTypeItemResource) IsSet() bool {
	return v.isSet
}

func (v *NullableProfilePrimaryAlbumTypeItemResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfilePrimaryAlbumTypeItemResource(val *ProfilePrimaryAlbumTypeItemResource) *NullableProfilePrimaryAlbumTypeItemResource {
	return &NullableProfilePrimaryAlbumTypeItemResource{value: val, isSet: true}
}

func (v NullableProfilePrimaryAlbumTypeItemResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfilePrimaryAlbumTypeItemResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


