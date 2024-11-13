package config

import "os"

type Config struct {
	DatabasePath string
}

func Load() *Config {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "data/tasks.db"
	}
	return &Config{
		DatabasePath: dbPath,
	}
}
