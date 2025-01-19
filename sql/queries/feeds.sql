-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedByURL :one
SELECT * 
FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT name, url, user_id
FROM feeds;