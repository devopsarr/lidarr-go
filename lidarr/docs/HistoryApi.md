# \HistoryApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryFailedById**](HistoryApi.md#CreateHistoryFailedById) | **Post** /api/v1/history/failed/{id} | 
[**GetHistory**](HistoryApi.md#GetHistory) | **Get** /api/v1/history | 
[**ListHistoryArtist**](HistoryApi.md#ListHistoryArtist) | **Get** /api/v1/history/artist | 
[**ListHistorySince**](HistoryApi.md#ListHistorySince) | **Get** /api/v1/history/since | 



## CreateHistoryFailedById

> CreateHistoryFailedById(ctx, id).Execute()



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
    resp, r, err := apiClient.HistoryApi.CreateHistoryFailedById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.CreateHistoryFailedById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateHistoryFailedByIdRequest struct via the builder pattern


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


## GetHistory

> HistoryResourcePagingResource GetHistory(ctx).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



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
    includeArtist := true // bool |  (optional) (default to false)
    includeAlbum := true // bool |  (optional) (default to false)
    includeTrack := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.GetHistory(context.Background()).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to false]
 **includeTrack** | **bool** |  | [default to false]

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryArtist

> []HistoryResource ListHistoryArtist(ctx).ArtistId(artistId).AlbumId(albumId).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



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
    albumId := int32(56) // int32 |  (optional)
    eventType := lidarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
    includeArtist := true // bool |  (optional) (default to false)
    includeAlbum := true // bool |  (optional) (default to false)
    includeTrack := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ListHistoryArtist(context.Background()).ArtistId(artistId).AlbumId(albumId).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistoryArtist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryArtist`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistoryArtist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumId** | **int32** |  | 
 **eventType** | [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | 
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to false]
 **includeTrack** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorySince

> []HistoryResource ListHistorySince(ctx).Date(date).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    lidarrClient "./openapi"
)

func main() {
    date := time.Now() // time.Time |  (optional)
    eventType := lidarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
    includeArtist := true // bool |  (optional) (default to false)
    includeAlbum := true // bool |  (optional) (default to false)
    includeTrack := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ListHistorySince(context.Background()).Date(date).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistorySinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | 
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to false]
 **includeTrack** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

