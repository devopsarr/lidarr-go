# TrackResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ArtistId** | Pointer to **int32** |  | [optional] 
**ForeignTrackId** | Pointer to **NullableString** |  | [optional] 
**ForeignRecordingId** | Pointer to **NullableString** |  | [optional] 
**TrackFileId** | Pointer to **int32** |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**Explicit** | Pointer to **bool** |  | [optional] 
**AbsoluteTrackNumber** | Pointer to **int32** |  | [optional] 
**TrackNumber** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**TrackFile** | Pointer to [**TrackFileResource**](TrackFileResource.md) |  | [optional] 
**MediumNumber** | Pointer to **int32** |  | [optional] 
**HasFile** | Pointer to **bool** |  | [optional] 
**Artist** | Pointer to [**ArtistResource**](ArtistResource.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 

## Methods

### NewTrackResource

`func NewTrackResource() *TrackResource`

NewTrackResource instantiates a new TrackResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackResourceWithDefaults

`func NewTrackResourceWithDefaults() *TrackResource`

NewTrackResourceWithDefaults instantiates a new TrackResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrackResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtistId

`func (o *TrackResource) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *TrackResource) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *TrackResource) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.

### HasArtistId

`func (o *TrackResource) HasArtistId() bool`

HasArtistId returns a boolean if a field has been set.

### GetForeignTrackId

`func (o *TrackResource) GetForeignTrackId() string`

GetForeignTrackId returns the ForeignTrackId field if non-nil, zero value otherwise.

### GetForeignTrackIdOk

`func (o *TrackResource) GetForeignTrackIdOk() (*string, bool)`

GetForeignTrackIdOk returns a tuple with the ForeignTrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTrackId

`func (o *TrackResource) SetForeignTrackId(v string)`

SetForeignTrackId sets ForeignTrackId field to given value.

### HasForeignTrackId

`func (o *TrackResource) HasForeignTrackId() bool`

HasForeignTrackId returns a boolean if a field has been set.

### SetForeignTrackIdNil

`func (o *TrackResource) SetForeignTrackIdNil(b bool)`

 SetForeignTrackIdNil sets the value for ForeignTrackId to be an explicit nil

### UnsetForeignTrackId
`func (o *TrackResource) UnsetForeignTrackId()`

UnsetForeignTrackId ensures that no value is present for ForeignTrackId, not even an explicit nil
### GetForeignRecordingId

`func (o *TrackResource) GetForeignRecordingId() string`

GetForeignRecordingId returns the ForeignRecordingId field if non-nil, zero value otherwise.

### GetForeignRecordingIdOk

`func (o *TrackResource) GetForeignRecordingIdOk() (*string, bool)`

GetForeignRecordingIdOk returns a tuple with the ForeignRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignRecordingId

`func (o *TrackResource) SetForeignRecordingId(v string)`

SetForeignRecordingId sets ForeignRecordingId field to given value.

### HasForeignRecordingId

`func (o *TrackResource) HasForeignRecordingId() bool`

HasForeignRecordingId returns a boolean if a field has been set.

### SetForeignRecordingIdNil

`func (o *TrackResource) SetForeignRecordingIdNil(b bool)`

 SetForeignRecordingIdNil sets the value for ForeignRecordingId to be an explicit nil

### UnsetForeignRecordingId
`func (o *TrackResource) UnsetForeignRecordingId()`

UnsetForeignRecordingId ensures that no value is present for ForeignRecordingId, not even an explicit nil
### GetTrackFileId

`func (o *TrackResource) GetTrackFileId() int32`

GetTrackFileId returns the TrackFileId field if non-nil, zero value otherwise.

### GetTrackFileIdOk

`func (o *TrackResource) GetTrackFileIdOk() (*int32, bool)`

GetTrackFileIdOk returns a tuple with the TrackFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFileId

`func (o *TrackResource) SetTrackFileId(v int32)`

SetTrackFileId sets TrackFileId field to given value.

### HasTrackFileId

`func (o *TrackResource) HasTrackFileId() bool`

HasTrackFileId returns a boolean if a field has been set.

### GetAlbumId

`func (o *TrackResource) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *TrackResource) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *TrackResource) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *TrackResource) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetExplicit

`func (o *TrackResource) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *TrackResource) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *TrackResource) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *TrackResource) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetAbsoluteTrackNumber

`func (o *TrackResource) GetAbsoluteTrackNumber() int32`

GetAbsoluteTrackNumber returns the AbsoluteTrackNumber field if non-nil, zero value otherwise.

### GetAbsoluteTrackNumberOk

`func (o *TrackResource) GetAbsoluteTrackNumberOk() (*int32, bool)`

GetAbsoluteTrackNumberOk returns a tuple with the AbsoluteTrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteTrackNumber

`func (o *TrackResource) SetAbsoluteTrackNumber(v int32)`

SetAbsoluteTrackNumber sets AbsoluteTrackNumber field to given value.

### HasAbsoluteTrackNumber

`func (o *TrackResource) HasAbsoluteTrackNumber() bool`

HasAbsoluteTrackNumber returns a boolean if a field has been set.

### GetTrackNumber

`func (o *TrackResource) GetTrackNumber() string`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *TrackResource) GetTrackNumberOk() (*string, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *TrackResource) SetTrackNumber(v string)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *TrackResource) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### SetTrackNumberNil

`func (o *TrackResource) SetTrackNumberNil(b bool)`

 SetTrackNumberNil sets the value for TrackNumber to be an explicit nil

### UnsetTrackNumber
`func (o *TrackResource) UnsetTrackNumber()`

UnsetTrackNumber ensures that no value is present for TrackNumber, not even an explicit nil
### GetTitle

`func (o *TrackResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrackResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrackResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrackResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TrackResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TrackResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDuration

`func (o *TrackResource) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TrackResource) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TrackResource) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TrackResource) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTrackFile

`func (o *TrackResource) GetTrackFile() TrackFileResource`

GetTrackFile returns the TrackFile field if non-nil, zero value otherwise.

### GetTrackFileOk

`func (o *TrackResource) GetTrackFileOk() (*TrackFileResource, bool)`

GetTrackFileOk returns a tuple with the TrackFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFile

`func (o *TrackResource) SetTrackFile(v TrackFileResource)`

SetTrackFile sets TrackFile field to given value.

### HasTrackFile

`func (o *TrackResource) HasTrackFile() bool`

HasTrackFile returns a boolean if a field has been set.

### GetMediumNumber

`func (o *TrackResource) GetMediumNumber() int32`

GetMediumNumber returns the MediumNumber field if non-nil, zero value otherwise.

### GetMediumNumberOk

`func (o *TrackResource) GetMediumNumberOk() (*int32, bool)`

GetMediumNumberOk returns a tuple with the MediumNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumNumber

`func (o *TrackResource) SetMediumNumber(v int32)`

SetMediumNumber sets MediumNumber field to given value.

### HasMediumNumber

`func (o *TrackResource) HasMediumNumber() bool`

HasMediumNumber returns a boolean if a field has been set.

### GetHasFile

`func (o *TrackResource) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *TrackResource) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *TrackResource) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *TrackResource) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetArtist

`func (o *TrackResource) GetArtist() ArtistResource`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *TrackResource) GetArtistOk() (*ArtistResource, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *TrackResource) SetArtist(v ArtistResource)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *TrackResource) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetRatings

`func (o *TrackResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *TrackResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *TrackResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *TrackResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


