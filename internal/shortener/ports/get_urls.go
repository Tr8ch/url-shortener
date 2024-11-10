package ports

import (
	"context"

	"url-shortener/internal/shortener/service"

	"github.com/go-kit/kit/endpoint"
)

func NewEndpointGetURLs(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetURLs(ctx)
	}
}
