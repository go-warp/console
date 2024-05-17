package pg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ Client = (*client)(nil)

// Client реализует клиента для работы PostgreSQL
type Client interface {
	Close() error
	GetPG() PG
}

// client стуктура клиента для работы с PostgreSQL
type client struct {
	pg PG
}

// NewClient создает и возвращает клиента для работы с PostgreSQL
func NewClient(ctx context.Context, pgCfg *pgxpool.Config) (*client, error) {
	db, err := pgxpool.NewWithConfig(ctx, pgCfg)
	if err != nil {
		log.Fatalf("failed to get database connection: %s", err.Error())
	}

	return &client{
		pg: newPG(db),
	}, nil
}

// GetPG Возвращает клиента для PostgreSQL
func (c *client) GetPG() PG {
	return c.pg
}

// Close закрывает коннект к PostgreSQL
func (c *client) Close() error {
	return c.pg.Close()
}
