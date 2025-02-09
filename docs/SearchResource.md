# SearchResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForeignId** | Pointer to **NullableString** |  | [optional] 
**Artist** | Pointer to [**ArtistResource**](ArtistResource.md) |  | [optional] 
**Album** | Pointer to [**AlbumResource**](AlbumResource.md) |  | [optional] 

## Methods

### NewSearchResource

`func NewSearchResource() *SearchResource`

NewSearchResource instantiates a new SearchResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResourceWithDefaults

`func NewSearchResourceWithDefaults() *SearchResource`

NewSearchResourceWithDefaults instantiates a new SearchResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SearchResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignId

`func (o *SearchResource) GetForeignId() string`

GetForeignId returns the ForeignId field if non-nil, zero value otherwise.

### GetForeignIdOk

`func (o *SearchResource) GetForeignIdOk() (*string, bool)`

GetForeignIdOk returns a tuple with the ForeignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignId

`func (o *SearchResource) SetForeignId(v string)`

SetForeignId sets ForeignId field to given value.

### HasForeignId

`func (o *SearchResource) HasForeignId() bool`

HasForeignId returns a boolean if a field has been set.

### SetForeignIdNil

`func (o *SearchResource) SetForeignIdNil(b bool)`

 SetForeignIdNil sets the value for ForeignId to be an explicit nil

### UnsetForeignId
`func (o *SearchResource) UnsetForeignId()`

UnsetForeignId ensures that no value is present for ForeignId, not even an explicit nil
### GetArtist

`func (o *SearchResource) GetArtist() ArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *SearchResource) GetArtistOk() (*ArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *SearchResource) SetArtist(v ArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *SearchResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetAlbum

`func (o *SearchResource) GetAlbum() AlbumResource`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *SearchResource) GetAlbumOk() (*AlbumResource, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *SearchResource) SetAlbum(v AlbumResource)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *SearchResource) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


