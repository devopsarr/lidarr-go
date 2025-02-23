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

// checks if the MetadataProviderConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataProviderConfigResource{}

// MetadataProviderConfigResource struct for MetadataProviderConfigResource
type MetadataProviderConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	MetadataSource NullableString `json:"metadataSource,omitempty"`
	WriteAudioTags *WriteAudioTagsType `json:"writeAudioTags,omitempty"`
	ScrubAudioTags *bool `json:"scrubAudioTags,omitempty"`
	EmbedCoverArt *bool `json:"embedCoverArt,omitempty"`
}

// NewMetadataProviderConfigResource instantiates a new MetadataProviderConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataProviderConfigResource() *MetadataProviderConfigResource {
	this := MetadataProviderConfigResource{}
	return &this
}

// NewMetadataProviderConfigResourceWithDefaults instantiates a new MetadataProviderConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataProviderConfigResourceWithDefaults() *MetadataProviderConfigResource {
	this := MetadataProviderConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataProviderConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetMetadataSource returns the MetadataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataProviderConfigResource) GetMetadataSource() string {
	if o == nil || IsNil(o.MetadataSource.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataSource.Get()
}

// GetMetadataSourceOk returns a tuple with the MetadataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataProviderConfigResource) GetMetadataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataSource.Get(), o.MetadataSource.IsSet()
}

// HasMetadataSource returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasMetadataSource() bool {
	if o != nil && o.MetadataSource.IsSet() {
		return true
	}

	return false
}

// SetMetadataSource gets a reference to the given NullableString and assigns it to the MetadataSource field.
func (o *MetadataProviderConfigResource) SetMetadataSource(v string) {
	o.MetadataSource.Set(&v)
}
// SetMetadataSourceNil sets the value for MetadataSource to be an explicit nil
func (o *MetadataProviderConfigResource) SetMetadataSourceNil() {
	o.MetadataSource.Set(nil)
}

// UnsetMetadataSource ensures that no value is present for MetadataSource, not even an explicit nil
func (o *MetadataProviderConfigResource) UnsetMetadataSource() {
	o.MetadataSource.Unset()
}

// GetWriteAudioTags returns the WriteAudioTags field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetWriteAudioTags() WriteAudioTagsType {
	if o == nil || IsNil(o.WriteAudioTags) {
		var ret WriteAudioTagsType
		return ret
	}
	return *o.WriteAudioTags
}

// GetWriteAudioTagsOk returns a tuple with the WriteAudioTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetWriteAudioTagsOk() (*WriteAudioTagsType, bool) {
	if o == nil || IsNil(o.WriteAudioTags) {
		return nil, false
	}
	return o.WriteAudioTags, true
}

// HasWriteAudioTags returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasWriteAudioTags() bool {
	if o != nil && !IsNil(o.WriteAudioTags) {
		return true
	}

	return false
}

// SetWriteAudioTags gets a reference to the given WriteAudioTagsType and assigns it to the WriteAudioTags field.
func (o *MetadataProviderConfigResource) SetWriteAudioTags(v WriteAudioTagsType) {
	o.WriteAudioTags = &v
}

// GetScrubAudioTags returns the ScrubAudioTags field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetScrubAudioTags() bool {
	if o == nil || IsNil(o.ScrubAudioTags) {
		var ret bool
		return ret
	}
	return *o.ScrubAudioTags
}

// GetScrubAudioTagsOk returns a tuple with the ScrubAudioTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetScrubAudioTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.ScrubAudioTags) {
		return nil, false
	}
	return o.ScrubAudioTags, true
}

// HasScrubAudioTags returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasScrubAudioTags() bool {
	if o != nil && !IsNil(o.ScrubAudioTags) {
		return true
	}

	return false
}

// SetScrubAudioTags gets a reference to the given bool and assigns it to the ScrubAudioTags field.
func (o *MetadataProviderConfigResource) SetScrubAudioTags(v bool) {
	o.ScrubAudioTags = &v
}

// GetEmbedCoverArt returns the EmbedCoverArt field value if set, zero value otherwise.
func (o *MetadataProviderConfigResource) GetEmbedCoverArt() bool {
	if o == nil || IsNil(o.EmbedCoverArt) {
		var ret bool
		return ret
	}
	return *o.EmbedCoverArt
}

// GetEmbedCoverArtOk returns a tuple with the EmbedCoverArt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataProviderConfigResource) GetEmbedCoverArtOk() (*bool, bool) {
	if o == nil || IsNil(o.EmbedCoverArt) {
		return nil, false
	}
	return o.EmbedCoverArt, true
}

// HasEmbedCoverArt returns a boolean if a field has been set.
func (o *MetadataProviderConfigResource) HasEmbedCoverArt() bool {
	if o != nil && !IsNil(o.EmbedCoverArt) {
		return true
	}

	return false
}

// SetEmbedCoverArt gets a reference to the given bool and assigns it to the EmbedCoverArt field.
func (o *MetadataProviderConfigResource) SetEmbedCoverArt(v bool) {
	o.EmbedCoverArt = &v
}

func (o MetadataProviderConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataProviderConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.MetadataSource.IsSet() {
		toSerialize["metadataSource"] = o.MetadataSource.Get()
	}
	if !IsNil(o.WriteAudioTags) {
		toSerialize["writeAudioTags"] = o.WriteAudioTags
	}
	if !IsNil(o.ScrubAudioTags) {
		toSerialize["scrubAudioTags"] = o.ScrubAudioTags
	}
	if !IsNil(o.EmbedCoverArt) {
		toSerialize["embedCoverArt"] = o.EmbedCoverArt
	}
	return toSerialize, nil
}

type NullableMetadataProviderConfigResource struct {
	value *MetadataProviderConfigResource
	isSet bool
}

func (v NullableMetadataProviderConfigResource) Get() *MetadataProviderConfigResource {
	return v.value
}

func (v *NullableMetadataProviderConfigResource) Set(val *MetadataProviderConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProviderConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProviderConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProviderConfigResource(val *MetadataProviderConfigResource) *NullableMetadataProviderConfigResource {
	return &NullableMetadataProviderConfigResource{value: val, isSet: true}
}

func (v NullableMetadataProviderConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProviderConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


