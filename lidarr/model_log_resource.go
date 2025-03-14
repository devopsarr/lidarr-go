/*
Lidarr

Lidarr API docs

API version: v2.9.6.4552
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"time"
)

// checks if the LogResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogResource{}

// LogResource struct for LogResource
type LogResource struct {
	Id *int32 `json:"id,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	Exception NullableString `json:"exception,omitempty"`
	ExceptionType NullableString `json:"exceptionType,omitempty"`
	Level NullableString `json:"level,omitempty"`
	Logger NullableString `json:"logger,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Method NullableString `json:"method,omitempty"`
}

// NewLogResource instantiates a new LogResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResource() *LogResource {
	this := LogResource{}
	return &this
}

// NewLogResourceWithDefaults instantiates a new LogResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResourceWithDefaults() *LogResource {
	this := LogResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LogResource) SetId(v int32) {
	o.Id = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LogResource) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResource) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LogResource) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *LogResource) SetTime(v time.Time) {
	o.Time = &v
}

// GetException returns the Exception field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResource) GetException() string {
	if o == nil || IsNil(o.Exception.Get()) {
		var ret string
		return ret
	}
	return *o.Exception.Get()
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResource) GetExceptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exception.Get(), o.Exception.IsSet()
}

// HasException returns a boolean if a field has been set.
func (o *LogResource) HasException() bool {
	if o != nil && o.Exception.IsSet() {
		return true
	}

	return false
}

// SetException gets a reference to the given NullableString and assigns it to the Exception field.
func (o *LogResource) SetException(v string) {
	o.Exception.Set(&v)
}
// SetExceptionNil sets the value for Exception to be an explicit nil
func (o *LogResource) SetExceptionNil() {
	o.Exception.Set(nil)
}

// UnsetException ensures that no value is present for Exception, not even an explicit nil
func (o *LogResource) UnsetException() {
	o.Exception.Unset()
}

// GetExceptionType returns the ExceptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResource) GetExceptionType() string {
	if o == nil || IsNil(o.ExceptionType.Get()) {
		var ret string
		return ret
	}
	return *o.ExceptionType.Get()
}

// GetExceptionTypeOk returns a tuple with the ExceptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResource) GetExceptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExceptionType.Get(), o.ExceptionType.IsSet()
}

// HasExceptionType returns a boolean if a field has been set.
func (o *LogResource) HasExceptionType() bool {
	if o != nil && o.ExceptionType.IsSet() {
		return true
	}

	return false
}

// SetExceptionType gets a reference to the given NullableString and assigns it to the ExceptionType field.
func (o *LogResource) SetExceptionType(v string) {
	o.ExceptionType.Set(&v)
}
// SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil
func (o *LogResource) SetExceptionTypeNil() {
	o.ExceptionType.Set(nil)
}

// UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
func (o *LogResource) UnsetExceptionType() {
	o.ExceptionType.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResource) GetLevel() string {
	if o == nil || IsNil(o.Level.Get()) {
		var ret string
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResource) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *LogResource) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableString and assigns it to the Level field.
func (o *LogResource) SetLevel(v string) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *LogResource) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *LogResource) UnsetLevel() {
	o.Level.Unset()
}

// GetLogger returns the Logger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResource) GetLogger() string {
	if o == nil || IsNil(o.Logger.Get()) {
		var ret string
		return ret
	}
	return *o.Logger.Get()
}

// GetLoggerOk returns a tuple with the Logger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResource) GetLoggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logger.Get(), o.Logger.IsSet()
}

// HasLogger returns a boolean if a field has been set.
func (o *LogResource) HasLogger() bool {
	if o != nil && o.Logger.IsSet() {
		return true
	}

	return false
}

// SetLogger gets a reference to the given NullableString and assigns it to the Logger field.
func (o *LogResource) SetLogger(v string) {
	o.Logger.Set(&v)
}
// SetLoggerNil sets the value for Logger to be an explicit nil
func (o *LogResource) SetLoggerNil() {
	o.Logger.Set(nil)
}

// UnsetLogger ensures that no value is present for Logger, not even an explicit nil
func (o *LogResource) UnsetLogger() {
	o.Logger.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResource) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResource) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *LogResource) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *LogResource) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *LogResource) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *LogResource) UnsetMessage() {
	o.Message.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResource) GetMethod() string {
	if o == nil || IsNil(o.Method.Get()) {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResource) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *LogResource) HasMethod() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *LogResource) SetMethod(v string) {
	o.Method.Set(&v)
}
// SetMethodNil sets the value for Method to be an explicit nil
func (o *LogResource) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *LogResource) UnsetMethod() {
	o.Method.Unset()
}

func (o LogResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if o.Exception.IsSet() {
		toSerialize["exception"] = o.Exception.Get()
	}
	if o.ExceptionType.IsSet() {
		toSerialize["exceptionType"] = o.ExceptionType.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.Logger.IsSet() {
		toSerialize["logger"] = o.Logger.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Method.IsSet() {
		toSerialize["method"] = o.Method.Get()
	}
	return toSerialize, nil
}

type NullableLogResource struct {
	value *LogResource
	isSet bool
}

func (v NullableLogResource) Get() *LogResource {
	return v.value
}

func (v *NullableLogResource) Set(val *LogResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResource(val *LogResource) *NullableLogResource {
	return &NullableLogResource{value: val, isSet: true}
}

func (v NullableLogResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


