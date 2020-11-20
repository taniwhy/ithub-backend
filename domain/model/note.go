package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
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
func NewNote(nID, uID, noteTitle, noteText string) *Note {
	return &Note{
		NoteID:    nID,
		UserID:    uID,
		NoteTitle: noteTitle,
		NoteText:  noteText,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Time: clock.Now(), Valid: false},
	}
}
