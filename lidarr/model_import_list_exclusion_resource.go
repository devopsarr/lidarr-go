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

// ImportListExclusionResource struct for ImportListExclusionResource
type ImportListExclusionResource struct {
	Id *int32 `json:"id,omitempty"`
	ForeignId NullableString `json:"foreignId,omitempty"`
	ArtistName NullableString `json:"artistName,omitempty"`
}

// NewImportListExclusionResource instantiates a new ImportListExclusionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListExclusionResource() *ImportListExclusionResource {
	this := ImportListExclusionResource{}
	return &this
}

// NewImportListExclusionResourceWithDefaults instantiates a new ImportListExclusionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListExclusionResourceWithDefaults() *ImportListExclusionResource {
	this := ImportListExclusionResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportListExclusionResource) SetId(v int32) {
	o.Id = &v
}

// GetForeignId returns the ForeignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetForeignId() string {
	if o == nil || isNil(o.ForeignId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignId.Get()
}

// GetForeignIdOk returns a tuple with the ForeignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetForeignIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignId.Get(), o.ForeignId.IsSet()
}

// HasForeignId returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasForeignId() bool {
	if o != nil && o.ForeignId.IsSet() {
		return true
	}

	return false
}

// SetForeignId gets a reference to the given NullableString and assigns it to the ForeignId field.
func (o *ImportListExclusionResource) SetForeignId(v string) {
	o.ForeignId.Set(&v)
}
// SetForeignIdNil sets the value for ForeignId to be an explicit nil
func (o *ImportListExclusionResource) SetForeignIdNil() {
	o.ForeignId.Set(nil)
}

// UnsetForeignId ensures that no value is present for ForeignId, not even an explicit nil
func (o *ImportListExclusionResource) UnsetForeignId() {
	o.ForeignId.Unset()
}

// GetArtistName returns the ArtistName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetArtistName() string {
	if o == nil || isNil(o.ArtistName.Get()) {
		var ret string
		return ret
	}
	return *o.ArtistName.Get()
}

// GetArtistNameOk returns a tuple with the ArtistName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetArtistNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArtistName.Get(), o.ArtistName.IsSet()
}

// HasArtistName returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasArtistName() bool {
	if o != nil && o.ArtistName.IsSet() {
		return true
	}

	return false
}

// SetArtistName gets a reference to the given NullableString and assigns it to the ArtistName field.
func (o *ImportListExclusionResource) SetArtistName(v string) {
	o.ArtistName.Set(&v)
}
// SetArtistNameNil sets the value for ArtistName to be an explicit nil
func (o *ImportListExclusionResource) SetArtistNameNil() {
	o.ArtistName.Set(nil)
}

// UnsetArtistName ensures that no value is present for ArtistName, not even an explicit nil
func (o *ImportListExclusionResource) UnsetArtistName() {
	o.ArtistName.Unset()
}

func (o ImportListExclusionResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ForeignId.IsSet() {
		toSerialize["foreignId"] = o.ForeignId.Get()
	}
	if o.ArtistName.IsSet() {
		toSerialize["artistName"] = o.ArtistName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableImportListExclusionResource struct {
	value *ImportListExclusionResource
	isSet bool
}

func (v NullableImportListExclusionResource) Get() *ImportListExclusionResource {
	return v.value
}

func (v *NullableImportListExclusionResource) Set(val *ImportListExclusionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListExclusionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListExclusionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListExclusionResource(val *ImportListExclusionResource) *NullableImportListExclusionResource {
	return &NullableImportListExclusionResource{value: val, isSet: true}
}

func (v NullableImportListExclusionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListExclusionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


