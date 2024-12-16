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

// checks if the DownloadClientResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadClientResource{}

// DownloadClientResource struct for DownloadClientResource
type DownloadClientResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Presets []DownloadClientResource `json:"presets,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Protocol *DownloadProtocol `json:"protocol,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	RemoveCompletedDownloads *bool `json:"removeCompletedDownloads,omitempty"`
	RemoveFailedDownloads *bool `json:"removeFailedDownloads,omitempty"`
}

// NewDownloadClientResource instantiates a new DownloadClientResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadClientResource() *DownloadClientResource {
	this := DownloadClientResource{}
	return &this
}

// NewDownloadClientResourceWithDefaults instantiates a new DownloadClientResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadClientResourceWithDefaults() *DownloadClientResource {
	this := DownloadClientResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DownloadClientResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DownloadClientResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DownloadClientResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DownloadClientResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DownloadClientResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DownloadClientResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DownloadClientResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetFields() []Field {
	if o == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *DownloadClientResource) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *DownloadClientResource) SetFields(v []Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetImplementationName() string {
	if o == nil || IsNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *DownloadClientResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *DownloadClientResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *DownloadClientResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *DownloadClientResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetImplementation() string {
	if o == nil || IsNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetImplementationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *DownloadClientResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *DownloadClientResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *DownloadClientResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *DownloadClientResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetConfigContract() string {
	if o == nil || IsNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *DownloadClientResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *DownloadClientResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *DownloadClientResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *DownloadClientResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetInfoLink() string {
	if o == nil || IsNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *DownloadClientResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *DownloadClientResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *DownloadClientResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *DownloadClientResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DownloadClientResource) GetMessage() ProviderMessage {
	if o == nil || IsNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DownloadClientResource) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *DownloadClientResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DownloadClientResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *DownloadClientResource) SetTags(v []int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientResource) GetPresets() []DownloadClientResource {
	if o == nil {
		var ret []DownloadClientResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientResource) GetPresetsOk() ([]DownloadClientResource, bool) {
	if o == nil || IsNil(o.Presets) {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *DownloadClientResource) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []DownloadClientResource and assigns it to the Presets field.
func (o *DownloadClientResource) SetPresets(v []DownloadClientResource) {
	o.Presets = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *DownloadClientResource) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *DownloadClientResource) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *DownloadClientResource) SetEnable(v bool) {
	o.Enable = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *DownloadClientResource) GetProtocol() DownloadProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret DownloadProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetProtocolOk() (*DownloadProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *DownloadClientResource) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given DownloadProtocol and assigns it to the Protocol field.
func (o *DownloadClientResource) SetProtocol(v DownloadProtocol) {
	o.Protocol = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DownloadClientResource) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DownloadClientResource) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *DownloadClientResource) SetPriority(v int32) {
	o.Priority = &v
}

// GetRemoveCompletedDownloads returns the RemoveCompletedDownloads field value if set, zero value otherwise.
func (o *DownloadClientResource) GetRemoveCompletedDownloads() bool {
	if o == nil || IsNil(o.RemoveCompletedDownloads) {
		var ret bool
		return ret
	}
	return *o.RemoveCompletedDownloads
}

// GetRemoveCompletedDownloadsOk returns a tuple with the RemoveCompletedDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetRemoveCompletedDownloadsOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveCompletedDownloads) {
		return nil, false
	}
	return o.RemoveCompletedDownloads, true
}

// HasRemoveCompletedDownloads returns a boolean if a field has been set.
func (o *DownloadClientResource) HasRemoveCompletedDownloads() bool {
	if o != nil && !IsNil(o.RemoveCompletedDownloads) {
		return true
	}

	return false
}

// SetRemoveCompletedDownloads gets a reference to the given bool and assigns it to the RemoveCompletedDownloads field.
func (o *DownloadClientResource) SetRemoveCompletedDownloads(v bool) {
	o.RemoveCompletedDownloads = &v
}

// GetRemoveFailedDownloads returns the RemoveFailedDownloads field value if set, zero value otherwise.
func (o *DownloadClientResource) GetRemoveFailedDownloads() bool {
	if o == nil || IsNil(o.RemoveFailedDownloads) {
		var ret bool
		return ret
	}
	return *o.RemoveFailedDownloads
}

// GetRemoveFailedDownloadsOk returns a tuple with the RemoveFailedDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientResource) GetRemoveFailedDownloadsOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveFailedDownloads) {
		return nil, false
	}
	return o.RemoveFailedDownloads, true
}

// HasRemoveFailedDownloads returns a boolean if a field has been set.
func (o *DownloadClientResource) HasRemoveFailedDownloads() bool {
	if o != nil && !IsNil(o.RemoveFailedDownloads) {
		return true
	}

	return false
}

// SetRemoveFailedDownloads gets a reference to the given bool and assigns it to the RemoveFailedDownloads field.
func (o *DownloadClientResource) SetRemoveFailedDownloads(v bool) {
	o.RemoveFailedDownloads = &v
}

func (o DownloadClientResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadClientResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.ImplementationName.IsSet() {
		toSerialize["implementationName"] = o.ImplementationName.Get()
	}
	if o.Implementation.IsSet() {
		toSerialize["implementation"] = o.Implementation.Get()
	}
	if o.ConfigContract.IsSet() {
		toSerialize["configContract"] = o.ConfigContract.Get()
	}
	if o.InfoLink.IsSet() {
		toSerialize["infoLink"] = o.InfoLink.Get()
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.RemoveCompletedDownloads) {
		toSerialize["removeCompletedDownloads"] = o.RemoveCompletedDownloads
	}
	if !IsNil(o.RemoveFailedDownloads) {
		toSerialize["removeFailedDownloads"] = o.RemoveFailedDownloads
	}
	return toSerialize, nil
}

type NullableDownloadClientResource struct {
	value *DownloadClientResource
	isSet bool
}

func (v NullableDownloadClientResource) Get() *DownloadClientResource {
	return v.value
}

func (v *NullableDownloadClientResource) Set(val *DownloadClientResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadClientResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadClientResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadClientResource(val *DownloadClientResource) *NullableDownloadClientResource {
	return &NullableDownloadClientResource{value: val, isSet: true}
}

func (v NullableDownloadClientResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadClientResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


