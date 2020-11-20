package model

import (
	//"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
)

// NoteTag :
type NoteTag struct {
	NoteID    string
	TagID     string
	TagName   string
	IsMain    bool
	CreatedAt time.Time
}

// NewNote_Tag : Tagテーブルのレコードモデル生成
func NewNoteTag(nID, tID string) *NoteTag {
	return &NoteTag{
		NoteID:    nID,
		TagID:     tID,
		IsMain:    false,
		CreatedAt: clock.Now(),
	}
}
