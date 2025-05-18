-- name: SessionCreate :one
INSERT INTO
  SESSION (id, user_id, expires_at)
VALUES
  (?, ?, ?)
RETURNING
  *;

-- name: SessionGet :one
SELECT
  SESSION.id,
  SESSION.user_id,
  SESSION.expires_at
FROM
  SESSION
  INNER JOIN user ON user.id = SESSION.user_id
WHERE
  SESSION.id = ?;

-- name: SessionDelete :exec
DELETE FROM
  SESSION
WHERE
  id = ?;

-- name: SessionDeleteUser :exec
DELETE FROM
  SESSION
WHERE
  user_id = ?;

-- name: SessionUpdate :exec
UPDATE
  SESSION
SET
  expires_at = ?
WHERE
  id = ?;
