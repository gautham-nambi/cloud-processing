// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "parallelSystems/user_gateway/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AuthenticateEndpoint endpoint.Endpoint
	SignUpEndpoint       endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.UserGatewayService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AuthenticateEndpoint: MakeAuthenticateEndpoint(s),
		SignUpEndpoint:       MakeSignUpEndpoint(s),
	}
	for _, m := range mdw["Authenticate"] {
		eps.AuthenticateEndpoint = m(eps.AuthenticateEndpoint)
	}
	for _, m := range mdw["SignUp"] {
		eps.SignUpEndpoint = m(eps.SignUpEndpoint)
	}
	return eps
}
