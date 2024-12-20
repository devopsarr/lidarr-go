/*
Lidarr

Lidarr API docs

API version: v2.8.2.4493
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the ArtistStatisticsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtistStatisticsResource{}

// ArtistStatisticsResource struct for ArtistStatisticsResource
type ArtistStatisticsResource struct {
	AlbumCount *int32 `json:"albumCount,omitempty"`
	TrackFileCount *int32 `json:"trackFileCount,omitempty"`
	TrackCount *int32 `json:"trackCount,omitempty"`
	TotalTrackCount *int32 `json:"totalTrackCount,omitempty"`
	SizeOnDisk *int64 `json:"sizeOnDisk,omitempty"`
	PercentOfTracks *float64 `json:"percentOfTracks,omitempty"`
}

// NewArtistStatisticsResource instantiates a new ArtistStatisticsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtistStatisticsResource() *ArtistStatisticsResource {
	this := ArtistStatisticsResource{}
	return &this
}

// NewArtistStatisticsResourceWithDefaults instantiates a new ArtistStatisticsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtistStatisticsResourceWithDefaults() *ArtistStatisticsResource {
	this := ArtistStatisticsResource{}
	return &this
}

// GetAlbumCount returns the AlbumCount field value if set, zero value otherwise.
func (o *ArtistStatisticsResource) GetAlbumCount() int32 {
	if o == nil || IsNil(o.AlbumCount) {
		var ret int32
		return ret
	}
	return *o.AlbumCount
}

// GetAlbumCountOk returns a tuple with the AlbumCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistStatisticsResource) GetAlbumCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AlbumCount) {
		return nil, false
	}
	return o.AlbumCount, true
}

// HasAlbumCount returns a boolean if a field has been set.
func (o *ArtistStatisticsResource) HasAlbumCount() bool {
	if o != nil && !IsNil(o.AlbumCount) {
		return true
	}

	return false
}

// SetAlbumCount gets a reference to the given int32 and assigns it to the AlbumCount field.
func (o *ArtistStatisticsResource) SetAlbumCount(v int32) {
	o.AlbumCount = &v
}

// GetTrackFileCount returns the TrackFileCount field value if set, zero value otherwise.
func (o *ArtistStatisticsResource) GetTrackFileCount() int32 {
	if o == nil || IsNil(o.TrackFileCount) {
		var ret int32
		return ret
	}
	return *o.TrackFileCount
}

// GetTrackFileCountOk returns a tuple with the TrackFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistStatisticsResource) GetTrackFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TrackFileCount) {
		return nil, false
	}
	return o.TrackFileCount, true
}

// HasTrackFileCount returns a boolean if a field has been set.
func (o *ArtistStatisticsResource) HasTrackFileCount() bool {
	if o != nil && !IsNil(o.TrackFileCount) {
		return true
	}

	return false
}

// SetTrackFileCount gets a reference to the given int32 and assigns it to the TrackFileCount field.
func (o *ArtistStatisticsResource) SetTrackFileCount(v int32) {
	o.TrackFileCount = &v
}

// GetTrackCount returns the TrackCount field value if set, zero value otherwise.
func (o *ArtistStatisticsResource) GetTrackCount() int32 {
	if o == nil || IsNil(o.TrackCount) {
		var ret int32
		return ret
	}
	return *o.TrackCount
}

// GetTrackCountOk returns a tuple with the TrackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistStatisticsResource) GetTrackCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TrackCount) {
		return nil, false
	}
	return o.TrackCount, true
}

// HasTrackCount returns a boolean if a field has been set.
func (o *ArtistStatisticsResource) HasTrackCount() bool {
	if o != nil && !IsNil(o.TrackCount) {
		return true
	}

	return false
}

// SetTrackCount gets a reference to the given int32 and assigns it to the TrackCount field.
func (o *ArtistStatisticsResource) SetTrackCount(v int32) {
	o.TrackCount = &v
}

// GetTotalTrackCount returns the TotalTrackCount field value if set, zero value otherwise.
func (o *ArtistStatisticsResource) GetTotalTrackCount() int32 {
	if o == nil || IsNil(o.TotalTrackCount) {
		var ret int32
		return ret
	}
	return *o.TotalTrackCount
}

// GetTotalTrackCountOk returns a tuple with the TotalTrackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistStatisticsResource) GetTotalTrackCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalTrackCount) {
		return nil, false
	}
	return o.TotalTrackCount, true
}

// HasTotalTrackCount returns a boolean if a field has been set.
func (o *ArtistStatisticsResource) HasTotalTrackCount() bool {
	if o != nil && !IsNil(o.TotalTrackCount) {
		return true
	}

	return false
}

// SetTotalTrackCount gets a reference to the given int32 and assigns it to the TotalTrackCount field.
func (o *ArtistStatisticsResource) SetTotalTrackCount(v int32) {
	o.TotalTrackCount = &v
}

// GetSizeOnDisk returns the SizeOnDisk field value if set, zero value otherwise.
func (o *ArtistStatisticsResource) GetSizeOnDisk() int64 {
	if o == nil || IsNil(o.SizeOnDisk) {
		var ret int64
		return ret
	}
	return *o.SizeOnDisk
}

// GetSizeOnDiskOk returns a tuple with the SizeOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistStatisticsResource) GetSizeOnDiskOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeOnDisk) {
		return nil, false
	}
	return o.SizeOnDisk, true
}

// HasSizeOnDisk returns a boolean if a field has been set.
func (o *ArtistStatisticsResource) HasSizeOnDisk() bool {
	if o != nil && !IsNil(o.SizeOnDisk) {
		return true
	}

	return false
}

// SetSizeOnDisk gets a reference to the given int64 and assigns it to the SizeOnDisk field.
func (o *ArtistStatisticsResource) SetSizeOnDisk(v int64) {
	o.SizeOnDisk = &v
}

// GetPercentOfTracks returns the PercentOfTracks field value if set, zero value otherwise.
func (o *ArtistStatisticsResource) GetPercentOfTracks() float64 {
	if o == nil || IsNil(o.PercentOfTracks) {
		var ret float64
		return ret
	}
	return *o.PercentOfTracks
}

// GetPercentOfTracksOk returns a tuple with the PercentOfTracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistStatisticsResource) GetPercentOfTracksOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentOfTracks) {
		return nil, false
	}
	return o.PercentOfTracks, true
}

// HasPercentOfTracks returns a boolean if a field has been set.
func (o *ArtistStatisticsResource) HasPercentOfTracks() bool {
	if o != nil && !IsNil(o.PercentOfTracks) {
		return true
	}

	return false
}

// SetPercentOfTracks gets a reference to the given float64 and assigns it to the PercentOfTracks field.
func (o *ArtistStatisticsResource) SetPercentOfTracks(v float64) {
	o.PercentOfTracks = &v
}

func (o ArtistStatisticsResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtistStatisticsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlbumCount) {
		toSerialize["albumCount"] = o.AlbumCount
	}
	if !IsNil(o.TrackFileCount) {
		toSerialize["trackFileCount"] = o.TrackFileCount
	}
	if !IsNil(o.TrackCount) {
		toSerialize["trackCount"] = o.TrackCount
	}
	if !IsNil(o.TotalTrackCount) {
		toSerialize["totalTrackCount"] = o.TotalTrackCount
	}
	if !IsNil(o.SizeOnDisk) {
		toSerialize["sizeOnDisk"] = o.SizeOnDisk
	}
	if !IsNil(o.PercentOfTracks) {
		toSerialize["percentOfTracks"] = o.PercentOfTracks
	}
	return toSerialize, nil
}

type NullableArtistStatisticsResource struct {
	value *ArtistStatisticsResource
	isSet bool
}

func (v NullableArtistStatisticsResource) Get() *ArtistStatisticsResource {
	return v.value
}

func (v *NullableArtistStatisticsResource) Set(val *ArtistStatisticsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableArtistStatisticsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableArtistStatisticsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtistStatisticsResource(val *ArtistStatisticsResource) *NullableArtistStatisticsResource {
	return &NullableArtistStatisticsResource{value: val, isSet: true}
}

func (v NullableArtistStatisticsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtistStatisticsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


