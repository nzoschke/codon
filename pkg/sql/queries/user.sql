-- name: UserCreate :one
INSERT INTO
  users (email, name)
VALUES
  (?, ?)
RETURNING
  *;

-- name: UserRead :one
SELECT
  *
FROM
  users
WHERE
  id = ?
LIMIT
  1;

-- name: UserUpdate :exec
UPDATE
  users
SET
  email = ?,
  name = ?
WHERE
  id = ?;

-- name: UserDelete :exec
DELETE FROM
  users
WHERE
  id = ?;
