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

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $2, last_fetched_at = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * 
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;