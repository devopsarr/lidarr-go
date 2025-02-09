# \ArtistLookupAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListArtistLookup**](ArtistLookupAPI.md#ListArtistLookup) | **Get** /api/v1/artist/lookup | 



## ListArtistLookup

> []ArtistResource ListArtistLookup(ctx).Term(term).Execute()



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
	resp, r, err := apiClient.ArtistLookupAPI.ListArtistLookup(context.Background()).Term(term).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistLookupAPI.ListArtistLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArtistLookup`: []ArtistResource
	fmt.Fprintf(os.Stdout, "Response from `ArtistLookupAPI.ListArtistLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArtistLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** |  | 

### Return type

[**[]ArtistResource**](ArtistResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

