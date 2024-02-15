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

// checks if the NamingConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamingConfigResource{}

// NamingConfigResource struct for NamingConfigResource
type NamingConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	RenameTracks *bool `json:"renameTracks,omitempty"`
	ReplaceIllegalCharacters *bool `json:"replaceIllegalCharacters,omitempty"`
	ColonReplacementFormat *int32 `json:"colonReplacementFormat,omitempty"`
	StandardTrackFormat NullableString `json:"standardTrackFormat,omitempty"`
	MultiDiscTrackFormat NullableString `json:"multiDiscTrackFormat,omitempty"`
	ArtistFolderFormat NullableString `json:"artistFolderFormat,omitempty"`
	IncludeArtistName *bool `json:"includeArtistName,omitempty"`
	IncludeAlbumTitle *bool `json:"includeAlbumTitle,omitempty"`
	IncludeQuality *bool `json:"includeQuality,omitempty"`
	ReplaceSpaces *bool `json:"replaceSpaces,omitempty"`
	Separator NullableString `json:"separator,omitempty"`
	NumberStyle NullableString `json:"numberStyle,omitempty"`
}

// NewNamingConfigResource instantiates a new NamingConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamingConfigResource() *NamingConfigResource {
	this := NamingConfigResource{}
	return &this
}

// NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamingConfigResourceWithDefaults() *NamingConfigResource {
	this := NamingConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NamingConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamingConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NamingConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetRenameTracks returns the RenameTracks field value if set, zero value otherwise.
func (o *NamingConfigResource) GetRenameTracks() bool {
	if o == nil || IsNil(o.RenameTracks) {
		var ret bool
		return ret
	}
	return *o.RenameTracks
}

// GetRenameTracksOk returns a tuple with the RenameTracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetRenameTracksOk() (*bool, bool) {
	if o == nil || IsNil(o.RenameTracks) {
		return nil, false
	}
	return o.RenameTracks, true
}

// HasRenameTracks returns a boolean if a field has been set.
func (o *NamingConfigResource) HasRenameTracks() bool {
	if o != nil && !IsNil(o.RenameTracks) {
		return true
	}

	return false
}

// SetRenameTracks gets a reference to the given bool and assigns it to the RenameTracks field.
func (o *NamingConfigResource) SetRenameTracks(v bool) {
	o.RenameTracks = &v
}

// GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field value if set, zero value otherwise.
func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool {
	if o == nil || IsNil(o.ReplaceIllegalCharacters) {
		var ret bool
		return ret
	}
	return *o.ReplaceIllegalCharacters
}

// GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceIllegalCharacters) {
		return nil, false
	}
	return o.ReplaceIllegalCharacters, true
}

// HasReplaceIllegalCharacters returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool {
	if o != nil && !IsNil(o.ReplaceIllegalCharacters) {
		return true
	}

	return false
}

// SetReplaceIllegalCharacters gets a reference to the given bool and assigns it to the ReplaceIllegalCharacters field.
func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool) {
	o.ReplaceIllegalCharacters = &v
}

// GetColonReplacementFormat returns the ColonReplacementFormat field value if set, zero value otherwise.
func (o *NamingConfigResource) GetColonReplacementFormat() int32 {
	if o == nil || IsNil(o.ColonReplacementFormat) {
		var ret int32
		return ret
	}
	return *o.ColonReplacementFormat
}

// GetColonReplacementFormatOk returns a tuple with the ColonReplacementFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetColonReplacementFormatOk() (*int32, bool) {
	if o == nil || IsNil(o.ColonReplacementFormat) {
		return nil, false
	}
	return o.ColonReplacementFormat, true
}

// HasColonReplacementFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasColonReplacementFormat() bool {
	if o != nil && !IsNil(o.ColonReplacementFormat) {
		return true
	}

	return false
}

// SetColonReplacementFormat gets a reference to the given int32 and assigns it to the ColonReplacementFormat field.
func (o *NamingConfigResource) SetColonReplacementFormat(v int32) {
	o.ColonReplacementFormat = &v
}

// GetStandardTrackFormat returns the StandardTrackFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetStandardTrackFormat() string {
	if o == nil || IsNil(o.StandardTrackFormat.Get()) {
		var ret string
		return ret
	}
	return *o.StandardTrackFormat.Get()
}

// GetStandardTrackFormatOk returns a tuple with the StandardTrackFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetStandardTrackFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StandardTrackFormat.Get(), o.StandardTrackFormat.IsSet()
}

// HasStandardTrackFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasStandardTrackFormat() bool {
	if o != nil && o.StandardTrackFormat.IsSet() {
		return true
	}

	return false
}

// SetStandardTrackFormat gets a reference to the given NullableString and assigns it to the StandardTrackFormat field.
func (o *NamingConfigResource) SetStandardTrackFormat(v string) {
	o.StandardTrackFormat.Set(&v)
}
// SetStandardTrackFormatNil sets the value for StandardTrackFormat to be an explicit nil
func (o *NamingConfigResource) SetStandardTrackFormatNil() {
	o.StandardTrackFormat.Set(nil)
}

// UnsetStandardTrackFormat ensures that no value is present for StandardTrackFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetStandardTrackFormat() {
	o.StandardTrackFormat.Unset()
}

// GetMultiDiscTrackFormat returns the MultiDiscTrackFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetMultiDiscTrackFormat() string {
	if o == nil || IsNil(o.MultiDiscTrackFormat.Get()) {
		var ret string
		return ret
	}
	return *o.MultiDiscTrackFormat.Get()
}

// GetMultiDiscTrackFormatOk returns a tuple with the MultiDiscTrackFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetMultiDiscTrackFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultiDiscTrackFormat.Get(), o.MultiDiscTrackFormat.IsSet()
}

// HasMultiDiscTrackFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasMultiDiscTrackFormat() bool {
	if o != nil && o.MultiDiscTrackFormat.IsSet() {
		return true
	}

	return false
}

// SetMultiDiscTrackFormat gets a reference to the given NullableString and assigns it to the MultiDiscTrackFormat field.
func (o *NamingConfigResource) SetMultiDiscTrackFormat(v string) {
	o.MultiDiscTrackFormat.Set(&v)
}
// SetMultiDiscTrackFormatNil sets the value for MultiDiscTrackFormat to be an explicit nil
func (o *NamingConfigResource) SetMultiDiscTrackFormatNil() {
	o.MultiDiscTrackFormat.Set(nil)
}

// UnsetMultiDiscTrackFormat ensures that no value is present for MultiDiscTrackFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetMultiDiscTrackFormat() {
	o.MultiDiscTrackFormat.Unset()
}

// GetArtistFolderFormat returns the ArtistFolderFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetArtistFolderFormat() string {
	if o == nil || IsNil(o.ArtistFolderFormat.Get()) {
		var ret string
		return ret
	}
	return *o.ArtistFolderFormat.Get()
}

// GetArtistFolderFormatOk returns a tuple with the ArtistFolderFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetArtistFolderFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArtistFolderFormat.Get(), o.ArtistFolderFormat.IsSet()
}

// HasArtistFolderFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasArtistFolderFormat() bool {
	if o != nil && o.ArtistFolderFormat.IsSet() {
		return true
	}

	return false
}

// SetArtistFolderFormat gets a reference to the given NullableString and assigns it to the ArtistFolderFormat field.
func (o *NamingConfigResource) SetArtistFolderFormat(v string) {
	o.ArtistFolderFormat.Set(&v)
}
// SetArtistFolderFormatNil sets the value for ArtistFolderFormat to be an explicit nil
func (o *NamingConfigResource) SetArtistFolderFormatNil() {
	o.ArtistFolderFormat.Set(nil)
}

// UnsetArtistFolderFormat ensures that no value is present for ArtistFolderFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetArtistFolderFormat() {
	o.ArtistFolderFormat.Unset()
}

// GetIncludeArtistName returns the IncludeArtistName field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeArtistName() bool {
	if o == nil || IsNil(o.IncludeArtistName) {
		var ret bool
		return ret
	}
	return *o.IncludeArtistName
}

// GetIncludeArtistNameOk returns a tuple with the IncludeArtistName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeArtistNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeArtistName) {
		return nil, false
	}
	return o.IncludeArtistName, true
}

// HasIncludeArtistName returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeArtistName() bool {
	if o != nil && !IsNil(o.IncludeArtistName) {
		return true
	}

	return false
}

// SetIncludeArtistName gets a reference to the given bool and assigns it to the IncludeArtistName field.
func (o *NamingConfigResource) SetIncludeArtistName(v bool) {
	o.IncludeArtistName = &v
}

// GetIncludeAlbumTitle returns the IncludeAlbumTitle field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeAlbumTitle() bool {
	if o == nil || IsNil(o.IncludeAlbumTitle) {
		var ret bool
		return ret
	}
	return *o.IncludeAlbumTitle
}

// GetIncludeAlbumTitleOk returns a tuple with the IncludeAlbumTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeAlbumTitleOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAlbumTitle) {
		return nil, false
	}
	return o.IncludeAlbumTitle, true
}

// HasIncludeAlbumTitle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeAlbumTitle() bool {
	if o != nil && !IsNil(o.IncludeAlbumTitle) {
		return true
	}

	return false
}

// SetIncludeAlbumTitle gets a reference to the given bool and assigns it to the IncludeAlbumTitle field.
func (o *NamingConfigResource) SetIncludeAlbumTitle(v bool) {
	o.IncludeAlbumTitle = &v
}

// GetIncludeQuality returns the IncludeQuality field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeQuality() bool {
	if o == nil || IsNil(o.IncludeQuality) {
		var ret bool
		return ret
	}
	return *o.IncludeQuality
}

// GetIncludeQualityOk returns a tuple with the IncludeQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeQualityOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeQuality) {
		return nil, false
	}
	return o.IncludeQuality, true
}

// HasIncludeQuality returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeQuality() bool {
	if o != nil && !IsNil(o.IncludeQuality) {
		return true
	}

	return false
}

// SetIncludeQuality gets a reference to the given bool and assigns it to the IncludeQuality field.
func (o *NamingConfigResource) SetIncludeQuality(v bool) {
	o.IncludeQuality = &v
}

// GetReplaceSpaces returns the ReplaceSpaces field value if set, zero value otherwise.
func (o *NamingConfigResource) GetReplaceSpaces() bool {
	if o == nil || IsNil(o.ReplaceSpaces) {
		var ret bool
		return ret
	}
	return *o.ReplaceSpaces
}

// GetReplaceSpacesOk returns a tuple with the ReplaceSpaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceSpacesOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceSpaces) {
		return nil, false
	}
	return o.ReplaceSpaces, true
}

// HasReplaceSpaces returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceSpaces() bool {
	if o != nil && !IsNil(o.ReplaceSpaces) {
		return true
	}

	return false
}

// SetReplaceSpaces gets a reference to the given bool and assigns it to the ReplaceSpaces field.
func (o *NamingConfigResource) SetReplaceSpaces(v bool) {
	o.ReplaceSpaces = &v
}

// GetSeparator returns the Separator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetSeparator() string {
	if o == nil || IsNil(o.Separator.Get()) {
		var ret string
		return ret
	}
	return *o.Separator.Get()
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetSeparatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Separator.Get(), o.Separator.IsSet()
}

// HasSeparator returns a boolean if a field has been set.
func (o *NamingConfigResource) HasSeparator() bool {
	if o != nil && o.Separator.IsSet() {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given NullableString and assigns it to the Separator field.
func (o *NamingConfigResource) SetSeparator(v string) {
	o.Separator.Set(&v)
}
// SetSeparatorNil sets the value for Separator to be an explicit nil
func (o *NamingConfigResource) SetSeparatorNil() {
	o.Separator.Set(nil)
}

// UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
func (o *NamingConfigResource) UnsetSeparator() {
	o.Separator.Unset()
}

// GetNumberStyle returns the NumberStyle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetNumberStyle() string {
	if o == nil || IsNil(o.NumberStyle.Get()) {
		var ret string
		return ret
	}
	return *o.NumberStyle.Get()
}

// GetNumberStyleOk returns a tuple with the NumberStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetNumberStyleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberStyle.Get(), o.NumberStyle.IsSet()
}

// HasNumberStyle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasNumberStyle() bool {
	if o != nil && o.NumberStyle.IsSet() {
		return true
	}

	return false
}

// SetNumberStyle gets a reference to the given NullableString and assigns it to the NumberStyle field.
func (o *NamingConfigResource) SetNumberStyle(v string) {
	o.NumberStyle.Set(&v)
}
// SetNumberStyleNil sets the value for NumberStyle to be an explicit nil
func (o *NamingConfigResource) SetNumberStyleNil() {
	o.NumberStyle.Set(nil)
}

// UnsetNumberStyle ensures that no value is present for NumberStyle, not even an explicit nil
func (o *NamingConfigResource) UnsetNumberStyle() {
	o.NumberStyle.Unset()
}

func (o NamingConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamingConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RenameTracks) {
		toSerialize["renameTracks"] = o.RenameTracks
	}
	if !IsNil(o.ReplaceIllegalCharacters) {
		toSerialize["replaceIllegalCharacters"] = o.ReplaceIllegalCharacters
	}
	if !IsNil(o.ColonReplacementFormat) {
		toSerialize["colonReplacementFormat"] = o.ColonReplacementFormat
	}
	if o.StandardTrackFormat.IsSet() {
		toSerialize["standardTrackFormat"] = o.StandardTrackFormat.Get()
	}
	if o.MultiDiscTrackFormat.IsSet() {
		toSerialize["multiDiscTrackFormat"] = o.MultiDiscTrackFormat.Get()
	}
	if o.ArtistFolderFormat.IsSet() {
		toSerialize["artistFolderFormat"] = o.ArtistFolderFormat.Get()
	}
	if !IsNil(o.IncludeArtistName) {
		toSerialize["includeArtistName"] = o.IncludeArtistName
	}
	if !IsNil(o.IncludeAlbumTitle) {
		toSerialize["includeAlbumTitle"] = o.IncludeAlbumTitle
	}
	if !IsNil(o.IncludeQuality) {
		toSerialize["includeQuality"] = o.IncludeQuality
	}
	if !IsNil(o.ReplaceSpaces) {
		toSerialize["replaceSpaces"] = o.ReplaceSpaces
	}
	if o.Separator.IsSet() {
		toSerialize["separator"] = o.Separator.Get()
	}
	if o.NumberStyle.IsSet() {
		toSerialize["numberStyle"] = o.NumberStyle.Get()
	}
	return toSerialize, nil
}

type NullableNamingConfigResource struct {
	value *NamingConfigResource
	isSet bool
}

func (v NullableNamingConfigResource) Get() *NamingConfigResource {
	return v.value
}

func (v *NullableNamingConfigResource) Set(val *NamingConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableNamingConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableNamingConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamingConfigResource(val *NamingConfigResource) *NullableNamingConfigResource {
	return &NullableNamingConfigResource{value: val, isSet: true}
}

func (v NullableNamingConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamingConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


