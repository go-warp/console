package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	redisAddrEnvName     = "REDIS_ADDR"     // Address of the Redis server
	redisPasswordEnvName = "REDIS_PASSWORD" // Password for the Redis server
	redisDBEnvName       = "REDIS_DB"       // Database number for the Redis server
)

// Redis implements the Redis interface
type Redis interface {
	// GetAddr returns the address of the Redis server
	GetAddr() string
	// GetPassword returns the password for the Redis server
	GetPassword() string
	// GetDB returns the database number for the Redis server
	GetDB() int
}

// redis implements the Redis interface
type redis struct {
	Addr     string
	Password string
	DB       int
}

// NewRedis makes and returns a new Redis object
func NewRedis() (Redis, error) {
	addr := os.Getenv(redisAddrEnvName)
	password := os.Getenv(redisPasswordEnvName)
	db, err := strconv.Atoi(os.Getenv(redisDBEnvName))
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", redisDBEnvName, err)
	}

	return &redis{
		Addr:     addr,
		Password: password,
		DB:       db,
	}, nil
}

// GetAddr returns the address of the Redis server
func (c *redis) GetAddr() string {
	return c.Addr
}

// GetPassword returns the password for the Redis server
func (c *redis) GetPassword() string {
	return c.Password
}

// GetDB returns the database number for the Redis server
func (c *redis) GetDB() int {
	return c.DB
}
