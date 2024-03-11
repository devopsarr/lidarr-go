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

// checks if the MetadataProfileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataProfileResource{}

// MetadataProfileResource struct for MetadataProfileResource
type MetadataProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	PrimaryAlbumTypes []ProfilePrimaryAlbumTypeItemResource `json:"primaryAlbumTypes,omitempty"`
	SecondaryAlbumTypes []ProfileSecondaryAlbumTypeItemResource `json:"secondaryAlbumTypes,omitempty"`
	ReleaseStatuses []ProfileReleaseStatusItemResource `json:"releaseStatuses,omitempty"`
}

// NewMetadataProfileResource instantiates a new MetadataProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataProfileResource() *MetadataProfileResource {
	this := MetadataProfileResource{}
	return &this
}

// NewMetadataProfileResourceWithDefaults instantiates a new MetadataProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataProfileResourceWithDefaults() *MetadataProfileResource {
	this := MetadataProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataProfileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetadataProfileResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetadataProfileResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetadataProfileResource) UnsetName() {
	o.Name.Unset()
}

// GetPrimaryAlbumTypes returns the PrimaryAlbumTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetPrimaryAlbumTypes() []ProfilePrimaryAlbumTypeItemResource {
	if o == nil {
		var ret []ProfilePrimaryAlbumTypeItemResource
		return ret
	}
	return o.PrimaryAlbumTypes
}

// GetPrimaryAlbumTypesOk returns a tuple with the PrimaryAlbumTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetPrimaryAlbumTypesOk() ([]ProfilePrimaryAlbumTypeItemResource, bool) {
	if o == nil || IsNil(o.PrimaryAlbumTypes) {
		return nil, false
	}
	return o.PrimaryAlbumTypes, true
}

// HasPrimaryAlbumTypes returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasPrimaryAlbumTypes() bool {
	if o != nil && !IsNil(o.PrimaryAlbumTypes) {
		return true
	}

	return false
}

// SetPrimaryAlbumTypes gets a reference to the given []ProfilePrimaryAlbumTypeItemResource and assigns it to the PrimaryAlbumTypes field.
func (o *MetadataProfileResource) SetPrimaryAlbumTypes(v []ProfilePrimaryAlbumTypeItemResource) {
	o.PrimaryAlbumTypes = v
}

// GetSecondaryAlbumTypes returns the SecondaryAlbumTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetSecondaryAlbumTypes() []ProfileSecondaryAlbumTypeItemResource {
	if o == nil {
		var ret []ProfileSecondaryAlbumTypeItemResource
		return ret
	}
	return o.SecondaryAlbumTypes
}

// GetSecondaryAlbumTypesOk returns a tuple with the SecondaryAlbumTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetSecondaryAlbumTypesOk() ([]ProfileSecondaryAlbumTypeItemResource, bool) {
	if o == nil || IsNil(o.SecondaryAlbumTypes) {
		return nil, false
	}
	return o.SecondaryAlbumTypes, true
}

// HasSecondaryAlbumTypes returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasSecondaryAlbumTypes() bool {
	if o != nil && !IsNil(o.SecondaryAlbumTypes) {
		return true
	}

	return false
}

// SetSecondaryAlbumTypes gets a reference to the given []ProfileSecondaryAlbumTypeItemResource and assigns it to the SecondaryAlbumTypes field.
func (o *MetadataProfileResource) SetSecondaryAlbumTypes(v []ProfileSecondaryAlbumTypeItemResource) {
	o.SecondaryAlbumTypes = v
}

// GetReleaseStatuses returns the ReleaseStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProfileResource) GetReleaseStatuses() []ProfileReleaseStatusItemResource {
	if o == nil {
		var ret []ProfileReleaseStatusItemResource
		return ret
	}
	return o.ReleaseStatuses
}

// GetReleaseStatusesOk returns a tuple with the ReleaseStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProfileResource) GetReleaseStatusesOk() ([]ProfileReleaseStatusItemResource, bool) {
	if o == nil || IsNil(o.ReleaseStatuses) {
		return nil, false
	}
	return o.ReleaseStatuses, true
}

// HasReleaseStatuses returns a boolean if a field has been set.
func (o *MetadataProfileResource) HasReleaseStatuses() bool {
	if o != nil && !IsNil(o.ReleaseStatuses) {
		return true
	}

	return false
}

// SetReleaseStatuses gets a reference to the given []ProfileReleaseStatusItemResource and assigns it to the ReleaseStatuses field.
func (o *MetadataProfileResource) SetReleaseStatuses(v []ProfileReleaseStatusItemResource) {
	o.ReleaseStatuses = v
}

func (o MetadataProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataProfileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PrimaryAlbumTypes != nil {
		toSerialize["primaryAlbumTypes"] = o.PrimaryAlbumTypes
	}
	if o.SecondaryAlbumTypes != nil {
		toSerialize["secondaryAlbumTypes"] = o.SecondaryAlbumTypes
	}
	if o.ReleaseStatuses != nil {
		toSerialize["releaseStatuses"] = o.ReleaseStatuses
	}
	return toSerialize, nil
}

type NullableMetadataProfileResource struct {
	value *MetadataProfileResource
	isSet bool
}

func (v NullableMetadataProfileResource) Get() *MetadataProfileResource {
	return v.value
}

func (v *NullableMetadataProfileResource) Set(val *MetadataProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProfileResource(val *MetadataProfileResource) *NullableMetadataProfileResource {
	return &NullableMetadataProfileResource{value: val, isSet: true}
}

func (v NullableMetadataProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


