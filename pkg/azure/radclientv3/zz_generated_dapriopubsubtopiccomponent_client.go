// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclientv3

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DaprIoPubSubTopicComponentClient contains the methods for the DaprIoPubSubTopicComponent group.
// Don't use this type directly, use NewDaprIoPubSubTopicComponentClient() instead.
type DaprIoPubSubTopicComponentClient struct {
	con *armcore.Connection
	subscriptionID string
}

// NewDaprIoPubSubTopicComponentClient creates a new instance of DaprIoPubSubTopicComponentClient with the specified values.
func NewDaprIoPubSubTopicComponentClient(con *armcore.Connection, subscriptionID string) *DaprIoPubSubTopicComponentClient {
	return &DaprIoPubSubTopicComponentClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a dapr.io.PubSubTopicComponent resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoPubSubTopicComponentClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, daprPubSubTopicComponentName string, parameters DaprPubSubTopicComponentResource, options *DaprIoPubSubTopicComponentCreateOrUpdateOptions) (DaprPubSubTopicComponentResourceResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationName, daprPubSubTopicComponentName, parameters, options)
	if err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return DaprPubSubTopicComponentResourceResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprIoPubSubTopicComponentClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, daprPubSubTopicComponentName string, parameters DaprPubSubTopicComponentResource, options *DaprIoPubSubTopicComponentCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.PubSubTopicComponent/{daprPubSubTopicComponentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if daprPubSubTopicComponentName == "" {
		return nil, errors.New("parameter daprPubSubTopicComponentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprPubSubTopicComponentName}", url.PathEscape(daprPubSubTopicComponentName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprIoPubSubTopicComponentClient) createOrUpdateHandleResponse(resp *azcore.Response) (DaprPubSubTopicComponentResourceResponse, error) {
	var val *DaprPubSubTopicComponentResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
return DaprPubSubTopicComponentResourceResponse{RawResponse: resp.Response, DaprPubSubTopicComponentResource: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DaprIoPubSubTopicComponentClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes a dapr.io.PubSubTopicComponent resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoPubSubTopicComponentClient) Delete(ctx context.Context, resourceGroupName string, applicationName string, daprPubSubTopicComponentName string, options *DaprIoPubSubTopicComponentDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationName, daprPubSubTopicComponentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprIoPubSubTopicComponentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, daprPubSubTopicComponentName string, options *DaprIoPubSubTopicComponentDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.PubSubTopicComponent/{daprPubSubTopicComponentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if daprPubSubTopicComponentName == "" {
		return nil, errors.New("parameter daprPubSubTopicComponentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprPubSubTopicComponentName}", url.PathEscape(daprPubSubTopicComponentName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DaprIoPubSubTopicComponentClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets a dapr.io.PubSubTopicComponent resource by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoPubSubTopicComponentClient) Get(ctx context.Context, resourceGroupName string, applicationName string, daprPubSubTopicComponentName string, options *DaprIoPubSubTopicComponentGetOptions) (DaprPubSubTopicComponentResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationName, daprPubSubTopicComponentName, options)
	if err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DaprPubSubTopicComponentResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprIoPubSubTopicComponentClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, daprPubSubTopicComponentName string, options *DaprIoPubSubTopicComponentGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.PubSubTopicComponent/{daprPubSubTopicComponentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if daprPubSubTopicComponentName == "" {
		return nil, errors.New("parameter daprPubSubTopicComponentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprPubSubTopicComponentName}", url.PathEscape(daprPubSubTopicComponentName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprIoPubSubTopicComponentClient) getHandleResponse(resp *azcore.Response) (DaprPubSubTopicComponentResourceResponse, error) {
	var val *DaprPubSubTopicComponentResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
return DaprPubSubTopicComponentResourceResponse{RawResponse: resp.Response, DaprPubSubTopicComponentResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *DaprIoPubSubTopicComponentClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - List the dapr.io.PubSubTopicComponent resources deployed in the application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoPubSubTopicComponentClient) List(ctx context.Context, resourceGroupName string, applicationName string, options *DaprIoPubSubTopicComponentListOptions) (DaprPubSubTopicComponentListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return DaprPubSubTopicComponentListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DaprPubSubTopicComponentListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DaprPubSubTopicComponentListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DaprIoPubSubTopicComponentClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *DaprIoPubSubTopicComponentListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.PubSubTopicComponent"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DaprIoPubSubTopicComponentClient) listHandleResponse(resp *azcore.Response) (DaprPubSubTopicComponentListResponse, error) {
	var val *DaprPubSubTopicComponentList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DaprPubSubTopicComponentListResponse{}, err
	}
return DaprPubSubTopicComponentListResponse{RawResponse: resp.Response, DaprPubSubTopicComponentList: val}, nil
}

// listHandleError handles the List error response.
func (client *DaprIoPubSubTopicComponentClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
