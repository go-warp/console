package config

import (
	"context"

	"github.com/joho/godotenv"
)

// Init инициализурует конфиг приложения
func Init(_ context.Context) error {
	return godotenv.Load(".env")
}
