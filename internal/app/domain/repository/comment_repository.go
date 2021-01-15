package repository

import (
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
)

// ICommentRepository :
type ICommentRepository interface {
	FindByNoteID(id string) ([]*model.Comment, error)
	Insert(comment *model.Comment) error
}
