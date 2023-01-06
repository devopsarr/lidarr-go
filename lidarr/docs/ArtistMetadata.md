# ArtistMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForeignArtistId** | Pointer to **NullableString** |  | [optional] 
**OldForeignArtistIds** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Disambiguation** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**ArtistStatusType**](ArtistStatusType.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Links** | Pointer to [**[]Links**](Links.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Members** | Pointer to [**[]Member**](Member.md) |  | [optional] 

## Methods

### NewArtistMetadata

`func NewArtistMetadata() *ArtistMetadata`

NewArtistMetadata instantiates a new ArtistMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistMetadataWithDefaults

`func NewArtistMetadataWithDefaults() *ArtistMetadata`

NewArtistMetadataWithDefaults instantiates a new ArtistMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArtistMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtistMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtistMetadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArtistMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignArtistId

`func (o *ArtistMetadata) GetForeignArtistId() string`

GetForeignArtistId returns the ForeignArtistId field if non-nil, zero value otherwise.

### GetForeignArtistIdOk

`func (o *ArtistMetadata) GetForeignArtistIdOk() (*string, bool)`

GetForeignArtistIdOk returns a tuple with the ForeignArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignArtistId

`func (o *ArtistMetadata) SetForeignArtistId(v string)`

SetForeignArtistId sets ForeignArtistId field to given value.

### HasForeignArtistId

`func (o *ArtistMetadata) HasForeignArtistId() bool`

HasForeignArtistId returns a boolean if a field has been set.

### SetForeignArtistIdNil

`func (o *ArtistMetadata) SetForeignArtistIdNil(b bool)`

 SetForeignArtistIdNil sets the value for ForeignArtistId to be an explicit nil

### UnsetForeignArtistId
`func (o *ArtistMetadata) UnsetForeignArtistId()`

UnsetForeignArtistId ensures that no value is present for ForeignArtistId, not even an explicit nil
### GetOldForeignArtistIds

`func (o *ArtistMetadata) GetOldForeignArtistIds() []string`

GetOldForeignArtistIds returns the OldForeignArtistIds field if non-nil, zero value otherwise.

### GetOldForeignArtistIdsOk

`func (o *ArtistMetadata) GetOldForeignArtistIdsOk() (*[]string, bool)`

GetOldForeignArtistIdsOk returns a tuple with the OldForeignArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldForeignArtistIds

`func (o *ArtistMetadata) SetOldForeignArtistIds(v []string)`

SetOldForeignArtistIds sets OldForeignArtistIds field to given value.

### HasOldForeignArtistIds

`func (o *ArtistMetadata) HasOldForeignArtistIds() bool`

HasOldForeignArtistIds returns a boolean if a field has been set.

### SetOldForeignArtistIdsNil

`func (o *ArtistMetadata) SetOldForeignArtistIdsNil(b bool)`

 SetOldForeignArtistIdsNil sets the value for OldForeignArtistIds to be an explicit nil

### UnsetOldForeignArtistIds
`func (o *ArtistMetadata) UnsetOldForeignArtistIds()`

UnsetOldForeignArtistIds ensures that no value is present for OldForeignArtistIds, not even an explicit nil
### GetName

`func (o *ArtistMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtistMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtistMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtistMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ArtistMetadata) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ArtistMetadata) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAliases

`func (o *ArtistMetadata) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ArtistMetadata) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ArtistMetadata) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ArtistMetadata) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *ArtistMetadata) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *ArtistMetadata) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil
### GetOverview

`func (o *ArtistMetadata) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ArtistMetadata) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ArtistMetadata) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ArtistMetadata) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *ArtistMetadata) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *ArtistMetadata) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetDisambiguation

`func (o *ArtistMetadata) GetDisambiguation() string`

GetDisambiguation returns the Disambiguation field if non-nil, zero value otherwise.

### GetDisambiguationOk

`func (o *ArtistMetadata) GetDisambiguationOk() (*string, bool)`

GetDisambiguationOk returns a tuple with the Disambiguation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguation

`func (o *ArtistMetadata) SetDisambiguation(v string)`

SetDisambiguation sets Disambiguation field to given value.

### HasDisambiguation

`func (o *ArtistMetadata) HasDisambiguation() bool`

HasDisambiguation returns a boolean if a field has been set.

### SetDisambiguationNil

`func (o *ArtistMetadata) SetDisambiguationNil(b bool)`

 SetDisambiguationNil sets the value for Disambiguation to be an explicit nil

### UnsetDisambiguation
`func (o *ArtistMetadata) UnsetDisambiguation()`

UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
### GetType

`func (o *ArtistMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtistMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtistMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArtistMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ArtistMetadata) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ArtistMetadata) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStatus

`func (o *ArtistMetadata) GetStatus() ArtistStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArtistMetadata) GetStatusOk() (*ArtistStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArtistMetadata) SetStatus(v ArtistStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArtistMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImages

`func (o *ArtistMetadata) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ArtistMetadata) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ArtistMetadata) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *ArtistMetadata) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *ArtistMetadata) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *ArtistMetadata) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetLinks

`func (o *ArtistMetadata) GetLinks() []Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ArtistMetadata) GetLinksOk() (*[]Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ArtistMetadata) SetLinks(v []Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ArtistMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ArtistMetadata) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ArtistMetadata) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetGenres

`func (o *ArtistMetadata) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ArtistMetadata) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ArtistMetadata) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ArtistMetadata) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *ArtistMetadata) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *ArtistMetadata) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetRatings

`func (o *ArtistMetadata) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *ArtistMetadata) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *ArtistMetadata) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *ArtistMetadata) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetMembers

`func (o *ArtistMetadata) GetMembers() []Member`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ArtistMetadata) GetMembersOk() (*[]Member, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ArtistMetadata) SetMembers(v []Member)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ArtistMetadata) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *ArtistMetadata) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *ArtistMetadata) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


