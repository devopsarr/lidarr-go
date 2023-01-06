# \AlbumApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Album**](AlbumApi.md#CreateApiV1Album) | **Post** /api/v1/album | 
[**DeleteApiV1Album**](AlbumApi.md#DeleteApiV1Album) | **Delete** /api/v1/album/{id} | 
[**GetApiV1AlbumById**](AlbumApi.md#GetApiV1AlbumById) | **Get** /api/v1/album/{id} | 
[**ListApiV1Album**](AlbumApi.md#ListApiV1Album) | **Get** /api/v1/album | 
[**PutApiV1AlbumMonitor**](AlbumApi.md#PutApiV1AlbumMonitor) | **Put** /api/v1/album/monitor | 
[**UpdateApiV1Album**](AlbumApi.md#UpdateApiV1Album) | **Put** /api/v1/album/{id} | 



## CreateApiV1Album

> AlbumResource CreateApiV1Album(ctx).AlbumResource(albumResource).Execute()



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
    albumResource := *lidarrClient.NewAlbumResource() // AlbumResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.CreateApiV1Album(context.Background()).AlbumResource(albumResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.CreateApiV1Album``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Album`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.CreateApiV1Album`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1AlbumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumResource** | [**AlbumResource**](AlbumResource.md) |  | 

### Return type

[**AlbumResource**](AlbumResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1Album

> DeleteApiV1Album(ctx, id).Execute()



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
    resp, r, err := apiClient.AlbumApi.DeleteApiV1Album(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.DeleteApiV1Album``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1AlbumRequest struct via the builder pattern


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


## GetApiV1AlbumById

> AlbumResource GetApiV1AlbumById(ctx, id).Execute()



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
    resp, r, err := apiClient.AlbumApi.GetApiV1AlbumById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.GetApiV1AlbumById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1AlbumById`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.GetApiV1AlbumById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AlbumByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlbumResource**](AlbumResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Album

> []AlbumResource ListApiV1Album(ctx).ArtistId(artistId).AlbumIds(albumIds).ForeignAlbumId(foreignAlbumId).IncludeAllArtistAlbums(includeAllArtistAlbums).Execute()



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
    albumIds := []int32{int32(123)} // []int32 |  (optional)
    foreignAlbumId := "foreignAlbumId_example" // string |  (optional)
    includeAllArtistAlbums := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.ListApiV1Album(context.Background()).ArtistId(artistId).AlbumIds(albumIds).ForeignAlbumId(foreignAlbumId).IncludeAllArtistAlbums(includeAllArtistAlbums).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.ListApiV1Album``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Album`: []AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.ListApiV1Album`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1AlbumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumIds** | **[]int32** |  | 
 **foreignAlbumId** | **string** |  | 
 **includeAllArtistAlbums** | **bool** |  | [default to false]

### Return type

[**[]AlbumResource**](AlbumResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1AlbumMonitor

> PutApiV1AlbumMonitor(ctx).AlbumsMonitoredResource(albumsMonitoredResource).Execute()



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
    albumsMonitoredResource := *lidarrClient.NewAlbumsMonitoredResource() // AlbumsMonitoredResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.PutApiV1AlbumMonitor(context.Background()).AlbumsMonitoredResource(albumsMonitoredResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.PutApiV1AlbumMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1AlbumMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumsMonitoredResource** | [**AlbumsMonitoredResource**](AlbumsMonitoredResource.md) |  | 

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


## UpdateApiV1Album

> AlbumResource UpdateApiV1Album(ctx, id).AlbumResource(albumResource).Execute()



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
    albumResource := *lidarrClient.NewAlbumResource() // AlbumResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumApi.UpdateApiV1Album(context.Background(), id).AlbumResource(albumResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.UpdateApiV1Album``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Album`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.UpdateApiV1Album`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1AlbumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **albumResource** | [**AlbumResource**](AlbumResource.md) |  | 

### Return type

[**AlbumResource**](AlbumResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

