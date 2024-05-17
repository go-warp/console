package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// pgUserEnvName имя переменной окружения с пользователем БД
	pgUserEnvName = "PG_USER"
	// pgPasswordEnvName имя переменной окружения с паролем БД
	pgPasswordEnvName = "PG_PASSWORD"
	// pgHostEnvName имя переменной окружения с хостом БД
	pgHostEnvName = "PG_HOST"
	// pgPortEnvName имя переменной окружения с портом БД
	pgPortEnvName = "PG_PORT"
	// pgDBNameEnvName имя переменной окружения с именем БД
	pgDBNameEnvName = "PG_DATABASE"
)

// PGConfig Реализует конфиг БД PostgreSQL
type PGConfig interface {
	// GetDSN возвращает DSN для коннекта к БД
	GetDSN() string
}

// pgConfig Структура конфига БД PostgreSQL
type pgConfig struct {
	User     string // Пользователь
	Password string // Пароль
	Host     string // Хост
	Port     int    // Порт
	DBName   string // Имя БД
}

// NewPGConfig создает и возвращает объекта конфига БД PostgreSQL
func NewPGConfig() (PGConfig, error) {
	user := os.Getenv(pgUserEnvName)
	psw := os.Getenv(pgPasswordEnvName)
	host := os.Getenv(pgHostEnvName)
	port, err := strconv.Atoi(os.Getenv(pgPortEnvName))
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", pgPortEnvName, err)
	}
	dbName := os.Getenv(pgDBNameEnvName)

	return &pgConfig{
		User:     user,
		Password: psw,
		Host:     host,
		Port:     port,
		DBName:   dbName,
	}, nil
}

// GetDSN возвращает DSN для коннекта к БД
func (c *pgConfig) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}
