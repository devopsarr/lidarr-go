# ArtistEditorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtistIds** | Pointer to **[]int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**MonitorNewItems** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 
**QualityProfileId** | Pointer to **NullableInt32** |  | [optional] 
**MetadataProfileId** | Pointer to **NullableInt32** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**MoveFiles** | Pointer to **bool** |  | [optional] 
**DeleteFiles** | Pointer to **bool** |  | [optional] 

## Methods

### NewArtistEditorResource

`func NewArtistEditorResource() *ArtistEditorResource`

NewArtistEditorResource instantiates a new ArtistEditorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistEditorResourceWithDefaults

`func NewArtistEditorResourceWithDefaults() *ArtistEditorResource`

NewArtistEditorResourceWithDefaults instantiates a new ArtistEditorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtistIds

`func (o *ArtistEditorResource) GetArtistIds() []int32`

GetArtistIds returns the ArtistIds field if non-nil, zero value otherwise.

### GetArtistIdsOk

`func (o *ArtistEditorResource) GetArtistIdsOk() (*[]int32, bool)`

GetArtistIdsOk returns a tuple with the ArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistIds

`func (o *ArtistEditorResource) SetArtistIds(v []int32)`

SetArtistIds sets ArtistIds field to given value.

### HasArtistIds

`func (o *ArtistEditorResource) HasArtistIds() bool`

HasArtistIds returns a boolean if a field has been set.

### SetArtistIdsNil

`func (o *ArtistEditorResource) SetArtistIdsNil(b bool)`

 SetArtistIdsNil sets the value for ArtistIds to be an explicit nil

### UnsetArtistIds
`func (o *ArtistEditorResource) UnsetArtistIds()`

UnsetArtistIds ensures that no value is present for ArtistIds, not even an explicit nil
### GetMonitored

`func (o *ArtistEditorResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *ArtistEditorResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *ArtistEditorResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *ArtistEditorResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *ArtistEditorResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *ArtistEditorResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetMonitorNewItems

`func (o *ArtistEditorResource) GetMonitorNewItems() NewItemMonitorTypes`

GetMonitorNewItems returns the MonitorNewItems field if non-nil, zero value otherwise.

### GetMonitorNewItemsOk

`func (o *ArtistEditorResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool)`

GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNewItems

`func (o *ArtistEditorResource) SetMonitorNewItems(v NewItemMonitorTypes)`

SetMonitorNewItems sets MonitorNewItems field to given value.

### HasMonitorNewItems

`func (o *ArtistEditorResource) HasMonitorNewItems() bool`

HasMonitorNewItems returns a boolean if a field has been set.

### GetQualityProfileId

`func (o *ArtistEditorResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *ArtistEditorResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *ArtistEditorResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *ArtistEditorResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### SetQualityProfileIdNil

`func (o *ArtistEditorResource) SetQualityProfileIdNil(b bool)`

 SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil

### UnsetQualityProfileId
`func (o *ArtistEditorResource) UnsetQualityProfileId()`

UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
### GetMetadataProfileId

`func (o *ArtistEditorResource) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *ArtistEditorResource) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *ArtistEditorResource) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *ArtistEditorResource) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### SetMetadataProfileIdNil

`func (o *ArtistEditorResource) SetMetadataProfileIdNil(b bool)`

 SetMetadataProfileIdNil sets the value for MetadataProfileId to be an explicit nil

### UnsetMetadataProfileId
`func (o *ArtistEditorResource) UnsetMetadataProfileId()`

UnsetMetadataProfileId ensures that no value is present for MetadataProfileId, not even an explicit nil
### GetRootFolderPath

`func (o *ArtistEditorResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *ArtistEditorResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *ArtistEditorResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *ArtistEditorResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *ArtistEditorResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *ArtistEditorResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetTags

`func (o *ArtistEditorResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArtistEditorResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArtistEditorResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArtistEditorResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ArtistEditorResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ArtistEditorResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *ArtistEditorResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *ArtistEditorResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *ArtistEditorResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *ArtistEditorResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetMoveFiles

`func (o *ArtistEditorResource) GetMoveFiles() bool`

GetMoveFiles returns the MoveFiles field if non-nil, zero value otherwise.

### GetMoveFilesOk

`func (o *ArtistEditorResource) GetMoveFilesOk() (*bool, bool)`

GetMoveFilesOk returns a tuple with the MoveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveFiles

`func (o *ArtistEditorResource) SetMoveFiles(v bool)`

SetMoveFiles sets MoveFiles field to given value.

### HasMoveFiles

`func (o *ArtistEditorResource) HasMoveFiles() bool`

HasMoveFiles returns a boolean if a field has been set.

### GetDeleteFiles

`func (o *ArtistEditorResource) GetDeleteFiles() bool`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *ArtistEditorResource) GetDeleteFilesOk() (*bool, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *ArtistEditorResource) SetDeleteFiles(v bool)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *ArtistEditorResource) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


