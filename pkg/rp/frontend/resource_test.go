// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package frontend

import (
	"github.com/Azure/go-autorest/autorest/to"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/rp"
	"github.com/project-radius/radius/pkg/rp/outputresource"
)

const (
	testHeaderfile = "resource-request-headers.json"
	testAPIVersion = "2022-03-15-privatepreview"
)

// TestResourceDataModel represents test resource.
type TestResourceDataModel struct {
	v1.BaseResource

	// Properties is the properties of the resource.
	Properties *TestResourceDataModelProperties `json:"properties"`
}

// ResourceTypeName returns the qualified name of the resource
func (r *TestResourceDataModel) ResourceTypeName() string {
	return "Applications.Core/resources"
}

// ApplyDeploymentOutput applies the properties changes based on the deployment output.
func (c *TestResourceDataModel) ApplyDeploymentOutput(do rp.DeploymentOutput) {
	c.Properties.Status.OutputResources = do.DeployedOutputResources
}

// OutputResources returns the output resources array.
func (c *TestResourceDataModel) OutputResources() []outputresource.OutputResource {
	return c.Properties.Status.OutputResources
}

// ResourceMetadata returns the application resource metadata.
func (h *TestResourceDataModel) ResourceMetadata() *rp.BasicResourceProperties {
	return &h.Properties.BasicResourceProperties
}

// TestResourceDataModelProperties represents the properties of TestResourceDataModel.
type TestResourceDataModelProperties struct {
	rp.BasicResourceProperties
	PropertyA string `json:"propertyA,omitempty"`
	PropertyB string `json:"propertyB,omitempty"`
}

// TestResource represents test resource for api version.
type TestResource struct {
	ID         *string                 `json:"id,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	SystemData *v1.SystemData          `json:"systemData,omitempty"`
	Type       *string                 `json:"type,omitempty"`
	Location   *string                 `json:"location,omitempty"`
	Properties *TestResourceProperties `json:"properties,omitempty"`
	Tags       map[string]*string      `json:"tags,omitempty"`
}

// TestResourceProperties - HTTP Route properties
type TestResourceProperties struct {
	ProvisioningState *v1.ProvisioningState `json:"provisioningState,omitempty"`
	Environment       *string               `json:"environment,omitempty"`
	Application       *string               `json:"application,omitempty"`
	PropertyA         *string               `json:"propertyA,omitempty"`
	PropertyB         *string               `json:"propertyB,omitempty"`
	Status            *ResourceStatus       `json:"status,omitempty"`
}

// ResourceStatus - Status of a resource.
type ResourceStatus struct {
	OutputResources []map[string]any `json:"outputResources,omitempty"`
}

func (src *TestResource) ConvertTo() (v1.DataModelInterface, error) {
	converted := &TestResourceDataModel{
		BaseResource: v1.BaseResource{
			TrackedResource: v1.TrackedResource{
				ID:       to.String(src.ID),
				Name:     to.String(src.Name),
				Type:     to.String(src.Type),
				Location: to.String(src.Location),
				Tags:     to.StringMap(src.Tags),
			},
			InternalMetadata: v1.InternalMetadata{
				UpdatedAPIVersion:      testAPIVersion,
				AsyncProvisioningState: toProvisioningStateDataModel(src.Properties.ProvisioningState),
			},
		},
		Properties: &TestResourceDataModelProperties{
			BasicResourceProperties: rp.BasicResourceProperties{
				Environment: to.String(src.Properties.Environment),
				Application: to.String(src.Properties.Application),
			},
			PropertyA: to.String(src.Properties.PropertyA),
			PropertyB: to.String(src.Properties.PropertyB),
		},
	}
	return converted, nil
}

func (dst *TestResource) ConvertFrom(src v1.DataModelInterface) error {
	dm, ok := src.(*TestResourceDataModel)
	if !ok {
		return v1.ErrInvalidModelConversion
	}

	dst.ID = to.StringPtr(dm.ID)
	dst.Name = to.StringPtr(dm.Name)
	dst.Type = to.StringPtr(dm.Type)
	dst.SystemData = &dm.SystemData
	dst.Location = to.StringPtr(dm.Location)
	dst.Tags = *to.StringMapPtr(dm.Tags)
	dst.Properties = &TestResourceProperties{
		Status: &ResourceStatus{
			OutputResources: rp.BuildExternalOutputResources(dm.Properties.Status.OutputResources),
		},
		ProvisioningState: fromProvisioningStateDataModel(dm.InternalMetadata.AsyncProvisioningState),
		Environment:       to.StringPtr(dm.Properties.Environment),
		Application:       to.StringPtr(dm.Properties.Application),
		PropertyA:         to.StringPtr(dm.Properties.PropertyA),
		PropertyB:         to.StringPtr(dm.Properties.PropertyB),
	}

	return nil
}

func toProvisioningStateDataModel(state *v1.ProvisioningState) v1.ProvisioningState {
	if state == nil {
		return v1.ProvisioningStateAccepted
	}
	return *state
}

func fromProvisioningStateDataModel(state v1.ProvisioningState) *v1.ProvisioningState {
	converted := v1.ProvisioningStateAccepted
	if state != "" {
		converted = state
	}

	return &converted
}
