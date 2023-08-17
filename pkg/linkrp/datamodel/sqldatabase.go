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

package datamodel

import (
	"fmt"
	"strings"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
)

// Recipe returns the LinkRecipe associated with the SqlDatabase instance if the ResourceProvisioning is not
// set to Manual, otherwise it returns nil.
func (sql *SqlDatabase) Recipe() *linkrp.LinkRecipe {
	if sql.Properties.ResourceProvisioning == linkrp.ResourceProvisioningManual {
		return nil
	}
	return &sql.Properties.Recipe
}

// SqlDatabase represents SqlDatabase link resource.
type SqlDatabase struct {
	v1.BaseResource

	// Properties is the properties of the resource.
	Properties SqlDatabaseProperties `json:"properties"`

	// LinkMetadata represents internal DataModel properties common to all link types.
	LinkMetadata
}

// ApplyDeploymentOutput updates the output resources of a SqlDatabase resource with the output resources of a DeploymentOutput
// object and returns no error.
func (r *SqlDatabase) ApplyDeploymentOutput(do rpv1.DeploymentOutput) error {
	r.Properties.Status.OutputResources = do.DeployedOutputResources
	return nil
}

// OutputResources returns the OutputResources of the SqlDatabase resource.
func (r *SqlDatabase) OutputResources() []rpv1.OutputResource {
	return r.Properties.Status.OutputResources
}

// ResourceMetadata returns the BasicResourceProperties of the SqlDatabase resource.
func (r *SqlDatabase) ResourceMetadata() *rpv1.BasicResourceProperties {
	return &r.Properties.BasicResourceProperties
}

// ResourceTypeName returns the resource type of the SqlDatabase resource.
func (sql *SqlDatabase) ResourceTypeName() string {
	return linkrp.SqlDatabasesResourceType
}

// SqlDatabaseProperties represents the properties of SqlDatabase resource.
type SqlDatabaseProperties struct {
	rpv1.BasicResourceProperties
	// The recipe used to automatically deploy underlying infrastructure for the SqlDB link
	Recipe linkrp.LinkRecipe `json:"recipe,omitempty"`
	// Database name of the target Sql database
	Database string `json:"database,omitempty"`
	// The fully qualified domain name of the Sql database
	Server string `json:"server,omitempty"`
	// Port value of the target Sql database
	Port int32 `json:"port,omitempty"`
	// Specifies how the underlying service/resource is provisioned and managed
	ResourceProvisioning linkrp.ResourceProvisioning `json:"resourceProvisioning,omitempty"`
	// List of the resource IDs that support the SqlDB resource
	Resources []*linkrp.ResourceReference `json:"resources,omitempty"`
	// Username of the Sql database
	Username string `json:"username,omitempty"`
	// Secrets values provided for the resource
	Secrets SqlDatabaseSecrets `json:"secrets,omitempty"`
}

// Secrets values consisting of secrets provided for the resource
type SqlDatabaseSecrets struct {
	Password         string `json:"password"`
	ConnectionString string `json:"connectionString"`
}

// VerifyInputs checks that the inputs for manual resource provisioning are all provided
//

// VerifyInputs checks if the required fields are set when the resourceProvisioning is set to manual and returns an error
// if any of the required fields are not set.
func (sql *SqlDatabase) VerifyInputs() error {
	msgs := []string{}
	if sql.Properties.ResourceProvisioning != "" && sql.Properties.ResourceProvisioning == linkrp.ResourceProvisioningManual {
		if sql.Properties.Server == "" {
			msgs = append(msgs, "server must be specified when resourceProvisioning is set to manual")
		}
		if sql.Properties.Port == 0 {
			msgs = append(msgs, "port must be specified when resourceProvisioning is set to manual")
		}
		if sql.Properties.Database == "" {
			msgs = append(msgs, "database must be specified when resourceProvisioning is set to manual")
		}
	}

	if len(msgs) == 1 {
		return &v1.ErrClientRP{
			Code:    v1.CodeInvalid,
			Message: msgs[0],
		}
	} else if len(msgs) > 1 {
		return &v1.ErrClientRP{
			Code:    v1.CodeInvalid,
			Message: fmt.Sprintf("multiple errors were found:\n\t%v", strings.Join(msgs, "\n\t")),
		}
	}

	return nil
}

// IsEmpty checks if the SqlDatabaseSecrets struct is empty.
func (sqlSecrets SqlDatabaseSecrets) IsEmpty() bool {
	return sqlSecrets == SqlDatabaseSecrets{}
}

// ResourceTypeName returns the resource type of the SqlDatabase resource.
func (sqlSecrets *SqlDatabaseSecrets) ResourceTypeName() string {
	return linkrp.SqlDatabasesResourceType
}
