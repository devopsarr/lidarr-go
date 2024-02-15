# \CutoffAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWantedCutoff**](CutoffAPI.md#GetWantedCutoff) | **Get** /api/v1/wanted/cutoff | 
[**GetWantedCutoffById**](CutoffAPI.md#GetWantedCutoffById) | **Get** /api/v1/wanted/cutoff/{id} | 



## GetWantedCutoff

> AlbumResourcePagingResource GetWantedCutoff(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeArtist(includeArtist).Monitored(monitored).Execute()



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sortKey := "sortKey_example" // string |  (optional)
	sortDirection := lidarrClient.SortDirection("default") // SortDirection |  (optional)
	includeArtist := true // bool |  (optional) (default to false)
	monitored := true // bool |  (optional) (default to true)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.CutoffAPI.GetWantedCutoff(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeArtist(includeArtist).Monitored(monitored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CutoffAPI.GetWantedCutoff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWantedCutoff`: AlbumResourcePagingResource
	fmt.Fprintf(os.Stdout, "Response from `CutoffAPI.GetWantedCutoff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedCutoffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **includeArtist** | **bool** |  | [default to false]
 **monitored** | **bool** |  | [default to true]

### Return type

[**AlbumResourcePagingResource**](AlbumResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWantedCutoffById

> AlbumResource GetWantedCutoffById(ctx, id).Execute()



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
	id := int32(56) // int32 | 

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.CutoffAPI.GetWantedCutoffById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CutoffAPI.GetWantedCutoffById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWantedCutoffById`: AlbumResource
	fmt.Fprintf(os.Stdout, "Response from `CutoffAPI.GetWantedCutoffById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedCutoffByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlbumResource**](AlbumResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

