package support

import (
	"url-shortener/internal/shortener/adapters"
	"url-shortener/internal/shortener/domain"
	"url-shortener/internal/shortener/mocks"
	"url-shortener/internal/shortener/service"

	"github.com/redis/go-redis/v9"
)

func NewService(opts ...Option) service.Service {
	deps := &dependencies{}
	deps.setDefaults()
	for _, opt := range opts {
		opt(deps)
	}

	var svc service.Service = service.New(
		deps.redis,
		deps.urlLen,
	)
	return svc
}

type Option func(*dependencies)

type dependencies struct {
	redis  domain.RedisRepository
	urlLen int
}

func (d *dependencies) setDefaults() {
	d.urlLen = 8
	d.redis = mocks.NewMockRedisRepository(d.urlLen)
}

func WithRedisRepository(redis *redis.Client, expirationInDays int) Option {
	return func(d *dependencies) {
		d.redis = adapters.NewRedisRepository(redis, expirationInDays)
	}
}

func WithURLLen(urlLen int) Option {
	return func(d *dependencies) {
		d.urlLen = urlLen
	}
}
