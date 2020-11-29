package repository

import (
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/handler/json"
)

// ICommentRepository :
type ICommentRepository interface {
	FindByID(id string) ([]*json.GetFollowsResJSON, error)
	Insert(comment *model.Comment) error
	Update(comment *model.Comment) error
	Delete(userID, noteID string) error
}
