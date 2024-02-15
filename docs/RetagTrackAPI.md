# \RetagTrackAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRetag**](RetagTrackAPI.md#ListRetag) | **Get** /api/v1/retag | 



## ListRetag

> []RetagTrackResource ListRetag(ctx).ArtistId(artistId).AlbumId(albumId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	lidarrClient "github.com/devopsarr/lidarr-go/lidarr"
)

func main() {
	artistId := int32(56) // int32 |  (optional)
	albumId := int32(56) // int32 |  (optional)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetagTrackAPI.ListRetag(context.Background()).ArtistId(artistId).AlbumId(albumId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetagTrackAPI.ListRetag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRetag`: []RetagTrackResource
	fmt.Fprintf(os.Stdout, "Response from `RetagTrackAPI.ListRetag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRetagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumId** | **int32** |  | 

### Return type

[**[]RetagTrackResource**](RetagTrackResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

