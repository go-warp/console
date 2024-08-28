package config

import (
	"os"
	"strconv"
)

var (
	ENV = os.Getenv("ENV")
	DB_IDLE_MAX, _ = strconv.Atoi(os.Getenv("DB_IDLE_MAX"))
)
