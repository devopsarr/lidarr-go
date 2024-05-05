/*
Lidarr

Lidarr API docs

API version: v2.2.5.4141
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the ParsedAlbumInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParsedAlbumInfo{}

// ParsedAlbumInfo struct for ParsedAlbumInfo
type ParsedAlbumInfo struct {
	AlbumTitle NullableString `json:"albumTitle,omitempty"`
	ArtistName NullableString `json:"artistName,omitempty"`
	AlbumType NullableString `json:"albumType,omitempty"`
	ArtistTitleInfo *ArtistTitleInfo `json:"artistTitleInfo,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	ReleaseDate NullableString `json:"releaseDate,omitempty"`
	Discography *bool `json:"discography,omitempty"`
	DiscographyStart *int32 `json:"discographyStart,omitempty"`
	DiscographyEnd *int32 `json:"discographyEnd,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	ReleaseHash NullableString `json:"releaseHash,omitempty"`
	ReleaseVersion NullableString `json:"releaseVersion,omitempty"`
	ReleaseTitle NullableString `json:"releaseTitle,omitempty"`
}

// NewParsedAlbumInfo instantiates a new ParsedAlbumInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsedAlbumInfo() *ParsedAlbumInfo {
	this := ParsedAlbumInfo{}
	return &this
}

// NewParsedAlbumInfoWithDefaults instantiates a new ParsedAlbumInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsedAlbumInfoWithDefaults() *ParsedAlbumInfo {
	this := ParsedAlbumInfo{}
	return &this
}

// GetAlbumTitle returns the AlbumTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetAlbumTitle() string {
	if o == nil || IsNil(o.AlbumTitle.Get()) {
		var ret string
		return ret
	}
	return *o.AlbumTitle.Get()
}

// GetAlbumTitleOk returns a tuple with the AlbumTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetAlbumTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlbumTitle.Get(), o.AlbumTitle.IsSet()
}

// HasAlbumTitle returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasAlbumTitle() bool {
	if o != nil && o.AlbumTitle.IsSet() {
		return true
	}

	return false
}

// SetAlbumTitle gets a reference to the given NullableString and assigns it to the AlbumTitle field.
func (o *ParsedAlbumInfo) SetAlbumTitle(v string) {
	o.AlbumTitle.Set(&v)
}
// SetAlbumTitleNil sets the value for AlbumTitle to be an explicit nil
func (o *ParsedAlbumInfo) SetAlbumTitleNil() {
	o.AlbumTitle.Set(nil)
}

// UnsetAlbumTitle ensures that no value is present for AlbumTitle, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetAlbumTitle() {
	o.AlbumTitle.Unset()
}

// GetArtistName returns the ArtistName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetArtistName() string {
	if o == nil || IsNil(o.ArtistName.Get()) {
		var ret string
		return ret
	}
	return *o.ArtistName.Get()
}

// GetArtistNameOk returns a tuple with the ArtistName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetArtistNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArtistName.Get(), o.ArtistName.IsSet()
}

// HasArtistName returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasArtistName() bool {
	if o != nil && o.ArtistName.IsSet() {
		return true
	}

	return false
}

// SetArtistName gets a reference to the given NullableString and assigns it to the ArtistName field.
func (o *ParsedAlbumInfo) SetArtistName(v string) {
	o.ArtistName.Set(&v)
}
// SetArtistNameNil sets the value for ArtistName to be an explicit nil
func (o *ParsedAlbumInfo) SetArtistNameNil() {
	o.ArtistName.Set(nil)
}

// UnsetArtistName ensures that no value is present for ArtistName, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetArtistName() {
	o.ArtistName.Unset()
}

// GetAlbumType returns the AlbumType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetAlbumType() string {
	if o == nil || IsNil(o.AlbumType.Get()) {
		var ret string
		return ret
	}
	return *o.AlbumType.Get()
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetAlbumTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlbumType.Get(), o.AlbumType.IsSet()
}

// HasAlbumType returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasAlbumType() bool {
	if o != nil && o.AlbumType.IsSet() {
		return true
	}

	return false
}

// SetAlbumType gets a reference to the given NullableString and assigns it to the AlbumType field.
func (o *ParsedAlbumInfo) SetAlbumType(v string) {
	o.AlbumType.Set(&v)
}
// SetAlbumTypeNil sets the value for AlbumType to be an explicit nil
func (o *ParsedAlbumInfo) SetAlbumTypeNil() {
	o.AlbumType.Set(nil)
}

// UnsetAlbumType ensures that no value is present for AlbumType, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetAlbumType() {
	o.AlbumType.Unset()
}

// GetArtistTitleInfo returns the ArtistTitleInfo field value if set, zero value otherwise.
func (o *ParsedAlbumInfo) GetArtistTitleInfo() ArtistTitleInfo {
	if o == nil || IsNil(o.ArtistTitleInfo) {
		var ret ArtistTitleInfo
		return ret
	}
	return *o.ArtistTitleInfo
}

// GetArtistTitleInfoOk returns a tuple with the ArtistTitleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedAlbumInfo) GetArtistTitleInfoOk() (*ArtistTitleInfo, bool) {
	if o == nil || IsNil(o.ArtistTitleInfo) {
		return nil, false
	}
	return o.ArtistTitleInfo, true
}

// HasArtistTitleInfo returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasArtistTitleInfo() bool {
	if o != nil && !IsNil(o.ArtistTitleInfo) {
		return true
	}

	return false
}

// SetArtistTitleInfo gets a reference to the given ArtistTitleInfo and assigns it to the ArtistTitleInfo field.
func (o *ParsedAlbumInfo) SetArtistTitleInfo(v ArtistTitleInfo) {
	o.ArtistTitleInfo = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ParsedAlbumInfo) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedAlbumInfo) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ParsedAlbumInfo) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate.IsSet() {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given NullableString and assigns it to the ReleaseDate field.
func (o *ParsedAlbumInfo) SetReleaseDate(v string) {
	o.ReleaseDate.Set(&v)
}
// SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil
func (o *ParsedAlbumInfo) SetReleaseDateNil() {
	o.ReleaseDate.Set(nil)
}

// UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetReleaseDate() {
	o.ReleaseDate.Unset()
}

// GetDiscography returns the Discography field value if set, zero value otherwise.
func (o *ParsedAlbumInfo) GetDiscography() bool {
	if o == nil || IsNil(o.Discography) {
		var ret bool
		return ret
	}
	return *o.Discography
}

// GetDiscographyOk returns a tuple with the Discography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedAlbumInfo) GetDiscographyOk() (*bool, bool) {
	if o == nil || IsNil(o.Discography) {
		return nil, false
	}
	return o.Discography, true
}

// HasDiscography returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasDiscography() bool {
	if o != nil && !IsNil(o.Discography) {
		return true
	}

	return false
}

// SetDiscography gets a reference to the given bool and assigns it to the Discography field.
func (o *ParsedAlbumInfo) SetDiscography(v bool) {
	o.Discography = &v
}

// GetDiscographyStart returns the DiscographyStart field value if set, zero value otherwise.
func (o *ParsedAlbumInfo) GetDiscographyStart() int32 {
	if o == nil || IsNil(o.DiscographyStart) {
		var ret int32
		return ret
	}
	return *o.DiscographyStart
}

// GetDiscographyStartOk returns a tuple with the DiscographyStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedAlbumInfo) GetDiscographyStartOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscographyStart) {
		return nil, false
	}
	return o.DiscographyStart, true
}

// HasDiscographyStart returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasDiscographyStart() bool {
	if o != nil && !IsNil(o.DiscographyStart) {
		return true
	}

	return false
}

// SetDiscographyStart gets a reference to the given int32 and assigns it to the DiscographyStart field.
func (o *ParsedAlbumInfo) SetDiscographyStart(v int32) {
	o.DiscographyStart = &v
}

// GetDiscographyEnd returns the DiscographyEnd field value if set, zero value otherwise.
func (o *ParsedAlbumInfo) GetDiscographyEnd() int32 {
	if o == nil || IsNil(o.DiscographyEnd) {
		var ret int32
		return ret
	}
	return *o.DiscographyEnd
}

// GetDiscographyEndOk returns a tuple with the DiscographyEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedAlbumInfo) GetDiscographyEndOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscographyEnd) {
		return nil, false
	}
	return o.DiscographyEnd, true
}

// HasDiscographyEnd returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasDiscographyEnd() bool {
	if o != nil && !IsNil(o.DiscographyEnd) {
		return true
	}

	return false
}

// SetDiscographyEnd gets a reference to the given int32 and assigns it to the DiscographyEnd field.
func (o *ParsedAlbumInfo) SetDiscographyEnd(v int32) {
	o.DiscographyEnd = &v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ParsedAlbumInfo) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ParsedAlbumInfo) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetReleaseHash returns the ReleaseHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetReleaseHash() string {
	if o == nil || IsNil(o.ReleaseHash.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseHash.Get()
}

// GetReleaseHashOk returns a tuple with the ReleaseHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetReleaseHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseHash.Get(), o.ReleaseHash.IsSet()
}

// HasReleaseHash returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasReleaseHash() bool {
	if o != nil && o.ReleaseHash.IsSet() {
		return true
	}

	return false
}

// SetReleaseHash gets a reference to the given NullableString and assigns it to the ReleaseHash field.
func (o *ParsedAlbumInfo) SetReleaseHash(v string) {
	o.ReleaseHash.Set(&v)
}
// SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil
func (o *ParsedAlbumInfo) SetReleaseHashNil() {
	o.ReleaseHash.Set(nil)
}

// UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetReleaseHash() {
	o.ReleaseHash.Unset()
}

// GetReleaseVersion returns the ReleaseVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetReleaseVersion() string {
	if o == nil || IsNil(o.ReleaseVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseVersion.Get()
}

// GetReleaseVersionOk returns a tuple with the ReleaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetReleaseVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseVersion.Get(), o.ReleaseVersion.IsSet()
}

// HasReleaseVersion returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasReleaseVersion() bool {
	if o != nil && o.ReleaseVersion.IsSet() {
		return true
	}

	return false
}

// SetReleaseVersion gets a reference to the given NullableString and assigns it to the ReleaseVersion field.
func (o *ParsedAlbumInfo) SetReleaseVersion(v string) {
	o.ReleaseVersion.Set(&v)
}
// SetReleaseVersionNil sets the value for ReleaseVersion to be an explicit nil
func (o *ParsedAlbumInfo) SetReleaseVersionNil() {
	o.ReleaseVersion.Set(nil)
}

// UnsetReleaseVersion ensures that no value is present for ReleaseVersion, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetReleaseVersion() {
	o.ReleaseVersion.Unset()
}

// GetReleaseTitle returns the ReleaseTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedAlbumInfo) GetReleaseTitle() string {
	if o == nil || IsNil(o.ReleaseTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseTitle.Get()
}

// GetReleaseTitleOk returns a tuple with the ReleaseTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedAlbumInfo) GetReleaseTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseTitle.Get(), o.ReleaseTitle.IsSet()
}

// HasReleaseTitle returns a boolean if a field has been set.
func (o *ParsedAlbumInfo) HasReleaseTitle() bool {
	if o != nil && o.ReleaseTitle.IsSet() {
		return true
	}

	return false
}

// SetReleaseTitle gets a reference to the given NullableString and assigns it to the ReleaseTitle field.
func (o *ParsedAlbumInfo) SetReleaseTitle(v string) {
	o.ReleaseTitle.Set(&v)
}
// SetReleaseTitleNil sets the value for ReleaseTitle to be an explicit nil
func (o *ParsedAlbumInfo) SetReleaseTitleNil() {
	o.ReleaseTitle.Set(nil)
}

// UnsetReleaseTitle ensures that no value is present for ReleaseTitle, not even an explicit nil
func (o *ParsedAlbumInfo) UnsetReleaseTitle() {
	o.ReleaseTitle.Unset()
}

func (o ParsedAlbumInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParsedAlbumInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AlbumTitle.IsSet() {
		toSerialize["albumTitle"] = o.AlbumTitle.Get()
	}
	if o.ArtistName.IsSet() {
		toSerialize["artistName"] = o.ArtistName.Get()
	}
	if o.AlbumType.IsSet() {
		toSerialize["albumType"] = o.AlbumType.Get()
	}
	if !IsNil(o.ArtistTitleInfo) {
		toSerialize["artistTitleInfo"] = o.ArtistTitleInfo
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.ReleaseDate.IsSet() {
		toSerialize["releaseDate"] = o.ReleaseDate.Get()
	}
	if !IsNil(o.Discography) {
		toSerialize["discography"] = o.Discography
	}
	if !IsNil(o.DiscographyStart) {
		toSerialize["discographyStart"] = o.DiscographyStart
	}
	if !IsNil(o.DiscographyEnd) {
		toSerialize["discographyEnd"] = o.DiscographyEnd
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.ReleaseHash.IsSet() {
		toSerialize["releaseHash"] = o.ReleaseHash.Get()
	}
	if o.ReleaseVersion.IsSet() {
		toSerialize["releaseVersion"] = o.ReleaseVersion.Get()
	}
	if o.ReleaseTitle.IsSet() {
		toSerialize["releaseTitle"] = o.ReleaseTitle.Get()
	}
	return toSerialize, nil
}

type NullableParsedAlbumInfo struct {
	value *ParsedAlbumInfo
	isSet bool
}

func (v NullableParsedAlbumInfo) Get() *ParsedAlbumInfo {
	return v.value
}

func (v *NullableParsedAlbumInfo) Set(val *ParsedAlbumInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableParsedAlbumInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableParsedAlbumInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsedAlbumInfo(val *ParsedAlbumInfo) *NullableParsedAlbumInfo {
	return &NullableParsedAlbumInfo{value: val, isSet: true}
}

func (v NullableParsedAlbumInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsedAlbumInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


