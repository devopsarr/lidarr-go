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

// checks if the AlbumResourcePagingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumResourcePagingResource{}

// AlbumResourcePagingResource struct for AlbumResourcePagingResource
type AlbumResourcePagingResource struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SortKey NullableString `json:"sortKey,omitempty"`
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	Records []AlbumResource `json:"records,omitempty"`
}

// NewAlbumResourcePagingResource instantiates a new AlbumResourcePagingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumResourcePagingResource() *AlbumResourcePagingResource {
	this := AlbumResourcePagingResource{}
	return &this
}

// NewAlbumResourcePagingResourceWithDefaults instantiates a new AlbumResourcePagingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumResourcePagingResourceWithDefaults() *AlbumResourcePagingResource {
	this := AlbumResourcePagingResource{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AlbumResourcePagingResource) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResourcePagingResource) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AlbumResourcePagingResource) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AlbumResourcePagingResource) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlbumResourcePagingResource) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResourcePagingResource) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlbumResourcePagingResource) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlbumResourcePagingResource) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResourcePagingResource) GetSortKey() string {
	if o == nil || IsNil(o.SortKey.Get()) {
		var ret string
		return ret
	}
	return *o.SortKey.Get()
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResourcePagingResource) GetSortKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortKey.Get(), o.SortKey.IsSet()
}

// HasSortKey returns a boolean if a field has been set.
func (o *AlbumResourcePagingResource) HasSortKey() bool {
	if o != nil && o.SortKey.IsSet() {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given NullableString and assigns it to the SortKey field.
func (o *AlbumResourcePagingResource) SetSortKey(v string) {
	o.SortKey.Set(&v)
}
// SetSortKeyNil sets the value for SortKey to be an explicit nil
func (o *AlbumResourcePagingResource) SetSortKeyNil() {
	o.SortKey.Set(nil)
}

// UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
func (o *AlbumResourcePagingResource) UnsetSortKey() {
	o.SortKey.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise.
func (o *AlbumResourcePagingResource) GetSortDirection() SortDirection {
	if o == nil || IsNil(o.SortDirection) {
		var ret SortDirection
		return ret
	}
	return *o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool) {
	if o == nil || IsNil(o.SortDirection) {
		return nil, false
	}
	return o.SortDirection, true
}

// HasSortDirection returns a boolean if a field has been set.
func (o *AlbumResourcePagingResource) HasSortDirection() bool {
	if o != nil && !IsNil(o.SortDirection) {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given SortDirection and assigns it to the SortDirection field.
func (o *AlbumResourcePagingResource) SetSortDirection(v SortDirection) {
	o.SortDirection = &v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *AlbumResourcePagingResource) GetTotalRecords() int32 {
	if o == nil || IsNil(o.TotalRecords) {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumResourcePagingResource) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecords) {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *AlbumResourcePagingResource) HasTotalRecords() bool {
	if o != nil && !IsNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *AlbumResourcePagingResource) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumResourcePagingResource) GetRecords() []AlbumResource {
	if o == nil {
		var ret []AlbumResource
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumResourcePagingResource) GetRecordsOk() ([]AlbumResource, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *AlbumResourcePagingResource) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []AlbumResource and assigns it to the Records field.
func (o *AlbumResourcePagingResource) SetRecords(v []AlbumResource) {
	o.Records = v
}

func (o AlbumResourcePagingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumResourcePagingResource) ToMap() (map[string]interface{}, error) {
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

type NullableAlbumResourcePagingResource struct {
	value *AlbumResourcePagingResource
	isSet bool
}

func (v NullableAlbumResourcePagingResource) Get() *AlbumResourcePagingResource {
	return v.value
}

func (v *NullableAlbumResourcePagingResource) Set(val *AlbumResourcePagingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumResourcePagingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumResourcePagingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumResourcePagingResource(val *AlbumResourcePagingResource) *NullableAlbumResourcePagingResource {
	return &NullableAlbumResourcePagingResource{value: val, isSet: true}
}

func (v NullableAlbumResourcePagingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumResourcePagingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


