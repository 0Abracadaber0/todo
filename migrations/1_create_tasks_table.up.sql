CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    due_date TEXT,
    overdue INTEGER DEFAULT 0,
    completed INTEGER DEFAULT 0
);
