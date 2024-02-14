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

// RetagTrackResource struct for RetagTrackResource
type RetagTrackResource struct {
	Id *int32 `json:"id,omitempty"`
	ArtistId *int32 `json:"artistId,omitempty"`
	AlbumId *int32 `json:"albumId,omitempty"`
	TrackNumbers []*int32 `json:"trackNumbers,omitempty"`
	TrackFileId *int32 `json:"trackFileId,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Changes []*TagDifference `json:"changes,omitempty"`
}

// NewRetagTrackResource instantiates a new RetagTrackResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetagTrackResource() *RetagTrackResource {
	this := RetagTrackResource{}
	return &this
}

// NewRetagTrackResourceWithDefaults instantiates a new RetagTrackResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetagTrackResourceWithDefaults() *RetagTrackResource {
	this := RetagTrackResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RetagTrackResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetagTrackResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RetagTrackResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RetagTrackResource) SetId(v int32) {
	o.Id = &v
}

// GetArtistId returns the ArtistId field value if set, zero value otherwise.
func (o *RetagTrackResource) GetArtistId() int32 {
	if o == nil || isNil(o.ArtistId) {
		var ret int32
		return ret
	}
	return *o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetagTrackResource) GetArtistIdOk() (*int32, bool) {
	if o == nil || isNil(o.ArtistId) {
    return nil, false
	}
	return o.ArtistId, true
}

// HasArtistId returns a boolean if a field has been set.
func (o *RetagTrackResource) HasArtistId() bool {
	if o != nil && !isNil(o.ArtistId) {
		return true
	}

	return false
}

// SetArtistId gets a reference to the given int32 and assigns it to the ArtistId field.
func (o *RetagTrackResource) SetArtistId(v int32) {
	o.ArtistId = &v
}

// GetAlbumId returns the AlbumId field value if set, zero value otherwise.
func (o *RetagTrackResource) GetAlbumId() int32 {
	if o == nil || isNil(o.AlbumId) {
		var ret int32
		return ret
	}
	return *o.AlbumId
}

// GetAlbumIdOk returns a tuple with the AlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetagTrackResource) GetAlbumIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlbumId) {
    return nil, false
	}
	return o.AlbumId, true
}

// HasAlbumId returns a boolean if a field has been set.
func (o *RetagTrackResource) HasAlbumId() bool {
	if o != nil && !isNil(o.AlbumId) {
		return true
	}

	return false
}

// SetAlbumId gets a reference to the given int32 and assigns it to the AlbumId field.
func (o *RetagTrackResource) SetAlbumId(v int32) {
	o.AlbumId = &v
}

// GetTrackNumbers returns the TrackNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetagTrackResource) GetTrackNumbers() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.TrackNumbers
}

// GetTrackNumbersOk returns a tuple with the TrackNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetagTrackResource) GetTrackNumbersOk() ([]*int32, bool) {
	if o == nil || isNil(o.TrackNumbers) {
    return nil, false
	}
	return o.TrackNumbers, true
}

// HasTrackNumbers returns a boolean if a field has been set.
func (o *RetagTrackResource) HasTrackNumbers() bool {
	if o != nil && isNil(o.TrackNumbers) {
		return true
	}

	return false
}

// SetTrackNumbers gets a reference to the given []int32 and assigns it to the TrackNumbers field.
func (o *RetagTrackResource) SetTrackNumbers(v []*int32) {
	o.TrackNumbers = v
}

// GetTrackFileId returns the TrackFileId field value if set, zero value otherwise.
func (o *RetagTrackResource) GetTrackFileId() int32 {
	if o == nil || isNil(o.TrackFileId) {
		var ret int32
		return ret
	}
	return *o.TrackFileId
}

// GetTrackFileIdOk returns a tuple with the TrackFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetagTrackResource) GetTrackFileIdOk() (*int32, bool) {
	if o == nil || isNil(o.TrackFileId) {
    return nil, false
	}
	return o.TrackFileId, true
}

// HasTrackFileId returns a boolean if a field has been set.
func (o *RetagTrackResource) HasTrackFileId() bool {
	if o != nil && !isNil(o.TrackFileId) {
		return true
	}

	return false
}

// SetTrackFileId gets a reference to the given int32 and assigns it to the TrackFileId field.
func (o *RetagTrackResource) SetTrackFileId(v int32) {
	o.TrackFileId = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetagTrackResource) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetagTrackResource) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *RetagTrackResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *RetagTrackResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *RetagTrackResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *RetagTrackResource) UnsetPath() {
	o.Path.Unset()
}

// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetagTrackResource) GetChanges() []*TagDifference {
	if o == nil {
		var ret []*TagDifference
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetagTrackResource) GetChangesOk() ([]*TagDifference, bool) {
	if o == nil || isNil(o.Changes) {
    return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *RetagTrackResource) HasChanges() bool {
	if o != nil && isNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []TagDifference and assigns it to the Changes field.
func (o *RetagTrackResource) SetChanges(v []*TagDifference) {
	o.Changes = v
}

func (o RetagTrackResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ArtistId) {
		toSerialize["artistId"] = o.ArtistId
	}
	if !isNil(o.AlbumId) {
		toSerialize["albumId"] = o.AlbumId
	}
	if o.TrackNumbers != nil {
		toSerialize["trackNumbers"] = o.TrackNumbers
	}
	if !isNil(o.TrackFileId) {
		toSerialize["trackFileId"] = o.TrackFileId
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	return json.Marshal(toSerialize)
}

type NullableRetagTrackResource struct {
	value *RetagTrackResource
	isSet bool
}

func (v NullableRetagTrackResource) Get() *RetagTrackResource {
	return v.value
}

func (v *NullableRetagTrackResource) Set(val *RetagTrackResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRetagTrackResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRetagTrackResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetagTrackResource(val *RetagTrackResource) *NullableRetagTrackResource {
	return &NullableRetagTrackResource{value: val, isSet: true}
}

func (v NullableRetagTrackResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetagTrackResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


