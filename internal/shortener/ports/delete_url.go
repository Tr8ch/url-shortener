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

func DecodeDeleteURL(_ context.Context, r *http.Request) (interface{}, error) {
	return service.DeleteShortURLInput{
		ShortURL: chi.URLParam(r, "link"),
	}, nil
}

func NewEndpointDeleteURL(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input, ok := request.(service.DeleteShortURLInput)
		if !ok {
			return nil, errors.WithStack(kithelper.ErrorCastFailed)
		}
		return nil, s.DeleteURL(ctx, input)
	}
}
