# AddAlbumOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddType** | Pointer to [**AlbumAddType**](AlbumAddType.md) |  | [optional] 
**SearchForNewAlbum** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddAlbumOptions

`func NewAddAlbumOptions() *AddAlbumOptions`

NewAddAlbumOptions instantiates a new AddAlbumOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlbumOptionsWithDefaults

`func NewAddAlbumOptionsWithDefaults() *AddAlbumOptions`

NewAddAlbumOptionsWithDefaults instantiates a new AddAlbumOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddType

`func (o *AddAlbumOptions) GetAddType() AlbumAddType`

GetAddType returns the AddType field if non-nil, zero value otherwise.

### GetAddTypeOk

`func (o *AddAlbumOptions) GetAddTypeOk() (*AlbumAddType, bool)`

GetAddTypeOk returns a tuple with the AddType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddType

`func (o *AddAlbumOptions) SetAddType(v AlbumAddType)`

SetAddType sets AddType field to given value.

### HasAddType

`func (o *AddAlbumOptions) HasAddType() bool`

HasAddType returns a boolean if a field has been set.

### GetSearchForNewAlbum

`func (o *AddAlbumOptions) GetSearchForNewAlbum() bool`

GetSearchForNewAlbum returns the SearchForNewAlbum field if non-nil, zero value otherwise.

### GetSearchForNewAlbumOk

`func (o *AddAlbumOptions) GetSearchForNewAlbumOk() (*bool, bool)`

GetSearchForNewAlbumOk returns a tuple with the SearchForNewAlbum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForNewAlbum

`func (o *AddAlbumOptions) SetSearchForNewAlbum(v bool)`

SetSearchForNewAlbum sets SearchForNewAlbum field to given value.

### HasSearchForNewAlbum

`func (o *AddAlbumOptions) HasSearchForNewAlbum() bool`

HasSearchForNewAlbum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


