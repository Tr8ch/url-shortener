package service

import (
	"context"

	"url-shortener/internal/shortener/domain"
)

type Service interface {
	CreateShortURL(context.Context, CreateShortURLInput) (string, error)
	GetURLs(context.Context) (*GetURLsResponse, error)
	Redirect(context.Context, RedirectInput) (*domain.URLInfo, error)
	GetStats(context.Context, GetStatsInput) (*domain.URLInfo, error)
	DeleteURL(context.Context, DeleteShortURLInput) error
}

type service struct {
	redis  domain.RedisRepository
	urlLen int
}

var _ Service = (*service)(nil)

func New(
	redis domain.RedisRepository,
	urlLen int,
) Service {
	return &service{
		redis:  redis,
		urlLen: urlLen,
	}
}
