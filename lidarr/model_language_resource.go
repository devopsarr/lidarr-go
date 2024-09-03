/*
Lidarr

Lidarr API docs

API version: v2.5.3.4341
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the LanguageResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguageResource{}

// LanguageResource struct for LanguageResource
type LanguageResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	NameLower NullableString `json:"nameLower,omitempty"`
}

// NewLanguageResource instantiates a new LanguageResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageResource() *LanguageResource {
	this := LanguageResource{}
	return &this
}

// NewLanguageResourceWithDefaults instantiates a new LanguageResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageResourceWithDefaults() *LanguageResource {
	this := LanguageResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LanguageResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LanguageResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LanguageResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LanguageResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanguageResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LanguageResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LanguageResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *LanguageResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LanguageResource) UnsetName() {
	o.Name.Unset()
}

// GetNameLower returns the NameLower field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LanguageResource) GetNameLower() string {
	if o == nil || IsNil(o.NameLower.Get()) {
		var ret string
		return ret
	}
	return *o.NameLower.Get()
}

// GetNameLowerOk returns a tuple with the NameLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanguageResource) GetNameLowerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameLower.Get(), o.NameLower.IsSet()
}

// HasNameLower returns a boolean if a field has been set.
func (o *LanguageResource) HasNameLower() bool {
	if o != nil && o.NameLower.IsSet() {
		return true
	}

	return false
}

// SetNameLower gets a reference to the given NullableString and assigns it to the NameLower field.
func (o *LanguageResource) SetNameLower(v string) {
	o.NameLower.Set(&v)
}
// SetNameLowerNil sets the value for NameLower to be an explicit nil
func (o *LanguageResource) SetNameLowerNil() {
	o.NameLower.Set(nil)
}

// UnsetNameLower ensures that no value is present for NameLower, not even an explicit nil
func (o *LanguageResource) UnsetNameLower() {
	o.NameLower.Unset()
}

func (o LanguageResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguageResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.NameLower.IsSet() {
		toSerialize["nameLower"] = o.NameLower.Get()
	}
	return toSerialize, nil
}

type NullableLanguageResource struct {
	value *LanguageResource
	isSet bool
}

func (v NullableLanguageResource) Get() *LanguageResource {
	return v.value
}

func (v *NullableLanguageResource) Set(val *LanguageResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageResource(val *LanguageResource) *NullableLanguageResource {
	return &NullableLanguageResource{value: val, isSet: true}
}

func (v NullableLanguageResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


