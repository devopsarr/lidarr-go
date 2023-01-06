# AlbumStudioResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artist** | Pointer to [**[]AlbumStudioArtistResource**](AlbumStudioArtistResource.md) |  | [optional] 
**MonitoringOptions** | Pointer to [**MonitoringOptions**](MonitoringOptions.md) |  | [optional] 
**MonitorNewItems** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 

## Methods

### NewAlbumStudioResource

`func NewAlbumStudioResource() *AlbumStudioResource`

NewAlbumStudioResource instantiates a new AlbumStudioResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumStudioResourceWithDefaults

`func NewAlbumStudioResourceWithDefaults() *AlbumStudioResource`

NewAlbumStudioResourceWithDefaults instantiates a new AlbumStudioResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtist

`func (o *AlbumStudioResource) GetArtist() []AlbumStudioArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *AlbumStudioResource) GetArtistOk() (*[]AlbumStudioArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *AlbumStudioResource) SetArtist(v []AlbumStudioArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *AlbumStudioResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### SetArtistNil

`func (o *AlbumStudioResource) SetArtistNil(b bool)`

 SetArtistNil sets the value for Artist to be an explicit nil

### UnsetArtist
`func (o *AlbumStudioResource) UnsetArtist()`

UnsetArtist ensures that no value is present for Artist, not even an explicit nil
### GetMonitoringOptions

`func (o *AlbumStudioResource) GetMonitoringOptions() MonitoringOptions`

GetMonitoringOptions returns the MonitoringOptions field if non-nil, zero value otherwise.

### GetMonitoringOptionsOk

`func (o *AlbumStudioResource) GetMonitoringOptionsOk() (*MonitoringOptions, bool)`

GetMonitoringOptionsOk returns a tuple with the MonitoringOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringOptions

`func (o *AlbumStudioResource) SetMonitoringOptions(v MonitoringOptions)`

SetMonitoringOptions sets MonitoringOptions field to given value.

### HasMonitoringOptions

`func (o *AlbumStudioResource) HasMonitoringOptions() bool`

HasMonitoringOptions returns a boolean if a field has been set.

### GetMonitorNewItems

`func (o *AlbumStudioResource) GetMonitorNewItems() NewItemMonitorTypes`

GetMonitorNewItems returns the MonitorNewItems field if non-nil, zero value otherwise.

### GetMonitorNewItemsOk

`func (o *AlbumStudioResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool)`

GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorNewItems

`func (o *AlbumStudioResource) SetMonitorNewItems(v NewItemMonitorTypes)`

SetMonitorNewItems sets MonitorNewItems field to given value.

### HasMonitorNewItems

`func (o *AlbumStudioResource) HasMonitorNewItems() bool`

HasMonitorNewItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


