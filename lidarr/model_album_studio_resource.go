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

// checks if the AlbumStudioResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumStudioResource{}

// AlbumStudioResource struct for AlbumStudioResource
type AlbumStudioResource struct {
	Artist []AlbumStudioArtistResource `json:"artist,omitempty"`
	MonitoringOptions *MonitoringOptions `json:"monitoringOptions,omitempty"`
	MonitorNewItems *NewItemMonitorTypes `json:"monitorNewItems,omitempty"`
}

// NewAlbumStudioResource instantiates a new AlbumStudioResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumStudioResource() *AlbumStudioResource {
	this := AlbumStudioResource{}
	return &this
}

// NewAlbumStudioResourceWithDefaults instantiates a new AlbumStudioResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumStudioResourceWithDefaults() *AlbumStudioResource {
	this := AlbumStudioResource{}
	return &this
}

// GetArtist returns the Artist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumStudioResource) GetArtist() []AlbumStudioArtistResource {
	if o == nil {
		var ret []AlbumStudioArtistResource
		return ret
	}
	return o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumStudioResource) GetArtistOk() ([]AlbumStudioArtistResource, bool) {
	if o == nil || IsNil(o.Artist) {
		return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *AlbumStudioResource) HasArtist() bool {
	if o != nil && !IsNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given []AlbumStudioArtistResource and assigns it to the Artist field.
func (o *AlbumStudioResource) SetArtist(v []AlbumStudioArtistResource) {
	o.Artist = v
}

// GetMonitoringOptions returns the MonitoringOptions field value if set, zero value otherwise.
func (o *AlbumStudioResource) GetMonitoringOptions() MonitoringOptions {
	if o == nil || IsNil(o.MonitoringOptions) {
		var ret MonitoringOptions
		return ret
	}
	return *o.MonitoringOptions
}

// GetMonitoringOptionsOk returns a tuple with the MonitoringOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStudioResource) GetMonitoringOptionsOk() (*MonitoringOptions, bool) {
	if o == nil || IsNil(o.MonitoringOptions) {
		return nil, false
	}
	return o.MonitoringOptions, true
}

// HasMonitoringOptions returns a boolean if a field has been set.
func (o *AlbumStudioResource) HasMonitoringOptions() bool {
	if o != nil && !IsNil(o.MonitoringOptions) {
		return true
	}

	return false
}

// SetMonitoringOptions gets a reference to the given MonitoringOptions and assigns it to the MonitoringOptions field.
func (o *AlbumStudioResource) SetMonitoringOptions(v MonitoringOptions) {
	o.MonitoringOptions = &v
}

// GetMonitorNewItems returns the MonitorNewItems field value if set, zero value otherwise.
func (o *AlbumStudioResource) GetMonitorNewItems() NewItemMonitorTypes {
	if o == nil || IsNil(o.MonitorNewItems) {
		var ret NewItemMonitorTypes
		return ret
	}
	return *o.MonitorNewItems
}

// GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStudioResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool) {
	if o == nil || IsNil(o.MonitorNewItems) {
		return nil, false
	}
	return o.MonitorNewItems, true
}

// HasMonitorNewItems returns a boolean if a field has been set.
func (o *AlbumStudioResource) HasMonitorNewItems() bool {
	if o != nil && !IsNil(o.MonitorNewItems) {
		return true
	}

	return false
}

// SetMonitorNewItems gets a reference to the given NewItemMonitorTypes and assigns it to the MonitorNewItems field.
func (o *AlbumStudioResource) SetMonitorNewItems(v NewItemMonitorTypes) {
	o.MonitorNewItems = &v
}

func (o AlbumStudioResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumStudioResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Artist != nil {
		toSerialize["artist"] = o.Artist
	}
	if !IsNil(o.MonitoringOptions) {
		toSerialize["monitoringOptions"] = o.MonitoringOptions
	}
	if !IsNil(o.MonitorNewItems) {
		toSerialize["monitorNewItems"] = o.MonitorNewItems
	}
	return toSerialize, nil
}

type NullableAlbumStudioResource struct {
	value *AlbumStudioResource
	isSet bool
}

func (v NullableAlbumStudioResource) Get() *AlbumStudioResource {
	return v.value
}

func (v *NullableAlbumStudioResource) Set(val *AlbumStudioResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumStudioResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumStudioResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumStudioResource(val *AlbumStudioResource) *NullableAlbumStudioResource {
	return &NullableAlbumStudioResource{value: val, isSet: true}
}

func (v NullableAlbumStudioResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumStudioResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


