-- +goose Up

CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIME NOT NULL,
  updated_at TIME NOT NULL
);

-- +goose Down

DROP TABLE users;