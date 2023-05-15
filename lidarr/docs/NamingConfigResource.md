# NamingConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RenameTracks** | Pointer to **bool** |  | [optional] 
**ReplaceIllegalCharacters** | Pointer to **bool** |  | [optional] 
**ColonReplacementFormat** | Pointer to **int32** |  | [optional] 
**StandardTrackFormat** | Pointer to **NullableString** |  | [optional] 
**MultiDiscTrackFormat** | Pointer to **NullableString** |  | [optional] 
**ArtistFolderFormat** | Pointer to **NullableString** |  | [optional] 
**IncludeArtistName** | Pointer to **bool** |  | [optional] 
**IncludeAlbumTitle** | Pointer to **bool** |  | [optional] 
**IncludeQuality** | Pointer to **bool** |  | [optional] 
**ReplaceSpaces** | Pointer to **bool** |  | [optional] 
**Separator** | Pointer to **NullableString** |  | [optional] 
**NumberStyle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNamingConfigResource

`func NewNamingConfigResource() *NamingConfigResource`

NewNamingConfigResource instantiates a new NamingConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingConfigResourceWithDefaults

`func NewNamingConfigResourceWithDefaults() *NamingConfigResource`

NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NamingConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRenameTracks

`func (o *NamingConfigResource) GetRenameTracks() bool`

GetRenameTracks returns the RenameTracks field if non-nil, zero value otherwise.

### GetRenameTracksOk

`func (o *NamingConfigResource) GetRenameTracksOk() (*bool, bool)`

GetRenameTracksOk returns a tuple with the RenameTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameTracks

`func (o *NamingConfigResource) SetRenameTracks(v bool)`

SetRenameTracks sets RenameTracks field to given value.

### HasRenameTracks

`func (o *NamingConfigResource) HasRenameTracks() bool`

HasRenameTracks returns a boolean if a field has been set.

### GetReplaceIllegalCharacters

`func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool`

GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field if non-nil, zero value otherwise.

### GetReplaceIllegalCharactersOk

`func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool)`

GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceIllegalCharacters

`func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool)`

SetReplaceIllegalCharacters sets ReplaceIllegalCharacters field to given value.

### HasReplaceIllegalCharacters

`func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool`

HasReplaceIllegalCharacters returns a boolean if a field has been set.

### GetColonReplacementFormat

`func (o *NamingConfigResource) GetColonReplacementFormat() int32`

GetColonReplacementFormat returns the ColonReplacementFormat field if non-nil, zero value otherwise.

### GetColonReplacementFormatOk

`func (o *NamingConfigResource) GetColonReplacementFormatOk() (*int32, bool)`

GetColonReplacementFormatOk returns a tuple with the ColonReplacementFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColonReplacementFormat

`func (o *NamingConfigResource) SetColonReplacementFormat(v int32)`

SetColonReplacementFormat sets ColonReplacementFormat field to given value.

### HasColonReplacementFormat

`func (o *NamingConfigResource) HasColonReplacementFormat() bool`

HasColonReplacementFormat returns a boolean if a field has been set.

### GetStandardTrackFormat

`func (o *NamingConfigResource) GetStandardTrackFormat() string`

GetStandardTrackFormat returns the StandardTrackFormat field if non-nil, zero value otherwise.

### GetStandardTrackFormatOk

`func (o *NamingConfigResource) GetStandardTrackFormatOk() (*string, bool)`

GetStandardTrackFormatOk returns a tuple with the StandardTrackFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardTrackFormat

`func (o *NamingConfigResource) SetStandardTrackFormat(v string)`

SetStandardTrackFormat sets StandardTrackFormat field to given value.

### HasStandardTrackFormat

`func (o *NamingConfigResource) HasStandardTrackFormat() bool`

HasStandardTrackFormat returns a boolean if a field has been set.

### SetStandardTrackFormatNil

`func (o *NamingConfigResource) SetStandardTrackFormatNil(b bool)`

 SetStandardTrackFormatNil sets the value for StandardTrackFormat to be an explicit nil

### UnsetStandardTrackFormat
`func (o *NamingConfigResource) UnsetStandardTrackFormat()`

UnsetStandardTrackFormat ensures that no value is present for StandardTrackFormat, not even an explicit nil
### GetMultiDiscTrackFormat

`func (o *NamingConfigResource) GetMultiDiscTrackFormat() string`

GetMultiDiscTrackFormat returns the MultiDiscTrackFormat field if non-nil, zero value otherwise.

### GetMultiDiscTrackFormatOk

`func (o *NamingConfigResource) GetMultiDiscTrackFormatOk() (*string, bool)`

GetMultiDiscTrackFormatOk returns a tuple with the MultiDiscTrackFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiDiscTrackFormat

`func (o *NamingConfigResource) SetMultiDiscTrackFormat(v string)`

SetMultiDiscTrackFormat sets MultiDiscTrackFormat field to given value.

### HasMultiDiscTrackFormat

`func (o *NamingConfigResource) HasMultiDiscTrackFormat() bool`

HasMultiDiscTrackFormat returns a boolean if a field has been set.

### SetMultiDiscTrackFormatNil

`func (o *NamingConfigResource) SetMultiDiscTrackFormatNil(b bool)`

 SetMultiDiscTrackFormatNil sets the value for MultiDiscTrackFormat to be an explicit nil

### UnsetMultiDiscTrackFormat
`func (o *NamingConfigResource) UnsetMultiDiscTrackFormat()`

UnsetMultiDiscTrackFormat ensures that no value is present for MultiDiscTrackFormat, not even an explicit nil
### GetArtistFolderFormat

`func (o *NamingConfigResource) GetArtistFolderFormat() string`

GetArtistFolderFormat returns the ArtistFolderFormat field if non-nil, zero value otherwise.

### GetArtistFolderFormatOk

`func (o *NamingConfigResource) GetArtistFolderFormatOk() (*string, bool)`

GetArtistFolderFormatOk returns a tuple with the ArtistFolderFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistFolderFormat

`func (o *NamingConfigResource) SetArtistFolderFormat(v string)`

SetArtistFolderFormat sets ArtistFolderFormat field to given value.

### HasArtistFolderFormat

`func (o *NamingConfigResource) HasArtistFolderFormat() bool`

HasArtistFolderFormat returns a boolean if a field has been set.

### SetArtistFolderFormatNil

`func (o *NamingConfigResource) SetArtistFolderFormatNil(b bool)`

 SetArtistFolderFormatNil sets the value for ArtistFolderFormat to be an explicit nil

### UnsetArtistFolderFormat
`func (o *NamingConfigResource) UnsetArtistFolderFormat()`

UnsetArtistFolderFormat ensures that no value is present for ArtistFolderFormat, not even an explicit nil
### GetIncludeArtistName

`func (o *NamingConfigResource) GetIncludeArtistName() bool`

GetIncludeArtistName returns the IncludeArtistName field if non-nil, zero value otherwise.

### GetIncludeArtistNameOk

`func (o *NamingConfigResource) GetIncludeArtistNameOk() (*bool, bool)`

GetIncludeArtistNameOk returns a tuple with the IncludeArtistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeArtistName

`func (o *NamingConfigResource) SetIncludeArtistName(v bool)`

SetIncludeArtistName sets IncludeArtistName field to given value.

### HasIncludeArtistName

`func (o *NamingConfigResource) HasIncludeArtistName() bool`

HasIncludeArtistName returns a boolean if a field has been set.

### GetIncludeAlbumTitle

`func (o *NamingConfigResource) GetIncludeAlbumTitle() bool`

GetIncludeAlbumTitle returns the IncludeAlbumTitle field if non-nil, zero value otherwise.

### GetIncludeAlbumTitleOk

`func (o *NamingConfigResource) GetIncludeAlbumTitleOk() (*bool, bool)`

GetIncludeAlbumTitleOk returns a tuple with the IncludeAlbumTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAlbumTitle

`func (o *NamingConfigResource) SetIncludeAlbumTitle(v bool)`

SetIncludeAlbumTitle sets IncludeAlbumTitle field to given value.

### HasIncludeAlbumTitle

`func (o *NamingConfigResource) HasIncludeAlbumTitle() bool`

HasIncludeAlbumTitle returns a boolean if a field has been set.

### GetIncludeQuality

`func (o *NamingConfigResource) GetIncludeQuality() bool`

GetIncludeQuality returns the IncludeQuality field if non-nil, zero value otherwise.

### GetIncludeQualityOk

`func (o *NamingConfigResource) GetIncludeQualityOk() (*bool, bool)`

GetIncludeQualityOk returns a tuple with the IncludeQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQuality

`func (o *NamingConfigResource) SetIncludeQuality(v bool)`

SetIncludeQuality sets IncludeQuality field to given value.

### HasIncludeQuality

`func (o *NamingConfigResource) HasIncludeQuality() bool`

HasIncludeQuality returns a boolean if a field has been set.

### GetReplaceSpaces

`func (o *NamingConfigResource) GetReplaceSpaces() bool`

GetReplaceSpaces returns the ReplaceSpaces field if non-nil, zero value otherwise.

### GetReplaceSpacesOk

`func (o *NamingConfigResource) GetReplaceSpacesOk() (*bool, bool)`

GetReplaceSpacesOk returns a tuple with the ReplaceSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceSpaces

`func (o *NamingConfigResource) SetReplaceSpaces(v bool)`

SetReplaceSpaces sets ReplaceSpaces field to given value.

### HasReplaceSpaces

`func (o *NamingConfigResource) HasReplaceSpaces() bool`

HasReplaceSpaces returns a boolean if a field has been set.

### GetSeparator

`func (o *NamingConfigResource) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *NamingConfigResource) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *NamingConfigResource) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *NamingConfigResource) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *NamingConfigResource) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *NamingConfigResource) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetNumberStyle

`func (o *NamingConfigResource) GetNumberStyle() string`

GetNumberStyle returns the NumberStyle field if non-nil, zero value otherwise.

### GetNumberStyleOk

`func (o *NamingConfigResource) GetNumberStyleOk() (*string, bool)`

GetNumberStyleOk returns a tuple with the NumberStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberStyle

`func (o *NamingConfigResource) SetNumberStyle(v string)`

SetNumberStyle sets NumberStyle field to given value.

### HasNumberStyle

`func (o *NamingConfigResource) HasNumberStyle() bool`

HasNumberStyle returns a boolean if a field has been set.

### SetNumberStyleNil

`func (o *NamingConfigResource) SetNumberStyleNil(b bool)`

 SetNumberStyleNil sets the value for NumberStyle to be an explicit nil

### UnsetNumberStyle
`func (o *NamingConfigResource) UnsetNumberStyle()`

UnsetNumberStyle ensures that no value is present for NumberStyle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


