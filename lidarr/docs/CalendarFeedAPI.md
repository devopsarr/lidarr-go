# \CalendarFeedAPI

All URIs are relative to *http://localhost:8686*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedV1CalendarLidarrIcs**](CalendarFeedAPI.md#GetFeedV1CalendarLidarrIcs) | **Get** /feed/v1/calendar/lidarr.ics | 



## GetFeedV1CalendarLidarrIcs

> GetFeedV1CalendarLidarrIcs(ctx).PastDays(pastDays).FutureDays(futureDays).Tags(tags).Unmonitored(unmonitored).Execute()



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
    pastDays := int32(56) // int32 |  (optional) (default to 7)
    futureDays := int32(56) // int32 |  (optional) (default to 28)
    tags := "tags_example" // string |  (optional) (default to "")
    unmonitored := true // bool |  (optional) (default to false)

    configuration := lidarrClient.NewConfiguration()
    apiClient := lidarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarFeedAPI.GetFeedV1CalendarLidarrIcs(context.Background()).PastDays(pastDays).FutureDays(futureDays).Tags(tags).Unmonitored(unmonitored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarFeedAPI.GetFeedV1CalendarLidarrIcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedV1CalendarLidarrIcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pastDays** | **int32** |  | [default to 7]
 **futureDays** | **int32** |  | [default to 28]
 **tags** | **string** |  | [default to &quot;&quot;]
 **unmonitored** | **bool** |  | [default to false]

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

