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

// checks if the RenameTrackResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenameTrackResource{}

// RenameTrackResource struct for RenameTrackResource
type RenameTrackResource struct {
	Id *int32 `json:"id,omitempty"`
	ArtistId *int32 `json:"artistId,omitempty"`
	AlbumId *int32 `json:"albumId,omitempty"`
	TrackNumbers []int32 `json:"trackNumbers,omitempty"`
	TrackFileId *int32 `json:"trackFileId,omitempty"`
	ExistingPath NullableString `json:"existingPath,omitempty"`
	NewPath NullableString `json:"newPath,omitempty"`
}

// NewRenameTrackResource instantiates a new RenameTrackResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameTrackResource() *RenameTrackResource {
	this := RenameTrackResource{}
	return &this
}

// NewRenameTrackResourceWithDefaults instantiates a new RenameTrackResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameTrackResourceWithDefaults() *RenameTrackResource {
	this := RenameTrackResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RenameTrackResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameTrackResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RenameTrackResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RenameTrackResource) SetId(v int32) {
	o.Id = &v
}

// GetArtistId returns the ArtistId field value if set, zero value otherwise.
func (o *RenameTrackResource) GetArtistId() int32 {
	if o == nil || IsNil(o.ArtistId) {
		var ret int32
		return ret
	}
	return *o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameTrackResource) GetArtistIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ArtistId) {
		return nil, false
	}
	return o.ArtistId, true
}

// HasArtistId returns a boolean if a field has been set.
func (o *RenameTrackResource) HasArtistId() bool {
	if o != nil && !IsNil(o.ArtistId) {
		return true
	}

	return false
}

// SetArtistId gets a reference to the given int32 and assigns it to the ArtistId field.
func (o *RenameTrackResource) SetArtistId(v int32) {
	o.ArtistId = &v
}

// GetAlbumId returns the AlbumId field value if set, zero value otherwise.
func (o *RenameTrackResource) GetAlbumId() int32 {
	if o == nil || IsNil(o.AlbumId) {
		var ret int32
		return ret
	}
	return *o.AlbumId
}

// GetAlbumIdOk returns a tuple with the AlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameTrackResource) GetAlbumIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AlbumId) {
		return nil, false
	}
	return o.AlbumId, true
}

// HasAlbumId returns a boolean if a field has been set.
func (o *RenameTrackResource) HasAlbumId() bool {
	if o != nil && !IsNil(o.AlbumId) {
		return true
	}

	return false
}

// SetAlbumId gets a reference to the given int32 and assigns it to the AlbumId field.
func (o *RenameTrackResource) SetAlbumId(v int32) {
	o.AlbumId = &v
}

// GetTrackNumbers returns the TrackNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameTrackResource) GetTrackNumbers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.TrackNumbers
}

// GetTrackNumbersOk returns a tuple with the TrackNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameTrackResource) GetTrackNumbersOk() ([]int32, bool) {
	if o == nil || IsNil(o.TrackNumbers) {
		return nil, false
	}
	return o.TrackNumbers, true
}

// HasTrackNumbers returns a boolean if a field has been set.
func (o *RenameTrackResource) HasTrackNumbers() bool {
	if o != nil && !IsNil(o.TrackNumbers) {
		return true
	}

	return false
}

// SetTrackNumbers gets a reference to the given []int32 and assigns it to the TrackNumbers field.
func (o *RenameTrackResource) SetTrackNumbers(v []int32) {
	o.TrackNumbers = v
}

// GetTrackFileId returns the TrackFileId field value if set, zero value otherwise.
func (o *RenameTrackResource) GetTrackFileId() int32 {
	if o == nil || IsNil(o.TrackFileId) {
		var ret int32
		return ret
	}
	return *o.TrackFileId
}

// GetTrackFileIdOk returns a tuple with the TrackFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameTrackResource) GetTrackFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TrackFileId) {
		return nil, false
	}
	return o.TrackFileId, true
}

// HasTrackFileId returns a boolean if a field has been set.
func (o *RenameTrackResource) HasTrackFileId() bool {
	if o != nil && !IsNil(o.TrackFileId) {
		return true
	}

	return false
}

// SetTrackFileId gets a reference to the given int32 and assigns it to the TrackFileId field.
func (o *RenameTrackResource) SetTrackFileId(v int32) {
	o.TrackFileId = &v
}

// GetExistingPath returns the ExistingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameTrackResource) GetExistingPath() string {
	if o == nil || IsNil(o.ExistingPath.Get()) {
		var ret string
		return ret
	}
	return *o.ExistingPath.Get()
}

// GetExistingPathOk returns a tuple with the ExistingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameTrackResource) GetExistingPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExistingPath.Get(), o.ExistingPath.IsSet()
}

// HasExistingPath returns a boolean if a field has been set.
func (o *RenameTrackResource) HasExistingPath() bool {
	if o != nil && o.ExistingPath.IsSet() {
		return true
	}

	return false
}

// SetExistingPath gets a reference to the given NullableString and assigns it to the ExistingPath field.
func (o *RenameTrackResource) SetExistingPath(v string) {
	o.ExistingPath.Set(&v)
}
// SetExistingPathNil sets the value for ExistingPath to be an explicit nil
func (o *RenameTrackResource) SetExistingPathNil() {
	o.ExistingPath.Set(nil)
}

// UnsetExistingPath ensures that no value is present for ExistingPath, not even an explicit nil
func (o *RenameTrackResource) UnsetExistingPath() {
	o.ExistingPath.Unset()
}

// GetNewPath returns the NewPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameTrackResource) GetNewPath() string {
	if o == nil || IsNil(o.NewPath.Get()) {
		var ret string
		return ret
	}
	return *o.NewPath.Get()
}

// GetNewPathOk returns a tuple with the NewPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameTrackResource) GetNewPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewPath.Get(), o.NewPath.IsSet()
}

// HasNewPath returns a boolean if a field has been set.
func (o *RenameTrackResource) HasNewPath() bool {
	if o != nil && o.NewPath.IsSet() {
		return true
	}

	return false
}

// SetNewPath gets a reference to the given NullableString and assigns it to the NewPath field.
func (o *RenameTrackResource) SetNewPath(v string) {
	o.NewPath.Set(&v)
}
// SetNewPathNil sets the value for NewPath to be an explicit nil
func (o *RenameTrackResource) SetNewPathNil() {
	o.NewPath.Set(nil)
}

// UnsetNewPath ensures that no value is present for NewPath, not even an explicit nil
func (o *RenameTrackResource) UnsetNewPath() {
	o.NewPath.Unset()
}

func (o RenameTrackResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenameTrackResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ArtistId) {
		toSerialize["artistId"] = o.ArtistId
	}
	if !IsNil(o.AlbumId) {
		toSerialize["albumId"] = o.AlbumId
	}
	if o.TrackNumbers != nil {
		toSerialize["trackNumbers"] = o.TrackNumbers
	}
	if !IsNil(o.TrackFileId) {
		toSerialize["trackFileId"] = o.TrackFileId
	}
	if o.ExistingPath.IsSet() {
		toSerialize["existingPath"] = o.ExistingPath.Get()
	}
	if o.NewPath.IsSet() {
		toSerialize["newPath"] = o.NewPath.Get()
	}
	return toSerialize, nil
}

type NullableRenameTrackResource struct {
	value *RenameTrackResource
	isSet bool
}

func (v NullableRenameTrackResource) Get() *RenameTrackResource {
	return v.value
}

func (v *NullableRenameTrackResource) Set(val *RenameTrackResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameTrackResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameTrackResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameTrackResource(val *RenameTrackResource) *NullableRenameTrackResource {
	return &NullableRenameTrackResource{value: val, isSet: true}
}

func (v NullableRenameTrackResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameTrackResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


