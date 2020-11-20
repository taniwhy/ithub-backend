
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    user_id TEXT NOT NULL,
    user_name TEXT,
    name TEXT NOT NULL,
    twitter_username TEXT,
    github_username TEXT,
    user_text TEXT,
    user_icon TEXT,
    email TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(user_id),
    UNIQUE(user_id, email)
);
-- +migrate Down
DROP TABLE IF EXISTS users;