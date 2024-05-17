package config

import (
	"errors"
	"os"
)

const grpcHostEnvName = "GRPC_HOST" // gRPC host

var _ GRPCConfig = (*grpcConfing)(nil)

// GRPCConfig represents the gRPC server configuration.
type GRPCConfig interface {
	// GetHost возвращает хост сервера gRPC
	GetHost() string
}

// grpcConfing represents the gRPC server configuration.
type grpcConfing struct {
	Host string // Хост сервера gRPC
}

// NewGRPCConfig creates and returns a new gRPC server configuration object.
func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if host == "" {
		return nil, errors.New("grpc host not found")
	}

	return &grpcConfing{
		Host: host,
	}, nil
}

// GetHost returns the gRPC server host.
func (c *grpcConfing) GetHost() string {
	return c.Host
}
