package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type Dependencies struct {
	postgres *pgxpool.Pool
	redis    *redis.Client
	logger   *slog.Logger
}

func (d *Dependencies) Close() {
	if d == nil {
		return
	}
	if d.postgres != nil {
		d.postgres.Close()
	}
	if d.redis != nil {
		d.redis.Close()
	}
}

func NewDependencies(ctx context.Context, opts ...Option) (deps *Dependencies, err error) {
	defer func() {
		if err != nil {
			deps.Close()
		}
	}()

	deps = &Dependencies{}
	for _, opt := range opts {
		if err := opt(ctx, deps); err != nil {
			return nil, err
		}
	}

	return deps, nil
}

type Option func(context.Context, *Dependencies) error

func WithRedis(addr string, db int) Option {
	return func(ctx context.Context, d *Dependencies) error {
		client := redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   db,
		})

		if err := client.Ping(ctx).Err(); err != nil {
			return err
		}

		d.redis = client
		return nil
	}
}

func WithLogger(env string) Option {
	return func(_ context.Context, d *Dependencies) error {
		var slogLevel slog.Level
		switch env {
		case EnvLocal:
			slogLevel = slog.LevelDebug
		case EnvDev:
			slogLevel = slog.LevelInfo
		case EnvProd:
			slogLevel = slog.LevelInfo
		default:
			slogLevel = slog.LevelInfo
		}

		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel,
		}))
		slog.SetDefault(logger)
		d.logger = logger

		return nil
	}
}
