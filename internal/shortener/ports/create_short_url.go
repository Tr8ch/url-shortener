package ports

import (
	"context"
	"encoding/json"
	"net/http"

	"url-shortener/internal/shortener/service"
	"url-shortener/pkg/kithelper"

	"github.com/go-kit/kit/endpoint"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

func DecodeCreateShortURL(_ context.Context, r *http.Request) (interface{}, error) {
	var request struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, herodot.ErrBadRequest.WithReason(`Unable to parse JSON body`).WithWrap(err)
	}

	return service.CreateShortURLInput{
		OriginalURL: request.URL,
	}, nil
}

func NewEndpointCreateShortURL(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input, ok := request.(service.CreateShortURLInput)
		if !ok {
			return nil, errors.WithStack(kithelper.ErrorCastFailed)
		}

		return s.CreateShortURL(ctx, input)
	}
}
