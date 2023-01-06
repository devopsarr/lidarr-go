# \TrackApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1TrackById**](TrackApi.md#GetApiV1TrackById) | **Get** /api/v1/track/{id} | 
[**ListApiV1Track**](TrackApi.md#ListApiV1Track) | **Get** /api/v1/track | 



## GetApiV1TrackById

> TrackResource GetApiV1TrackById(ctx, id).Execute()



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
    resp, r, err := apiClient.TrackApi.GetApiV1TrackById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackApi.GetApiV1TrackById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1TrackById`: TrackResource
    fmt.Fprintf(os.Stdout, "Response from `TrackApi.GetApiV1TrackById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1TrackByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrackResource**](TrackResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1Track

> []TrackResource ListApiV1Track(ctx).ArtistId(artistId).AlbumId(albumId).AlbumReleaseId(albumReleaseId).TrackIds(trackIds).Execute()



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
    albumReleaseId := int32(56) // int32 |  (optional)
    trackIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrackApi.ListApiV1Track(context.Background()).ArtistId(artistId).AlbumId(albumId).AlbumReleaseId(albumReleaseId).TrackIds(trackIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackApi.ListApiV1Track``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Track`: []TrackResource
    fmt.Fprintf(os.Stdout, "Response from `TrackApi.ListApiV1Track`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1TrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumId** | **int32** |  | 
 **albumReleaseId** | **int32** |  | 
 **trackIds** | **[]int32** |  | 

### Return type

[**[]TrackResource**](TrackResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

