# RetagTrackResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistId** | Pointer to **int32** |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**TrackNumbers** | Pointer to **[]int32** |  | [optional] 
**TrackFileId** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Changes** | Pointer to [**[]TagDifference**](TagDifference.md) |  | [optional] 

## Methods

### NewRetagTrackResource

`func NewRetagTrackResource() *RetagTrackResource`

NewRetagTrackResource instantiates a new RetagTrackResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetagTrackResourceWithDefaults

`func NewRetagTrackResourceWithDefaults() *RetagTrackResource`

NewRetagTrackResourceWithDefaults instantiates a new RetagTrackResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RetagTrackResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetagTrackResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetagTrackResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RetagTrackResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistId

`func (o *RetagTrackResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *RetagTrackResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *RetagTrackResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *RetagTrackResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### GetAlbumId

`func (o *RetagTrackResource) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *RetagTrackResource) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *RetagTrackResource) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *RetagTrackResource) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetTrackNumbers

`func (o *RetagTrackResource) GetTrackNumbers() []int32`

GetTrackNumbers returns the TrackNumbers field if non-nil, zero value otherwise.

### GetTrackNumbersOk

`func (o *RetagTrackResource) GetTrackNumbersOk() (*[]int32, bool)`

GetTrackNumbersOk returns a tuple with the TrackNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumbers

`func (o *RetagTrackResource) SetTrackNumbers(v []int32)`

SetTrackNumbers sets TrackNumbers field to given value.

### HasTrackNumbers

`func (o *RetagTrackResource) HasTrackNumbers() bool`

HasTrackNumbers returns a boolean if a field has been set.

### SetTrackNumbersNil

`func (o *RetagTrackResource) SetTrackNumbersNil(b bool)`

 SetTrackNumbersNil sets the value for TrackNumbers to be an explicit nil

### UnsetTrackNumbers
`func (o *RetagTrackResource) UnsetTrackNumbers()`

UnsetTrackNumbers ensures that no value is present for TrackNumbers, not even an explicit nil
### GetTrackFileId

`func (o *RetagTrackResource) GetTrackFileId() int32`

GetTrackFileId returns the TrackFileId field if non-nil, zero value otherwise.

### GetTrackFileIdOk

`func (o *RetagTrackResource) GetTrackFileIdOk() (*int32, bool)`

GetTrackFileIdOk returns a tuple with the TrackFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFileId

`func (o *RetagTrackResource) SetTrackFileId(v int32)`

SetTrackFileId sets TrackFileId field to given value.

### HasTrackFileId

`func (o *RetagTrackResource) HasTrackFileId() bool`

HasTrackFileId returns a boolean if a field has been set.

### GetPath

`func (o *RetagTrackResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RetagTrackResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RetagTrackResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RetagTrackResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *RetagTrackResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *RetagTrackResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetChanges

`func (o *RetagTrackResource) GetChanges() []TagDifference`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *RetagTrackResource) GetChangesOk() (*[]TagDifference, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *RetagTrackResource) SetChanges(v []TagDifference)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *RetagTrackResource) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *RetagTrackResource) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *RetagTrackResource) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


