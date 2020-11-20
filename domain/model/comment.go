package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
)

// Comment :
type Comment struct {
	CommentID string
	NoteID    string
	UserID    string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewComment : Commentテーブルのレコードモデル生成
func NewComment(cID, gID, nID, comment string) *Comment {
	return &Comment{
		CommentID: cID,
		UserID:    gID,
		NoteID:    nID,
		Comment:   comment,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Time: clock.Now(), Valid: false},
	}
}
