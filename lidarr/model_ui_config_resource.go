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

// checks if the UiConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiConfigResource{}

// UiConfigResource struct for UiConfigResource
type UiConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	FirstDayOfWeek *int32 `json:"firstDayOfWeek,omitempty"`
	CalendarWeekColumnHeader NullableString `json:"calendarWeekColumnHeader,omitempty"`
	ShortDateFormat NullableString `json:"shortDateFormat,omitempty"`
	LongDateFormat NullableString `json:"longDateFormat,omitempty"`
	TimeFormat NullableString `json:"timeFormat,omitempty"`
	ShowRelativeDates *bool `json:"showRelativeDates,omitempty"`
	EnableColorImpairedMode *bool `json:"enableColorImpairedMode,omitempty"`
	UiLanguage *int32 `json:"uiLanguage,omitempty"`
	ExpandAlbumByDefault *bool `json:"expandAlbumByDefault,omitempty"`
	ExpandSingleByDefault *bool `json:"expandSingleByDefault,omitempty"`
	ExpandEPByDefault *bool `json:"expandEPByDefault,omitempty"`
	ExpandBroadcastByDefault *bool `json:"expandBroadcastByDefault,omitempty"`
	ExpandOtherByDefault *bool `json:"expandOtherByDefault,omitempty"`
	Theme NullableString `json:"theme,omitempty"`
}

// NewUiConfigResource instantiates a new UiConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiConfigResource() *UiConfigResource {
	this := UiConfigResource{}
	return &this
}

// NewUiConfigResourceWithDefaults instantiates a new UiConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiConfigResourceWithDefaults() *UiConfigResource {
	this := UiConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UiConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UiConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UiConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetFirstDayOfWeek returns the FirstDayOfWeek field value if set, zero value otherwise.
func (o *UiConfigResource) GetFirstDayOfWeek() int32 {
	if o == nil || IsNil(o.FirstDayOfWeek) {
		var ret int32
		return ret
	}
	return *o.FirstDayOfWeek
}

// GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetFirstDayOfWeekOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstDayOfWeek) {
		return nil, false
	}
	return o.FirstDayOfWeek, true
}

// HasFirstDayOfWeek returns a boolean if a field has been set.
func (o *UiConfigResource) HasFirstDayOfWeek() bool {
	if o != nil && !IsNil(o.FirstDayOfWeek) {
		return true
	}

	return false
}

// SetFirstDayOfWeek gets a reference to the given int32 and assigns it to the FirstDayOfWeek field.
func (o *UiConfigResource) SetFirstDayOfWeek(v int32) {
	o.FirstDayOfWeek = &v
}

// GetCalendarWeekColumnHeader returns the CalendarWeekColumnHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetCalendarWeekColumnHeader() string {
	if o == nil || IsNil(o.CalendarWeekColumnHeader.Get()) {
		var ret string
		return ret
	}
	return *o.CalendarWeekColumnHeader.Get()
}

// GetCalendarWeekColumnHeaderOk returns a tuple with the CalendarWeekColumnHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetCalendarWeekColumnHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CalendarWeekColumnHeader.Get(), o.CalendarWeekColumnHeader.IsSet()
}

// HasCalendarWeekColumnHeader returns a boolean if a field has been set.
func (o *UiConfigResource) HasCalendarWeekColumnHeader() bool {
	if o != nil && o.CalendarWeekColumnHeader.IsSet() {
		return true
	}

	return false
}

// SetCalendarWeekColumnHeader gets a reference to the given NullableString and assigns it to the CalendarWeekColumnHeader field.
func (o *UiConfigResource) SetCalendarWeekColumnHeader(v string) {
	o.CalendarWeekColumnHeader.Set(&v)
}
// SetCalendarWeekColumnHeaderNil sets the value for CalendarWeekColumnHeader to be an explicit nil
func (o *UiConfigResource) SetCalendarWeekColumnHeaderNil() {
	o.CalendarWeekColumnHeader.Set(nil)
}

// UnsetCalendarWeekColumnHeader ensures that no value is present for CalendarWeekColumnHeader, not even an explicit nil
func (o *UiConfigResource) UnsetCalendarWeekColumnHeader() {
	o.CalendarWeekColumnHeader.Unset()
}

// GetShortDateFormat returns the ShortDateFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetShortDateFormat() string {
	if o == nil || IsNil(o.ShortDateFormat.Get()) {
		var ret string
		return ret
	}
	return *o.ShortDateFormat.Get()
}

// GetShortDateFormatOk returns a tuple with the ShortDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetShortDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortDateFormat.Get(), o.ShortDateFormat.IsSet()
}

// HasShortDateFormat returns a boolean if a field has been set.
func (o *UiConfigResource) HasShortDateFormat() bool {
	if o != nil && o.ShortDateFormat.IsSet() {
		return true
	}

	return false
}

// SetShortDateFormat gets a reference to the given NullableString and assigns it to the ShortDateFormat field.
func (o *UiConfigResource) SetShortDateFormat(v string) {
	o.ShortDateFormat.Set(&v)
}
// SetShortDateFormatNil sets the value for ShortDateFormat to be an explicit nil
func (o *UiConfigResource) SetShortDateFormatNil() {
	o.ShortDateFormat.Set(nil)
}

// UnsetShortDateFormat ensures that no value is present for ShortDateFormat, not even an explicit nil
func (o *UiConfigResource) UnsetShortDateFormat() {
	o.ShortDateFormat.Unset()
}

// GetLongDateFormat returns the LongDateFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetLongDateFormat() string {
	if o == nil || IsNil(o.LongDateFormat.Get()) {
		var ret string
		return ret
	}
	return *o.LongDateFormat.Get()
}

// GetLongDateFormatOk returns a tuple with the LongDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetLongDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LongDateFormat.Get(), o.LongDateFormat.IsSet()
}

// HasLongDateFormat returns a boolean if a field has been set.
func (o *UiConfigResource) HasLongDateFormat() bool {
	if o != nil && o.LongDateFormat.IsSet() {
		return true
	}

	return false
}

// SetLongDateFormat gets a reference to the given NullableString and assigns it to the LongDateFormat field.
func (o *UiConfigResource) SetLongDateFormat(v string) {
	o.LongDateFormat.Set(&v)
}
// SetLongDateFormatNil sets the value for LongDateFormat to be an explicit nil
func (o *UiConfigResource) SetLongDateFormatNil() {
	o.LongDateFormat.Set(nil)
}

// UnsetLongDateFormat ensures that no value is present for LongDateFormat, not even an explicit nil
func (o *UiConfigResource) UnsetLongDateFormat() {
	o.LongDateFormat.Unset()
}

// GetTimeFormat returns the TimeFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetTimeFormat() string {
	if o == nil || IsNil(o.TimeFormat.Get()) {
		var ret string
		return ret
	}
	return *o.TimeFormat.Get()
}

// GetTimeFormatOk returns a tuple with the TimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetTimeFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeFormat.Get(), o.TimeFormat.IsSet()
}

// HasTimeFormat returns a boolean if a field has been set.
func (o *UiConfigResource) HasTimeFormat() bool {
	if o != nil && o.TimeFormat.IsSet() {
		return true
	}

	return false
}

// SetTimeFormat gets a reference to the given NullableString and assigns it to the TimeFormat field.
func (o *UiConfigResource) SetTimeFormat(v string) {
	o.TimeFormat.Set(&v)
}
// SetTimeFormatNil sets the value for TimeFormat to be an explicit nil
func (o *UiConfigResource) SetTimeFormatNil() {
	o.TimeFormat.Set(nil)
}

// UnsetTimeFormat ensures that no value is present for TimeFormat, not even an explicit nil
func (o *UiConfigResource) UnsetTimeFormat() {
	o.TimeFormat.Unset()
}

// GetShowRelativeDates returns the ShowRelativeDates field value if set, zero value otherwise.
func (o *UiConfigResource) GetShowRelativeDates() bool {
	if o == nil || IsNil(o.ShowRelativeDates) {
		var ret bool
		return ret
	}
	return *o.ShowRelativeDates
}

// GetShowRelativeDatesOk returns a tuple with the ShowRelativeDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetShowRelativeDatesOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowRelativeDates) {
		return nil, false
	}
	return o.ShowRelativeDates, true
}

// HasShowRelativeDates returns a boolean if a field has been set.
func (o *UiConfigResource) HasShowRelativeDates() bool {
	if o != nil && !IsNil(o.ShowRelativeDates) {
		return true
	}

	return false
}

// SetShowRelativeDates gets a reference to the given bool and assigns it to the ShowRelativeDates field.
func (o *UiConfigResource) SetShowRelativeDates(v bool) {
	o.ShowRelativeDates = &v
}

// GetEnableColorImpairedMode returns the EnableColorImpairedMode field value if set, zero value otherwise.
func (o *UiConfigResource) GetEnableColorImpairedMode() bool {
	if o == nil || IsNil(o.EnableColorImpairedMode) {
		var ret bool
		return ret
	}
	return *o.EnableColorImpairedMode
}

// GetEnableColorImpairedModeOk returns a tuple with the EnableColorImpairedMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetEnableColorImpairedModeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableColorImpairedMode) {
		return nil, false
	}
	return o.EnableColorImpairedMode, true
}

// HasEnableColorImpairedMode returns a boolean if a field has been set.
func (o *UiConfigResource) HasEnableColorImpairedMode() bool {
	if o != nil && !IsNil(o.EnableColorImpairedMode) {
		return true
	}

	return false
}

// SetEnableColorImpairedMode gets a reference to the given bool and assigns it to the EnableColorImpairedMode field.
func (o *UiConfigResource) SetEnableColorImpairedMode(v bool) {
	o.EnableColorImpairedMode = &v
}

// GetUiLanguage returns the UiLanguage field value if set, zero value otherwise.
func (o *UiConfigResource) GetUiLanguage() int32 {
	if o == nil || IsNil(o.UiLanguage) {
		var ret int32
		return ret
	}
	return *o.UiLanguage
}

// GetUiLanguageOk returns a tuple with the UiLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetUiLanguageOk() (*int32, bool) {
	if o == nil || IsNil(o.UiLanguage) {
		return nil, false
	}
	return o.UiLanguage, true
}

// HasUiLanguage returns a boolean if a field has been set.
func (o *UiConfigResource) HasUiLanguage() bool {
	if o != nil && !IsNil(o.UiLanguage) {
		return true
	}

	return false
}

// SetUiLanguage gets a reference to the given int32 and assigns it to the UiLanguage field.
func (o *UiConfigResource) SetUiLanguage(v int32) {
	o.UiLanguage = &v
}

// GetExpandAlbumByDefault returns the ExpandAlbumByDefault field value if set, zero value otherwise.
func (o *UiConfigResource) GetExpandAlbumByDefault() bool {
	if o == nil || IsNil(o.ExpandAlbumByDefault) {
		var ret bool
		return ret
	}
	return *o.ExpandAlbumByDefault
}

// GetExpandAlbumByDefaultOk returns a tuple with the ExpandAlbumByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetExpandAlbumByDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpandAlbumByDefault) {
		return nil, false
	}
	return o.ExpandAlbumByDefault, true
}

// HasExpandAlbumByDefault returns a boolean if a field has been set.
func (o *UiConfigResource) HasExpandAlbumByDefault() bool {
	if o != nil && !IsNil(o.ExpandAlbumByDefault) {
		return true
	}

	return false
}

// SetExpandAlbumByDefault gets a reference to the given bool and assigns it to the ExpandAlbumByDefault field.
func (o *UiConfigResource) SetExpandAlbumByDefault(v bool) {
	o.ExpandAlbumByDefault = &v
}

// GetExpandSingleByDefault returns the ExpandSingleByDefault field value if set, zero value otherwise.
func (o *UiConfigResource) GetExpandSingleByDefault() bool {
	if o == nil || IsNil(o.ExpandSingleByDefault) {
		var ret bool
		return ret
	}
	return *o.ExpandSingleByDefault
}

// GetExpandSingleByDefaultOk returns a tuple with the ExpandSingleByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetExpandSingleByDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpandSingleByDefault) {
		return nil, false
	}
	return o.ExpandSingleByDefault, true
}

// HasExpandSingleByDefault returns a boolean if a field has been set.
func (o *UiConfigResource) HasExpandSingleByDefault() bool {
	if o != nil && !IsNil(o.ExpandSingleByDefault) {
		return true
	}

	return false
}

// SetExpandSingleByDefault gets a reference to the given bool and assigns it to the ExpandSingleByDefault field.
func (o *UiConfigResource) SetExpandSingleByDefault(v bool) {
	o.ExpandSingleByDefault = &v
}

// GetExpandEPByDefault returns the ExpandEPByDefault field value if set, zero value otherwise.
func (o *UiConfigResource) GetExpandEPByDefault() bool {
	if o == nil || IsNil(o.ExpandEPByDefault) {
		var ret bool
		return ret
	}
	return *o.ExpandEPByDefault
}

// GetExpandEPByDefaultOk returns a tuple with the ExpandEPByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetExpandEPByDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpandEPByDefault) {
		return nil, false
	}
	return o.ExpandEPByDefault, true
}

// HasExpandEPByDefault returns a boolean if a field has been set.
func (o *UiConfigResource) HasExpandEPByDefault() bool {
	if o != nil && !IsNil(o.ExpandEPByDefault) {
		return true
	}

	return false
}

// SetExpandEPByDefault gets a reference to the given bool and assigns it to the ExpandEPByDefault field.
func (o *UiConfigResource) SetExpandEPByDefault(v bool) {
	o.ExpandEPByDefault = &v
}

// GetExpandBroadcastByDefault returns the ExpandBroadcastByDefault field value if set, zero value otherwise.
func (o *UiConfigResource) GetExpandBroadcastByDefault() bool {
	if o == nil || IsNil(o.ExpandBroadcastByDefault) {
		var ret bool
		return ret
	}
	return *o.ExpandBroadcastByDefault
}

// GetExpandBroadcastByDefaultOk returns a tuple with the ExpandBroadcastByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetExpandBroadcastByDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpandBroadcastByDefault) {
		return nil, false
	}
	return o.ExpandBroadcastByDefault, true
}

// HasExpandBroadcastByDefault returns a boolean if a field has been set.
func (o *UiConfigResource) HasExpandBroadcastByDefault() bool {
	if o != nil && !IsNil(o.ExpandBroadcastByDefault) {
		return true
	}

	return false
}

// SetExpandBroadcastByDefault gets a reference to the given bool and assigns it to the ExpandBroadcastByDefault field.
func (o *UiConfigResource) SetExpandBroadcastByDefault(v bool) {
	o.ExpandBroadcastByDefault = &v
}

// GetExpandOtherByDefault returns the ExpandOtherByDefault field value if set, zero value otherwise.
func (o *UiConfigResource) GetExpandOtherByDefault() bool {
	if o == nil || IsNil(o.ExpandOtherByDefault) {
		var ret bool
		return ret
	}
	return *o.ExpandOtherByDefault
}

// GetExpandOtherByDefaultOk returns a tuple with the ExpandOtherByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetExpandOtherByDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpandOtherByDefault) {
		return nil, false
	}
	return o.ExpandOtherByDefault, true
}

// HasExpandOtherByDefault returns a boolean if a field has been set.
func (o *UiConfigResource) HasExpandOtherByDefault() bool {
	if o != nil && !IsNil(o.ExpandOtherByDefault) {
		return true
	}

	return false
}

// SetExpandOtherByDefault gets a reference to the given bool and assigns it to the ExpandOtherByDefault field.
func (o *UiConfigResource) SetExpandOtherByDefault(v bool) {
	o.ExpandOtherByDefault = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetTheme() string {
	if o == nil || IsNil(o.Theme.Get()) {
		var ret string
		return ret
	}
	return *o.Theme.Get()
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Theme.Get(), o.Theme.IsSet()
}

// HasTheme returns a boolean if a field has been set.
func (o *UiConfigResource) HasTheme() bool {
	if o != nil && o.Theme.IsSet() {
		return true
	}

	return false
}

// SetTheme gets a reference to the given NullableString and assigns it to the Theme field.
func (o *UiConfigResource) SetTheme(v string) {
	o.Theme.Set(&v)
}
// SetThemeNil sets the value for Theme to be an explicit nil
func (o *UiConfigResource) SetThemeNil() {
	o.Theme.Set(nil)
}

// UnsetTheme ensures that no value is present for Theme, not even an explicit nil
func (o *UiConfigResource) UnsetTheme() {
	o.Theme.Unset()
}

func (o UiConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FirstDayOfWeek) {
		toSerialize["firstDayOfWeek"] = o.FirstDayOfWeek
	}
	if o.CalendarWeekColumnHeader.IsSet() {
		toSerialize["calendarWeekColumnHeader"] = o.CalendarWeekColumnHeader.Get()
	}
	if o.ShortDateFormat.IsSet() {
		toSerialize["shortDateFormat"] = o.ShortDateFormat.Get()
	}
	if o.LongDateFormat.IsSet() {
		toSerialize["longDateFormat"] = o.LongDateFormat.Get()
	}
	if o.TimeFormat.IsSet() {
		toSerialize["timeFormat"] = o.TimeFormat.Get()
	}
	if !IsNil(o.ShowRelativeDates) {
		toSerialize["showRelativeDates"] = o.ShowRelativeDates
	}
	if !IsNil(o.EnableColorImpairedMode) {
		toSerialize["enableColorImpairedMode"] = o.EnableColorImpairedMode
	}
	if !IsNil(o.UiLanguage) {
		toSerialize["uiLanguage"] = o.UiLanguage
	}
	if !IsNil(o.ExpandAlbumByDefault) {
		toSerialize["expandAlbumByDefault"] = o.ExpandAlbumByDefault
	}
	if !IsNil(o.ExpandSingleByDefault) {
		toSerialize["expandSingleByDefault"] = o.ExpandSingleByDefault
	}
	if !IsNil(o.ExpandEPByDefault) {
		toSerialize["expandEPByDefault"] = o.ExpandEPByDefault
	}
	if !IsNil(o.ExpandBroadcastByDefault) {
		toSerialize["expandBroadcastByDefault"] = o.ExpandBroadcastByDefault
	}
	if !IsNil(o.ExpandOtherByDefault) {
		toSerialize["expandOtherByDefault"] = o.ExpandOtherByDefault
	}
	if o.Theme.IsSet() {
		toSerialize["theme"] = o.Theme.Get()
	}
	return toSerialize, nil
}

type NullableUiConfigResource struct {
	value *UiConfigResource
	isSet bool
}

func (v NullableUiConfigResource) Get() *UiConfigResource {
	return v.value
}

func (v *NullableUiConfigResource) Set(val *UiConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableUiConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableUiConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiConfigResource(val *UiConfigResource) *NullableUiConfigResource {
	return &NullableUiConfigResource{value: val, isSet: true}
}

func (v NullableUiConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


