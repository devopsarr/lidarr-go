/*
Lidarr

Lidarr API docs

API version: v2.2.5.4141
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
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	page *int32
	pageSize *int32
	sortKey *string
	sortDirection *SortDirection
	includeArtist *bool
	includeAlbum *bool
	includeTrack *bool
	eventType *[]int32
	albumId *int32
	downloadId *string
	artistIds *[]int32
	quality *[]int32
}

func (r ApiGetHistoryRequest) Page(page int32) ApiGetHistoryRequest {
	r.page = &page
	return r
}

func (r ApiGetHistoryRequest) PageSize(pageSize int32) ApiGetHistoryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetHistoryRequest) SortKey(sortKey string) ApiGetHistoryRequest {
	r.sortKey = &sortKey
	return r
}

func (r ApiGetHistoryRequest) SortDirection(sortDirection SortDirection) ApiGetHistoryRequest {
	r.sortDirection = &sortDirection
	return r
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

func (r ApiGetHistoryRequest) EventType(eventType []int32) ApiGetHistoryRequest {
	r.eventType = &eventType
	return r
}

func (r ApiGetHistoryRequest) AlbumId(albumId int32) ApiGetHistoryRequest {
	r.albumId = &albumId
	return r
}

func (r ApiGetHistoryRequest) DownloadId(downloadId string) ApiGetHistoryRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetHistoryRequest) ArtistIds(artistIds []int32) ApiGetHistoryRequest {
	r.artistIds = &artistIds
	return r
}

func (r ApiGetHistoryRequest) Quality(quality []int32) ApiGetHistoryRequest {
	r.quality = &quality
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

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 10
		r.pageSize = &defaultValue
	}
	if r.sortKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortKey", r.sortKey, "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortDirection", r.sortDirection, "")
	}
	if r.includeArtist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeArtist", r.includeArtist, "")
	}
	if r.includeAlbum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAlbum", r.includeAlbum, "")
	}
	if r.includeTrack != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeTrack", r.includeTrack, "")
	}
	if r.eventType != nil {
		t := *r.eventType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "multi")
		}
	}
	if r.albumId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "albumId", r.albumId, "")
	}
	if r.downloadId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "")
	}
	if r.artistIds != nil {
		t := *r.artistIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "artistIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "artistIds", t, "multi")
		}
	}
	if r.quality != nil {
		t := *r.quality
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "quality", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "quality", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

func (r ApiListHistoryArtistRequest) Execute() ([]HistoryResource, *http.Response, error) {
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
func (a *HistoryAPIService) ListHistoryArtistExecute(r ApiListHistoryArtistRequest) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoryResource
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
		parameterAddToHeaderOrQuery(localVarQueryParams, "artistId", r.artistId, "")
	}
	if r.albumId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "albumId", r.albumId, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.includeArtist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeArtist", r.includeArtist, "")
	} else {
		var defaultValue bool = false
		r.includeArtist = &defaultValue
	}
	if r.includeAlbum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAlbum", r.includeAlbum, "")
	} else {
		var defaultValue bool = false
		r.includeAlbum = &defaultValue
	}
	if r.includeTrack != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeTrack", r.includeTrack, "")
	} else {
		var defaultValue bool = false
		r.includeTrack = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

func (r ApiListHistorySinceRequest) Execute() ([]HistoryResource, *http.Response, error) {
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
func (a *HistoryAPIService) ListHistorySinceExecute(r ApiListHistorySinceRequest) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoryResource
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
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.includeArtist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeArtist", r.includeArtist, "")
	} else {
		var defaultValue bool = false
		r.includeArtist = &defaultValue
	}
	if r.includeAlbum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAlbum", r.includeAlbum, "")
	} else {
		var defaultValue bool = false
		r.includeAlbum = &defaultValue
	}
	if r.includeTrack != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeTrack", r.includeTrack, "")
	} else {
		var defaultValue bool = false
		r.includeTrack = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
