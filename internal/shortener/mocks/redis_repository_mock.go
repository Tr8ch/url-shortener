package mocks

import (
	"context"
	"sync"
	"time"

	"url-shortener/internal/shortener/domain"
)

type MockRedisRepository struct {
	data            map[string]domain.URLInfo
	shortToOriginal map[string]string
	originalToShort map[string]string
	expiration      time.Duration
	mu              sync.RWMutex
}

var _ domain.RedisRepository = (*MockRedisRepository)(nil)

func NewMockRedisRepository(expirationInDays int) *MockRedisRepository {
	return &MockRedisRepository{
		data:            make(map[string]domain.URLInfo),
		shortToOriginal: make(map[string]string),
		originalToShort: make(map[string]string),
		expiration:      time.Duration(expirationInDays*24) * time.Hour,
	}
}

func (m *MockRedisRepository) SetURLInfo(ctx context.Context, input domain.SetInput) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[input.ShortURL] = input.URLInfo
	m.shortToOriginal[input.ShortURL] = input.URLInfo.OriginalURL
	m.originalToShort[input.URLInfo.OriginalURL] = input.ShortURL

	go func(shortURL string) {
		time.Sleep(m.expiration)
		m.mu.Lock()
		defer m.mu.Unlock()
		delete(m.data, shortURL)
		delete(m.shortToOriginal, shortURL)
	}(input.ShortURL)

	return nil
}

func (m *MockRedisRepository) GetURLInfo(ctx context.Context, shortURL string) (*domain.URLInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	urlInfo, exists := m.data[shortURL]
	if !exists {
		return nil, domain.ErrShortURLNotFound
	}

	return &urlInfo, nil
}

func (m *MockRedisRepository) SetOriginalURL(ctx context.Context, input domain.SetOriginalURLInput) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.originalToShort[input.OriginalURL] = input.ShortURL
	m.shortToOriginal[input.ShortURL] = input.OriginalURL

	return nil
}

func (m *MockRedisRepository) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	shortURL, exists := m.originalToShort[originalURL]
	if !exists {
		return "", domain.ErrShortURLNotFound
	}

	return shortURL, nil
}

func (m *MockRedisRepository) GetURLs(ctx context.Context) ([]domain.URLs, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	urls := make([]domain.URLs, 0, len(m.data))
	for originalURL, shortURL := range m.originalToShort {
		urls = append(urls, domain.URLs{
			OriginalURL: originalURL,
			ShortURL:    shortURL,
		})
	}

	if len(urls) == 0 {
		return nil, domain.ErrURLsNotFound
	}

	return urls, nil
}

func (m *MockRedisRepository) Exists(ctx context.Context, url string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, existsInData := m.data[url]
	_, existsInOriginal := m.originalToShort[url]
	_, existsInShort := m.shortToOriginal[url]

	return existsInData || existsInOriginal || existsInShort, nil
}

func (m *MockRedisRepository) DeleteURL(ctx context.Context, url string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.shortToOriginal[url]; exists {
		delete(m.data, url)
		delete(m.shortToOriginal, url)
		return nil
	}

	if _, exists := m.originalToShort[url]; exists {
		delete(m.originalToShort, url)
		return nil
	}

	return domain.ErrShortURLNotFound
}
