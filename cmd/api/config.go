package main

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// Config domain with application settings variables.
type Config struct {
	Redis struct {
		Addr string `env:"ADDR, default=localhost:6379"`
		DB   int    `env:"DB, default=5"`
	} `env:", prefix=REDIS_"`
	ENV              string `env:"ENV, default=prod"`
	Port             string `env:"PORT, default=8888"`
	ExpirationInDays int    `env:"EXPIRATION_IN_DAYS, default=30"`
	URLLen           int    `env:"URL_LEN, default=8"`
}

func getConfig(ctx context.Context) (*Config, error) {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
