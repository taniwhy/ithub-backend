package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
	"github.com/taniwhy/ithub-backend/util/uuid"
)

// Note :
type Note struct {
	NoteID    string
	UserID    string
	NoteTitle string
	NoteText  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewNote : Noteテーブルのレコードモデル生成
func NewNote(userID, noteTitle, noteText string) *Note {
	return &Note{
		NoteID:    uuid.UuID(),
		UserID:    userID,
		NoteTitle: noteTitle,
		NoteText:  noteText,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
}
