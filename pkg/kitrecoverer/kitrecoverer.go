package kitrecoverer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/go-kit/kit/endpoint"
)

func RecovererMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = convertRecoverToError(r)
				}
			}()

			return next(ctx, request)
		}
	}
}

func convertRecoverToError(r interface{}) error {
	return errors.New(fmt.Sprintf("%v", r))
}
