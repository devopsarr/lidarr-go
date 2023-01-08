# \TrackFileApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1TrackFile**](TrackFileApi.md#DeleteApiV1TrackFile) | **Delete** /api/v1/trackfile/{id} | 
[**DeleteApiV1TrackFileBulk**](TrackFileApi.md#DeleteApiV1TrackFileBulk) | **Delete** /api/v1/trackfile/bulk | 
[**GetApiV1TrackFileById**](TrackFileApi.md#GetApiV1TrackFileById) | **Get** /api/v1/trackfile/{id} | 
[**ListApiV1TrackFile**](TrackFileApi.md#ListApiV1TrackFile) | **Get** /api/v1/trackfile | 
[**PutApiV1TrackFileEditor**](TrackFileApi.md#PutApiV1TrackFileEditor) | **Put** /api/v1/trackfile/editor | 
[**UpdateApiV1TrackFile**](TrackFileApi.md#UpdateApiV1TrackFile) | **Put** /api/v1/trackfile/{id} | 



## DeleteApiV1TrackFile

> DeleteApiV1TrackFile(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackFileApi.DeleteApiV1TrackFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.DeleteApiV1TrackFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1TrackFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1TrackFileBulk

> DeleteApiV1TrackFileBulk(ctx).TrackFileListResource(trackFileListResource).Execute()



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
    trackFileListResource := *lidarrClient.NewTrackFileListResource() // TrackFileListResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackFileApi.DeleteApiV1TrackFileBulk(context.Background()).TrackFileListResource(trackFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.DeleteApiV1TrackFileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1TrackFileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackFileListResource** | [**TrackFileListResource**](TrackFileListResource.md) |  | 

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


## GetApiV1TrackFileById

> TrackFileResource GetApiV1TrackFileById(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackFileApi.GetApiV1TrackFileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.GetApiV1TrackFileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1TrackFileById`: TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.GetApiV1TrackFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1TrackFileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrackFileResource**](TrackFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1TrackFile

> []TrackFileResource ListApiV1TrackFile(ctx).ArtistId(artistId).TrackFileIds(trackFileIds).AlbumId(albumId).Unmapped(unmapped).Execute()



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
    artistId := int32(56) // int32 |  (optional)
    trackFileIds := []int32{int32(123)} // []int32 |  (optional)
    albumId := []int32{int32(123)} // []int32 |  (optional)
    unmapped := true // bool |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackFileApi.ListApiV1TrackFile(context.Background()).ArtistId(artistId).TrackFileIds(trackFileIds).AlbumId(albumId).Unmapped(unmapped).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.ListApiV1TrackFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1TrackFile`: []TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.ListApiV1TrackFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1TrackFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **trackFileIds** | **[]int32** |  | 
 **albumId** | **[]int32** |  | 
 **unmapped** | **bool** |  | 

### Return type

[**[]TrackFileResource**](TrackFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1TrackFileEditor

> PutApiV1TrackFileEditor(ctx).TrackFileListResource(trackFileListResource).Execute()



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
    trackFileListResource := *lidarrClient.NewTrackFileListResource() // TrackFileListResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackFileApi.PutApiV1TrackFileEditor(context.Background()).TrackFileListResource(trackFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.PutApiV1TrackFileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1TrackFileEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackFileListResource** | [**TrackFileListResource**](TrackFileListResource.md) |  | 

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


## UpdateApiV1TrackFile

> TrackFileResource UpdateApiV1TrackFile(ctx, id).TrackFileResource(trackFileResource).Execute()



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
    id := "id_example" // string | 
    trackFileResource := *lidarrClient.NewTrackFileResource() // TrackFileResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackFileApi.UpdateApiV1TrackFile(context.Background(), id).TrackFileResource(trackFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.UpdateApiV1TrackFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1TrackFile`: TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.UpdateApiV1TrackFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1TrackFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackFileResource** | [**TrackFileResource**](TrackFileResource.md) |  | 

### Return type

[**TrackFileResource**](TrackFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

