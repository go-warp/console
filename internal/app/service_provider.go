package app

import (
	"context"
	"log"

	"github.com/sitnikovik/go-grpc-api-template/internal/repository/some_repo"
	"github.com/sitnikovik/go-grpc-api-template/internal/service/some_service"

	"github.com/sitnikovik/go-grpc-api-template/internal/closer"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sitnikovik/go-grpc-api-template/internal/client/pg"

	"github.com/sitnikovik/go-grpc-api-template/internal/client/redis"
	"github.com/sitnikovik/go-grpc-api-template/internal/config"
)

// serviceProvider struct stored all app services
type serviceProvider struct {
	redisConfig config.Redis    // Redis config
	pgConfig    config.PGConfig // Redis config

	redisClient redis.Client // Redis client
	pgClient    pg.Client    // PostgreSQL client

	someService some_service.Service // Some service
	someRepo    some_repo.Repository // Some service
}

// newServiceProvider creates and returns a new service provider
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// SomeService returns the SomeService instance
func (sp *serviceProvider) SomeService() some_service.Service {
	if sp.someService == nil {
		repo := sp.SomeRepo()
		s := some_service.NewService(repo)
		sp.someService = s
	}

	return sp.someService
}

// SomeRepo returns the SomeRepo instance
func (sp *serviceProvider) SomeRepo() some_repo.Repository {
	if sp.someRepo == nil {
		db := sp.PgClient(context.Background())
		r := some_repo.NewRepository(db)
		sp.someRepo = r
	}

	return sp.someRepo
}

// RedisClient returns the Redis client instance
func (sp *serviceProvider) RedisClient() redis.Client {
	if sp.redisClient == nil {
		conf := sp.RedisConfig()
		c := redis.NewClient(conf.GetAddr(), conf.GetPassword(), conf.GetDB())
		if err := c.Ping(); err != nil {
			log.Fatalf("failed to ping redis: %s", err.Error())
		}

		sp.redisClient = c
	}

	return sp.redisClient
}

// RedisConfig returns the Redis config instance
func (sp *serviceProvider) RedisConfig() config.Redis {
	if sp.redisConfig == nil {
		c, err := config.NewRedis()
		if err != nil {
			log.Fatalf("redis config failed: %s", err.Error())
		}
		sp.redisConfig = c
	}

	return sp.redisConfig
}

// PgClient returns the PostgreSQL client instance(singleton)
func (sp *serviceProvider) PgClient(ctx context.Context) pg.Client {
	if sp.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(sp.PgConfig().GetDSN())
		if err != nil {
			log.Fatalf("failed to get database config: %s", err.Error())
		}

		c, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatalf("failed to get database client connection: %s", err.Error())
		}
		if err = c.GetPG().Ping(ctx); err != nil {
			log.Fatalf("failed to ping database: %s", err.Error())
		}
		closer.Add(c.Close)

		sp.pgClient = c
	}

	return sp.pgClient
}

// PgConfig returns the PostgreSQL config instance(singleton)
func (sp *serviceProvider) PgConfig() config.PGConfig {
	if sp.pgConfig == nil {
		c, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("postgresql config failed: %s", err.Error())
		}

		sp.pgConfig = c
	}

	return sp.pgConfig
}
