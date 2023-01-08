# \ManualImportApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiV1ManualImport**](ManualImportApi.md#ListApiV1ManualImport) | **Get** /api/v1/manualimport | 
[**PutApiV1ManualImport**](ManualImportApi.md#PutApiV1ManualImport) | **Put** /api/v1/manualimport | 



## ListApiV1ManualImport

> []ManualImportResource ListApiV1ManualImport(ctx).Folder(folder).DownloadId(downloadId).ArtistId(artistId).FilterExistingFiles(filterExistingFiles).ReplaceExistingFiles(replaceExistingFiles).Execute()



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
    resp, r, err := apiClient.ManualImportApi.ListApiV1ManualImport(context.Background()).Folder(folder).DownloadId(downloadId).ArtistId(artistId).FilterExistingFiles(filterExistingFiles).ReplaceExistingFiles(replaceExistingFiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.ListApiV1ManualImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1ManualImport`: []ManualImportResource
    fmt.Fprintf(os.Stdout, "Response from `ManualImportApi.ListApiV1ManualImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1ManualImportRequest struct via the builder pattern


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


## PutApiV1ManualImport

> PutApiV1ManualImport(ctx).ManualImportResource(manualImportResource).Execute()



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
    resp, r, err := apiClient.ManualImportApi.PutApiV1ManualImport(context.Background()).ManualImportResource(manualImportResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.PutApiV1ManualImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1ManualImportRequest struct via the builder pattern


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

