/*
Lidarr

Lidarr API docs

API version: v2.6.4.4402
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"encoding/json"
)

// checks if the MediaManagementConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaManagementConfigResource{}

// MediaManagementConfigResource struct for MediaManagementConfigResource
type MediaManagementConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	AutoUnmonitorPreviouslyDownloadedTracks *bool `json:"autoUnmonitorPreviouslyDownloadedTracks,omitempty"`
	RecycleBin NullableString `json:"recycleBin,omitempty"`
	RecycleBinCleanupDays *int32 `json:"recycleBinCleanupDays,omitempty"`
	DownloadPropersAndRepacks *ProperDownloadTypes `json:"downloadPropersAndRepacks,omitempty"`
	CreateEmptyArtistFolders *bool `json:"createEmptyArtistFolders,omitempty"`
	DeleteEmptyFolders *bool `json:"deleteEmptyFolders,omitempty"`
	FileDate *FileDateType `json:"fileDate,omitempty"`
	WatchLibraryForChanges *bool `json:"watchLibraryForChanges,omitempty"`
	RescanAfterRefresh *RescanAfterRefreshType `json:"rescanAfterRefresh,omitempty"`
	AllowFingerprinting *AllowFingerprinting `json:"allowFingerprinting,omitempty"`
	SetPermissionsLinux *bool `json:"setPermissionsLinux,omitempty"`
	ChmodFolder NullableString `json:"chmodFolder,omitempty"`
	ChownGroup NullableString `json:"chownGroup,omitempty"`
	SkipFreeSpaceCheckWhenImporting *bool `json:"skipFreeSpaceCheckWhenImporting,omitempty"`
	MinimumFreeSpaceWhenImporting *int32 `json:"minimumFreeSpaceWhenImporting,omitempty"`
	CopyUsingHardlinks *bool `json:"copyUsingHardlinks,omitempty"`
	ImportExtraFiles *bool `json:"importExtraFiles,omitempty"`
	ExtraFileExtensions NullableString `json:"extraFileExtensions,omitempty"`
}

// NewMediaManagementConfigResource instantiates a new MediaManagementConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaManagementConfigResource() *MediaManagementConfigResource {
	this := MediaManagementConfigResource{}
	return &this
}

// NewMediaManagementConfigResourceWithDefaults instantiates a new MediaManagementConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaManagementConfigResourceWithDefaults() *MediaManagementConfigResource {
	this := MediaManagementConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MediaManagementConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetAutoUnmonitorPreviouslyDownloadedTracks returns the AutoUnmonitorPreviouslyDownloadedTracks field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetAutoUnmonitorPreviouslyDownloadedTracks() bool {
	if o == nil || IsNil(o.AutoUnmonitorPreviouslyDownloadedTracks) {
		var ret bool
		return ret
	}
	return *o.AutoUnmonitorPreviouslyDownloadedTracks
}

// GetAutoUnmonitorPreviouslyDownloadedTracksOk returns a tuple with the AutoUnmonitorPreviouslyDownloadedTracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetAutoUnmonitorPreviouslyDownloadedTracksOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUnmonitorPreviouslyDownloadedTracks) {
		return nil, false
	}
	return o.AutoUnmonitorPreviouslyDownloadedTracks, true
}

// HasAutoUnmonitorPreviouslyDownloadedTracks returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasAutoUnmonitorPreviouslyDownloadedTracks() bool {
	if o != nil && !IsNil(o.AutoUnmonitorPreviouslyDownloadedTracks) {
		return true
	}

	return false
}

// SetAutoUnmonitorPreviouslyDownloadedTracks gets a reference to the given bool and assigns it to the AutoUnmonitorPreviouslyDownloadedTracks field.
func (o *MediaManagementConfigResource) SetAutoUnmonitorPreviouslyDownloadedTracks(v bool) {
	o.AutoUnmonitorPreviouslyDownloadedTracks = &v
}

// GetRecycleBin returns the RecycleBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetRecycleBin() string {
	if o == nil || IsNil(o.RecycleBin.Get()) {
		var ret string
		return ret
	}
	return *o.RecycleBin.Get()
}

// GetRecycleBinOk returns a tuple with the RecycleBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetRecycleBinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecycleBin.Get(), o.RecycleBin.IsSet()
}

// HasRecycleBin returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasRecycleBin() bool {
	if o != nil && o.RecycleBin.IsSet() {
		return true
	}

	return false
}

// SetRecycleBin gets a reference to the given NullableString and assigns it to the RecycleBin field.
func (o *MediaManagementConfigResource) SetRecycleBin(v string) {
	o.RecycleBin.Set(&v)
}
// SetRecycleBinNil sets the value for RecycleBin to be an explicit nil
func (o *MediaManagementConfigResource) SetRecycleBinNil() {
	o.RecycleBin.Set(nil)
}

// UnsetRecycleBin ensures that no value is present for RecycleBin, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetRecycleBin() {
	o.RecycleBin.Unset()
}

// GetRecycleBinCleanupDays returns the RecycleBinCleanupDays field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetRecycleBinCleanupDays() int32 {
	if o == nil || IsNil(o.RecycleBinCleanupDays) {
		var ret int32
		return ret
	}
	return *o.RecycleBinCleanupDays
}

// GetRecycleBinCleanupDaysOk returns a tuple with the RecycleBinCleanupDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetRecycleBinCleanupDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.RecycleBinCleanupDays) {
		return nil, false
	}
	return o.RecycleBinCleanupDays, true
}

// HasRecycleBinCleanupDays returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasRecycleBinCleanupDays() bool {
	if o != nil && !IsNil(o.RecycleBinCleanupDays) {
		return true
	}

	return false
}

// SetRecycleBinCleanupDays gets a reference to the given int32 and assigns it to the RecycleBinCleanupDays field.
func (o *MediaManagementConfigResource) SetRecycleBinCleanupDays(v int32) {
	o.RecycleBinCleanupDays = &v
}

// GetDownloadPropersAndRepacks returns the DownloadPropersAndRepacks field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetDownloadPropersAndRepacks() ProperDownloadTypes {
	if o == nil || IsNil(o.DownloadPropersAndRepacks) {
		var ret ProperDownloadTypes
		return ret
	}
	return *o.DownloadPropersAndRepacks
}

// GetDownloadPropersAndRepacksOk returns a tuple with the DownloadPropersAndRepacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetDownloadPropersAndRepacksOk() (*ProperDownloadTypes, bool) {
	if o == nil || IsNil(o.DownloadPropersAndRepacks) {
		return nil, false
	}
	return o.DownloadPropersAndRepacks, true
}

// HasDownloadPropersAndRepacks returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasDownloadPropersAndRepacks() bool {
	if o != nil && !IsNil(o.DownloadPropersAndRepacks) {
		return true
	}

	return false
}

// SetDownloadPropersAndRepacks gets a reference to the given ProperDownloadTypes and assigns it to the DownloadPropersAndRepacks field.
func (o *MediaManagementConfigResource) SetDownloadPropersAndRepacks(v ProperDownloadTypes) {
	o.DownloadPropersAndRepacks = &v
}

// GetCreateEmptyArtistFolders returns the CreateEmptyArtistFolders field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetCreateEmptyArtistFolders() bool {
	if o == nil || IsNil(o.CreateEmptyArtistFolders) {
		var ret bool
		return ret
	}
	return *o.CreateEmptyArtistFolders
}

// GetCreateEmptyArtistFoldersOk returns a tuple with the CreateEmptyArtistFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetCreateEmptyArtistFoldersOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateEmptyArtistFolders) {
		return nil, false
	}
	return o.CreateEmptyArtistFolders, true
}

// HasCreateEmptyArtistFolders returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasCreateEmptyArtistFolders() bool {
	if o != nil && !IsNil(o.CreateEmptyArtistFolders) {
		return true
	}

	return false
}

// SetCreateEmptyArtistFolders gets a reference to the given bool and assigns it to the CreateEmptyArtistFolders field.
func (o *MediaManagementConfigResource) SetCreateEmptyArtistFolders(v bool) {
	o.CreateEmptyArtistFolders = &v
}

// GetDeleteEmptyFolders returns the DeleteEmptyFolders field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetDeleteEmptyFolders() bool {
	if o == nil || IsNil(o.DeleteEmptyFolders) {
		var ret bool
		return ret
	}
	return *o.DeleteEmptyFolders
}

// GetDeleteEmptyFoldersOk returns a tuple with the DeleteEmptyFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetDeleteEmptyFoldersOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteEmptyFolders) {
		return nil, false
	}
	return o.DeleteEmptyFolders, true
}

// HasDeleteEmptyFolders returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasDeleteEmptyFolders() bool {
	if o != nil && !IsNil(o.DeleteEmptyFolders) {
		return true
	}

	return false
}

// SetDeleteEmptyFolders gets a reference to the given bool and assigns it to the DeleteEmptyFolders field.
func (o *MediaManagementConfigResource) SetDeleteEmptyFolders(v bool) {
	o.DeleteEmptyFolders = &v
}

// GetFileDate returns the FileDate field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetFileDate() FileDateType {
	if o == nil || IsNil(o.FileDate) {
		var ret FileDateType
		return ret
	}
	return *o.FileDate
}

// GetFileDateOk returns a tuple with the FileDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetFileDateOk() (*FileDateType, bool) {
	if o == nil || IsNil(o.FileDate) {
		return nil, false
	}
	return o.FileDate, true
}

// HasFileDate returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasFileDate() bool {
	if o != nil && !IsNil(o.FileDate) {
		return true
	}

	return false
}

// SetFileDate gets a reference to the given FileDateType and assigns it to the FileDate field.
func (o *MediaManagementConfigResource) SetFileDate(v FileDateType) {
	o.FileDate = &v
}

// GetWatchLibraryForChanges returns the WatchLibraryForChanges field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetWatchLibraryForChanges() bool {
	if o == nil || IsNil(o.WatchLibraryForChanges) {
		var ret bool
		return ret
	}
	return *o.WatchLibraryForChanges
}

// GetWatchLibraryForChangesOk returns a tuple with the WatchLibraryForChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetWatchLibraryForChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.WatchLibraryForChanges) {
		return nil, false
	}
	return o.WatchLibraryForChanges, true
}

// HasWatchLibraryForChanges returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasWatchLibraryForChanges() bool {
	if o != nil && !IsNil(o.WatchLibraryForChanges) {
		return true
	}

	return false
}

// SetWatchLibraryForChanges gets a reference to the given bool and assigns it to the WatchLibraryForChanges field.
func (o *MediaManagementConfigResource) SetWatchLibraryForChanges(v bool) {
	o.WatchLibraryForChanges = &v
}

// GetRescanAfterRefresh returns the RescanAfterRefresh field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetRescanAfterRefresh() RescanAfterRefreshType {
	if o == nil || IsNil(o.RescanAfterRefresh) {
		var ret RescanAfterRefreshType
		return ret
	}
	return *o.RescanAfterRefresh
}

// GetRescanAfterRefreshOk returns a tuple with the RescanAfterRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetRescanAfterRefreshOk() (*RescanAfterRefreshType, bool) {
	if o == nil || IsNil(o.RescanAfterRefresh) {
		return nil, false
	}
	return o.RescanAfterRefresh, true
}

// HasRescanAfterRefresh returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasRescanAfterRefresh() bool {
	if o != nil && !IsNil(o.RescanAfterRefresh) {
		return true
	}

	return false
}

// SetRescanAfterRefresh gets a reference to the given RescanAfterRefreshType and assigns it to the RescanAfterRefresh field.
func (o *MediaManagementConfigResource) SetRescanAfterRefresh(v RescanAfterRefreshType) {
	o.RescanAfterRefresh = &v
}

// GetAllowFingerprinting returns the AllowFingerprinting field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetAllowFingerprinting() AllowFingerprinting {
	if o == nil || IsNil(o.AllowFingerprinting) {
		var ret AllowFingerprinting
		return ret
	}
	return *o.AllowFingerprinting
}

// GetAllowFingerprintingOk returns a tuple with the AllowFingerprinting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetAllowFingerprintingOk() (*AllowFingerprinting, bool) {
	if o == nil || IsNil(o.AllowFingerprinting) {
		return nil, false
	}
	return o.AllowFingerprinting, true
}

// HasAllowFingerprinting returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasAllowFingerprinting() bool {
	if o != nil && !IsNil(o.AllowFingerprinting) {
		return true
	}

	return false
}

// SetAllowFingerprinting gets a reference to the given AllowFingerprinting and assigns it to the AllowFingerprinting field.
func (o *MediaManagementConfigResource) SetAllowFingerprinting(v AllowFingerprinting) {
	o.AllowFingerprinting = &v
}

// GetSetPermissionsLinux returns the SetPermissionsLinux field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetSetPermissionsLinux() bool {
	if o == nil || IsNil(o.SetPermissionsLinux) {
		var ret bool
		return ret
	}
	return *o.SetPermissionsLinux
}

// GetSetPermissionsLinuxOk returns a tuple with the SetPermissionsLinux field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetSetPermissionsLinuxOk() (*bool, bool) {
	if o == nil || IsNil(o.SetPermissionsLinux) {
		return nil, false
	}
	return o.SetPermissionsLinux, true
}

// HasSetPermissionsLinux returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasSetPermissionsLinux() bool {
	if o != nil && !IsNil(o.SetPermissionsLinux) {
		return true
	}

	return false
}

// SetSetPermissionsLinux gets a reference to the given bool and assigns it to the SetPermissionsLinux field.
func (o *MediaManagementConfigResource) SetSetPermissionsLinux(v bool) {
	o.SetPermissionsLinux = &v
}

// GetChmodFolder returns the ChmodFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetChmodFolder() string {
	if o == nil || IsNil(o.ChmodFolder.Get()) {
		var ret string
		return ret
	}
	return *o.ChmodFolder.Get()
}

// GetChmodFolderOk returns a tuple with the ChmodFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetChmodFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChmodFolder.Get(), o.ChmodFolder.IsSet()
}

// HasChmodFolder returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasChmodFolder() bool {
	if o != nil && o.ChmodFolder.IsSet() {
		return true
	}

	return false
}

// SetChmodFolder gets a reference to the given NullableString and assigns it to the ChmodFolder field.
func (o *MediaManagementConfigResource) SetChmodFolder(v string) {
	o.ChmodFolder.Set(&v)
}
// SetChmodFolderNil sets the value for ChmodFolder to be an explicit nil
func (o *MediaManagementConfigResource) SetChmodFolderNil() {
	o.ChmodFolder.Set(nil)
}

// UnsetChmodFolder ensures that no value is present for ChmodFolder, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetChmodFolder() {
	o.ChmodFolder.Unset()
}

// GetChownGroup returns the ChownGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetChownGroup() string {
	if o == nil || IsNil(o.ChownGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ChownGroup.Get()
}

// GetChownGroupOk returns a tuple with the ChownGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetChownGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChownGroup.Get(), o.ChownGroup.IsSet()
}

// HasChownGroup returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasChownGroup() bool {
	if o != nil && o.ChownGroup.IsSet() {
		return true
	}

	return false
}

// SetChownGroup gets a reference to the given NullableString and assigns it to the ChownGroup field.
func (o *MediaManagementConfigResource) SetChownGroup(v string) {
	o.ChownGroup.Set(&v)
}
// SetChownGroupNil sets the value for ChownGroup to be an explicit nil
func (o *MediaManagementConfigResource) SetChownGroupNil() {
	o.ChownGroup.Set(nil)
}

// UnsetChownGroup ensures that no value is present for ChownGroup, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetChownGroup() {
	o.ChownGroup.Unset()
}

// GetSkipFreeSpaceCheckWhenImporting returns the SkipFreeSpaceCheckWhenImporting field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetSkipFreeSpaceCheckWhenImporting() bool {
	if o == nil || IsNil(o.SkipFreeSpaceCheckWhenImporting) {
		var ret bool
		return ret
	}
	return *o.SkipFreeSpaceCheckWhenImporting
}

// GetSkipFreeSpaceCheckWhenImportingOk returns a tuple with the SkipFreeSpaceCheckWhenImporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetSkipFreeSpaceCheckWhenImportingOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipFreeSpaceCheckWhenImporting) {
		return nil, false
	}
	return o.SkipFreeSpaceCheckWhenImporting, true
}

// HasSkipFreeSpaceCheckWhenImporting returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasSkipFreeSpaceCheckWhenImporting() bool {
	if o != nil && !IsNil(o.SkipFreeSpaceCheckWhenImporting) {
		return true
	}

	return false
}

// SetSkipFreeSpaceCheckWhenImporting gets a reference to the given bool and assigns it to the SkipFreeSpaceCheckWhenImporting field.
func (o *MediaManagementConfigResource) SetSkipFreeSpaceCheckWhenImporting(v bool) {
	o.SkipFreeSpaceCheckWhenImporting = &v
}

// GetMinimumFreeSpaceWhenImporting returns the MinimumFreeSpaceWhenImporting field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetMinimumFreeSpaceWhenImporting() int32 {
	if o == nil || IsNil(o.MinimumFreeSpaceWhenImporting) {
		var ret int32
		return ret
	}
	return *o.MinimumFreeSpaceWhenImporting
}

// GetMinimumFreeSpaceWhenImportingOk returns a tuple with the MinimumFreeSpaceWhenImporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetMinimumFreeSpaceWhenImportingOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumFreeSpaceWhenImporting) {
		return nil, false
	}
	return o.MinimumFreeSpaceWhenImporting, true
}

// HasMinimumFreeSpaceWhenImporting returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasMinimumFreeSpaceWhenImporting() bool {
	if o != nil && !IsNil(o.MinimumFreeSpaceWhenImporting) {
		return true
	}

	return false
}

// SetMinimumFreeSpaceWhenImporting gets a reference to the given int32 and assigns it to the MinimumFreeSpaceWhenImporting field.
func (o *MediaManagementConfigResource) SetMinimumFreeSpaceWhenImporting(v int32) {
	o.MinimumFreeSpaceWhenImporting = &v
}

// GetCopyUsingHardlinks returns the CopyUsingHardlinks field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetCopyUsingHardlinks() bool {
	if o == nil || IsNil(o.CopyUsingHardlinks) {
		var ret bool
		return ret
	}
	return *o.CopyUsingHardlinks
}

// GetCopyUsingHardlinksOk returns a tuple with the CopyUsingHardlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetCopyUsingHardlinksOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyUsingHardlinks) {
		return nil, false
	}
	return o.CopyUsingHardlinks, true
}

// HasCopyUsingHardlinks returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasCopyUsingHardlinks() bool {
	if o != nil && !IsNil(o.CopyUsingHardlinks) {
		return true
	}

	return false
}

// SetCopyUsingHardlinks gets a reference to the given bool and assigns it to the CopyUsingHardlinks field.
func (o *MediaManagementConfigResource) SetCopyUsingHardlinks(v bool) {
	o.CopyUsingHardlinks = &v
}

// GetImportExtraFiles returns the ImportExtraFiles field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetImportExtraFiles() bool {
	if o == nil || IsNil(o.ImportExtraFiles) {
		var ret bool
		return ret
	}
	return *o.ImportExtraFiles
}

// GetImportExtraFilesOk returns a tuple with the ImportExtraFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetImportExtraFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportExtraFiles) {
		return nil, false
	}
	return o.ImportExtraFiles, true
}

// HasImportExtraFiles returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasImportExtraFiles() bool {
	if o != nil && !IsNil(o.ImportExtraFiles) {
		return true
	}

	return false
}

// SetImportExtraFiles gets a reference to the given bool and assigns it to the ImportExtraFiles field.
func (o *MediaManagementConfigResource) SetImportExtraFiles(v bool) {
	o.ImportExtraFiles = &v
}

// GetExtraFileExtensions returns the ExtraFileExtensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetExtraFileExtensions() string {
	if o == nil || IsNil(o.ExtraFileExtensions.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraFileExtensions.Get()
}

// GetExtraFileExtensionsOk returns a tuple with the ExtraFileExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetExtraFileExtensionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraFileExtensions.Get(), o.ExtraFileExtensions.IsSet()
}

// HasExtraFileExtensions returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasExtraFileExtensions() bool {
	if o != nil && o.ExtraFileExtensions.IsSet() {
		return true
	}

	return false
}

// SetExtraFileExtensions gets a reference to the given NullableString and assigns it to the ExtraFileExtensions field.
func (o *MediaManagementConfigResource) SetExtraFileExtensions(v string) {
	o.ExtraFileExtensions.Set(&v)
}
// SetExtraFileExtensionsNil sets the value for ExtraFileExtensions to be an explicit nil
func (o *MediaManagementConfigResource) SetExtraFileExtensionsNil() {
	o.ExtraFileExtensions.Set(nil)
}

// UnsetExtraFileExtensions ensures that no value is present for ExtraFileExtensions, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetExtraFileExtensions() {
	o.ExtraFileExtensions.Unset()
}

func (o MediaManagementConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaManagementConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AutoUnmonitorPreviouslyDownloadedTracks) {
		toSerialize["autoUnmonitorPreviouslyDownloadedTracks"] = o.AutoUnmonitorPreviouslyDownloadedTracks
	}
	if o.RecycleBin.IsSet() {
		toSerialize["recycleBin"] = o.RecycleBin.Get()
	}
	if !IsNil(o.RecycleBinCleanupDays) {
		toSerialize["recycleBinCleanupDays"] = o.RecycleBinCleanupDays
	}
	if !IsNil(o.DownloadPropersAndRepacks) {
		toSerialize["downloadPropersAndRepacks"] = o.DownloadPropersAndRepacks
	}
	if !IsNil(o.CreateEmptyArtistFolders) {
		toSerialize["createEmptyArtistFolders"] = o.CreateEmptyArtistFolders
	}
	if !IsNil(o.DeleteEmptyFolders) {
		toSerialize["deleteEmptyFolders"] = o.DeleteEmptyFolders
	}
	if !IsNil(o.FileDate) {
		toSerialize["fileDate"] = o.FileDate
	}
	if !IsNil(o.WatchLibraryForChanges) {
		toSerialize["watchLibraryForChanges"] = o.WatchLibraryForChanges
	}
	if !IsNil(o.RescanAfterRefresh) {
		toSerialize["rescanAfterRefresh"] = o.RescanAfterRefresh
	}
	if !IsNil(o.AllowFingerprinting) {
		toSerialize["allowFingerprinting"] = o.AllowFingerprinting
	}
	if !IsNil(o.SetPermissionsLinux) {
		toSerialize["setPermissionsLinux"] = o.SetPermissionsLinux
	}
	if o.ChmodFolder.IsSet() {
		toSerialize["chmodFolder"] = o.ChmodFolder.Get()
	}
	if o.ChownGroup.IsSet() {
		toSerialize["chownGroup"] = o.ChownGroup.Get()
	}
	if !IsNil(o.SkipFreeSpaceCheckWhenImporting) {
		toSerialize["skipFreeSpaceCheckWhenImporting"] = o.SkipFreeSpaceCheckWhenImporting
	}
	if !IsNil(o.MinimumFreeSpaceWhenImporting) {
		toSerialize["minimumFreeSpaceWhenImporting"] = o.MinimumFreeSpaceWhenImporting
	}
	if !IsNil(o.CopyUsingHardlinks) {
		toSerialize["copyUsingHardlinks"] = o.CopyUsingHardlinks
	}
	if !IsNil(o.ImportExtraFiles) {
		toSerialize["importExtraFiles"] = o.ImportExtraFiles
	}
	if o.ExtraFileExtensions.IsSet() {
		toSerialize["extraFileExtensions"] = o.ExtraFileExtensions.Get()
	}
	return toSerialize, nil
}

type NullableMediaManagementConfigResource struct {
	value *MediaManagementConfigResource
	isSet bool
}

func (v NullableMediaManagementConfigResource) Get() *MediaManagementConfigResource {
	return v.value
}

func (v *NullableMediaManagementConfigResource) Set(val *MediaManagementConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaManagementConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaManagementConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaManagementConfigResource(val *MediaManagementConfigResource) *NullableMediaManagementConfigResource {
	return &NullableMediaManagementConfigResource{value: val, isSet: true}
}

func (v NullableMediaManagementConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaManagementConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


