package json

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

// LoginRequest :
type LoginRequest struct {
	IDToken string `json:"id_token" validate:"required"`
}


// GetUserResJSON :
type GetUserResJSON struct {
	UserID          null.String      `json:"user_id" binding:"required"`
	Name            string           `json:"name" binding:"required"`
	IconLink        null.String      `json:"icon_link" binding:"required"`
	GithubUsername  null.String      `json:"github_link" binding:"required"`
	TwitterUsername null.String      `json:"twitter_link" binding:"required"`
	UserText        null.String      `json:"user_text" binding:"required"`
	FollowCount     int              `json:"follow_count" binding:"required"`
	FollowerCount   int              `json:"follower_count" binding:"required"`
	FollowTags      []GetTagsResJSON `json:"follow_tags" binding:"required"`
	IsYou           bool             `json:"is_you" binding:"required"`
	CreatedAt       time.Time        `json:"created_at" binding:"required"`
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
	NoteID    string    `json:"user_id" binding:"required"`
	UserName  string    `json:"user_name" binding:"required"`
	NoteTitle string    `json:"note_title" binding:"required"`
	NoteText  string    `json:"note_text" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

// CreateNoteReqJSON : 新規ノート作成リクエストボディ
type CreateNoteReqJSON struct {
	NoteTitle string `json:"note_title" binding:"required"`
	NoteText  string `json:"note_text" binding:"required"`
}

// UpdateNoteReqJSON : 新規ノート更新リクエストボディ
type UpdateNoteReqJSON struct {
	NoteTitle string `json:"note_title" binding:"required"`
	NoteText  string `json:"note_text" binding:"required"`
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
