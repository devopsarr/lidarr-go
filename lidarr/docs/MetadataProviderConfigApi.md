# \MetadataProviderConfigApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataProviderConfig**](MetadataProviderConfigApi.md#GetMetadataProviderConfig) | **Get** /api/v1/config/metadataprovider | 
[**GetMetadataProviderConfigById**](MetadataProviderConfigApi.md#GetMetadataProviderConfigById) | **Get** /api/v1/config/metadataprovider/{id} | 
[**UpdateMetadataProviderConfig**](MetadataProviderConfigApi.md#UpdateMetadataProviderConfig) | **Put** /api/v1/config/metadataprovider/{id} | 



## GetMetadataProviderConfig

> MetadataProviderConfigResource GetMetadataProviderConfig(ctx).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigApi.GetMetadataProviderConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigApi.GetMetadataProviderConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataProviderConfig`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigApi.GetMetadataProviderConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataProviderConfigRequest struct via the builder pattern


### Return type

[**MetadataProviderConfigResource**](MetadataProviderConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataProviderConfigById

> MetadataProviderConfigResource GetMetadataProviderConfigById(ctx, id).Execute()



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
    resp, r, err := apiClient.MetadataProviderConfigApi.GetMetadataProviderConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigApi.GetMetadataProviderConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataProviderConfigById`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigApi.GetMetadataProviderConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataProviderConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataProviderConfigResource**](MetadataProviderConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadataProviderConfig

> MetadataProviderConfigResource UpdateMetadataProviderConfig(ctx, id).MetadataProviderConfigResource(metadataProviderConfigResource).Execute()



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
    metadataProviderConfigResource := *lidarrClient.NewMetadataProviderConfigResource() // MetadataProviderConfigResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProviderConfigApi.UpdateMetadataProviderConfig(context.Background(), id).MetadataProviderConfigResource(metadataProviderConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProviderConfigApi.UpdateMetadataProviderConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetadataProviderConfig`: MetadataProviderConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MetadataProviderConfigApi.UpdateMetadataProviderConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataProviderConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataProviderConfigResource** | [**MetadataProviderConfigResource**](MetadataProviderConfigResource.md) |  | 

### Return type

[**MetadataProviderConfigResource**](MetadataProviderConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

