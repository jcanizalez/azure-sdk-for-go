// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteLinksOperations contains the methods for the ExpressRouteLinks group.
type ExpressRouteLinksOperations interface {
	// Get - Retrieves the specified ExpressRouteLink resource.
	Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (*ExpressRouteLinkResponse, error)
	// List - Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
	List(resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) ExpressRouteLinkListResultPager
}

// ExpressRouteLinksClient implements the ExpressRouteLinksOperations interface.
// Don't use this type directly, use NewExpressRouteLinksClient() instead.
type ExpressRouteLinksClient struct {
	*Client
	subscriptionID string
}

// NewExpressRouteLinksClient creates a new instance of ExpressRouteLinksClient with the specified values.
func NewExpressRouteLinksClient(c *Client, subscriptionID string) ExpressRouteLinksOperations {
	return &ExpressRouteLinksClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ExpressRouteLinksClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Retrieves the specified ExpressRouteLink resource.
func (client *ExpressRouteLinksClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (*ExpressRouteLinkResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, expressRoutePortName, linkName, options)
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
func (client *ExpressRouteLinksClient) GetCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *ExpressRouteLinksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	urlPath = strings.ReplaceAll(urlPath, "{linkName}", url.PathEscape(linkName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *ExpressRouteLinksClient) GetHandleResponse(resp *azcore.Response) (*ExpressRouteLinkResponse, error) {
	result := ExpressRouteLinkResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteLink)
}

// GetHandleError handles the Get error response.
func (client *ExpressRouteLinksClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
func (client *ExpressRouteLinksClient) List(resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) ExpressRouteLinkListResultPager {
	return &expressRouteLinkListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteLinkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteLinkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ExpressRouteLinksClient) ListCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRouteLinksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *ExpressRouteLinksClient) ListHandleResponse(resp *azcore.Response) (*ExpressRouteLinkListResultResponse, error) {
	result := ExpressRouteLinkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteLinkListResult)
}

// ListHandleError handles the List error response.
func (client *ExpressRouteLinksClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
