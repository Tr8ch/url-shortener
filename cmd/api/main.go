package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"url-shortener/pkg/kithelper"
	"url-shortener/pkg/kitlogger"
	"url-shortener/pkg/kitrecoverer"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	shortener "url-shortener/internal/shortener"
)

func main() {
	ctx := context.Background()

	conf, err := getConfig(ctx)
	if err != nil {
		slog.Error("Unable to get configuration", "error", err)
		return
	}

	deps, err := NewDependencies(
		ctx,
		WithRedis(conf.Redis.Addr, conf.Redis.DB),
		WithLogger(conf.ENV),
	)
	if err != nil {
		slog.Error("Unable to initialize dependencies", "error", err)
		return
	}
	defer deps.Close()

	var (
		endpointMdlws = []endpoint.Middleware{
			kitlogger.LoggingMiddleware(deps.logger),
			kitrecoverer.RecovererMiddleware(),
		}
		serverOpts = []kithttp.ServerOption{
			kithttp.ServerErrorEncoder(kithelper.ErrorEncoder),
		}
	)

	shortenerSvc := shortener.NewService(
		shortener.WithRedisRepository(deps.redis, conf.ExpirationInDays),
		shortener.WithURLLen(conf.URLLen),
	)

	router := chi.NewRouter()

	router.Mount("/", shortener.NewHandler(shortenerSvc, endpointMdlws, serverOpts))

	http.Handle("/", kithelper.AccessControl(router))
	run(conf.Port)
}

func run(port string) {
	addr := ":" + port
	server := &http.Server{
		Addr: addr,
	}

	errs := make(chan error, 2)
	go func() {
		slog.Info(
			"listening",
			"address", addr,
		)

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errs <- err
		}
		slog.Info("Stopped serving new connections")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			errs <- err
		}

		errs <- fmt.Errorf("%s", sig)
	}()

	slog.Info("terminated", "err", <-errs)
}
