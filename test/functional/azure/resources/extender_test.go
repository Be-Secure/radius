// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package resources_test

import (
	"testing"

	"github.com/project-radius/radius/pkg/providers"
	"github.com/project-radius/radius/pkg/radrp/outputresource"
	"github.com/project-radius/radius/pkg/radrp/rest"
	"github.com/project-radius/radius/pkg/renderers/containerv1alpha3"
	"github.com/project-radius/radius/pkg/renderers/extenderv1alpha3"
	"github.com/project-radius/radius/pkg/resourcekinds"
	"github.com/project-radius/radius/test/functional"
	"github.com/project-radius/radius/test/functional/azure"
	"github.com/project-radius/radius/test/step"
	"github.com/project-radius/radius/test/validation"
)

func Test_ExtenderResource(t *testing.T) {
	application := "azure-resources-extender"
	template := "testdata/azure-resources-extender.bicep"

	test := azure.NewApplicationTest(t, application, []azure.TestStep{
		{
			Executor:       step.NewDeployExecutor(template, functional.GetMagpieImage()),
			AzureResources: &validation.AzureResourceSet{},
			RadiusResources: &validation.ResourceSet{
				Resources: []validation.RadiusResource{
					{
						ApplicationName: application,
						ResourceName:    "twilio",
						ResourceType:    extenderv1alpha3.ResourceType,
						OutputResources: map[string]validation.ExpectedOutputResource{},
					},
					{
						ApplicationName: application,
						ResourceName:    "myapp",
						ResourceType:    containerv1alpha3.ResourceType,
						OutputResources: map[string]validation.ExpectedOutputResource{
							outputresource.LocalIDDeployment: validation.NewOutputResource(outputresource.LocalIDDeployment, rest.ResourceType{Type: resourcekinds.Deployment, Provider: providers.ProviderKubernetes}, false, rest.OutputResourceStatus{}),
						},
					},
				},
			},
			Objects: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					application: {
						validation.NewK8sPodForResource(application, "myapp"),
					},
				},
			},
		},
	})

	test.Test(t)
}
