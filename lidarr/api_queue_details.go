/*
Lidarr

Lidarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lidarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// QueueDetailsApiService QueueDetailsApi service
type QueueDetailsApiService service
type ApiGetQueueDetailsByIdRequest struct {
	ctx context.Context
	ApiService *QueueDetailsApiService
	id int32
}

func (r ApiGetQueueDetailsByIdRequest) Execute() (*QueueResource, *http.Response, error) {
	return r.ApiService.GetQueueDetailsByIdExecute(r)
}

/*
GetQueueDetailsById Method for GetQueueDetailsById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetQueueDetailsByIdRequest
*/
func (a *QueueDetailsApiService) GetQueueDetailsById(ctx context.Context, id int32) ApiGetQueueDetailsByIdRequest {
	return ApiGetQueueDetailsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return QueueResource
func (a *QueueDetailsApiService) GetQueueDetailsByIdExecute(r ApiGetQueueDetailsByIdRequest) (*QueueResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueueResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueDetailsApiService.GetQueueDetailsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/queue/details/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type ApiListQueueDetailsRequest struct {
	ctx context.Context
	ApiService *QueueDetailsApiService
	artistId *int32
	albumIds *[]int32
	includeArtist *bool
	includeAlbum *bool
}

func (r ApiListQueueDetailsRequest) ArtistId(artistId int32) ApiListQueueDetailsRequest {
	r.artistId = &artistId
	return r
}

func (r ApiListQueueDetailsRequest) AlbumIds(albumIds []int32) ApiListQueueDetailsRequest {
	r.albumIds = &albumIds
	return r
}

func (r ApiListQueueDetailsRequest) IncludeArtist(includeArtist bool) ApiListQueueDetailsRequest {
	r.includeArtist = &includeArtist
	return r
}

func (r ApiListQueueDetailsRequest) IncludeAlbum(includeAlbum bool) ApiListQueueDetailsRequest {
	r.includeAlbum = &includeAlbum
	return r
}

func (r ApiListQueueDetailsRequest) Execute() ([]*QueueResource, *http.Response, error) {
	return r.ApiService.ListQueueDetailsExecute(r)
}

/*
ListQueueDetails Method for ListQueueDetails

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListQueueDetailsRequest
*/
func (a *QueueDetailsApiService) ListQueueDetails(ctx context.Context) ApiListQueueDetailsRequest {
	return ApiListQueueDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []QueueResource
func (a *QueueDetailsApiService) ListQueueDetailsExecute(r ApiListQueueDetailsRequest) ([]*QueueResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*QueueResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueDetailsApiService.ListQueueDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/queue/details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.artistId != nil {
		localVarQueryParams.Add("artistId", parameterToString(*r.artistId, ""))
	}
	if r.albumIds != nil {
		t := *r.albumIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("albumIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("albumIds", parameterToString(t, "multi"))
		}
	}
	if r.includeArtist != nil {
		localVarQueryParams.Add("includeArtist", parameterToString(*r.includeArtist, ""))
	}
	if r.includeAlbum != nil {
		localVarQueryParams.Add("includeAlbum", parameterToString(*r.includeAlbum, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
