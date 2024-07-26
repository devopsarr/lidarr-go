# AlbumResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**ArtistId** | Pointer to **int32** |  | [optional] 
**ForeignAlbumId** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**AnyReleaseOk** | Pointer to **bool** |  | [optional] 
**ProfileId** | Pointer to **int32** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**AlbumType** | Pointer to **NullableString** |  | [optional] 
**SecondaryTypes** | Pointer to **[]string** |  | [optional] 
**MediumCount** | Pointer to **int32** |  | [optional] [readonly] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 
**Releases** | Pointer to [**[]AlbumReleaseResource**](AlbumReleaseResource.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Media** | Pointer to [**[]MediumResource**](MediumResource.md) |  | [optional] 
**Artist** | Pointer to [**ArtistResource**](ArtistResource.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Statistics** | Pointer to [**AlbumStatisticsResource**](AlbumStatisticsResource.md) |  | [optional] 
**AddOptions** | Pointer to [**AddAlbumOptions**](AddAlbumOptions.md) |  | [optional] 
**RemoteCover** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlbumResource

`func NewAlbumResource() *AlbumResource`

NewAlbumResource instantiates a new AlbumResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumResourceWithDefaults

`func NewAlbumResourceWithDefaults() *AlbumResource`

NewAlbumResourceWithDefaults instantiates a new AlbumResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlbumResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlbumResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlbumResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlbumResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *AlbumResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlbumResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlbumResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlbumResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AlbumResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AlbumResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDisambiguation

`func (o *AlbumResource) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *AlbumResource) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *AlbumResource) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *AlbumResource) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *AlbumResource) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *AlbumResource) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetOverview

`func (o *AlbumResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *AlbumResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *AlbumResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *AlbumResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *AlbumResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *AlbumResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetArtistId

`func (o *AlbumResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *AlbumResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *AlbumResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *AlbumResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### GetForeignAlbumId

`func (o *AlbumResource) GetForeignAlbumId() string`

GetForeignAlbumId returns the ForeignAlbumId field if non-nil, zero value otherwise.

### GetForeignAlbumIdOk

`func (o *AlbumResource) GetForeignAlbumIdOk() (*string, bool)`

GetForeignAlbumIdOk returns a tuple with the ForeignAlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAlbumId

`func (o *AlbumResource) SetForeignAlbumId(v string)`

SetForeignAlbumId sets ForeignAlbumId field to given value.

### HasForeignAlbumId

`func (o *AlbumResource) HasForeignAlbumId() bool`

HasForeignAlbumId returns a boolean if a field has been set.

### SetForeignAlbumIdNil

`func (o *AlbumResource) SetForeignAlbumIdNil(b bool)`

 SetForeignAlbumIdNil sets the value for ForeignAlbumId to be an explicit nil

### UnsetForeignAlbumId
`func (o *AlbumResource) UnsetForeignAlbumId()`

UnsetForeignAlbumId ensures that no value is present for ForeignAlbumId, not even an explicit nil
### GetMonitored

`func (o *AlbumResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AlbumResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AlbumResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AlbumResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetAnyReleaseOk

`func (o *AlbumResource) GetAnyReleaseOk() bool`

GetAnyReleaseOk returns the AnyReleaseOk field if non-nil, zero value otherwise.

### GetAnyReleaseOkOk

`func (o *AlbumResource) GetAnyReleaseOkOk() (*bool, bool)`

GetAnyReleaseOkOk returns a tuple with the AnyReleaseOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyReleaseOk

`func (o *AlbumResource) SetAnyReleaseOk(v bool)`

SetAnyReleaseOk sets AnyReleaseOk field to given value.

### HasAnyReleaseOk

`func (o *AlbumResource) HasAnyReleaseOk() bool`

HasAnyReleaseOk returns a boolean if a field has been set.

### GetProfileId

`func (o *AlbumResource) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AlbumResource) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AlbumResource) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AlbumResource) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetDuration

`func (o *AlbumResource) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AlbumResource) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AlbumResource) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AlbumResource) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAlbumType

`func (o *AlbumResource) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *AlbumResource) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *AlbumResource) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.

### HasAlbumType

`func (o *AlbumResource) HasAlbumType() bool`

HasAlbumType returns a boolean if a field has been set.

### SetAlbumTypeNil

`func (o *AlbumResource) SetAlbumTypeNil(b bool)`

 SetAlbumTypeNil sets the value for AlbumType to be an explicit nil

### UnsetAlbumType
`func (o *AlbumResource) UnsetAlbumType()`

UnsetAlbumType ensures that no value is present for AlbumType, not even an explicit nil
### GetSecondaryTypes

`func (o *AlbumResource) GetSecondaryTypes() []string`

GetSecondaryTypes returns the SecondaryTypes field if non-nil, zero value otherwise.

### GetSecondaryTypesOk

`func (o *AlbumResource) GetSecondaryTypesOk() (*[]string, bool)`

GetSecondaryTypesOk returns a tuple with the SecondaryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTypes

`func (o *AlbumResource) SetSecondaryTypes(v []string)`

SetSecondaryTypes sets SecondaryTypes field to given value.

### HasSecondaryTypes

`func (o *AlbumResource) HasSecondaryTypes() bool`

HasSecondaryTypes returns a boolean if a field has been set.

### SetSecondaryTypesNil

`func (o *AlbumResource) SetSecondaryTypesNil(b bool)`

 SetSecondaryTypesNil sets the value for SecondaryTypes to be an explicit nil

### UnsetSecondaryTypes
`func (o *AlbumResource) UnsetSecondaryTypes()`

UnsetSecondaryTypes ensures that no value is present for SecondaryTypes, not even an explicit nil
### GetMediumCount

`func (o *AlbumResource) GetMediumCount() int32`

GetMediumCount returns the MediumCount field if non-nil, zero value otherwise.

### GetMediumCountOk

`func (o *AlbumResource) GetMediumCountOk() (*int32, bool)`

GetMediumCountOk returns a tuple with the MediumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumCount

`func (o *AlbumResource) SetMediumCount(v int32)`

SetMediumCount sets MediumCount field to given value.

### HasMediumCount

`func (o *AlbumResource) HasMediumCount() bool`

HasMediumCount returns a boolean if a field has been set.

### GetRatings

`func (o *AlbumResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *AlbumResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *AlbumResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *AlbumResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetReleaseDate

`func (o *AlbumResource) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AlbumResource) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AlbumResource) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AlbumResource) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *AlbumResource) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *AlbumResource) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetReleases

`func (o *AlbumResource) GetReleases() []AlbumReleaseResource`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *AlbumResource) GetReleasesOk() (*[]AlbumReleaseResource, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *AlbumResource) SetReleases(v []AlbumReleaseResource)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *AlbumResource) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### SetReleasesNil

`func (o *AlbumResource) SetReleasesNil(b bool)`

 SetReleasesNil sets the value for Releases to be an explicit nil

### UnsetReleases
`func (o *AlbumResource) UnsetReleases()`

UnsetReleases ensures that no value is present for Releases, not even an explicit nil
### GetGenres

`func (o *AlbumResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AlbumResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AlbumResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AlbumResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *AlbumResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *AlbumResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetMedia

`func (o *AlbumResource) GetMedia() []MediumResource`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *AlbumResource) GetMediaOk() (*[]MediumResource, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *AlbumResource) SetMedia(v []MediumResource)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *AlbumResource) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *AlbumResource) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *AlbumResource) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetArtist

`func (o *AlbumResource) GetArtist() ArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *AlbumResource) GetArtistOk() (*ArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *AlbumResource) SetArtist(v ArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *AlbumResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetImages

`func (o *AlbumResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AlbumResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AlbumResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *AlbumResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *AlbumResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *AlbumResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *AlbumResource) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlbumResource) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlbumResource) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlbumResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AlbumResource) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AlbumResource) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStatistics

`func (o *AlbumResource) GetStatistics() AlbumStatisticsResource`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *AlbumResource) GetStatisticsOk() (*AlbumStatisticsResource, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *AlbumResource) SetStatistics(v AlbumStatisticsResource)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *AlbumResource) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetAddOptions

`func (o *AlbumResource) GetAddOptions() AddAlbumOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *AlbumResource) GetAddOptionsOk() (*AddAlbumOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *AlbumResource) SetAddOptions(v AddAlbumOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *AlbumResource) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetRemoteCover

`func (o *AlbumResource) GetRemoteCover() string`

GetRemoteCover returns the RemoteCover field if non-nil, zero value otherwise.

### GetRemoteCoverOk

`func (o *AlbumResource) GetRemoteCoverOk() (*string, bool)`

GetRemoteCoverOk returns a tuple with the RemoteCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCover

`func (o *AlbumResource) SetRemoteCover(v string)`

SetRemoteCover sets RemoteCover field to given value.

### HasRemoteCover

`func (o *AlbumResource) HasRemoteCover() bool`

HasRemoteCover returns a boolean if a field has been set.

### SetRemoteCoverNil

`func (o *AlbumResource) SetRemoteCoverNil(b bool)`

 SetRemoteCoverNil sets the value for RemoteCover to be an explicit nil

### UnsetRemoteCover
`func (o *AlbumResource) UnsetRemoteCover()`

UnsetRemoteCover ensures that no value is present for RemoteCover, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


