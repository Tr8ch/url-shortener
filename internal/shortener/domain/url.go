package domain

import (
	"context"
	"time"

	"github.com/ory/herodot"
)

type URLInfo struct {
	OriginalURL   string    `redis:"original_url"`
	CreatedAt     time.Time `redis:"created_at"`
	ClickCounts   int64     `redis:"click_counts"`
	LastEnteredAt time.Time `redis:"last_entered_at"`
}

type URLs struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

type RedisRepository interface {
	SetURLInfo(ctx context.Context, input SetInput) error
	GetURLInfo(ctx context.Context, shortURL string) (*URLInfo, error)

	SetOriginalURL(ctx context.Context, input SetOriginalURLInput) error
	GetShortURL(ctx context.Context, originalURL string) (string, error)

	GetURLs(ctx context.Context) ([]URLs, error)

	Exists(ctx context.Context, url string) (bool, error)

	DeleteURL(ctx context.Context, shortURL string) error
}

type (
	SetInput struct {
		ShortURL string
		URLInfo  URLInfo
	}
	SetOriginalURLInput struct {
		ShortURL    string
		OriginalURL string
	}
)

var (
	// validate errors
	ErrValidateProtocol = herodot.ErrBadRequest.WithReason(`incorrect "http/https" protocol`)
	ErrValidateURL      = herodot.ErrBadRequest.WithReason(`incorrect "url" field`)
	ErrValidateShortURL = herodot.ErrBadRequest.WithReason(`incorrect "short_url" field`)

	// repository errors
	ErrShortURLNotFound = herodot.ErrNotFound.WithReason(`short URL not found`)
	ErrURLsNotFound     = herodot.ErrNotFound.WithReason(`URLs not found`)
)
