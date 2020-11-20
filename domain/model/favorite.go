package model

import (
	"database/sql"
	"time"

	"github.com/taniwhy/ithub-backend/util/clock"
)

// Favorite :
type Favorite struct {
	FavoriteID      string
	NoteID          string
	UserID          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}

// NewFavorite : Favoriteテーブルのレコードモデル生成
func NewFavorite(favoID, uID, nID string) *Favorite {
	return &Favorite{
		FavoriteID:      favoID,
		UserID:          uID,
		NoteID:          nID,
		CreatedAt:       clock.Now(),
		UpdatedAt:       clock.Now(),
		DeletedAt:       sql.NullTime{Time: clock.Now(), Valid: false},
	}
}