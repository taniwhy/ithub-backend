-- +migrate Up
CREATE TABLE IF NOT EXISTS comments (
    comment_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    note_id TEXT NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(comment_id),
    FOREIGN KEY(user_name) REFERENCES users(user_name) ON UPDATE CASCADE ON DELETE
    SET
        NULL,
        FOREIGN KEY(note_id) REFERENCES notes(note_id) ON UPDATE CASCADE ON DELETE
    SET
        NULL
);

-- +migrate Down
DROP TABLE IF EXISTS note_tags;