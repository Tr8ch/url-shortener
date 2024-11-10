package support

import (
	"net/http"

	"url-shortener/internal/shortener/ports"
	"url-shortener/internal/shortener/service"
	"url-shortener/pkg/kithelper"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHandler(
	s service.Service,
	endpointMiddlewars []endpoint.Middleware,
	serverOptions []kithttp.ServerOption,
) http.Handler {
	middleware := kithelper.ChainMiddlewares(endpointMiddlewars)

	r := chi.NewRouter()

	r.Handle("POST /shortener", kithttp.NewServer(
		middleware(ports.NewEndpointCreateShortURL(s)),
		ports.DecodeCreateShortURL,
		kithttp.EncodeJSONResponse,
		serverOptions...,
	))

	r.Handle("GET /shortener", kithttp.NewServer(
		middleware(ports.NewEndpointGetURLs(s)),
		kithelper.EmptyRequest,
		kithttp.EncodeJSONResponse,
		serverOptions...,
	))

	r.Handle("GET /{link}", kithttp.NewServer(
		middleware(ports.NewEndpointRedirect(s)),
		ports.DecodeRedirect,
		ports.EncodeRedirect,
		serverOptions...,
	))

	r.Handle("GET /stats/{link}", kithttp.NewServer(
		middleware(ports.NewEndpointGetStats(s)),
		ports.DecodeGetStats,
		kithttp.EncodeJSONResponse,
		serverOptions...,
	))

	r.Handle("DELETE /{link}", kithttp.NewServer(
		middleware(ports.NewEndpointDeleteURL(s)),
		ports.DecodeDeleteURL,
		kithelper.EmptyResponse,
		serverOptions...,
	))

	return r
}
