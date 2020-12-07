package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
)

// User :
type User struct {
	ID            string
	UserID        sql.NullString
	Name          string
	TwitterLink   sql.NullString
	GithubLink    sql.NullString
	UserText      sql.NullString
	UserIcon      sql.NullString
	Email         string
	IsAdmin       bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
	FollowCount   int
	FollowerCount int
	CommentCount  int
	IsYou         bool
}

// NewUser : userテーブルのレコードモデル生成
func NewUser(gID, name, icon, email string) *User {
	return &User{
		ID:            gID,
		UserID:        sql.NullString{String: "", Valid: false},
		Name:          name,
		TwitterLink:   sql.NullString{String: "", Valid: false},
		GithubLink:    sql.NullString{String: "", Valid: false},
		UserText:      sql.NullString{String: "", Valid: false},
		UserIcon:      sql.NullString{String: icon, Valid: true},
		Email:         email,
		IsAdmin:       false,
		CreatedAt:     clock.Now(),
		UpdatedAt:     clock.Now(),
		DeletedAt:     sql.NullTime{Valid: false},
		FollowCount:   0,
		FollowerCount: 0,
		CommentCount:  0,
		IsYou:         false,
	}
}
