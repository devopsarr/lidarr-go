/*
Lidarr

Lidarr API docs

API version: v2.8.2.4493
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
)


// CutoffAPIService CutoffAPI service
type CutoffAPIService service

type ApiGetWantedCutoffRequest struct {
	ctx context.Context
	ApiService *CutoffAPIService
	page *int32
	pageSize *int32
	sortKey *string
	sortDirection *SortDirection
	includeArtist *bool
	monitored *bool
}

func (r ApiGetWantedCutoffRequest) Page(page int32) ApiGetWantedCutoffRequest {
	r.page = &page
	return r
}

func (r ApiGetWantedCutoffRequest) PageSize(pageSize int32) ApiGetWantedCutoffRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetWantedCutoffRequest) SortKey(sortKey string) ApiGetWantedCutoffRequest {
	r.sortKey = &sortKey
	return r
}

func (r ApiGetWantedCutoffRequest) SortDirection(sortDirection SortDirection) ApiGetWantedCutoffRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiGetWantedCutoffRequest) IncludeArtist(includeArtist bool) ApiGetWantedCutoffRequest {
	r.includeArtist = &includeArtist
	return r
}

func (r ApiGetWantedCutoffRequest) Monitored(monitored bool) ApiGetWantedCutoffRequest {
	r.monitored = &monitored
	return r
}

func (r ApiGetWantedCutoffRequest) Execute() (*AlbumResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetWantedCutoffExecute(r)
}

/*
GetWantedCutoff Method for GetWantedCutoff

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWantedCutoffRequest
*/
func (a *CutoffAPIService) GetWantedCutoff(ctx context.Context) ApiGetWantedCutoffRequest {
	return ApiGetWantedCutoffRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlbumResourcePagingResource
func (a *CutoffAPIService) GetWantedCutoffExecute(r ApiGetWantedCutoffRequest) (*AlbumResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlbumResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CutoffAPIService.GetWantedCutoff")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/wanted/cutoff"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 10
		r.pageSize = &defaultValue
	}
	if r.sortKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortKey", r.sortKey, "form", "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortDirection", r.sortDirection, "form", "")
	}
	if r.includeArtist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeArtist", r.includeArtist, "form", "")
	} else {
		var defaultValue bool = false
		r.includeArtist = &defaultValue
	}
	if r.monitored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "monitored", r.monitored, "form", "")
	} else {
		var defaultValue bool = true
		r.monitored = &defaultValue
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

type ApiGetWantedCutoffByIdRequest struct {
	ctx context.Context
	ApiService *CutoffAPIService
	id int32
}

func (r ApiGetWantedCutoffByIdRequest) Execute() (*AlbumResource, *http.Response, error) {
	return r.ApiService.GetWantedCutoffByIdExecute(r)
}

/*
GetWantedCutoffById Method for GetWantedCutoffById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetWantedCutoffByIdRequest
*/
func (a *CutoffAPIService) GetWantedCutoffById(ctx context.Context, id int32) ApiGetWantedCutoffByIdRequest {
	return ApiGetWantedCutoffByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AlbumResource
func (a *CutoffAPIService) GetWantedCutoffByIdExecute(r ApiGetWantedCutoffByIdRequest) (*AlbumResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlbumResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CutoffAPIService.GetWantedCutoffById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/wanted/cutoff/{id}"
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
