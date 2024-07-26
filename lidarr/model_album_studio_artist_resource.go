/*
Lidarr

Lidarr API docs

API version: v2.4.3.4248
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the AlbumStudioArtistResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumStudioArtistResource{}

// AlbumStudioArtistResource struct for AlbumStudioArtistResource
type AlbumStudioArtistResource struct {
	Id *int32 `json:"id,omitempty"`
	Monitored NullableBool `json:"monitored,omitempty"`
	Albums []AlbumResource `json:"albums,omitempty"`
}

// NewAlbumStudioArtistResource instantiates a new AlbumStudioArtistResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumStudioArtistResource() *AlbumStudioArtistResource {
	this := AlbumStudioArtistResource{}
	return &this
}

// NewAlbumStudioArtistResourceWithDefaults instantiates a new AlbumStudioArtistResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumStudioArtistResourceWithDefaults() *AlbumStudioArtistResource {
	this := AlbumStudioArtistResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlbumStudioArtistResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumStudioArtistResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlbumStudioArtistResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlbumStudioArtistResource) SetId(v int32) {
	o.Id = &v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumStudioArtistResource) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored.Get()) {
		var ret bool
		return ret
	}
	return *o.Monitored.Get()
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumStudioArtistResource) GetMonitoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Monitored.Get(), o.Monitored.IsSet()
}

// HasMonitored returns a boolean if a field has been set.
func (o *AlbumStudioArtistResource) HasMonitored() bool {
	if o != nil && o.Monitored.IsSet() {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given NullableBool and assigns it to the Monitored field.
func (o *AlbumStudioArtistResource) SetMonitored(v bool) {
	o.Monitored.Set(&v)
}
// SetMonitoredNil sets the value for Monitored to be an explicit nil
func (o *AlbumStudioArtistResource) SetMonitoredNil() {
	o.Monitored.Set(nil)
}

// UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
func (o *AlbumStudioArtistResource) UnsetMonitored() {
	o.Monitored.Unset()
}

// GetAlbums returns the Albums field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumStudioArtistResource) GetAlbums() []AlbumResource {
	if o == nil {
		var ret []AlbumResource
		return ret
	}
	return o.Albums
}

// GetAlbumsOk returns a tuple with the Albums field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumStudioArtistResource) GetAlbumsOk() ([]AlbumResource, bool) {
	if o == nil || IsNil(o.Albums) {
		return nil, false
	}
	return o.Albums, true
}

// HasAlbums returns a boolean if a field has been set.
func (o *AlbumStudioArtistResource) HasAlbums() bool {
	if o != nil && !IsNil(o.Albums) {
		return true
	}

	return false
}

// SetAlbums gets a reference to the given []AlbumResource and assigns it to the Albums field.
func (o *AlbumStudioArtistResource) SetAlbums(v []AlbumResource) {
	o.Albums = v
}

func (o AlbumStudioArtistResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumStudioArtistResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Monitored.IsSet() {
		toSerialize["monitored"] = o.Monitored.Get()
	}
	if o.Albums != nil {
		toSerialize["albums"] = o.Albums
	}
	return toSerialize, nil
}

type NullableAlbumStudioArtistResource struct {
	value *AlbumStudioArtistResource
	isSet bool
}

func (v NullableAlbumStudioArtistResource) Get() *AlbumStudioArtistResource {
	return v.value
}

func (v *NullableAlbumStudioArtistResource) Set(val *AlbumStudioArtistResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumStudioArtistResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumStudioArtistResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumStudioArtistResource(val *AlbumStudioArtistResource) *NullableAlbumStudioArtistResource {
	return &NullableAlbumStudioArtistResource{value: val, isSet: true}
}

func (v NullableAlbumStudioArtistResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumStudioArtistResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


