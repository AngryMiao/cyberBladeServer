package cache

import (
	"angrymiao-ai/config"
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var (
	Cache *cache
)

type cache struct {
	Client *redis.Client
	Prefix string
}

func Init(c *config.Config) {
	Client := NewRedis(c.RedisCache)
	Cache = &cache{
		Client: Client,
		Prefix: c.RedisCache.Service,
	}
}

func NewRedis(c *config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	return rdb
}

func KeyGenerate(prefix, key string) string {
	var b strings.Builder
	b.WriteString(Cache.Prefix)
	b.WriteString("::")
	b.WriteString(prefix)

	if key != "" {
		b.WriteString(":")
		b.WriteString(key)
	}

	return b.String()
}

func Get(key string) (string, error) {
	var ctx = context.Background()
	val, err := Cache.Client.Get(ctx, key).Result()
	return val, err
}

func Set(key string, value interface{}, expiration time.Duration) (string, error) {
	var ctx = context.Background()
	val, err := Cache.Client.Set(ctx, key, value, expiration).Result()
	return val, err
}

func SetEX(key string, value interface{}, ttl time.Duration) (string, error) {
	var ctx = context.Background()
	val, err := Cache.Client.SetEX(ctx, key, value, ttl).Result()
	return val, err
}

func SetNX(key string, value interface{}, ttl time.Duration) (bool, error) {
	var ctx = context.Background()
	val, err := Cache.Client.SetNX(ctx, key, value, ttl).Result()
	return val, err
}

func Del(key ...string) (int64, error) {
	var ctx = context.Background()
	val, err := Cache.Client.Del(ctx, key...).Result()
	return val, err
}

func TTL(key string) (time.Duration, error) {
	var ctx = context.Background()
	val, err := Cache.Client.TTL(ctx, key).Result()
	return val, err
}

func Incr(key string) (int64, error) {
	var ctx = context.Background()
	val, err := Cache.Client.Incr(ctx, key).Result()
	return val, err
}

func Expire(key string, ttl time.Duration) (bool, error) {
	var ctx = context.Background()
	val, err := Cache.Client.Expire(ctx, key, ttl).Result()
	return val, err
}

func HMSet(key string, values ...interface{}) (bool, error) {
	var ctx = context.Background()
	return Cache.Client.HMSet(ctx, key, values).Result()
}

func HDel(key string, fields ...string) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.HDel(ctx, key, fields...).Result()
}

func LPush(key string, values ...interface{}) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.LPush(ctx, key, values).Result()
}

func RPush(key string, values ...interface{}) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.RPush(ctx, key, values).Result()
}

func LRange(key string, start, end int64) ([]string, error) {
	var ctx = context.Background()
	return Cache.Client.LRange(ctx, key, start, end).Result()
}

func LRem(key string, count int64, value interface{}) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.LRem(ctx, key, count, value).Result()
}

func LAll(key string) ([]string, error) {
	var ctx = context.Background()
	return Cache.Client.LRange(ctx, key, 0, -1).Result()
}

func HGetAll(key string) (map[string]string, error) {
	var ctx = context.Background()
	return Cache.Client.HGetAll(ctx, key).Result()
}

func HGet(key, field string) (string, error) {
	var ctx = context.Background()
	return Cache.Client.HGet(ctx, key, field).Result()
}

func HSet(key string, values ...interface{}) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.HSet(ctx, key, values).Result()
}

func HIncrBy(key string, fields string, number int64) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.HIncrBy(ctx, key, fields, number).Result()
}

func EvalSha(sha string, keys []string, args ...interface{}) (interface{}, error) {
	var ctx = context.Background()
	return Cache.Client.EvalSha(ctx, sha, keys, args).Result()
}

func Exists(key string) (int64, error) {
	var ctx = context.Background()
	return Cache.Client.Exists(ctx, key).Result()
}
