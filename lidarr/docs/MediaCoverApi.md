# \MediaCoverApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1MediaCoverAlbumalbumIdByFilename**](MediaCoverApi.md#GetApiV1MediaCoverAlbumalbumIdByFilename) | **Get** /api/v1/mediacover/album/{albumId}/{filename} | 
[**GetApiV1MediaCoverArtistartistIdByFilename**](MediaCoverApi.md#GetApiV1MediaCoverArtistartistIdByFilename) | **Get** /api/v1/mediacover/artist/{artistId}/{filename} | 



## GetApiV1MediaCoverAlbumalbumIdByFilename

> GetApiV1MediaCoverAlbumalbumIdByFilename(ctx, albumId, filename).Execute()



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
    albumId := int32(56) // int32 | 
    filename := "filename_example" // string | 

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaCoverApi.GetApiV1MediaCoverAlbumalbumIdByFilename(context.Background(), albumId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetApiV1MediaCoverAlbumalbumIdByFilename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MediaCoverAlbumalbumIdByFilenameRequest struct via the builder pattern


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


## GetApiV1MediaCoverArtistartistIdByFilename

> GetApiV1MediaCoverArtistartistIdByFilename(ctx, artistId, filename).Execute()



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
    artistId := int32(56) // int32 | 
    filename := "filename_example" // string | 

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaCoverApi.GetApiV1MediaCoverArtistartistIdByFilename(context.Background(), artistId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetApiV1MediaCoverArtistartistIdByFilename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artistId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MediaCoverArtistartistIdByFilenameRequest struct via the builder pattern


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

