package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
	"github.com/taniwhy/ithub-backend/internal/pkg/util/uuid"
)

// Comment :
type Comment struct {
	CommentID string
	UserName  string
	NoteID    string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewComment : Commentテーブルのレコードモデル生成
func NewComment(UserName, noteID, comment string) *Comment {
	return &Comment{
		CommentID: uuid.New(),
		UserName:  UserName,
		NoteID:    noteID,
		Comment:   comment,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
}
