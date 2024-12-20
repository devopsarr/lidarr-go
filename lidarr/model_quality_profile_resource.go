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

// checks if the QualityProfileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QualityProfileResource{}

// QualityProfileResource struct for QualityProfileResource
type QualityProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	UpgradeAllowed *bool `json:"upgradeAllowed,omitempty"`
	Cutoff *int32 `json:"cutoff,omitempty"`
	Items []QualityProfileQualityItemResource `json:"items,omitempty"`
	MinFormatScore *int32 `json:"minFormatScore,omitempty"`
	CutoffFormatScore *int32 `json:"cutoffFormatScore,omitempty"`
	FormatItems []ProfileFormatItemResource `json:"formatItems,omitempty"`
}

// NewQualityProfileResource instantiates a new QualityProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityProfileResource() *QualityProfileResource {
	this := QualityProfileResource{}
	return &this
}

// NewQualityProfileResourceWithDefaults instantiates a new QualityProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityProfileResourceWithDefaults() *QualityProfileResource {
	this := QualityProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QualityProfileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QualityProfileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *QualityProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *QualityProfileResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *QualityProfileResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *QualityProfileResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *QualityProfileResource) UnsetName() {
	o.Name.Unset()
}

// GetUpgradeAllowed returns the UpgradeAllowed field value if set, zero value otherwise.
func (o *QualityProfileResource) GetUpgradeAllowed() bool {
	if o == nil || IsNil(o.UpgradeAllowed) {
		var ret bool
		return ret
	}
	return *o.UpgradeAllowed
}

// GetUpgradeAllowedOk returns a tuple with the UpgradeAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileResource) GetUpgradeAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.UpgradeAllowed) {
		return nil, false
	}
	return o.UpgradeAllowed, true
}

// HasUpgradeAllowed returns a boolean if a field has been set.
func (o *QualityProfileResource) HasUpgradeAllowed() bool {
	if o != nil && !IsNil(o.UpgradeAllowed) {
		return true
	}

	return false
}

// SetUpgradeAllowed gets a reference to the given bool and assigns it to the UpgradeAllowed field.
func (o *QualityProfileResource) SetUpgradeAllowed(v bool) {
	o.UpgradeAllowed = &v
}

// GetCutoff returns the Cutoff field value if set, zero value otherwise.
func (o *QualityProfileResource) GetCutoff() int32 {
	if o == nil || IsNil(o.Cutoff) {
		var ret int32
		return ret
	}
	return *o.Cutoff
}

// GetCutoffOk returns a tuple with the Cutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileResource) GetCutoffOk() (*int32, bool) {
	if o == nil || IsNil(o.Cutoff) {
		return nil, false
	}
	return o.Cutoff, true
}

// HasCutoff returns a boolean if a field has been set.
func (o *QualityProfileResource) HasCutoff() bool {
	if o != nil && !IsNil(o.Cutoff) {
		return true
	}

	return false
}

// SetCutoff gets a reference to the given int32 and assigns it to the Cutoff field.
func (o *QualityProfileResource) SetCutoff(v int32) {
	o.Cutoff = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileResource) GetItems() []QualityProfileQualityItemResource {
	if o == nil {
		var ret []QualityProfileQualityItemResource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileResource) GetItemsOk() ([]QualityProfileQualityItemResource, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QualityProfileResource) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QualityProfileQualityItemResource and assigns it to the Items field.
func (o *QualityProfileResource) SetItems(v []QualityProfileQualityItemResource) {
	o.Items = v
}

// GetMinFormatScore returns the MinFormatScore field value if set, zero value otherwise.
func (o *QualityProfileResource) GetMinFormatScore() int32 {
	if o == nil || IsNil(o.MinFormatScore) {
		var ret int32
		return ret
	}
	return *o.MinFormatScore
}

// GetMinFormatScoreOk returns a tuple with the MinFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileResource) GetMinFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.MinFormatScore) {
		return nil, false
	}
	return o.MinFormatScore, true
}

// HasMinFormatScore returns a boolean if a field has been set.
func (o *QualityProfileResource) HasMinFormatScore() bool {
	if o != nil && !IsNil(o.MinFormatScore) {
		return true
	}

	return false
}

// SetMinFormatScore gets a reference to the given int32 and assigns it to the MinFormatScore field.
func (o *QualityProfileResource) SetMinFormatScore(v int32) {
	o.MinFormatScore = &v
}

// GetCutoffFormatScore returns the CutoffFormatScore field value if set, zero value otherwise.
func (o *QualityProfileResource) GetCutoffFormatScore() int32 {
	if o == nil || IsNil(o.CutoffFormatScore) {
		var ret int32
		return ret
	}
	return *o.CutoffFormatScore
}

// GetCutoffFormatScoreOk returns a tuple with the CutoffFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileResource) GetCutoffFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CutoffFormatScore) {
		return nil, false
	}
	return o.CutoffFormatScore, true
}

// HasCutoffFormatScore returns a boolean if a field has been set.
func (o *QualityProfileResource) HasCutoffFormatScore() bool {
	if o != nil && !IsNil(o.CutoffFormatScore) {
		return true
	}

	return false
}

// SetCutoffFormatScore gets a reference to the given int32 and assigns it to the CutoffFormatScore field.
func (o *QualityProfileResource) SetCutoffFormatScore(v int32) {
	o.CutoffFormatScore = &v
}

// GetFormatItems returns the FormatItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileResource) GetFormatItems() []ProfileFormatItemResource {
	if o == nil {
		var ret []ProfileFormatItemResource
		return ret
	}
	return o.FormatItems
}

// GetFormatItemsOk returns a tuple with the FormatItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileResource) GetFormatItemsOk() ([]ProfileFormatItemResource, bool) {
	if o == nil || IsNil(o.FormatItems) {
		return nil, false
	}
	return o.FormatItems, true
}

// HasFormatItems returns a boolean if a field has been set.
func (o *QualityProfileResource) HasFormatItems() bool {
	if o != nil && !IsNil(o.FormatItems) {
		return true
	}

	return false
}

// SetFormatItems gets a reference to the given []ProfileFormatItemResource and assigns it to the FormatItems field.
func (o *QualityProfileResource) SetFormatItems(v []ProfileFormatItemResource) {
	o.FormatItems = v
}

func (o QualityProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QualityProfileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.UpgradeAllowed) {
		toSerialize["upgradeAllowed"] = o.UpgradeAllowed
	}
	if !IsNil(o.Cutoff) {
		toSerialize["cutoff"] = o.Cutoff
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.MinFormatScore) {
		toSerialize["minFormatScore"] = o.MinFormatScore
	}
	if !IsNil(o.CutoffFormatScore) {
		toSerialize["cutoffFormatScore"] = o.CutoffFormatScore
	}
	if o.FormatItems != nil {
		toSerialize["formatItems"] = o.FormatItems
	}
	return toSerialize, nil
}

type NullableQualityProfileResource struct {
	value *QualityProfileResource
	isSet bool
}

func (v NullableQualityProfileResource) Get() *QualityProfileResource {
	return v.value
}

func (v *NullableQualityProfileResource) Set(val *QualityProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityProfileResource(val *QualityProfileResource) *NullableQualityProfileResource {
	return &NullableQualityProfileResource{value: val, isSet: true}
}

func (v NullableQualityProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


