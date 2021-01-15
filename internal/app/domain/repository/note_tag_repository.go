package repository

import "github.com/taniwhy/ithub-backend/internal/app/domain/model"

// INoteTagRepository :
type INoteTagRepository interface {
	FindByID(id string) ([]*model.NoteTag, error)
	Insert(tag *model.NoteTag) error
	Delete(id string) error
}
