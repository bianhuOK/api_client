package persistence

import (
	"time"

	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
	cache "github.com/patrickmn/go-cache"
)

type SqlLocalCache struct {
	cache *cache.Cache
}

type LocalCacheConfig struct {
	DefaultTTL      time.Duration
	CleanupInterval time.Duration
}

func ProviderSqlLocalCacheConfig() *LocalCacheConfig {
	defaultTTL := 5 * time.Minute
	defaultCleanupTTL := 10 * time.Minute
	return &LocalCacheConfig{
		DefaultTTL:      defaultTTL,
		CleanupInterval: defaultCleanupTTL,
	}
}

func NewSqlLocalCache(config *LocalCacheConfig) *SqlLocalCache {
	return &SqlLocalCache{
		cache: cache.New(config.DefaultTTL, config.CleanupInterval),
	}
}

func (c *SqlLocalCache) Get(key string) (*domain.SqlTemplate, bool) {
	value, found := c.cache.Get(key)
	if !found {
		return nil, false
	}
	template, ok := value.(*domain.SqlTemplate)
	if !ok {
		return nil, false
	}
	return template, true
}

func (c *SqlLocalCache) Set(key string, value *domain.SqlTemplate, ttl time.Duration) {
	c.cache.Set(key, value, ttl)
}
