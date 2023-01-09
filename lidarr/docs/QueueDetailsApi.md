# \QueueDetailsApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueueDetailsById**](QueueDetailsApi.md#GetQueueDetailsById) | **Get** /api/v1/queue/details/{id} | 
[**ListQueueDetails**](QueueDetailsApi.md#ListQueueDetails) | **Get** /api/v1/queue/details | 



## GetQueueDetailsById

> QueueResource GetQueueDetailsById(ctx, id).Execute()



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
    resp, r, err := apiClient.QueueDetailsApi.GetQueueDetailsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.GetQueueDetailsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueDetailsById`: QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.GetQueueDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueResource**](QueueResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueueDetails

> []QueueResource ListQueueDetails(ctx).ArtistId(artistId).AlbumIds(albumIds).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).Execute()



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
    includeArtist := true // bool |  (optional) (default to false)
    includeAlbum := true // bool |  (optional) (default to true)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueDetailsApi.ListQueueDetails(context.Background()).ArtistId(artistId).AlbumIds(albumIds).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.ListQueueDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueueDetails`: []QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.ListQueueDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueueDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumIds** | **[]int32** |  | 
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to true]

### Return type

[**[]QueueResource**](QueueResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

