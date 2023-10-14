# \AlbumAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlbum**](AlbumAPI.md#CreateAlbum) | **Post** /api/v1/album | 
[**DeleteAlbum**](AlbumAPI.md#DeleteAlbum) | **Delete** /api/v1/album/{id} | 
[**GetAlbumById**](AlbumAPI.md#GetAlbumById) | **Get** /api/v1/album/{id} | 
[**ListAlbum**](AlbumAPI.md#ListAlbum) | **Get** /api/v1/album | 
[**PutAlbumMonitor**](AlbumAPI.md#PutAlbumMonitor) | **Put** /api/v1/album/monitor | 
[**UpdateAlbum**](AlbumAPI.md#UpdateAlbum) | **Put** /api/v1/album/{id} | 



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
    resp, r, err := apiClient.AlbumAPI.CreateAlbum(context.Background()).AlbumResource(albumResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumAPI.CreateAlbum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlbum`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumAPI.CreateAlbum`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlbum

> DeleteAlbum(ctx, id).DeleteFiles(deleteFiles).AddImportListExclusion(addImportListExclusion).Execute()



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
    deleteFiles := true // bool |  (optional) (default to false)
    addImportListExclusion := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumAPI.DeleteAlbum(context.Background(), id).DeleteFiles(deleteFiles).AddImportListExclusion(addImportListExclusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumAPI.DeleteAlbum``: %v\n", err)
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

 **deleteFiles** | **bool** |  | [default to false]
 **addImportListExclusion** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

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
    resp, r, err := apiClient.AlbumAPI.GetAlbumById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumAPI.GetAlbumById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlbumById`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumAPI.GetAlbumById`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

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
    resp, r, err := apiClient.AlbumAPI.ListAlbum(context.Background()).ArtistId(artistId).AlbumIds(albumIds).ForeignAlbumId(foreignAlbumId).IncludeAllArtistAlbums(includeAllArtistAlbums).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumAPI.ListAlbum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlbum`: []AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumAPI.ListAlbum`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

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
    resp, r, err := apiClient.AlbumAPI.PutAlbumMonitor(context.Background()).AlbumsMonitoredResource(albumsMonitoredResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumAPI.PutAlbumMonitor``: %v\n", err)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

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
    resp, r, err := apiClient.AlbumAPI.UpdateAlbum(context.Background(), id).AlbumResource(albumResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumAPI.UpdateAlbum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlbum`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `AlbumAPI.UpdateAlbum`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

