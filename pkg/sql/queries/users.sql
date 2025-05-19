-- name: UserCreate :one
INSERT INTO
  users (email, password_hash)
VALUES
  (?, ?)
RETURNING
  *;

-- name: UserGet :one
SELECT
  email,
  id
FROM
  users
WHERE
  id = ?
LIMIT
  1;

-- name: UserGetByEmail :one
SELECT
  *
FROM
  users
WHERE
  email = ?
LIMIT
  1;

-- name: SessionCreate :one
INSERT INTO
  sessions (id, user_id, expires_at)
VALUES
  (?, ?, ?)
RETURNING
  *;

-- name: SessionGet :one
SELECT
  *
FROM
  sessions
WHERE
  sessions.id = ?;

-- name: SessionDelete :exec
DELETE FROM
  sessions
WHERE
  id = ?;

-- name: SessionDeleteUser :exec
DELETE FROM
  sessions
WHERE
  user_id = ?;

-- name: SessionUpdate :exec
UPDATE
  sessions
SET
  expires_at = ?
WHERE
  id = ?;
