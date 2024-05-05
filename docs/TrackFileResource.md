# TrackFileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistId** | Pointer to **int32** |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**DateAdded** | Pointer to **time.Time** |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**QualityWeight** | Pointer to **int32** |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**CustomFormatScore** | Pointer to **int32** |  | [optional] 
**IndexerFlags** | Pointer to **NullableInt32** |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoResource**](MediaInfoResource.md) |  | [optional] 
**QualityCutoffNotMet** | Pointer to **bool** |  | [optional] 
**AudioTags** | Pointer to [**ParsedTrackInfo**](ParsedTrackInfo.md) |  | [optional] 

## Methods

### NewTrackFileResource

`func NewTrackFileResource() *TrackFileResource`

NewTrackFileResource instantiates a new TrackFileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackFileResourceWithDefaults

`func NewTrackFileResourceWithDefaults() *TrackFileResource`

NewTrackFileResourceWithDefaults instantiates a new TrackFileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackFileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackFileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackFileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrackFileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistId

`func (o *TrackFileResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *TrackFileResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *TrackFileResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *TrackFileResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### GetAlbumId

`func (o *TrackFileResource) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *TrackFileResource) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *TrackFileResource) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *TrackFileResource) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetPath

`func (o *TrackFileResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TrackFileResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TrackFileResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TrackFileResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *TrackFileResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *TrackFileResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *TrackFileResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrackFileResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrackFileResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrackFileResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDateAdded

`func (o *TrackFileResource) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *TrackFileResource) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *TrackFileResource) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *TrackFileResource) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSceneName

`func (o *TrackFileResource) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *TrackFileResource) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *TrackFileResource) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *TrackFileResource) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *TrackFileResource) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *TrackFileResource) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetReleaseGroup

`func (o *TrackFileResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *TrackFileResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *TrackFileResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *TrackFileResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *TrackFileResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *TrackFileResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetQuality

`func (o *TrackFileResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *TrackFileResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *TrackFileResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *TrackFileResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetQualityWeight

`func (o *TrackFileResource) GetQualityWeight() int32`

GetQualityWeight returns the QualityWeight field if non-nil, zero value otherwise.

### GetQualityWeightOk

`func (o *TrackFileResource) GetQualityWeightOk() (*int32, bool)`

GetQualityWeightOk returns a tuple with the QualityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityWeight

`func (o *TrackFileResource) SetQualityWeight(v int32)`

SetQualityWeight sets QualityWeight field to given value.

### HasQualityWeight

`func (o *TrackFileResource) HasQualityWeight() bool`

HasQualityWeight returns a boolean if a field has been set.

### GetCustomFormats

`func (o *TrackFileResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *TrackFileResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *TrackFileResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *TrackFileResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *TrackFileResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *TrackFileResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetCustomFormatScore

`func (o *TrackFileResource) GetCustomFormatScore() int32`

GetCustomFormatScore returns the CustomFormatScore field if non-nil, zero value otherwise.

### GetCustomFormatScoreOk

`func (o *TrackFileResource) GetCustomFormatScoreOk() (*int32, bool)`

GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormatScore

`func (o *TrackFileResource) SetCustomFormatScore(v int32)`

SetCustomFormatScore sets CustomFormatScore field to given value.

### HasCustomFormatScore

`func (o *TrackFileResource) HasCustomFormatScore() bool`

HasCustomFormatScore returns a boolean if a field has been set.

### GetIndexerFlags

`func (o *TrackFileResource) GetIndexerFlags() int32`

GetIndexerFlags returns the IndexerFlags field if non-nil, zero value otherwise.

### GetIndexerFlagsOk

`func (o *TrackFileResource) GetIndexerFlagsOk() (*int32, bool)`

GetIndexerFlagsOk returns a tuple with the IndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerFlags

`func (o *TrackFileResource) SetIndexerFlags(v int32)`

SetIndexerFlags sets IndexerFlags field to given value.

### HasIndexerFlags

`func (o *TrackFileResource) HasIndexerFlags() bool`

HasIndexerFlags returns a boolean if a field has been set.

### SetIndexerFlagsNil

`func (o *TrackFileResource) SetIndexerFlagsNil(b bool)`

 SetIndexerFlagsNil sets the value for IndexerFlags to be an explicit nil

### UnsetIndexerFlags
`func (o *TrackFileResource) UnsetIndexerFlags()`

UnsetIndexerFlags ensures that no value is present for IndexerFlags, not even an explicit nil
### GetMediaInfo

`func (o *TrackFileResource) GetMediaInfo() MediaInfoResource`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *TrackFileResource) GetMediaInfoOk() (*MediaInfoResource, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *TrackFileResource) SetMediaInfo(v MediaInfoResource)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *TrackFileResource) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetQualityCutoffNotMet

`func (o *TrackFileResource) GetQualityCutoffNotMet() bool`

GetQualityCutoffNotMet returns the QualityCutoffNotMet field if non-nil, zero value otherwise.

### GetQualityCutoffNotMetOk

`func (o *TrackFileResource) GetQualityCutoffNotMetOk() (*bool, bool)`

GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityCutoffNotMet

`func (o *TrackFileResource) SetQualityCutoffNotMet(v bool)`

SetQualityCutoffNotMet sets QualityCutoffNotMet field to given value.

### HasQualityCutoffNotMet

`func (o *TrackFileResource) HasQualityCutoffNotMet() bool`

HasQualityCutoffNotMet returns a boolean if a field has been set.

### GetAudioTags

`func (o *TrackFileResource) GetAudioTags() ParsedTrackInfo`

GetAudioTags returns the AudioTags field if non-nil, zero value otherwise.

### GetAudioTagsOk

`func (o *TrackFileResource) GetAudioTagsOk() (*ParsedTrackInfo, bool)`

GetAudioTagsOk returns a tuple with the AudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTags

`func (o *TrackFileResource) SetAudioTags(v ParsedTrackInfo)`

SetAudioTags sets AudioTags field to given value.

### HasAudioTags

`func (o *TrackFileResource) HasAudioTags() bool`

HasAudioTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


