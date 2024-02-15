# RootFolderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**DefaultMetadataProfileId** | Pointer to **int32** |  | [optional] 
**DefaultQualityProfileId** | Pointer to **int32** |  | [optional] 
**DefaultMonitorOption** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**DefaultNewItemMonitorOption** | Pointer to [**NewItemMonitorTypes**](NewItemMonitorTypes.md) |  | [optional] 
**DefaultTags** | Pointer to **[]int32** |  | [optional] 
**Accessible** | Pointer to **bool** |  | [optional] 
**FreeSpace** | Pointer to **NullableInt64** |  | [optional] 
**TotalSpace** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewRootFolderResource

`func NewRootFolderResource() *RootFolderResource`

NewRootFolderResource instantiates a new RootFolderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootFolderResourceWithDefaults

`func NewRootFolderResourceWithDefaults() *RootFolderResource`

NewRootFolderResourceWithDefaults instantiates a new RootFolderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RootFolderResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RootFolderResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RootFolderResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RootFolderResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RootFolderResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RootFolderResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RootFolderResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RootFolderResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RootFolderResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RootFolderResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *RootFolderResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RootFolderResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RootFolderResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RootFolderResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *RootFolderResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *RootFolderResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetDefaultMetadataProfileId

`func (o *RootFolderResource) GetDefaultMetadataProfileId() int32`

GetDefaultMetadataProfileId returns the DefaultMetadataProfileId field if non-nil, zero value otherwise.

### GetDefaultMetadataProfileIdOk

`func (o *RootFolderResource) GetDefaultMetadataProfileIdOk() (*int32, bool)`

GetDefaultMetadataProfileIdOk returns a tuple with the DefaultMetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMetadataProfileId

`func (o *RootFolderResource) SetDefaultMetadataProfileId(v int32)`

SetDefaultMetadataProfileId sets DefaultMetadataProfileId field to given value.

### HasDefaultMetadataProfileId

`func (o *RootFolderResource) HasDefaultMetadataProfileId() bool`

HasDefaultMetadataProfileId returns a boolean if a field has been set.

### GetDefaultQualityProfileId

`func (o *RootFolderResource) GetDefaultQualityProfileId() int32`

GetDefaultQualityProfileId returns the DefaultQualityProfileId field if non-nil, zero value otherwise.

### GetDefaultQualityProfileIdOk

`func (o *RootFolderResource) GetDefaultQualityProfileIdOk() (*int32, bool)`

GetDefaultQualityProfileIdOk returns a tuple with the DefaultQualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQualityProfileId

`func (o *RootFolderResource) SetDefaultQualityProfileId(v int32)`

SetDefaultQualityProfileId sets DefaultQualityProfileId field to given value.

### HasDefaultQualityProfileId

`func (o *RootFolderResource) HasDefaultQualityProfileId() bool`

HasDefaultQualityProfileId returns a boolean if a field has been set.

### GetDefaultMonitorOption

`func (o *RootFolderResource) GetDefaultMonitorOption() MonitorTypes`

GetDefaultMonitorOption returns the DefaultMonitorOption field if non-nil, zero value otherwise.

### GetDefaultMonitorOptionOk

`func (o *RootFolderResource) GetDefaultMonitorOptionOk() (*MonitorTypes, bool)`

GetDefaultMonitorOptionOk returns a tuple with the DefaultMonitorOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMonitorOption

`func (o *RootFolderResource) SetDefaultMonitorOption(v MonitorTypes)`

SetDefaultMonitorOption sets DefaultMonitorOption field to given value.

### HasDefaultMonitorOption

`func (o *RootFolderResource) HasDefaultMonitorOption() bool`

HasDefaultMonitorOption returns a boolean if a field has been set.

### GetDefaultNewItemMonitorOption

`func (o *RootFolderResource) GetDefaultNewItemMonitorOption() NewItemMonitorTypes`

GetDefaultNewItemMonitorOption returns the DefaultNewItemMonitorOption field if non-nil, zero value otherwise.

### GetDefaultNewItemMonitorOptionOk

`func (o *RootFolderResource) GetDefaultNewItemMonitorOptionOk() (*NewItemMonitorTypes, bool)`

GetDefaultNewItemMonitorOptionOk returns a tuple with the DefaultNewItemMonitorOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNewItemMonitorOption

`func (o *RootFolderResource) SetDefaultNewItemMonitorOption(v NewItemMonitorTypes)`

SetDefaultNewItemMonitorOption sets DefaultNewItemMonitorOption field to given value.

### HasDefaultNewItemMonitorOption

`func (o *RootFolderResource) HasDefaultNewItemMonitorOption() bool`

HasDefaultNewItemMonitorOption returns a boolean if a field has been set.

### GetDefaultTags

`func (o *RootFolderResource) GetDefaultTags() []int32`

GetDefaultTags returns the DefaultTags field if non-nil, zero value otherwise.

### GetDefaultTagsOk

`func (o *RootFolderResource) GetDefaultTagsOk() (*[]int32, bool)`

GetDefaultTagsOk returns a tuple with the DefaultTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTags

`func (o *RootFolderResource) SetDefaultTags(v []int32)`

SetDefaultTags sets DefaultTags field to given value.

### HasDefaultTags

`func (o *RootFolderResource) HasDefaultTags() bool`

HasDefaultTags returns a boolean if a field has been set.

### SetDefaultTagsNil

`func (o *RootFolderResource) SetDefaultTagsNil(b bool)`

 SetDefaultTagsNil sets the value for DefaultTags to be an explicit nil

### UnsetDefaultTags
`func (o *RootFolderResource) UnsetDefaultTags()`

UnsetDefaultTags ensures that no value is present for DefaultTags, not even an explicit nil
### GetAccessible

`func (o *RootFolderResource) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *RootFolderResource) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *RootFolderResource) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *RootFolderResource) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetFreeSpace

`func (o *RootFolderResource) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *RootFolderResource) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *RootFolderResource) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *RootFolderResource) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *RootFolderResource) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *RootFolderResource) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetTotalSpace

`func (o *RootFolderResource) GetTotalSpace() int64`

GetTotalSpace returns the TotalSpace field if non-nil, zero value otherwise.

### GetTotalSpaceOk

`func (o *RootFolderResource) GetTotalSpaceOk() (*int64, bool)`

GetTotalSpaceOk returns a tuple with the TotalSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpace

`func (o *RootFolderResource) SetTotalSpace(v int64)`

SetTotalSpace sets TotalSpace field to given value.

### HasTotalSpace

`func (o *RootFolderResource) HasTotalSpace() bool`

HasTotalSpace returns a boolean if a field has been set.

### SetTotalSpaceNil

`func (o *RootFolderResource) SetTotalSpaceNil(b bool)`

 SetTotalSpaceNil sets the value for TotalSpace to be an explicit nil

### UnsetTotalSpace
`func (o *RootFolderResource) UnsetTotalSpace()`

UnsetTotalSpace ensures that no value is present for TotalSpace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


