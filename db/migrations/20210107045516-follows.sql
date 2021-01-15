-- +migrate Up
CREATE TABLE IF NOT EXISTS follows (
    follow_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    follow_user_name TEXT,
    created_at TIMESTAMP NOT NULL,
    UNIQUE(follow_id),
    PRIMARY KEY(user_name, follow_user_name),
    FOREIGN KEY(user_name) REFERENCES users(user_name) ON UPDATE CASCADE ON DELETE
    SET
        NULL,
        FOREIGN KEY(follow_user_name) REFERENCES users(user_name) ON UPDATE CASCADE ON DELETE
    SET
        NULL
);

-- +migrate Down
DROP TABLE IF EXISTS follows;