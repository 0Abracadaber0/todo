-- name: CreateTask :exec
INSERT INTO tasks (title, description, due_date, overdue, completed)
VALUES (?, ?, ?, ?, ?);

-- name: GetTask :one
SELECT id, title, description, due_date, overdue, completed
FROM tasks
WHERE id = ?;

-- name: ListTasks :many
SELECT id, title, description, due_date, overdue, completed
FROM tasks;

-- name: GetLastID :one
SELECT last_insert_rowid()
FROM tasks;
