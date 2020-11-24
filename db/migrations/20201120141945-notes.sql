-- +migrate Up
CREATE TABLE IF NOT EXISTS notes (
    note_id TEXT NOT NULL,
    user_name TEXT,
    note_title TEXT NOT NULL,
    note_text TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(note_id),
    FOREIGN KEY(user_name) REFERENCES users(user_name) ON UPDATE CASCADE ON DELETE
    SET
        NULL
);

-- +migrate Down
DROP TABLE IF EXISTS notes;