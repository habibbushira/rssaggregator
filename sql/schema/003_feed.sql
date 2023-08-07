-- +goose Up

CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIME NOT NULL,
  updated_at TIME NOT NULL
);

-- +goose Down

DROP TABLE feeds;