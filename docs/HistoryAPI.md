# \HistoryAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryFailedById**](HistoryAPI.md#CreateHistoryFailedById) | **Post** /api/v1/history/failed/{id} | 
[**GetHistory**](HistoryAPI.md#GetHistory) | **Get** /api/v1/history | 
[**ListHistoryArtist**](HistoryAPI.md#ListHistoryArtist) | **Get** /api/v1/history/artist | 
[**ListHistorySince**](HistoryAPI.md#ListHistorySince) | **Get** /api/v1/history/since | 



## CreateHistoryFailedById

> CreateHistoryFailedById(ctx, id).Execute()



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
	r, err := apiClient.HistoryAPI.CreateHistoryFailedById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.CreateHistoryFailedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHistoryFailedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistory

> HistoryResourcePagingResource GetHistory(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).EventType(eventType).AlbumId(albumId).DownloadId(downloadId).ArtistIds(artistIds).Quality(quality).Execute()



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
	includeArtist := true // bool |  (optional)
	includeAlbum := true // bool |  (optional)
	includeTrack := true // bool |  (optional)
	eventType := []int32{int32(123)} // []int32 |  (optional)
	albumId := int32(56) // int32 |  (optional)
	downloadId := "downloadId_example" // string |  (optional)
	artistIds := []int32{int32(123)} // []int32 |  (optional)
	quality := []int32{int32(123)} // []int32 |  (optional)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryAPI.GetHistory(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).EventType(eventType).AlbumId(albumId).DownloadId(downloadId).ArtistIds(artistIds).Quality(quality).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.GetHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistory`: HistoryResourcePagingResource
	fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **includeArtist** | **bool** |  | 
 **includeAlbum** | **bool** |  | 
 **includeTrack** | **bool** |  | 
 **eventType** | **[]int32** |  | 
 **albumId** | **int32** |  | 
 **downloadId** | **string** |  | 
 **artistIds** | **[]int32** |  | 
 **quality** | **[]int32** |  | 

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryArtist

> []HistoryResource ListHistoryArtist(ctx).ArtistId(artistId).AlbumId(albumId).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



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
	eventType := lidarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
	includeArtist := true // bool |  (optional) (default to false)
	includeAlbum := true // bool |  (optional) (default to false)
	includeTrack := true // bool |  (optional) (default to false)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryAPI.ListHistoryArtist(context.Background()).ArtistId(artistId).AlbumId(albumId).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistoryArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistoryArtist`: []HistoryResource
	fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistoryArtist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **int32** |  | 
 **albumId** | **int32** |  | 
 **eventType** | [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | 
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to false]
 **includeTrack** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorySince

> []HistoryResource ListHistorySince(ctx).Date(date).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	lidarrClient "github.com/devopsarr/lidarr-go/lidarr"
)

func main() {
	date := time.Now() // time.Time |  (optional)
	eventType := lidarrClient.EntityHistoryEventType("unknown") // EntityHistoryEventType |  (optional)
	includeArtist := true // bool |  (optional) (default to false)
	includeAlbum := true // bool |  (optional) (default to false)
	includeTrack := true // bool |  (optional) (default to false)

	configuration := lidarrClient.NewConfiguration()
	apiClient := lidarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoryAPI.ListHistorySince(context.Background()).Date(date).EventType(eventType).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).IncludeTrack(includeTrack).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistorySince``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistorySince`: []HistoryResource
	fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistorySinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**EntityHistoryEventType**](EntityHistoryEventType.md) |  | 
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to false]
 **includeTrack** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

