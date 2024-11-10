package kitlogger

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/ory/herodot"
)

func LoggingMiddleware(logger *slog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			resp, err := next(ctx, request)
			if err != nil {
				if carrier := herodot.StatusCodeCarrier(nil); !errors.As(err, &carrier) {
					logger.Error(
						"internal error",
						"request", request,
						"error", err,
						"stack_trace", encodeStackTrace(err),
					)
				}
			}

			return resp, err
		}
	}
}

type stackFrame struct {
	Frame string `json:"frame"`
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func encodeStackTrace(err error) []stackFrame {
	tracer, ok := err.(stackTracer)
	if !ok {
		return nil
	}

	trace := tracer.StackTrace()
	if len(trace) == 0 {
		return nil
	}

	s := make([]stackFrame, len(trace))

	for i, v := range trace {
		frame, _ := v.MarshalText()
		f := stackFrame{
			Frame: string(frame),
		}

		s[i] = f
	}

	return s
}
