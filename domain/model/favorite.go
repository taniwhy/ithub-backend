package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/package/util/clock"
	"github.com/taniwhy/ithub-backend/package/util/uuid"
)

// Favorite :
type Favorite struct {
	FavoriteID string
	NoteID     string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}

// NewFavorite : Favoriteテーブルのレコードモデル生成
func NewFavorite(userID, noteID string) *Favorite {
	return &Favorite{
		FavoriteID: uuid.UuID(),
		UserID:     userID,
		NoteID:     noteID,
		CreatedAt:  clock.Now(),
		UpdatedAt:  clock.Now(),
		DeletedAt:  sql.NullTime{Valid: false},
	}
}
