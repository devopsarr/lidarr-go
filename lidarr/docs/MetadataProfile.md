# MetadataProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**PrimaryAlbumTypes** | Pointer to [**[]ProfilePrimaryAlbumTypeItem**](ProfilePrimaryAlbumTypeItem.md) |  | [optional] 
**SecondaryAlbumTypes** | Pointer to [**[]ProfileSecondaryAlbumTypeItem**](ProfileSecondaryAlbumTypeItem.md) |  | [optional] 
**ReleaseStatuses** | Pointer to [**[]ProfileReleaseStatusItem**](ProfileReleaseStatusItem.md) |  | [optional] 

## Methods

### NewMetadataProfile

`func NewMetadataProfile() *MetadataProfile`

NewMetadataProfile instantiates a new MetadataProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProfileWithDefaults

`func NewMetadataProfileWithDefaults() *MetadataProfile`

NewMetadataProfileWithDefaults instantiates a new MetadataProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrimaryAlbumTypes

`func (o *MetadataProfile) GetPrimaryAlbumTypes() []ProfilePrimaryAlbumTypeItem`

GetPrimaryAlbumTypes returns the PrimaryAlbumTypes field if non-nil, zero value otherwise.

### GetPrimaryAlbumTypesOk

`func (o *MetadataProfile) GetPrimaryAlbumTypesOk() (*[]ProfilePrimaryAlbumTypeItem, bool)`

GetPrimaryAlbumTypesOk returns a tuple with the PrimaryAlbumTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAlbumTypes

`func (o *MetadataProfile) SetPrimaryAlbumTypes(v []ProfilePrimaryAlbumTypeItem)`

SetPrimaryAlbumTypes sets PrimaryAlbumTypes field to given value.

### HasPrimaryAlbumTypes

`func (o *MetadataProfile) HasPrimaryAlbumTypes() bool`

HasPrimaryAlbumTypes returns a boolean if a field has been set.

### SetPrimaryAlbumTypesNil

`func (o *MetadataProfile) SetPrimaryAlbumTypesNil(b bool)`

 SetPrimaryAlbumTypesNil sets the value for PrimaryAlbumTypes to be an explicit nil

### UnsetPrimaryAlbumTypes
`func (o *MetadataProfile) UnsetPrimaryAlbumTypes()`

UnsetPrimaryAlbumTypes ensures that no value is present for PrimaryAlbumTypes, not even an explicit nil
### GetSecondaryAlbumTypes

`func (o *MetadataProfile) GetSecondaryAlbumTypes() []ProfileSecondaryAlbumTypeItem`

GetSecondaryAlbumTypes returns the SecondaryAlbumTypes field if non-nil, zero value otherwise.

### GetSecondaryAlbumTypesOk

`func (o *MetadataProfile) GetSecondaryAlbumTypesOk() (*[]ProfileSecondaryAlbumTypeItem, bool)`

GetSecondaryAlbumTypesOk returns a tuple with the SecondaryAlbumTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAlbumTypes

`func (o *MetadataProfile) SetSecondaryAlbumTypes(v []ProfileSecondaryAlbumTypeItem)`

SetSecondaryAlbumTypes sets SecondaryAlbumTypes field to given value.

### HasSecondaryAlbumTypes

`func (o *MetadataProfile) HasSecondaryAlbumTypes() bool`

HasSecondaryAlbumTypes returns a boolean if a field has been set.

### SetSecondaryAlbumTypesNil

`func (o *MetadataProfile) SetSecondaryAlbumTypesNil(b bool)`

 SetSecondaryAlbumTypesNil sets the value for SecondaryAlbumTypes to be an explicit nil

### UnsetSecondaryAlbumTypes
`func (o *MetadataProfile) UnsetSecondaryAlbumTypes()`

UnsetSecondaryAlbumTypes ensures that no value is present for SecondaryAlbumTypes, not even an explicit nil
### GetReleaseStatuses

`func (o *MetadataProfile) GetReleaseStatuses() []ProfileReleaseStatusItem`

GetReleaseStatuses returns the ReleaseStatuses field if non-nil, zero value otherwise.

### GetReleaseStatusesOk

`func (o *MetadataProfile) GetReleaseStatusesOk() (*[]ProfileReleaseStatusItem, bool)`

GetReleaseStatusesOk returns a tuple with the ReleaseStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatuses

`func (o *MetadataProfile) SetReleaseStatuses(v []ProfileReleaseStatusItem)`

SetReleaseStatuses sets ReleaseStatuses field to given value.

### HasReleaseStatuses

`func (o *MetadataProfile) HasReleaseStatuses() bool`

HasReleaseStatuses returns a boolean if a field has been set.

### SetReleaseStatusesNil

`func (o *MetadataProfile) SetReleaseStatusesNil(b bool)`

 SetReleaseStatusesNil sets the value for ReleaseStatuses to be an explicit nil

### UnsetReleaseStatuses
`func (o *MetadataProfile) UnsetReleaseStatuses()`

UnsetReleaseStatuses ensures that no value is present for ReleaseStatuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


