/*
Lidarr

Lidarr API docs

API version: v2.4.3.4248
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the TrackedDownloadStatusMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackedDownloadStatusMessage{}

// TrackedDownloadStatusMessage struct for TrackedDownloadStatusMessage
type TrackedDownloadStatusMessage struct {
	Title NullableString `json:"title,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

// NewTrackedDownloadStatusMessage instantiates a new TrackedDownloadStatusMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackedDownloadStatusMessage() *TrackedDownloadStatusMessage {
	this := TrackedDownloadStatusMessage{}
	return &this
}

// NewTrackedDownloadStatusMessageWithDefaults instantiates a new TrackedDownloadStatusMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackedDownloadStatusMessageWithDefaults() *TrackedDownloadStatusMessage {
	this := TrackedDownloadStatusMessage{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackedDownloadStatusMessage) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackedDownloadStatusMessage) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TrackedDownloadStatusMessage) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TrackedDownloadStatusMessage) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TrackedDownloadStatusMessage) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TrackedDownloadStatusMessage) UnsetTitle() {
	o.Title.Unset()
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackedDownloadStatusMessage) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackedDownloadStatusMessage) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *TrackedDownloadStatusMessage) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *TrackedDownloadStatusMessage) SetMessages(v []string) {
	o.Messages = v
}

func (o TrackedDownloadStatusMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackedDownloadStatusMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableTrackedDownloadStatusMessage struct {
	value *TrackedDownloadStatusMessage
	isSet bool
}

func (v NullableTrackedDownloadStatusMessage) Get() *TrackedDownloadStatusMessage {
	return v.value
}

func (v *NullableTrackedDownloadStatusMessage) Set(val *TrackedDownloadStatusMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackedDownloadStatusMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackedDownloadStatusMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackedDownloadStatusMessage(val *TrackedDownloadStatusMessage) *NullableTrackedDownloadStatusMessage {
	return &NullableTrackedDownloadStatusMessage{value: val, isSet: true}
}

func (v NullableTrackedDownloadStatusMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackedDownloadStatusMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


