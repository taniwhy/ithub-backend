package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
)

// User :
type User struct {
	UserID          string
	UserName        sql.NullString
	Name            string
	TwitterUsername sql.NullString
	GithubUsername  sql.NullString
	UserIcon        string
	Email           string
	IsAdmin         bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}

// NewUser : userテーブルのレコードモデル生成
func NewUser(gID, name, icon, email string) *User {
	return &User{
		UserID:          gID,
		UserName:        sql.NullString{String: "", Valid: false},
		Name:            name,
		TwitterUsername: sql.NullString{String: "", Valid: false},
		GithubUsername:  sql.NullString{String: "", Valid: false},
		UserIcon:        icon,
		Email:           email,
		IsAdmin:         false,
		CreatedAt:       clock.Now(),
		UpdatedAt:       clock.Now(),
		DeletedAt:       sql.NullTime{Time: clock.Now(), Valid: false},
	}
}