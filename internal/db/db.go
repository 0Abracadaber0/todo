package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(databasePath string) (*sql.DB, error) {
	// TODO: errors
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTasksTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

// TODO: migrations
func createTasksTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        due_date TEXT,
        overdue BOOLEAN DEFAULT FALSE,
        completed BOOLEAN DEFAULT FALSE
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу tasks: %w", err)
	}
	return nil
}
