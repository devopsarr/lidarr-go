/*
Lidarr

Lidarr API docs

API version: v2.7.1.4417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the IndexerBulkResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexerBulkResource{}

// IndexerBulkResource struct for IndexerBulkResource
type IndexerBulkResource struct {
	Ids []int32 `json:"ids,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	EnableRss NullableBool `json:"enableRss,omitempty"`
	EnableAutomaticSearch NullableBool `json:"enableAutomaticSearch,omitempty"`
	EnableInteractiveSearch NullableBool `json:"enableInteractiveSearch,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
}

// NewIndexerBulkResource instantiates a new IndexerBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerBulkResource() *IndexerBulkResource {
	this := IndexerBulkResource{}
	return &this
}

// NewIndexerBulkResourceWithDefaults instantiates a new IndexerBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerBulkResourceWithDefaults() *IndexerBulkResource {
	this := IndexerBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *IndexerBulkResource) SetIds(v []int32) {
	o.Ids = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *IndexerBulkResource) SetTags(v []int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *IndexerBulkResource) GetApplyTags() ApplyTags {
	if o == nil || IsNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerBulkResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || IsNil(o.ApplyTags) {
		return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasApplyTags() bool {
	if o != nil && !IsNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *IndexerBulkResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetEnableRss returns the EnableRss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetEnableRss() bool {
	if o == nil || IsNil(o.EnableRss.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableRss.Get()
}

// GetEnableRssOk returns a tuple with the EnableRss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetEnableRssOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableRss.Get(), o.EnableRss.IsSet()
}

// HasEnableRss returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasEnableRss() bool {
	if o != nil && o.EnableRss.IsSet() {
		return true
	}

	return false
}

// SetEnableRss gets a reference to the given NullableBool and assigns it to the EnableRss field.
func (o *IndexerBulkResource) SetEnableRss(v bool) {
	o.EnableRss.Set(&v)
}
// SetEnableRssNil sets the value for EnableRss to be an explicit nil
func (o *IndexerBulkResource) SetEnableRssNil() {
	o.EnableRss.Set(nil)
}

// UnsetEnableRss ensures that no value is present for EnableRss, not even an explicit nil
func (o *IndexerBulkResource) UnsetEnableRss() {
	o.EnableRss.Unset()
}

// GetEnableAutomaticSearch returns the EnableAutomaticSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetEnableAutomaticSearch() bool {
	if o == nil || IsNil(o.EnableAutomaticSearch.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAutomaticSearch.Get()
}

// GetEnableAutomaticSearchOk returns a tuple with the EnableAutomaticSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetEnableAutomaticSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableAutomaticSearch.Get(), o.EnableAutomaticSearch.IsSet()
}

// HasEnableAutomaticSearch returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasEnableAutomaticSearch() bool {
	if o != nil && o.EnableAutomaticSearch.IsSet() {
		return true
	}

	return false
}

// SetEnableAutomaticSearch gets a reference to the given NullableBool and assigns it to the EnableAutomaticSearch field.
func (o *IndexerBulkResource) SetEnableAutomaticSearch(v bool) {
	o.EnableAutomaticSearch.Set(&v)
}
// SetEnableAutomaticSearchNil sets the value for EnableAutomaticSearch to be an explicit nil
func (o *IndexerBulkResource) SetEnableAutomaticSearchNil() {
	o.EnableAutomaticSearch.Set(nil)
}

// UnsetEnableAutomaticSearch ensures that no value is present for EnableAutomaticSearch, not even an explicit nil
func (o *IndexerBulkResource) UnsetEnableAutomaticSearch() {
	o.EnableAutomaticSearch.Unset()
}

// GetEnableInteractiveSearch returns the EnableInteractiveSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetEnableInteractiveSearch() bool {
	if o == nil || IsNil(o.EnableInteractiveSearch.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableInteractiveSearch.Get()
}

// GetEnableInteractiveSearchOk returns a tuple with the EnableInteractiveSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetEnableInteractiveSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableInteractiveSearch.Get(), o.EnableInteractiveSearch.IsSet()
}

// HasEnableInteractiveSearch returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasEnableInteractiveSearch() bool {
	if o != nil && o.EnableInteractiveSearch.IsSet() {
		return true
	}

	return false
}

// SetEnableInteractiveSearch gets a reference to the given NullableBool and assigns it to the EnableInteractiveSearch field.
func (o *IndexerBulkResource) SetEnableInteractiveSearch(v bool) {
	o.EnableInteractiveSearch.Set(&v)
}
// SetEnableInteractiveSearchNil sets the value for EnableInteractiveSearch to be an explicit nil
func (o *IndexerBulkResource) SetEnableInteractiveSearchNil() {
	o.EnableInteractiveSearch.Set(nil)
}

// UnsetEnableInteractiveSearch ensures that no value is present for EnableInteractiveSearch, not even an explicit nil
func (o *IndexerBulkResource) UnsetEnableInteractiveSearch() {
	o.EnableInteractiveSearch.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *IndexerBulkResource) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *IndexerBulkResource) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *IndexerBulkResource) UnsetPriority() {
	o.Priority.Unset()
}

func (o IndexerBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexerBulkResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	if o.EnableRss.IsSet() {
		toSerialize["enableRss"] = o.EnableRss.Get()
	}
	if o.EnableAutomaticSearch.IsSet() {
		toSerialize["enableAutomaticSearch"] = o.EnableAutomaticSearch.Get()
	}
	if o.EnableInteractiveSearch.IsSet() {
		toSerialize["enableInteractiveSearch"] = o.EnableInteractiveSearch.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	return toSerialize, nil
}

type NullableIndexerBulkResource struct {
	value *IndexerBulkResource
	isSet bool
}

func (v NullableIndexerBulkResource) Get() *IndexerBulkResource {
	return v.value
}

func (v *NullableIndexerBulkResource) Set(val *IndexerBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerBulkResource(val *IndexerBulkResource) *NullableIndexerBulkResource {
	return &NullableIndexerBulkResource{value: val, isSet: true}
}

func (v NullableIndexerBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


