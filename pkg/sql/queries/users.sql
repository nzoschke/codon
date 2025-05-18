-- name: UserCreate :one
INSERT INTO
  users (email, password_hash)
VALUES
  (?, ?)
RETURNING
  id,
  email;

-- name: UserGet :one
SELECT
  id,
  email
FROM
  users
WHERE
  id = ?
LIMIT
  1;

-- name: UserGetPasswordHash :one
SELECT
  id,
  password_hash
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
  sessions.id,
  sessions.user_id,
  sessions.expires_at
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
