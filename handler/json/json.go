package json

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

// GetUserResJSON :
type GetUserResJSON struct {
	UserID          string      `json:"user_id" binding:"required"`
	UserName        null.String `json:"user_name" binding:"required"`
	Name            string      `json:"name" binding:"required"`
	TwitterUsername null.String `json:"twitter_username" binding:"required"`
	GithubUsername  null.String `json:"github_username" binding:"required"`
	UserText        null.String `json:"user_text" binding:"required"`
	UserIcon        null.String `json:"user_icon" binding:"required"`
	CreatedAt       time.Time   `json:"created_at" binding:"required"`
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

// GetNoteResJSON :
type GetNoteResJSON struct {
	NoteID    string      `json:"user_id" binding:"required"`
	UserName  null.String `json:"user_name" binding:"required"`
	NoteTitle string      `json:"note_title" binding:"required"`
	NoteText  string      `json:"note_text" binding:"required"`
	CreatedAt time.Time   `json:"created_at" binding:"required"`
}

// GetFollowsResJSON :
type GetFollowsResJSON struct {
	UserName  null.String `json:"user_name" binding:"required"`
	Name      string      `json:"name" binding:"required"`
	UserText  null.String `json:"user_text" binding:"required"`
	UserIcon  null.String `json:"user_icon" binding:"required"`
	Following bool        `json:"following" binding:"required"`
	CreatedAt time.Time   `json:"created_at" binding:"required"`
}

// GetTagsResJSON : タグ取得リクエストボディ
type GetTagsResJSON struct {
	TagID     string      `json:"tag_id" binding:"required"`
	TagName   string      `json:"tag_name" binding:"required"`
	TagIcon   null.String `json:"tag_icon" binding:"required"`
	CreatedAt time.Time   `json:"created_at" binding:"required"`
}

// CreateTagReqJSON : 新規タグ作成リクエストボディ
type CreateTagReqJSON struct {
	TagName string `json:"tag_name" binding:"required"`
	TagIcon string `json:"tag_icon"`
}
