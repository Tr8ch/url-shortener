package service

import (
	"context"

	"url-shortener/internal/shortener/domain"
)

type DeleteShortURLInput struct {
	ShortURL string
}

func (s *service) DeleteURL(ctx context.Context, input DeleteShortURLInput) error {
	if err := input.validate(); err != nil {
		return err
	}

	exist, err := s.redis.Exists(ctx, input.ShortURL)
	if err != nil {
		return err
	}
	if !exist {
		return domain.ErrShortURLNotFound
	}

	urlInfo, err := s.redis.GetURLInfo(ctx, input.ShortURL)
	if err != nil {
		return err
	}

	if err := s.redis.DeleteURL(ctx, input.ShortURL); err != nil {
		return err
	}
	if err := s.redis.DeleteURL(ctx, urlInfo.OriginalURL); err != nil {
		return err
	}

	return nil
}

func (i *DeleteShortURLInput) validate() error {
	if len(i.ShortURL) == 0 {
		return domain.ErrValidateShortURL
	}
	return nil
}
