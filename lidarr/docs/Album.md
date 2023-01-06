# Album

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistMetadataId** | Pointer to **int32** |  | [optional] 
**ForeignAlbumId** | Pointer to **NullableString** |  | [optional] 
**OldForeignAlbumIds** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**AlbumType** | Pointer to **NullableString** |  | [optional] 
**SecondaryTypes** | Pointer to [**[]SecondaryAlbumType**](SecondaryAlbumType.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**ProfileId** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**AnyReleaseOk** | Pointer to **bool** |  | [optional] 
**LastInfoSync** | Pointer to **NullableTime** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**AddOptions** | Pointer to [**AddAlbumOptions**](AddAlbumOptions.md) |  | [optional] 
**ArtistMetadata** | Pointer to [**ArtistMetadataLazyLoaded**](ArtistMetadataLazyLoaded.md) |  | [optional] 
**AlbumReleases** | Pointer to [**AlbumReleaseListLazyLoaded**](AlbumReleaseListLazyLoaded.md) |  | [optional] 
**Artist** | Pointer to [**ArtistLazyLoaded**](ArtistLazyLoaded.md) |  | [optional] 

## Methods

### NewAlbum

`func NewAlbum() *Album`

NewAlbum instantiates a new Album object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumWithDefaults

`func NewAlbumWithDefaults() *Album`

NewAlbumWithDefaults instantiates a new Album object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Album) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Album) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Album) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Album) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistMetadataId

`func (o *Album) GetArtistMetadataId() int32`

GetArtistMetadataId returns the ArtistMetadataId field if non-nil, zero value otherwise.

### GetArtistMetadataIdOk

`func (o *Album) GetArtistMetadataIdOk() (*int32, bool)`

GetArtistMetadataIdOk returns a tuple with the ArtistMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMetadataId

`func (o *Album) SetArtistMetadataId(v int32)`

SetArtistMetadataId sets ArtistMetadataId field to given value.

### HasArtistMetadataId

`func (o *Album) HasArtistMetadataId() bool`

HasArtistMetadataId returns a boolean if a field has been set.

### GetForeignAlbumId

`func (o *Album) GetForeignAlbumId() string`

GetForeignAlbumId returns the ForeignAlbumId field if non-nil, zero value otherwise.

### GetForeignAlbumIdOk

`func (o *Album) GetForeignAlbumIdOk() (*string, bool)`

GetForeignAlbumIdOk returns a tuple with the ForeignAlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAlbumId

`func (o *Album) SetForeignAlbumId(v string)`

SetForeignAlbumId sets ForeignAlbumId field to given value.

### HasForeignAlbumId

`func (o *Album) HasForeignAlbumId() bool`

HasForeignAlbumId returns a boolean if a field has been set.

### SetForeignAlbumIdNil

`func (o *Album) SetForeignAlbumIdNil(b bool)`

 SetForeignAlbumIdNil sets the value for ForeignAlbumId to be an explicit nil

### UnsetForeignAlbumId
`func (o *Album) UnsetForeignAlbumId()`

UnsetForeignAlbumId ensures that no value is present for ForeignAlbumId, not even an explicit nil
### GetOldForeignAlbumIds

`func (o *Album) GetOldForeignAlbumIds() []string`

GetOldForeignAlbumIds returns the OldForeignAlbumIds field if non-nil, zero value otherwise.

### GetOldForeignAlbumIdsOk

`func (o *Album) GetOldForeignAlbumIdsOk() (*[]string, bool)`

GetOldForeignAlbumIdsOk returns a tuple with the OldForeignAlbumIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldForeignAlbumIds

`func (o *Album) SetOldForeignAlbumIds(v []string)`

SetOldForeignAlbumIds sets OldForeignAlbumIds field to given value.

### HasOldForeignAlbumIds

`func (o *Album) HasOldForeignAlbumIds() bool`

HasOldForeignAlbumIds returns a boolean if a field has been set.

### SetOldForeignAlbumIdsNil

`func (o *Album) SetOldForeignAlbumIdsNil(b bool)`

 SetOldForeignAlbumIdsNil sets the value for OldForeignAlbumIds to be an explicit nil

### UnsetOldForeignAlbumIds
`func (o *Album) UnsetOldForeignAlbumIds()`

UnsetOldForeignAlbumIds ensures that no value is present for OldForeignAlbumIds, not even an explicit nil
### GetTitle

`func (o *Album) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Album) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Album) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Album) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Album) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Album) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetOverview

`func (o *Album) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Album) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Album) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *Album) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *Album) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *Album) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetDisambiguation

`func (o *Album) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *Album) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *Album) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *Album) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *Album) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *Album) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetReleaseDate

`func (o *Album) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Album) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Album) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *Album) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *Album) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *Album) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetImages

`func (o *Album) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Album) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Album) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *Album) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *Album) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *Album) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *Album) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Album) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Album) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Album) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *Album) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *Album) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetGenres

`func (o *Album) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *Album) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *Album) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *Album) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *Album) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *Album) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetAlbumType

`func (o *Album) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *Album) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *Album) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.

### HasAlbumType

`func (o *Album) HasAlbumType() bool`

HasAlbumType returns a boolean if a field has been set.

### SetAlbumTypeNil

`func (o *Album) SetAlbumTypeNil(b bool)`

 SetAlbumTypeNil sets the value for AlbumType to be an explicit nil

### UnsetAlbumType
`func (o *Album) UnsetAlbumType()`

UnsetAlbumType ensures that no value is present for AlbumType, not even an explicit nil
### GetSecondaryTypes

`func (o *Album) GetSecondaryTypes() []SecondaryAlbumType`

GetSecondaryTypes returns the SecondaryTypes field if non-nil, zero value otherwise.

### GetSecondaryTypesOk

`func (o *Album) GetSecondaryTypesOk() (*[]SecondaryAlbumType, bool)`

GetSecondaryTypesOk returns a tuple with the SecondaryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTypes

`func (o *Album) SetSecondaryTypes(v []SecondaryAlbumType)`

SetSecondaryTypes sets SecondaryTypes field to given value.

### HasSecondaryTypes

`func (o *Album) HasSecondaryTypes() bool`

HasSecondaryTypes returns a boolean if a field has been set.

### SetSecondaryTypesNil

`func (o *Album) SetSecondaryTypesNil(b bool)`

 SetSecondaryTypesNil sets the value for SecondaryTypes to be an explicit nil

### UnsetSecondaryTypes
`func (o *Album) UnsetSecondaryTypes()`

UnsetSecondaryTypes ensures that no value is present for SecondaryTypes, not even an explicit nil
### GetRatings

`func (o *Album) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *Album) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *Album) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *Album) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetCleanTitle

`func (o *Album) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *Album) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *Album) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *Album) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *Album) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *Album) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetProfileId

`func (o *Album) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Album) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Album) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Album) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetMonitored

`func (o *Album) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *Album) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *Album) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *Album) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetAnyReleaseOk

`func (o *Album) GetAnyReleaseOk() bool`

GetAnyReleaseOk returns the AnyReleaseOk field if non-nil, zero value otherwise.

### GetAnyReleaseOkOk

`func (o *Album) GetAnyReleaseOkOk() (*bool, bool)`

GetAnyReleaseOkOk returns a tuple with the AnyReleaseOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyReleaseOk

`func (o *Album) SetAnyReleaseOk(v bool)`

SetAnyReleaseOk sets AnyReleaseOk field to given value.

### HasAnyReleaseOk

`func (o *Album) HasAnyReleaseOk() bool`

HasAnyReleaseOk returns a boolean if a field has been set.

### GetLastInfoSync

`func (o *Album) GetLastInfoSync() time.Time`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *Album) GetLastInfoSyncOk() (*time.Time, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *Album) SetLastInfoSync(v time.Time)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *Album) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *Album) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *Album) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetAdded

`func (o *Album) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *Album) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *Album) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *Album) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetAddOptions

`func (o *Album) GetAddOptions() AddAlbumOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *Album) GetAddOptionsOk() (*AddAlbumOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *Album) SetAddOptions(v AddAlbumOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *Album) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetArtistMetadata

`func (o *Album) GetArtistMetadata() ArtistMetadataLazyLoaded`

GetArtistMetadata returns the ArtistMetadata field if non-nil, zero value otherwise.

### GetArtistMetadataOk

`func (o *Album) GetArtistMetadataOk() (*ArtistMetadataLazyLoaded, bool)`

GetArtistMetadataOk returns a tuple with the ArtistMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMetadata

`func (o *Album) SetArtistMetadata(v ArtistMetadataLazyLoaded)`

SetArtistMetadata sets ArtistMetadata field to given value.

### HasArtistMetadata

`func (o *Album) HasArtistMetadata() bool`

HasArtistMetadata returns a boolean if a field has been set.

### GetAlbumReleases

`func (o *Album) GetAlbumReleases() AlbumReleaseListLazyLoaded`

GetAlbumReleases returns the AlbumReleases field if non-nil, zero value otherwise.

### GetAlbumReleasesOk

`func (o *Album) GetAlbumReleasesOk() (*AlbumReleaseListLazyLoaded, bool)`

GetAlbumReleasesOk returns a tuple with the AlbumReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumReleases

`func (o *Album) SetAlbumReleases(v AlbumReleaseListLazyLoaded)`

SetAlbumReleases sets AlbumReleases field to given value.

### HasAlbumReleases

`func (o *Album) HasAlbumReleases() bool`

HasAlbumReleases returns a boolean if a field has been set.

### GetArtist

`func (o *Album) GetArtist() ArtistLazyLoaded`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *Album) GetArtistOk() (*ArtistLazyLoaded, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *Album) SetArtist(v ArtistLazyLoaded)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *Album) HasArtist() bool`

HasArtist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


