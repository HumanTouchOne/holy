package redis

import (
	"context"
	"github.com/HumanTouchOne/holy/components/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client     *redis.Client
	ctx        context.Context
	defaultExp time.Duration
}

var _ cache.ICacheE = (*RedisCache)(nil)

func NewRedisHandle(ctx context.Context, client *redis.Client) *RedisCache {
	return &RedisCache{
		client:     client,
		ctx:        ctx,
		defaultExp: time.Second * 3600 * 24,
	}
}
func (r *RedisCache) SetDefaultExp(exp time.Duration) {
	r.defaultExp = exp
}
func (r *RedisCache) Get(key string) ([]byte, error) {
	return r.client.Get(r.ctx, key).Bytes()
}
func (r *RedisCache) Set(key string, val []byte) error {
	return r.client.Set(r.ctx, key, val, r.defaultExp).Err()
}
func (r *RedisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
func (r *RedisCache) SetE(key string, value []byte, exp time.Duration) error {
	return r.client.Set(r.ctx, key, value, exp).Err()
}
