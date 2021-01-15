-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    user_id TEXT NOT NULL,
    user_name TEXT UNIQUE,
    name TEXT NOT NULL,
    twitter_username TEXT,
    github_username TEXT,
    user_icon TEXT,
    user_text TEXT,
    email TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(user_id),
    UNIQUE(email, user_name)
);

-- +migrate Down
DROP TABLE IF EXISTS users;