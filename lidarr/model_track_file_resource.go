/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"time"
)

// TrackFileResource struct for TrackFileResource
type TrackFileResource struct {
	Id *int32 `json:"id,omitempty"`
	ArtistId *int32 `json:"artistId,omitempty"`
	AlbumId *int32 `json:"albumId,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Size *int64 `json:"size,omitempty"`
	DateAdded *time.Time `json:"dateAdded,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	QualityWeight *int32 `json:"qualityWeight,omitempty"`
	MediaInfo *MediaInfoResource `json:"mediaInfo,omitempty"`
	QualityCutoffNotMet *bool `json:"qualityCutoffNotMet,omitempty"`
	AudioTags *ParsedTrackInfo `json:"audioTags,omitempty"`
}

// NewTrackFileResource instantiates a new TrackFileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackFileResource() *TrackFileResource {
	this := TrackFileResource{}
	return &this
}

// NewTrackFileResourceWithDefaults instantiates a new TrackFileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackFileResourceWithDefaults() *TrackFileResource {
	this := TrackFileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackFileResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackFileResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TrackFileResource) SetId(v int32) {
	o.Id = &v
}

// GetArtistId returns the ArtistId field value if set, zero value otherwise.
func (o *TrackFileResource) GetArtistId() int32 {
	if o == nil || isNil(o.ArtistId) {
		var ret int32
		return ret
	}
	return *o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetArtistIdOk() (*int32, bool) {
	if o == nil || isNil(o.ArtistId) {
    return nil, false
	}
	return o.ArtistId, true
}

// HasArtistId returns a boolean if a field has been set.
func (o *TrackFileResource) HasArtistId() bool {
	if o != nil && !isNil(o.ArtistId) {
		return true
	}

	return false
}

// SetArtistId gets a reference to the given int32 and assigns it to the ArtistId field.
func (o *TrackFileResource) SetArtistId(v int32) {
	o.ArtistId = &v
}

// GetAlbumId returns the AlbumId field value if set, zero value otherwise.
func (o *TrackFileResource) GetAlbumId() int32 {
	if o == nil || isNil(o.AlbumId) {
		var ret int32
		return ret
	}
	return *o.AlbumId
}

// GetAlbumIdOk returns a tuple with the AlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetAlbumIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlbumId) {
    return nil, false
	}
	return o.AlbumId, true
}

// HasAlbumId returns a boolean if a field has been set.
func (o *TrackFileResource) HasAlbumId() bool {
	if o != nil && !isNil(o.AlbumId) {
		return true
	}

	return false
}

// SetAlbumId gets a reference to the given int32 and assigns it to the AlbumId field.
func (o *TrackFileResource) SetAlbumId(v int32) {
	o.AlbumId = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackFileResource) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackFileResource) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *TrackFileResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *TrackFileResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *TrackFileResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *TrackFileResource) UnsetPath() {
	o.Path.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TrackFileResource) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TrackFileResource) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *TrackFileResource) SetSize(v int64) {
	o.Size = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *TrackFileResource) GetDateAdded() time.Time {
	if o == nil || isNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateAdded) {
    return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *TrackFileResource) HasDateAdded() bool {
	if o != nil && !isNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *TrackFileResource) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *TrackFileResource) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *TrackFileResource) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *TrackFileResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetQualityWeight returns the QualityWeight field value if set, zero value otherwise.
func (o *TrackFileResource) GetQualityWeight() int32 {
	if o == nil || isNil(o.QualityWeight) {
		var ret int32
		return ret
	}
	return *o.QualityWeight
}

// GetQualityWeightOk returns a tuple with the QualityWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetQualityWeightOk() (*int32, bool) {
	if o == nil || isNil(o.QualityWeight) {
    return nil, false
	}
	return o.QualityWeight, true
}

// HasQualityWeight returns a boolean if a field has been set.
func (o *TrackFileResource) HasQualityWeight() bool {
	if o != nil && !isNil(o.QualityWeight) {
		return true
	}

	return false
}

// SetQualityWeight gets a reference to the given int32 and assigns it to the QualityWeight field.
func (o *TrackFileResource) SetQualityWeight(v int32) {
	o.QualityWeight = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *TrackFileResource) GetMediaInfo() MediaInfoResource {
	if o == nil || isNil(o.MediaInfo) {
		var ret MediaInfoResource
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetMediaInfoOk() (*MediaInfoResource, bool) {
	if o == nil || isNil(o.MediaInfo) {
    return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *TrackFileResource) HasMediaInfo() bool {
	if o != nil && !isNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfoResource and assigns it to the MediaInfo field.
func (o *TrackFileResource) SetMediaInfo(v MediaInfoResource) {
	o.MediaInfo = &v
}

// GetQualityCutoffNotMet returns the QualityCutoffNotMet field value if set, zero value otherwise.
func (o *TrackFileResource) GetQualityCutoffNotMet() bool {
	if o == nil || isNil(o.QualityCutoffNotMet) {
		var ret bool
		return ret
	}
	return *o.QualityCutoffNotMet
}

// GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetQualityCutoffNotMetOk() (*bool, bool) {
	if o == nil || isNil(o.QualityCutoffNotMet) {
    return nil, false
	}
	return o.QualityCutoffNotMet, true
}

// HasQualityCutoffNotMet returns a boolean if a field has been set.
func (o *TrackFileResource) HasQualityCutoffNotMet() bool {
	if o != nil && !isNil(o.QualityCutoffNotMet) {
		return true
	}

	return false
}

// SetQualityCutoffNotMet gets a reference to the given bool and assigns it to the QualityCutoffNotMet field.
func (o *TrackFileResource) SetQualityCutoffNotMet(v bool) {
	o.QualityCutoffNotMet = &v
}

// GetAudioTags returns the AudioTags field value if set, zero value otherwise.
func (o *TrackFileResource) GetAudioTags() ParsedTrackInfo {
	if o == nil || isNil(o.AudioTags) {
		var ret ParsedTrackInfo
		return ret
	}
	return *o.AudioTags
}

// GetAudioTagsOk returns a tuple with the AudioTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFileResource) GetAudioTagsOk() (*ParsedTrackInfo, bool) {
	if o == nil || isNil(o.AudioTags) {
    return nil, false
	}
	return o.AudioTags, true
}

// HasAudioTags returns a boolean if a field has been set.
func (o *TrackFileResource) HasAudioTags() bool {
	if o != nil && !isNil(o.AudioTags) {
		return true
	}

	return false
}

// SetAudioTags gets a reference to the given ParsedTrackInfo and assigns it to the AudioTags field.
func (o *TrackFileResource) SetAudioTags(v ParsedTrackInfo) {
	o.AudioTags = &v
}

func (o TrackFileResource) MarshalJSON() ([]byte, error) {
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
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.DateAdded) {
		toSerialize["dateAdded"] = o.DateAdded
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !isNil(o.QualityWeight) {
		toSerialize["qualityWeight"] = o.QualityWeight
	}
	if !isNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if !isNil(o.QualityCutoffNotMet) {
		toSerialize["qualityCutoffNotMet"] = o.QualityCutoffNotMet
	}
	if !isNil(o.AudioTags) {
		toSerialize["audioTags"] = o.AudioTags
	}
	return json.Marshal(toSerialize)
}

type NullableTrackFileResource struct {
	value *TrackFileResource
	isSet bool
}

func (v NullableTrackFileResource) Get() *TrackFileResource {
	return v.value
}

func (v *NullableTrackFileResource) Set(val *TrackFileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackFileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackFileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackFileResource(val *TrackFileResource) *NullableTrackFileResource {
	return &NullableTrackFileResource{value: val, isSet: true}
}

func (v NullableTrackFileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackFileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


