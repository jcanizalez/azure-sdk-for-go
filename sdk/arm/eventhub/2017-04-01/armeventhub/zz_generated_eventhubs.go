// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EventHubsOperations contains the methods for the EventHubs group.
type EventHubsOperations interface {
	// CreateOrUpdate - Creates or updates a new Event Hub as a nested resource within a Namespace.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters Eventhub, options *EventHubsCreateOrUpdateOptions) (*EventhubResponse, error)
	// CreateOrUpdateAuthorizationRule - Creates or updates an AuthorizationRule for the specified Event Hub. Creation/update of the AuthorizationRule will take a few seconds to take effect.
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters AuthorizationRule, options *EventHubsCreateOrUpdateAuthorizationRuleOptions) (*AuthorizationRuleResponse, error)
	// Delete - Deletes an Event Hub from the specified Namespace and resource group.
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsDeleteOptions) (*http.Response, error)
	// DeleteAuthorizationRule - Deletes an Event Hub AuthorizationRule.
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsDeleteAuthorizationRuleOptions) (*http.Response, error)
	// Get - Gets an Event Hubs description for the specified Event Hub.
	Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsGetOptions) (*EventhubResponse, error)
	// GetAuthorizationRule - Gets an AuthorizationRule for an Event Hub by rule name.
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsGetAuthorizationRuleOptions) (*AuthorizationRuleResponse, error)
	// ListAuthorizationRules - Gets the authorization rules for an Event Hub.
	ListAuthorizationRules(resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsListAuthorizationRulesOptions) AuthorizationRuleListResultPager
	// ListByNamespace - Gets all the Event Hubs in a Namespace.
	ListByNamespace(resourceGroupName string, namespaceName string, options *EventHubsListByNamespaceOptions) EventHubListResultPager
	// ListKeys - Gets the ACS and SAS connection strings for the Event Hub.
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsListKeysOptions) (*AccessKeysResponse, error)
	// RegenerateKeys - Regenerates the ACS and SAS connection strings for the Event Hub.
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *EventHubsRegenerateKeysOptions) (*AccessKeysResponse, error)
}

// EventHubsClient implements the EventHubsOperations interface.
// Don't use this type directly, use NewEventHubsClient() instead.
type EventHubsClient struct {
	*Client
	subscriptionID string
}

// NewEventHubsClient creates a new instance of EventHubsClient with the specified values.
func NewEventHubsClient(c *Client, subscriptionID string) EventHubsOperations {
	return &EventHubsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *EventHubsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates a new Event Hub as a nested resource within a Namespace.
func (client *EventHubsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters Eventhub, options *EventHubsCreateOrUpdateOptions) (*EventhubResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EventHubsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters Eventhub, options *EventHubsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EventHubsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*EventhubResponse, error) {
	result := EventhubResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Eventhub)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *EventHubsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateOrUpdateAuthorizationRule - Creates or updates an AuthorizationRule for the specified Event Hub. Creation/update of the AuthorizationRule will take a few seconds to take effect.
func (client *EventHubsClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters AuthorizationRule, options *EventHubsCreateOrUpdateAuthorizationRuleOptions) (*AuthorizationRuleResponse, error) {
	req, err := client.CreateOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateAuthorizationRuleHandleError(resp)
	}
	result, err := client.CreateOrUpdateAuthorizationRuleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *EventHubsClient) CreateOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters AuthorizationRule, options *EventHubsCreateOrUpdateAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *EventHubsClient) CreateOrUpdateAuthorizationRuleHandleResponse(resp *azcore.Response) (*AuthorizationRuleResponse, error) {
	result := AuthorizationRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AuthorizationRule)
}

// CreateOrUpdateAuthorizationRuleHandleError handles the CreateOrUpdateAuthorizationRule error response.
func (client *EventHubsClient) CreateOrUpdateAuthorizationRuleHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - Deletes an Event Hub from the specified Namespace and resource group.
func (client *EventHubsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsDeleteOptions) (*http.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *EventHubsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *EventHubsClient) DeleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteAuthorizationRule - Deletes an Event Hub AuthorizationRule.
func (client *EventHubsClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsDeleteAuthorizationRuleOptions) (*http.Response, error) {
	req, err := client.DeleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteAuthorizationRuleHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *EventHubsClient) DeleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsDeleteAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteAuthorizationRuleHandleError handles the DeleteAuthorizationRule error response.
func (client *EventHubsClient) DeleteAuthorizationRuleHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets an Event Hubs description for the specified Event Hub.
func (client *EventHubsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsGetOptions) (*EventhubResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
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
func (client *EventHubsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *EventHubsClient) GetHandleResponse(resp *azcore.Response) (*EventhubResponse, error) {
	result := EventhubResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Eventhub)
}

// GetHandleError handles the Get error response.
func (client *EventHubsClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetAuthorizationRule - Gets an AuthorizationRule for an Event Hub by rule name.
func (client *EventHubsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsGetAuthorizationRuleOptions) (*AuthorizationRuleResponse, error) {
	req, err := client.GetAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetAuthorizationRuleHandleError(resp)
	}
	result, err := client.GetAuthorizationRuleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *EventHubsClient) GetAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsGetAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *EventHubsClient) GetAuthorizationRuleHandleResponse(resp *azcore.Response) (*AuthorizationRuleResponse, error) {
	result := AuthorizationRuleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AuthorizationRule)
}

// GetAuthorizationRuleHandleError handles the GetAuthorizationRule error response.
func (client *EventHubsClient) GetAuthorizationRuleHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAuthorizationRules - Gets the authorization rules for an Event Hub.
func (client *EventHubsClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsListAuthorizationRulesOptions) AuthorizationRuleListResultPager {
	return &authorizationRuleListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
		},
		responder: client.ListAuthorizationRulesHandleResponse,
		errorer:   client.ListAuthorizationRulesHandleError,
		advancer: func(ctx context.Context, resp *AuthorizationRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AuthorizationRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *EventHubsClient) ListAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsListAuthorizationRulesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *EventHubsClient) ListAuthorizationRulesHandleResponse(resp *azcore.Response) (*AuthorizationRuleListResultResponse, error) {
	result := AuthorizationRuleListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AuthorizationRuleListResult)
}

// ListAuthorizationRulesHandleError handles the ListAuthorizationRules error response.
func (client *EventHubsClient) ListAuthorizationRulesHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByNamespace - Gets all the Event Hubs in a Namespace.
func (client *EventHubsClient) ListByNamespace(resourceGroupName string, namespaceName string, options *EventHubsListByNamespaceOptions) EventHubListResultPager {
	return &eventHubListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		responder: client.ListByNamespaceHandleResponse,
		errorer:   client.ListByNamespaceHandleError,
		advancer: func(ctx context.Context, resp *EventHubListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EventHubListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByNamespaceCreateRequest creates the ListByNamespace request.
func (client *EventHubsClient) ListByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *EventHubsListByNamespaceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	if options != nil && options.Skip != nil {
		query.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByNamespaceHandleResponse handles the ListByNamespace response.
func (client *EventHubsClient) ListByNamespaceHandleResponse(resp *azcore.Response) (*EventHubListResultResponse, error) {
	result := EventHubListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.EventHubListResult)
}

// ListByNamespaceHandleError handles the ListByNamespace error response.
func (client *EventHubsClient) ListByNamespaceHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListKeys - Gets the ACS and SAS connection strings for the Event Hub.
func (client *EventHubsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsListKeysOptions) (*AccessKeysResponse, error) {
	req, err := client.ListKeysCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListKeysHandleError(resp)
	}
	result, err := client.ListKeysHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListKeysCreateRequest creates the ListKeys request.
func (client *EventHubsClient) ListKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsListKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}/listKeys"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListKeysHandleResponse handles the ListKeys response.
func (client *EventHubsClient) ListKeysHandleResponse(resp *azcore.Response) (*AccessKeysResponse, error) {
	result := AccessKeysResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AccessKeys)
}

// ListKeysHandleError handles the ListKeys error response.
func (client *EventHubsClient) ListKeysHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RegenerateKeys - Regenerates the ACS and SAS connection strings for the Event Hub.
func (client *EventHubsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *EventHubsRegenerateKeysOptions) (*AccessKeysResponse, error) {
	req, err := client.RegenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.RegenerateKeysHandleError(resp)
	}
	result, err := client.RegenerateKeysHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RegenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *EventHubsClient) RegenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *EventHubsRegenerateKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2017-04-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// RegenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *EventHubsClient) RegenerateKeysHandleResponse(resp *azcore.Response) (*AccessKeysResponse, error) {
	result := AccessKeysResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AccessKeys)
}

// RegenerateKeysHandleError handles the RegenerateKeys error response.
func (client *EventHubsClient) RegenerateKeysHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
