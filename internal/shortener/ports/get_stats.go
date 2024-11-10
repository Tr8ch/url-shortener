package ports

import (
	"context"
	"net/http"

	"url-shortener/internal/shortener/service"
	"url-shortener/pkg/kithelper"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

func DecodeGetStats(_ context.Context, r *http.Request) (interface{}, error) {
	return service.GetStatsInput{
		ShortURL: chi.URLParam(r, "link"),
	}, nil
}

func NewEndpointGetStats(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input, ok := request.(service.GetStatsInput)
		if !ok {
			return nil, errors.WithStack(kithelper.ErrorCastFailed)
		}
		return s.GetStats(ctx, input)
	}
}
