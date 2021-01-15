package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
	"github.com/taniwhy/ithub-backend/internal/pkg/util/uuid"
)

// Tag :
type Tag struct {
	TagID     string
	TagName   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewTag : Tagテーブルのレコードモデル生成
func NewTag(tagName string) *Tag {
	return &Tag{
		TagID:     uuid.New(),
		TagName:   tagName,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
}
