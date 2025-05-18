-- name: SessionCreate :one
INSERT INTO
  user_session (id, user_id, expires_at)
VALUES
  (?, ?, ?)
RETURNING
  *;

-- name: SessionGet :one
SELECT
  user_session.id,
  user_session.user_id,
  user_session.expires_at
FROM
  user_session
  INNER JOIN user ON user.id = SESSION.user_id
WHERE
  user_session.id = ?;

-- name: SessionDelete :exec
DELETE FROM
  user_session
WHERE
  id = ?;

-- name: SessionDeleteUser :exec
DELETE FROM
  user_session
WHERE
  user_id = ?;

-- name: SessionUpdate :exec
UPDATE
  user_session
SET
  expires_at = ?
WHERE
  id = ?;
