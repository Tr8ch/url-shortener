package kithelper

import (
	"github.com/go-kit/kit/endpoint"
)

func ChainMiddlewares(middlewares []endpoint.Middleware) endpoint.Middleware {
	if len(middlewares) == 0 {
		return func(next endpoint.Endpoint) endpoint.Endpoint {
			return next
		}
	}

	return endpoint.Chain(middlewares[0], middlewares[1:]...)
}
