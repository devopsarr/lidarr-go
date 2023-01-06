# AlbumRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**ForeignReleaseId** | Pointer to **NullableString** |  | [optional] 
**OldForeignReleaseIds** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **[]string** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **[]string** |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 
**Media** | Pointer to [**[]Medium**](Medium.md) |  | [optional] 
**TrackCount** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**Album** | Pointer to [**AlbumLazyLoaded**](AlbumLazyLoaded.md) |  | [optional] 
**Tracks** | Pointer to [**TrackListLazyLoaded**](TrackListLazyLoaded.md) |  | [optional] 

## Methods

### NewAlbumRelease

`func NewAlbumRelease() *AlbumRelease`

NewAlbumRelease instantiates a new AlbumRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumReleaseWithDefaults

`func NewAlbumReleaseWithDefaults() *AlbumRelease`

NewAlbumReleaseWithDefaults instantiates a new AlbumRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlbumRelease) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlbumRelease) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlbumRelease) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlbumRelease) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlbumId

`func (o *AlbumRelease) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *AlbumRelease) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *AlbumRelease) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *AlbumRelease) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetForeignReleaseId

`func (o *AlbumRelease) GetForeignReleaseId() string`

GetForeignReleaseId returns the ForeignReleaseId field if non-nil, zero value otherwise.

### GetForeignReleaseIdOk

`func (o *AlbumRelease) GetForeignReleaseIdOk() (*string, bool)`

GetForeignReleaseIdOk returns a tuple with the ForeignReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignReleaseId

`func (o *AlbumRelease) SetForeignReleaseId(v string)`

SetForeignReleaseId sets ForeignReleaseId field to given value.

### HasForeignReleaseId

`func (o *AlbumRelease) HasForeignReleaseId() bool`

HasForeignReleaseId returns a boolean if a field has been set.

### SetForeignReleaseIdNil

`func (o *AlbumRelease) SetForeignReleaseIdNil(b bool)`

 SetForeignReleaseIdNil sets the value for ForeignReleaseId to be an explicit nil

### UnsetForeignReleaseId
`func (o *AlbumRelease) UnsetForeignReleaseId()`

UnsetForeignReleaseId ensures that no value is present for ForeignReleaseId, not even an explicit nil
### GetOldForeignReleaseIds

`func (o *AlbumRelease) GetOldForeignReleaseIds() []string`

GetOldForeignReleaseIds returns the OldForeignReleaseIds field if non-nil, zero value otherwise.

### GetOldForeignReleaseIdsOk

`func (o *AlbumRelease) GetOldForeignReleaseIdsOk() (*[]string, bool)`

GetOldForeignReleaseIdsOk returns a tuple with the OldForeignReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldForeignReleaseIds

`func (o *AlbumRelease) SetOldForeignReleaseIds(v []string)`

SetOldForeignReleaseIds sets OldForeignReleaseIds field to given value.

### HasOldForeignReleaseIds

`func (o *AlbumRelease) HasOldForeignReleaseIds() bool`

HasOldForeignReleaseIds returns a boolean if a field has been set.

### SetOldForeignReleaseIdsNil

`func (o *AlbumRelease) SetOldForeignReleaseIdsNil(b bool)`

 SetOldForeignReleaseIdsNil sets the value for OldForeignReleaseIds to be an explicit nil

### UnsetOldForeignReleaseIds
`func (o *AlbumRelease) UnsetOldForeignReleaseIds()`

UnsetOldForeignReleaseIds ensures that no value is present for OldForeignReleaseIds, not even an explicit nil
### GetTitle

`func (o *AlbumRelease) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlbumRelease) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlbumRelease) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlbumRelease) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AlbumRelease) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AlbumRelease) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetStatus

`func (o *AlbumRelease) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlbumRelease) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlbumRelease) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlbumRelease) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AlbumRelease) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AlbumRelease) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDuration

`func (o *AlbumRelease) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AlbumRelease) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AlbumRelease) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AlbumRelease) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLabel

`func (o *AlbumRelease) GetLabel() []string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AlbumRelease) GetLabelOk() (*[]string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AlbumRelease) SetLabel(v []string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AlbumRelease) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *AlbumRelease) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *AlbumRelease) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisambiguation

`func (o *AlbumRelease) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *AlbumRelease) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *AlbumRelease) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *AlbumRelease) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *AlbumRelease) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *AlbumRelease) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetCountry

`func (o *AlbumRelease) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AlbumRelease) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AlbumRelease) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AlbumRelease) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AlbumRelease) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AlbumRelease) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetReleaseDate

`func (o *AlbumRelease) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AlbumRelease) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AlbumRelease) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AlbumRelease) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *AlbumRelease) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *AlbumRelease) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetMedia

`func (o *AlbumRelease) GetMedia() []Medium`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *AlbumRelease) GetMediaOk() (*[]Medium, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *AlbumRelease) SetMedia(v []Medium)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *AlbumRelease) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *AlbumRelease) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *AlbumRelease) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetTrackCount

`func (o *AlbumRelease) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *AlbumRelease) GetTrackCountOk() (*int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackCount

`func (o *AlbumRelease) SetTrackCount(v int32)`

SetTrackCount sets TrackCount field to given value.

### HasTrackCount

`func (o *AlbumRelease) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### GetMonitored

`func (o *AlbumRelease) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AlbumRelease) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AlbumRelease) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AlbumRelease) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetAlbum

`func (o *AlbumRelease) GetAlbum() AlbumLazyLoaded`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *AlbumRelease) GetAlbumOk() (*AlbumLazyLoaded, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *AlbumRelease) SetAlbum(v AlbumLazyLoaded)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *AlbumRelease) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetTracks

`func (o *AlbumRelease) GetTracks() TrackListLazyLoaded`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *AlbumRelease) GetTracksOk() (*TrackListLazyLoaded, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *AlbumRelease) SetTracks(v TrackListLazyLoaded)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *AlbumRelease) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


