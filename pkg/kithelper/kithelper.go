package kithelper

import (
	"context"
	"encoding/json"
	"net/http"

	"url-shortener/pkg/jsonhelper"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

type ErrorJSON interface {
	JSON() any
}

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	if c := herodot.StatusCodeCarrier(nil); errors.As(err, &c) {
		writeErrorCode(w, c.StatusCode(), err)
	} else {
		writeErrorCode(w, http.StatusInternalServerError, herodot.ErrInternalServerError)
	}
}

type errorWithStack interface {
	Cause() error
}

func errorIsWithStack(err error) bool {
	_, ok := err.(errorWithStack)
	return ok
}

func writeErrorCode(w http.ResponseWriter, code int, err error) {
	var payload interface{} = err

	if errorIsWithStack(err) {
		payload = errors.Unwrap(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func EmptyResponse(_ context.Context, w http.ResponseWriter, _ interface{}) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func EncodeResponse[A, B comparable](fn jsonhelper.Encoder[A, B]) kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		output, ok := response.(A)
		if !ok {
			return http.ErrNotSupported
		}
		jsonOutput := fn(output)

		return kithttp.EncodeJSONResponse(ctx, w, jsonOutput)
	}
}

func EmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil //nolint:nilnil  // uneccessary error check.
}

var ErrorCastFailed = errors.New("cast failed")
