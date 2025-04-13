-- name: ContactCreate :one
INSERT INTO
  contacts (email, meta, name, phone)
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
  meta = ?,
  name = ?,
  phone = ?
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

-- name: ContactAge :one
SELECT
  CAST(meta ->> '$.age' AS INT) AS age
FROM
  contacts
WHERE
  id = ?
LIMIT
  1;
