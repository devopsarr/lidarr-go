# \HistoryApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1HistoryFailedById**](HistoryApi.md#CreateApiV1HistoryFailedById) | **Post** /api/v1/history/failed/{id} | 
[**GetApiV1History**](HistoryApi.md#GetApiV1History) | **Get** /api/v1/history | 
[**ListApiV1HistoryArtist**](HistoryApi.md#ListApiV1HistoryArtist) | **Get** /api/v1/history/artist | 
[**ListApiV1HistorySince**](HistoryApi.md#ListApiV1HistorySince) | **Get** /api/v1/history/since | 



## CreateApiV1HistoryFailedById

> CreateApiV1HistoryFailedById(ctx, id).Execute()



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
    resp, r, err := apiClient.HistoryApi.CreateApiV1HistoryFailedById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.CreateApiV1HistoryFailedById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateApiV1HistoryFailedByIdRequest struct via the builder pattern


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


## GetApiV1History

> HistoryResourcePagingResource GetApiV1History(ctx).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



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
    resp, r, err := apiClient.HistoryApi.GetApiV1History(context.Background()).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetApiV1History``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1History`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetApiV1History`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1HistoryRequest struct via the builder pattern


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


## ListApiV1HistoryArtist

> []HistoryResource ListApiV1HistoryArtist(ctx).ArtistId(artistId).AlbumId(albumId).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



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
    resp, r, err := apiClient.HistoryApi.ListApiV1HistoryArtist(context.Background()).ArtistId(artistId).AlbumId(albumId).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListApiV1HistoryArtist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1HistoryArtist`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListApiV1HistoryArtist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1HistoryArtistRequest struct via the builder pattern


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


## ListApiV1HistorySince

> []HistoryResource ListApiV1HistorySince(ctx).Date(date).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



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
    resp, r, err := apiClient.HistoryApi.ListApiV1HistorySince(context.Background()).Date(date).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListApiV1HistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1HistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListApiV1HistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1HistorySinceRequest struct via the builder pattern


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

