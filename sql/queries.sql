-- name: CreateTask :one
INSERT INTO tasks (title, description, due_date, overdue, completed)
VALUES (?, ?, ?, ?, ?)
RETURNING id;


-- name: GetTask :one
SELECT id, title, description, due_date, overdue, completed
FROM tasks
WHERE id = ?;

-- name: GetTasks :many
SELECT id, title, description, due_date, overdue, completed
FROM tasks;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;

-- name: UpdateTask :exec
UPDATE tasks
SET title = ?,
    description = ?,
    due_date = ?
WHERE id = ?;

-- name: CompleteTask :exec
UPDATE tasks
SET completed = true
WHERE id = ?;

-- name: MarkOverdueTasks :exec
UPDATE tasks
SET overdue = 1
WHERE due_date < CURRENT_DATE AND completed == 0 AND overdue == 0;

