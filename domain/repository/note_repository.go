package repository

import "github.com/taniwhy/ithub-backend/domain/model"

// INoteRepository :
type INoteRepository interface {
	FindList() ([]*model.Note, error)
	FindByID(id string) (*model.Note, error)
	Insert(note *model.Note) error
	Update(note *model.Note) error
	Delete(id string) error
}
