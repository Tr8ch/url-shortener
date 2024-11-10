package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"net/url"
	"regexp"
	"time"

	"url-shortener/internal/shortener/domain"
)

type CreateShortURLInput struct {
	OriginalURL string
}

func (s *service) CreateShortURL(ctx context.Context, input CreateShortURLInput) (string, error) {
	if err := input.validate(); err != nil {
		return "", err
	}

	exist, err := s.redis.Exists(ctx, input.OriginalURL)
	if err != nil {
		return "", err
	}
	if exist {
		return s.redis.GetShortURL(ctx, input.OriginalURL)
	}

	shortURL := generateShortURL(s.urlLen)

	if err := s.redis.SetOriginalURL(ctx, domain.SetOriginalURLInput{
		ShortURL:    shortURL,
		OriginalURL: input.OriginalURL,
	}); err != nil {
		return "", err
	}

	if err := s.redis.SetURLInfo(ctx, domain.SetInput{
		ShortURL: shortURL,
		URLInfo: domain.URLInfo{
			OriginalURL: input.OriginalURL,
			CreatedAt:   time.Now(),
		},
	}); err != nil {
		return "", err
	}

	return shortURL, nil
}

func (i *CreateShortURLInput) validate() error {
	re := regexp.MustCompile(`^(http|https)://`)
	if !re.MatchString(i.OriginalURL) {
		return domain.ErrValidateProtocol
	}

	parsedURL, err := url.ParseRequestURI(i.OriginalURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return domain.ErrValidateURL
	}

	return nil
}

func generateShortURL(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[charIndex.Int64()]
	}
	return string(result)
}
