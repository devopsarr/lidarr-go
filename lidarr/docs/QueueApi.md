# \QueueApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1Queue**](QueueApi.md#DeleteApiV1Queue) | **Delete** /api/v1/queue/{id} | 
[**DeleteApiV1QueueBulk**](QueueApi.md#DeleteApiV1QueueBulk) | **Delete** /api/v1/queue/bulk | 
[**GetApiV1Queue**](QueueApi.md#GetApiV1Queue) | **Get** /api/v1/queue | 
[**GetApiV1QueueById**](QueueApi.md#GetApiV1QueueById) | **Get** /api/v1/queue/{id} | 



## DeleteApiV1Queue

> DeleteApiV1Queue(ctx, id).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).Execute()



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
    id := int32(56) // int32 | 
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    skipReDownload := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.DeleteApiV1Queue(context.Background(), id).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.DeleteApiV1Queue``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1QueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **skipReDownload** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1QueueBulk

> DeleteApiV1QueueBulk(ctx).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).QueueBulkResource(queueBulkResource).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    skipReDownload := true // bool |  (optional) (default to false)
    queueBulkResource := *lidarrClient.NewQueueBulkResource() // QueueBulkResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.DeleteApiV1QueueBulk(context.Background()).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipReDownload(skipReDownload).QueueBulkResource(queueBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.DeleteApiV1QueueBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1QueueBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **skipReDownload** | **bool** |  | [default to false]
 **queueBulkResource** | [**QueueBulkResource**](QueueBulkResource.md) |  | 

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


## GetApiV1Queue

> QueueResourcePagingResource GetApiV1Queue(ctx).IncludeUnknownArtistItems(includeUnknownArtistItems).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).Execute()



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
    includeUnknownArtistItems := true // bool |  (optional) (default to false)
    includeArtist := true // bool |  (optional) (default to false)
    includeAlbum := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.GetApiV1Queue(context.Background()).IncludeUnknownArtistItems(includeUnknownArtistItems).IncludeArtist(includeArtist).IncludeAlbum(includeAlbum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.GetApiV1Queue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1Queue`: QueueResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `QueueApi.GetApiV1Queue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUnknownArtistItems** | **bool** |  | [default to false]
 **includeArtist** | **bool** |  | [default to false]
 **includeAlbum** | **bool** |  | [default to false]

### Return type

[**QueueResourcePagingResource**](QueueResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1QueueById

> QueueResource GetApiV1QueueById(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.GetApiV1QueueById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.GetApiV1QueueById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1QueueById`: QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueApi.GetApiV1QueueById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1QueueByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueResource**](QueueResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

