// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsCategoryOperations contains the methods for the DiagnosticSettingsCategory group.
type DiagnosticSettingsCategoryOperations interface {
	// Get - Gets the diagnostic settings category for the specified resource.
	Get(ctx context.Context, resourceUri string, name string, options *DiagnosticSettingsCategoryGetOptions) (*DiagnosticSettingsCategoryResourceResponse, error)
	// List - Lists the diagnostic settings categories for the specified resource.
	List(ctx context.Context, resourceUri string, options *DiagnosticSettingsCategoryListOptions) (*DiagnosticSettingsCategoryResourceCollectionResponse, error)
}

// DiagnosticSettingsCategoryClient implements the DiagnosticSettingsCategoryOperations interface.
// Don't use this type directly, use NewDiagnosticSettingsCategoryClient() instead.
type DiagnosticSettingsCategoryClient struct {
	*Client
}

// NewDiagnosticSettingsCategoryClient creates a new instance of DiagnosticSettingsCategoryClient with the specified values.
func NewDiagnosticSettingsCategoryClient(c *Client) DiagnosticSettingsCategoryOperations {
	return &DiagnosticSettingsCategoryClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *DiagnosticSettingsCategoryClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Gets the diagnostic settings category for the specified resource.
func (client *DiagnosticSettingsCategoryClient) Get(ctx context.Context, resourceUri string, name string, options *DiagnosticSettingsCategoryGetOptions) (*DiagnosticSettingsCategoryResourceResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceUri, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *DiagnosticSettingsCategoryClient) GetCreateRequest(ctx context.Context, resourceUri string, name string, options *DiagnosticSettingsCategoryGetOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-05-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *DiagnosticSettingsCategoryClient) GetHandleResponse(resp *azcore.Response) (*DiagnosticSettingsCategoryResourceResponse, error) {
	result := DiagnosticSettingsCategoryResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiagnosticSettingsCategoryResource)
}

// GetHandleError handles the Get error response.
func (client *DiagnosticSettingsCategoryClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists the diagnostic settings categories for the specified resource.
func (client *DiagnosticSettingsCategoryClient) List(ctx context.Context, resourceUri string, options *DiagnosticSettingsCategoryListOptions) (*DiagnosticSettingsCategoryResourceCollectionResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceUri, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *DiagnosticSettingsCategoryClient) ListCreateRequest(ctx context.Context, resourceUri string, options *DiagnosticSettingsCategoryListOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-05-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *DiagnosticSettingsCategoryClient) ListHandleResponse(resp *azcore.Response) (*DiagnosticSettingsCategoryResourceCollectionResponse, error) {
	result := DiagnosticSettingsCategoryResourceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiagnosticSettingsCategoryResourceCollection)
}

// ListHandleError handles the List error response.
func (client *DiagnosticSettingsCategoryClient) ListHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
