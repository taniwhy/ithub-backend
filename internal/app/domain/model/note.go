package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
	"github.com/taniwhy/ithub-backend/internal/pkg/util/uuid"
)

// Note :
type Note struct {
	NoteID    string
	UserName  string
	NoteTitle string
	NoteText  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewNote : Noteテーブルのレコードモデル生成
func NewNote(UserName, noteTitle, noteText string) *Note {
	return &Note{
		NoteID:    uuid.New(),
		UserName:  UserName,
		NoteTitle: noteTitle,
		NoteText:  noteText,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
}
