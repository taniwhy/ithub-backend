package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
	"github.com/taniwhy/ithub-backend/internal/pkg/util/uuid"
)

// Follow :
type Follow struct {
	FollowID       string
	UserName       string
	FollowUserName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
}

// NewFollow : Tagテーブルのレコードモデル生成
func NewFollow(userName, followUserName string) *Follow {
	return &Follow{
		FollowID:       uuid.New(),
		UserName:       userName,
		FollowUserName: followUserName,
		CreatedAt:      clock.Now(),
		UpdatedAt:      clock.Now(),
		DeletedAt:      sql.NullTime{Valid: false},
	}
}
