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
	"time"
)


// HistoryAPIService HistoryAPI service
type HistoryAPIService service
type ApiCreateHistoryFailedByIdRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	id int32
}

func (r ApiCreateHistoryFailedByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateHistoryFailedByIdExecute(r)
}

/*
CreateHistoryFailedById Method for CreateHistoryFailedById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCreateHistoryFailedByIdRequest
*/
func (a *HistoryAPIService) CreateHistoryFailedById(ctx context.Context, id int32) ApiCreateHistoryFailedByIdRequest {
	return ApiCreateHistoryFailedByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *HistoryAPIService) CreateHistoryFailedByIdExecute(r ApiCreateHistoryFailedByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.CreateHistoryFailedById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history/failed/{id}"
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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type ApiGetHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	includeArtist *bool
	includeAlbum *bool
	includeTrack *bool
}

func (r ApiGetHistoryRequest) IncludeArtist(includeArtist bool) ApiGetHistoryRequest {
	r.includeArtist = &includeArtist
	return r
}

func (r ApiGetHistoryRequest) IncludeAlbum(includeAlbum bool) ApiGetHistoryRequest {
	r.includeAlbum = &includeAlbum
	return r
}

func (r ApiGetHistoryRequest) IncludeTrack(includeTrack bool) ApiGetHistoryRequest {
	r.includeTrack = &includeTrack
	return r
}

func (r ApiGetHistoryRequest) Execute() (*HistoryResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetHistoryExecute(r)
}

/*
GetHistory Method for GetHistory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHistoryRequest
*/
func (a *HistoryAPIService) GetHistory(ctx context.Context) ApiGetHistoryRequest {
	return ApiGetHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistoryResourcePagingResource
func (a *HistoryAPIService) GetHistoryExecute(r ApiGetHistoryRequest) (*HistoryResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistoryResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.GetHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeArtist != nil {
		localVarQueryParams.Add("includeArtist", parameterToString(*r.includeArtist, ""))
	}
	if r.includeAlbum != nil {
		localVarQueryParams.Add("includeAlbum", parameterToString(*r.includeAlbum, ""))
	}
	if r.includeTrack != nil {
		localVarQueryParams.Add("includeTrack", parameterToString(*r.includeTrack, ""))
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
type ApiListHistoryArtistRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	artistId *int32
	albumId *int32
	eventType *EntityHistoryEventType
	includeArtist *bool
	includeAlbum *bool
	includeTrack *bool
}

func (r ApiListHistoryArtistRequest) ArtistId(artistId int32) ApiListHistoryArtistRequest {
	r.artistId = &artistId
	return r
}

func (r ApiListHistoryArtistRequest) AlbumId(albumId int32) ApiListHistoryArtistRequest {
	r.albumId = &albumId
	return r
}

func (r ApiListHistoryArtistRequest) EventType(eventType EntityHistoryEventType) ApiListHistoryArtistRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistoryArtistRequest) IncludeArtist(includeArtist bool) ApiListHistoryArtistRequest {
	r.includeArtist = &includeArtist
	return r
}

func (r ApiListHistoryArtistRequest) IncludeAlbum(includeAlbum bool) ApiListHistoryArtistRequest {
	r.includeAlbum = &includeAlbum
	return r
}

func (r ApiListHistoryArtistRequest) IncludeTrack(includeTrack bool) ApiListHistoryArtistRequest {
	r.includeTrack = &includeTrack
	return r
}

func (r ApiListHistoryArtistRequest) Execute() ([]*HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistoryArtistExecute(r)
}

/*
ListHistoryArtist Method for ListHistoryArtist

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistoryArtistRequest
*/
func (a *HistoryAPIService) ListHistoryArtist(ctx context.Context) ApiListHistoryArtistRequest {
	return ApiListHistoryArtistRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryAPIService) ListHistoryArtistExecute(r ApiListHistoryArtistRequest) ([]*HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.ListHistoryArtist")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history/artist"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.artistId != nil {
		localVarQueryParams.Add("artistId", parameterToString(*r.artistId, ""))
	}
	if r.albumId != nil {
		localVarQueryParams.Add("albumId", parameterToString(*r.albumId, ""))
	}
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.includeArtist != nil {
		localVarQueryParams.Add("includeArtist", parameterToString(*r.includeArtist, ""))
	}
	if r.includeAlbum != nil {
		localVarQueryParams.Add("includeAlbum", parameterToString(*r.includeAlbum, ""))
	}
	if r.includeTrack != nil {
		localVarQueryParams.Add("includeTrack", parameterToString(*r.includeTrack, ""))
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
type ApiListHistorySinceRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	date *time.Time
	eventType *EntityHistoryEventType
	includeArtist *bool
	includeAlbum *bool
	includeTrack *bool
}

func (r ApiListHistorySinceRequest) Date(date time.Time) ApiListHistorySinceRequest {
	r.date = &date
	return r
}

func (r ApiListHistorySinceRequest) EventType(eventType EntityHistoryEventType) ApiListHistorySinceRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistorySinceRequest) IncludeArtist(includeArtist bool) ApiListHistorySinceRequest {
	r.includeArtist = &includeArtist
	return r
}

func (r ApiListHistorySinceRequest) IncludeAlbum(includeAlbum bool) ApiListHistorySinceRequest {
	r.includeAlbum = &includeAlbum
	return r
}

func (r ApiListHistorySinceRequest) IncludeTrack(includeTrack bool) ApiListHistorySinceRequest {
	r.includeTrack = &includeTrack
	return r
}

func (r ApiListHistorySinceRequest) Execute() ([]*HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistorySinceExecute(r)
}

/*
ListHistorySince Method for ListHistorySince

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistorySinceRequest
*/
func (a *HistoryAPIService) ListHistorySince(ctx context.Context) ApiListHistorySinceRequest {
	return ApiListHistorySinceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryAPIService) ListHistorySinceExecute(r ApiListHistorySinceRequest) ([]*HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.ListHistorySince")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/history/since"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		localVarQueryParams.Add("date", parameterToString(*r.date, ""))
	}
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.includeArtist != nil {
		localVarQueryParams.Add("includeArtist", parameterToString(*r.includeArtist, ""))
	}
	if r.includeAlbum != nil {
		localVarQueryParams.Add("includeAlbum", parameterToString(*r.includeAlbum, ""))
	}
	if r.includeTrack != nil {
		localVarQueryParams.Add("includeTrack", parameterToString(*r.includeTrack, ""))
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
