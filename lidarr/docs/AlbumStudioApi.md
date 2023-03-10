# \AlbumStudioApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlbumStudio**](AlbumStudioApi.md#CreateAlbumStudio) | **Post** /api/v1/albumstudio | 



## CreateAlbumStudio

> CreateAlbumStudio(ctx).AlbumStudioResource(albumStudioResource).Execute()



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
    resp, r, err := apiClient.AlbumStudioApi.CreateAlbumStudio(context.Background()).AlbumStudioResource(albumStudioResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlbumStudioApi.CreateAlbumStudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlbumStudioRequest struct via the builder pattern


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

