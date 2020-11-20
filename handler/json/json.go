package json

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

// UserJSON :
type UserJSON struct {
	UserID          string      `json:"user_id" binding:"required"`
	UserName        null.String `json:"user_name" binding:"required"`
	Name            string      `json:"name" binding:"required"`
	TwitterUsername null.String `json:"twitter_username" binding:"required"`
	GithubUsername  null.String `json:"github_username" binding:"required"`
	UserText        null.String `json:"user_text" binding:"required"`
	UserIcon        null.String `json:"user_icon" binding:"required"`
	CreatedAt       time.Time   `json:"created_at" binding:"required"`
}

// GetUserResJSON :
type GetUserResJSON struct {
	User UserJSON `json:"user" binding:"required"`
}

// UpdateUserReqJSON : ユーザー更新リクエストボディ
type UpdateUserReqJSON struct {
	UserName        string      `json:"user_name" binding:"required"`
	Name            string      `json:"name" binding:"required"`
	TwitterUsername null.String `json:"twitter_username" binding:"required"`
	GithubUsername  null.String `json:"github_username" binding:"required"`
	UserText        null.String `json:"user_text" binding:"required"`
	UserIcon        null.String `json:"user_icon" binding:"required"`
}
