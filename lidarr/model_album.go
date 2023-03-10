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

// Album struct for Album
type Album struct {
	Id *int32 `json:"id,omitempty"`
	ArtistMetadataId *int32 `json:"artistMetadataId,omitempty"`
	ForeignAlbumId NullableString `json:"foreignAlbumId,omitempty"`
	OldForeignAlbumIds []*string `json:"oldForeignAlbumIds,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Overview NullableString `json:"overview,omitempty"`
	Disambiguation NullableString `json:"disambiguation,omitempty"`
	ReleaseDate NullableTime `json:"releaseDate,omitempty"`
	Images []*MediaCover `json:"images,omitempty"`
	Links []*Links `json:"links,omitempty"`
	Genres []*string `json:"genres,omitempty"`
	AlbumType NullableString `json:"albumType,omitempty"`
	SecondaryTypes []*SecondaryAlbumType `json:"secondaryTypes,omitempty"`
	Ratings *Ratings `json:"ratings,omitempty"`
	CleanTitle NullableString `json:"cleanTitle,omitempty"`
	ProfileId *int32 `json:"profileId,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	AnyReleaseOk *bool `json:"anyReleaseOk,omitempty"`
	LastInfoSync NullableTime `json:"lastInfoSync,omitempty"`
	Added *time.Time `json:"added,omitempty"`
	AddOptions *AddAlbumOptions `json:"addOptions,omitempty"`
	ArtistMetadata *ArtistMetadataLazyLoaded `json:"artistMetadata,omitempty"`
	AlbumReleases *AlbumReleaseListLazyLoaded `json:"albumReleases,omitempty"`
	Artist *ArtistLazyLoaded `json:"artist,omitempty"`
}

// NewAlbum instantiates a new Album object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbum() *Album {
	this := Album{}
	return &this
}

// NewAlbumWithDefaults instantiates a new Album object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumWithDefaults() *Album {
	this := Album{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Album) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Album) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Album) SetId(v int32) {
	o.Id = &v
}

// GetArtistMetadataId returns the ArtistMetadataId field value if set, zero value otherwise.
func (o *Album) GetArtistMetadataId() int32 {
	if o == nil || isNil(o.ArtistMetadataId) {
		var ret int32
		return ret
	}
	return *o.ArtistMetadataId
}

// GetArtistMetadataIdOk returns a tuple with the ArtistMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetArtistMetadataIdOk() (*int32, bool) {
	if o == nil || isNil(o.ArtistMetadataId) {
    return nil, false
	}
	return o.ArtistMetadataId, true
}

// HasArtistMetadataId returns a boolean if a field has been set.
func (o *Album) HasArtistMetadataId() bool {
	if o != nil && !isNil(o.ArtistMetadataId) {
		return true
	}

	return false
}

// SetArtistMetadataId gets a reference to the given int32 and assigns it to the ArtistMetadataId field.
func (o *Album) SetArtistMetadataId(v int32) {
	o.ArtistMetadataId = &v
}

// GetForeignAlbumId returns the ForeignAlbumId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetForeignAlbumId() string {
	if o == nil || isNil(o.ForeignAlbumId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignAlbumId.Get()
}

// GetForeignAlbumIdOk returns a tuple with the ForeignAlbumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetForeignAlbumIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignAlbumId.Get(), o.ForeignAlbumId.IsSet()
}

// HasForeignAlbumId returns a boolean if a field has been set.
func (o *Album) HasForeignAlbumId() bool {
	if o != nil && o.ForeignAlbumId.IsSet() {
		return true
	}

	return false
}

// SetForeignAlbumId gets a reference to the given NullableString and assigns it to the ForeignAlbumId field.
func (o *Album) SetForeignAlbumId(v string) {
	o.ForeignAlbumId.Set(&v)
}
// SetForeignAlbumIdNil sets the value for ForeignAlbumId to be an explicit nil
func (o *Album) SetForeignAlbumIdNil() {
	o.ForeignAlbumId.Set(nil)
}

// UnsetForeignAlbumId ensures that no value is present for ForeignAlbumId, not even an explicit nil
func (o *Album) UnsetForeignAlbumId() {
	o.ForeignAlbumId.Unset()
}

// GetOldForeignAlbumIds returns the OldForeignAlbumIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetOldForeignAlbumIds() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.OldForeignAlbumIds
}

// GetOldForeignAlbumIdsOk returns a tuple with the OldForeignAlbumIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetOldForeignAlbumIdsOk() ([]*string, bool) {
	if o == nil || isNil(o.OldForeignAlbumIds) {
    return nil, false
	}
	return o.OldForeignAlbumIds, true
}

// HasOldForeignAlbumIds returns a boolean if a field has been set.
func (o *Album) HasOldForeignAlbumIds() bool {
	if o != nil && isNil(o.OldForeignAlbumIds) {
		return true
	}

	return false
}

// SetOldForeignAlbumIds gets a reference to the given []string and assigns it to the OldForeignAlbumIds field.
func (o *Album) SetOldForeignAlbumIds(v []*string) {
	o.OldForeignAlbumIds = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Album) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Album) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Album) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Album) UnsetTitle() {
	o.Title.Unset()
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetOverview() string {
	if o == nil || isNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetOverviewOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *Album) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *Album) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *Album) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *Album) UnsetOverview() {
	o.Overview.Unset()
}

// GetDisambiguation returns the Disambiguation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetDisambiguation() string {
	if o == nil || isNil(o.Disambiguation.Get()) {
		var ret string
		return ret
	}
	return *o.Disambiguation.Get()
}

// GetDisambiguationOk returns a tuple with the Disambiguation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetDisambiguationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Disambiguation.Get(), o.Disambiguation.IsSet()
}

// HasDisambiguation returns a boolean if a field has been set.
func (o *Album) HasDisambiguation() bool {
	if o != nil && o.Disambiguation.IsSet() {
		return true
	}

	return false
}

// SetDisambiguation gets a reference to the given NullableString and assigns it to the Disambiguation field.
func (o *Album) SetDisambiguation(v string) {
	o.Disambiguation.Set(&v)
}
// SetDisambiguationNil sets the value for Disambiguation to be an explicit nil
func (o *Album) SetDisambiguationNil() {
	o.Disambiguation.Set(nil)
}

// UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
func (o *Album) UnsetDisambiguation() {
	o.Disambiguation.Unset()
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetReleaseDate() time.Time {
	if o == nil || isNil(o.ReleaseDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *Album) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate.IsSet() {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given NullableTime and assigns it to the ReleaseDate field.
func (o *Album) SetReleaseDate(v time.Time) {
	o.ReleaseDate.Set(&v)
}
// SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil
func (o *Album) SetReleaseDateNil() {
	o.ReleaseDate.Set(nil)
}

// UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
func (o *Album) UnsetReleaseDate() {
	o.ReleaseDate.Unset()
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetImages() []*MediaCover {
	if o == nil {
		var ret []*MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetImagesOk() ([]*MediaCover, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Album) HasImages() bool {
	if o != nil && isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *Album) SetImages(v []*MediaCover) {
	o.Images = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetLinks() []*Links {
	if o == nil {
		var ret []*Links
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetLinksOk() ([]*Links, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Album) HasLinks() bool {
	if o != nil && isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Links and assigns it to the Links field.
func (o *Album) SetLinks(v []*Links) {
	o.Links = v
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetGenres() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetGenresOk() ([]*string, bool) {
	if o == nil || isNil(o.Genres) {
    return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *Album) HasGenres() bool {
	if o != nil && isNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *Album) SetGenres(v []*string) {
	o.Genres = v
}

// GetAlbumType returns the AlbumType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetAlbumType() string {
	if o == nil || isNil(o.AlbumType.Get()) {
		var ret string
		return ret
	}
	return *o.AlbumType.Get()
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetAlbumTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AlbumType.Get(), o.AlbumType.IsSet()
}

// HasAlbumType returns a boolean if a field has been set.
func (o *Album) HasAlbumType() bool {
	if o != nil && o.AlbumType.IsSet() {
		return true
	}

	return false
}

// SetAlbumType gets a reference to the given NullableString and assigns it to the AlbumType field.
func (o *Album) SetAlbumType(v string) {
	o.AlbumType.Set(&v)
}
// SetAlbumTypeNil sets the value for AlbumType to be an explicit nil
func (o *Album) SetAlbumTypeNil() {
	o.AlbumType.Set(nil)
}

// UnsetAlbumType ensures that no value is present for AlbumType, not even an explicit nil
func (o *Album) UnsetAlbumType() {
	o.AlbumType.Unset()
}

// GetSecondaryTypes returns the SecondaryTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetSecondaryTypes() []*SecondaryAlbumType {
	if o == nil {
		var ret []*SecondaryAlbumType
		return ret
	}
	return o.SecondaryTypes
}

// GetSecondaryTypesOk returns a tuple with the SecondaryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetSecondaryTypesOk() ([]*SecondaryAlbumType, bool) {
	if o == nil || isNil(o.SecondaryTypes) {
    return nil, false
	}
	return o.SecondaryTypes, true
}

// HasSecondaryTypes returns a boolean if a field has been set.
func (o *Album) HasSecondaryTypes() bool {
	if o != nil && isNil(o.SecondaryTypes) {
		return true
	}

	return false
}

// SetSecondaryTypes gets a reference to the given []SecondaryAlbumType and assigns it to the SecondaryTypes field.
func (o *Album) SetSecondaryTypes(v []*SecondaryAlbumType) {
	o.SecondaryTypes = v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *Album) GetRatings() Ratings {
	if o == nil || isNil(o.Ratings) {
		var ret Ratings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetRatingsOk() (*Ratings, bool) {
	if o == nil || isNil(o.Ratings) {
    return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *Album) HasRatings() bool {
	if o != nil && !isNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given Ratings and assigns it to the Ratings field.
func (o *Album) SetRatings(v Ratings) {
	o.Ratings = &v
}

// GetCleanTitle returns the CleanTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetCleanTitle() string {
	if o == nil || isNil(o.CleanTitle.Get()) {
		var ret string
		return ret
	}
	return *o.CleanTitle.Get()
}

// GetCleanTitleOk returns a tuple with the CleanTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetCleanTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CleanTitle.Get(), o.CleanTitle.IsSet()
}

// HasCleanTitle returns a boolean if a field has been set.
func (o *Album) HasCleanTitle() bool {
	if o != nil && o.CleanTitle.IsSet() {
		return true
	}

	return false
}

// SetCleanTitle gets a reference to the given NullableString and assigns it to the CleanTitle field.
func (o *Album) SetCleanTitle(v string) {
	o.CleanTitle.Set(&v)
}
// SetCleanTitleNil sets the value for CleanTitle to be an explicit nil
func (o *Album) SetCleanTitleNil() {
	o.CleanTitle.Set(nil)
}

// UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
func (o *Album) UnsetCleanTitle() {
	o.CleanTitle.Unset()
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *Album) GetProfileId() int32 {
	if o == nil || isNil(o.ProfileId) {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetProfileIdOk() (*int32, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *Album) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *Album) SetProfileId(v int32) {
	o.ProfileId = &v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *Album) GetMonitored() bool {
	if o == nil || isNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetMonitoredOk() (*bool, bool) {
	if o == nil || isNil(o.Monitored) {
    return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *Album) HasMonitored() bool {
	if o != nil && !isNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *Album) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetAnyReleaseOk returns the AnyReleaseOk field value if set, zero value otherwise.
func (o *Album) GetAnyReleaseOk() bool {
	if o == nil || isNil(o.AnyReleaseOk) {
		var ret bool
		return ret
	}
	return *o.AnyReleaseOk
}

// GetAnyReleaseOkOk returns a tuple with the AnyReleaseOk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetAnyReleaseOkOk() (*bool, bool) {
	if o == nil || isNil(o.AnyReleaseOk) {
    return nil, false
	}
	return o.AnyReleaseOk, true
}

// HasAnyReleaseOk returns a boolean if a field has been set.
func (o *Album) HasAnyReleaseOk() bool {
	if o != nil && !isNil(o.AnyReleaseOk) {
		return true
	}

	return false
}

// SetAnyReleaseOk gets a reference to the given bool and assigns it to the AnyReleaseOk field.
func (o *Album) SetAnyReleaseOk(v bool) {
	o.AnyReleaseOk = &v
}

// GetLastInfoSync returns the LastInfoSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Album) GetLastInfoSync() time.Time {
	if o == nil || isNil(o.LastInfoSync.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastInfoSync.Get()
}

// GetLastInfoSyncOk returns a tuple with the LastInfoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Album) GetLastInfoSyncOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastInfoSync.Get(), o.LastInfoSync.IsSet()
}

// HasLastInfoSync returns a boolean if a field has been set.
func (o *Album) HasLastInfoSync() bool {
	if o != nil && o.LastInfoSync.IsSet() {
		return true
	}

	return false
}

// SetLastInfoSync gets a reference to the given NullableTime and assigns it to the LastInfoSync field.
func (o *Album) SetLastInfoSync(v time.Time) {
	o.LastInfoSync.Set(&v)
}
// SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil
func (o *Album) SetLastInfoSyncNil() {
	o.LastInfoSync.Set(nil)
}

// UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
func (o *Album) UnsetLastInfoSync() {
	o.LastInfoSync.Unset()
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *Album) GetAdded() time.Time {
	if o == nil || isNil(o.Added) {
		var ret time.Time
		return ret
	}
	return *o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Added) {
    return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *Album) HasAdded() bool {
	if o != nil && !isNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given time.Time and assigns it to the Added field.
func (o *Album) SetAdded(v time.Time) {
	o.Added = &v
}

// GetAddOptions returns the AddOptions field value if set, zero value otherwise.
func (o *Album) GetAddOptions() AddAlbumOptions {
	if o == nil || isNil(o.AddOptions) {
		var ret AddAlbumOptions
		return ret
	}
	return *o.AddOptions
}

// GetAddOptionsOk returns a tuple with the AddOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetAddOptionsOk() (*AddAlbumOptions, bool) {
	if o == nil || isNil(o.AddOptions) {
    return nil, false
	}
	return o.AddOptions, true
}

// HasAddOptions returns a boolean if a field has been set.
func (o *Album) HasAddOptions() bool {
	if o != nil && !isNil(o.AddOptions) {
		return true
	}

	return false
}

// SetAddOptions gets a reference to the given AddAlbumOptions and assigns it to the AddOptions field.
func (o *Album) SetAddOptions(v AddAlbumOptions) {
	o.AddOptions = &v
}

// GetArtistMetadata returns the ArtistMetadata field value if set, zero value otherwise.
func (o *Album) GetArtistMetadata() ArtistMetadataLazyLoaded {
	if o == nil || isNil(o.ArtistMetadata) {
		var ret ArtistMetadataLazyLoaded
		return ret
	}
	return *o.ArtistMetadata
}

// GetArtistMetadataOk returns a tuple with the ArtistMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetArtistMetadataOk() (*ArtistMetadataLazyLoaded, bool) {
	if o == nil || isNil(o.ArtistMetadata) {
    return nil, false
	}
	return o.ArtistMetadata, true
}

// HasArtistMetadata returns a boolean if a field has been set.
func (o *Album) HasArtistMetadata() bool {
	if o != nil && !isNil(o.ArtistMetadata) {
		return true
	}

	return false
}

// SetArtistMetadata gets a reference to the given ArtistMetadataLazyLoaded and assigns it to the ArtistMetadata field.
func (o *Album) SetArtistMetadata(v ArtistMetadataLazyLoaded) {
	o.ArtistMetadata = &v
}

// GetAlbumReleases returns the AlbumReleases field value if set, zero value otherwise.
func (o *Album) GetAlbumReleases() AlbumReleaseListLazyLoaded {
	if o == nil || isNil(o.AlbumReleases) {
		var ret AlbumReleaseListLazyLoaded
		return ret
	}
	return *o.AlbumReleases
}

// GetAlbumReleasesOk returns a tuple with the AlbumReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetAlbumReleasesOk() (*AlbumReleaseListLazyLoaded, bool) {
	if o == nil || isNil(o.AlbumReleases) {
    return nil, false
	}
	return o.AlbumReleases, true
}

// HasAlbumReleases returns a boolean if a field has been set.
func (o *Album) HasAlbumReleases() bool {
	if o != nil && !isNil(o.AlbumReleases) {
		return true
	}

	return false
}

// SetAlbumReleases gets a reference to the given AlbumReleaseListLazyLoaded and assigns it to the AlbumReleases field.
func (o *Album) SetAlbumReleases(v AlbumReleaseListLazyLoaded) {
	o.AlbumReleases = &v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *Album) GetArtist() ArtistLazyLoaded {
	if o == nil || isNil(o.Artist) {
		var ret ArtistLazyLoaded
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Album) GetArtistOk() (*ArtistLazyLoaded, bool) {
	if o == nil || isNil(o.Artist) {
    return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *Album) HasArtist() bool {
	if o != nil && !isNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistLazyLoaded and assigns it to the Artist field.
func (o *Album) SetArtist(v ArtistLazyLoaded) {
	o.Artist = &v
}

func (o Album) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ArtistMetadataId) {
		toSerialize["artistMetadataId"] = o.ArtistMetadataId
	}
	if o.ForeignAlbumId.IsSet() {
		toSerialize["foreignAlbumId"] = o.ForeignAlbumId.Get()
	}
	if o.OldForeignAlbumIds != nil {
		toSerialize["oldForeignAlbumIds"] = o.OldForeignAlbumIds
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Overview.IsSet() {
		toSerialize["overview"] = o.Overview.Get()
	}
	if o.Disambiguation.IsSet() {
		toSerialize["disambiguation"] = o.Disambiguation.Get()
	}
	if o.ReleaseDate.IsSet() {
		toSerialize["releaseDate"] = o.ReleaseDate.Get()
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if o.AlbumType.IsSet() {
		toSerialize["albumType"] = o.AlbumType.Get()
	}
	if o.SecondaryTypes != nil {
		toSerialize["secondaryTypes"] = o.SecondaryTypes
	}
	if !isNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if o.CleanTitle.IsSet() {
		toSerialize["cleanTitle"] = o.CleanTitle.Get()
	}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !isNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !isNil(o.AnyReleaseOk) {
		toSerialize["anyReleaseOk"] = o.AnyReleaseOk
	}
	if o.LastInfoSync.IsSet() {
		toSerialize["lastInfoSync"] = o.LastInfoSync.Get()
	}
	if !isNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !isNil(o.AddOptions) {
		toSerialize["addOptions"] = o.AddOptions
	}
	if !isNil(o.ArtistMetadata) {
		toSerialize["artistMetadata"] = o.ArtistMetadata
	}
	if !isNil(o.AlbumReleases) {
		toSerialize["albumReleases"] = o.AlbumReleases
	}
	if !isNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	return json.Marshal(toSerialize)
}

type NullableAlbum struct {
	value *Album
	isSet bool
}

func (v NullableAlbum) Get() *Album {
	return v.value
}

func (v *NullableAlbum) Set(val *Album) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbum) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbum(val *Album) *NullableAlbum {
	return &NullableAlbum{value: val, isSet: true}
}

func (v NullableAlbum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


