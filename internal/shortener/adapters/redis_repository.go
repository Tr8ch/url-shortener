package adapters

import (
	"context"
	"time"

	"url-shortener/internal/shortener/domain"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	rd         *redis.Client
	expiration time.Duration
}

var _ domain.RedisRepository = (*RedisRepository)(nil)

func NewRedisRepository(rd *redis.Client, expirationInDays int) *RedisRepository {
	return &RedisRepository{
		rd:         rd,
		expiration: time.Duration(expirationInDays*24) * time.Hour,
	}
}

func (r *RedisRepository) SetURLInfo(ctx context.Context, input domain.SetInput) error {
	if err := r.rd.HSet(ctx, input.ShortURL, input.URLInfo).Err(); err != nil {
		return errors.WithStack(err)
	}

	if err := r.rd.Expire(ctx, input.ShortURL, r.expiration).Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *RedisRepository) GetURLInfo(ctx context.Context, shortURL string) (*domain.URLInfo, error) {
	hget := r.rd.HGetAll(ctx, shortURL)
	if err := hget.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, domain.ErrShortURLNotFound
		}
		return nil, errors.WithStack(err)
	}

	var urlInfo domain.URLInfo
	if err := hget.Scan(&urlInfo); err != nil {
		return nil, errors.WithStack(err)
	}
	return &urlInfo, nil
}

func (r *RedisRepository) SetOriginalURL(ctx context.Context, input domain.SetOriginalURLInput) error {
	if err := r.rd.Set(ctx, input.OriginalURL, input.ShortURL, r.expiration).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *RedisRepository) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	shortURL, err := r.rd.Get(ctx, originalURL).Result()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return shortURL, nil
}

func (r *RedisRepository) GetURLs(ctx context.Context) ([]domain.URLs, error) {
	originalURLs, err := r.rd.Keys(ctx, "http*").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, domain.ErrURLsNotFound
		}
		return nil, errors.WithStack(err)
	}

	urls := make([]domain.URLs, 0, len(originalURLs))
	for _, originalURL := range originalURLs {
		shortURL, err := r.rd.Get(ctx, originalURL).Result()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		urls = append(urls, domain.URLs{
			OriginalURL: originalURL,
			ShortURL:    shortURL,
		})
	}

	return urls, nil
}

func (r *RedisRepository) Exists(ctx context.Context, url string) (bool, error) {
	exists, err := r.rd.Exists(ctx, url).Result()
	if err != nil {
		return false, errors.WithStack(err)
	}
	return exists > 0, nil
}

func (r *RedisRepository) DeleteURL(ctx context.Context, url string) error {
	err := r.rd.Del(ctx, url).Err()
	return errors.WithStack(err)
}
