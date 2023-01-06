# TrackFileListResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackFileIds** | Pointer to **[]int32** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 

## Methods

### NewTrackFileListResource

`func NewTrackFileListResource() *TrackFileListResource`

NewTrackFileListResource instantiates a new TrackFileListResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackFileListResourceWithDefaults

`func NewTrackFileListResourceWithDefaults() *TrackFileListResource`

NewTrackFileListResourceWithDefaults instantiates a new TrackFileListResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackFileIds

`func (o *TrackFileListResource) GetTrackFileIds() []int32`

GetTrackFileIds returns the TrackFileIds field if non-nil, zero value otherwise.

### GetTrackFileIdsOk

`func (o *TrackFileListResource) GetTrackFileIdsOk() (*[]int32, bool)`

GetTrackFileIdsOk returns a tuple with the TrackFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFileIds

`func (o *TrackFileListResource) SetTrackFileIds(v []int32)`

SetTrackFileIds sets TrackFileIds field to given value.

### HasTrackFileIds

`func (o *TrackFileListResource) HasTrackFileIds() bool`

HasTrackFileIds returns a boolean if a field has been set.

### SetTrackFileIdsNil

`func (o *TrackFileListResource) SetTrackFileIdsNil(b bool)`

 SetTrackFileIdsNil sets the value for TrackFileIds to be an explicit nil

### UnsetTrackFileIds
`func (o *TrackFileListResource) UnsetTrackFileIds()`

UnsetTrackFileIds ensures that no value is present for TrackFileIds, not even an explicit nil
### GetQuality

`func (o *TrackFileListResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *TrackFileListResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *TrackFileListResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *TrackFileListResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


