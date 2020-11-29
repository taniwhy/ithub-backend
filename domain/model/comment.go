package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
	"github.com/taniwhy/ithub-backend/util/uuid"
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
func NewComment(userID, noteID, comment string) *Comment {
	return &Comment{
		CommentID: uuid.New(),
		UserID:    userID,
		NoteID:    noteID,
		Comment:   comment,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
}
