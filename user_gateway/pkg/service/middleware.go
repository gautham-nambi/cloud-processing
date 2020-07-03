package service

import (
	"context"
	io "parallelSystems/user_gateway/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(GatewayService) GatewayService

type loggingMiddleware struct {
	logger log.Logger
	next   GatewayService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a GatewayService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next GatewayService) GatewayService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Authenticate(ctx context.Context, details io.Login) (credentials io.Credentials, err error) {
	defer func() {
		l.logger.Log("method", "Authenticate", "details", details, "credentials", credentials, "err", err)
	}()
	return l.next.Authenticate(ctx, details)
}
