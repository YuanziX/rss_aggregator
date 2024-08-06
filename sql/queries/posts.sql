-- name: CreatePost :one
INSERT INTO
    posts (
        id,
        created_at,
        updated_at,
        title,
        description,
        published_at,
        url,
        feed_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetAllPosts :many
SELECT
    *
FROM
    posts
ORDER BY
    published_at DESC;

-- name: GetPostsForUser :many
SELECT
    posts.*
FROM
    posts
    JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE
    feed_follows.user_id = ?
ORDER BY
    posts.published_at DESC
LIMIT
    ?;