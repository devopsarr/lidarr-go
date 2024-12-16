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

// checks if the QueueResourcePagingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueResourcePagingResource{}

// QueueResourcePagingResource struct for QueueResourcePagingResource
type QueueResourcePagingResource struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SortKey NullableString `json:"sortKey,omitempty"`
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	Records []QueueResource `json:"records,omitempty"`
}

// NewQueueResourcePagingResource instantiates a new QueueResourcePagingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueResourcePagingResource() *QueueResourcePagingResource {
	this := QueueResourcePagingResource{}
	return &this
}

// NewQueueResourcePagingResourceWithDefaults instantiates a new QueueResourcePagingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueResourcePagingResourceWithDefaults() *QueueResourcePagingResource {
	this := QueueResourcePagingResource{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *QueueResourcePagingResource) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResourcePagingResource) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *QueueResourcePagingResource) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *QueueResourcePagingResource) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *QueueResourcePagingResource) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResourcePagingResource) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *QueueResourcePagingResource) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *QueueResourcePagingResource) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResourcePagingResource) GetSortKey() string {
	if o == nil || IsNil(o.SortKey.Get()) {
		var ret string
		return ret
	}
	return *o.SortKey.Get()
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResourcePagingResource) GetSortKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortKey.Get(), o.SortKey.IsSet()
}

// HasSortKey returns a boolean if a field has been set.
func (o *QueueResourcePagingResource) HasSortKey() bool {
	if o != nil && o.SortKey.IsSet() {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given NullableString and assigns it to the SortKey field.
func (o *QueueResourcePagingResource) SetSortKey(v string) {
	o.SortKey.Set(&v)
}
// SetSortKeyNil sets the value for SortKey to be an explicit nil
func (o *QueueResourcePagingResource) SetSortKeyNil() {
	o.SortKey.Set(nil)
}

// UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
func (o *QueueResourcePagingResource) UnsetSortKey() {
	o.SortKey.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise.
func (o *QueueResourcePagingResource) GetSortDirection() SortDirection {
	if o == nil || IsNil(o.SortDirection) {
		var ret SortDirection
		return ret
	}
	return *o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool) {
	if o == nil || IsNil(o.SortDirection) {
		return nil, false
	}
	return o.SortDirection, true
}

// HasSortDirection returns a boolean if a field has been set.
func (o *QueueResourcePagingResource) HasSortDirection() bool {
	if o != nil && !IsNil(o.SortDirection) {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given SortDirection and assigns it to the SortDirection field.
func (o *QueueResourcePagingResource) SetSortDirection(v SortDirection) {
	o.SortDirection = &v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *QueueResourcePagingResource) GetTotalRecords() int32 {
	if o == nil || IsNil(o.TotalRecords) {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueResourcePagingResource) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecords) {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *QueueResourcePagingResource) HasTotalRecords() bool {
	if o != nil && !IsNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *QueueResourcePagingResource) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueResourcePagingResource) GetRecords() []QueueResource {
	if o == nil {
		var ret []QueueResource
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueResourcePagingResource) GetRecordsOk() ([]QueueResource, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *QueueResourcePagingResource) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []QueueResource and assigns it to the Records field.
func (o *QueueResourcePagingResource) SetRecords(v []QueueResource) {
	o.Records = v
}

func (o QueueResourcePagingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueResourcePagingResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.SortKey.IsSet() {
		toSerialize["sortKey"] = o.SortKey.Get()
	}
	if !IsNil(o.SortDirection) {
		toSerialize["sortDirection"] = o.SortDirection
	}
	if !IsNil(o.TotalRecords) {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableQueueResourcePagingResource struct {
	value *QueueResourcePagingResource
	isSet bool
}

func (v NullableQueueResourcePagingResource) Get() *QueueResourcePagingResource {
	return v.value
}

func (v *NullableQueueResourcePagingResource) Set(val *QueueResourcePagingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueResourcePagingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueResourcePagingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueResourcePagingResource(val *QueueResourcePagingResource) *NullableQueueResourcePagingResource {
	return &NullableQueueResourcePagingResource{value: val, isSet: true}
}

func (v NullableQueueResourcePagingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueResourcePagingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


