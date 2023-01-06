/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
	"time"
)

// Command struct for Command
type Command struct {
	SendUpdatesToClient *bool `json:"sendUpdatesToClient,omitempty"`
	UpdateScheduledTask *bool `json:"updateScheduledTask,omitempty"`
	CompletionMessage NullableString `json:"completionMessage,omitempty"`
	RequiresDiskAccess *bool `json:"requiresDiskAccess,omitempty"`
	IsExclusive *bool `json:"isExclusive,omitempty"`
	IsTypeExclusive *bool `json:"isTypeExclusive,omitempty"`
	Name NullableString `json:"name,omitempty"`
	LastExecutionTime NullableTime `json:"lastExecutionTime,omitempty"`
	LastStartTime NullableTime `json:"lastStartTime,omitempty"`
	Trigger *CommandTrigger `json:"trigger,omitempty"`
	SuppressMessages *bool `json:"suppressMessages,omitempty"`
	ClientUserAgent NullableString `json:"clientUserAgent,omitempty"`
}

// NewCommand instantiates a new Command object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommand() *Command {
	this := Command{}
	return &this
}

// NewCommandWithDefaults instantiates a new Command object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandWithDefaults() *Command {
	this := Command{}
	return &this
}

// GetSendUpdatesToClient returns the SendUpdatesToClient field value if set, zero value otherwise.
func (o *Command) GetSendUpdatesToClient() bool {
	if o == nil || isNil(o.SendUpdatesToClient) {
		var ret bool
		return ret
	}
	return *o.SendUpdatesToClient
}

// GetSendUpdatesToClientOk returns a tuple with the SendUpdatesToClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetSendUpdatesToClientOk() (*bool, bool) {
	if o == nil || isNil(o.SendUpdatesToClient) {
    return nil, false
	}
	return o.SendUpdatesToClient, true
}

// HasSendUpdatesToClient returns a boolean if a field has been set.
func (o *Command) HasSendUpdatesToClient() bool {
	if o != nil && !isNil(o.SendUpdatesToClient) {
		return true
	}

	return false
}

// SetSendUpdatesToClient gets a reference to the given bool and assigns it to the SendUpdatesToClient field.
func (o *Command) SetSendUpdatesToClient(v bool) {
	o.SendUpdatesToClient = &v
}

// GetUpdateScheduledTask returns the UpdateScheduledTask field value if set, zero value otherwise.
func (o *Command) GetUpdateScheduledTask() bool {
	if o == nil || isNil(o.UpdateScheduledTask) {
		var ret bool
		return ret
	}
	return *o.UpdateScheduledTask
}

// GetUpdateScheduledTaskOk returns a tuple with the UpdateScheduledTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetUpdateScheduledTaskOk() (*bool, bool) {
	if o == nil || isNil(o.UpdateScheduledTask) {
    return nil, false
	}
	return o.UpdateScheduledTask, true
}

// HasUpdateScheduledTask returns a boolean if a field has been set.
func (o *Command) HasUpdateScheduledTask() bool {
	if o != nil && !isNil(o.UpdateScheduledTask) {
		return true
	}

	return false
}

// SetUpdateScheduledTask gets a reference to the given bool and assigns it to the UpdateScheduledTask field.
func (o *Command) SetUpdateScheduledTask(v bool) {
	o.UpdateScheduledTask = &v
}

// GetCompletionMessage returns the CompletionMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetCompletionMessage() string {
	if o == nil || isNil(o.CompletionMessage.Get()) {
		var ret string
		return ret
	}
	return *o.CompletionMessage.Get()
}

// GetCompletionMessageOk returns a tuple with the CompletionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetCompletionMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CompletionMessage.Get(), o.CompletionMessage.IsSet()
}

// HasCompletionMessage returns a boolean if a field has been set.
func (o *Command) HasCompletionMessage() bool {
	if o != nil && o.CompletionMessage.IsSet() {
		return true
	}

	return false
}

// SetCompletionMessage gets a reference to the given NullableString and assigns it to the CompletionMessage field.
func (o *Command) SetCompletionMessage(v string) {
	o.CompletionMessage.Set(&v)
}
// SetCompletionMessageNil sets the value for CompletionMessage to be an explicit nil
func (o *Command) SetCompletionMessageNil() {
	o.CompletionMessage.Set(nil)
}

// UnsetCompletionMessage ensures that no value is present for CompletionMessage, not even an explicit nil
func (o *Command) UnsetCompletionMessage() {
	o.CompletionMessage.Unset()
}

// GetRequiresDiskAccess returns the RequiresDiskAccess field value if set, zero value otherwise.
func (o *Command) GetRequiresDiskAccess() bool {
	if o == nil || isNil(o.RequiresDiskAccess) {
		var ret bool
		return ret
	}
	return *o.RequiresDiskAccess
}

// GetRequiresDiskAccessOk returns a tuple with the RequiresDiskAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetRequiresDiskAccessOk() (*bool, bool) {
	if o == nil || isNil(o.RequiresDiskAccess) {
    return nil, false
	}
	return o.RequiresDiskAccess, true
}

// HasRequiresDiskAccess returns a boolean if a field has been set.
func (o *Command) HasRequiresDiskAccess() bool {
	if o != nil && !isNil(o.RequiresDiskAccess) {
		return true
	}

	return false
}

// SetRequiresDiskAccess gets a reference to the given bool and assigns it to the RequiresDiskAccess field.
func (o *Command) SetRequiresDiskAccess(v bool) {
	o.RequiresDiskAccess = &v
}

// GetIsExclusive returns the IsExclusive field value if set, zero value otherwise.
func (o *Command) GetIsExclusive() bool {
	if o == nil || isNil(o.IsExclusive) {
		var ret bool
		return ret
	}
	return *o.IsExclusive
}

// GetIsExclusiveOk returns a tuple with the IsExclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetIsExclusiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsExclusive) {
    return nil, false
	}
	return o.IsExclusive, true
}

// HasIsExclusive returns a boolean if a field has been set.
func (o *Command) HasIsExclusive() bool {
	if o != nil && !isNil(o.IsExclusive) {
		return true
	}

	return false
}

// SetIsExclusive gets a reference to the given bool and assigns it to the IsExclusive field.
func (o *Command) SetIsExclusive(v bool) {
	o.IsExclusive = &v
}

// GetIsTypeExclusive returns the IsTypeExclusive field value if set, zero value otherwise.
func (o *Command) GetIsTypeExclusive() bool {
	if o == nil || isNil(o.IsTypeExclusive) {
		var ret bool
		return ret
	}
	return *o.IsTypeExclusive
}

// GetIsTypeExclusiveOk returns a tuple with the IsTypeExclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetIsTypeExclusiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsTypeExclusive) {
    return nil, false
	}
	return o.IsTypeExclusive, true
}

// HasIsTypeExclusive returns a boolean if a field has been set.
func (o *Command) HasIsTypeExclusive() bool {
	if o != nil && !isNil(o.IsTypeExclusive) {
		return true
	}

	return false
}

// SetIsTypeExclusive gets a reference to the given bool and assigns it to the IsTypeExclusive field.
func (o *Command) SetIsTypeExclusive(v bool) {
	o.IsTypeExclusive = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Command) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Command) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Command) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Command) UnsetName() {
	o.Name.Unset()
}

// GetLastExecutionTime returns the LastExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetLastExecutionTime() time.Time {
	if o == nil || isNil(o.LastExecutionTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastExecutionTime.Get()
}

// GetLastExecutionTimeOk returns a tuple with the LastExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetLastExecutionTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastExecutionTime.Get(), o.LastExecutionTime.IsSet()
}

// HasLastExecutionTime returns a boolean if a field has been set.
func (o *Command) HasLastExecutionTime() bool {
	if o != nil && o.LastExecutionTime.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionTime gets a reference to the given NullableTime and assigns it to the LastExecutionTime field.
func (o *Command) SetLastExecutionTime(v time.Time) {
	o.LastExecutionTime.Set(&v)
}
// SetLastExecutionTimeNil sets the value for LastExecutionTime to be an explicit nil
func (o *Command) SetLastExecutionTimeNil() {
	o.LastExecutionTime.Set(nil)
}

// UnsetLastExecutionTime ensures that no value is present for LastExecutionTime, not even an explicit nil
func (o *Command) UnsetLastExecutionTime() {
	o.LastExecutionTime.Unset()
}

// GetLastStartTime returns the LastStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetLastStartTime() time.Time {
	if o == nil || isNil(o.LastStartTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastStartTime.Get()
}

// GetLastStartTimeOk returns a tuple with the LastStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetLastStartTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastStartTime.Get(), o.LastStartTime.IsSet()
}

// HasLastStartTime returns a boolean if a field has been set.
func (o *Command) HasLastStartTime() bool {
	if o != nil && o.LastStartTime.IsSet() {
		return true
	}

	return false
}

// SetLastStartTime gets a reference to the given NullableTime and assigns it to the LastStartTime field.
func (o *Command) SetLastStartTime(v time.Time) {
	o.LastStartTime.Set(&v)
}
// SetLastStartTimeNil sets the value for LastStartTime to be an explicit nil
func (o *Command) SetLastStartTimeNil() {
	o.LastStartTime.Set(nil)
}

// UnsetLastStartTime ensures that no value is present for LastStartTime, not even an explicit nil
func (o *Command) UnsetLastStartTime() {
	o.LastStartTime.Unset()
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *Command) GetTrigger() CommandTrigger {
	if o == nil || isNil(o.Trigger) {
		var ret CommandTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetTriggerOk() (*CommandTrigger, bool) {
	if o == nil || isNil(o.Trigger) {
    return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *Command) HasTrigger() bool {
	if o != nil && !isNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given CommandTrigger and assigns it to the Trigger field.
func (o *Command) SetTrigger(v CommandTrigger) {
	o.Trigger = &v
}

// GetSuppressMessages returns the SuppressMessages field value if set, zero value otherwise.
func (o *Command) GetSuppressMessages() bool {
	if o == nil || isNil(o.SuppressMessages) {
		var ret bool
		return ret
	}
	return *o.SuppressMessages
}

// GetSuppressMessagesOk returns a tuple with the SuppressMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Command) GetSuppressMessagesOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressMessages) {
    return nil, false
	}
	return o.SuppressMessages, true
}

// HasSuppressMessages returns a boolean if a field has been set.
func (o *Command) HasSuppressMessages() bool {
	if o != nil && !isNil(o.SuppressMessages) {
		return true
	}

	return false
}

// SetSuppressMessages gets a reference to the given bool and assigns it to the SuppressMessages field.
func (o *Command) SetSuppressMessages(v bool) {
	o.SuppressMessages = &v
}

// GetClientUserAgent returns the ClientUserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Command) GetClientUserAgent() string {
	if o == nil || isNil(o.ClientUserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.ClientUserAgent.Get()
}

// GetClientUserAgentOk returns a tuple with the ClientUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Command) GetClientUserAgentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ClientUserAgent.Get(), o.ClientUserAgent.IsSet()
}

// HasClientUserAgent returns a boolean if a field has been set.
func (o *Command) HasClientUserAgent() bool {
	if o != nil && o.ClientUserAgent.IsSet() {
		return true
	}

	return false
}

// SetClientUserAgent gets a reference to the given NullableString and assigns it to the ClientUserAgent field.
func (o *Command) SetClientUserAgent(v string) {
	o.ClientUserAgent.Set(&v)
}
// SetClientUserAgentNil sets the value for ClientUserAgent to be an explicit nil
func (o *Command) SetClientUserAgentNil() {
	o.ClientUserAgent.Set(nil)
}

// UnsetClientUserAgent ensures that no value is present for ClientUserAgent, not even an explicit nil
func (o *Command) UnsetClientUserAgent() {
	o.ClientUserAgent.Unset()
}

func (o Command) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SendUpdatesToClient) {
		toSerialize["sendUpdatesToClient"] = o.SendUpdatesToClient
	}
	if !isNil(o.UpdateScheduledTask) {
		toSerialize["updateScheduledTask"] = o.UpdateScheduledTask
	}
	if o.CompletionMessage.IsSet() {
		toSerialize["completionMessage"] = o.CompletionMessage.Get()
	}
	if !isNil(o.RequiresDiskAccess) {
		toSerialize["requiresDiskAccess"] = o.RequiresDiskAccess
	}
	if !isNil(o.IsExclusive) {
		toSerialize["isExclusive"] = o.IsExclusive
	}
	if !isNil(o.IsTypeExclusive) {
		toSerialize["isTypeExclusive"] = o.IsTypeExclusive
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.LastExecutionTime.IsSet() {
		toSerialize["lastExecutionTime"] = o.LastExecutionTime.Get()
	}
	if o.LastStartTime.IsSet() {
		toSerialize["lastStartTime"] = o.LastStartTime.Get()
	}
	if !isNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !isNil(o.SuppressMessages) {
		toSerialize["suppressMessages"] = o.SuppressMessages
	}
	if o.ClientUserAgent.IsSet() {
		toSerialize["clientUserAgent"] = o.ClientUserAgent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommand struct {
	value *Command
	isSet bool
}

func (v NullableCommand) Get() *Command {
	return v.value
}

func (v *NullableCommand) Set(val *Command) {
	v.value = val
	v.isSet = true
}

func (v NullableCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommand(val *Command) *NullableCommand {
	return &NullableCommand{value: val, isSet: true}
}

func (v NullableCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


