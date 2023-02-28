# BlocklistResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistId** | Pointer to **int32** |  | [optional] 
**AlbumIds** | Pointer to **[]int32** |  | [optional] 
**SourceTitle** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**Indexer** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Artist** | Pointer to [**ArtistResource**](ArtistResource.md) |  | [optional] 

## Methods

### NewBlocklistResource

`func NewBlocklistResource() *BlocklistResource`

NewBlocklistResource instantiates a new BlocklistResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistResourceWithDefaults

`func NewBlocklistResourceWithDefaults() *BlocklistResource`

NewBlocklistResourceWithDefaults instantiates a new BlocklistResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlocklistResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlocklistResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlocklistResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BlocklistResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistId

`func (o *BlocklistResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *BlocklistResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *BlocklistResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *BlocklistResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### GetAlbumIds

`func (o *BlocklistResource) GetAlbumIds() []int32`

GetAlbumIds returns the AlbumIds field if non-nil, zero value otherwise.

### GetAlbumIdsOk

`func (o *BlocklistResource) GetAlbumIdsOk() (*[]int32, bool)`

GetAlbumIdsOk returns a tuple with the AlbumIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumIds

`func (o *BlocklistResource) SetAlbumIds(v []int32)`

SetAlbumIds sets AlbumIds field to given value.

### HasAlbumIds

`func (o *BlocklistResource) HasAlbumIds() bool`

HasAlbumIds returns a boolean if a field has been set.

### SetAlbumIdsNil

`func (o *BlocklistResource) SetAlbumIdsNil(b bool)`

 SetAlbumIdsNil sets the value for AlbumIds to be an explicit nil

### UnsetAlbumIds
`func (o *BlocklistResource) UnsetAlbumIds()`

UnsetAlbumIds ensures that no value is present for AlbumIds, not even an explicit nil
### GetSourceTitle

`func (o *BlocklistResource) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *BlocklistResource) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *BlocklistResource) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *BlocklistResource) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### SetSourceTitleNil

`func (o *BlocklistResource) SetSourceTitleNil(b bool)`

 SetSourceTitleNil sets the value for SourceTitle to be an explicit nil

### UnsetSourceTitle
`func (o *BlocklistResource) UnsetSourceTitle()`

UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
### GetQuality

`func (o *BlocklistResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BlocklistResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BlocklistResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BlocklistResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *BlocklistResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *BlocklistResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *BlocklistResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *BlocklistResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *BlocklistResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *BlocklistResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetDate

`func (o *BlocklistResource) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BlocklistResource) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BlocklistResource) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BlocklistResource) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetProtocol

`func (o *BlocklistResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BlocklistResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BlocklistResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BlocklistResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetIndexer

`func (o *BlocklistResource) GetIndexer() string`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *BlocklistResource) GetIndexerOk() (*string, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *BlocklistResource) SetIndexer(v string)`

SetIndexer sets Indexer field to given value.

### HasIndexer

`func (o *BlocklistResource) HasIndexer() bool`

HasIndexer returns a boolean if a field has been set.

### SetIndexerNil

`func (o *BlocklistResource) SetIndexerNil(b bool)`

 SetIndexerNil sets the value for Indexer to be an explicit nil

### UnsetIndexer
`func (o *BlocklistResource) UnsetIndexer()`

UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
### GetMessage

`func (o *BlocklistResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BlocklistResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BlocklistResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BlocklistResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BlocklistResource) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BlocklistResource) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetArtist

`func (o *BlocklistResource) GetArtist() ArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *BlocklistResource) GetArtistOk() (*ArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *BlocklistResource) SetArtist(v ArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *BlocklistResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


