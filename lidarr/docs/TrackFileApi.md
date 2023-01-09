# \TrackFileApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrackFile**](TrackFileApi.md#DeleteTrackFile) | **Delete** /api/v1/trackfile/{id} | 
[**DeleteTrackFileBulk**](TrackFileApi.md#DeleteTrackFileBulk) | **Delete** /api/v1/trackfile/bulk | 
[**GetTrackFileById**](TrackFileApi.md#GetTrackFileById) | **Get** /api/v1/trackfile/{id} | 
[**ListTrackFile**](TrackFileApi.md#ListTrackFile) | **Get** /api/v1/trackfile | 
[**PutTrackFileEditor**](TrackFileApi.md#PutTrackFileEditor) | **Put** /api/v1/trackfile/editor | 
[**UpdateTrackFile**](TrackFileApi.md#UpdateTrackFile) | **Put** /api/v1/trackfile/{id} | 



## DeleteTrackFile

> DeleteTrackFile(ctx, id).Execute()



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
    resp, r, err := apiClient.TrackFileApi.DeleteTrackFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.DeleteTrackFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTrackFileRequest struct via the builder pattern


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


## DeleteTrackFileBulk

> DeleteTrackFileBulk(ctx).TrackFileListResource(trackFileListResource).Execute()



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
    resp, r, err := apiClient.TrackFileApi.DeleteTrackFileBulk(context.Background()).TrackFileListResource(trackFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.DeleteTrackFileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackFileBulkRequest struct via the builder pattern


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


## GetTrackFileById

> TrackFileResource GetTrackFileById(ctx, id).Execute()



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
    resp, r, err := apiClient.TrackFileApi.GetTrackFileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.GetTrackFileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrackFileById`: TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.GetTrackFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackFileByIdRequest struct via the builder pattern


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


## ListTrackFile

> []TrackFileResource ListTrackFile(ctx).ArtistId(artistId).TrackFileIds(trackFileIds).AlbumId(albumId).Unmapped(unmapped).Execute()



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
    resp, r, err := apiClient.TrackFileApi.ListTrackFile(context.Background()).ArtistId(artistId).TrackFileIds(trackFileIds).AlbumId(albumId).Unmapped(unmapped).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.ListTrackFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrackFile`: []TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.ListTrackFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrackFileRequest struct via the builder pattern


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


## PutTrackFileEditor

> PutTrackFileEditor(ctx).TrackFileListResource(trackFileListResource).Execute()



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
    resp, r, err := apiClient.TrackFileApi.PutTrackFileEditor(context.Background()).TrackFileListResource(trackFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.PutTrackFileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTrackFileEditorRequest struct via the builder pattern


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


## UpdateTrackFile

> TrackFileResource UpdateTrackFile(ctx, id).TrackFileResource(trackFileResource).Execute()



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
    resp, r, err := apiClient.TrackFileApi.UpdateTrackFile(context.Background(), id).TrackFileResource(trackFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.UpdateTrackFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrackFile`: TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.UpdateTrackFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackFileRequest struct via the builder pattern


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

