package config

import (
	"os"
	"strconv"
)

// Config is a struct that represents the configuration
type Config struct {
	env        string
	apiVersion int
	debug      bool
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	c := &Config{}

	c.env = os.Getenv("ENV")
	c.apiVersion, _ = strconv.Atoi(os.Getenv("API_VERSION"))
	c.debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	return c
}

// GetEnv returns the ENV value
func (c *Config) GetEnv() string {
	return c.env
}

// GetApiVersion returns the API_VERSION value
func (c *Config) GetApiVersion() int {
	return c.apiVersion
}

// GetDebug returns the DEBUG value
func (c *Config) GetDebug() bool {
	return c.debug
}
