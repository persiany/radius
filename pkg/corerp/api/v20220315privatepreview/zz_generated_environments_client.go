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

// EnvironmentsClient contains the methods for the Environments group.
// Don't use this type directly, use NewEnvironmentsClient() instead.
type EnvironmentsClient struct {
	host string
	rootScope string
	pl runtime.Pipeline
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEnvironmentsClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnvironmentsClient, error) {
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
	client := &EnvironmentsClient{
		rootScope: rootScope,
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an Environment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// environmentName - The name of the environment
// environmentResource - environment details
// options - EnvironmentsClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentsClient.CreateOrUpdate
// method.
func (client *EnvironmentsClient) CreateOrUpdate(ctx context.Context, environmentName string, environmentResource EnvironmentResource, options *EnvironmentsClientCreateOrUpdateOptions) (EnvironmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, environmentName, environmentResource, options)
	if err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return EnvironmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, environmentName string, environmentResource EnvironmentResource, options *EnvironmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, environmentResource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EnvironmentsClient) createOrUpdateHandleResponse(resp *http.Response) (EnvironmentsClientCreateOrUpdateResponse, error) {
	result := EnvironmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResource); err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an Environment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// environmentName - The name of the environment
// options - EnvironmentsClientDeleteOptions contains the optional parameters for the EnvironmentsClient.Delete method.
func (client *EnvironmentsClient) Delete(ctx context.Context, environmentName string, options *EnvironmentsClientDeleteOptions) (EnvironmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, environmentName, options)
	if err != nil {
		return EnvironmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnvironmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return EnvironmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return EnvironmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EnvironmentsClient) deleteCreateRequest(ctx context.Context, environmentName string, options *EnvironmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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

// Get - Gets the properties of an Environment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// environmentName - The name of the environment
// options - EnvironmentsClientGetOptions contains the optional parameters for the EnvironmentsClient.Get method.
func (client *EnvironmentsClient) Get(ctx context.Context, environmentName string, options *EnvironmentsClientGetOptions) (EnvironmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, environmentName, options)
	if err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnvironmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EnvironmentsClient) getCreateRequest(ctx context.Context, environmentName string, options *EnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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
func (client *EnvironmentsClient) getHandleResponse(resp *http.Response) (EnvironmentsClientGetResponse, error) {
	result := EnvironmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResource); err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - List all environments in a scope.
// Generated from API version 2022-03-15-privatepreview
// options - EnvironmentsClientListByScopeOptions contains the optional parameters for the EnvironmentsClient.ListByScope
// method.
func (client *EnvironmentsClient) NewListByScopePager(options *EnvironmentsClientListByScopeOptions) (*runtime.Pager[EnvironmentsClientListByScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[EnvironmentsClientListByScopeResponse]{
		More: func(page EnvironmentsClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnvironmentsClientListByScopeResponse) (EnvironmentsClientListByScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnvironmentsClientListByScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EnvironmentsClientListByScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnvironmentsClientListByScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByScopeHandleResponse(resp)
		},
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *EnvironmentsClient) listByScopeCreateRequest(ctx context.Context, options *EnvironmentsClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments"
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

// listByScopeHandleResponse handles the ListByScope response.
func (client *EnvironmentsClient) listByScopeHandleResponse(resp *http.Response) (EnvironmentsClientListByScopeResponse, error) {
	result := EnvironmentsClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResourceList); err != nil {
		return EnvironmentsClientListByScopeResponse{}, err
	}
	return result, nil
}

// Update - Update the properties of an existing Environment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// environmentName - The name of the environment
// environmentResource - environment details
// options - EnvironmentsClientUpdateOptions contains the optional parameters for the EnvironmentsClient.Update method.
func (client *EnvironmentsClient) Update(ctx context.Context, environmentName string, environmentResource EnvironmentResource, options *EnvironmentsClientUpdateOptions) (EnvironmentsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, environmentName, environmentResource, options)
	if err != nil {
		return EnvironmentsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnvironmentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return EnvironmentsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EnvironmentsClient) updateCreateRequest(ctx context.Context, environmentName string, environmentResource EnvironmentResource, options *EnvironmentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, environmentResource)
}

// updateHandleResponse handles the Update response.
func (client *EnvironmentsClient) updateHandleResponse(resp *http.Response) (EnvironmentsClientUpdateResponse, error) {
	result := EnvironmentsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResource); err != nil {
		return EnvironmentsClientUpdateResponse{}, err
	}
	return result, nil
}

