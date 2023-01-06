/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// ArtistMetadata struct for ArtistMetadata
type ArtistMetadata struct {
	Id *int32 `json:"id,omitempty"`
	ForeignArtistId NullableString `json:"foreignArtistId,omitempty"`
	OldForeignArtistIds []*string `json:"oldForeignArtistIds,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Aliases []*string `json:"aliases,omitempty"`
	Overview NullableString `json:"overview,omitempty"`
	Disambiguation NullableString `json:"disambiguation,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Status *ArtistStatusType `json:"status,omitempty"`
	Images []*MediaCover `json:"images,omitempty"`
	Links []*Links `json:"links,omitempty"`
	Genres []*string `json:"genres,omitempty"`
	Ratings *Ratings `json:"ratings,omitempty"`
	Members []*Member `json:"members,omitempty"`
}

// NewArtistMetadata instantiates a new ArtistMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtistMetadata() *ArtistMetadata {
	this := ArtistMetadata{}
	return &this
}

// NewArtistMetadataWithDefaults instantiates a new ArtistMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtistMetadataWithDefaults() *ArtistMetadata {
	this := ArtistMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArtistMetadata) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistMetadata) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArtistMetadata) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ArtistMetadata) SetId(v int32) {
	o.Id = &v
}

// GetForeignArtistId returns the ForeignArtistId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetForeignArtistId() string {
	if o == nil || isNil(o.ForeignArtistId.Get()) {
		var ret string
		return ret
	}
	return *o.ForeignArtistId.Get()
}

// GetForeignArtistIdOk returns a tuple with the ForeignArtistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetForeignArtistIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ForeignArtistId.Get(), o.ForeignArtistId.IsSet()
}

// HasForeignArtistId returns a boolean if a field has been set.
func (o *ArtistMetadata) HasForeignArtistId() bool {
	if o != nil && o.ForeignArtistId.IsSet() {
		return true
	}

	return false
}

// SetForeignArtistId gets a reference to the given NullableString and assigns it to the ForeignArtistId field.
func (o *ArtistMetadata) SetForeignArtistId(v string) {
	o.ForeignArtistId.Set(&v)
}
// SetForeignArtistIdNil sets the value for ForeignArtistId to be an explicit nil
func (o *ArtistMetadata) SetForeignArtistIdNil() {
	o.ForeignArtistId.Set(nil)
}

// UnsetForeignArtistId ensures that no value is present for ForeignArtistId, not even an explicit nil
func (o *ArtistMetadata) UnsetForeignArtistId() {
	o.ForeignArtistId.Unset()
}

// GetOldForeignArtistIds returns the OldForeignArtistIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetOldForeignArtistIds() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.OldForeignArtistIds
}

// GetOldForeignArtistIdsOk returns a tuple with the OldForeignArtistIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetOldForeignArtistIdsOk() ([]*string, bool) {
	if o == nil || isNil(o.OldForeignArtistIds) {
    return nil, false
	}
	return o.OldForeignArtistIds, true
}

// HasOldForeignArtistIds returns a boolean if a field has been set.
func (o *ArtistMetadata) HasOldForeignArtistIds() bool {
	if o != nil && isNil(o.OldForeignArtistIds) {
		return true
	}

	return false
}

// SetOldForeignArtistIds gets a reference to the given []string and assigns it to the OldForeignArtistIds field.
func (o *ArtistMetadata) SetOldForeignArtistIds(v []*string) {
	o.OldForeignArtistIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ArtistMetadata) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ArtistMetadata) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ArtistMetadata) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ArtistMetadata) UnsetName() {
	o.Name.Unset()
}

// GetAliases returns the Aliases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetAliases() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetAliasesOk() ([]*string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *ArtistMetadata) HasAliases() bool {
	if o != nil && isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *ArtistMetadata) SetAliases(v []*string) {
	o.Aliases = v
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetOverview() string {
	if o == nil || isNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetOverviewOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *ArtistMetadata) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *ArtistMetadata) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *ArtistMetadata) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *ArtistMetadata) UnsetOverview() {
	o.Overview.Unset()
}

// GetDisambiguation returns the Disambiguation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetDisambiguation() string {
	if o == nil || isNil(o.Disambiguation.Get()) {
		var ret string
		return ret
	}
	return *o.Disambiguation.Get()
}

// GetDisambiguationOk returns a tuple with the Disambiguation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetDisambiguationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Disambiguation.Get(), o.Disambiguation.IsSet()
}

// HasDisambiguation returns a boolean if a field has been set.
func (o *ArtistMetadata) HasDisambiguation() bool {
	if o != nil && o.Disambiguation.IsSet() {
		return true
	}

	return false
}

// SetDisambiguation gets a reference to the given NullableString and assigns it to the Disambiguation field.
func (o *ArtistMetadata) SetDisambiguation(v string) {
	o.Disambiguation.Set(&v)
}
// SetDisambiguationNil sets the value for Disambiguation to be an explicit nil
func (o *ArtistMetadata) SetDisambiguationNil() {
	o.Disambiguation.Set(nil)
}

// UnsetDisambiguation ensures that no value is present for Disambiguation, not even an explicit nil
func (o *ArtistMetadata) UnsetDisambiguation() {
	o.Disambiguation.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetType() string {
	if o == nil || isNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ArtistMetadata) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ArtistMetadata) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ArtistMetadata) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ArtistMetadata) UnsetType() {
	o.Type.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ArtistMetadata) GetStatus() ArtistStatusType {
	if o == nil || isNil(o.Status) {
		var ret ArtistStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistMetadata) GetStatusOk() (*ArtistStatusType, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ArtistMetadata) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ArtistStatusType and assigns it to the Status field.
func (o *ArtistMetadata) SetStatus(v ArtistStatusType) {
	o.Status = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetImages() []*MediaCover {
	if o == nil {
		var ret []*MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetImagesOk() ([]*MediaCover, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ArtistMetadata) HasImages() bool {
	if o != nil && isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *ArtistMetadata) SetImages(v []*MediaCover) {
	o.Images = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetLinks() []*Links {
	if o == nil {
		var ret []*Links
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetLinksOk() ([]*Links, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ArtistMetadata) HasLinks() bool {
	if o != nil && isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Links and assigns it to the Links field.
func (o *ArtistMetadata) SetLinks(v []*Links) {
	o.Links = v
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetGenres() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetGenresOk() ([]*string, bool) {
	if o == nil || isNil(o.Genres) {
    return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *ArtistMetadata) HasGenres() bool {
	if o != nil && isNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *ArtistMetadata) SetGenres(v []*string) {
	o.Genres = v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *ArtistMetadata) GetRatings() Ratings {
	if o == nil || isNil(o.Ratings) {
		var ret Ratings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtistMetadata) GetRatingsOk() (*Ratings, bool) {
	if o == nil || isNil(o.Ratings) {
    return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *ArtistMetadata) HasRatings() bool {
	if o != nil && !isNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given Ratings and assigns it to the Ratings field.
func (o *ArtistMetadata) SetRatings(v Ratings) {
	o.Ratings = &v
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtistMetadata) GetMembers() []*Member {
	if o == nil {
		var ret []*Member
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtistMetadata) GetMembersOk() ([]*Member, bool) {
	if o == nil || isNil(o.Members) {
    return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ArtistMetadata) HasMembers() bool {
	if o != nil && isNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Member and assigns it to the Members field.
func (o *ArtistMetadata) SetMembers(v []*Member) {
	o.Members = v
}

func (o ArtistMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ForeignArtistId.IsSet() {
		toSerialize["foreignArtistId"] = o.ForeignArtistId.Get()
	}
	if o.OldForeignArtistIds != nil {
		toSerialize["oldForeignArtistIds"] = o.OldForeignArtistIds
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Aliases != nil {
		toSerialize["aliases"] = o.Aliases
	}
	if o.Overview.IsSet() {
		toSerialize["overview"] = o.Overview.Get()
	}
	if o.Disambiguation.IsSet() {
		toSerialize["disambiguation"] = o.Disambiguation.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if !isNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableArtistMetadata struct {
	value *ArtistMetadata
	isSet bool
}

func (v NullableArtistMetadata) Get() *ArtistMetadata {
	return v.value
}

func (v *NullableArtistMetadata) Set(val *ArtistMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableArtistMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableArtistMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtistMetadata(val *ArtistMetadata) *NullableArtistMetadata {
	return &NullableArtistMetadata{value: val, isSet: true}
}

func (v NullableArtistMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtistMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


