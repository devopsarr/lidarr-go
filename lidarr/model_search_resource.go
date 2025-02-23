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

// checks if the SearchResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResource{}

// SearchResource struct for SearchResource
type SearchResource struct {
	Id *int32 `json:"id,omitempty"`
	ForeignId NullableString `json:"foreignId,omitempty"`
	Artist *ArtistResource `json:"artist,omitempty"`
	Album *AlbumResource `json:"album,omitempty"`
}

// NewSearchResource instantiates a new SearchResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResource() *SearchResource {
	this := SearchResource{}
	return &this
}

// NewSearchResourceWithDefaults instantiates a new SearchResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResourceWithDefaults() *SearchResource {
	this := SearchResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SearchResource) SetId(v int32) {
	o.Id = &v
}

// GetForeignId returns the ForeignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResource) GetForeignId() string {
	if o == nil || IsNil(o.ForeignId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignId.Get()
}

// GetForeignIdOk returns a tuple with the ForeignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResource) GetForeignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForeignId.Get(), o.ForeignId.IsSet()
}

// HasForeignId returns a boolean if a field has been set.
func (o *SearchResource) HasForeignId() bool {
	if o != nil && o.ForeignId.IsSet() {
		return true
	}

	return false
}

// SetForeignId gets a reference to the given NullableString and assigns it to the ForeignId field.
func (o *SearchResource) SetForeignId(v string) {
	o.ForeignId.Set(&v)
}
// SetForeignIdNil sets the value for ForeignId to be an explicit nil
func (o *SearchResource) SetForeignIdNil() {
	o.ForeignId.Set(nil)
}

// UnsetForeignId ensures that no value is present for ForeignId, not even an explicit nil
func (o *SearchResource) UnsetForeignId() {
	o.ForeignId.Unset()
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *SearchResource) GetArtist() ArtistResource {
	if o == nil || IsNil(o.Artist) {
		var ret ArtistResource
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResource) GetArtistOk() (*ArtistResource, bool) {
	if o == nil || IsNil(o.Artist) {
		return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *SearchResource) HasArtist() bool {
	if o != nil && !IsNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given ArtistResource and assigns it to the Artist field.
func (o *SearchResource) SetArtist(v ArtistResource) {
	o.Artist = &v
}

// GetAlbum returns the Album field value if set, zero value otherwise.
func (o *SearchResource) GetAlbum() AlbumResource {
	if o == nil || IsNil(o.Album) {
		var ret AlbumResource
		return ret
	}
	return *o.Album
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResource) GetAlbumOk() (*AlbumResource, bool) {
	if o == nil || IsNil(o.Album) {
		return nil, false
	}
	return o.Album, true
}

// HasAlbum returns a boolean if a field has been set.
func (o *SearchResource) HasAlbum() bool {
	if o != nil && !IsNil(o.Album) {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given AlbumResource and assigns it to the Album field.
func (o *SearchResource) SetAlbum(v AlbumResource) {
	o.Album = &v
}

func (o SearchResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ForeignId.IsSet() {
		toSerialize["foreignId"] = o.ForeignId.Get()
	}
	if !IsNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if !IsNil(o.Album) {
		toSerialize["album"] = o.Album
	}
	return toSerialize, nil
}

type NullableSearchResource struct {
	value *SearchResource
	isSet bool
}

func (v NullableSearchResource) Get() *SearchResource {
	return v.value
}

func (v *NullableSearchResource) Set(val *SearchResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResource(val *SearchResource) *NullableSearchResource {
	return &NullableSearchResource{value: val, isSet: true}
}

func (v NullableSearchResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


