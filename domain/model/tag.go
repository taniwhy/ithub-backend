package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
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
func NewTag(tID, tagName string) *Tag {
	return &Tag{
		TagID:     tID,
		TagName:   tagName,
		TagIcon:   sql.NullString{String: "", Valid: false},
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Time: clock.Now(), Valid: false},
	}
}
