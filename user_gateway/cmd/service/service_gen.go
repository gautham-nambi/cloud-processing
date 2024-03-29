// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "parallelSystems/user_gateway/pkg/endpoint"
	http1 "parallelSystems/user_gateway/pkg/http"
	"parallelSystems/user_gateway/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Authenticate": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Authenticate", logger))},
		"SignUp":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SignUp", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Authenticate"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Authenticate")), endpoint.InstrumentingMiddleware(duration.With("method", "Authenticate"))}
	mw["SignUp"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SignUp")), endpoint.InstrumentingMiddleware(duration.With("method", "SignUp"))}
}

func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Authenticate", "SignUp"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
