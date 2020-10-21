// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ConfigurationStoresOperations contains the methods for the ConfigurationStores group.
type ConfigurationStoresOperations interface {
	// BeginCreate - Creates a configuration store with the specified parameters.
	BeginCreate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresCreateOptions) (*ConfigurationStorePollerResponse, error)
	// ResumeCreate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreate(token string) (ConfigurationStorePoller, error)
	// BeginDelete - Deletes a configuration store.
	BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the properties of the specified configuration store.
	Get(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresGetOptions) (*ConfigurationStoreResponse, error)
	// List - Lists the configuration stores for a given subscription.
	List(options *ConfigurationStoresListOptions) ConfigurationStoreListResultPager
	// ListByResourceGroup - Lists the configuration stores for a given resource group.
	ListByResourceGroup(resourceGroupName string, options *ConfigurationStoresListByResourceGroupOptions) ConfigurationStoreListResultPager
	// ListKeyValue - Lists a configuration store key-value.
	ListKeyValue(ctx context.Context, resourceGroupName string, configStoreName string, listKeyValueParameters ListKeyValueParameters, options *ConfigurationStoresListKeyValueOptions) (*KeyValueResponse, error)
	// ListKeys - Lists the access key for the specified configuration store.
	ListKeys(resourceGroupName string, configStoreName string, options *ConfigurationStoresListKeysOptions) APIKeyListResultPager
	// RegenerateKey - Regenerates an access key for the specified configuration store.
	RegenerateKey(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresRegenerateKeyOptions) (*APIKeyResponse, error)
	// BeginUpdate - Updates a configuration store with the specified parameters.
	BeginUpdate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresUpdateOptions) (*ConfigurationStorePollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (ConfigurationStorePoller, error)
}

// ConfigurationStoresClient implements the ConfigurationStoresOperations interface.
// Don't use this type directly, use NewConfigurationStoresClient() instead.
type ConfigurationStoresClient struct {
	*Client
	subscriptionID string
}

// NewConfigurationStoresClient creates a new instance of ConfigurationStoresClient with the specified values.
func NewConfigurationStoresClient(c *Client, subscriptionID string) ConfigurationStoresOperations {
	return &ConfigurationStoresClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ConfigurationStoresClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *ConfigurationStoresClient) BeginCreate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresCreateOptions) (*ConfigurationStorePollerResponse, error) {
	resp, err := client.Create(ctx, resourceGroupName, configStoreName, configStoreCreationParameters, options)
	if err != nil {
		return nil, err
	}
	result := &ConfigurationStorePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConfigurationStoresClient.Create", "", resp, client.CreateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &configurationStorePoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConfigurationStoreResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ConfigurationStoresClient) ResumeCreate(token string) (ConfigurationStorePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConfigurationStoresClient.Create", token, client.CreateHandleError)
	if err != nil {
		return nil, err
	}
	return &configurationStorePoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Create - Creates a configuration store with the specified parameters.
func (client *ConfigurationStoresClient) Create(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresCreateOptions) (*azcore.Response, error) {
	req, err := client.CreateCreateRequest(ctx, resourceGroupName, configStoreName, configStoreCreationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateHandleError(resp)
	}
	return resp, nil
}

// CreateCreateRequest creates the Create request.
func (client *ConfigurationStoresClient) CreateCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(configStoreCreationParameters)
}

// CreateHandleResponse handles the Create response.
func (client *ConfigurationStoresClient) CreateHandleResponse(resp *azcore.Response) (*ConfigurationStoreResponse, error) {
	result := ConfigurationStoreResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConfigurationStore)
}

// CreateHandleError handles the Create error response.
func (client *ConfigurationStoresClient) CreateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ConfigurationStoresClient) BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, configStoreName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConfigurationStoresClient.Delete", "", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ConfigurationStoresClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConfigurationStoresClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes a configuration store.
func (client *ConfigurationStoresClient) Delete(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, configStoreName, options)
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
func (client *ConfigurationStoresClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *ConfigurationStoresClient) DeleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the properties of the specified configuration store.
func (client *ConfigurationStoresClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresGetOptions) (*ConfigurationStoreResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, configStoreName, options)
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
func (client *ConfigurationStoresClient) GetCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *ConfigurationStoresClient) GetHandleResponse(resp *azcore.Response) (*ConfigurationStoreResponse, error) {
	result := ConfigurationStoreResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConfigurationStore)
}

// GetHandleError handles the Get error response.
func (client *ConfigurationStoresClient) GetHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists the configuration stores for a given subscription.
func (client *ConfigurationStoresClient) List(options *ConfigurationStoresListOptions) ConfigurationStoreListResultPager {
	return &configurationStoreListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ConfigurationStoreListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ConfigurationStoreListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ConfigurationStoresClient) ListCreateRequest(ctx context.Context, options *ConfigurationStoresListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	if options != nil && options.SkipToken != nil {
		query.Set("$skipToken", *options.SkipToken)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *ConfigurationStoresClient) ListHandleResponse(resp *azcore.Response) (*ConfigurationStoreListResultResponse, error) {
	result := ConfigurationStoreListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConfigurationStoreListResult)
}

// ListHandleError handles the List error response.
func (client *ConfigurationStoresClient) ListHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists the configuration stores for a given resource group.
func (client *ConfigurationStoresClient) ListByResourceGroup(resourceGroupName string, options *ConfigurationStoresListByResourceGroupOptions) ConfigurationStoreListResultPager {
	return &configurationStoreListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ConfigurationStoreListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ConfigurationStoreListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationStoresClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationStoresListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	if options != nil && options.SkipToken != nil {
		query.Set("$skipToken", *options.SkipToken)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationStoresClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ConfigurationStoreListResultResponse, error) {
	result := ConfigurationStoreListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConfigurationStoreListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ConfigurationStoresClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListKeyValue - Lists a configuration store key-value.
func (client *ConfigurationStoresClient) ListKeyValue(ctx context.Context, resourceGroupName string, configStoreName string, listKeyValueParameters ListKeyValueParameters, options *ConfigurationStoresListKeyValueOptions) (*KeyValueResponse, error) {
	req, err := client.ListKeyValueCreateRequest(ctx, resourceGroupName, configStoreName, listKeyValueParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListKeyValueHandleError(resp)
	}
	result, err := client.ListKeyValueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListKeyValueCreateRequest creates the ListKeyValue request.
func (client *ConfigurationStoresClient) ListKeyValueCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, listKeyValueParameters ListKeyValueParameters, options *ConfigurationStoresListKeyValueOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/listKeyValue"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(listKeyValueParameters)
}

// ListKeyValueHandleResponse handles the ListKeyValue response.
func (client *ConfigurationStoresClient) ListKeyValueHandleResponse(resp *azcore.Response) (*KeyValueResponse, error) {
	result := KeyValueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.KeyValue)
}

// ListKeyValueHandleError handles the ListKeyValue error response.
func (client *ConfigurationStoresClient) ListKeyValueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListKeys - Lists the access key for the specified configuration store.
func (client *ConfigurationStoresClient) ListKeys(resourceGroupName string, configStoreName string, options *ConfigurationStoresListKeysOptions) APIKeyListResultPager {
	return &apiKeyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListKeysCreateRequest(ctx, resourceGroupName, configStoreName, options)
		},
		responder: client.ListKeysHandleResponse,
		errorer:   client.ListKeysHandleError,
		advancer: func(ctx context.Context, resp *APIKeyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.APIKeyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListKeysCreateRequest creates the ListKeys request.
func (client *ConfigurationStoresClient) ListKeysCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresListKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/ListKeys"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	if options != nil && options.SkipToken != nil {
		query.Set("$skipToken", *options.SkipToken)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListKeysHandleResponse handles the ListKeys response.
func (client *ConfigurationStoresClient) ListKeysHandleResponse(resp *azcore.Response) (*APIKeyListResultResponse, error) {
	result := APIKeyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.APIKeyListResult)
}

// ListKeysHandleError handles the ListKeys error response.
func (client *ConfigurationStoresClient) ListKeysHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RegenerateKey - Regenerates an access key for the specified configuration store.
func (client *ConfigurationStoresClient) RegenerateKey(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresRegenerateKeyOptions) (*APIKeyResponse, error) {
	req, err := client.RegenerateKeyCreateRequest(ctx, resourceGroupName, configStoreName, regenerateKeyParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.RegenerateKeyHandleError(resp)
	}
	result, err := client.RegenerateKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RegenerateKeyCreateRequest creates the RegenerateKey request.
func (client *ConfigurationStoresClient) RegenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresRegenerateKeyOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/RegenerateKey"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(regenerateKeyParameters)
}

// RegenerateKeyHandleResponse handles the RegenerateKey response.
func (client *ConfigurationStoresClient) RegenerateKeyHandleResponse(resp *azcore.Response) (*APIKeyResponse, error) {
	result := APIKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.APIKey)
}

// RegenerateKeyHandleError handles the RegenerateKey error response.
func (client *ConfigurationStoresClient) RegenerateKeyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *ConfigurationStoresClient) BeginUpdate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresUpdateOptions) (*ConfigurationStorePollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	result := &ConfigurationStorePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConfigurationStoresClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &configurationStorePoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConfigurationStoreResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ConfigurationStoresClient) ResumeUpdate(token string) (ConfigurationStorePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConfigurationStoresClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &configurationStorePoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Update - Updates a configuration store with the specified parameters.
func (client *ConfigurationStoresClient) Update(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *ConfigurationStoresClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-10-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(configStoreUpdateParameters)
}

// UpdateHandleResponse handles the Update response.
func (client *ConfigurationStoresClient) UpdateHandleResponse(resp *azcore.Response) (*ConfigurationStoreResponse, error) {
	result := ConfigurationStoreResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConfigurationStore)
}

// UpdateHandleError handles the Update error response.
func (client *ConfigurationStoresClient) UpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
