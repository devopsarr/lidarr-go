# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Instrument** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Member) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Member) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Member) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInstrument

`func (o *Member) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *Member) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *Member) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *Member) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### SetInstrumentNil

`func (o *Member) SetInstrumentNil(b bool)`

 SetInstrumentNil sets the value for Instrument to be an explicit nil

### UnsetInstrument
`func (o *Member) UnsetInstrument()`

UnsetInstrument ensures that no value is present for Instrument, not even an explicit nil
### GetImages

`func (o *Member) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Member) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Member) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *Member) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *Member) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *Member) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


