// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// swagger endpoints
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design -o
// $(GOPATH)/src/goa.design/goa/examples/cellar

package swagger

import (
	goa "goa.design/goa"
)

// Endpoints wraps the "swagger" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "swagger" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "swagger" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
