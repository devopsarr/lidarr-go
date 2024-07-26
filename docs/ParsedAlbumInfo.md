# ParsedAlbumInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseTitle** | Pointer to **NullableString** |  | [optional] 
**AlbumTitle** | Pointer to **NullableString** |  | [optional] 
**ArtistName** | Pointer to **NullableString** |  | [optional] 
**AlbumType** | Pointer to **NullableString** |  | [optional] 
**ArtistTitleInfo** | Pointer to [**ArtistTitleInfo**](ArtistTitleInfo.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**Discography** | Pointer to **bool** |  | [optional] 
**DiscographyStart** | Pointer to **int32** |  | [optional] 
**DiscographyEnd** | Pointer to **int32** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 
**ReleaseVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewParsedAlbumInfo

`func NewParsedAlbumInfo() *ParsedAlbumInfo`

NewParsedAlbumInfo instantiates a new ParsedAlbumInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedAlbumInfoWithDefaults

`func NewParsedAlbumInfoWithDefaults() *ParsedAlbumInfo`

NewParsedAlbumInfoWithDefaults instantiates a new ParsedAlbumInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseTitle

`func (o *ParsedAlbumInfo) GetReleaseTitle() string`

GetReleaseTitle returns the ReleaseTitle field if non-nil, zero value otherwise.

### GetReleaseTitleOk

`func (o *ParsedAlbumInfo) GetReleaseTitleOk() (*string, bool)`

GetReleaseTitleOk returns a tuple with the ReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTitle

`func (o *ParsedAlbumInfo) SetReleaseTitle(v string)`

SetReleaseTitle sets ReleaseTitle field to given value.

### HasReleaseTitle

`func (o *ParsedAlbumInfo) HasReleaseTitle() bool`

HasReleaseTitle returns a boolean if a field has been set.

### SetReleaseTitleNil

`func (o *ParsedAlbumInfo) SetReleaseTitleNil(b bool)`

 SetReleaseTitleNil sets the value for ReleaseTitle to be an explicit nil

### UnsetReleaseTitle
`func (o *ParsedAlbumInfo) UnsetReleaseTitle()`

UnsetReleaseTitle ensures that no value is present for ReleaseTitle, not even an explicit nil
### GetAlbumTitle

`func (o *ParsedAlbumInfo) GetAlbumTitle() string`

GetAlbumTitle returns the AlbumTitle field if non-nil, zero value otherwise.

### GetAlbumTitleOk

`func (o *ParsedAlbumInfo) GetAlbumTitleOk() (*string, bool)`

GetAlbumTitleOk returns a tuple with the AlbumTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumTitle

`func (o *ParsedAlbumInfo) SetAlbumTitle(v string)`

SetAlbumTitle sets AlbumTitle field to given value.

### HasAlbumTitle

`func (o *ParsedAlbumInfo) HasAlbumTitle() bool`

HasAlbumTitle returns a boolean if a field has been set.

### SetAlbumTitleNil

`func (o *ParsedAlbumInfo) SetAlbumTitleNil(b bool)`

 SetAlbumTitleNil sets the value for AlbumTitle to be an explicit nil

### UnsetAlbumTitle
`func (o *ParsedAlbumInfo) UnsetAlbumTitle()`

UnsetAlbumTitle ensures that no value is present for AlbumTitle, not even an explicit nil
### GetArtistName

`func (o *ParsedAlbumInfo) GetArtistName() string`

GetArtistName returns the ArtistName field if non-nil, zero value otherwise.

### GetArtistNameOk

`func (o *ParsedAlbumInfo) GetArtistNameOk() (*string, bool)`

GetArtistNameOk returns a tuple with the ArtistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistName

`func (o *ParsedAlbumInfo) SetArtistName(v string)`

SetArtistName sets ArtistName field to given value.

### HasArtistName

`func (o *ParsedAlbumInfo) HasArtistName() bool`

HasArtistName returns a boolean if a field has been set.

### SetArtistNameNil

`func (o *ParsedAlbumInfo) SetArtistNameNil(b bool)`

 SetArtistNameNil sets the value for ArtistName to be an explicit nil

### UnsetArtistName
`func (o *ParsedAlbumInfo) UnsetArtistName()`

UnsetArtistName ensures that no value is present for ArtistName, not even an explicit nil
### GetAlbumType

`func (o *ParsedAlbumInfo) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *ParsedAlbumInfo) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *ParsedAlbumInfo) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.

### HasAlbumType

`func (o *ParsedAlbumInfo) HasAlbumType() bool`

HasAlbumType returns a boolean if a field has been set.

### SetAlbumTypeNil

`func (o *ParsedAlbumInfo) SetAlbumTypeNil(b bool)`

 SetAlbumTypeNil sets the value for AlbumType to be an explicit nil

### UnsetAlbumType
`func (o *ParsedAlbumInfo) UnsetAlbumType()`

UnsetAlbumType ensures that no value is present for AlbumType, not even an explicit nil
### GetArtistTitleInfo

`func (o *ParsedAlbumInfo) GetArtistTitleInfo() ArtistTitleInfo`

GetArtistTitleInfo returns the ArtistTitleInfo field if non-nil, zero value otherwise.

### GetArtistTitleInfoOk

`func (o *ParsedAlbumInfo) GetArtistTitleInfoOk() (*ArtistTitleInfo, bool)`

GetArtistTitleInfoOk returns a tuple with the ArtistTitleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistTitleInfo

`func (o *ParsedAlbumInfo) SetArtistTitleInfo(v ArtistTitleInfo)`

SetArtistTitleInfo sets ArtistTitleInfo field to given value.

### HasArtistTitleInfo

`func (o *ParsedAlbumInfo) HasArtistTitleInfo() bool`

HasArtistTitleInfo returns a boolean if a field has been set.

### GetQuality

`func (o *ParsedAlbumInfo) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ParsedAlbumInfo) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ParsedAlbumInfo) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ParsedAlbumInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ParsedAlbumInfo) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ParsedAlbumInfo) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ParsedAlbumInfo) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ParsedAlbumInfo) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ParsedAlbumInfo) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ParsedAlbumInfo) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetDiscography

`func (o *ParsedAlbumInfo) GetDiscography() bool`

GetDiscography returns the Discography field if non-nil, zero value otherwise.

### GetDiscographyOk

`func (o *ParsedAlbumInfo) GetDiscographyOk() (*bool, bool)`

GetDiscographyOk returns a tuple with the Discography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscography

`func (o *ParsedAlbumInfo) SetDiscography(v bool)`

SetDiscography sets Discography field to given value.

### HasDiscography

`func (o *ParsedAlbumInfo) HasDiscography() bool`

HasDiscography returns a boolean if a field has been set.

### GetDiscographyStart

`func (o *ParsedAlbumInfo) GetDiscographyStart() int32`

GetDiscographyStart returns the DiscographyStart field if non-nil, zero value otherwise.

### GetDiscographyStartOk

`func (o *ParsedAlbumInfo) GetDiscographyStartOk() (*int32, bool)`

GetDiscographyStartOk returns a tuple with the DiscographyStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscographyStart

`func (o *ParsedAlbumInfo) SetDiscographyStart(v int32)`

SetDiscographyStart sets DiscographyStart field to given value.

### HasDiscographyStart

`func (o *ParsedAlbumInfo) HasDiscographyStart() bool`

HasDiscographyStart returns a boolean if a field has been set.

### GetDiscographyEnd

`func (o *ParsedAlbumInfo) GetDiscographyEnd() int32`

GetDiscographyEnd returns the DiscographyEnd field if non-nil, zero value otherwise.

### GetDiscographyEndOk

`func (o *ParsedAlbumInfo) GetDiscographyEndOk() (*int32, bool)`

GetDiscographyEndOk returns a tuple with the DiscographyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscographyEnd

`func (o *ParsedAlbumInfo) SetDiscographyEnd(v int32)`

SetDiscographyEnd sets DiscographyEnd field to given value.

### HasDiscographyEnd

`func (o *ParsedAlbumInfo) HasDiscographyEnd() bool`

HasDiscographyEnd returns a boolean if a field has been set.

### GetReleaseGroup

`func (o *ParsedAlbumInfo) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ParsedAlbumInfo) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ParsedAlbumInfo) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ParsedAlbumInfo) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ParsedAlbumInfo) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ParsedAlbumInfo) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetReleaseHash

`func (o *ParsedAlbumInfo) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ParsedAlbumInfo) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ParsedAlbumInfo) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ParsedAlbumInfo) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ParsedAlbumInfo) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ParsedAlbumInfo) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
### GetReleaseVersion

`func (o *ParsedAlbumInfo) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *ParsedAlbumInfo) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *ParsedAlbumInfo) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *ParsedAlbumInfo) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### SetReleaseVersionNil

`func (o *ParsedAlbumInfo) SetReleaseVersionNil(b bool)`

 SetReleaseVersionNil sets the value for ReleaseVersion to be an explicit nil

### UnsetReleaseVersion
`func (o *ParsedAlbumInfo) UnsetReleaseVersion()`

UnsetReleaseVersion ensures that no value is present for ReleaseVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


