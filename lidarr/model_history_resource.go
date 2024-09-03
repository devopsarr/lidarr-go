/*
Lidarr

Lidarr API docs

API version: v2.5.3.4341
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"time"
)

// checks if the HistoryResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryResource{}

// HistoryResource struct for HistoryResource
type HistoryResource struct {
	Id *int32 `json:"id,omitempty"`
	AlbumId *int32 `json:"albumId,omitempty"`
	ArtistId *int32 `json:"artistId,omitempty"`
	TrackId *int32 `json:"trackId,omitempty"`
	SourceTitle NullableString `json:"sourceTitle,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	QualityCutoffNotMet *bool `json:"qualityCutoffNotMet,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	EventType *EntityHistoryEventType `json:"eventType,omitempty"`
	Data map[string]string `json:"data,omitempty"`
	Album *AlbumResource `json:"album,omitempty"`
	Artist *ArtistResource `json:"artist,omitempty"`
	Track *TrackResource `json:"track,omitempty"`
}

// NewHistoryResource instantiates a new HistoryResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryResource() *HistoryResource {
	this := HistoryResource{}
	return &this
}

// NewHistoryResourceWithDefaults instantiates a new HistoryResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryResourceWithDefaults() *HistoryResource {
	this := HistoryResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HistoryResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HistoryResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *HistoryResource) SetId(v int32) {
	o.Id = &v
}

// GetAlbumId returns the AlbumId field value if set, zero value otherwise.
func (o *HistoryResource) GetAlbumId() int32 {
	if o == nil || IsNil(o.AlbumId) {
		var ret int32
		return ret
	}
	return *o.AlbumId
}

// GetAlbumIdOk returns a tuple with the AlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetAlbumIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AlbumId) {
		return nil, false
	}
	return o.AlbumId, true
}

// HasAlbumId returns a boolean if a field has been set.
func (o *HistoryResource) HasAlbumId() bool {
	if o != nil && !IsNil(o.AlbumId) {
		return true
	}

	return false
}

// SetAlbumId gets a reference to the given int32 and assigns it to the AlbumId field.
func (o *HistoryResource) SetAlbumId(v int32) {
	o.AlbumId = &v
}

// GetArtistId returns the ArtistId field value if set, zero value otherwise.
func (o *HistoryResource) GetArtistId() int32 {
	if o == nil || IsNil(o.ArtistId) {
		var ret int32
		return ret
	}
	return *o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetArtistIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ArtistId) {
		return nil, false
	}
	return o.ArtistId, true
}

// HasArtistId returns a boolean if a field has been set.
func (o *HistoryResource) HasArtistId() bool {
	if o != nil && !IsNil(o.ArtistId) {
		return true
	}

	return false
}

// SetArtistId gets a reference to the given int32 and assigns it to the ArtistId field.
func (o *HistoryResource) SetArtistId(v int32) {
	o.ArtistId = &v
}

// GetTrackId returns the TrackId field value if set, zero value otherwise.
func (o *HistoryResource) GetTrackId() int32 {
	if o == nil || IsNil(o.TrackId) {
		var ret int32
		return ret
	}
	return *o.TrackId
}

// GetTrackIdOk returns a tuple with the TrackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetTrackIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TrackId) {
		return nil, false
	}
	return o.TrackId, true
}

// HasTrackId returns a boolean if a field has been set.
func (o *HistoryResource) HasTrackId() bool {
	if o != nil && !IsNil(o.TrackId) {
		return true
	}

	return false
}

// SetTrackId gets a reference to the given int32 and assigns it to the TrackId field.
func (o *HistoryResource) SetTrackId(v int32) {
	o.TrackId = &v
}

// GetSourceTitle returns the SourceTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetSourceTitle() string {
	if o == nil || IsNil(o.SourceTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SourceTitle.Get()
}

// GetSourceTitleOk returns a tuple with the SourceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetSourceTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceTitle.Get(), o.SourceTitle.IsSet()
}

// HasSourceTitle returns a boolean if a field has been set.
func (o *HistoryResource) HasSourceTitle() bool {
	if o != nil && o.SourceTitle.IsSet() {
		return true
	}

	return false
}

// SetSourceTitle gets a reference to the given NullableString and assigns it to the SourceTitle field.
func (o *HistoryResource) SetSourceTitle(v string) {
	o.SourceTitle.Set(&v)
}
// SetSourceTitleNil sets the value for SourceTitle to be an explicit nil
func (o *HistoryResource) SetSourceTitleNil() {
	o.SourceTitle.Set(nil)
}

// UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
func (o *HistoryResource) UnsetSourceTitle() {
	o.SourceTitle.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *HistoryResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *HistoryResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *HistoryResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *HistoryResource) HasCustomFormats() bool {
	if o != nil && !IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *HistoryResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *HistoryResource) GetCustomFormatScore() int32 {
	if o == nil || IsNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomFormatScore) {
		return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *HistoryResource) HasCustomFormatScore() bool {
	if o != nil && !IsNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *HistoryResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetQualityCutoffNotMet returns the QualityCutoffNotMet field value if set, zero value otherwise.
func (o *HistoryResource) GetQualityCutoffNotMet() bool {
	if o == nil || IsNil(o.QualityCutoffNotMet) {
		var ret bool
		return ret
	}
	return *o.QualityCutoffNotMet
}

// GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetQualityCutoffNotMetOk() (*bool, bool) {
	if o == nil || IsNil(o.QualityCutoffNotMet) {
		return nil, false
	}
	return o.QualityCutoffNotMet, true
}

// HasQualityCutoffNotMet returns a boolean if a field has been set.
func (o *HistoryResource) HasQualityCutoffNotMet() bool {
	if o != nil && !IsNil(o.QualityCutoffNotMet) {
		return true
	}

	return false
}

// SetQualityCutoffNotMet gets a reference to the given bool and assigns it to the QualityCutoffNotMet field.
func (o *HistoryResource) SetQualityCutoffNotMet(v bool) {
	o.QualityCutoffNotMet = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *HistoryResource) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *HistoryResource) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *HistoryResource) SetDate(v time.Time) {
	o.Date = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *HistoryResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *HistoryResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *HistoryResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *HistoryResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *HistoryResource) GetEventType() EntityHistoryEventType {
	if o == nil || IsNil(o.EventType) {
		var ret EntityHistoryEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetEventTypeOk() (*EntityHistoryEventType, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *HistoryResource) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given EntityHistoryEventType and assigns it to the EventType field.
func (o *HistoryResource) SetEventType(v EntityHistoryEventType) {
	o.EventType = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoryResource) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *HistoryResource) SetData(v map[string]string) {
	o.Data = v
}

// GetAlbum returns the Album field value if set, zero value otherwise.
func (o *HistoryResource) GetAlbum() AlbumResource {
	if o == nil || IsNil(o.Album) {
		var ret AlbumResource
		return ret
	}
	return *o.Album
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetAlbumOk() (*AlbumResource, bool) {
	if o == nil || IsNil(o.Album) {
		return nil, false
	}
	return o.Album, true
}

// HasAlbum returns a boolean if a field has been set.
func (o *HistoryResource) HasAlbum() bool {
	if o != nil && !IsNil(o.Album) {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given AlbumResource and assigns it to the Album field.
func (o *HistoryResource) SetAlbum(v AlbumResource) {
	o.Album = &v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *HistoryResource) GetArtist() ArtistResource {
	if o == nil || IsNil(o.Artist) {
		var ret ArtistResource
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetArtistOk() (*ArtistResource, bool) {
	if o == nil || IsNil(o.Artist) {
		return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *HistoryResource) HasArtist() bool {
	if o != nil && !IsNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistResource and assigns it to the Artist field.
func (o *HistoryResource) SetArtist(v ArtistResource) {
	o.Artist = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *HistoryResource) GetTrack() TrackResource {
	if o == nil || IsNil(o.Track) {
		var ret TrackResource
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetTrackOk() (*TrackResource, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *HistoryResource) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given TrackResource and assigns it to the Track field.
func (o *HistoryResource) SetTrack(v TrackResource) {
	o.Track = &v
}

func (o HistoryResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AlbumId) {
		toSerialize["albumId"] = o.AlbumId
	}
	if !IsNil(o.ArtistId) {
		toSerialize["artistId"] = o.ArtistId
	}
	if !IsNil(o.TrackId) {
		toSerialize["trackId"] = o.TrackId
	}
	if o.SourceTitle.IsSet() {
		toSerialize["sourceTitle"] = o.SourceTitle.Get()
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !IsNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	if !IsNil(o.QualityCutoffNotMet) {
		toSerialize["qualityCutoffNotMet"] = o.QualityCutoffNotMet
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Album) {
		toSerialize["album"] = o.Album
	}
	if !IsNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	return toSerialize, nil
}

type NullableHistoryResource struct {
	value *HistoryResource
	isSet bool
}

func (v NullableHistoryResource) Get() *HistoryResource {
	return v.value
}

func (v *NullableHistoryResource) Set(val *HistoryResource) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryResource) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryResource(val *HistoryResource) *NullableHistoryResource {
	return &NullableHistoryResource{value: val, isSet: true}
}

func (v NullableHistoryResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


