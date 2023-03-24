# \MediaCoverApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaCoverAlbumalbumIdByFilename**](MediaCoverApi.md#GetMediaCoverAlbumalbumIdByFilename) | **Get** /api/v1/mediacover/album/{albumId}/{filename} | 
[**GetMediaCoverArtistartistIdByFilename**](MediaCoverApi.md#GetMediaCoverArtistartistIdByFilename) | **Get** /api/v1/mediacover/artist/{artistId}/{filename} | 



## GetMediaCoverAlbumalbumIdByFilename

> GetMediaCoverAlbumalbumIdByFilename(ctx, albumId, filename).Execute()



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
    resp, r, err := apiClient.MediaCoverApi.GetMediaCoverAlbumalbumIdByFilename(context.Background(), albumId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetMediaCoverAlbumalbumIdByFilename``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMediaCoverAlbumalbumIdByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetMediaCoverArtistartistIdByFilename

> GetMediaCoverArtistartistIdByFilename(ctx, artistId, filename).Execute()



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
    resp, r, err := apiClient.MediaCoverApi.GetMediaCoverArtistartistIdByFilename(context.Background(), artistId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.GetMediaCoverArtistartistIdByFilename``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMediaCoverArtistartistIdByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

