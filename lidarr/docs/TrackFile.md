# TrackFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**DateAdded** | Pointer to **time.Time** |  | [optional] 
**OriginalFilePath** | Pointer to **NullableString** |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoModel**](MediaInfoModel.md) |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**Tracks** | Pointer to [**TrackListLazyLoaded**](TrackListLazyLoaded.md) |  | [optional] 
**Artist** | Pointer to [**ArtistLazyLoaded**](ArtistLazyLoaded.md) |  | [optional] 
**Album** | Pointer to [**AlbumLazyLoaded**](AlbumLazyLoaded.md) |  | [optional] 

## Methods

### NewTrackFile

`func NewTrackFile() *TrackFile`

NewTrackFile instantiates a new TrackFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackFileWithDefaults

`func NewTrackFileWithDefaults() *TrackFile`

NewTrackFileWithDefaults instantiates a new TrackFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrackFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *TrackFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TrackFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TrackFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TrackFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *TrackFile) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *TrackFile) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *TrackFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrackFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrackFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrackFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetModified

`func (o *TrackFile) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TrackFile) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TrackFile) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *TrackFile) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDateAdded

`func (o *TrackFile) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *TrackFile) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *TrackFile) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *TrackFile) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetOriginalFilePath

`func (o *TrackFile) GetOriginalFilePath() string`

GetOriginalFilePath returns the OriginalFilePath field if non-nil, zero value otherwise.

### GetOriginalFilePathOk

`func (o *TrackFile) GetOriginalFilePathOk() (*string, bool)`

GetOriginalFilePathOk returns a tuple with the OriginalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFilePath

`func (o *TrackFile) SetOriginalFilePath(v string)`

SetOriginalFilePath sets OriginalFilePath field to given value.

### HasOriginalFilePath

`func (o *TrackFile) HasOriginalFilePath() bool`

HasOriginalFilePath returns a boolean if a field has been set.

### SetOriginalFilePathNil

`func (o *TrackFile) SetOriginalFilePathNil(b bool)`

 SetOriginalFilePathNil sets the value for OriginalFilePath to be an explicit nil

### UnsetOriginalFilePath
`func (o *TrackFile) UnsetOriginalFilePath()`

UnsetOriginalFilePath ensures that no value is present for OriginalFilePath, not even an explicit nil
### GetSceneName

`func (o *TrackFile) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *TrackFile) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *TrackFile) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *TrackFile) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *TrackFile) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *TrackFile) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetReleaseGroup

`func (o *TrackFile) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *TrackFile) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *TrackFile) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *TrackFile) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *TrackFile) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *TrackFile) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetQuality

`func (o *TrackFile) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *TrackFile) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *TrackFile) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *TrackFile) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetMediaInfo

`func (o *TrackFile) GetMediaInfo() MediaInfoModel`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *TrackFile) GetMediaInfoOk() (*MediaInfoModel, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *TrackFile) SetMediaInfo(v MediaInfoModel)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *TrackFile) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetAlbumId

`func (o *TrackFile) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *TrackFile) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *TrackFile) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *TrackFile) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetTracks

`func (o *TrackFile) GetTracks() TrackListLazyLoaded`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *TrackFile) GetTracksOk() (*TrackListLazyLoaded, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *TrackFile) SetTracks(v TrackListLazyLoaded)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *TrackFile) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### GetArtist

`func (o *TrackFile) GetArtist() ArtistLazyLoaded`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *TrackFile) GetArtistOk() (*ArtistLazyLoaded, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *TrackFile) SetArtist(v ArtistLazyLoaded)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *TrackFile) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetAlbum

`func (o *TrackFile) GetAlbum() AlbumLazyLoaded`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *TrackFile) GetAlbumOk() (*AlbumLazyLoaded, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *TrackFile) SetAlbum(v AlbumLazyLoaded)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *TrackFile) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


