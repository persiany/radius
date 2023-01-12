// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v20220315privatepreview

import (
	"encoding/json"
	"testing"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/corerp/datamodel"
	radiustesting "github.com/project-radius/radius/pkg/corerp/testing"
	"github.com/project-radius/radius/pkg/rp/outputresource"
	"github.com/stretchr/testify/require"
)

func TestVolumeConvertVersionedToDataModel(t *testing.T) {
	// arrange
	r := &VolumeResource{}
	err := json.Unmarshal(radiustesting.ReadFixture("volume-az-kv.json"), r)
	require.NoError(t, err)

	expected := &datamodel.VolumeResource{}
	err = json.Unmarshal(radiustesting.ReadFixture("volume-az-kv-datamodel.json"), expected)
	require.NoError(t, err)

	// act
	dm, err := r.ConvertTo()

	// assert
	require.NoError(t, err)
	ct := dm.(*datamodel.VolumeResource)
	require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/volumes/azkeyvault0", ct.ID)
	require.Equal(t, "azkeyvault0", ct.Name)
	require.Equal(t, "Applications.Core/volumes", ct.Type)
	require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testGroup/providers/Applications.Core/applications/app0", ct.Properties.Application)
	require.Equal(t, []outputresource.OutputResource(nil), ct.Properties.Status.OutputResources)
	require.Equal(t, "2022-03-15-privatepreview", ct.InternalMetadata.UpdatedAPIVersion)
	require.Equal(t, expected.Properties.AzureKeyVault, ct.Properties.AzureKeyVault)
}

func TestVolumeConvertDataModelToVersioned(t *testing.T) {
	// arrange
	r := &datamodel.VolumeResource{}
	err := json.Unmarshal(radiustesting.ReadFixture("volume-az-kv-datamodel.json"), r)
	require.NoError(t, err)

	expected := &VolumeResource{}
	err = json.Unmarshal(radiustesting.ReadFixture("volume-az-kv.json"), expected)
	require.NoError(t, err)
	expected.Properties.GetVolumeProperties().Status.OutputResources[0]["Identity"] = nil

	// act
	versioned := &VolumeResource{}
	err = versioned.ConvertFrom(r)

	// assert
	require.NoError(t, err)
	require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/radius-test-rg/providers/Applications.Core/volumes/azkeyvault0", r.ID)
	require.Equal(t, "azkeyvault0", r.Name)
	require.Equal(t, "Applications.Core/volumes", r.Type)
	require.Equal(t, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testGroup/providers/Applications.Core/applications/app0", r.Properties.Application)
	require.Equal(t, expected.Properties, versioned.Properties)
}

func TestVolumeConvertFromValidation(t *testing.T) {
	validationTests := []struct {
		src v1.DataModelInterface
		err error
	}{
		{&fakeResource{}, v1.ErrInvalidModelConversion},
		{nil, v1.ErrInvalidModelConversion},
	}

	for _, tc := range validationTests {
		versioned := &HTTPRouteResource{}
		err := versioned.ConvertFrom(tc.src)
		require.ErrorAs(t, tc.err, &err)
	}
}
