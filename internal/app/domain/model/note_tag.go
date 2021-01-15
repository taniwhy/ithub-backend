package model

import (
	//"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
)

// NoteTag :
type NoteTag struct {
	NoteID    string
	TagName   string
	CreatedAt time.Time
}

// NewNoteTag : NoteTagテーブルのレコードモデル生成
func NewNoteTag(noteID, tagName string) *NoteTag {
	return &NoteTag{
		NoteID:    noteID,
		TagName:   tagName,
		CreatedAt: clock.Now(),
	}
}
