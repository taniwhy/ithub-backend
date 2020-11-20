package model

import (
	//"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
)

// Note_Tag :
type Note_Tag struct {
	NoteID    string
	TagID     string
	TagName   string
	IsMain    bool
	CreatedAt time.Time
}

// NewNote_Tag : Tagテーブルのレコードモデル生成
func NewNote_Tag(nID, tID string) *Note_Tag {
	return &Note_Tag{
		NoteID:    nID,
		TagID:     tID,
		IsMain:    false,
		CreatedAt: clock.Now(),
	}
}
