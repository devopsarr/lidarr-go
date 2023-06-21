# \ReleaseApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelease**](ReleaseApi.md#CreateRelease) | **Post** /api/v1/release | 
[**ListRelease**](ReleaseApi.md#ListRelease) | **Get** /api/v1/release | 



## CreateRelease

> ReleaseResource CreateRelease(ctx).ReleaseResource(releaseResource).Execute()



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
    releaseResource := *lidarrClient.NewReleaseResource() // ReleaseResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.CreateRelease(context.Background()).ReleaseResource(releaseResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.CreateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelease`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.CreateRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md) |  | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRelease

> []ReleaseResource ListRelease(ctx).AlbumId(albumId).ArtistId(artistId).Execute()



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
    albumId := int32(56) // int32 |  (optional)
    artistId := int32(56) // int32 |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.ListRelease(context.Background()).AlbumId(albumId).ArtistId(artistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.ListRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRelease`: []ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.ListRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumId** | **int32** |  | 
 **artistId** | **int32** |  | 

### Return type

[**[]ReleaseResource**](ReleaseResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

