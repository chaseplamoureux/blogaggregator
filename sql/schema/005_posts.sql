-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL UNIQUE,
    url   TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID REFERENCES feeds(id) NOT NULL
);

-- +goose Down
DROP TABLE posts;