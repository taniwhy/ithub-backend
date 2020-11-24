package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
	"github.com/taniwhy/ithub-backend/util/uuid"
)

// Tag :
type Tag struct {
	TagID     string
	TagName   string
	TagIcon   sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewTag : Tagテーブルのレコードモデル生成
func NewTag(tagName, TagIcon string) *Tag {
	return &Tag{
		TagID:     uuid.UuID(),
		TagName:   tagName,
		TagIcon:   sql.NullString{String: TagIcon, Valid: TagIcon != ""},
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
}
