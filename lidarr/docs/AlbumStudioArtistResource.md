# AlbumStudioArtistResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**Albums** | Pointer to [**[]AlbumResource**](AlbumResource.md) |  | [optional] 

## Methods

### NewAlbumStudioArtistResource

`func NewAlbumStudioArtistResource() *AlbumStudioArtistResource`

NewAlbumStudioArtistResource instantiates a new AlbumStudioArtistResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumStudioArtistResourceWithDefaults

`func NewAlbumStudioArtistResourceWithDefaults() *AlbumStudioArtistResource`

NewAlbumStudioArtistResourceWithDefaults instantiates a new AlbumStudioArtistResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlbumStudioArtistResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlbumStudioArtistResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlbumStudioArtistResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlbumStudioArtistResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonitored

`func (o *AlbumStudioArtistResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AlbumStudioArtistResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AlbumStudioArtistResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AlbumStudioArtistResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *AlbumStudioArtistResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *AlbumStudioArtistResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetAlbums

`func (o *AlbumStudioArtistResource) GetAlbums() []AlbumResource`

GetAlbums returns the Albums field if non-nil, zero value otherwise.

### GetAlbumsOk

`func (o *AlbumStudioArtistResource) GetAlbumsOk() (*[]AlbumResource, bool)`

GetAlbumsOk returns a tuple with the Albums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbums

`func (o *AlbumStudioArtistResource) SetAlbums(v []AlbumResource)`

SetAlbums sets Albums field to given value.

### HasAlbums

`func (o *AlbumStudioArtistResource) HasAlbums() bool`

HasAlbums returns a boolean if a field has been set.

### SetAlbumsNil

`func (o *AlbumStudioArtistResource) SetAlbumsNil(b bool)`

 SetAlbumsNil sets the value for Albums to be an explicit nil

### UnsetAlbums
`func (o *AlbumStudioArtistResource) UnsetAlbums()`

UnsetAlbums ensures that no value is present for Albums, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


