/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package show

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/radius-project/radius/pkg/cli/clierrors"
	"github.com/radius-project/radius/pkg/cli/connections"
	cli_credential "github.com/radius-project/radius/pkg/cli/credential"
	"github.com/radius-project/radius/pkg/cli/framework"
	"github.com/radius-project/radius/pkg/cli/objectformats"
	"github.com/radius-project/radius/pkg/cli/output"
	"github.com/radius-project/radius/pkg/cli/workspaces"
	"github.com/radius-project/radius/test/radcli"
	"github.com/stretchr/testify/require"
)

func Test_CommandValidation(t *testing.T) {
	radcli.SharedCommandValidation(t, NewCommand)
}

func Test_Validate(t *testing.T) {
	configWithWorkspace := radcli.LoadConfigWithWorkspace(t)
	testcases := []radcli.ValidateInput{
		{
			Name:          "Valid Show Command",
			Input:         []string{"azure"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
		{
			Name:          "Show Command with fallback workspace",
			Input:         []string{"Azure"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         radcli.LoadEmptyConfig(t),
			},
		},
		{
			Name:          "Show Command with unsupported provider type",
			Input:         []string{"invalidProviderType"},
			ExpectedValid: false,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
		{
			Name:          "Show Command with insufficient args",
			Input:         []string{},
			ExpectedValid: false,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
		{
			Name:          "Show Command with too many args",
			Input:         []string{"azure", "a", "b"},
			ExpectedValid: false,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
		{
			Name:          "Valid Show AWS Command",
			Input:         []string{"aws"},
			ExpectedValid: true,
			ConfigHolder: framework.ConfigHolder{
				ConfigFilePath: "",
				Config:         configWithWorkspace,
			},
		},
	}
	radcli.SharedValidateValidation(t, NewCommand, testcases)
}

func Test_Run(t *testing.T) {
	connection := map[string]any{
		"kind":    workspaces.KindKubernetes,
		"context": "my-context",
	}

	t.Run("Show azure provider", func(t *testing.T) {
		t.Run("Exists", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			provider := cli_credential.ProviderCredentialConfiguration{
				CloudProviderStatus: cli_credential.CloudProviderStatus{
					Name:    "azure",
					Enabled: true,
				},
			}

			client := cli_credential.NewMockCredentialManagementClient(ctrl)
			client.EXPECT().
				Get(gomock.Any(), "azure").
				Return(provider, nil).
				Times(1)

			outputSink := &output.MockOutput{}

			runner := &Runner{
				ConnectionFactory: &connections.MockFactory{CredentialManagementClient: client},
				Output:            outputSink,
				Workspace:         &workspaces.Workspace{Connection: connection},
				Kind:              "azure",
				Format:            "table",
			}

			err := runner.Run(context.Background())
			require.NoError(t, err)

			expected := []any{
				output.LogOutput{
					Format: "Showing credential for cloud provider %q for Radius installation %q...",
					Params: []any{"azure", "Kubernetes (context=my-context)"},
				},
				output.FormattedOutput{
					Format:  "table",
					Obj:     provider,
					Options: objectformats.GetCloudProviderTableFormat(runner.Kind),
				},
			}
			require.Equal(t, expected, outputSink.Writes)
		})
		t.Run("Not Found", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			client := cli_credential.NewMockCredentialManagementClient(ctrl)
			client.EXPECT().
				Get(gomock.Any(), "azure").
				Return(cli_credential.ProviderCredentialConfiguration{}, nil).
				Times(1)

			outputSink := &output.MockOutput{}

			runner := &Runner{
				ConnectionFactory: &connections.MockFactory{CredentialManagementClient: client},
				Output:            outputSink,
				Workspace:         &workspaces.Workspace{Connection: connection},
				Kind:              "azure",
				Format:            "table",
			}

			err := runner.Run(context.Background())
			expected := clierrors.Message("The credentials for cloud provider %q could not be found.", runner.Kind)
			require.Equal(t, expected, err)
		})
	})
	t.Run("Show aws provider", func(t *testing.T) {
		t.Run("Exists", func(t *testing.T) {
			ctrl := gomock.NewController(t)

			provider := cli_credential.ProviderCredentialConfiguration{
				CloudProviderStatus: cli_credential.CloudProviderStatus{
					Name:    "aws",
					Enabled: true,
				},
			}

			client := cli_credential.NewMockCredentialManagementClient(ctrl)
			client.EXPECT().
				Get(gomock.Any(), "aws").
				Return(provider, nil).
				Times(1)

			outputSink := &output.MockOutput{}

			runner := &Runner{
				ConnectionFactory: &connections.MockFactory{CredentialManagementClient: client},
				Output:            outputSink,
				Workspace:         &workspaces.Workspace{Connection: connection},
				Kind:              "aws",
				Format:            "table",
			}

			err := runner.Run(context.Background())
			require.NoError(t, err)

			expected := []any{
				output.LogOutput{
					Format: "Showing credential for cloud provider %q for Radius installation %q...",
					Params: []any{"aws", "Kubernetes (context=my-context)"},
				},
				output.FormattedOutput{
					Format:  "table",
					Obj:     provider,
					Options: objectformats.GetCloudProviderTableFormat(runner.Kind),
				},
			}
			require.Equal(t, expected, outputSink.Writes)
		})
		t.Run("Not Found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			provider := cli_credential.ProviderCredentialConfiguration{
				CloudProviderStatus: cli_credential.CloudProviderStatus{
					Name:    "aws",
					Enabled: false,
				},
			}
			client := cli_credential.NewMockCredentialManagementClient(ctrl)
			client.EXPECT().
				Get(gomock.Any(), "aws").
				Return(provider, nil).
				Times(1)

			outputSink := &output.MockOutput{}

			runner := &Runner{
				ConnectionFactory: &connections.MockFactory{CredentialManagementClient: client},
				Output:            outputSink,
				Workspace:         &workspaces.Workspace{Connection: connection},
				Kind:              "aws",
				Format:            "table",
			}

			err := runner.Run(context.Background())
			expected := clierrors.Message("The credentials for cloud provider %q could not be found.", runner.Kind)
			require.Equal(t, expected, err)
		})
	})

}
