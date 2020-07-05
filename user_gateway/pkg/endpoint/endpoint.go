package endpoint

import (
	"context"
	io "parallelSystems/user_gateway/pkg/io"
	service "parallelSystems/user_gateway/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// AuthenticateRequest collects the request parameters for the Authenticate method.
type AuthenticateRequest struct {
	Details io.Login `json:"details"`
}

// AuthenticateResponse collects the response parameters for the Authenticate method.
type AuthenticateResponse struct {
	Credentials io.Credentials `json:"items"`
	Err         error          `json:"err"`
}

// MakeAuthenticateEndpoint returns an endpoint that invokes Authenticate on the service.
func MakeAuthenticateEndpoint(s service.UserGatewayService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthenticateRequest)
		credentials, err := s.Authenticate(ctx, req.Details)
		return AuthenticateResponse{
			Credentials: credentials,
			Err:         err,
		}, nil
	}
}

// Failed implements Failer.
func (r AuthenticateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Authenticate implements Service. Primarily useful in a client.
func (e Endpoints) Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	request := AuthenticateRequest{Details: details}
	response, err := e.AuthenticateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AuthenticateResponse).Credentials, response.(AuthenticateResponse).Err
}

// SignUpRequest collects the request parameters for the SignUp method.
type SignUpRequest struct {
	Details io.Login `json:"details"`
}

// SignUpResponse collects the response parameters for the SignUp method.
type SignUpResponse struct {
	Credentials io.Credentials `json:"credentials"`
	Err         error          `json:"err"`
}

// MakeSignUpEndpoint returns an endpoint that invokes SignUp on the service.
func MakeSignUpEndpoint(s service.UserGatewayService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		credentials, err := s.SignUp(ctx, req.Details)
		return SignUpResponse{
			Credentials: credentials,
			Err:         err,
		}, nil
	}
}

// Failed implements Failer.
func (r SignUpResponse) Failed() error {
	return r.Err
}

// SignUp implements Service. Primarily useful in a client.
func (e Endpoints) SignUp(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	request := SignUpRequest{Details: details}
	response, err := e.SignUpEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignUpResponse).Credentials, response.(SignUpResponse).Err
}
