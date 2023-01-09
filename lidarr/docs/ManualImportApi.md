# \ManualImportApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListManualImport**](ManualImportApi.md#ListManualImport) | **Get** /api/v1/manualimport | 
[**PutManualImport**](ManualImportApi.md#PutManualImport) | **Put** /api/v1/manualimport | 



## ListManualImport

> []ManualImportResource ListManualImport(ctx).Folder(folder).DownloadId(downloadId).ArtistId(artistId).FilterExistingFiles(filterExistingFiles).ReplaceExistingFiles(replaceExistingFiles).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    lidarrClient "./openapi"
)

func main() {
    folder := "folder_example" // string |  (optional)
    downloadId := "downloadId_example" // string |  (optional)
    artistId := int32(56) // int32 |  (optional)
    filterExistingFiles := true // bool |  (optional) (default to true)
    replaceExistingFiles := true // bool |  (optional) (default to true)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualImportApi.ListManualImport(context.Background()).Folder(folder).DownloadId(downloadId).ArtistId(artistId).FilterExistingFiles(filterExistingFiles).ReplaceExistingFiles(replaceExistingFiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.ListManualImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManualImport`: []ManualImportResource
    fmt.Fprintf(os.Stdout, "Response from `ManualImportApi.ListManualImport`: %v\n", resp)
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

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManualImport

> PutManualImport(ctx).ManualImportResource(manualImportResource).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    lidarrClient "./openapi"
)

func main() {
    manualImportResource := []lidarrClient.ManualImportResource{*lidarrClient.NewManualImportResource()} // []ManualImportResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualImportApi.PutManualImport(context.Background()).ManualImportResource(manualImportResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.PutManualImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualImportResource** | [**[]ManualImportResource**](ManualImportResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

