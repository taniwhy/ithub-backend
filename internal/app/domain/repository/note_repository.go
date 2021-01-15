package repository

import "github.com/taniwhy/ithub-backend/internal/app/domain/model"

// INoteRepository :
type INoteRepository interface {
	FindListByName(name string) ([]*model.Note, error)
	FindByID(id string) (*model.Note, error)
	Insert(note *model.Note) error
	Update(note *model.Note) error
	Delete(noteID string) error
	PostCount(name string) (int, error)
}
