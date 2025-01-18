// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, name, url, created_at, updated_at, user_id
`

type CreateFeedParams struct {
	ID        uuid.UUID
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const getFeed = `-- name: getFeed :one
SELECT id, name, url, created_at, updated_at, user_id 
FROM feeds
WHERE name = $1
`

func (q *Queries) getFeed(ctx context.Context, name string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed, name)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}
