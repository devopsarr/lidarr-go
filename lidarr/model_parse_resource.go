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

// checks if the ParseResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParseResource{}

// ParseResource struct for ParseResource
type ParseResource struct {
	Id *int32 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	ParsedAlbumInfo *ParsedAlbumInfo `json:"parsedAlbumInfo,omitempty"`
	Artist *ArtistResource `json:"artist,omitempty"`
	Albums []AlbumResource `json:"albums,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
}

// NewParseResource instantiates a new ParseResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseResource() *ParseResource {
	this := ParseResource{}
	return &this
}

// NewParseResourceWithDefaults instantiates a new ParseResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseResourceWithDefaults() *ParseResource {
	this := ParseResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParseResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParseResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ParseResource) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseResource) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseResource) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ParseResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ParseResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ParseResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ParseResource) UnsetTitle() {
	o.Title.Unset()
}

// GetParsedAlbumInfo returns the ParsedAlbumInfo field value if set, zero value otherwise.
func (o *ParseResource) GetParsedAlbumInfo() ParsedAlbumInfo {
	if o == nil || IsNil(o.ParsedAlbumInfo) {
		var ret ParsedAlbumInfo
		return ret
	}
	return *o.ParsedAlbumInfo
}

// GetParsedAlbumInfoOk returns a tuple with the ParsedAlbumInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetParsedAlbumInfoOk() (*ParsedAlbumInfo, bool) {
	if o == nil || IsNil(o.ParsedAlbumInfo) {
		return nil, false
	}
	return o.ParsedAlbumInfo, true
}

// HasParsedAlbumInfo returns a boolean if a field has been set.
func (o *ParseResource) HasParsedAlbumInfo() bool {
	if o != nil && !IsNil(o.ParsedAlbumInfo) {
		return true
	}

	return false
}

// SetParsedAlbumInfo gets a reference to the given ParsedAlbumInfo and assigns it to the ParsedAlbumInfo field.
func (o *ParseResource) SetParsedAlbumInfo(v ParsedAlbumInfo) {
	o.ParsedAlbumInfo = &v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *ParseResource) GetArtist() ArtistResource {
	if o == nil || IsNil(o.Artist) {
		var ret ArtistResource
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetArtistOk() (*ArtistResource, bool) {
	if o == nil || IsNil(o.Artist) {
		return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *ParseResource) HasArtist() bool {
	if o != nil && !IsNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistResource and assigns it to the Artist field.
func (o *ParseResource) SetArtist(v ArtistResource) {
	o.Artist = &v
}

// GetAlbums returns the Albums field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseResource) GetAlbums() []AlbumResource {
	if o == nil {
		var ret []AlbumResource
		return ret
	}
	return o.Albums
}

// GetAlbumsOk returns a tuple with the Albums field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseResource) GetAlbumsOk() ([]AlbumResource, bool) {
	if o == nil || IsNil(o.Albums) {
		return nil, false
	}
	return o.Albums, true
}

// HasAlbums returns a boolean if a field has been set.
func (o *ParseResource) HasAlbums() bool {
	if o != nil && !IsNil(o.Albums) {
		return true
	}

	return false
}

// SetAlbums gets a reference to the given []AlbumResource and assigns it to the Albums field.
func (o *ParseResource) SetAlbums(v []AlbumResource) {
	o.Albums = v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *ParseResource) HasCustomFormats() bool {
	if o != nil && !IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *ParseResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *ParseResource) GetCustomFormatScore() int32 {
	if o == nil || IsNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomFormatScore) {
		return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *ParseResource) HasCustomFormatScore() bool {
	if o != nil && !IsNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *ParseResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

func (o ParseResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParseResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.ParsedAlbumInfo) {
		toSerialize["parsedAlbumInfo"] = o.ParsedAlbumInfo
	}
	if !IsNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if o.Albums != nil {
		toSerialize["albums"] = o.Albums
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !IsNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	return toSerialize, nil
}

type NullableParseResource struct {
	value *ParseResource
	isSet bool
}

func (v NullableParseResource) Get() *ParseResource {
	return v.value
}

func (v *NullableParseResource) Set(val *ParseResource) {
	v.value = val
	v.isSet = true
}

func (v NullableParseResource) IsSet() bool {
	return v.isSet
}

func (v *NullableParseResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseResource(val *ParseResource) *NullableParseResource {
	return &NullableParseResource{value: val, isSet: true}
}

func (v NullableParseResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


