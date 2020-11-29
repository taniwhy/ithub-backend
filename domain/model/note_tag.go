package model

import (
	//"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/package/util/clock"
)

// NoteTag :
type NoteTag struct {
	NoteID    string
	TagID     string
	TagName   string
	IsMain    bool
	CreatedAt time.Time
}

// NewNoteTag : NoteTagテーブルのレコードモデル生成
func NewNoteTag(noteID, tagID string, isMain bool) *NoteTag {
	return &NoteTag{
		NoteID:    noteID,
		TagID:     tagID,
		IsMain:    isMain,
		CreatedAt: clock.Now(),
	}
}
