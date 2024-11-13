package config

import "os"

type Config struct {
	DatabasePath   string
	MigrationsPath string
}

func Load() *Config {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "data/tasks.db"
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		migrationsPath = "migrations/"
	}

	return &Config{
		DatabasePath:   dbPath,
		MigrationsPath: migrationsPath,
	}
}
