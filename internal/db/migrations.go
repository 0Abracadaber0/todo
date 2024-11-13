package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RunMigrations(db *sql.DB, migrationsPath string) error {
	if err := mkDir(migrationsPath); err != nil {
		return err
	}

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("reading migrations directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".up.sql") {
			continue
		}

		migrationPath := filepath.Join(migrationsPath, file.Name())
		sqlContent, _ := os.ReadFile(migrationPath)
		_, err = db.Exec(string(sqlContent))
		if err != nil {
			return fmt.Errorf("failed to run migration %s: %s", file.Name(), err)
		}
	}

	return nil
}
