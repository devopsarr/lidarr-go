# ParsedTrackInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**ArtistTitle** | Pointer to **NullableString** |  | [optional] 
**AlbumTitle** | Pointer to **NullableString** |  | [optional] 
**ArtistTitleInfo** | Pointer to [**ArtistTitleInfo**](ArtistTitleInfo.md) |  | [optional] 
**ArtistMBId** | Pointer to **NullableString** |  | [optional] 
**AlbumMBId** | Pointer to **NullableString** |  | [optional] 
**ReleaseMBId** | Pointer to **NullableString** |  | [optional] 
**RecordingMBId** | Pointer to **NullableString** |  | [optional] 
**TrackMBId** | Pointer to **NullableString** |  | [optional] 
**DiscNumber** | Pointer to **int32** |  | [optional] 
**DiscCount** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to [**IsoCountry**](IsoCountry.md) |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**CatalogNumber** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to [**TimeSpan**](TimeSpan.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoModel**](MediaInfoModel.md) |  | [optional] 
**TrackNumbers** | Pointer to **[]int32** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewParsedTrackInfo

`func NewParsedTrackInfo() *ParsedTrackInfo`

NewParsedTrackInfo instantiates a new ParsedTrackInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedTrackInfoWithDefaults

`func NewParsedTrackInfoWithDefaults() *ParsedTrackInfo`

NewParsedTrackInfoWithDefaults instantiates a new ParsedTrackInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ParsedTrackInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParsedTrackInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParsedTrackInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ParsedTrackInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ParsedTrackInfo) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ParsedTrackInfo) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *ParsedTrackInfo) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *ParsedTrackInfo) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *ParsedTrackInfo) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *ParsedTrackInfo) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *ParsedTrackInfo) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *ParsedTrackInfo) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetArtistTitle

`func (o *ParsedTrackInfo) GetArtistTitle() string`

GetArtistTitle returns the ArtistTitle field if non-nil, zero value otherwise.

### GetArtistTitleOk

`func (o *ParsedTrackInfo) GetArtistTitleOk() (*string, bool)`

GetArtistTitleOk returns a tuple with the ArtistTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistTitle

`func (o *ParsedTrackInfo) SetArtistTitle(v string)`

SetArtistTitle sets ArtistTitle field to given value.

### HasArtistTitle

`func (o *ParsedTrackInfo) HasArtistTitle() bool`

HasArtistTitle returns a boolean if a field has been set.

### SetArtistTitleNil

`func (o *ParsedTrackInfo) SetArtistTitleNil(b bool)`

 SetArtistTitleNil sets the value for ArtistTitle to be an explicit nil

### UnsetArtistTitle
`func (o *ParsedTrackInfo) UnsetArtistTitle()`

UnsetArtistTitle ensures that no value is present for ArtistTitle, not even an explicit nil
### GetAlbumTitle

`func (o *ParsedTrackInfo) GetAlbumTitle() string`

GetAlbumTitle returns the AlbumTitle field if non-nil, zero value otherwise.

### GetAlbumTitleOk

`func (o *ParsedTrackInfo) GetAlbumTitleOk() (*string, bool)`

GetAlbumTitleOk returns a tuple with the AlbumTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumTitle

`func (o *ParsedTrackInfo) SetAlbumTitle(v string)`

SetAlbumTitle sets AlbumTitle field to given value.

### HasAlbumTitle

`func (o *ParsedTrackInfo) HasAlbumTitle() bool`

HasAlbumTitle returns a boolean if a field has been set.

### SetAlbumTitleNil

`func (o *ParsedTrackInfo) SetAlbumTitleNil(b bool)`

 SetAlbumTitleNil sets the value for AlbumTitle to be an explicit nil

### UnsetAlbumTitle
`func (o *ParsedTrackInfo) UnsetAlbumTitle()`

UnsetAlbumTitle ensures that no value is present for AlbumTitle, not even an explicit nil
### GetArtistTitleInfo

`func (o *ParsedTrackInfo) GetArtistTitleInfo() ArtistTitleInfo`

GetArtistTitleInfo returns the ArtistTitleInfo field if non-nil, zero value otherwise.

### GetArtistTitleInfoOk

`func (o *ParsedTrackInfo) GetArtistTitleInfoOk() (*ArtistTitleInfo, bool)`

GetArtistTitleInfoOk returns a tuple with the ArtistTitleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistTitleInfo

`func (o *ParsedTrackInfo) SetArtistTitleInfo(v ArtistTitleInfo)`

SetArtistTitleInfo sets ArtistTitleInfo field to given value.

### HasArtistTitleInfo

`func (o *ParsedTrackInfo) HasArtistTitleInfo() bool`

HasArtistTitleInfo returns a boolean if a field has been set.

### GetArtistMBId

`func (o *ParsedTrackInfo) GetArtistMBId() string`

GetArtistMBId returns the ArtistMBId field if non-nil, zero value otherwise.

### GetArtistMBIdOk

`func (o *ParsedTrackInfo) GetArtistMBIdOk() (*string, bool)`

GetArtistMBIdOk returns a tuple with the ArtistMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMBId

`func (o *ParsedTrackInfo) SetArtistMBId(v string)`

SetArtistMBId sets ArtistMBId field to given value.

### HasArtistMBId

`func (o *ParsedTrackInfo) HasArtistMBId() bool`

HasArtistMBId returns a boolean if a field has been set.

### SetArtistMBIdNil

`func (o *ParsedTrackInfo) SetArtistMBIdNil(b bool)`

 SetArtistMBIdNil sets the value for ArtistMBId to be an explicit nil

### UnsetArtistMBId
`func (o *ParsedTrackInfo) UnsetArtistMBId()`

UnsetArtistMBId ensures that no value is present for ArtistMBId, not even an explicit nil
### GetAlbumMBId

`func (o *ParsedTrackInfo) GetAlbumMBId() string`

GetAlbumMBId returns the AlbumMBId field if non-nil, zero value otherwise.

### GetAlbumMBIdOk

`func (o *ParsedTrackInfo) GetAlbumMBIdOk() (*string, bool)`

GetAlbumMBIdOk returns a tuple with the AlbumMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumMBId

`func (o *ParsedTrackInfo) SetAlbumMBId(v string)`

SetAlbumMBId sets AlbumMBId field to given value.

### HasAlbumMBId

`func (o *ParsedTrackInfo) HasAlbumMBId() bool`

HasAlbumMBId returns a boolean if a field has been set.

### SetAlbumMBIdNil

`func (o *ParsedTrackInfo) SetAlbumMBIdNil(b bool)`

 SetAlbumMBIdNil sets the value for AlbumMBId to be an explicit nil

### UnsetAlbumMBId
`func (o *ParsedTrackInfo) UnsetAlbumMBId()`

UnsetAlbumMBId ensures that no value is present for AlbumMBId, not even an explicit nil
### GetReleaseMBId

`func (o *ParsedTrackInfo) GetReleaseMBId() string`

GetReleaseMBId returns the ReleaseMBId field if non-nil, zero value otherwise.

### GetReleaseMBIdOk

`func (o *ParsedTrackInfo) GetReleaseMBIdOk() (*string, bool)`

GetReleaseMBIdOk returns a tuple with the ReleaseMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMBId

`func (o *ParsedTrackInfo) SetReleaseMBId(v string)`

SetReleaseMBId sets ReleaseMBId field to given value.

### HasReleaseMBId

`func (o *ParsedTrackInfo) HasReleaseMBId() bool`

HasReleaseMBId returns a boolean if a field has been set.

### SetReleaseMBIdNil

`func (o *ParsedTrackInfo) SetReleaseMBIdNil(b bool)`

 SetReleaseMBIdNil sets the value for ReleaseMBId to be an explicit nil

### UnsetReleaseMBId
`func (o *ParsedTrackInfo) UnsetReleaseMBId()`

UnsetReleaseMBId ensures that no value is present for ReleaseMBId, not even an explicit nil
### GetRecordingMBId

`func (o *ParsedTrackInfo) GetRecordingMBId() string`

GetRecordingMBId returns the RecordingMBId field if non-nil, zero value otherwise.

### GetRecordingMBIdOk

`func (o *ParsedTrackInfo) GetRecordingMBIdOk() (*string, bool)`

GetRecordingMBIdOk returns a tuple with the RecordingMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingMBId

`func (o *ParsedTrackInfo) SetRecordingMBId(v string)`

SetRecordingMBId sets RecordingMBId field to given value.

### HasRecordingMBId

`func (o *ParsedTrackInfo) HasRecordingMBId() bool`

HasRecordingMBId returns a boolean if a field has been set.

### SetRecordingMBIdNil

`func (o *ParsedTrackInfo) SetRecordingMBIdNil(b bool)`

 SetRecordingMBIdNil sets the value for RecordingMBId to be an explicit nil

### UnsetRecordingMBId
`func (o *ParsedTrackInfo) UnsetRecordingMBId()`

UnsetRecordingMBId ensures that no value is present for RecordingMBId, not even an explicit nil
### GetTrackMBId

`func (o *ParsedTrackInfo) GetTrackMBId() string`

GetTrackMBId returns the TrackMBId field if non-nil, zero value otherwise.

### GetTrackMBIdOk

`func (o *ParsedTrackInfo) GetTrackMBIdOk() (*string, bool)`

GetTrackMBIdOk returns a tuple with the TrackMBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackMBId

`func (o *ParsedTrackInfo) SetTrackMBId(v string)`

SetTrackMBId sets TrackMBId field to given value.

### HasTrackMBId

`func (o *ParsedTrackInfo) HasTrackMBId() bool`

HasTrackMBId returns a boolean if a field has been set.

### SetTrackMBIdNil

`func (o *ParsedTrackInfo) SetTrackMBIdNil(b bool)`

 SetTrackMBIdNil sets the value for TrackMBId to be an explicit nil

### UnsetTrackMBId
`func (o *ParsedTrackInfo) UnsetTrackMBId()`

UnsetTrackMBId ensures that no value is present for TrackMBId, not even an explicit nil
### GetDiscNumber

`func (o *ParsedTrackInfo) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *ParsedTrackInfo) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *ParsedTrackInfo) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *ParsedTrackInfo) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDiscCount

`func (o *ParsedTrackInfo) GetDiscCount() int32`

GetDiscCount returns the DiscCount field if non-nil, zero value otherwise.

### GetDiscCountOk

`func (o *ParsedTrackInfo) GetDiscCountOk() (*int32, bool)`

GetDiscCountOk returns a tuple with the DiscCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscCount

`func (o *ParsedTrackInfo) SetDiscCount(v int32)`

SetDiscCount sets DiscCount field to given value.

### HasDiscCount

`func (o *ParsedTrackInfo) HasDiscCount() bool`

HasDiscCount returns a boolean if a field has been set.

### GetCountry

`func (o *ParsedTrackInfo) GetCountry() IsoCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ParsedTrackInfo) GetCountryOk() (*IsoCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ParsedTrackInfo) SetCountry(v IsoCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ParsedTrackInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetYear

`func (o *ParsedTrackInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ParsedTrackInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ParsedTrackInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ParsedTrackInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetLabel

`func (o *ParsedTrackInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ParsedTrackInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ParsedTrackInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ParsedTrackInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ParsedTrackInfo) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ParsedTrackInfo) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetCatalogNumber

`func (o *ParsedTrackInfo) GetCatalogNumber() string`

GetCatalogNumber returns the CatalogNumber field if non-nil, zero value otherwise.

### GetCatalogNumberOk

`func (o *ParsedTrackInfo) GetCatalogNumberOk() (*string, bool)`

GetCatalogNumberOk returns a tuple with the CatalogNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogNumber

`func (o *ParsedTrackInfo) SetCatalogNumber(v string)`

SetCatalogNumber sets CatalogNumber field to given value.

### HasCatalogNumber

`func (o *ParsedTrackInfo) HasCatalogNumber() bool`

HasCatalogNumber returns a boolean if a field has been set.

### SetCatalogNumberNil

`func (o *ParsedTrackInfo) SetCatalogNumberNil(b bool)`

 SetCatalogNumberNil sets the value for CatalogNumber to be an explicit nil

### UnsetCatalogNumber
`func (o *ParsedTrackInfo) UnsetCatalogNumber()`

UnsetCatalogNumber ensures that no value is present for CatalogNumber, not even an explicit nil
### GetDisambiguation

`func (o *ParsedTrackInfo) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *ParsedTrackInfo) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *ParsedTrackInfo) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *ParsedTrackInfo) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *ParsedTrackInfo) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *ParsedTrackInfo) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetDuration

`func (o *ParsedTrackInfo) GetDuration() TimeSpan`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ParsedTrackInfo) GetDurationOk() (*TimeSpan, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ParsedTrackInfo) SetDuration(v TimeSpan)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ParsedTrackInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetQuality

`func (o *ParsedTrackInfo) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ParsedTrackInfo) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ParsedTrackInfo) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ParsedTrackInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetMediaInfo

`func (o *ParsedTrackInfo) GetMediaInfo() MediaInfoModel`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *ParsedTrackInfo) GetMediaInfoOk() (*MediaInfoModel, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *ParsedTrackInfo) SetMediaInfo(v MediaInfoModel)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *ParsedTrackInfo) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetTrackNumbers

`func (o *ParsedTrackInfo) GetTrackNumbers() []int32`

GetTrackNumbers returns the TrackNumbers field if non-nil, zero value otherwise.

### GetTrackNumbersOk

`func (o *ParsedTrackInfo) GetTrackNumbersOk() (*[]int32, bool)`

GetTrackNumbersOk returns a tuple with the TrackNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumbers

`func (o *ParsedTrackInfo) SetTrackNumbers(v []int32)`

SetTrackNumbers sets TrackNumbers field to given value.

### HasTrackNumbers

`func (o *ParsedTrackInfo) HasTrackNumbers() bool`

HasTrackNumbers returns a boolean if a field has been set.

### SetTrackNumbersNil

`func (o *ParsedTrackInfo) SetTrackNumbersNil(b bool)`

 SetTrackNumbersNil sets the value for TrackNumbers to be an explicit nil

### UnsetTrackNumbers
`func (o *ParsedTrackInfo) UnsetTrackNumbers()`

UnsetTrackNumbers ensures that no value is present for TrackNumbers, not even an explicit nil
### GetReleaseGroup

`func (o *ParsedTrackInfo) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ParsedTrackInfo) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ParsedTrackInfo) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ParsedTrackInfo) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ParsedTrackInfo) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ParsedTrackInfo) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetReleaseHash

`func (o *ParsedTrackInfo) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ParsedTrackInfo) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ParsedTrackInfo) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ParsedTrackInfo) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ParsedTrackInfo) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ParsedTrackInfo) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


