# \RenameTrackApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiV1Rename**](RenameTrackApi.md#ListApiV1Rename) | **Get** /api/v1/rename | 



## ListApiV1Rename

> []RenameTrackResource ListApiV1Rename(ctx).ArtistId(artistId).AlbumId(albumId).Execute()



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

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenameTrackApi.ListApiV1Rename(context.Background()).ArtistId(artistId).AlbumId(albumId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenameTrackApi.ListApiV1Rename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Rename`: []RenameTrackResource
    fmt.Fprintf(os.Stdout, "Response from `RenameTrackApi.ListApiV1Rename`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1RenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumId** | **int32** |  | 

### Return type

[**[]RenameTrackResource**](RenameTrackResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

