package service

import (
	"context"

	"url-shortener/internal/shortener/domain"
)

type GetURLsResponse struct {
	Total int           `json:"total"`
	URLs  []domain.URLs `json:"urls"`
}

func (s *service) GetURLs(ctx context.Context) (*GetURLsResponse, error) {
	urls, err := s.redis.GetURLs(ctx)
	if err != nil {
		return nil, err
	}

	return &GetURLsResponse{
		URLs:  urls,
		Total: len(urls),
	}, nil
}
