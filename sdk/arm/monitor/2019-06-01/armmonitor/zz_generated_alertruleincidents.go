// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AlertRuleIncidentsOperations contains the methods for the AlertRuleIncidents group.
type AlertRuleIncidentsOperations interface {
	// Get - Gets an incident associated to an alert rule
	Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string, options *AlertRuleIncidentsGetOptions) (*IncidentResponse, error)
	// ListByAlertRule - Gets a list of incidents associated to an alert rule
	ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRuleIncidentsListByAlertRuleOptions) (*IncidentListResultResponse, error)
}

// AlertRuleIncidentsClient implements the AlertRuleIncidentsOperations interface.
// Don't use this type directly, use NewAlertRuleIncidentsClient() instead.
type AlertRuleIncidentsClient struct {
	*Client
	subscriptionID string
}

// NewAlertRuleIncidentsClient creates a new instance of AlertRuleIncidentsClient with the specified values.
func NewAlertRuleIncidentsClient(c *Client, subscriptionID string) AlertRuleIncidentsOperations {
	return &AlertRuleIncidentsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AlertRuleIncidentsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Gets an incident associated to an alert rule
func (client *AlertRuleIncidentsClient) Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string, options *AlertRuleIncidentsGetOptions) (*IncidentResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, ruleName, incidentName, options)
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
func (client *AlertRuleIncidentsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, incidentName string, options *AlertRuleIncidentsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents/{incidentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{incidentName}", url.PathEscape(incidentName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *AlertRuleIncidentsClient) GetHandleResponse(resp *azcore.Response) (*IncidentResponse, error) {
	result := IncidentResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Incident)
}

// GetHandleError handles the Get error response.
func (client *AlertRuleIncidentsClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByAlertRule - Gets a list of incidents associated to an alert rule
func (client *AlertRuleIncidentsClient) ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRuleIncidentsListByAlertRuleOptions) (*IncidentListResultResponse, error) {
	req, err := client.ListByAlertRuleCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListByAlertRuleHandleError(resp)
	}
	result, err := client.ListByAlertRuleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListByAlertRuleCreateRequest creates the ListByAlertRule request.
func (client *AlertRuleIncidentsClient) ListByAlertRuleCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRuleIncidentsListByAlertRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByAlertRuleHandleResponse handles the ListByAlertRule response.
func (client *AlertRuleIncidentsClient) ListByAlertRuleHandleResponse(resp *azcore.Response) (*IncidentListResultResponse, error) {
	result := IncidentListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.IncidentListResult)
}

// ListByAlertRuleHandleError handles the ListByAlertRule error response.
func (client *AlertRuleIncidentsClient) ListByAlertRuleHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
