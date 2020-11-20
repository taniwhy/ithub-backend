package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
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
func NewFollow(fID, uID, fouID string) *Follow {
	return &Follow{
		FollowID:     fID,
		UserID:       uID,
		FollowUserID: fouID,
		CreatedAt:    clock.Now(),
		UpdatedAt:    clock.Now(),
		DeletedAt:    sql.NullTime{Time: clock.Now(), Valid: false},
	}
}
