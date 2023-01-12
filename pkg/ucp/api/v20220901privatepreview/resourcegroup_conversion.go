// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v20220901privatepreview

import (
	to "github.com/Azure/go-autorest/autorest/to"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/ucp/datamodel"
)

// ConvertTo converts from the versioned ResourceGroup resource to version-agnostic datamodel.
func (src *ResourceGroupResource) ConvertTo() (v1.DataModelInterface, error) {
	// Note: SystemData conversion isn't required since this property comes ARM and datastore.

	converted := &datamodel.ResourceGroup{
		TrackedResource: v1.TrackedResource{
			ID:       to.String(src.ID),
			Name:     to.String(src.Name),
			Type:     to.String(src.Type),
			Location: to.String(src.Location),
			Tags:     to.StringMap(src.Tags),
		},
	}

	return converted, nil
}

// ConvertFrom converts from version-agnostic datamodel to the versioned ResourceGroup resource.
func (dst *ResourceGroupResource) ConvertFrom(src v1.DataModelInterface) error {
	// TODO: Improve the validation.
	rg, ok := src.(*datamodel.ResourceGroup)
	if !ok {
		return v1.ErrInvalidModelConversion
	}

	dst.ID = to.StringPtr(rg.ID)
	dst.Name = to.StringPtr(rg.Name)
	dst.Type = to.StringPtr(rg.Type)
	dst.Location = to.StringPtr(rg.Location)
	dst.Tags = *to.StringMapPtr(rg.Tags)

	return nil
}
