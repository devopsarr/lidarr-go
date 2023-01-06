/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// TrackResource struct for TrackResource
type TrackResource struct {
	Id *int32 `json:"id,omitempty"`
	ArtistId *int32 `json:"artistId,omitempty"`
	ForeignTrackId NullableString `json:"foreignTrackId,omitempty"`
	ForeignRecordingId NullableString `json:"foreignRecordingId,omitempty"`
	TrackFileId *int32 `json:"trackFileId,omitempty"`
	AlbumId *int32 `json:"albumId,omitempty"`
	Explicit *bool `json:"explicit,omitempty"`
	AbsoluteTrackNumber *int32 `json:"absoluteTrackNumber,omitempty"`
	TrackNumber NullableString `json:"trackNumber,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	TrackFile *TrackFileResource `json:"trackFile,omitempty"`
	MediumNumber *int32 `json:"mediumNumber,omitempty"`
	HasFile *bool `json:"hasFile,omitempty"`
	Artist *ArtistResource `json:"artist,omitempty"`
	Ratings *Ratings `json:"ratings,omitempty"`
	Grabbed *bool `json:"grabbed,omitempty"`
}

// NewTrackResource instantiates a new TrackResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackResource() *TrackResource {
	this := TrackResource{}
	return &this
}

// NewTrackResourceWithDefaults instantiates a new TrackResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackResourceWithDefaults() *TrackResource {
	this := TrackResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TrackResource) SetId(v int32) {
	o.Id = &v
}

// GetArtistId returns the ArtistId field value if set, zero value otherwise.
func (o *TrackResource) GetArtistId() int32 {
	if o == nil || isNil(o.ArtistId) {
		var ret int32
		return ret
	}
	return *o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetArtistIdOk() (*int32, bool) {
	if o == nil || isNil(o.ArtistId) {
    return nil, false
	}
	return o.ArtistId, true
}

// HasArtistId returns a boolean if a field has been set.
func (o *TrackResource) HasArtistId() bool {
	if o != nil && !isNil(o.ArtistId) {
		return true
	}

	return false
}

// SetArtistId gets a reference to the given int32 and assigns it to the ArtistId field.
func (o *TrackResource) SetArtistId(v int32) {
	o.ArtistId = &v
}

// GetForeignTrackId returns the ForeignTrackId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackResource) GetForeignTrackId() string {
	if o == nil || isNil(o.ForeignTrackId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignTrackId.Get()
}

// GetForeignTrackIdOk returns a tuple with the ForeignTrackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackResource) GetForeignTrackIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignTrackId.Get(), o.ForeignTrackId.IsSet()
}

// HasForeignTrackId returns a boolean if a field has been set.
func (o *TrackResource) HasForeignTrackId() bool {
	if o != nil && o.ForeignTrackId.IsSet() {
		return true
	}

	return false
}

// SetForeignTrackId gets a reference to the given NullableString and assigns it to the ForeignTrackId field.
func (o *TrackResource) SetForeignTrackId(v string) {
	o.ForeignTrackId.Set(&v)
}
// SetForeignTrackIdNil sets the value for ForeignTrackId to be an explicit nil
func (o *TrackResource) SetForeignTrackIdNil() {
	o.ForeignTrackId.Set(nil)
}

// UnsetForeignTrackId ensures that no value is present for ForeignTrackId, not even an explicit nil
func (o *TrackResource) UnsetForeignTrackId() {
	o.ForeignTrackId.Unset()
}

// GetForeignRecordingId returns the ForeignRecordingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackResource) GetForeignRecordingId() string {
	if o == nil || isNil(o.ForeignRecordingId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignRecordingId.Get()
}

// GetForeignRecordingIdOk returns a tuple with the ForeignRecordingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackResource) GetForeignRecordingIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignRecordingId.Get(), o.ForeignRecordingId.IsSet()
}

// HasForeignRecordingId returns a boolean if a field has been set.
func (o *TrackResource) HasForeignRecordingId() bool {
	if o != nil && o.ForeignRecordingId.IsSet() {
		return true
	}

	return false
}

// SetForeignRecordingId gets a reference to the given NullableString and assigns it to the ForeignRecordingId field.
func (o *TrackResource) SetForeignRecordingId(v string) {
	o.ForeignRecordingId.Set(&v)
}
// SetForeignRecordingIdNil sets the value for ForeignRecordingId to be an explicit nil
func (o *TrackResource) SetForeignRecordingIdNil() {
	o.ForeignRecordingId.Set(nil)
}

// UnsetForeignRecordingId ensures that no value is present for ForeignRecordingId, not even an explicit nil
func (o *TrackResource) UnsetForeignRecordingId() {
	o.ForeignRecordingId.Unset()
}

// GetTrackFileId returns the TrackFileId field value if set, zero value otherwise.
func (o *TrackResource) GetTrackFileId() int32 {
	if o == nil || isNil(o.TrackFileId) {
		var ret int32
		return ret
	}
	return *o.TrackFileId
}

// GetTrackFileIdOk returns a tuple with the TrackFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetTrackFileIdOk() (*int32, bool) {
	if o == nil || isNil(o.TrackFileId) {
    return nil, false
	}
	return o.TrackFileId, true
}

// HasTrackFileId returns a boolean if a field has been set.
func (o *TrackResource) HasTrackFileId() bool {
	if o != nil && !isNil(o.TrackFileId) {
		return true
	}

	return false
}

// SetTrackFileId gets a reference to the given int32 and assigns it to the TrackFileId field.
func (o *TrackResource) SetTrackFileId(v int32) {
	o.TrackFileId = &v
}

// GetAlbumId returns the AlbumId field value if set, zero value otherwise.
func (o *TrackResource) GetAlbumId() int32 {
	if o == nil || isNil(o.AlbumId) {
		var ret int32
		return ret
	}
	return *o.AlbumId
}

// GetAlbumIdOk returns a tuple with the AlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetAlbumIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlbumId) {
    return nil, false
	}
	return o.AlbumId, true
}

// HasAlbumId returns a boolean if a field has been set.
func (o *TrackResource) HasAlbumId() bool {
	if o != nil && !isNil(o.AlbumId) {
		return true
	}

	return false
}

// SetAlbumId gets a reference to the given int32 and assigns it to the AlbumId field.
func (o *TrackResource) SetAlbumId(v int32) {
	o.AlbumId = &v
}

// GetExplicit returns the Explicit field value if set, zero value otherwise.
func (o *TrackResource) GetExplicit() bool {
	if o == nil || isNil(o.Explicit) {
		var ret bool
		return ret
	}
	return *o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetExplicitOk() (*bool, bool) {
	if o == nil || isNil(o.Explicit) {
    return nil, false
	}
	return o.Explicit, true
}

// HasExplicit returns a boolean if a field has been set.
func (o *TrackResource) HasExplicit() bool {
	if o != nil && !isNil(o.Explicit) {
		return true
	}

	return false
}

// SetExplicit gets a reference to the given bool and assigns it to the Explicit field.
func (o *TrackResource) SetExplicit(v bool) {
	o.Explicit = &v
}

// GetAbsoluteTrackNumber returns the AbsoluteTrackNumber field value if set, zero value otherwise.
func (o *TrackResource) GetAbsoluteTrackNumber() int32 {
	if o == nil || isNil(o.AbsoluteTrackNumber) {
		var ret int32
		return ret
	}
	return *o.AbsoluteTrackNumber
}

// GetAbsoluteTrackNumberOk returns a tuple with the AbsoluteTrackNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetAbsoluteTrackNumberOk() (*int32, bool) {
	if o == nil || isNil(o.AbsoluteTrackNumber) {
    return nil, false
	}
	return o.AbsoluteTrackNumber, true
}

// HasAbsoluteTrackNumber returns a boolean if a field has been set.
func (o *TrackResource) HasAbsoluteTrackNumber() bool {
	if o != nil && !isNil(o.AbsoluteTrackNumber) {
		return true
	}

	return false
}

// SetAbsoluteTrackNumber gets a reference to the given int32 and assigns it to the AbsoluteTrackNumber field.
func (o *TrackResource) SetAbsoluteTrackNumber(v int32) {
	o.AbsoluteTrackNumber = &v
}

// GetTrackNumber returns the TrackNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackResource) GetTrackNumber() string {
	if o == nil || isNil(o.TrackNumber.Get()) {
		var ret string
		return ret
	}
	return *o.TrackNumber.Get()
}

// GetTrackNumberOk returns a tuple with the TrackNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackResource) GetTrackNumberOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrackNumber.Get(), o.TrackNumber.IsSet()
}

// HasTrackNumber returns a boolean if a field has been set.
func (o *TrackResource) HasTrackNumber() bool {
	if o != nil && o.TrackNumber.IsSet() {
		return true
	}

	return false
}

// SetTrackNumber gets a reference to the given NullableString and assigns it to the TrackNumber field.
func (o *TrackResource) SetTrackNumber(v string) {
	o.TrackNumber.Set(&v)
}
// SetTrackNumberNil sets the value for TrackNumber to be an explicit nil
func (o *TrackResource) SetTrackNumberNil() {
	o.TrackNumber.Set(nil)
}

// UnsetTrackNumber ensures that no value is present for TrackNumber, not even an explicit nil
func (o *TrackResource) UnsetTrackNumber() {
	o.TrackNumber.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TrackResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TrackResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TrackResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TrackResource) UnsetTitle() {
	o.Title.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TrackResource) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TrackResource) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *TrackResource) SetDuration(v int32) {
	o.Duration = &v
}

// GetTrackFile returns the TrackFile field value if set, zero value otherwise.
func (o *TrackResource) GetTrackFile() TrackFileResource {
	if o == nil || isNil(o.TrackFile) {
		var ret TrackFileResource
		return ret
	}
	return *o.TrackFile
}

// GetTrackFileOk returns a tuple with the TrackFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetTrackFileOk() (*TrackFileResource, bool) {
	if o == nil || isNil(o.TrackFile) {
    return nil, false
	}
	return o.TrackFile, true
}

// HasTrackFile returns a boolean if a field has been set.
func (o *TrackResource) HasTrackFile() bool {
	if o != nil && !isNil(o.TrackFile) {
		return true
	}

	return false
}

// SetTrackFile gets a reference to the given TrackFileResource and assigns it to the TrackFile field.
func (o *TrackResource) SetTrackFile(v TrackFileResource) {
	o.TrackFile = &v
}

// GetMediumNumber returns the MediumNumber field value if set, zero value otherwise.
func (o *TrackResource) GetMediumNumber() int32 {
	if o == nil || isNil(o.MediumNumber) {
		var ret int32
		return ret
	}
	return *o.MediumNumber
}

// GetMediumNumberOk returns a tuple with the MediumNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetMediumNumberOk() (*int32, bool) {
	if o == nil || isNil(o.MediumNumber) {
    return nil, false
	}
	return o.MediumNumber, true
}

// HasMediumNumber returns a boolean if a field has been set.
func (o *TrackResource) HasMediumNumber() bool {
	if o != nil && !isNil(o.MediumNumber) {
		return true
	}

	return false
}

// SetMediumNumber gets a reference to the given int32 and assigns it to the MediumNumber field.
func (o *TrackResource) SetMediumNumber(v int32) {
	o.MediumNumber = &v
}

// GetHasFile returns the HasFile field value if set, zero value otherwise.
func (o *TrackResource) GetHasFile() bool {
	if o == nil || isNil(o.HasFile) {
		var ret bool
		return ret
	}
	return *o.HasFile
}

// GetHasFileOk returns a tuple with the HasFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetHasFileOk() (*bool, bool) {
	if o == nil || isNil(o.HasFile) {
    return nil, false
	}
	return o.HasFile, true
}

// HasHasFile returns a boolean if a field has been set.
func (o *TrackResource) HasHasFile() bool {
	if o != nil && !isNil(o.HasFile) {
		return true
	}

	return false
}

// SetHasFile gets a reference to the given bool and assigns it to the HasFile field.
func (o *TrackResource) SetHasFile(v bool) {
	o.HasFile = &v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *TrackResource) GetArtist() ArtistResource {
	if o == nil || isNil(o.Artist) {
		var ret ArtistResource
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetArtistOk() (*ArtistResource, bool) {
	if o == nil || isNil(o.Artist) {
    return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *TrackResource) HasArtist() bool {
	if o != nil && !isNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistResource and assigns it to the Artist field.
func (o *TrackResource) SetArtist(v ArtistResource) {
	o.Artist = &v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *TrackResource) GetRatings() Ratings {
	if o == nil || isNil(o.Ratings) {
		var ret Ratings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetRatingsOk() (*Ratings, bool) {
	if o == nil || isNil(o.Ratings) {
    return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *TrackResource) HasRatings() bool {
	if o != nil && !isNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given Ratings and assigns it to the Ratings field.
func (o *TrackResource) SetRatings(v Ratings) {
	o.Ratings = &v
}

// GetGrabbed returns the Grabbed field value if set, zero value otherwise.
func (o *TrackResource) GetGrabbed() bool {
	if o == nil || isNil(o.Grabbed) {
		var ret bool
		return ret
	}
	return *o.Grabbed
}

// GetGrabbedOk returns a tuple with the Grabbed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackResource) GetGrabbedOk() (*bool, bool) {
	if o == nil || isNil(o.Grabbed) {
    return nil, false
	}
	return o.Grabbed, true
}

// HasGrabbed returns a boolean if a field has been set.
func (o *TrackResource) HasGrabbed() bool {
	if o != nil && !isNil(o.Grabbed) {
		return true
	}

	return false
}

// SetGrabbed gets a reference to the given bool and assigns it to the Grabbed field.
func (o *TrackResource) SetGrabbed(v bool) {
	o.Grabbed = &v
}

func (o TrackResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ArtistId) {
		toSerialize["artistId"] = o.ArtistId
	}
	if o.ForeignTrackId.IsSet() {
		toSerialize["foreignTrackId"] = o.ForeignTrackId.Get()
	}
	if o.ForeignRecordingId.IsSet() {
		toSerialize["foreignRecordingId"] = o.ForeignRecordingId.Get()
	}
	if !isNil(o.TrackFileId) {
		toSerialize["trackFileId"] = o.TrackFileId
	}
	if !isNil(o.AlbumId) {
		toSerialize["albumId"] = o.AlbumId
	}
	if !isNil(o.Explicit) {
		toSerialize["explicit"] = o.Explicit
	}
	if !isNil(o.AbsoluteTrackNumber) {
		toSerialize["absoluteTrackNumber"] = o.AbsoluteTrackNumber
	}
	if o.TrackNumber.IsSet() {
		toSerialize["trackNumber"] = o.TrackNumber.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.TrackFile) {
		toSerialize["trackFile"] = o.TrackFile
	}
	if !isNil(o.MediumNumber) {
		toSerialize["mediumNumber"] = o.MediumNumber
	}
	if !isNil(o.HasFile) {
		toSerialize["hasFile"] = o.HasFile
	}
	if !isNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if !isNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if !isNil(o.Grabbed) {
		toSerialize["grabbed"] = o.Grabbed
	}
	return json.Marshal(toSerialize)
}

type NullableTrackResource struct {
	value *TrackResource
	isSet bool
}

func (v NullableTrackResource) Get() *TrackResource {
	return v.value
}

func (v *NullableTrackResource) Set(val *TrackResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackResource(val *TrackResource) *NullableTrackResource {
	return &NullableTrackResource{value: val, isSet: true}
}

func (v NullableTrackResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


