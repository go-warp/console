package some_service

import "github.com/sitnikovik/go-grpc-api-template/internal/repository/some_repo"

// Service is the interface that wraps some service methods
type Service interface {
	// Implement some methods
}

// service implements the Service interface
type service struct {
	repo some_repo.Repository
}

// NewService creates and returns a new service
func NewService(repo some_repo.Repository) Service {
	return &service{repo: repo}
}
