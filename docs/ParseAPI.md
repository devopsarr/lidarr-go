# \ParseAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParse**](ParseAPI.md#GetParse) | **Get** /api/v1/parse | 



## GetParse

> ParseResource GetParse(ctx).Title(title).Execute()



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
	title := "title_example" // string |  (optional)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParseAPI.GetParse(context.Background()).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParseAPI.GetParse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParse`: ParseResource
	fmt.Fprintf(os.Stdout, "Response from `ParseAPI.GetParse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** |  | 

### Return type

[**ParseResource**](ParseResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

