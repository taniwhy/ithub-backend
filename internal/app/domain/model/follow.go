package model

import (
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
}

// NewFollow : Tagテーブルのレコードモデル生成
func NewFollow(userName, followUserName string) *Follow {
	return &Follow{
		FollowID:       uuid.New(),
		UserName:       userName,
		FollowUserName: followUserName,
		CreatedAt:      clock.Now(),
	}
}
