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

// TrackFile struct for TrackFile
type TrackFile struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	DateAdded *time.Time `json:"dateAdded,omitempty"`
	SceneName NullableString `json:"sceneName,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	MediaInfo *MediaInfoModel `json:"mediaInfo,omitempty"`
	AlbumId *int32 `json:"albumId,omitempty"`
	Tracks *TrackListLazyLoaded `json:"tracks,omitempty"`
	Artist *ArtistLazyLoaded `json:"artist,omitempty"`
	Album *AlbumLazyLoaded `json:"album,omitempty"`
}

// NewTrackFile instantiates a new TrackFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackFile() *TrackFile {
	this := TrackFile{}
	return &this
}

// NewTrackFileWithDefaults instantiates a new TrackFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackFileWithDefaults() *TrackFile {
	this := TrackFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackFile) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackFile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TrackFile) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackFile) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackFile) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *TrackFile) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *TrackFile) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *TrackFile) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *TrackFile) UnsetPath() {
	o.Path.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TrackFile) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TrackFile) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *TrackFile) SetSize(v int64) {
	o.Size = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *TrackFile) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
    return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *TrackFile) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *TrackFile) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *TrackFile) GetDateAdded() time.Time {
	if o == nil || isNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateAdded) {
    return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *TrackFile) HasDateAdded() bool {
	if o != nil && !isNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *TrackFile) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackFile) GetSceneName() string {
	if o == nil || isNil(o.SceneName.Get()) {
		var ret string
		return ret
	}
	return *o.SceneName.Get()
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackFile) GetSceneNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SceneName.Get(), o.SceneName.IsSet()
}

// HasSceneName returns a boolean if a field has been set.
func (o *TrackFile) HasSceneName() bool {
	if o != nil && o.SceneName.IsSet() {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given NullableString and assigns it to the SceneName field.
func (o *TrackFile) SetSceneName(v string) {
	o.SceneName.Set(&v)
}
// SetSceneNameNil sets the value for SceneName to be an explicit nil
func (o *TrackFile) SetSceneNameNil() {
	o.SceneName.Set(nil)
}

// UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
func (o *TrackFile) UnsetSceneName() {
	o.SceneName.Unset()
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackFile) GetReleaseGroup() string {
	if o == nil || isNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackFile) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *TrackFile) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *TrackFile) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *TrackFile) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *TrackFile) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *TrackFile) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *TrackFile) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *TrackFile) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *TrackFile) GetMediaInfo() MediaInfoModel {
	if o == nil || isNil(o.MediaInfo) {
		var ret MediaInfoModel
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetMediaInfoOk() (*MediaInfoModel, bool) {
	if o == nil || isNil(o.MediaInfo) {
    return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *TrackFile) HasMediaInfo() bool {
	if o != nil && !isNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfoModel and assigns it to the MediaInfo field.
func (o *TrackFile) SetMediaInfo(v MediaInfoModel) {
	o.MediaInfo = &v
}

// GetAlbumId returns the AlbumId field value if set, zero value otherwise.
func (o *TrackFile) GetAlbumId() int32 {
	if o == nil || isNil(o.AlbumId) {
		var ret int32
		return ret
	}
	return *o.AlbumId
}

// GetAlbumIdOk returns a tuple with the AlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetAlbumIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlbumId) {
    return nil, false
	}
	return o.AlbumId, true
}

// HasAlbumId returns a boolean if a field has been set.
func (o *TrackFile) HasAlbumId() bool {
	if o != nil && !isNil(o.AlbumId) {
		return true
	}

	return false
}

// SetAlbumId gets a reference to the given int32 and assigns it to the AlbumId field.
func (o *TrackFile) SetAlbumId(v int32) {
	o.AlbumId = &v
}

// GetTracks returns the Tracks field value if set, zero value otherwise.
func (o *TrackFile) GetTracks() TrackListLazyLoaded {
	if o == nil || isNil(o.Tracks) {
		var ret TrackListLazyLoaded
		return ret
	}
	return *o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetTracksOk() (*TrackListLazyLoaded, bool) {
	if o == nil || isNil(o.Tracks) {
    return nil, false
	}
	return o.Tracks, true
}

// HasTracks returns a boolean if a field has been set.
func (o *TrackFile) HasTracks() bool {
	if o != nil && !isNil(o.Tracks) {
		return true
	}

	return false
}

// SetTracks gets a reference to the given TrackListLazyLoaded and assigns it to the Tracks field.
func (o *TrackFile) SetTracks(v TrackListLazyLoaded) {
	o.Tracks = &v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *TrackFile) GetArtist() ArtistLazyLoaded {
	if o == nil || isNil(o.Artist) {
		var ret ArtistLazyLoaded
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetArtistOk() (*ArtistLazyLoaded, bool) {
	if o == nil || isNil(o.Artist) {
    return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *TrackFile) HasArtist() bool {
	if o != nil && !isNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistLazyLoaded and assigns it to the Artist field.
func (o *TrackFile) SetArtist(v ArtistLazyLoaded) {
	o.Artist = &v
}

// GetAlbum returns the Album field value if set, zero value otherwise.
func (o *TrackFile) GetAlbum() AlbumLazyLoaded {
	if o == nil || isNil(o.Album) {
		var ret AlbumLazyLoaded
		return ret
	}
	return *o.Album
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackFile) GetAlbumOk() (*AlbumLazyLoaded, bool) {
	if o == nil || isNil(o.Album) {
    return nil, false
	}
	return o.Album, true
}

// HasAlbum returns a boolean if a field has been set.
func (o *TrackFile) HasAlbum() bool {
	if o != nil && !isNil(o.Album) {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given AlbumLazyLoaded and assigns it to the Album field.
func (o *TrackFile) SetAlbum(v AlbumLazyLoaded) {
	o.Album = &v
}

func (o TrackFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.DateAdded) {
		toSerialize["dateAdded"] = o.DateAdded
	}
	if o.SceneName.IsSet() {
		toSerialize["sceneName"] = o.SceneName.Get()
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !isNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if !isNil(o.AlbumId) {
		toSerialize["albumId"] = o.AlbumId
	}
	if !isNil(o.Tracks) {
		toSerialize["tracks"] = o.Tracks
	}
	if !isNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if !isNil(o.Album) {
		toSerialize["album"] = o.Album
	}
	return json.Marshal(toSerialize)
}

type NullableTrackFile struct {
	value *TrackFile
	isSet bool
}

func (v NullableTrackFile) Get() *TrackFile {
	return v.value
}

func (v *NullableTrackFile) Set(val *TrackFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackFile(val *TrackFile) *NullableTrackFile {
	return &NullableTrackFile{value: val, isSet: true}
}

func (v NullableTrackFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


