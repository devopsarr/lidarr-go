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

// checks if the RootFolderResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootFolderResource{}

// RootFolderResource struct for RootFolderResource
type RootFolderResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Path NullableString `json:"path,omitempty"`
	DefaultMetadataProfileId *int32 `json:"defaultMetadataProfileId,omitempty"`
	DefaultQualityProfileId *int32 `json:"defaultQualityProfileId,omitempty"`
	DefaultMonitorOption *MonitorTypes `json:"defaultMonitorOption,omitempty"`
	DefaultNewItemMonitorOption *NewItemMonitorTypes `json:"defaultNewItemMonitorOption,omitempty"`
	DefaultTags []int32 `json:"defaultTags,omitempty"`
	Accessible *bool `json:"accessible,omitempty"`
	FreeSpace NullableInt64 `json:"freeSpace,omitempty"`
	TotalSpace NullableInt64 `json:"totalSpace,omitempty"`
}

// NewRootFolderResource instantiates a new RootFolderResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootFolderResource() *RootFolderResource {
	this := RootFolderResource{}
	return &this
}

// NewRootFolderResourceWithDefaults instantiates a new RootFolderResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootFolderResourceWithDefaults() *RootFolderResource {
	this := RootFolderResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RootFolderResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RootFolderResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RootFolderResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RootFolderResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RootFolderResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RootFolderResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RootFolderResource) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *RootFolderResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *RootFolderResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *RootFolderResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *RootFolderResource) UnsetPath() {
	o.Path.Unset()
}

// GetDefaultMetadataProfileId returns the DefaultMetadataProfileId field value if set, zero value otherwise.
func (o *RootFolderResource) GetDefaultMetadataProfileId() int32 {
	if o == nil || IsNil(o.DefaultMetadataProfileId) {
		var ret int32
		return ret
	}
	return *o.DefaultMetadataProfileId
}

// GetDefaultMetadataProfileIdOk returns a tuple with the DefaultMetadataProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetDefaultMetadataProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultMetadataProfileId) {
		return nil, false
	}
	return o.DefaultMetadataProfileId, true
}

// HasDefaultMetadataProfileId returns a boolean if a field has been set.
func (o *RootFolderResource) HasDefaultMetadataProfileId() bool {
	if o != nil && !IsNil(o.DefaultMetadataProfileId) {
		return true
	}

	return false
}

// SetDefaultMetadataProfileId gets a reference to the given int32 and assigns it to the DefaultMetadataProfileId field.
func (o *RootFolderResource) SetDefaultMetadataProfileId(v int32) {
	o.DefaultMetadataProfileId = &v
}

// GetDefaultQualityProfileId returns the DefaultQualityProfileId field value if set, zero value otherwise.
func (o *RootFolderResource) GetDefaultQualityProfileId() int32 {
	if o == nil || IsNil(o.DefaultQualityProfileId) {
		var ret int32
		return ret
	}
	return *o.DefaultQualityProfileId
}

// GetDefaultQualityProfileIdOk returns a tuple with the DefaultQualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetDefaultQualityProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultQualityProfileId) {
		return nil, false
	}
	return o.DefaultQualityProfileId, true
}

// HasDefaultQualityProfileId returns a boolean if a field has been set.
func (o *RootFolderResource) HasDefaultQualityProfileId() bool {
	if o != nil && !IsNil(o.DefaultQualityProfileId) {
		return true
	}

	return false
}

// SetDefaultQualityProfileId gets a reference to the given int32 and assigns it to the DefaultQualityProfileId field.
func (o *RootFolderResource) SetDefaultQualityProfileId(v int32) {
	o.DefaultQualityProfileId = &v
}

// GetDefaultMonitorOption returns the DefaultMonitorOption field value if set, zero value otherwise.
func (o *RootFolderResource) GetDefaultMonitorOption() MonitorTypes {
	if o == nil || IsNil(o.DefaultMonitorOption) {
		var ret MonitorTypes
		return ret
	}
	return *o.DefaultMonitorOption
}

// GetDefaultMonitorOptionOk returns a tuple with the DefaultMonitorOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetDefaultMonitorOptionOk() (*MonitorTypes, bool) {
	if o == nil || IsNil(o.DefaultMonitorOption) {
		return nil, false
	}
	return o.DefaultMonitorOption, true
}

// HasDefaultMonitorOption returns a boolean if a field has been set.
func (o *RootFolderResource) HasDefaultMonitorOption() bool {
	if o != nil && !IsNil(o.DefaultMonitorOption) {
		return true
	}

	return false
}

// SetDefaultMonitorOption gets a reference to the given MonitorTypes and assigns it to the DefaultMonitorOption field.
func (o *RootFolderResource) SetDefaultMonitorOption(v MonitorTypes) {
	o.DefaultMonitorOption = &v
}

// GetDefaultNewItemMonitorOption returns the DefaultNewItemMonitorOption field value if set, zero value otherwise.
func (o *RootFolderResource) GetDefaultNewItemMonitorOption() NewItemMonitorTypes {
	if o == nil || IsNil(o.DefaultNewItemMonitorOption) {
		var ret NewItemMonitorTypes
		return ret
	}
	return *o.DefaultNewItemMonitorOption
}

// GetDefaultNewItemMonitorOptionOk returns a tuple with the DefaultNewItemMonitorOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetDefaultNewItemMonitorOptionOk() (*NewItemMonitorTypes, bool) {
	if o == nil || IsNil(o.DefaultNewItemMonitorOption) {
		return nil, false
	}
	return o.DefaultNewItemMonitorOption, true
}

// HasDefaultNewItemMonitorOption returns a boolean if a field has been set.
func (o *RootFolderResource) HasDefaultNewItemMonitorOption() bool {
	if o != nil && !IsNil(o.DefaultNewItemMonitorOption) {
		return true
	}

	return false
}

// SetDefaultNewItemMonitorOption gets a reference to the given NewItemMonitorTypes and assigns it to the DefaultNewItemMonitorOption field.
func (o *RootFolderResource) SetDefaultNewItemMonitorOption(v NewItemMonitorTypes) {
	o.DefaultNewItemMonitorOption = &v
}

// GetDefaultTags returns the DefaultTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetDefaultTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.DefaultTags
}

// GetDefaultTagsOk returns a tuple with the DefaultTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetDefaultTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DefaultTags) {
		return nil, false
	}
	return o.DefaultTags, true
}

// HasDefaultTags returns a boolean if a field has been set.
func (o *RootFolderResource) HasDefaultTags() bool {
	if o != nil && !IsNil(o.DefaultTags) {
		return true
	}

	return false
}

// SetDefaultTags gets a reference to the given []int32 and assigns it to the DefaultTags field.
func (o *RootFolderResource) SetDefaultTags(v []int32) {
	o.DefaultTags = v
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *RootFolderResource) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible) {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Accessible) {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *RootFolderResource) HasAccessible() bool {
	if o != nil && !IsNil(o.Accessible) {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *RootFolderResource) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetFreeSpace returns the FreeSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetFreeSpace() int64 {
	if o == nil || IsNil(o.FreeSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.FreeSpace.Get()
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetFreeSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FreeSpace.Get(), o.FreeSpace.IsSet()
}

// HasFreeSpace returns a boolean if a field has been set.
func (o *RootFolderResource) HasFreeSpace() bool {
	if o != nil && o.FreeSpace.IsSet() {
		return true
	}

	return false
}

// SetFreeSpace gets a reference to the given NullableInt64 and assigns it to the FreeSpace field.
func (o *RootFolderResource) SetFreeSpace(v int64) {
	o.FreeSpace.Set(&v)
}
// SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil
func (o *RootFolderResource) SetFreeSpaceNil() {
	o.FreeSpace.Set(nil)
}

// UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
func (o *RootFolderResource) UnsetFreeSpace() {
	o.FreeSpace.Unset()
}

// GetTotalSpace returns the TotalSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetTotalSpace() int64 {
	if o == nil || IsNil(o.TotalSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.TotalSpace.Get()
}

// GetTotalSpaceOk returns a tuple with the TotalSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetTotalSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalSpace.Get(), o.TotalSpace.IsSet()
}

// HasTotalSpace returns a boolean if a field has been set.
func (o *RootFolderResource) HasTotalSpace() bool {
	if o != nil && o.TotalSpace.IsSet() {
		return true
	}

	return false
}

// SetTotalSpace gets a reference to the given NullableInt64 and assigns it to the TotalSpace field.
func (o *RootFolderResource) SetTotalSpace(v int64) {
	o.TotalSpace.Set(&v)
}
// SetTotalSpaceNil sets the value for TotalSpace to be an explicit nil
func (o *RootFolderResource) SetTotalSpaceNil() {
	o.TotalSpace.Set(nil)
}

// UnsetTotalSpace ensures that no value is present for TotalSpace, not even an explicit nil
func (o *RootFolderResource) UnsetTotalSpace() {
	o.TotalSpace.Unset()
}

func (o RootFolderResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootFolderResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.DefaultMetadataProfileId) {
		toSerialize["defaultMetadataProfileId"] = o.DefaultMetadataProfileId
	}
	if !IsNil(o.DefaultQualityProfileId) {
		toSerialize["defaultQualityProfileId"] = o.DefaultQualityProfileId
	}
	if !IsNil(o.DefaultMonitorOption) {
		toSerialize["defaultMonitorOption"] = o.DefaultMonitorOption
	}
	if !IsNil(o.DefaultNewItemMonitorOption) {
		toSerialize["defaultNewItemMonitorOption"] = o.DefaultNewItemMonitorOption
	}
	if o.DefaultTags != nil {
		toSerialize["defaultTags"] = o.DefaultTags
	}
	if !IsNil(o.Accessible) {
		toSerialize["accessible"] = o.Accessible
	}
	if o.FreeSpace.IsSet() {
		toSerialize["freeSpace"] = o.FreeSpace.Get()
	}
	if o.TotalSpace.IsSet() {
		toSerialize["totalSpace"] = o.TotalSpace.Get()
	}
	return toSerialize, nil
}

type NullableRootFolderResource struct {
	value *RootFolderResource
	isSet bool
}

func (v NullableRootFolderResource) Get() *RootFolderResource {
	return v.value
}

func (v *NullableRootFolderResource) Set(val *RootFolderResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRootFolderResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRootFolderResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootFolderResource(val *RootFolderResource) *NullableRootFolderResource {
	return &NullableRootFolderResource{value: val, isSet: true}
}

func (v NullableRootFolderResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootFolderResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


