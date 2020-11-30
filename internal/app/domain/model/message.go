package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
)

// Message :
type Message struct {
	MessageID string
	NoteID    string
	UserID    string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewMessage : Messageテーブルのレコードモデル生成
func NewMessage(cID, uID, nID, comment string) *Message {
	return &Message{
		MessageID: cID,
		UserID:    uID,
		NoteID:    nID,
		Comment:   comment,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Time: clock.Now(), Valid: false},
	}
}
