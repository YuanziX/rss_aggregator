-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows
WHERE user_id = ?;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE id = ? AND user_id = ?;