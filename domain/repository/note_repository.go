package repository

import "github.com/taniwhy/ithub-backend/domain/model"

// INoteRepository :
type INoteRepository interface {
	FindByID(id string) (*model.Note, error)
	Insert(note *model.Note) error
	Update(note *model.Note) error
	Delete(id string) error
}
