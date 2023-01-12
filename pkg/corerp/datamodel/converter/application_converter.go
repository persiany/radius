// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package converter

import (
	"encoding/json"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	v20220315privatepreview "github.com/project-radius/radius/pkg/corerp/api/v20220315privatepreview"
	"github.com/project-radius/radius/pkg/corerp/datamodel"
)

// ApplicationDataModelToVersioned converts version agnostic application datamodel to versioned model.
func ApplicationDataModelToVersioned(model *datamodel.Application, version string) (v1.VersionedModelInterface, error) {
	switch version {
	case v20220315privatepreview.Version:
		versioned := &v20220315privatepreview.ApplicationResource{}
		err := versioned.ConvertFrom(model)
		return versioned, err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}

// ApplicationDataModelFromVersioned converts versioned application model to datamodel.
func ApplicationDataModelFromVersioned(content []byte, version string) (*datamodel.Application, error) {
	switch version {
	case v20220315privatepreview.Version:
		am := &v20220315privatepreview.ApplicationResource{}
		if err := json.Unmarshal(content, am); err != nil {
			return nil, err
		}
		dm, err := am.ConvertTo()
		return dm.(*datamodel.Application), err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}
