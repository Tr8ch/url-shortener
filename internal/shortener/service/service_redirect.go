package service

import (
	"context"
	"time"

	"url-shortener/internal/shortener/domain"
)

type RedirectInput struct {
	ShortURL string
}

func (s *service) Redirect(ctx context.Context, input RedirectInput) (*domain.URLInfo, error) {
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

	urlInfo.ClickCounts++
	urlInfo.LastEnteredAt = time.Now()

	if err := s.redis.SetURLInfo(ctx, domain.SetInput{
		ShortURL: input.ShortURL,
		URLInfo:  *urlInfo,
	}); err != nil {
		return nil, err
	}

	return urlInfo, nil
}

func (i *RedirectInput) validate() error {
	if len(i.ShortURL) == 0 {
		return domain.ErrValidateShortURL
	}
	return nil
}
