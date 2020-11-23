package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
	"github.com/taniwhy/ithub-backend/util/uuid"
)

// Follow :
type Follow struct {
	FollowID     string
	UserID       string
	FollowUserID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}

// NewFollow : Tagテーブルのレコードモデル生成
func NewFollow(userID, followUserID string) *Follow {
	return &Follow{
		FollowID:     uuid.UuID(),
		UserID:       userID,
		FollowUserID: followUserID,
		CreatedAt:    clock.Now(),
		UpdatedAt:    clock.Now(),
		DeletedAt:    sql.NullTime{Valid: false},
	}
}
