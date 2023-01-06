# Artist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistMetadataId** | Pointer to **int32** |  | [optional] 
**CleanName** | Pointer to **NullableString** |  | [optional] 
**SortName** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**MonitorNewItems** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 
**LastInfoSync** | Pointer to **NullableTime** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**AddOptions** | Pointer to [**AddArtistOptions**](AddArtistOptions.md) |  | [optional] 
**Metadata** | Pointer to [**ArtistMetadataLazyLoaded**](ArtistMetadataLazyLoaded.md) |  | [optional] 
**QualityProfile** | Pointer to [**QualityProfileLazyLoaded**](QualityProfileLazyLoaded.md) |  | [optional] 
**MetadataProfile** | Pointer to [**MetadataProfileLazyLoaded**](MetadataProfileLazyLoaded.md) |  | [optional] 
**Albums** | Pointer to [**AlbumListLazyLoaded**](AlbumListLazyLoaded.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ForeignArtistId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewArtist

`func NewArtist() *Artist`

NewArtist instantiates a new Artist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistWithDefaults

`func NewArtistWithDefaults() *Artist`

NewArtistWithDefaults instantiates a new Artist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Artist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Artist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Artist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Artist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistMetadataId

`func (o *Artist) GetArtistMetadataId() int32`

GetArtistMetadataId returns the ArtistMetadataId field if non-nil, zero value otherwise.

### GetArtistMetadataIdOk

`func (o *Artist) GetArtistMetadataIdOk() (*int32, bool)`

GetArtistMetadataIdOk returns a tuple with the ArtistMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMetadataId

`func (o *Artist) SetArtistMetadataId(v int32)`

SetArtistMetadataId sets ArtistMetadataId field to given value.

### HasArtistMetadataId

`func (o *Artist) HasArtistMetadataId() bool`

HasArtistMetadataId returns a boolean if a field has been set.

### GetCleanName

`func (o *Artist) GetCleanName() string`

GetCleanName returns the CleanName field if non-nil, zero value otherwise.

### GetCleanNameOk

`func (o *Artist) GetCleanNameOk() (*string, bool)`

GetCleanNameOk returns a tuple with the CleanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanName

`func (o *Artist) SetCleanName(v string)`

SetCleanName sets CleanName field to given value.

### HasCleanName

`func (o *Artist) HasCleanName() bool`

HasCleanName returns a boolean if a field has been set.

### SetCleanNameNil

`func (o *Artist) SetCleanNameNil(b bool)`

 SetCleanNameNil sets the value for CleanName to be an explicit nil

### UnsetCleanName
`func (o *Artist) UnsetCleanName()`

UnsetCleanName ensures that no value is present for CleanName, not even an explicit nil
### GetSortName

`func (o *Artist) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *Artist) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *Artist) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *Artist) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *Artist) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *Artist) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil
### GetMonitored

`func (o *Artist) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *Artist) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *Artist) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *Artist) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetMonitorNewItems

`func (o *Artist) GetMonitorNewItems() NewItemMonitorTypes`

GetMonitorNewItems returns the MonitorNewItems field if non-nil, zero value otherwise.

### GetMonitorNewItemsOk

`func (o *Artist) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool)`

GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNewItems

`func (o *Artist) SetMonitorNewItems(v NewItemMonitorTypes)`

SetMonitorNewItems sets MonitorNewItems field to given value.

### HasMonitorNewItems

`func (o *Artist) HasMonitorNewItems() bool`

HasMonitorNewItems returns a boolean if a field has been set.

### GetLastInfoSync

`func (o *Artist) GetLastInfoSync() time.Time`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *Artist) GetLastInfoSyncOk() (*time.Time, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *Artist) SetLastInfoSync(v time.Time)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *Artist) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *Artist) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *Artist) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetPath

`func (o *Artist) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Artist) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Artist) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Artist) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Artist) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Artist) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRootFolderPath

`func (o *Artist) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *Artist) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *Artist) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *Artist) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *Artist) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *Artist) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetAdded

`func (o *Artist) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *Artist) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *Artist) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *Artist) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetQualityProfileId

`func (o *Artist) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *Artist) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *Artist) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *Artist) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *Artist) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *Artist) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *Artist) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *Artist) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetTags

`func (o *Artist) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Artist) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Artist) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Artist) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Artist) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Artist) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAddOptions

`func (o *Artist) GetAddOptions() AddArtistOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *Artist) GetAddOptionsOk() (*AddArtistOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *Artist) SetAddOptions(v AddArtistOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *Artist) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *Artist) GetMetadata() ArtistMetadataLazyLoaded`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Artist) GetMetadataOk() (*ArtistMetadataLazyLoaded, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Artist) SetMetadata(v ArtistMetadataLazyLoaded)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Artist) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQualityProfile

`func (o *Artist) GetQualityProfile() QualityProfileLazyLoaded`

GetQualityProfile returns the QualityProfile field if non-nil, zero value otherwise.

### GetQualityProfileOk

`func (o *Artist) GetQualityProfileOk() (*QualityProfileLazyLoaded, bool)`

GetQualityProfileOk returns a tuple with the QualityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfile

`func (o *Artist) SetQualityProfile(v QualityProfileLazyLoaded)`

SetQualityProfile sets QualityProfile field to given value.

### HasQualityProfile

`func (o *Artist) HasQualityProfile() bool`

HasQualityProfile returns a boolean if a field has been set.

### GetMetadataProfile

`func (o *Artist) GetMetadataProfile() MetadataProfileLazyLoaded`

GetMetadataProfile returns the MetadataProfile field if non-nil, zero value otherwise.

### GetMetadataProfileOk

`func (o *Artist) GetMetadataProfileOk() (*MetadataProfileLazyLoaded, bool)`

GetMetadataProfileOk returns a tuple with the MetadataProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfile

`func (o *Artist) SetMetadataProfile(v MetadataProfileLazyLoaded)`

SetMetadataProfile sets MetadataProfile field to given value.

### HasMetadataProfile

`func (o *Artist) HasMetadataProfile() bool`

HasMetadataProfile returns a boolean if a field has been set.

### GetAlbums

`func (o *Artist) GetAlbums() AlbumListLazyLoaded`

GetAlbums returns the Albums field if non-nil, zero value otherwise.

### GetAlbumsOk

`func (o *Artist) GetAlbumsOk() (*AlbumListLazyLoaded, bool)`

GetAlbumsOk returns a tuple with the Albums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbums

`func (o *Artist) SetAlbums(v AlbumListLazyLoaded)`

SetAlbums sets Albums field to given value.

### HasAlbums

`func (o *Artist) HasAlbums() bool`

HasAlbums returns a boolean if a field has been set.

### GetName

`func (o *Artist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Artist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Artist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Artist) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Artist) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Artist) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetForeignArtistId

`func (o *Artist) GetForeignArtistId() string`

GetForeignArtistId returns the ForeignArtistId field if non-nil, zero value otherwise.

### GetForeignArtistIdOk

`func (o *Artist) GetForeignArtistIdOk() (*string, bool)`

GetForeignArtistIdOk returns a tuple with the ForeignArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignArtistId

`func (o *Artist) SetForeignArtistId(v string)`

SetForeignArtistId sets ForeignArtistId field to given value.

### HasForeignArtistId

`func (o *Artist) HasForeignArtistId() bool`

HasForeignArtistId returns a boolean if a field has been set.

### SetForeignArtistIdNil

`func (o *Artist) SetForeignArtistIdNil(b bool)`

 SetForeignArtistIdNil sets the value for ForeignArtistId to be an explicit nil

### UnsetForeignArtistId
`func (o *Artist) UnsetForeignArtistId()`

UnsetForeignArtistId ensures that no value is present for ForeignArtistId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


