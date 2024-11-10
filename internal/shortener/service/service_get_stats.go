package service

import (
	"context"

	"url-shortener/internal/shortener/domain"
)

type GetStatsInput struct {
	ShortURL string
}

func (s *service) GetStats(ctx context.Context, input GetStatsInput) (*domain.URLInfo, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	exist, err := s.redis.Exists(ctx, input.ShortURL)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, domain.ErrShortURLNotFound
	}

	urlInfo, err := s.redis.GetURLInfo(ctx, input.ShortURL)
	if err != nil {
		return nil, err
	}

	return urlInfo, nil
}

func (i *GetStatsInput) validate() error {
	if len(i.ShortURL) == 0 {
		return domain.ErrValidateShortURL
	}
	return nil
}
