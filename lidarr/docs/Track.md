# Track

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForeignTrackId** | Pointer to **NullableString** |  | [optional] 
**OldForeignTrackIds** | Pointer to **[]string** |  | [optional] 
**ForeignRecordingId** | Pointer to **NullableString** |  | [optional] 
**OldForeignRecordingIds** | Pointer to **[]string** |  | [optional] 
**AlbumReleaseId** | Pointer to **int32** |  | [optional] 
**ArtistMetadataId** | Pointer to **int32** |  | [optional] 
**TrackNumber** | Pointer to **NullableString** |  | [optional] 
**AbsoluteTrackNumber** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Explicit** | Pointer to **bool** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**MediumNumber** | Pointer to **int32** |  | [optional] 
**TrackFileId** | Pointer to **int32** |  | [optional] 
**HasFile** | Pointer to **bool** |  | [optional] [readonly] 
**AlbumRelease** | Pointer to [**AlbumReleaseLazyLoaded**](AlbumReleaseLazyLoaded.md) |  | [optional] 
**ArtistMetadata** | Pointer to [**ArtistMetadataLazyLoaded**](ArtistMetadataLazyLoaded.md) |  | [optional] 
**TrackFile** | Pointer to [**TrackFileLazyLoaded**](TrackFileLazyLoaded.md) |  | [optional] 
**Artist** | Pointer to [**ArtistLazyLoaded**](ArtistLazyLoaded.md) |  | [optional] 
**AlbumId** | Pointer to **int32** |  | [optional] 
**Album** | Pointer to [**Album**](Album.md) |  | [optional] 

## Methods

### NewTrack

`func NewTrack() *Track`

NewTrack instantiates a new Track object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackWithDefaults

`func NewTrackWithDefaults() *Track`

NewTrackWithDefaults instantiates a new Track object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Track) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Track) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Track) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Track) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignTrackId

`func (o *Track) GetForeignTrackId() string`

GetForeignTrackId returns the ForeignTrackId field if non-nil, zero value otherwise.

### GetForeignTrackIdOk

`func (o *Track) GetForeignTrackIdOk() (*string, bool)`

GetForeignTrackIdOk returns a tuple with the ForeignTrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTrackId

`func (o *Track) SetForeignTrackId(v string)`

SetForeignTrackId sets ForeignTrackId field to given value.

### HasForeignTrackId

`func (o *Track) HasForeignTrackId() bool`

HasForeignTrackId returns a boolean if a field has been set.

### SetForeignTrackIdNil

`func (o *Track) SetForeignTrackIdNil(b bool)`

 SetForeignTrackIdNil sets the value for ForeignTrackId to be an explicit nil

### UnsetForeignTrackId
`func (o *Track) UnsetForeignTrackId()`

UnsetForeignTrackId ensures that no value is present for ForeignTrackId, not even an explicit nil
### GetOldForeignTrackIds

`func (o *Track) GetOldForeignTrackIds() []string`

GetOldForeignTrackIds returns the OldForeignTrackIds field if non-nil, zero value otherwise.

### GetOldForeignTrackIdsOk

`func (o *Track) GetOldForeignTrackIdsOk() (*[]string, bool)`

GetOldForeignTrackIdsOk returns a tuple with the OldForeignTrackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldForeignTrackIds

`func (o *Track) SetOldForeignTrackIds(v []string)`

SetOldForeignTrackIds sets OldForeignTrackIds field to given value.

### HasOldForeignTrackIds

`func (o *Track) HasOldForeignTrackIds() bool`

HasOldForeignTrackIds returns a boolean if a field has been set.

### SetOldForeignTrackIdsNil

`func (o *Track) SetOldForeignTrackIdsNil(b bool)`

 SetOldForeignTrackIdsNil sets the value for OldForeignTrackIds to be an explicit nil

### UnsetOldForeignTrackIds
`func (o *Track) UnsetOldForeignTrackIds()`

UnsetOldForeignTrackIds ensures that no value is present for OldForeignTrackIds, not even an explicit nil
### GetForeignRecordingId

`func (o *Track) GetForeignRecordingId() string`

GetForeignRecordingId returns the ForeignRecordingId field if non-nil, zero value otherwise.

### GetForeignRecordingIdOk

`func (o *Track) GetForeignRecordingIdOk() (*string, bool)`

GetForeignRecordingIdOk returns a tuple with the ForeignRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignRecordingId

`func (o *Track) SetForeignRecordingId(v string)`

SetForeignRecordingId sets ForeignRecordingId field to given value.

### HasForeignRecordingId

`func (o *Track) HasForeignRecordingId() bool`

HasForeignRecordingId returns a boolean if a field has been set.

### SetForeignRecordingIdNil

`func (o *Track) SetForeignRecordingIdNil(b bool)`

 SetForeignRecordingIdNil sets the value for ForeignRecordingId to be an explicit nil

### UnsetForeignRecordingId
`func (o *Track) UnsetForeignRecordingId()`

UnsetForeignRecordingId ensures that no value is present for ForeignRecordingId, not even an explicit nil
### GetOldForeignRecordingIds

`func (o *Track) GetOldForeignRecordingIds() []string`

GetOldForeignRecordingIds returns the OldForeignRecordingIds field if non-nil, zero value otherwise.

### GetOldForeignRecordingIdsOk

`func (o *Track) GetOldForeignRecordingIdsOk() (*[]string, bool)`

GetOldForeignRecordingIdsOk returns a tuple with the OldForeignRecordingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldForeignRecordingIds

`func (o *Track) SetOldForeignRecordingIds(v []string)`

SetOldForeignRecordingIds sets OldForeignRecordingIds field to given value.

### HasOldForeignRecordingIds

`func (o *Track) HasOldForeignRecordingIds() bool`

HasOldForeignRecordingIds returns a boolean if a field has been set.

### SetOldForeignRecordingIdsNil

`func (o *Track) SetOldForeignRecordingIdsNil(b bool)`

 SetOldForeignRecordingIdsNil sets the value for OldForeignRecordingIds to be an explicit nil

### UnsetOldForeignRecordingIds
`func (o *Track) UnsetOldForeignRecordingIds()`

UnsetOldForeignRecordingIds ensures that no value is present for OldForeignRecordingIds, not even an explicit nil
### GetAlbumReleaseId

`func (o *Track) GetAlbumReleaseId() int32`

GetAlbumReleaseId returns the AlbumReleaseId field if non-nil, zero value otherwise.

### GetAlbumReleaseIdOk

`func (o *Track) GetAlbumReleaseIdOk() (*int32, bool)`

GetAlbumReleaseIdOk returns a tuple with the AlbumReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumReleaseId

`func (o *Track) SetAlbumReleaseId(v int32)`

SetAlbumReleaseId sets AlbumReleaseId field to given value.

### HasAlbumReleaseId

`func (o *Track) HasAlbumReleaseId() bool`

HasAlbumReleaseId returns a boolean if a field has been set.

### GetArtistMetadataId

`func (o *Track) GetArtistMetadataId() int32`

GetArtistMetadataId returns the ArtistMetadataId field if non-nil, zero value otherwise.

### GetArtistMetadataIdOk

`func (o *Track) GetArtistMetadataIdOk() (*int32, bool)`

GetArtistMetadataIdOk returns a tuple with the ArtistMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMetadataId

`func (o *Track) SetArtistMetadataId(v int32)`

SetArtistMetadataId sets ArtistMetadataId field to given value.

### HasArtistMetadataId

`func (o *Track) HasArtistMetadataId() bool`

HasArtistMetadataId returns a boolean if a field has been set.

### GetTrackNumber

`func (o *Track) GetTrackNumber() string`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *Track) GetTrackNumberOk() (*string, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *Track) SetTrackNumber(v string)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *Track) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### SetTrackNumberNil

`func (o *Track) SetTrackNumberNil(b bool)`

 SetTrackNumberNil sets the value for TrackNumber to be an explicit nil

### UnsetTrackNumber
`func (o *Track) UnsetTrackNumber()`

UnsetTrackNumber ensures that no value is present for TrackNumber, not even an explicit nil
### GetAbsoluteTrackNumber

`func (o *Track) GetAbsoluteTrackNumber() int32`

GetAbsoluteTrackNumber returns the AbsoluteTrackNumber field if non-nil, zero value otherwise.

### GetAbsoluteTrackNumberOk

`func (o *Track) GetAbsoluteTrackNumberOk() (*int32, bool)`

GetAbsoluteTrackNumberOk returns a tuple with the AbsoluteTrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteTrackNumber

`func (o *Track) SetAbsoluteTrackNumber(v int32)`

SetAbsoluteTrackNumber sets AbsoluteTrackNumber field to given value.

### HasAbsoluteTrackNumber

`func (o *Track) HasAbsoluteTrackNumber() bool`

HasAbsoluteTrackNumber returns a boolean if a field has been set.

### GetTitle

`func (o *Track) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Track) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Track) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Track) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Track) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Track) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDuration

`func (o *Track) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Track) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Track) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Track) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetExplicit

`func (o *Track) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *Track) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *Track) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *Track) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetRatings

`func (o *Track) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *Track) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *Track) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *Track) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetMediumNumber

`func (o *Track) GetMediumNumber() int32`

GetMediumNumber returns the MediumNumber field if non-nil, zero value otherwise.

### GetMediumNumberOk

`func (o *Track) GetMediumNumberOk() (*int32, bool)`

GetMediumNumberOk returns a tuple with the MediumNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumNumber

`func (o *Track) SetMediumNumber(v int32)`

SetMediumNumber sets MediumNumber field to given value.

### HasMediumNumber

`func (o *Track) HasMediumNumber() bool`

HasMediumNumber returns a boolean if a field has been set.

### GetTrackFileId

`func (o *Track) GetTrackFileId() int32`

GetTrackFileId returns the TrackFileId field if non-nil, zero value otherwise.

### GetTrackFileIdOk

`func (o *Track) GetTrackFileIdOk() (*int32, bool)`

GetTrackFileIdOk returns a tuple with the TrackFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFileId

`func (o *Track) SetTrackFileId(v int32)`

SetTrackFileId sets TrackFileId field to given value.

### HasTrackFileId

`func (o *Track) HasTrackFileId() bool`

HasTrackFileId returns a boolean if a field has been set.

### GetHasFile

`func (o *Track) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *Track) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *Track) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *Track) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetAlbumRelease

`func (o *Track) GetAlbumRelease() AlbumReleaseLazyLoaded`

GetAlbumRelease returns the AlbumRelease field if non-nil, zero value otherwise.

### GetAlbumReleaseOk

`func (o *Track) GetAlbumReleaseOk() (*AlbumReleaseLazyLoaded, bool)`

GetAlbumReleaseOk returns a tuple with the AlbumRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumRelease

`func (o *Track) SetAlbumRelease(v AlbumReleaseLazyLoaded)`

SetAlbumRelease sets AlbumRelease field to given value.

### HasAlbumRelease

`func (o *Track) HasAlbumRelease() bool`

HasAlbumRelease returns a boolean if a field has been set.

### GetArtistMetadata

`func (o *Track) GetArtistMetadata() ArtistMetadataLazyLoaded`

GetArtistMetadata returns the ArtistMetadata field if non-nil, zero value otherwise.

### GetArtistMetadataOk

`func (o *Track) GetArtistMetadataOk() (*ArtistMetadataLazyLoaded, bool)`

GetArtistMetadataOk returns a tuple with the ArtistMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistMetadata

`func (o *Track) SetArtistMetadata(v ArtistMetadataLazyLoaded)`

SetArtistMetadata sets ArtistMetadata field to given value.

### HasArtistMetadata

`func (o *Track) HasArtistMetadata() bool`

HasArtistMetadata returns a boolean if a field has been set.

### GetTrackFile

`func (o *Track) GetTrackFile() TrackFileLazyLoaded`

GetTrackFile returns the TrackFile field if non-nil, zero value otherwise.

### GetTrackFileOk

`func (o *Track) GetTrackFileOk() (*TrackFileLazyLoaded, bool)`

GetTrackFileOk returns a tuple with the TrackFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFile

`func (o *Track) SetTrackFile(v TrackFileLazyLoaded)`

SetTrackFile sets TrackFile field to given value.

### HasTrackFile

`func (o *Track) HasTrackFile() bool`

HasTrackFile returns a boolean if a field has been set.

### GetArtist

`func (o *Track) GetArtist() ArtistLazyLoaded`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *Track) GetArtistOk() (*ArtistLazyLoaded, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *Track) SetArtist(v ArtistLazyLoaded)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *Track) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetAlbumId

`func (o *Track) GetAlbumId() int32`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *Track) GetAlbumIdOk() (*int32, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *Track) SetAlbumId(v int32)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *Track) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetAlbum

`func (o *Track) GetAlbum() Album`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *Track) GetAlbumOk() (*Album, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *Track) SetAlbum(v Album)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *Track) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


