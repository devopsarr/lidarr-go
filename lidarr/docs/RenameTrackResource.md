# RenameTrackResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistId** | Pointer to **int32** |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**TrackNumbers** | Pointer to **[]int32** |  | [optional] 
**TrackFileId** | Pointer to **int32** |  | [optional] 
**ExistingPath** | Pointer to **NullableString** |  | [optional] 
**NewPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRenameTrackResource

`func NewRenameTrackResource() *RenameTrackResource`

NewRenameTrackResource instantiates a new RenameTrackResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameTrackResourceWithDefaults

`func NewRenameTrackResourceWithDefaults() *RenameTrackResource`

NewRenameTrackResourceWithDefaults instantiates a new RenameTrackResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RenameTrackResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RenameTrackResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RenameTrackResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RenameTrackResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistId

`func (o *RenameTrackResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *RenameTrackResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *RenameTrackResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *RenameTrackResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### GetAlbumId

`func (o *RenameTrackResource) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *RenameTrackResource) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *RenameTrackResource) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *RenameTrackResource) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetTrackNumbers

`func (o *RenameTrackResource) GetTrackNumbers() []int32`

GetTrackNumbers returns the TrackNumbers field if non-nil, zero value otherwise.

### GetTrackNumbersOk

`func (o *RenameTrackResource) GetTrackNumbersOk() (*[]int32, bool)`

GetTrackNumbersOk returns a tuple with the TrackNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumbers

`func (o *RenameTrackResource) SetTrackNumbers(v []int32)`

SetTrackNumbers sets TrackNumbers field to given value.

### HasTrackNumbers

`func (o *RenameTrackResource) HasTrackNumbers() bool`

HasTrackNumbers returns a boolean if a field has been set.

### SetTrackNumbersNil

`func (o *RenameTrackResource) SetTrackNumbersNil(b bool)`

 SetTrackNumbersNil sets the value for TrackNumbers to be an explicit nil

### UnsetTrackNumbers
`func (o *RenameTrackResource) UnsetTrackNumbers()`

UnsetTrackNumbers ensures that no value is present for TrackNumbers, not even an explicit nil
### GetTrackFileId

`func (o *RenameTrackResource) GetTrackFileId() int32`

GetTrackFileId returns the TrackFileId field if non-nil, zero value otherwise.

### GetTrackFileIdOk

`func (o *RenameTrackResource) GetTrackFileIdOk() (*int32, bool)`

GetTrackFileIdOk returns a tuple with the TrackFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFileId

`func (o *RenameTrackResource) SetTrackFileId(v int32)`

SetTrackFileId sets TrackFileId field to given value.

### HasTrackFileId

`func (o *RenameTrackResource) HasTrackFileId() bool`

HasTrackFileId returns a boolean if a field has been set.

### GetExistingPath

`func (o *RenameTrackResource) GetExistingPath() string`

GetExistingPath returns the ExistingPath field if non-nil, zero value otherwise.

### GetExistingPathOk

`func (o *RenameTrackResource) GetExistingPathOk() (*string, bool)`

GetExistingPathOk returns a tuple with the ExistingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingPath

`func (o *RenameTrackResource) SetExistingPath(v string)`

SetExistingPath sets ExistingPath field to given value.

### HasExistingPath

`func (o *RenameTrackResource) HasExistingPath() bool`

HasExistingPath returns a boolean if a field has been set.

### SetExistingPathNil

`func (o *RenameTrackResource) SetExistingPathNil(b bool)`

 SetExistingPathNil sets the value for ExistingPath to be an explicit nil

### UnsetExistingPath
`func (o *RenameTrackResource) UnsetExistingPath()`

UnsetExistingPath ensures that no value is present for ExistingPath, not even an explicit nil
### GetNewPath

`func (o *RenameTrackResource) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *RenameTrackResource) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *RenameTrackResource) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *RenameTrackResource) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### SetNewPathNil

`func (o *RenameTrackResource) SetNewPathNil(b bool)`

 SetNewPathNil sets the value for NewPath to be an explicit nil

### UnsetNewPath
`func (o *RenameTrackResource) UnsetNewPath()`

UnsetNewPath ensures that no value is present for NewPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


