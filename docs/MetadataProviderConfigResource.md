# MetadataProviderConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MetadataSource** | Pointer to **NullableString** |  | [optional] 
**WriteAudioTags** | Pointer to [**WriteAudioTagsType**](WriteAudioTagsType.md) |  | [optional] 
**ScrubAudioTags** | Pointer to **bool** |  | [optional] 
**EmbedCoverArt** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetadataProviderConfigResource

`func NewMetadataProviderConfigResource() *MetadataProviderConfigResource`

NewMetadataProviderConfigResource instantiates a new MetadataProviderConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProviderConfigResourceWithDefaults

`func NewMetadataProviderConfigResourceWithDefaults() *MetadataProviderConfigResource`

NewMetadataProviderConfigResourceWithDefaults instantiates a new MetadataProviderConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataProviderConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataProviderConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataProviderConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataProviderConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataSource

`func (o *MetadataProviderConfigResource) GetMetadataSource() string`

GetMetadataSource returns the MetadataSource field if non-nil, zero value otherwise.

### GetMetadataSourceOk

`func (o *MetadataProviderConfigResource) GetMetadataSourceOk() (*string, bool)`

GetMetadataSourceOk returns a tuple with the MetadataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSource

`func (o *MetadataProviderConfigResource) SetMetadataSource(v string)`

SetMetadataSource sets MetadataSource field to given value.

### HasMetadataSource

`func (o *MetadataProviderConfigResource) HasMetadataSource() bool`

HasMetadataSource returns a boolean if a field has been set.

### SetMetadataSourceNil

`func (o *MetadataProviderConfigResource) SetMetadataSourceNil(b bool)`

 SetMetadataSourceNil sets the value for MetadataSource to be an explicit nil

### UnsetMetadataSource
`func (o *MetadataProviderConfigResource) UnsetMetadataSource()`

UnsetMetadataSource ensures that no value is present for MetadataSource, not even an explicit nil
### GetWriteAudioTags

`func (o *MetadataProviderConfigResource) GetWriteAudioTags() WriteAudioTagsType`

GetWriteAudioTags returns the WriteAudioTags field if non-nil, zero value otherwise.

### GetWriteAudioTagsOk

`func (o *MetadataProviderConfigResource) GetWriteAudioTagsOk() (*WriteAudioTagsType, bool)`

GetWriteAudioTagsOk returns a tuple with the WriteAudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAudioTags

`func (o *MetadataProviderConfigResource) SetWriteAudioTags(v WriteAudioTagsType)`

SetWriteAudioTags sets WriteAudioTags field to given value.

### HasWriteAudioTags

`func (o *MetadataProviderConfigResource) HasWriteAudioTags() bool`

HasWriteAudioTags returns a boolean if a field has been set.

### GetScrubAudioTags

`func (o *MetadataProviderConfigResource) GetScrubAudioTags() bool`

GetScrubAudioTags returns the ScrubAudioTags field if non-nil, zero value otherwise.

### GetScrubAudioTagsOk

`func (o *MetadataProviderConfigResource) GetScrubAudioTagsOk() (*bool, bool)`

GetScrubAudioTagsOk returns a tuple with the ScrubAudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubAudioTags

`func (o *MetadataProviderConfigResource) SetScrubAudioTags(v bool)`

SetScrubAudioTags sets ScrubAudioTags field to given value.

### HasScrubAudioTags

`func (o *MetadataProviderConfigResource) HasScrubAudioTags() bool`

HasScrubAudioTags returns a boolean if a field has been set.

### GetEmbedCoverArt

`func (o *MetadataProviderConfigResource) GetEmbedCoverArt() bool`

GetEmbedCoverArt returns the EmbedCoverArt field if non-nil, zero value otherwise.

### GetEmbedCoverArtOk

`func (o *MetadataProviderConfigResource) GetEmbedCoverArtOk() (*bool, bool)`

GetEmbedCoverArtOk returns a tuple with the EmbedCoverArt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedCoverArt

`func (o *MetadataProviderConfigResource) SetEmbedCoverArt(v bool)`

SetEmbedCoverArt sets EmbedCoverArt field to given value.

### HasEmbedCoverArt

`func (o *MetadataProviderConfigResource) HasEmbedCoverArt() bool`

HasEmbedCoverArt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


