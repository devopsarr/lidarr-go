# \MissingApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWantedMissing**](MissingApi.md#GetWantedMissing) | **Get** /api/v1/wanted/missing | 
[**GetWantedMissingById**](MissingApi.md#GetWantedMissingById) | **Get** /api/v1/wanted/missing/{id} | 



## GetWantedMissing

> AlbumResourcePagingResource GetWantedMissing(ctx).IncludeArtist(includeArtist).Execute()



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

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MissingApi.GetWantedMissing(context.Background()).IncludeArtist(includeArtist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MissingApi.GetWantedMissing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWantedMissing`: AlbumResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `MissingApi.GetWantedMissing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedMissingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArtist** | **bool** |  | [default to false]

### Return type

[**AlbumResourcePagingResource**](AlbumResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWantedMissingById

> AlbumResource GetWantedMissingById(ctx, id).Execute()



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
    resp, r, err := apiClient.MissingApi.GetWantedMissingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MissingApi.GetWantedMissingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWantedMissingById`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `MissingApi.GetWantedMissingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedMissingByIdRequest struct via the builder pattern


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

