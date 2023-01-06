# MetadataProfileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PrimaryAlbumTypes** | Pointer to [**[]ProfilePrimaryAlbumTypeItemResource**](ProfilePrimaryAlbumTypeItemResource.md) |  | [optional] 
**SecondaryAlbumTypes** | Pointer to [**[]ProfileSecondaryAlbumTypeItemResource**](ProfileSecondaryAlbumTypeItemResource.md) |  | [optional] 
**ReleaseStatuses** | Pointer to [**[]ProfileReleaseStatusItemResource**](ProfileReleaseStatusItemResource.md) |  | [optional] 

## Methods

### NewMetadataProfileResource

`func NewMetadataProfileResource() *MetadataProfileResource`

NewMetadataProfileResource instantiates a new MetadataProfileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProfileResourceWithDefaults

`func NewMetadataProfileResourceWithDefaults() *MetadataProfileResource`

NewMetadataProfileResourceWithDefaults instantiates a new MetadataProfileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataProfileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataProfileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataProfileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataProfileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataProfileResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataProfileResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataProfileResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataProfileResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataProfileResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataProfileResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrimaryAlbumTypes

`func (o *MetadataProfileResource) GetPrimaryAlbumTypes() []ProfilePrimaryAlbumTypeItemResource`

GetPrimaryAlbumTypes returns the PrimaryAlbumTypes field if non-nil, zero value otherwise.

### GetPrimaryAlbumTypesOk

`func (o *MetadataProfileResource) GetPrimaryAlbumTypesOk() (*[]ProfilePrimaryAlbumTypeItemResource, bool)`

GetPrimaryAlbumTypesOk returns a tuple with the PrimaryAlbumTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAlbumTypes

`func (o *MetadataProfileResource) SetPrimaryAlbumTypes(v []ProfilePrimaryAlbumTypeItemResource)`

SetPrimaryAlbumTypes sets PrimaryAlbumTypes field to given value.

### HasPrimaryAlbumTypes

`func (o *MetadataProfileResource) HasPrimaryAlbumTypes() bool`

HasPrimaryAlbumTypes returns a boolean if a field has been set.

### SetPrimaryAlbumTypesNil

`func (o *MetadataProfileResource) SetPrimaryAlbumTypesNil(b bool)`

 SetPrimaryAlbumTypesNil sets the value for PrimaryAlbumTypes to be an explicit nil

### UnsetPrimaryAlbumTypes
`func (o *MetadataProfileResource) UnsetPrimaryAlbumTypes()`

UnsetPrimaryAlbumTypes ensures that no value is present for PrimaryAlbumTypes, not even an explicit nil
### GetSecondaryAlbumTypes

`func (o *MetadataProfileResource) GetSecondaryAlbumTypes() []ProfileSecondaryAlbumTypeItemResource`

GetSecondaryAlbumTypes returns the SecondaryAlbumTypes field if non-nil, zero value otherwise.

### GetSecondaryAlbumTypesOk

`func (o *MetadataProfileResource) GetSecondaryAlbumTypesOk() (*[]ProfileSecondaryAlbumTypeItemResource, bool)`

GetSecondaryAlbumTypesOk returns a tuple with the SecondaryAlbumTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAlbumTypes

`func (o *MetadataProfileResource) SetSecondaryAlbumTypes(v []ProfileSecondaryAlbumTypeItemResource)`

SetSecondaryAlbumTypes sets SecondaryAlbumTypes field to given value.

### HasSecondaryAlbumTypes

`func (o *MetadataProfileResource) HasSecondaryAlbumTypes() bool`

HasSecondaryAlbumTypes returns a boolean if a field has been set.

### SetSecondaryAlbumTypesNil

`func (o *MetadataProfileResource) SetSecondaryAlbumTypesNil(b bool)`

 SetSecondaryAlbumTypesNil sets the value for SecondaryAlbumTypes to be an explicit nil

### UnsetSecondaryAlbumTypes
`func (o *MetadataProfileResource) UnsetSecondaryAlbumTypes()`

UnsetSecondaryAlbumTypes ensures that no value is present for SecondaryAlbumTypes, not even an explicit nil
### GetReleaseStatuses

`func (o *MetadataProfileResource) GetReleaseStatuses() []ProfileReleaseStatusItemResource`

GetReleaseStatuses returns the ReleaseStatuses field if non-nil, zero value otherwise.

### GetReleaseStatusesOk

`func (o *MetadataProfileResource) GetReleaseStatusesOk() (*[]ProfileReleaseStatusItemResource, bool)`

GetReleaseStatusesOk returns a tuple with the ReleaseStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatuses

`func (o *MetadataProfileResource) SetReleaseStatuses(v []ProfileReleaseStatusItemResource)`

SetReleaseStatuses sets ReleaseStatuses field to given value.

### HasReleaseStatuses

`func (o *MetadataProfileResource) HasReleaseStatuses() bool`

HasReleaseStatuses returns a boolean if a field has been set.

### SetReleaseStatusesNil

`func (o *MetadataProfileResource) SetReleaseStatusesNil(b bool)`

 SetReleaseStatusesNil sets the value for ReleaseStatuses to be an explicit nil

### UnsetReleaseStatuses
`func (o *MetadataProfileResource) UnsetReleaseStatuses()`

UnsetReleaseStatuses ensures that no value is present for ReleaseStatuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


