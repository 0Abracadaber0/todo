CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    due_date TEXT,
    overdue INTEGER DEFAULT 0,
    completed INTEGER DEFAULT 0
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tasks_id ON tasks(id);
CREATE INDEX IF NOT EXISTS idx_tasks_due_date_overdue_completed ON tasks(due_date, overdue, completed);
