package service

import (
	"context"
	"errors"
	"parallelSystems/user_gateway/pkg/io"
	"parallelSystems/user_gateway/pkg/model"

	"github.com/dgrijalva/jwt-go"
	log "github.com/go-kit/kit/log"
)

// UserGatewayService describes the service.
type UserGatewayService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error)
	SignUp(ctx context.Context, details io.Login) (credentials io.Credentials, err error)
}

type basicUserGatewayService struct {
	logger log.Logger
}

func (b basicUserGatewayService) Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	var user = &model.User{
		Username: details.Username,
		Password: details.Password,
	}
	b.logger.Log("enter")
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

func (b *basicUserGatewayService) SignUp(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	// TODO implement the business logic of SignUp
	return credentials, err
}
