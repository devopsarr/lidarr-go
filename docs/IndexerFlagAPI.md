# \IndexerFlagAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIndexerFlag**](IndexerFlagAPI.md#ListIndexerFlag) | **Get** /api/v1/indexerflag | 



## ListIndexerFlag

> []IndexerFlagResource ListIndexerFlag(ctx).Execute()



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

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexerFlagAPI.ListIndexerFlag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerFlagAPI.ListIndexerFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndexerFlag`: []IndexerFlagResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerFlagAPI.ListIndexerFlag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerFlagRequest struct via the builder pattern


### Return type

[**[]IndexerFlagResource**](IndexerFlagResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

