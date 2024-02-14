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

// ImportListBulkResource struct for ImportListBulkResource
type ImportListBulkResource struct {
	Ids []*int32 `json:"ids,omitempty"`
	Tags []*int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	EnableAutomaticAdd NullableBool `json:"enableAutomaticAdd,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	QualityProfileId NullableInt32 `json:"qualityProfileId,omitempty"`
}

// NewImportListBulkResource instantiates a new ImportListBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListBulkResource() *ImportListBulkResource {
	this := ImportListBulkResource{}
	return &this
}

// NewImportListBulkResourceWithDefaults instantiates a new ImportListBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListBulkResourceWithDefaults() *ImportListBulkResource {
	this := ImportListBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListBulkResource) GetIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListBulkResource) GetIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ImportListBulkResource) HasIds() bool {
	if o != nil && isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *ImportListBulkResource) SetIds(v []*int32) {
	o.Ids = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListBulkResource) GetTags() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListBulkResource) GetTagsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImportListBulkResource) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ImportListBulkResource) SetTags(v []*int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *ImportListBulkResource) GetApplyTags() ApplyTags {
	if o == nil || isNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListBulkResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || isNil(o.ApplyTags) {
    return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *ImportListBulkResource) HasApplyTags() bool {
	if o != nil && !isNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *ImportListBulkResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetEnableAutomaticAdd returns the EnableAutomaticAdd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListBulkResource) GetEnableAutomaticAdd() bool {
	if o == nil || isNil(o.EnableAutomaticAdd.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAutomaticAdd.Get()
}

// GetEnableAutomaticAddOk returns a tuple with the EnableAutomaticAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListBulkResource) GetEnableAutomaticAddOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnableAutomaticAdd.Get(), o.EnableAutomaticAdd.IsSet()
}

// HasEnableAutomaticAdd returns a boolean if a field has been set.
func (o *ImportListBulkResource) HasEnableAutomaticAdd() bool {
	if o != nil && o.EnableAutomaticAdd.IsSet() {
		return true
	}

	return false
}

// SetEnableAutomaticAdd gets a reference to the given NullableBool and assigns it to the EnableAutomaticAdd field.
func (o *ImportListBulkResource) SetEnableAutomaticAdd(v bool) {
	o.EnableAutomaticAdd.Set(&v)
}
// SetEnableAutomaticAddNil sets the value for EnableAutomaticAdd to be an explicit nil
func (o *ImportListBulkResource) SetEnableAutomaticAddNil() {
	o.EnableAutomaticAdd.Set(nil)
}

// UnsetEnableAutomaticAdd ensures that no value is present for EnableAutomaticAdd, not even an explicit nil
func (o *ImportListBulkResource) UnsetEnableAutomaticAdd() {
	o.EnableAutomaticAdd.Unset()
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListBulkResource) GetRootFolderPath() string {
	if o == nil || isNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListBulkResource) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *ImportListBulkResource) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *ImportListBulkResource) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *ImportListBulkResource) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *ImportListBulkResource) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListBulkResource) GetQualityProfileId() int32 {
	if o == nil || isNil(o.QualityProfileId.Get()) {
		var ret int32
		return ret
	}
	return *o.QualityProfileId.Get()
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListBulkResource) GetQualityProfileIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.QualityProfileId.Get(), o.QualityProfileId.IsSet()
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *ImportListBulkResource) HasQualityProfileId() bool {
	if o != nil && o.QualityProfileId.IsSet() {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given NullableInt32 and assigns it to the QualityProfileId field.
func (o *ImportListBulkResource) SetQualityProfileId(v int32) {
	o.QualityProfileId.Set(&v)
}
// SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil
func (o *ImportListBulkResource) SetQualityProfileIdNil() {
	o.QualityProfileId.Set(nil)
}

// UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
func (o *ImportListBulkResource) UnsetQualityProfileId() {
	o.QualityProfileId.Unset()
}

func (o ImportListBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	if o.EnableAutomaticAdd.IsSet() {
		toSerialize["enableAutomaticAdd"] = o.EnableAutomaticAdd.Get()
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if o.QualityProfileId.IsSet() {
		toSerialize["qualityProfileId"] = o.QualityProfileId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableImportListBulkResource struct {
	value *ImportListBulkResource
	isSet bool
}

func (v NullableImportListBulkResource) Get() *ImportListBulkResource {
	return v.value
}

func (v *NullableImportListBulkResource) Set(val *ImportListBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListBulkResource(val *ImportListBulkResource) *NullableImportListBulkResource {
	return &NullableImportListBulkResource{value: val, isSet: true}
}

func (v NullableImportListBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


