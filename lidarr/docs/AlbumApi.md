# \AlbumApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlbum**](AlbumApi.md#CreateAlbum) | **Post** /api/v1/album | 
[**DeleteAlbum**](AlbumApi.md#DeleteAlbum) | **Delete** /api/v1/album/{id} | 
[**GetAlbumById**](AlbumApi.md#GetAlbumById) | **Get** /api/v1/album/{id} | 
[**ListAlbum**](AlbumApi.md#ListAlbum) | **Get** /api/v1/album | 
[**PutAlbumMonitor**](AlbumApi.md#PutAlbumMonitor) | **Put** /api/v1/album/monitor | 
[**UpdateAlbum**](AlbumApi.md#UpdateAlbum) | **Put** /api/v1/album/{id} | 



## CreateAlbum

> AlbumResource CreateAlbum(ctx).AlbumResource(albumResource).Execute()



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
    resp, r, err := apiClient.AlbumApi.CreateAlbum(context.Background()).AlbumResource(albumResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.CreateAlbum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlbum`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.CreateAlbum`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlbumRequest struct via the builder pattern


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


## DeleteAlbum

> DeleteAlbum(ctx, id).Execute()



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
    resp, r, err := apiClient.AlbumApi.DeleteAlbum(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.DeleteAlbum``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAlbumRequest struct via the builder pattern


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


## GetAlbumById

> AlbumResource GetAlbumById(ctx, id).Execute()



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
    resp, r, err := apiClient.AlbumApi.GetAlbumById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.GetAlbumById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlbumById`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.GetAlbumById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlbumByIdRequest struct via the builder pattern


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


## ListAlbum

> []AlbumResource ListAlbum(ctx).ArtistId(artistId).AlbumIds(albumIds).ForeignAlbumId(foreignAlbumId).IncludeAllArtistAlbums(includeAllArtistAlbums).Execute()



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
    resp, r, err := apiClient.AlbumApi.ListAlbum(context.Background()).ArtistId(artistId).AlbumIds(albumIds).ForeignAlbumId(foreignAlbumId).IncludeAllArtistAlbums(includeAllArtistAlbums).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.ListAlbum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlbum`: []AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.ListAlbum`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlbumRequest struct via the builder pattern


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


## PutAlbumMonitor

> PutAlbumMonitor(ctx).AlbumsMonitoredResource(albumsMonitoredResource).Execute()



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
    resp, r, err := apiClient.AlbumApi.PutAlbumMonitor(context.Background()).AlbumsMonitoredResource(albumsMonitoredResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.PutAlbumMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAlbumMonitorRequest struct via the builder pattern


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


## UpdateAlbum

> AlbumResource UpdateAlbum(ctx, id).AlbumResource(albumResource).Execute()



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
    resp, r, err := apiClient.AlbumApi.UpdateAlbum(context.Background(), id).AlbumResource(albumResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumApi.UpdateAlbum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlbum`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumApi.UpdateAlbum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlbumRequest struct via the builder pattern


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

