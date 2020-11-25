-- +migrate Up
CREATE TABLE IF NOT EXISTS tags (
    tag_id TEXT NOT NULL,
    tag_name TEXT NOT NULL,
    tag_icon TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(tag_id),
    UNIQUE(tag_name)
);

-- +migrate Down
DROP TABLE IF EXISTS tags;