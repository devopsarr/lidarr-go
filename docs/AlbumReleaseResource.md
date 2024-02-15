# AlbumReleaseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**ForeignReleaseId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**TrackCount** | Pointer to **int32** |  | [optional] 
**Media** | Pointer to [**[]MediumResource**](MediumResource.md) |  | [optional] 
**MediumCount** | Pointer to **int32** |  | [optional] [readonly] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **[]string** |  | [optional] 
**Label** | Pointer to **[]string** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 

## Methods

### NewAlbumReleaseResource

`func NewAlbumReleaseResource() *AlbumReleaseResource`

NewAlbumReleaseResource instantiates a new AlbumReleaseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumReleaseResourceWithDefaults

`func NewAlbumReleaseResourceWithDefaults() *AlbumReleaseResource`

NewAlbumReleaseResourceWithDefaults instantiates a new AlbumReleaseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlbumReleaseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlbumReleaseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlbumReleaseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlbumReleaseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlbumId

`func (o *AlbumReleaseResource) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *AlbumReleaseResource) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *AlbumReleaseResource) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *AlbumReleaseResource) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetForeignReleaseId

`func (o *AlbumReleaseResource) GetForeignReleaseId() string`

GetForeignReleaseId returns the ForeignReleaseId field if non-nil, zero value otherwise.

### GetForeignReleaseIdOk

`func (o *AlbumReleaseResource) GetForeignReleaseIdOk() (*string, bool)`

GetForeignReleaseIdOk returns a tuple with the ForeignReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignReleaseId

`func (o *AlbumReleaseResource) SetForeignReleaseId(v string)`

SetForeignReleaseId sets ForeignReleaseId field to given value.

### HasForeignReleaseId

`func (o *AlbumReleaseResource) HasForeignReleaseId() bool`

HasForeignReleaseId returns a boolean if a field has been set.

### SetForeignReleaseIdNil

`func (o *AlbumReleaseResource) SetForeignReleaseIdNil(b bool)`

 SetForeignReleaseIdNil sets the value for ForeignReleaseId to be an explicit nil

### UnsetForeignReleaseId
`func (o *AlbumReleaseResource) UnsetForeignReleaseId()`

UnsetForeignReleaseId ensures that no value is present for ForeignReleaseId, not even an explicit nil
### GetTitle

`func (o *AlbumReleaseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlbumReleaseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlbumReleaseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlbumReleaseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AlbumReleaseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AlbumReleaseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *AlbumReleaseResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlbumReleaseResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlbumReleaseResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlbumReleaseResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AlbumReleaseResource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AlbumReleaseResource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDuration

`func (o *AlbumReleaseResource) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AlbumReleaseResource) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AlbumReleaseResource) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AlbumReleaseResource) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTrackCount

`func (o *AlbumReleaseResource) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *AlbumReleaseResource) GetTrackCountOk() (*int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackCount

`func (o *AlbumReleaseResource) SetTrackCount(v int32)`

SetTrackCount sets TrackCount field to given value.

### HasTrackCount

`func (o *AlbumReleaseResource) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### GetMedia

`func (o *AlbumReleaseResource) GetMedia() []MediumResource`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *AlbumReleaseResource) GetMediaOk() (*[]MediumResource, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *AlbumReleaseResource) SetMedia(v []MediumResource)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *AlbumReleaseResource) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *AlbumReleaseResource) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *AlbumReleaseResource) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetMediumCount

`func (o *AlbumReleaseResource) GetMediumCount() int32`

GetMediumCount returns the MediumCount field if non-nil, zero value otherwise.

### GetMediumCountOk

`func (o *AlbumReleaseResource) GetMediumCountOk() (*int32, bool)`

GetMediumCountOk returns a tuple with the MediumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumCount

`func (o *AlbumReleaseResource) SetMediumCount(v int32)`

SetMediumCount sets MediumCount field to given value.

### HasMediumCount

`func (o *AlbumReleaseResource) HasMediumCount() bool`

HasMediumCount returns a boolean if a field has been set.

### GetDisambiguation

`func (o *AlbumReleaseResource) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *AlbumReleaseResource) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *AlbumReleaseResource) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *AlbumReleaseResource) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *AlbumReleaseResource) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *AlbumReleaseResource) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetCountry

`func (o *AlbumReleaseResource) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AlbumReleaseResource) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AlbumReleaseResource) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AlbumReleaseResource) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AlbumReleaseResource) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AlbumReleaseResource) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLabel

`func (o *AlbumReleaseResource) GetLabel() []string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AlbumReleaseResource) GetLabelOk() (*[]string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AlbumReleaseResource) SetLabel(v []string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AlbumReleaseResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *AlbumReleaseResource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *AlbumReleaseResource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetFormat

`func (o *AlbumReleaseResource) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AlbumReleaseResource) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AlbumReleaseResource) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AlbumReleaseResource) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *AlbumReleaseResource) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *AlbumReleaseResource) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetMonitored

`func (o *AlbumReleaseResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AlbumReleaseResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AlbumReleaseResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AlbumReleaseResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


