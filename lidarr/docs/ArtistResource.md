# ArtistResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistMetadataId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**ArtistStatusType**](ArtistStatusType.md) |  | [optional] 
**Ended** | Pointer to **bool** |  | [optional] [readonly] 
**ArtistName** | Pointer to **NullableString** |  | [optional] 
**ForeignArtistId** | Pointer to **NullableString** |  | [optional] 
**MbId** | Pointer to **NullableString** |  | [optional] 
**TadbId** | Pointer to **int32** |  | [optional] 
**DiscogsId** | Pointer to **int32** |  | [optional] 
**AllMusicId** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**ArtistType** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**NextAlbum** | Pointer to [**Album**](Album.md) |  | [optional] 
**LastAlbum** | Pointer to [**Album**](Album.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Members** | Pointer to [**[]Member**](Member.md) |  | [optional] 
**RemotePoster** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**MonitorNewItems** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Folder** | Pointer to **NullableString** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**CleanName** | Pointer to **NullableString** |  | [optional] 
**SortName** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**AddOptions** | Pointer to [**AddArtistOptions**](AddArtistOptions.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Statistics** | Pointer to [**ArtistStatisticsResource**](ArtistStatisticsResource.md) |  | [optional] 

## Methods

### NewArtistResource

`func NewArtistResource() *ArtistResource`

NewArtistResource instantiates a new ArtistResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistResourceWithDefaults

`func NewArtistResourceWithDefaults() *ArtistResource`

NewArtistResourceWithDefaults instantiates a new ArtistResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArtistResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtistResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtistResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArtistResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistMetadataId

`func (o *ArtistResource) GetArtistMetadataId() int32`

GetArtistMetadataId returns the ArtistMetadataId field if non-nil, zero value otherwise.

### GetArtistMetadataIdOk

`func (o *ArtistResource) GetArtistMetadataIdOk() (*int32, bool)`

GetArtistMetadataIdOk returns a tuple with the ArtistMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMetadataId

`func (o *ArtistResource) SetArtistMetadataId(v int32)`

SetArtistMetadataId sets ArtistMetadataId field to given value.

### HasArtistMetadataId

`func (o *ArtistResource) HasArtistMetadataId() bool`

HasArtistMetadataId returns a boolean if a field has been set.

### GetStatus

`func (o *ArtistResource) GetStatus() ArtistStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArtistResource) GetStatusOk() (*ArtistStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArtistResource) SetStatus(v ArtistStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArtistResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnded

`func (o *ArtistResource) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *ArtistResource) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *ArtistResource) SetEnded(v bool)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *ArtistResource) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetArtistName

`func (o *ArtistResource) GetArtistName() string`

GetArtistName returns the ArtistName field if non-nil, zero value otherwise.

### GetArtistNameOk

`func (o *ArtistResource) GetArtistNameOk() (*string, bool)`

GetArtistNameOk returns a tuple with the ArtistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistName

`func (o *ArtistResource) SetArtistName(v string)`

SetArtistName sets ArtistName field to given value.

### HasArtistName

`func (o *ArtistResource) HasArtistName() bool`

HasArtistName returns a boolean if a field has been set.

### SetArtistNameNil

`func (o *ArtistResource) SetArtistNameNil(b bool)`

 SetArtistNameNil sets the value for ArtistName to be an explicit nil

### UnsetArtistName
`func (o *ArtistResource) UnsetArtistName()`

UnsetArtistName ensures that no value is present for ArtistName, not even an explicit nil
### GetForeignArtistId

`func (o *ArtistResource) GetForeignArtistId() string`

GetForeignArtistId returns the ForeignArtistId field if non-nil, zero value otherwise.

### GetForeignArtistIdOk

`func (o *ArtistResource) GetForeignArtistIdOk() (*string, bool)`

GetForeignArtistIdOk returns a tuple with the ForeignArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignArtistId

`func (o *ArtistResource) SetForeignArtistId(v string)`

SetForeignArtistId sets ForeignArtistId field to given value.

### HasForeignArtistId

`func (o *ArtistResource) HasForeignArtistId() bool`

HasForeignArtistId returns a boolean if a field has been set.

### SetForeignArtistIdNil

`func (o *ArtistResource) SetForeignArtistIdNil(b bool)`

 SetForeignArtistIdNil sets the value for ForeignArtistId to be an explicit nil

### UnsetForeignArtistId
`func (o *ArtistResource) UnsetForeignArtistId()`

UnsetForeignArtistId ensures that no value is present for ForeignArtistId, not even an explicit nil
### GetMbId

`func (o *ArtistResource) GetMbId() string`

GetMbId returns the MbId field if non-nil, zero value otherwise.

### GetMbIdOk

`func (o *ArtistResource) GetMbIdOk() (*string, bool)`

GetMbIdOk returns a tuple with the MbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbId

`func (o *ArtistResource) SetMbId(v string)`

SetMbId sets MbId field to given value.

### HasMbId

`func (o *ArtistResource) HasMbId() bool`

HasMbId returns a boolean if a field has been set.

### SetMbIdNil

`func (o *ArtistResource) SetMbIdNil(b bool)`

 SetMbIdNil sets the value for MbId to be an explicit nil

### UnsetMbId
`func (o *ArtistResource) UnsetMbId()`

UnsetMbId ensures that no value is present for MbId, not even an explicit nil
### GetTadbId

`func (o *ArtistResource) GetTadbId() int32`

GetTadbId returns the TadbId field if non-nil, zero value otherwise.

### GetTadbIdOk

`func (o *ArtistResource) GetTadbIdOk() (*int32, bool)`

GetTadbIdOk returns a tuple with the TadbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTadbId

`func (o *ArtistResource) SetTadbId(v int32)`

SetTadbId sets TadbId field to given value.

### HasTadbId

`func (o *ArtistResource) HasTadbId() bool`

HasTadbId returns a boolean if a field has been set.

### GetDiscogsId

`func (o *ArtistResource) GetDiscogsId() int32`

GetDiscogsId returns the DiscogsId field if non-nil, zero value otherwise.

### GetDiscogsIdOk

`func (o *ArtistResource) GetDiscogsIdOk() (*int32, bool)`

GetDiscogsIdOk returns a tuple with the DiscogsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscogsId

`func (o *ArtistResource) SetDiscogsId(v int32)`

SetDiscogsId sets DiscogsId field to given value.

### HasDiscogsId

`func (o *ArtistResource) HasDiscogsId() bool`

HasDiscogsId returns a boolean if a field has been set.

### GetAllMusicId

`func (o *ArtistResource) GetAllMusicId() string`

GetAllMusicId returns the AllMusicId field if non-nil, zero value otherwise.

### GetAllMusicIdOk

`func (o *ArtistResource) GetAllMusicIdOk() (*string, bool)`

GetAllMusicIdOk returns a tuple with the AllMusicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMusicId

`func (o *ArtistResource) SetAllMusicId(v string)`

SetAllMusicId sets AllMusicId field to given value.

### HasAllMusicId

`func (o *ArtistResource) HasAllMusicId() bool`

HasAllMusicId returns a boolean if a field has been set.

### SetAllMusicIdNil

`func (o *ArtistResource) SetAllMusicIdNil(b bool)`

 SetAllMusicIdNil sets the value for AllMusicId to be an explicit nil

### UnsetAllMusicId
`func (o *ArtistResource) UnsetAllMusicId()`

UnsetAllMusicId ensures that no value is present for AllMusicId, not even an explicit nil
### GetOverview

`func (o *ArtistResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ArtistResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ArtistResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ArtistResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *ArtistResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *ArtistResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetArtistType

`func (o *ArtistResource) GetArtistType() string`

GetArtistType returns the ArtistType field if non-nil, zero value otherwise.

### GetArtistTypeOk

`func (o *ArtistResource) GetArtistTypeOk() (*string, bool)`

GetArtistTypeOk returns a tuple with the ArtistType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistType

`func (o *ArtistResource) SetArtistType(v string)`

SetArtistType sets ArtistType field to given value.

### HasArtistType

`func (o *ArtistResource) HasArtistType() bool`

HasArtistType returns a boolean if a field has been set.

### SetArtistTypeNil

`func (o *ArtistResource) SetArtistTypeNil(b bool)`

 SetArtistTypeNil sets the value for ArtistType to be an explicit nil

### UnsetArtistType
`func (o *ArtistResource) UnsetArtistType()`

UnsetArtistType ensures that no value is present for ArtistType, not even an explicit nil
### GetDisambiguation

`func (o *ArtistResource) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *ArtistResource) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *ArtistResource) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *ArtistResource) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *ArtistResource) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *ArtistResource) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetLinks

`func (o *ArtistResource) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ArtistResource) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ArtistResource) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ArtistResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ArtistResource) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ArtistResource) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNextAlbum

`func (o *ArtistResource) GetNextAlbum() Album`

GetNextAlbum returns the NextAlbum field if non-nil, zero value otherwise.

### GetNextAlbumOk

`func (o *ArtistResource) GetNextAlbumOk() (*Album, bool)`

GetNextAlbumOk returns a tuple with the NextAlbum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAlbum

`func (o *ArtistResource) SetNextAlbum(v Album)`

SetNextAlbum sets NextAlbum field to given value.

### HasNextAlbum

`func (o *ArtistResource) HasNextAlbum() bool`

HasNextAlbum returns a boolean if a field has been set.

### GetLastAlbum

`func (o *ArtistResource) GetLastAlbum() Album`

GetLastAlbum returns the LastAlbum field if non-nil, zero value otherwise.

### GetLastAlbumOk

`func (o *ArtistResource) GetLastAlbumOk() (*Album, bool)`

GetLastAlbumOk returns a tuple with the LastAlbum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAlbum

`func (o *ArtistResource) SetLastAlbum(v Album)`

SetLastAlbum sets LastAlbum field to given value.

### HasLastAlbum

`func (o *ArtistResource) HasLastAlbum() bool`

HasLastAlbum returns a boolean if a field has been set.

### GetImages

`func (o *ArtistResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ArtistResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ArtistResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *ArtistResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *ArtistResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *ArtistResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetMembers

`func (o *ArtistResource) GetMembers() []Member`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ArtistResource) GetMembersOk() (*[]Member, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ArtistResource) SetMembers(v []Member)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ArtistResource) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *ArtistResource) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *ArtistResource) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetRemotePoster

`func (o *ArtistResource) GetRemotePoster() string`

GetRemotePoster returns the RemotePoster field if non-nil, zero value otherwise.

### GetRemotePosterOk

`func (o *ArtistResource) GetRemotePosterOk() (*string, bool)`

GetRemotePosterOk returns a tuple with the RemotePoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePoster

`func (o *ArtistResource) SetRemotePoster(v string)`

SetRemotePoster sets RemotePoster field to given value.

### HasRemotePoster

`func (o *ArtistResource) HasRemotePoster() bool`

HasRemotePoster returns a boolean if a field has been set.

### SetRemotePosterNil

`func (o *ArtistResource) SetRemotePosterNil(b bool)`

 SetRemotePosterNil sets the value for RemotePoster to be an explicit nil

### UnsetRemotePoster
`func (o *ArtistResource) UnsetRemotePoster()`

UnsetRemotePoster ensures that no value is present for RemotePoster, not even an explicit nil
### GetPath

`func (o *ArtistResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ArtistResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ArtistResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ArtistResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ArtistResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ArtistResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQualityProfileId

`func (o *ArtistResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *ArtistResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *ArtistResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *ArtistResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *ArtistResource) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *ArtistResource) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *ArtistResource) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *ArtistResource) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetMonitored

`func (o *ArtistResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *ArtistResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *ArtistResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *ArtistResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetMonitorNewItems

`func (o *ArtistResource) GetMonitorNewItems() NewItemMonitorTypes`

GetMonitorNewItems returns the MonitorNewItems field if non-nil, zero value otherwise.

### GetMonitorNewItemsOk

`func (o *ArtistResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool)`

GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNewItems

`func (o *ArtistResource) SetMonitorNewItems(v NewItemMonitorTypes)`

SetMonitorNewItems sets MonitorNewItems field to given value.

### HasMonitorNewItems

`func (o *ArtistResource) HasMonitorNewItems() bool`

HasMonitorNewItems returns a boolean if a field has been set.

### GetRootFolderPath

`func (o *ArtistResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *ArtistResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *ArtistResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *ArtistResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *ArtistResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *ArtistResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetFolder

`func (o *ArtistResource) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *ArtistResource) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *ArtistResource) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *ArtistResource) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *ArtistResource) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *ArtistResource) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetGenres

`func (o *ArtistResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ArtistResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ArtistResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ArtistResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *ArtistResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *ArtistResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetCleanName

`func (o *ArtistResource) GetCleanName() string`

GetCleanName returns the CleanName field if non-nil, zero value otherwise.

### GetCleanNameOk

`func (o *ArtistResource) GetCleanNameOk() (*string, bool)`

GetCleanNameOk returns a tuple with the CleanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanName

`func (o *ArtistResource) SetCleanName(v string)`

SetCleanName sets CleanName field to given value.

### HasCleanName

`func (o *ArtistResource) HasCleanName() bool`

HasCleanName returns a boolean if a field has been set.

### SetCleanNameNil

`func (o *ArtistResource) SetCleanNameNil(b bool)`

 SetCleanNameNil sets the value for CleanName to be an explicit nil

### UnsetCleanName
`func (o *ArtistResource) UnsetCleanName()`

UnsetCleanName ensures that no value is present for CleanName, not even an explicit nil
### GetSortName

`func (o *ArtistResource) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *ArtistResource) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *ArtistResource) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *ArtistResource) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *ArtistResource) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *ArtistResource) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil
### GetTags

`func (o *ArtistResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArtistResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArtistResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArtistResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ArtistResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ArtistResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAdded

`func (o *ArtistResource) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *ArtistResource) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *ArtistResource) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *ArtistResource) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetAddOptions

`func (o *ArtistResource) GetAddOptions() AddArtistOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *ArtistResource) GetAddOptionsOk() (*AddArtistOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *ArtistResource) SetAddOptions(v AddArtistOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *ArtistResource) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetRatings

`func (o *ArtistResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *ArtistResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *ArtistResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *ArtistResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetStatistics

`func (o *ArtistResource) GetStatistics() ArtistStatisticsResource`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ArtistResource) GetStatisticsOk() (*ArtistStatisticsResource, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ArtistResource) SetStatistics(v ArtistStatisticsResource)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *ArtistResource) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


