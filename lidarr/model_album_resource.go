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

// AlbumResource struct for AlbumResource
type AlbumResource struct {
	Id *int32 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Disambiguation NullableString `json:"disambiguation,omitempty"`
	Overview NullableString `json:"overview,omitempty"`
	ArtistId *int32 `json:"artistId,omitempty"`
	ForeignAlbumId NullableString `json:"foreignAlbumId,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	AnyReleaseOk *bool `json:"anyReleaseOk,omitempty"`
	ProfileId *int32 `json:"profileId,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	AlbumType NullableString `json:"albumType,omitempty"`
	SecondaryTypes []*string `json:"secondaryTypes,omitempty"`
	MediumCount *int32 `json:"mediumCount,omitempty"`
	Ratings *Ratings `json:"ratings,omitempty"`
	ReleaseDate NullableTime `json:"releaseDate,omitempty"`
	Releases []*AlbumReleaseResource `json:"releases,omitempty"`
	Genres []*string `json:"genres,omitempty"`
	Media []*MediumResource `json:"media,omitempty"`
	Artist *ArtistResource `json:"artist,omitempty"`
	Images []*MediaCover `json:"images,omitempty"`
	Links []*Links `json:"links,omitempty"`
	Statistics *AlbumStatisticsResource `json:"statistics,omitempty"`
	AddOptions *AddAlbumOptions `json:"addOptions,omitempty"`
	RemoteCover NullableString `json:"remoteCover,omitempty"`
	Grabbed *bool `json:"grabbed,omitempty"`
}

// NewAlbumResource instantiates a new AlbumResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumResource() *AlbumResource {
	this := AlbumResource{}
	return &this
}

// NewAlbumResourceWithDefaults instantiates a new AlbumResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumResourceWithDefaults() *AlbumResource {
	this := AlbumResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlbumResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlbumResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlbumResource) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AlbumResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AlbumResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AlbumResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AlbumResource) UnsetTitle() {
	o.Title.Unset()
}

// GetDisambiguation returns the Disambiguation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetDisambiguation() string {
	if o == nil || isNil(o.Disambiguation.Get()) {
		var ret string
		return ret
	}
	return *o.Disambiguation.Get()
}

// GetDisambiguationOk returns a tuple with the Disambiguation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetDisambiguationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Disambiguation.Get(), o.Disambiguation.IsSet()
}

// HasDisambiguation returns a boolean if a field has been set.
func (o *AlbumResource) HasDisambiguation() bool {
	if o != nil && o.Disambiguation.IsSet() {
		return true
	}

	return false
}

// SetDisambiguation gets a reference to the given NullableString and assigns it to the Disambiguation field.
func (o *AlbumResource) SetDisambiguation(v string) {
	o.Disambiguation.Set(&v)
}
// SetDisambiguationNil sets the value for Disambiguation to be an explicit nil
func (o *AlbumResource) SetDisambiguationNil() {
	o.Disambiguation.Set(nil)
}

// UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
func (o *AlbumResource) UnsetDisambiguation() {
	o.Disambiguation.Unset()
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetOverview() string {
	if o == nil || isNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetOverviewOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *AlbumResource) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *AlbumResource) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *AlbumResource) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *AlbumResource) UnsetOverview() {
	o.Overview.Unset()
}

// GetArtistId returns the ArtistId field value if set, zero value otherwise.
func (o *AlbumResource) GetArtistId() int32 {
	if o == nil || isNil(o.ArtistId) {
		var ret int32
		return ret
	}
	return *o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetArtistIdOk() (*int32, bool) {
	if o == nil || isNil(o.ArtistId) {
    return nil, false
	}
	return o.ArtistId, true
}

// HasArtistId returns a boolean if a field has been set.
func (o *AlbumResource) HasArtistId() bool {
	if o != nil && !isNil(o.ArtistId) {
		return true
	}

	return false
}

// SetArtistId gets a reference to the given int32 and assigns it to the ArtistId field.
func (o *AlbumResource) SetArtistId(v int32) {
	o.ArtistId = &v
}

// GetForeignAlbumId returns the ForeignAlbumId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetForeignAlbumId() string {
	if o == nil || isNil(o.ForeignAlbumId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignAlbumId.Get()
}

// GetForeignAlbumIdOk returns a tuple with the ForeignAlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetForeignAlbumIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignAlbumId.Get(), o.ForeignAlbumId.IsSet()
}

// HasForeignAlbumId returns a boolean if a field has been set.
func (o *AlbumResource) HasForeignAlbumId() bool {
	if o != nil && o.ForeignAlbumId.IsSet() {
		return true
	}

	return false
}

// SetForeignAlbumId gets a reference to the given NullableString and assigns it to the ForeignAlbumId field.
func (o *AlbumResource) SetForeignAlbumId(v string) {
	o.ForeignAlbumId.Set(&v)
}
// SetForeignAlbumIdNil sets the value for ForeignAlbumId to be an explicit nil
func (o *AlbumResource) SetForeignAlbumIdNil() {
	o.ForeignAlbumId.Set(nil)
}

// UnsetForeignAlbumId ensures that no value is present for ForeignAlbumId, not even an explicit nil
func (o *AlbumResource) UnsetForeignAlbumId() {
	o.ForeignAlbumId.Unset()
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *AlbumResource) GetMonitored() bool {
	if o == nil || isNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetMonitoredOk() (*bool, bool) {
	if o == nil || isNil(o.Monitored) {
    return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *AlbumResource) HasMonitored() bool {
	if o != nil && !isNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *AlbumResource) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetAnyReleaseOk returns the AnyReleaseOk field value if set, zero value otherwise.
func (o *AlbumResource) GetAnyReleaseOk() bool {
	if o == nil || isNil(o.AnyReleaseOk) {
		var ret bool
		return ret
	}
	return *o.AnyReleaseOk
}

// GetAnyReleaseOkOk returns a tuple with the AnyReleaseOk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetAnyReleaseOkOk() (*bool, bool) {
	if o == nil || isNil(o.AnyReleaseOk) {
    return nil, false
	}
	return o.AnyReleaseOk, true
}

// HasAnyReleaseOk returns a boolean if a field has been set.
func (o *AlbumResource) HasAnyReleaseOk() bool {
	if o != nil && !isNil(o.AnyReleaseOk) {
		return true
	}

	return false
}

// SetAnyReleaseOk gets a reference to the given bool and assigns it to the AnyReleaseOk field.
func (o *AlbumResource) SetAnyReleaseOk(v bool) {
	o.AnyReleaseOk = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AlbumResource) GetProfileId() int32 {
	if o == nil || isNil(o.ProfileId) {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetProfileIdOk() (*int32, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AlbumResource) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *AlbumResource) SetProfileId(v int32) {
	o.ProfileId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AlbumResource) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AlbumResource) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AlbumResource) SetDuration(v int32) {
	o.Duration = &v
}

// GetAlbumType returns the AlbumType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetAlbumType() string {
	if o == nil || isNil(o.AlbumType.Get()) {
		var ret string
		return ret
	}
	return *o.AlbumType.Get()
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetAlbumTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AlbumType.Get(), o.AlbumType.IsSet()
}

// HasAlbumType returns a boolean if a field has been set.
func (o *AlbumResource) HasAlbumType() bool {
	if o != nil && o.AlbumType.IsSet() {
		return true
	}

	return false
}

// SetAlbumType gets a reference to the given NullableString and assigns it to the AlbumType field.
func (o *AlbumResource) SetAlbumType(v string) {
	o.AlbumType.Set(&v)
}
// SetAlbumTypeNil sets the value for AlbumType to be an explicit nil
func (o *AlbumResource) SetAlbumTypeNil() {
	o.AlbumType.Set(nil)
}

// UnsetAlbumType ensures that no value is present for AlbumType, not even an explicit nil
func (o *AlbumResource) UnsetAlbumType() {
	o.AlbumType.Unset()
}

// GetSecondaryTypes returns the SecondaryTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetSecondaryTypes() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.SecondaryTypes
}

// GetSecondaryTypesOk returns a tuple with the SecondaryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetSecondaryTypesOk() ([]*string, bool) {
	if o == nil || isNil(o.SecondaryTypes) {
    return nil, false
	}
	return o.SecondaryTypes, true
}

// HasSecondaryTypes returns a boolean if a field has been set.
func (o *AlbumResource) HasSecondaryTypes() bool {
	if o != nil && isNil(o.SecondaryTypes) {
		return true
	}

	return false
}

// SetSecondaryTypes gets a reference to the given []string and assigns it to the SecondaryTypes field.
func (o *AlbumResource) SetSecondaryTypes(v []*string) {
	o.SecondaryTypes = v
}

// GetMediumCount returns the MediumCount field value if set, zero value otherwise.
func (o *AlbumResource) GetMediumCount() int32 {
	if o == nil || isNil(o.MediumCount) {
		var ret int32
		return ret
	}
	return *o.MediumCount
}

// GetMediumCountOk returns a tuple with the MediumCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetMediumCountOk() (*int32, bool) {
	if o == nil || isNil(o.MediumCount) {
    return nil, false
	}
	return o.MediumCount, true
}

// HasMediumCount returns a boolean if a field has been set.
func (o *AlbumResource) HasMediumCount() bool {
	if o != nil && !isNil(o.MediumCount) {
		return true
	}

	return false
}

// SetMediumCount gets a reference to the given int32 and assigns it to the MediumCount field.
func (o *AlbumResource) SetMediumCount(v int32) {
	o.MediumCount = &v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *AlbumResource) GetRatings() Ratings {
	if o == nil || isNil(o.Ratings) {
		var ret Ratings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetRatingsOk() (*Ratings, bool) {
	if o == nil || isNil(o.Ratings) {
    return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *AlbumResource) HasRatings() bool {
	if o != nil && !isNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given Ratings and assigns it to the Ratings field.
func (o *AlbumResource) SetRatings(v Ratings) {
	o.Ratings = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetReleaseDate() time.Time {
	if o == nil || isNil(o.ReleaseDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *AlbumResource) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate.IsSet() {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given NullableTime and assigns it to the ReleaseDate field.
func (o *AlbumResource) SetReleaseDate(v time.Time) {
	o.ReleaseDate.Set(&v)
}
// SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil
func (o *AlbumResource) SetReleaseDateNil() {
	o.ReleaseDate.Set(nil)
}

// UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
func (o *AlbumResource) UnsetReleaseDate() {
	o.ReleaseDate.Unset()
}

// GetReleases returns the Releases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetReleases() []*AlbumReleaseResource {
	if o == nil {
		var ret []*AlbumReleaseResource
		return ret
	}
	return o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetReleasesOk() ([]*AlbumReleaseResource, bool) {
	if o == nil || isNil(o.Releases) {
    return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *AlbumResource) HasReleases() bool {
	if o != nil && isNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given []AlbumReleaseResource and assigns it to the Releases field.
func (o *AlbumResource) SetReleases(v []*AlbumReleaseResource) {
	o.Releases = v
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetGenres() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetGenresOk() ([]*string, bool) {
	if o == nil || isNil(o.Genres) {
    return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *AlbumResource) HasGenres() bool {
	if o != nil && isNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *AlbumResource) SetGenres(v []*string) {
	o.Genres = v
}

// GetMedia returns the Media field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetMedia() []*MediumResource {
	if o == nil {
		var ret []*MediumResource
		return ret
	}
	return o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetMediaOk() ([]*MediumResource, bool) {
	if o == nil || isNil(o.Media) {
    return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *AlbumResource) HasMedia() bool {
	if o != nil && isNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given []MediumResource and assigns it to the Media field.
func (o *AlbumResource) SetMedia(v []*MediumResource) {
	o.Media = v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *AlbumResource) GetArtist() ArtistResource {
	if o == nil || isNil(o.Artist) {
		var ret ArtistResource
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetArtistOk() (*ArtistResource, bool) {
	if o == nil || isNil(o.Artist) {
    return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *AlbumResource) HasArtist() bool {
	if o != nil && !isNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistResource and assigns it to the Artist field.
func (o *AlbumResource) SetArtist(v ArtistResource) {
	o.Artist = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetImages() []*MediaCover {
	if o == nil {
		var ret []*MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetImagesOk() ([]*MediaCover, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *AlbumResource) HasImages() bool {
	if o != nil && isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *AlbumResource) SetImages(v []*MediaCover) {
	o.Images = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetLinks() []*Links {
	if o == nil {
		var ret []*Links
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetLinksOk() ([]*Links, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlbumResource) HasLinks() bool {
	if o != nil && isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Links and assigns it to the Links field.
func (o *AlbumResource) SetLinks(v []*Links) {
	o.Links = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *AlbumResource) GetStatistics() AlbumStatisticsResource {
	if o == nil || isNil(o.Statistics) {
		var ret AlbumStatisticsResource
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetStatisticsOk() (*AlbumStatisticsResource, bool) {
	if o == nil || isNil(o.Statistics) {
    return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *AlbumResource) HasStatistics() bool {
	if o != nil && !isNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given AlbumStatisticsResource and assigns it to the Statistics field.
func (o *AlbumResource) SetStatistics(v AlbumStatisticsResource) {
	o.Statistics = &v
}

// GetAddOptions returns the AddOptions field value if set, zero value otherwise.
func (o *AlbumResource) GetAddOptions() AddAlbumOptions {
	if o == nil || isNil(o.AddOptions) {
		var ret AddAlbumOptions
		return ret
	}
	return *o.AddOptions
}

// GetAddOptionsOk returns a tuple with the AddOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetAddOptionsOk() (*AddAlbumOptions, bool) {
	if o == nil || isNil(o.AddOptions) {
    return nil, false
	}
	return o.AddOptions, true
}

// HasAddOptions returns a boolean if a field has been set.
func (o *AlbumResource) HasAddOptions() bool {
	if o != nil && !isNil(o.AddOptions) {
		return true
	}

	return false
}

// SetAddOptions gets a reference to the given AddAlbumOptions and assigns it to the AddOptions field.
func (o *AlbumResource) SetAddOptions(v AddAlbumOptions) {
	o.AddOptions = &v
}

// GetRemoteCover returns the RemoteCover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResource) GetRemoteCover() string {
	if o == nil || isNil(o.RemoteCover.Get()) {
		var ret string
		return ret
	}
	return *o.RemoteCover.Get()
}

// GetRemoteCoverOk returns a tuple with the RemoteCover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResource) GetRemoteCoverOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RemoteCover.Get(), o.RemoteCover.IsSet()
}

// HasRemoteCover returns a boolean if a field has been set.
func (o *AlbumResource) HasRemoteCover() bool {
	if o != nil && o.RemoteCover.IsSet() {
		return true
	}

	return false
}

// SetRemoteCover gets a reference to the given NullableString and assigns it to the RemoteCover field.
func (o *AlbumResource) SetRemoteCover(v string) {
	o.RemoteCover.Set(&v)
}
// SetRemoteCoverNil sets the value for RemoteCover to be an explicit nil
func (o *AlbumResource) SetRemoteCoverNil() {
	o.RemoteCover.Set(nil)
}

// UnsetRemoteCover ensures that no value is present for RemoteCover, not even an explicit nil
func (o *AlbumResource) UnsetRemoteCover() {
	o.RemoteCover.Unset()
}

// GetGrabbed returns the Grabbed field value if set, zero value otherwise.
func (o *AlbumResource) GetGrabbed() bool {
	if o == nil || isNil(o.Grabbed) {
		var ret bool
		return ret
	}
	return *o.Grabbed
}

// GetGrabbedOk returns a tuple with the Grabbed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResource) GetGrabbedOk() (*bool, bool) {
	if o == nil || isNil(o.Grabbed) {
    return nil, false
	}
	return o.Grabbed, true
}

// HasGrabbed returns a boolean if a field has been set.
func (o *AlbumResource) HasGrabbed() bool {
	if o != nil && !isNil(o.Grabbed) {
		return true
	}

	return false
}

// SetGrabbed gets a reference to the given bool and assigns it to the Grabbed field.
func (o *AlbumResource) SetGrabbed(v bool) {
	o.Grabbed = &v
}

func (o AlbumResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Disambiguation.IsSet() {
		toSerialize["disambiguation"] = o.Disambiguation.Get()
	}
	if o.Overview.IsSet() {
		toSerialize["overview"] = o.Overview.Get()
	}
	if !isNil(o.ArtistId) {
		toSerialize["artistId"] = o.ArtistId
	}
	if o.ForeignAlbumId.IsSet() {
		toSerialize["foreignAlbumId"] = o.ForeignAlbumId.Get()
	}
	if !isNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !isNil(o.AnyReleaseOk) {
		toSerialize["anyReleaseOk"] = o.AnyReleaseOk
	}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if o.AlbumType.IsSet() {
		toSerialize["albumType"] = o.AlbumType.Get()
	}
	if o.SecondaryTypes != nil {
		toSerialize["secondaryTypes"] = o.SecondaryTypes
	}
	if !isNil(o.MediumCount) {
		toSerialize["mediumCount"] = o.MediumCount
	}
	if !isNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if o.ReleaseDate.IsSet() {
		toSerialize["releaseDate"] = o.ReleaseDate.Get()
	}
	if o.Releases != nil {
		toSerialize["releases"] = o.Releases
	}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if o.Media != nil {
		toSerialize["media"] = o.Media
	}
	if !isNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !isNil(o.AddOptions) {
		toSerialize["addOptions"] = o.AddOptions
	}
	if o.RemoteCover.IsSet() {
		toSerialize["remoteCover"] = o.RemoteCover.Get()
	}
	if !isNil(o.Grabbed) {
		toSerialize["grabbed"] = o.Grabbed
	}
	return json.Marshal(toSerialize)
}

type NullableAlbumResource struct {
	value *AlbumResource
	isSet bool
}

func (v NullableAlbumResource) Get() *AlbumResource {
	return v.value
}

func (v *NullableAlbumResource) Set(val *AlbumResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumResource(val *AlbumResource) *NullableAlbumResource {
	return &NullableAlbumResource{value: val, isSet: true}
}

func (v NullableAlbumResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


