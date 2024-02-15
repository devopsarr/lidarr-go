# AddArtistOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**AlbumsToMonitor** | Pointer to **[]string** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**SearchForMissingAlbums** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddArtistOptions

`func NewAddArtistOptions() *AddArtistOptions`

NewAddArtistOptions instantiates a new AddArtistOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArtistOptionsWithDefaults

`func NewAddArtistOptionsWithDefaults() *AddArtistOptions`

NewAddArtistOptionsWithDefaults instantiates a new AddArtistOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *AddArtistOptions) GetMonitor() MonitorTypes`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *AddArtistOptions) GetMonitorOk() (*MonitorTypes, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *AddArtistOptions) SetMonitor(v MonitorTypes)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *AddArtistOptions) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetAlbumsToMonitor

`func (o *AddArtistOptions) GetAlbumsToMonitor() []string`

GetAlbumsToMonitor returns the AlbumsToMonitor field if non-nil, zero value otherwise.

### GetAlbumsToMonitorOk

`func (o *AddArtistOptions) GetAlbumsToMonitorOk() (*[]string, bool)`

GetAlbumsToMonitorOk returns a tuple with the AlbumsToMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumsToMonitor

`func (o *AddArtistOptions) SetAlbumsToMonitor(v []string)`

SetAlbumsToMonitor sets AlbumsToMonitor field to given value.

### HasAlbumsToMonitor

`func (o *AddArtistOptions) HasAlbumsToMonitor() bool`

HasAlbumsToMonitor returns a boolean if a field has been set.

### SetAlbumsToMonitorNil

`func (o *AddArtistOptions) SetAlbumsToMonitorNil(b bool)`

 SetAlbumsToMonitorNil sets the value for AlbumsToMonitor to be an explicit nil

### UnsetAlbumsToMonitor
`func (o *AddArtistOptions) UnsetAlbumsToMonitor()`

UnsetAlbumsToMonitor ensures that no value is present for AlbumsToMonitor, not even an explicit nil
### GetMonitored

`func (o *AddArtistOptions) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *AddArtistOptions) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *AddArtistOptions) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *AddArtistOptions) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetSearchForMissingAlbums

`func (o *AddArtistOptions) GetSearchForMissingAlbums() bool`

GetSearchForMissingAlbums returns the SearchForMissingAlbums field if non-nil, zero value otherwise.

### GetSearchForMissingAlbumsOk

`func (o *AddArtistOptions) GetSearchForMissingAlbumsOk() (*bool, bool)`

GetSearchForMissingAlbumsOk returns a tuple with the SearchForMissingAlbums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForMissingAlbums

`func (o *AddArtistOptions) SetSearchForMissingAlbums(v bool)`

SetSearchForMissingAlbums sets SearchForMissingAlbums field to given value.

### HasSearchForMissingAlbums

`func (o *AddArtistOptions) HasSearchForMissingAlbums() bool`

HasSearchForMissingAlbums returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


