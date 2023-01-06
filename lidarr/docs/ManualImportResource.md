# ManualImportResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Artist** | Pointer to [**ArtistResource**](ArtistResource.md) |  | [optional] 
**Album** | Pointer to [**AlbumResource**](AlbumResource.md) |  | [optional] 
**AlbumReleaseId** | Pointer to **int32** |  | [optional] 
**Tracks** | Pointer to [**[]TrackResource**](TrackResource.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**QualityWeight** | Pointer to **int32** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**Rejections** | Pointer to [**[]Rejection**](Rejection.md) |  | [optional] 
**AudioTags** | Pointer to [**ParsedTrackInfo**](ParsedTrackInfo.md) |  | [optional] 
**AdditionalFile** | Pointer to **bool** |  | [optional] 
**ReplaceExistingFiles** | Pointer to **bool** |  | [optional] 
**DisableReleaseSwitching** | Pointer to **bool** |  | [optional] 

## Methods

### NewManualImportResource

`func NewManualImportResource() *ManualImportResource`

NewManualImportResource instantiates a new ManualImportResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualImportResourceWithDefaults

`func NewManualImportResourceWithDefaults() *ManualImportResource`

NewManualImportResourceWithDefaults instantiates a new ManualImportResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualImportResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualImportResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualImportResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualImportResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ManualImportResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManualImportResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManualImportResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManualImportResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ManualImportResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ManualImportResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetName

`func (o *ManualImportResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManualImportResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManualImportResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManualImportResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ManualImportResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ManualImportResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *ManualImportResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ManualImportResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ManualImportResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ManualImportResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetArtist

`func (o *ManualImportResource) GetArtist() ArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *ManualImportResource) GetArtistOk() (*ArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *ManualImportResource) SetArtist(v ArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *ManualImportResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetAlbum

`func (o *ManualImportResource) GetAlbum() AlbumResource`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *ManualImportResource) GetAlbumOk() (*AlbumResource, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *ManualImportResource) SetAlbum(v AlbumResource)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *ManualImportResource) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetAlbumReleaseId

`func (o *ManualImportResource) GetAlbumReleaseId() int32`

GetAlbumReleaseId returns the AlbumReleaseId field if non-nil, zero value otherwise.

### GetAlbumReleaseIdOk

`func (o *ManualImportResource) GetAlbumReleaseIdOk() (*int32, bool)`

GetAlbumReleaseIdOk returns a tuple with the AlbumReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumReleaseId

`func (o *ManualImportResource) SetAlbumReleaseId(v int32)`

SetAlbumReleaseId sets AlbumReleaseId field to given value.

### HasAlbumReleaseId

`func (o *ManualImportResource) HasAlbumReleaseId() bool`

HasAlbumReleaseId returns a boolean if a field has been set.

### GetTracks

`func (o *ManualImportResource) GetTracks() []TrackResource`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *ManualImportResource) GetTracksOk() (*[]TrackResource, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *ManualImportResource) SetTracks(v []TrackResource)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *ManualImportResource) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### SetTracksNil

`func (o *ManualImportResource) SetTracksNil(b bool)`

 SetTracksNil sets the value for Tracks to be an explicit nil

### UnsetTracks
`func (o *ManualImportResource) UnsetTracks()`

UnsetTracks ensures that no value is present for Tracks, not even an explicit nil
### GetQuality

`func (o *ManualImportResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ManualImportResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ManualImportResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ManualImportResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetQualityWeight

`func (o *ManualImportResource) GetQualityWeight() int32`

GetQualityWeight returns the QualityWeight field if non-nil, zero value otherwise.

### GetQualityWeightOk

`func (o *ManualImportResource) GetQualityWeightOk() (*int32, bool)`

GetQualityWeightOk returns a tuple with the QualityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityWeight

`func (o *ManualImportResource) SetQualityWeight(v int32)`

SetQualityWeight sets QualityWeight field to given value.

### HasQualityWeight

`func (o *ManualImportResource) HasQualityWeight() bool`

HasQualityWeight returns a boolean if a field has been set.

### GetDownloadId

`func (o *ManualImportResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *ManualImportResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *ManualImportResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *ManualImportResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *ManualImportResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *ManualImportResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetRejections

`func (o *ManualImportResource) GetRejections() []Rejection`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ManualImportResource) GetRejectionsOk() (*[]Rejection, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ManualImportResource) SetRejections(v []Rejection)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ManualImportResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ManualImportResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ManualImportResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil
### GetAudioTags

`func (o *ManualImportResource) GetAudioTags() ParsedTrackInfo`

GetAudioTags returns the AudioTags field if non-nil, zero value otherwise.

### GetAudioTagsOk

`func (o *ManualImportResource) GetAudioTagsOk() (*ParsedTrackInfo, bool)`

GetAudioTagsOk returns a tuple with the AudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTags

`func (o *ManualImportResource) SetAudioTags(v ParsedTrackInfo)`

SetAudioTags sets AudioTags field to given value.

### HasAudioTags

`func (o *ManualImportResource) HasAudioTags() bool`

HasAudioTags returns a boolean if a field has been set.

### GetAdditionalFile

`func (o *ManualImportResource) GetAdditionalFile() bool`

GetAdditionalFile returns the AdditionalFile field if non-nil, zero value otherwise.

### GetAdditionalFileOk

`func (o *ManualImportResource) GetAdditionalFileOk() (*bool, bool)`

GetAdditionalFileOk returns a tuple with the AdditionalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFile

`func (o *ManualImportResource) SetAdditionalFile(v bool)`

SetAdditionalFile sets AdditionalFile field to given value.

### HasAdditionalFile

`func (o *ManualImportResource) HasAdditionalFile() bool`

HasAdditionalFile returns a boolean if a field has been set.

### GetReplaceExistingFiles

`func (o *ManualImportResource) GetReplaceExistingFiles() bool`

GetReplaceExistingFiles returns the ReplaceExistingFiles field if non-nil, zero value otherwise.

### GetReplaceExistingFilesOk

`func (o *ManualImportResource) GetReplaceExistingFilesOk() (*bool, bool)`

GetReplaceExistingFilesOk returns a tuple with the ReplaceExistingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceExistingFiles

`func (o *ManualImportResource) SetReplaceExistingFiles(v bool)`

SetReplaceExistingFiles sets ReplaceExistingFiles field to given value.

### HasReplaceExistingFiles

`func (o *ManualImportResource) HasReplaceExistingFiles() bool`

HasReplaceExistingFiles returns a boolean if a field has been set.

### GetDisableReleaseSwitching

`func (o *ManualImportResource) GetDisableReleaseSwitching() bool`

GetDisableReleaseSwitching returns the DisableReleaseSwitching field if non-nil, zero value otherwise.

### GetDisableReleaseSwitchingOk

`func (o *ManualImportResource) GetDisableReleaseSwitchingOk() (*bool, bool)`

GetDisableReleaseSwitchingOk returns a tuple with the DisableReleaseSwitching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReleaseSwitching

`func (o *ManualImportResource) SetDisableReleaseSwitching(v bool)`

SetDisableReleaseSwitching sets DisableReleaseSwitching field to given value.

### HasDisableReleaseSwitching

`func (o *ManualImportResource) HasDisableReleaseSwitching() bool`

HasDisableReleaseSwitching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


