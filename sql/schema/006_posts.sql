-- +goose Up

CREATE TABLE posts (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  url TEXT UNIQUE NOT NULL,
  published_at TIME NOT NULL,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  created_at TIME NOT NULL,
  updated_at TIME NOT NULL
);

-- +goose Down

DROP TABLE posts;