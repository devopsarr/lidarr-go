# \AlbumStudioApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1AlbumStudio**](AlbumStudioApi.md#CreateApiV1AlbumStudio) | **Post** /api/v1/albumstudio | 



## CreateApiV1AlbumStudio

> CreateApiV1AlbumStudio(ctx).AlbumStudioResource(albumStudioResource).Execute()



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
    albumStudioResource := *lidarrClient.NewAlbumStudioResource() // AlbumStudioResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlbumStudioApi.CreateApiV1AlbumStudio(context.Background()).AlbumStudioResource(albumStudioResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumStudioApi.CreateApiV1AlbumStudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1AlbumStudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumStudioResource** | [**AlbumStudioResource**](AlbumStudioResource.md) |  | 

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

