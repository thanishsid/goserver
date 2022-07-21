package rediscache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type CacheStore struct {
	Client *redis.Client
	Cache  *cache.Cache
}

func NewCacheStore(address string, db int) (*CacheStore, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})

	status := rdb.Ping(context.Background())
	if status.Err() != nil {
		return nil, status.Err()
	}

	rdbCache := cache.New(&cache.Options{
		Redis:      rdb,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &CacheStore{
		Client: rdb,
		Cache:  rdbCache,
	}, nil
}
