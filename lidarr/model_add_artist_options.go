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

// checks if the AddArtistOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddArtistOptions{}

// AddArtistOptions struct for AddArtistOptions
type AddArtistOptions struct {
	Monitor *MonitorTypes `json:"monitor,omitempty"`
	AlbumsToMonitor []string `json:"albumsToMonitor,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	SearchForMissingAlbums *bool `json:"searchForMissingAlbums,omitempty"`
}

// NewAddArtistOptions instantiates a new AddArtistOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddArtistOptions() *AddArtistOptions {
	this := AddArtistOptions{}
	return &this
}

// NewAddArtistOptionsWithDefaults instantiates a new AddArtistOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddArtistOptionsWithDefaults() *AddArtistOptions {
	this := AddArtistOptions{}
	return &this
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *AddArtistOptions) GetMonitor() MonitorTypes {
	if o == nil || IsNil(o.Monitor) {
		var ret MonitorTypes
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddArtistOptions) GetMonitorOk() (*MonitorTypes, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *AddArtistOptions) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorTypes and assigns it to the Monitor field.
func (o *AddArtistOptions) SetMonitor(v MonitorTypes) {
	o.Monitor = &v
}

// GetAlbumsToMonitor returns the AlbumsToMonitor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddArtistOptions) GetAlbumsToMonitor() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AlbumsToMonitor
}

// GetAlbumsToMonitorOk returns a tuple with the AlbumsToMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddArtistOptions) GetAlbumsToMonitorOk() ([]string, bool) {
	if o == nil || IsNil(o.AlbumsToMonitor) {
		return nil, false
	}
	return o.AlbumsToMonitor, true
}

// HasAlbumsToMonitor returns a boolean if a field has been set.
func (o *AddArtistOptions) HasAlbumsToMonitor() bool {
	if o != nil && !IsNil(o.AlbumsToMonitor) {
		return true
	}

	return false
}

// SetAlbumsToMonitor gets a reference to the given []string and assigns it to the AlbumsToMonitor field.
func (o *AddArtistOptions) SetAlbumsToMonitor(v []string) {
	o.AlbumsToMonitor = v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *AddArtistOptions) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddArtistOptions) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *AddArtistOptions) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *AddArtistOptions) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetSearchForMissingAlbums returns the SearchForMissingAlbums field value if set, zero value otherwise.
func (o *AddArtistOptions) GetSearchForMissingAlbums() bool {
	if o == nil || IsNil(o.SearchForMissingAlbums) {
		var ret bool
		return ret
	}
	return *o.SearchForMissingAlbums
}

// GetSearchForMissingAlbumsOk returns a tuple with the SearchForMissingAlbums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddArtistOptions) GetSearchForMissingAlbumsOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchForMissingAlbums) {
		return nil, false
	}
	return o.SearchForMissingAlbums, true
}

// HasSearchForMissingAlbums returns a boolean if a field has been set.
func (o *AddArtistOptions) HasSearchForMissingAlbums() bool {
	if o != nil && !IsNil(o.SearchForMissingAlbums) {
		return true
	}

	return false
}

// SetSearchForMissingAlbums gets a reference to the given bool and assigns it to the SearchForMissingAlbums field.
func (o *AddArtistOptions) SetSearchForMissingAlbums(v bool) {
	o.SearchForMissingAlbums = &v
}

func (o AddArtistOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddArtistOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if o.AlbumsToMonitor != nil {
		toSerialize["albumsToMonitor"] = o.AlbumsToMonitor
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !IsNil(o.SearchForMissingAlbums) {
		toSerialize["searchForMissingAlbums"] = o.SearchForMissingAlbums
	}
	return toSerialize, nil
}

type NullableAddArtistOptions struct {
	value *AddArtistOptions
	isSet bool
}

func (v NullableAddArtistOptions) Get() *AddArtistOptions {
	return v.value
}

func (v *NullableAddArtistOptions) Set(val *AddArtistOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAddArtistOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAddArtistOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddArtistOptions(val *AddArtistOptions) *NullableAddArtistOptions {
	return &NullableAddArtistOptions{value: val, isSet: true}
}

func (v NullableAddArtistOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddArtistOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


