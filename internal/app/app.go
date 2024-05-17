package app

import (
	"context"
	"log"
	"sync"

	"github.com/sitnikovik/go-grpc-api-template/internal/closer"

	"github.com/sitnikovik/go-grpc-api-template/internal/config"
)

// App Main application structure
type App struct {
	sp *serviceProvider // Service provider that stores app services to work with

	// TODO: Add gRPC server
}

// NewApp creates and returns a new application object
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts the application
func (a *App) Run() error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
		}
		a.Close()
		closer.CloseAll()
		closer.Wait()
	}()

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		log.Println("gRPC server is running...")
		if err := a.runGRPCServer(); err != nil {
			log.Fatalf("gRPC server error: %s\n", err.Error())
		}
		wg.Done()
	}()

	// TODO: add some other services here

	wg.Wait()

	return nil
}

// initDeps initializes dependencies for the application to work
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initServiceProvider инициализирует сервис-провайдер
func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider()

	return nil
}

// initGRPCServer инициализирует gRPC сервер
func (a *App) initGRPCServer(_ context.Context) error {
	panic("gRPC server initialization is not implemented yet!")
}

// runGRPCServer запускает gRPC сервер
func (a *App) runGRPCServer() error {
	panic("gRPC server initialization is not implemented yet!")
}

// Close closes the application
func (a *App) Close() {
	// TODO: implement closing logic
}
