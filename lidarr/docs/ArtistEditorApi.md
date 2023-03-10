# \ArtistEditorApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArtistEditor**](ArtistEditorApi.md#DeleteArtistEditor) | **Delete** /api/v1/artist/editor | 
[**PutArtistEditor**](ArtistEditorApi.md#PutArtistEditor) | **Put** /api/v1/artist/editor | 



## DeleteArtistEditor

> DeleteArtistEditor(ctx).ArtistEditorResource(artistEditorResource).Execute()



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
    artistEditorResource := *lidarrClient.NewArtistEditorResource() // ArtistEditorResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistEditorApi.DeleteArtistEditor(context.Background()).ArtistEditorResource(artistEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistEditorApi.DeleteArtistEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtistEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistEditorResource** | [**ArtistEditorResource**](ArtistEditorResource.md) |  | 

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


## PutArtistEditor

> PutArtistEditor(ctx).ArtistEditorResource(artistEditorResource).Execute()



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
    artistEditorResource := *lidarrClient.NewArtistEditorResource() // ArtistEditorResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistEditorApi.PutArtistEditor(context.Background()).ArtistEditorResource(artistEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistEditorApi.PutArtistEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutArtistEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistEditorResource** | [**ArtistEditorResource**](ArtistEditorResource.md) |  | 

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

