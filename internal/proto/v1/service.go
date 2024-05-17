package v1

import "github.com/sitnikovik/go-grpc-api-template/internal/service/some_service"

// Implementation is the implementation of the proto service
type Implementation struct {
	// TODO: Specify proto pkg service here!

	some_service.Service
}

// NewImplementation creates and returns a new implementation
func NewImplementation(s some_service.Service) *Implementation {
	return &Implementation{s}
}

// TODO: implement proto service methods here!
