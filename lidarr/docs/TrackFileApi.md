# \TrackFileApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1Trackfile**](TrackFileApi.md#DeleteApiV1Trackfile) | **Delete** /api/v1/trackfile/{id} | 
[**DeleteApiV1TrackfileBulk**](TrackFileApi.md#DeleteApiV1TrackfileBulk) | **Delete** /api/v1/trackfile/bulk | 
[**GetApiV1TrackfileById**](TrackFileApi.md#GetApiV1TrackfileById) | **Get** /api/v1/trackfile/{id} | 
[**ListApiV1Trackfile**](TrackFileApi.md#ListApiV1Trackfile) | **Get** /api/v1/trackfile | 
[**PutApiV1TrackfileEditor**](TrackFileApi.md#PutApiV1TrackfileEditor) | **Put** /api/v1/trackfile/editor | 
[**UpdateApiV1Trackfile**](TrackFileApi.md#UpdateApiV1Trackfile) | **Put** /api/v1/trackfile/{id} | 



## DeleteApiV1Trackfile

> DeleteApiV1Trackfile(ctx, id).Execute()



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
    resp, r, err := apiClient.TrackFileApi.DeleteApiV1Trackfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.DeleteApiV1Trackfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1TrackfileRequest struct via the builder pattern


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


## DeleteApiV1TrackfileBulk

> DeleteApiV1TrackfileBulk(ctx).TrackFileListResource(trackFileListResource).Execute()



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
    resp, r, err := apiClient.TrackFileApi.DeleteApiV1TrackfileBulk(context.Background()).TrackFileListResource(trackFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.DeleteApiV1TrackfileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1TrackfileBulkRequest struct via the builder pattern


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


## GetApiV1TrackfileById

> TrackFileResource GetApiV1TrackfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.TrackFileApi.GetApiV1TrackfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.GetApiV1TrackfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1TrackfileById`: TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.GetApiV1TrackfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1TrackfileByIdRequest struct via the builder pattern


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


## ListApiV1Trackfile

> []TrackFileResource ListApiV1Trackfile(ctx).ArtistId(artistId).TrackFileIds(trackFileIds).AlbumId(albumId).Unmapped(unmapped).Execute()



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
    resp, r, err := apiClient.TrackFileApi.ListApiV1Trackfile(context.Background()).ArtistId(artistId).TrackFileIds(trackFileIds).AlbumId(albumId).Unmapped(unmapped).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.ListApiV1Trackfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Trackfile`: []TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.ListApiV1Trackfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1TrackfileRequest struct via the builder pattern


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


## PutApiV1TrackfileEditor

> PutApiV1TrackfileEditor(ctx).TrackFileListResource(trackFileListResource).Execute()



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
    resp, r, err := apiClient.TrackFileApi.PutApiV1TrackfileEditor(context.Background()).TrackFileListResource(trackFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.PutApiV1TrackfileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1TrackfileEditorRequest struct via the builder pattern


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


## UpdateApiV1Trackfile

> TrackFileResource UpdateApiV1Trackfile(ctx, id).TrackFileResource(trackFileResource).Execute()



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
    resp, r, err := apiClient.TrackFileApi.UpdateApiV1Trackfile(context.Background(), id).TrackFileResource(trackFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackFileApi.UpdateApiV1Trackfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Trackfile`: TrackFileResource
    fmt.Fprintf(os.Stdout, "Response from `TrackFileApi.UpdateApiV1Trackfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1TrackfileRequest struct via the builder pattern


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

