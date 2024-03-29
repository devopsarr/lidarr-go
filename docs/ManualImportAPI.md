# \ManualImportAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualImport**](ManualImportAPI.md#CreateManualImport) | **Post** /api/v1/manualimport | 
[**ListManualImport**](ManualImportAPI.md#ListManualImport) | **Get** /api/v1/manualimport | 



## CreateManualImport

> CreateManualImport(ctx).ManualImportUpdateResource(manualImportUpdateResource).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	lidarrClient "github.com/devopsarr/lidarr-go/lidarr"
)

func main() {
	manualImportUpdateResource := []lidarrClient.ManualImportUpdateResource{*lidarrClient.NewManualImportUpdateResource()} // []ManualImportUpdateResource |  (optional)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	r, err := apiClient.ManualImportAPI.CreateManualImport(context.Background()).ManualImportUpdateResource(manualImportUpdateResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualImportAPI.CreateManualImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualImportUpdateResource** | [**[]ManualImportUpdateResource**](ManualImportUpdateResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManualImport

> []ManualImportResource ListManualImport(ctx).Folder(folder).DownloadId(downloadId).ArtistId(artistId).FilterExistingFiles(filterExistingFiles).ReplaceExistingFiles(replaceExistingFiles).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	lidarrClient "github.com/devopsarr/lidarr-go/lidarr"
)

func main() {
	folder := "folder_example" // string |  (optional)
	downloadId := "downloadId_example" // string |  (optional)
	artistId := int32(56) // int32 |  (optional)
	filterExistingFiles := true // bool |  (optional) (default to true)
	replaceExistingFiles := true // bool |  (optional) (default to true)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManualImportAPI.ListManualImport(context.Background()).Folder(folder).DownloadId(downloadId).ArtistId(artistId).FilterExistingFiles(filterExistingFiles).ReplaceExistingFiles(replaceExistingFiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualImportAPI.ListManualImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListManualImport`: []ManualImportResource
	fmt.Fprintf(os.Stdout, "Response from `ManualImportAPI.ListManualImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folder** | **string** |  | 
 **downloadId** | **string** |  | 
 **artistId** | **int32** |  | 
 **filterExistingFiles** | **bool** |  | [default to true]
 **replaceExistingFiles** | **bool** |  | [default to true]

### Return type

[**[]ManualImportResource**](ManualImportResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

