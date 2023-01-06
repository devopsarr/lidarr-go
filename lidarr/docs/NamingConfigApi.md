# \NamingConfigApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigNaming**](NamingConfigApi.md#GetApiV1ConfigNaming) | **Get** /api/v1/config/naming | 
[**GetApiV1ConfigNamingById**](NamingConfigApi.md#GetApiV1ConfigNamingById) | **Get** /api/v1/config/naming/{id} | 
[**GetApiV1ConfigNamingExamples**](NamingConfigApi.md#GetApiV1ConfigNamingExamples) | **Get** /api/v1/config/naming/examples | 
[**UpdateApiV1ConfigNaming**](NamingConfigApi.md#UpdateApiV1ConfigNaming) | **Put** /api/v1/config/naming/{id} | 



## GetApiV1ConfigNaming

> NamingConfigResource GetApiV1ConfigNaming(ctx).Execute()



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

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigApi.GetApiV1ConfigNaming(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.GetApiV1ConfigNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigNaming`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigApi.GetApiV1ConfigNaming`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigNamingRequest struct via the builder pattern


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ConfigNamingById

> NamingConfigResource GetApiV1ConfigNamingById(ctx, id).Execute()



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
    resp, r, err := apiClient.NamingConfigApi.GetApiV1ConfigNamingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.GetApiV1ConfigNamingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigNamingById`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigApi.GetApiV1ConfigNamingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigNamingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ConfigNamingExamples

> GetApiV1ConfigNamingExamples(ctx).RenameTracks(renameTracks).ReplaceIllegalCharacters(replaceIllegalCharacters).StandardTrackFormat(standardTrackFormat).MultiDiscTrackFormat(multiDiscTrackFormat).ArtistFolderFormat(artistFolderFormat).IncludeArtistName(includeArtistName).IncludeAlbumTitle(includeAlbumTitle).IncludeQuality(includeQuality).ReplaceSpaces(replaceSpaces).Separator(separator).NumberStyle(numberStyle).Id(id).ResourceName(resourceName).Execute()



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
    renameTracks := true // bool |  (optional)
    replaceIllegalCharacters := true // bool |  (optional)
    standardTrackFormat := "standardTrackFormat_example" // string |  (optional)
    multiDiscTrackFormat := "multiDiscTrackFormat_example" // string |  (optional)
    artistFolderFormat := "artistFolderFormat_example" // string |  (optional)
    includeArtistName := true // bool |  (optional)
    includeAlbumTitle := true // bool |  (optional)
    includeQuality := true // bool |  (optional)
    replaceSpaces := true // bool |  (optional)
    separator := "separator_example" // string |  (optional)
    numberStyle := "numberStyle_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    resourceName := "resourceName_example" // string |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigApi.GetApiV1ConfigNamingExamples(context.Background()).RenameTracks(renameTracks).ReplaceIllegalCharacters(replaceIllegalCharacters).StandardTrackFormat(standardTrackFormat).MultiDiscTrackFormat(multiDiscTrackFormat).ArtistFolderFormat(artistFolderFormat).IncludeArtistName(includeArtistName).IncludeAlbumTitle(includeAlbumTitle).IncludeQuality(includeQuality).ReplaceSpaces(replaceSpaces).Separator(separator).NumberStyle(numberStyle).Id(id).ResourceName(resourceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.GetApiV1ConfigNamingExamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigNamingExamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameTracks** | **bool** |  | 
 **replaceIllegalCharacters** | **bool** |  | 
 **standardTrackFormat** | **string** |  | 
 **multiDiscTrackFormat** | **string** |  | 
 **artistFolderFormat** | **string** |  | 
 **includeArtistName** | **bool** |  | 
 **includeAlbumTitle** | **bool** |  | 
 **includeQuality** | **bool** |  | 
 **replaceSpaces** | **bool** |  | 
 **separator** | **string** |  | 
 **numberStyle** | **string** |  | 
 **id** | **int32** |  | 
 **resourceName** | **string** |  | 

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


## UpdateApiV1ConfigNaming

> NamingConfigResource UpdateApiV1ConfigNaming(ctx, id).NamingConfigResource(namingConfigResource).Execute()



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
    id := "id_example" // string | 
    namingConfigResource := *lidarrClient.NewNamingConfigResource() // NamingConfigResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigApi.UpdateApiV1ConfigNaming(context.Background(), id).NamingConfigResource(namingConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.UpdateApiV1ConfigNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigNaming`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigApi.UpdateApiV1ConfigNaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namingConfigResource** | [**NamingConfigResource**](NamingConfigResource.md) |  | 

### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

