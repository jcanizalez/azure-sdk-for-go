// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// PrivateEndpointConnectionsOperations contains the methods for the PrivateEndpointConnections group.
type PrivateEndpointConnectionsOperations interface {
	// BeginDelete - Deletes the specified private endpoint connection associated with the key vault.
	BeginDelete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsDeleteOptions) (*PrivateEndpointConnectionPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (PrivateEndpointConnectionPoller, error)
	// Get - Gets the specified private endpoint connection associated with the key vault.
	Get(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsGetOptions) (*PrivateEndpointConnectionResponse, error)
	// Put - Updates the specified private endpoint connection associated with the key vault.
	Put(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsPutOptions) (*PrivateEndpointConnectionResponse, error)
}

// PrivateEndpointConnectionsClient implements the PrivateEndpointConnectionsOperations interface.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	*Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
func NewPrivateEndpointConnectionsClient(c *Client, subscriptionID string) PrivateEndpointConnectionsOperations {
	return &PrivateEndpointConnectionsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PrivateEndpointConnectionsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsDeleteOptions) (*PrivateEndpointConnectionPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	result := &PrivateEndpointConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateEndpointConnectionsClient.Delete", "", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &privateEndpointConnectionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PrivateEndpointConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *PrivateEndpointConnectionsClient) ResumeDelete(token string) (PrivateEndpointConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &privateEndpointConnectionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified private endpoint connection associated with the key vault.
func (client *PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-02-14")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *PrivateEndpointConnectionsClient) DeleteHandleResponse(resp *azcore.Response) (*PrivateEndpointConnectionResponse, error) {
	result := PrivateEndpointConnectionResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointConnection)
}

// DeleteHandleError handles the Delete error response.
func (client *PrivateEndpointConnectionsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified private endpoint connection associated with the key vault.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsGetOptions) (*PrivateEndpointConnectionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, options)
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
func (client *PrivateEndpointConnectionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-02-14")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) GetHandleResponse(resp *azcore.Response) (*PrivateEndpointConnectionResponse, error) {
	result := PrivateEndpointConnectionResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointConnection)
}

// GetHandleError handles the Get error response.
func (client *PrivateEndpointConnectionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put - Updates the specified private endpoint connection associated with the key vault.
func (client *PrivateEndpointConnectionsClient) Put(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsPutOptions) (*PrivateEndpointConnectionResponse, error) {
	req, err := client.PutCreateRequest(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutHandleError(resp)
	}
	result, err := client.PutHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutCreateRequest creates the Put request.
func (client *PrivateEndpointConnectionsClient) PutCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsPutOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-02-14")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(properties)
}

// PutHandleResponse handles the Put response.
func (client *PrivateEndpointConnectionsClient) PutHandleResponse(resp *azcore.Response) (*PrivateEndpointConnectionResponse, error) {
	result := PrivateEndpointConnectionResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return nil, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointConnection)
}

// PutHandleError handles the Put error response.
func (client *PrivateEndpointConnectionsClient) PutHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
