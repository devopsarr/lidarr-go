# ParseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ParsedAlbumInfo** | Pointer to [**ParsedAlbumInfo**](ParsedAlbumInfo.md) |  | [optional] 
**Artist** | Pointer to [**ArtistResource**](ArtistResource.md) |  | [optional] 
**Albums** | Pointer to [**[]AlbumResource**](AlbumResource.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**CustomFormatScore** | Pointer to **int32** |  | [optional] 

## Methods

### NewParseResource

`func NewParseResource() *ParseResource`

NewParseResource instantiates a new ParseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseResourceWithDefaults

`func NewParseResourceWithDefaults() *ParseResource`

NewParseResourceWithDefaults instantiates a new ParseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ParseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ParseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ParseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ParseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ParseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetParsedAlbumInfo

`func (o *ParseResource) GetParsedAlbumInfo() ParsedAlbumInfo`

GetParsedAlbumInfo returns the ParsedAlbumInfo field if non-nil, zero value otherwise.

### GetParsedAlbumInfoOk

`func (o *ParseResource) GetParsedAlbumInfoOk() (*ParsedAlbumInfo, bool)`

GetParsedAlbumInfoOk returns a tuple with the ParsedAlbumInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedAlbumInfo

`func (o *ParseResource) SetParsedAlbumInfo(v ParsedAlbumInfo)`

SetParsedAlbumInfo sets ParsedAlbumInfo field to given value.

### HasParsedAlbumInfo

`func (o *ParseResource) HasParsedAlbumInfo() bool`

HasParsedAlbumInfo returns a boolean if a field has been set.

### GetArtist

`func (o *ParseResource) GetArtist() ArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *ParseResource) GetArtistOk() (*ArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *ParseResource) SetArtist(v ArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *ParseResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetAlbums

`func (o *ParseResource) GetAlbums() []AlbumResource`

GetAlbums returns the Albums field if non-nil, zero value otherwise.

### GetAlbumsOk

`func (o *ParseResource) GetAlbumsOk() (*[]AlbumResource, bool)`

GetAlbumsOk returns a tuple with the Albums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbums

`func (o *ParseResource) SetAlbums(v []AlbumResource)`

SetAlbums sets Albums field to given value.

### HasAlbums

`func (o *ParseResource) HasAlbums() bool`

HasAlbums returns a boolean if a field has been set.

### SetAlbumsNil

`func (o *ParseResource) SetAlbumsNil(b bool)`

 SetAlbumsNil sets the value for Albums to be an explicit nil

### UnsetAlbums
`func (o *ParseResource) UnsetAlbums()`

UnsetAlbums ensures that no value is present for Albums, not even an explicit nil
### GetCustomFormats

`func (o *ParseResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *ParseResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *ParseResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *ParseResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *ParseResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *ParseResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetCustomFormatScore

`func (o *ParseResource) GetCustomFormatScore() int32`

GetCustomFormatScore returns the CustomFormatScore field if non-nil, zero value otherwise.

### GetCustomFormatScoreOk

`func (o *ParseResource) GetCustomFormatScoreOk() (*int32, bool)`

GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormatScore

`func (o *ParseResource) SetCustomFormatScore(v int32)`

SetCustomFormatScore sets CustomFormatScore field to given value.

### HasCustomFormatScore

`func (o *ParseResource) HasCustomFormatScore() bool`

HasCustomFormatScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


