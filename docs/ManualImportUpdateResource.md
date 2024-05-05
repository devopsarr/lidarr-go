# ManualImportUpdateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ArtistId** | Pointer to **NullableInt32** |  | [optional] 
**AlbumId** | Pointer to **NullableInt32** |  | [optional] 
**AlbumReleaseId** | Pointer to **NullableInt32** |  | [optional] 
**Tracks** | Pointer to [**[]TrackResource**](TrackResource.md) |  | [optional] 
**TrackIds** | Pointer to **[]int32** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**IndexerFlags** | Pointer to **int32** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFile** | Pointer to **bool** |  | [optional] 
**ReplaceExistingFiles** | Pointer to **bool** |  | [optional] 
**DisableReleaseSwitching** | Pointer to **bool** |  | [optional] 
**Rejections** | Pointer to [**[]Rejection**](Rejection.md) |  | [optional] 

## Methods

### NewManualImportUpdateResource

`func NewManualImportUpdateResource() *ManualImportUpdateResource`

NewManualImportUpdateResource instantiates a new ManualImportUpdateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualImportUpdateResourceWithDefaults

`func NewManualImportUpdateResourceWithDefaults() *ManualImportUpdateResource`

NewManualImportUpdateResourceWithDefaults instantiates a new ManualImportUpdateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualImportUpdateResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualImportUpdateResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualImportUpdateResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualImportUpdateResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ManualImportUpdateResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManualImportUpdateResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManualImportUpdateResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManualImportUpdateResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ManualImportUpdateResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ManualImportUpdateResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetName

`func (o *ManualImportUpdateResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManualImportUpdateResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManualImportUpdateResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManualImportUpdateResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ManualImportUpdateResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ManualImportUpdateResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetArtistId

`func (o *ManualImportUpdateResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *ManualImportUpdateResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *ManualImportUpdateResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *ManualImportUpdateResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### SetArtistIdNil

`func (o *ManualImportUpdateResource) SetArtistIdNil(b bool)`

 SetArtistIdNil sets the value for ArtistId to be an explicit nil

### UnsetArtistId
`func (o *ManualImportUpdateResource) UnsetArtistId()`

UnsetArtistId ensures that no value is present for ArtistId, not even an explicit nil
### GetAlbumId

`func (o *ManualImportUpdateResource) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *ManualImportUpdateResource) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *ManualImportUpdateResource) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *ManualImportUpdateResource) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### SetAlbumIdNil

`func (o *ManualImportUpdateResource) SetAlbumIdNil(b bool)`

 SetAlbumIdNil sets the value for AlbumId to be an explicit nil

### UnsetAlbumId
`func (o *ManualImportUpdateResource) UnsetAlbumId()`

UnsetAlbumId ensures that no value is present for AlbumId, not even an explicit nil
### GetAlbumReleaseId

`func (o *ManualImportUpdateResource) GetAlbumReleaseId() int32`

GetAlbumReleaseId returns the AlbumReleaseId field if non-nil, zero value otherwise.

### GetAlbumReleaseIdOk

`func (o *ManualImportUpdateResource) GetAlbumReleaseIdOk() (*int32, bool)`

GetAlbumReleaseIdOk returns a tuple with the AlbumReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumReleaseId

`func (o *ManualImportUpdateResource) SetAlbumReleaseId(v int32)`

SetAlbumReleaseId sets AlbumReleaseId field to given value.

### HasAlbumReleaseId

`func (o *ManualImportUpdateResource) HasAlbumReleaseId() bool`

HasAlbumReleaseId returns a boolean if a field has been set.

### SetAlbumReleaseIdNil

`func (o *ManualImportUpdateResource) SetAlbumReleaseIdNil(b bool)`

 SetAlbumReleaseIdNil sets the value for AlbumReleaseId to be an explicit nil

### UnsetAlbumReleaseId
`func (o *ManualImportUpdateResource) UnsetAlbumReleaseId()`

UnsetAlbumReleaseId ensures that no value is present for AlbumReleaseId, not even an explicit nil
### GetTracks

`func (o *ManualImportUpdateResource) GetTracks() []TrackResource`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *ManualImportUpdateResource) GetTracksOk() (*[]TrackResource, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *ManualImportUpdateResource) SetTracks(v []TrackResource)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *ManualImportUpdateResource) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### SetTracksNil

`func (o *ManualImportUpdateResource) SetTracksNil(b bool)`

 SetTracksNil sets the value for Tracks to be an explicit nil

### UnsetTracks
`func (o *ManualImportUpdateResource) UnsetTracks()`

UnsetTracks ensures that no value is present for Tracks, not even an explicit nil
### GetTrackIds

`func (o *ManualImportUpdateResource) GetTrackIds() []int32`

GetTrackIds returns the TrackIds field if non-nil, zero value otherwise.

### GetTrackIdsOk

`func (o *ManualImportUpdateResource) GetTrackIdsOk() (*[]int32, bool)`

GetTrackIdsOk returns a tuple with the TrackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackIds

`func (o *ManualImportUpdateResource) SetTrackIds(v []int32)`

SetTrackIds sets TrackIds field to given value.

### HasTrackIds

`func (o *ManualImportUpdateResource) HasTrackIds() bool`

HasTrackIds returns a boolean if a field has been set.

### SetTrackIdsNil

`func (o *ManualImportUpdateResource) SetTrackIdsNil(b bool)`

 SetTrackIdsNil sets the value for TrackIds to be an explicit nil

### UnsetTrackIds
`func (o *ManualImportUpdateResource) UnsetTrackIds()`

UnsetTrackIds ensures that no value is present for TrackIds, not even an explicit nil
### GetQuality

`func (o *ManualImportUpdateResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ManualImportUpdateResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ManualImportUpdateResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ManualImportUpdateResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetReleaseGroup

`func (o *ManualImportUpdateResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ManualImportUpdateResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ManualImportUpdateResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ManualImportUpdateResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ManualImportUpdateResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ManualImportUpdateResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetIndexerFlags

`func (o *ManualImportUpdateResource) GetIndexerFlags() int32`

GetIndexerFlags returns the IndexerFlags field if non-nil, zero value otherwise.

### GetIndexerFlagsOk

`func (o *ManualImportUpdateResource) GetIndexerFlagsOk() (*int32, bool)`

GetIndexerFlagsOk returns a tuple with the IndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerFlags

`func (o *ManualImportUpdateResource) SetIndexerFlags(v int32)`

SetIndexerFlags sets IndexerFlags field to given value.

### HasIndexerFlags

`func (o *ManualImportUpdateResource) HasIndexerFlags() bool`

HasIndexerFlags returns a boolean if a field has been set.

### GetDownloadId

`func (o *ManualImportUpdateResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *ManualImportUpdateResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *ManualImportUpdateResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *ManualImportUpdateResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *ManualImportUpdateResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *ManualImportUpdateResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetAdditionalFile

`func (o *ManualImportUpdateResource) GetAdditionalFile() bool`

GetAdditionalFile returns the AdditionalFile field if non-nil, zero value otherwise.

### GetAdditionalFileOk

`func (o *ManualImportUpdateResource) GetAdditionalFileOk() (*bool, bool)`

GetAdditionalFileOk returns a tuple with the AdditionalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFile

`func (o *ManualImportUpdateResource) SetAdditionalFile(v bool)`

SetAdditionalFile sets AdditionalFile field to given value.

### HasAdditionalFile

`func (o *ManualImportUpdateResource) HasAdditionalFile() bool`

HasAdditionalFile returns a boolean if a field has been set.

### GetReplaceExistingFiles

`func (o *ManualImportUpdateResource) GetReplaceExistingFiles() bool`

GetReplaceExistingFiles returns the ReplaceExistingFiles field if non-nil, zero value otherwise.

### GetReplaceExistingFilesOk

`func (o *ManualImportUpdateResource) GetReplaceExistingFilesOk() (*bool, bool)`

GetReplaceExistingFilesOk returns a tuple with the ReplaceExistingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceExistingFiles

`func (o *ManualImportUpdateResource) SetReplaceExistingFiles(v bool)`

SetReplaceExistingFiles sets ReplaceExistingFiles field to given value.

### HasReplaceExistingFiles

`func (o *ManualImportUpdateResource) HasReplaceExistingFiles() bool`

HasReplaceExistingFiles returns a boolean if a field has been set.

### GetDisableReleaseSwitching

`func (o *ManualImportUpdateResource) GetDisableReleaseSwitching() bool`

GetDisableReleaseSwitching returns the DisableReleaseSwitching field if non-nil, zero value otherwise.

### GetDisableReleaseSwitchingOk

`func (o *ManualImportUpdateResource) GetDisableReleaseSwitchingOk() (*bool, bool)`

GetDisableReleaseSwitchingOk returns a tuple with the DisableReleaseSwitching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReleaseSwitching

`func (o *ManualImportUpdateResource) SetDisableReleaseSwitching(v bool)`

SetDisableReleaseSwitching sets DisableReleaseSwitching field to given value.

### HasDisableReleaseSwitching

`func (o *ManualImportUpdateResource) HasDisableReleaseSwitching() bool`

HasDisableReleaseSwitching returns a boolean if a field has been set.

### GetRejections

`func (o *ManualImportUpdateResource) GetRejections() []Rejection`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ManualImportUpdateResource) GetRejectionsOk() (*[]Rejection, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ManualImportUpdateResource) SetRejections(v []Rejection)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ManualImportUpdateResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ManualImportUpdateResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ManualImportUpdateResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


