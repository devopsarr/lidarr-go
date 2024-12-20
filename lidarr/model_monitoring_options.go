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

// checks if the MonitoringOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringOptions{}

// MonitoringOptions struct for MonitoringOptions
type MonitoringOptions struct {
	Monitor *MonitorTypes `json:"monitor,omitempty"`
	AlbumsToMonitor []string `json:"albumsToMonitor,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
}

// NewMonitoringOptions instantiates a new MonitoringOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringOptions() *MonitoringOptions {
	this := MonitoringOptions{}
	return &this
}

// NewMonitoringOptionsWithDefaults instantiates a new MonitoringOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringOptionsWithDefaults() *MonitoringOptions {
	this := MonitoringOptions{}
	return &this
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *MonitoringOptions) GetMonitor() MonitorTypes {
	if o == nil || IsNil(o.Monitor) {
		var ret MonitorTypes
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetMonitorOk() (*MonitorTypes, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *MonitoringOptions) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorTypes and assigns it to the Monitor field.
func (o *MonitoringOptions) SetMonitor(v MonitorTypes) {
	o.Monitor = &v
}

// GetAlbumsToMonitor returns the AlbumsToMonitor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringOptions) GetAlbumsToMonitor() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AlbumsToMonitor
}

// GetAlbumsToMonitorOk returns a tuple with the AlbumsToMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringOptions) GetAlbumsToMonitorOk() ([]string, bool) {
	if o == nil || IsNil(o.AlbumsToMonitor) {
		return nil, false
	}
	return o.AlbumsToMonitor, true
}

// HasAlbumsToMonitor returns a boolean if a field has been set.
func (o *MonitoringOptions) HasAlbumsToMonitor() bool {
	if o != nil && !IsNil(o.AlbumsToMonitor) {
		return true
	}

	return false
}

// SetAlbumsToMonitor gets a reference to the given []string and assigns it to the AlbumsToMonitor field.
func (o *MonitoringOptions) SetAlbumsToMonitor(v []string) {
	o.AlbumsToMonitor = v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *MonitoringOptions) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *MonitoringOptions) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *MonitoringOptions) SetMonitored(v bool) {
	o.Monitored = &v
}

func (o MonitoringOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if o.AlbumsToMonitor != nil {
		toSerialize["albumsToMonitor"] = o.AlbumsToMonitor
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	return toSerialize, nil
}

type NullableMonitoringOptions struct {
	value *MonitoringOptions
	isSet bool
}

func (v NullableMonitoringOptions) Get() *MonitoringOptions {
	return v.value
}

func (v *NullableMonitoringOptions) Set(val *MonitoringOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringOptions(val *MonitoringOptions) *NullableMonitoringOptions {
	return &NullableMonitoringOptions{value: val, isSet: true}
}

func (v NullableMonitoringOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


