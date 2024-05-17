package some_repo

import "github.com/sitnikovik/go-grpc-api-template/internal/client/pg"

// Repository is the interface that wraps the basic repository methods
type Repository interface {
}

// repository implements the Repository interface
type repository struct {
	db pg.Client
}

// NewRepository creates and returns a new repository
func NewRepository(db pg.Client) Repository {
	return &repository{db: db}
}
