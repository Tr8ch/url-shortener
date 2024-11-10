package ports

import (
	"context"
	"net/http"

	"url-shortener/internal/shortener/domain"
	"url-shortener/internal/shortener/service"
	"url-shortener/pkg/kithelper"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

func DecodeRedirect(_ context.Context, r *http.Request) (interface{}, error) {
	return service.RedirectInput{
		ShortURL: chi.URLParam(r, "link"),
	}, nil
}

func EncodeRedirect(_ context.Context, w http.ResponseWriter, response interface{}) error {
	urlInfo, ok := response.(*domain.URLInfo)
	if !ok {
		return errors.WithStack(kithelper.ErrorCastFailed)
	}

	http.Redirect(w, &http.Request{}, urlInfo.OriginalURL, http.StatusMovedPermanently)
	return nil
}

func NewEndpointRedirect(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input, ok := request.(service.RedirectInput)
		if !ok {
			return nil, errors.WithStack(kithelper.ErrorCastFailed)
		}

		return s.Redirect(ctx, input)
	}
}
