-- +migrate Up
CREATE TABLE IF NOT EXISTS note_tags (
    note_id TEXT NOT NULL,
    tag_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(note_id, tag_name),
    FOREIGN KEY(note_id) REFERENCES notes(note_id) ON UPDATE CASCADE ON DELETE
    SET
        NULL,
        FOREIGN KEY(tag_name) REFERENCES tags(tag_name) ON UPDATE CASCADE ON DELETE
    SET
        NULL
);

-- +migrate Down
DROP TABLE IF EXISTS note_tags;