# \CalendarApi

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCalendarById**](CalendarApi.md#GetCalendarById) | **Get** /api/v1/calendar/{id} | 
[**ListCalendar**](CalendarApi.md#ListCalendar) | **Get** /api/v1/calendar | 



## GetCalendarById

> AlbumResource GetCalendarById(ctx, id).Execute()



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
    resp, r, err := apiClient.CalendarApi.GetCalendarById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.GetCalendarById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCalendarById`: AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.GetCalendarById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCalendarByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlbumResource**](AlbumResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCalendar

> []AlbumResource ListCalendar(ctx).Start(start).End(end).Unmonitored(unmonitored).IncludeArtist(includeArtist).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    lidarrClient "./openapi"
)

func main() {
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    unmonitored := true // bool |  (optional) (default to false)
    includeArtist := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarApi.ListCalendar(context.Background()).Start(start).End(end).Unmonitored(unmonitored).IncludeArtist(includeArtist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.ListCalendar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCalendar`: []AlbumResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.ListCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **unmonitored** | **bool** |  | [default to false]
 **includeArtist** | **bool** |  | [default to false]

### Return type

[**[]AlbumResource**](AlbumResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

