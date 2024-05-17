package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ PG = (*pg)(nil)

// Query Структура SQL-запроса к БД
type Query struct {
	Name     string // Имя запроса для дебага
	QueryRaw string // Сформированный SQL-запрос
}

// NewQuery создает и возвращает объект запроса к БД
func NewQuery(name string, queryRaw string) *Query {
	return &Query{
		Name:     name,
		QueryRaw: queryRaw,
	}
}

// QueryExecer Реализует методы запросов к БД
type QueryExecer interface {
	// ExecContext выполянет запрос
	ExecContext(ctx context.Context, q *Query, args ...interface{}) (pgconn.CommandTag, error)
	// QueryContext выполняет запрос и возвращает строки в ответе
	QueryContext(ctx context.Context, q *Query, args ...interface{}) (pgx.Rows, error)
	// QueryRowContext выполняет запрос и возвращает одну строку в ответе
	QueryRowContext(ctx context.Context, q *Query, args ...interface{}) pgx.Row
}

// Pinger Реализует базовый пингер к БД
type Pinger interface {
	// Ping пингует БД
	Ping(ctx context.Context) error
}

// PG Реализует взаимодействие с PostgreSQL
type PG interface {
	QueryExecer
	Pinger
	// Close закрывает пул коннектов к БД
	Close() error
}

// pg Стуктура для работы с БД PostgreSQL
type pg struct {
	pgxPool *pgxpool.Pool // Пул коннектов к БД
}

// newPG создает и возвращает клиента БД
func newPG(db *pgxpool.Pool) *pg {
	return &pg{
		pgxPool: db,
	}
}

// Close закрывает коннект к БД
func (p *pg) Close() error {
	p.pgxPool.Close()

	return nil
}

// Ping пингует БД
func (p *pg) Ping(ctx context.Context) error {
	return p.pgxPool.Ping(ctx)
}

// ExecContext выполянет запрос к БД
func (p *pg) ExecContext(ctx context.Context, q *Query, args ...interface{}) (pgconn.CommandTag, error) {
	return p.pgxPool.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext выполняет запрос и возвращает строки в ответе
func (p *pg) QueryContext(ctx context.Context, q *Query, args ...interface{}) (pgx.Rows, error) {
	return p.pgxPool.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext выполняет запрос и возвращает одну строку в ответе
func (p *pg) QueryRowContext(ctx context.Context, q *Query, args ...interface{}) pgx.Row {
	return p.pgxPool.QueryRow(ctx, q.QueryRaw, args...)
}
