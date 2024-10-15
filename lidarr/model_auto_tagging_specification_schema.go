/*
Lidarr

Lidarr API docs

API version: v2.6.4.4402
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the AutoTaggingSpecificationSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTaggingSpecificationSchema{}

// AutoTaggingSpecificationSchema struct for AutoTaggingSpecificationSchema
type AutoTaggingSpecificationSchema struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Negate *bool `json:"negate,omitempty"`
	Required *bool `json:"required,omitempty"`
	Fields []Field `json:"fields,omitempty"`
}

// NewAutoTaggingSpecificationSchema instantiates a new AutoTaggingSpecificationSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTaggingSpecificationSchema() *AutoTaggingSpecificationSchema {
	this := AutoTaggingSpecificationSchema{}
	return &this
}

// NewAutoTaggingSpecificationSchemaWithDefaults instantiates a new AutoTaggingSpecificationSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTaggingSpecificationSchemaWithDefaults() *AutoTaggingSpecificationSchema {
	this := AutoTaggingSpecificationSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoTaggingSpecificationSchema) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTaggingSpecificationSchema) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AutoTaggingSpecificationSchema) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingSpecificationSchema) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingSpecificationSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AutoTaggingSpecificationSchema) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AutoTaggingSpecificationSchema) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AutoTaggingSpecificationSchema) UnsetName() {
	o.Name.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingSpecificationSchema) GetImplementation() string {
	if o == nil || IsNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingSpecificationSchema) GetImplementationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *AutoTaggingSpecificationSchema) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *AutoTaggingSpecificationSchema) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *AutoTaggingSpecificationSchema) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingSpecificationSchema) GetImplementationName() string {
	if o == nil || IsNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingSpecificationSchema) GetImplementationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *AutoTaggingSpecificationSchema) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *AutoTaggingSpecificationSchema) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *AutoTaggingSpecificationSchema) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *AutoTaggingSpecificationSchema) GetNegate() bool {
	if o == nil || IsNil(o.Negate) {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTaggingSpecificationSchema) GetNegateOk() (*bool, bool) {
	if o == nil || IsNil(o.Negate) {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasNegate() bool {
	if o != nil && !IsNil(o.Negate) {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *AutoTaggingSpecificationSchema) SetNegate(v bool) {
	o.Negate = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *AutoTaggingSpecificationSchema) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTaggingSpecificationSchema) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *AutoTaggingSpecificationSchema) SetRequired(v bool) {
	o.Required = &v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingSpecificationSchema) GetFields() []Field {
	if o == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingSpecificationSchema) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AutoTaggingSpecificationSchema) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *AutoTaggingSpecificationSchema) SetFields(v []Field) {
	o.Fields = v
}

func (o AutoTaggingSpecificationSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTaggingSpecificationSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Implementation.IsSet() {
		toSerialize["implementation"] = o.Implementation.Get()
	}
	if o.ImplementationName.IsSet() {
		toSerialize["implementationName"] = o.ImplementationName.Get()
	}
	if !IsNil(o.Negate) {
		toSerialize["negate"] = o.Negate
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableAutoTaggingSpecificationSchema struct {
	value *AutoTaggingSpecificationSchema
	isSet bool
}

func (v NullableAutoTaggingSpecificationSchema) Get() *AutoTaggingSpecificationSchema {
	return v.value
}

func (v *NullableAutoTaggingSpecificationSchema) Set(val *AutoTaggingSpecificationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTaggingSpecificationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTaggingSpecificationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTaggingSpecificationSchema(val *AutoTaggingSpecificationSchema) *NullableAutoTaggingSpecificationSchema {
	return &NullableAutoTaggingSpecificationSchema{value: val, isSet: true}
}

func (v NullableAutoTaggingSpecificationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTaggingSpecificationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


