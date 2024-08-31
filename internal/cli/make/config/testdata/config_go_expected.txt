package config

import (
	"os"
)

// Config is a struct that represents the configuration
type Config struct {
	foo string
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	c := &Config{}

	c.foo = os.Getenv("FOO")

	return c
}

// Foo returns the FOO value
func (c *Config) Foo() string {
	return c.foo
}
