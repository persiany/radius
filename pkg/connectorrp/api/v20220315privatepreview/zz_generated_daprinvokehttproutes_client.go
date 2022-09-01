//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DaprInvokeHTTPRoutesClient contains the methods for the DaprInvokeHTTPRoutes group.
// Don't use this type directly, use NewDaprInvokeHTTPRoutesClient() instead.
type DaprInvokeHTTPRoutesClient struct {
	host string
	rootScope string
	pl runtime.Pipeline
}

// NewDaprInvokeHTTPRoutesClient creates a new instance of DaprInvokeHTTPRoutesClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDaprInvokeHTTPRoutesClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprInvokeHTTPRoutesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DaprInvokeHTTPRoutesClient{
		rootScope: rootScope,
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a DaprInvokeHttpRoute resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprInvokeHTTPRouteName - The name of the DaprInvokeHttpRoute connector resource
// daprInvokeHTTPRouteParameters - daprInvokeHttpRoute create parameters
// options - DaprInvokeHTTPRoutesClientCreateOrUpdateOptions contains the optional parameters for the DaprInvokeHTTPRoutesClient.CreateOrUpdate
// method.
func (client *DaprInvokeHTTPRoutesClient) CreateOrUpdate(ctx context.Context, daprInvokeHTTPRouteName string, daprInvokeHTTPRouteParameters DaprInvokeHTTPRouteResource, options *DaprInvokeHTTPRoutesClientCreateOrUpdateOptions) (DaprInvokeHTTPRoutesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, daprInvokeHTTPRouteName, daprInvokeHTTPRouteParameters, options)
	if err != nil {
		return DaprInvokeHTTPRoutesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRoutesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DaprInvokeHTTPRoutesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprInvokeHTTPRoutesClient) createOrUpdateCreateRequest(ctx context.Context, daprInvokeHTTPRouteName string, daprInvokeHTTPRouteParameters DaprInvokeHTTPRouteResource, options *DaprInvokeHTTPRoutesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, daprInvokeHTTPRouteParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprInvokeHTTPRoutesClient) createOrUpdateHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesClientCreateOrUpdateResponse, error) {
	result := DaprInvokeHTTPRoutesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResource); err != nil {
		return DaprInvokeHTTPRoutesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing daprInvokeHttpRoute resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprInvokeHTTPRouteName - The name of the DaprInvokeHttpRoute connector resource
// options - DaprInvokeHTTPRoutesClientDeleteOptions contains the optional parameters for the DaprInvokeHTTPRoutesClient.Delete
// method.
func (client *DaprInvokeHTTPRoutesClient) Delete(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesClientDeleteOptions) (DaprInvokeHTTPRoutesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, daprInvokeHTTPRouteName, options)
	if err != nil {
		return DaprInvokeHTTPRoutesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRoutesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return DaprInvokeHTTPRoutesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DaprInvokeHTTPRoutesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprInvokeHTTPRoutesClient) deleteCreateRequest(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a daprInvokeHttpRoute resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprInvokeHTTPRouteName - The name of the DaprInvokeHttpRoute connector resource
// options - DaprInvokeHTTPRoutesClientGetOptions contains the optional parameters for the DaprInvokeHTTPRoutesClient.Get
// method.
func (client *DaprInvokeHTTPRoutesClient) Get(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesClientGetOptions) (DaprInvokeHTTPRoutesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, daprInvokeHTTPRouteName, options)
	if err != nil {
		return DaprInvokeHTTPRoutesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRoutesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprInvokeHTTPRoutesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprInvokeHTTPRoutesClient) getCreateRequest(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRoutesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprInvokeHTTPRoutesClient) getHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesClientGetResponse, error) {
	result := DaprInvokeHTTPRoutesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResource); err != nil {
		return DaprInvokeHTTPRoutesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRootScopePager - Lists information about all daprInvokeHttpRoute resources in the given root scope
// Generated from API version 2022-03-15-privatepreview
// options - DaprInvokeHTTPRoutesClientListByRootScopeOptions contains the optional parameters for the DaprInvokeHTTPRoutesClient.ListByRootScope
// method.
func (client *DaprInvokeHTTPRoutesClient) NewListByRootScopePager(options *DaprInvokeHTTPRoutesClientListByRootScopeOptions) (*runtime.Pager[DaprInvokeHTTPRoutesClientListByRootScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[DaprInvokeHTTPRoutesClientListByRootScopeResponse]{
		More: func(page DaprInvokeHTTPRoutesClientListByRootScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprInvokeHTTPRoutesClientListByRootScopeResponse) (DaprInvokeHTTPRoutesClientListByRootScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRootScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DaprInvokeHTTPRoutesClientListByRootScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DaprInvokeHTTPRoutesClientListByRootScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DaprInvokeHTTPRoutesClientListByRootScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRootScopeHandleResponse(resp)
		},
	})
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *DaprInvokeHTTPRoutesClient) listByRootScopeCreateRequest(ctx context.Context, options *DaprInvokeHTTPRoutesClientListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/daprInvokeHttpRoutes"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *DaprInvokeHTTPRoutesClient) listByRootScopeHandleResponse(resp *http.Response) (DaprInvokeHTTPRoutesClientListByRootScopeResponse, error) {
	result := DaprInvokeHTTPRoutesClientListByRootScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteList); err != nil {
		return DaprInvokeHTTPRoutesClientListByRootScopeResponse{}, err
	}
	return result, nil
}

