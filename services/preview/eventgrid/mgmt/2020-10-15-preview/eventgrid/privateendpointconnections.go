package eventgrid

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateEndpointConnectionsClient is the azure EventGrid Management Client
type PrivateEndpointConnectionsClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionsClient creates an instance of the PrivateEndpointConnectionsClient client.
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return NewPrivateEndpointConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionsClientWithBaseURI creates an instance of the PrivateEndpointConnectionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return PrivateEndpointConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete a specific private endpoint connection under a topic or domain.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// parentType - the type of the parent resource. This can be either \'topics\' or \'domains\'.
// parentName - the name of the parent resource (namely, either, the topic name or domain name).
// privateEndpointConnectionName - the name of the private endpoint connection connection.
func (client PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string) (result PrivateEndpointConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateEndpointConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentName":                    autorest.Encode("path", parentName),
		"parentType":                    autorest.Encode("path", parentType),
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) DeleteSender(req *http.Request) (future PrivateEndpointConnectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a specific private endpoint connection under a topic or domain.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// parentType - the type of the parent resource. This can be either \'topics\' or \'domains\'.
// parentName - the name of the parent resource (namely, either, the topic name or domain name).
// privateEndpointConnectionName - the name of the private endpoint connection connection.
func (client PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentName":                    autorest.Encode("path", parentName),
		"parentType":                    autorest.Encode("path", parentType),
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) GetResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResource get all private endpoint connections under a topic or domain.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// parentType - the type of the parent resource. This can be either \'topics\' or \'domains\'.
// parentName - the name of the parent resource (namely, either, the topic name or domain name).
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PrivateEndpointConnectionsClient) ListByResource(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result PrivateEndpointConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.ListByResource")
		defer func() {
			sc := -1
			if result.peclr.Response.Response != nil {
				sc = result.peclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceNextResults
	req, err := client.ListByResourcePreparer(ctx, resourceGroupName, parentType, parentName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "ListByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.peclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "ListByResource", resp, "Failure sending request")
		return
	}

	result.peclr, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "ListByResource", resp, "Failure responding to request")
		return
	}
	if result.peclr.hasNextLink() && result.peclr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourcePreparer prepares the ListByResource request.
func (client PrivateEndpointConnectionsClient) ListByResourcePreparer(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentName":        autorest.Encode("path", parentName),
		"parentType":        autorest.Encode("path", parentType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceSender sends the ListByResource request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) ListByResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceResponder handles the response to the ListByResource request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) ListByResourceResponder(resp *http.Response) (result PrivateEndpointConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceNextResults retrieves the next set of results, if any.
func (client PrivateEndpointConnectionsClient) listByResourceNextResults(ctx context.Context, lastResults PrivateEndpointConnectionListResult) (result PrivateEndpointConnectionListResult, err error) {
	req, err := lastResults.privateEndpointConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "listByResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "listByResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "listByResourceNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndpointConnectionsClient) ListByResourceComplete(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result PrivateEndpointConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.ListByResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResource(ctx, resourceGroupName, parentType, parentName, filter, top)
	return
}

// Update update a specific private endpoint connection under a topic or domain.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// parentType - the type of the parent resource. This can be either \'topics\' or \'domains\'.
// parentName - the name of the parent resource (namely, either, the topic name or domain name).
// privateEndpointConnectionName - the name of the private endpoint connection connection.
// privateEndpointConnection - the private endpoint connection object to update.
func (client PrivateEndpointConnectionsClient) Update(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection) (result PrivateEndpointConnectionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName, privateEndpointConnection)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateEndpointConnectionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PrivateEndpointConnectionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentName":                    autorest.Encode("path", parentName),
		"parentType":                    autorest.Encode("path", parentType),
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithJSON(privateEndpointConnection),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) UpdateSender(req *http.Request) (future PrivateEndpointConnectionsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) UpdateResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
