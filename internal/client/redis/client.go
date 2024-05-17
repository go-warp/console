package redis

import (
	"context"
	"errors"
	"time"

	redisError "github.com/sitnikovik/go-grpc-api-template/internal/errors/client/redis"

	"github.com/go-redis/redis"
)

// DefaultExpiration is the default expiration time for keys.
const DefaultExpiration = 24 * time.Hour

// Client is the interface that wraps the basic Redis operations.
type Client interface {
	// Get returns the value of key.
	Get(ctx context.Context, key string) (string, bool, error)

	// Set sets the value of key with expiration.
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error

	// Del deletes one or more keys.
	Del(ctx context.Context, keys ...string) error

	// Exists returns if key exists.
	Exists(ctx context.Context, key string) (bool, error)

	// Expire sets a timeout on key.
	Expire(ctx context.Context, key string, expiration time.Duration) error

	// Keys finds all keys matching the given pattern.
	Keys(ctx context.Context, pattern string) ([]string, error)

	// Ping checks the connection to the Redis server.
	Ping() error

	// Close closes the client, releasing any open resources.
	Close() error
}

// client is the implementation of the Client interface.
type client struct {
	redis *redis.Client
}

// NewClient creates and returns a new Redis client.
func NewClient(address, password string, db int) Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return &client{redis: rdb}
}

// Get returns the value of key.
func (c *client) Get(_ context.Context, key string) (string, bool, error) {
	val, err := c.redis.Get(key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", false, nil
		}

		return "", false, redisError.WrapError(err)
	}

	return val, true, nil
}

// Set sets the value of key.
func (c *client) Set(_ context.Context, key string, value interface{}, exp time.Duration) error {
	if exp == 0 {
		exp = DefaultExpiration
	}

	return redisError.WrapError(c.redis.Set(key, value, exp).Err())
}

// Del deletes one or more keys.
func (c *client) Del(_ context.Context, keys ...string) error {
	return redisError.WrapError(c.redis.Del(keys...).Err())
}

// Exists returns if key exists.
func (c *client) Exists(_ context.Context, key string) (bool, error) {
	i, err := c.redis.Exists(key).Result()
	if err != nil {
		return false, redisError.WrapError(err)
	}

	return i != 0, nil
}

// Expire sets a timeout on key.
func (c *client) Expire(_ context.Context, key string, expiration time.Duration) error {
	err := c.redis.Expire(key, expiration).Err()
	if err != nil {
		return redisError.WrapError(err)
	}

	return nil
}

// Keys finds all keys matching the given pattern.
func (c *client) Keys(_ context.Context, pattern string) ([]string, error) {
	ss, err := c.redis.Keys(pattern).Result()
	if err != nil {
		return nil, redisError.WrapError(err)
	}

	return ss, nil
}

// Ping checks the connection to the Redis server.
func (c *client) Ping() error {
	return redisError.WrapError(c.redis.Ping().Err())
}

// Close closes the client, releasing any open resources.
func (c *client) Close() error {
	return redisError.WrapError(c.redis.Close())
}
