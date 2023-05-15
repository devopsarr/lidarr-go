# \DownloadClientAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownloadClient**](DownloadClientAPI.md#CreateDownloadClient) | **Post** /api/v1/downloadclient | 
[**CreateDownloadClientActionByName**](DownloadClientAPI.md#CreateDownloadClientActionByName) | **Post** /api/v1/downloadclient/action/{name} | 
[**DeleteDownloadClient**](DownloadClientAPI.md#DeleteDownloadClient) | **Delete** /api/v1/downloadclient/{id} | 
[**GetDownloadClientById**](DownloadClientAPI.md#GetDownloadClientById) | **Get** /api/v1/downloadclient/{id} | 
[**ListDownloadClient**](DownloadClientAPI.md#ListDownloadClient) | **Get** /api/v1/downloadclient | 
[**ListDownloadClientSchema**](DownloadClientAPI.md#ListDownloadClientSchema) | **Get** /api/v1/downloadclient/schema | 
[**TestDownloadClient**](DownloadClientAPI.md#TestDownloadClient) | **Post** /api/v1/downloadclient/test | 
[**TestallDownloadClient**](DownloadClientAPI.md#TestallDownloadClient) | **Post** /api/v1/downloadclient/testall | 
[**UpdateDownloadClient**](DownloadClientAPI.md#UpdateDownloadClient) | **Put** /api/v1/downloadclient/{id} | 



## CreateDownloadClient

> DownloadClientResource CreateDownloadClient(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *lidarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientAPI.CreateDownloadClient(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.CreateDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDownloadClient`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.CreateDownloadClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDownloadClientActionByName

> CreateDownloadClientActionByName(ctx, name).DownloadClientResource(downloadClientResource).Execute()



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
    name := "name_example" // string | 
    downloadClientResource := *lidarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientAPI.CreateDownloadClientActionByName(context.Background(), name).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.CreateDownloadClientActionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadClientActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDownloadClient

> DeleteDownloadClient(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientAPI.DeleteDownloadClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.DeleteDownloadClient``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDownloadClientRequest struct via the builder pattern


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


## GetDownloadClientById

> DownloadClientResource GetDownloadClientById(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientAPI.GetDownloadClientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.GetDownloadClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadClientById`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.GetDownloadClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDownloadClient

> []DownloadClientResource ListDownloadClient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientAPI.ListDownloadClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.ListDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDownloadClient`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.ListDownloadClient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDownloadClientRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDownloadClientSchema

> []DownloadClientResource ListDownloadClientSchema(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientAPI.ListDownloadClientSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.ListDownloadClientSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDownloadClientSchema`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.ListDownloadClientSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDownloadClientSchemaRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDownloadClient

> TestDownloadClient(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *lidarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientAPI.TestDownloadClient(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.TestDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestallDownloadClient

> TestallDownloadClient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientAPI.TestallDownloadClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.TestallDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallDownloadClientRequest struct via the builder pattern


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


## UpdateDownloadClient

> DownloadClientResource UpdateDownloadClient(ctx, id).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *lidarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientAPI.UpdateDownloadClient(context.Background(), id).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.UpdateDownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDownloadClient`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.UpdateDownloadClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
