package repository

import (
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/pkg/json"
)

// ICommentRepository :
type ICommentRepository interface {
	FindByID(id string) ([]*json.GetFollowsResJSON, error)
	Insert(comment *model.Comment) error
	Update(comment *model.Comment) error
	Delete(userID, noteID string) error
}
