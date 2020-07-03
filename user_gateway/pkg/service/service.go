package service

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	"parallelSystems/user_gateway/pkg/io"
)

// UserGatewayService describes the service.
type UserGatewayService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error)
}

type basicUserGatewayService struct{}

func (b *basicUserGatewayService) Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	token, err := jwt.New(jwt.SigningMethodHS256).SignedString(io.JwtSecret)
	credentials.Id = uuid.New()
	credentials.AccessToken = token
	return credentials, err
}

// NewBasicUserGatewayService returns a naive, stateless implementation of UserGatewayService.
func NewBasicUserGatewayService() UserGatewayService {
	return &basicUserGatewayService{}
}

// New returns a UserGatewayService with all of the expected middleware wired in.
func New(middleware []Middleware) UserGatewayService {
	var svc UserGatewayService = NewBasicUserGatewayService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
