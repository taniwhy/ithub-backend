package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
)

// User :
type User struct {
	ID        string
	UserID    sql.NullString
	UserName  string
	UserIcon  string
	Email     string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// NewUser : userテーブルのレコードモデル生成
func NewUser(gID, name, icon, email string) *User {
	return &User{
		ID:        gID,
		UserID:    sql.NullString{String: "", Valid: false},
		UserName:  name,
		UserIcon:  icon,
		Email:     email,
		IsAdmin:   false,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		DeletedAt: sql.NullTime{Time: clock.Now(), Valid: false},
	}
}
