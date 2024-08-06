-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;
