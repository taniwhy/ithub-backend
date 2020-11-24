package repository

import "github.com/taniwhy/ithub-backend/domain/model"

// IFollowRepository :
type IFollowRepository interface {
	FindByID(id string) (*model.Note, error)
	Insert(note *model.Follow) error
	Delete(id string) error
}
