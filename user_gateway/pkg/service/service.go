package service

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"parallelSystems/user_gateway/pkg/io"
	"parallelSystems/user_gateway/pkg/model"
)

// GatewayService describes the service.
type GatewayService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error)
}

type basicUserGatewayService struct{}

func (b *basicUserGatewayService) Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	var user = &model.User{
		Username: details.Username,
		Password: details.Password,
	}
	isVerified, err := user.CheckPassword()
	if err != nil {
		return io.Credentials{}, err
	}
	if !isVerified {
		err = errors.New("incorrect password")
		return io.Credentials{}, err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	}).SigningString()
	if err != nil {
		return io.Credentials{}, err
	}
	credentials = io.Credentials{
		Id:          user.Id,
		AccessToken: token,
		Username:    user.Username,
	}
	return credentials, err
}

// NewBasicUserGatewayService returns a naive, stateless implementation of GatewayService.
func NewBasicUserGatewayService() GatewayService {
	return &basicUserGatewayService{}
}

// New returns a GatewayService with all of the expected middleware wired in.
func New(middleware []Middleware) GatewayService {
	var svc GatewayService = NewBasicUserGatewayService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
