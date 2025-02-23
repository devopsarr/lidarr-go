# \AlbumLookupAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAlbumLookup**](AlbumLookupAPI.md#ListAlbumLookup) | **Get** /api/v1/album/lookup | 



## ListAlbumLookup

> []AlbumResource ListAlbumLookup(ctx).Term(term).Execute()



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
	term := "term_example" // string |  (optional)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlbumLookupAPI.ListAlbumLookup(context.Background()).Term(term).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlbumLookupAPI.ListAlbumLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlbumLookup`: []AlbumResource
	fmt.Fprintf(os.Stdout, "Response from `AlbumLookupAPI.ListAlbumLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlbumLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** |  | 

### Return type

[**[]AlbumResource**](AlbumResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

