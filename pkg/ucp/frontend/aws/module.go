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

package aws

import (
	"github.com/gorilla/mux"
	ucp_aws "github.com/project-radius/radius/pkg/ucp/aws"
	"github.com/project-radius/radius/pkg/ucp/frontend/modules"
	"github.com/project-radius/radius/pkg/validator"
)

// NewModule creates a new AWS module.
func NewModule(options modules.Options) *Module {
	m := Module{options: options}
	m.router = mux.NewRouter()
	m.router.NotFoundHandler = validator.APINotFoundHandler()
	m.router.MethodNotAllowedHandler = validator.APIMethodNotAllowedHandler()

	return &m
}

var _ modules.Initializer = &Module{}

// Module defines the module for AWS functionality.
type Module struct {
	options modules.Options
	router  *mux.Router

	// AWSClients provides access to AWS services. This field can be overridden by tests.
	AWSClients ucp_aws.Clients
}

// PlaneType returns the type of plane this module is for.
func (m *Module) PlaneType() string {
	return "aws"
}
