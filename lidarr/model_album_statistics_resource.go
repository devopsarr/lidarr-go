/*
Lidarr

Lidarr API docs

API version: v2.1.7.4030
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// AlbumStatisticsResource struct for AlbumStatisticsResource
type AlbumStatisticsResource struct {
	TrackFileCount *int32 `json:"trackFileCount,omitempty"`
	TrackCount *int32 `json:"trackCount,omitempty"`
	TotalTrackCount *int32 `json:"totalTrackCount,omitempty"`
	SizeOnDisk *int64 `json:"sizeOnDisk,omitempty"`
	PercentOfTracks *float64 `json:"percentOfTracks,omitempty"`
}

// NewAlbumStatisticsResource instantiates a new AlbumStatisticsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumStatisticsResource() *AlbumStatisticsResource {
	this := AlbumStatisticsResource{}
	return &this
}

// NewAlbumStatisticsResourceWithDefaults instantiates a new AlbumStatisticsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumStatisticsResourceWithDefaults() *AlbumStatisticsResource {
	this := AlbumStatisticsResource{}
	return &this
}

// GetTrackFileCount returns the TrackFileCount field value if set, zero value otherwise.
func (o *AlbumStatisticsResource) GetTrackFileCount() int32 {
	if o == nil || isNil(o.TrackFileCount) {
		var ret int32
		return ret
	}
	return *o.TrackFileCount
}

// GetTrackFileCountOk returns a tuple with the TrackFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStatisticsResource) GetTrackFileCountOk() (*int32, bool) {
	if o == nil || isNil(o.TrackFileCount) {
    return nil, false
	}
	return o.TrackFileCount, true
}

// HasTrackFileCount returns a boolean if a field has been set.
func (o *AlbumStatisticsResource) HasTrackFileCount() bool {
	if o != nil && !isNil(o.TrackFileCount) {
		return true
	}

	return false
}

// SetTrackFileCount gets a reference to the given int32 and assigns it to the TrackFileCount field.
func (o *AlbumStatisticsResource) SetTrackFileCount(v int32) {
	o.TrackFileCount = &v
}

// GetTrackCount returns the TrackCount field value if set, zero value otherwise.
func (o *AlbumStatisticsResource) GetTrackCount() int32 {
	if o == nil || isNil(o.TrackCount) {
		var ret int32
		return ret
	}
	return *o.TrackCount
}

// GetTrackCountOk returns a tuple with the TrackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStatisticsResource) GetTrackCountOk() (*int32, bool) {
	if o == nil || isNil(o.TrackCount) {
    return nil, false
	}
	return o.TrackCount, true
}

// HasTrackCount returns a boolean if a field has been set.
func (o *AlbumStatisticsResource) HasTrackCount() bool {
	if o != nil && !isNil(o.TrackCount) {
		return true
	}

	return false
}

// SetTrackCount gets a reference to the given int32 and assigns it to the TrackCount field.
func (o *AlbumStatisticsResource) SetTrackCount(v int32) {
	o.TrackCount = &v
}

// GetTotalTrackCount returns the TotalTrackCount field value if set, zero value otherwise.
func (o *AlbumStatisticsResource) GetTotalTrackCount() int32 {
	if o == nil || isNil(o.TotalTrackCount) {
		var ret int32
		return ret
	}
	return *o.TotalTrackCount
}

// GetTotalTrackCountOk returns a tuple with the TotalTrackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStatisticsResource) GetTotalTrackCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalTrackCount) {
    return nil, false
	}
	return o.TotalTrackCount, true
}

// HasTotalTrackCount returns a boolean if a field has been set.
func (o *AlbumStatisticsResource) HasTotalTrackCount() bool {
	if o != nil && !isNil(o.TotalTrackCount) {
		return true
	}

	return false
}

// SetTotalTrackCount gets a reference to the given int32 and assigns it to the TotalTrackCount field.
func (o *AlbumStatisticsResource) SetTotalTrackCount(v int32) {
	o.TotalTrackCount = &v
}

// GetSizeOnDisk returns the SizeOnDisk field value if set, zero value otherwise.
func (o *AlbumStatisticsResource) GetSizeOnDisk() int64 {
	if o == nil || isNil(o.SizeOnDisk) {
		var ret int64
		return ret
	}
	return *o.SizeOnDisk
}

// GetSizeOnDiskOk returns a tuple with the SizeOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStatisticsResource) GetSizeOnDiskOk() (*int64, bool) {
	if o == nil || isNil(o.SizeOnDisk) {
    return nil, false
	}
	return o.SizeOnDisk, true
}

// HasSizeOnDisk returns a boolean if a field has been set.
func (o *AlbumStatisticsResource) HasSizeOnDisk() bool {
	if o != nil && !isNil(o.SizeOnDisk) {
		return true
	}

	return false
}

// SetSizeOnDisk gets a reference to the given int64 and assigns it to the SizeOnDisk field.
func (o *AlbumStatisticsResource) SetSizeOnDisk(v int64) {
	o.SizeOnDisk = &v
}

// GetPercentOfTracks returns the PercentOfTracks field value if set, zero value otherwise.
func (o *AlbumStatisticsResource) GetPercentOfTracks() float64 {
	if o == nil || isNil(o.PercentOfTracks) {
		var ret float64
		return ret
	}
	return *o.PercentOfTracks
}

// GetPercentOfTracksOk returns a tuple with the PercentOfTracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStatisticsResource) GetPercentOfTracksOk() (*float64, bool) {
	if o == nil || isNil(o.PercentOfTracks) {
    return nil, false
	}
	return o.PercentOfTracks, true
}

// HasPercentOfTracks returns a boolean if a field has been set.
func (o *AlbumStatisticsResource) HasPercentOfTracks() bool {
	if o != nil && !isNil(o.PercentOfTracks) {
		return true
	}

	return false
}

// SetPercentOfTracks gets a reference to the given float64 and assigns it to the PercentOfTracks field.
func (o *AlbumStatisticsResource) SetPercentOfTracks(v float64) {
	o.PercentOfTracks = &v
}

func (o AlbumStatisticsResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TrackFileCount) {
		toSerialize["trackFileCount"] = o.TrackFileCount
	}
	if !isNil(o.TrackCount) {
		toSerialize["trackCount"] = o.TrackCount
	}
	if !isNil(o.TotalTrackCount) {
		toSerialize["totalTrackCount"] = o.TotalTrackCount
	}
	if !isNil(o.SizeOnDisk) {
		toSerialize["sizeOnDisk"] = o.SizeOnDisk
	}
	if !isNil(o.PercentOfTracks) {
		toSerialize["percentOfTracks"] = o.PercentOfTracks
	}
	return json.Marshal(toSerialize)
}

type NullableAlbumStatisticsResource struct {
	value *AlbumStatisticsResource
	isSet bool
}

func (v NullableAlbumStatisticsResource) Get() *AlbumStatisticsResource {
	return v.value
}

func (v *NullableAlbumStatisticsResource) Set(val *AlbumStatisticsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumStatisticsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumStatisticsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumStatisticsResource(val *AlbumStatisticsResource) *NullableAlbumStatisticsResource {
	return &NullableAlbumStatisticsResource{value: val, isSet: true}
}

func (v NullableAlbumStatisticsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumStatisticsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


