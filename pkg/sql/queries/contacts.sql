-- name: ContactCreate :one
INSERT INTO
  contacts (email, name, phone, meta)
VALUES
  (?, ?, ?, ?)
RETURNING
  *;

-- name: ContactRead :one
SELECT
  *
FROM
  contacts
WHERE
  id = ?
LIMIT
  1;

-- name: ContactUpdate :exec
UPDATE
  contacts
SET
  email = ?,
  name = ?,
  phone = ?,
  meta = ?
WHERE
  id = ?;

-- name: ContactDelete :exec
DELETE FROM
  contacts
WHERE
  id = ?;

-- name: ContactList :many
SELECT
  *
FROM
  contacts
ORDER BY
  created_at DESC
LIMIT
  ?;
