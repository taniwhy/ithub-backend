package repository

import "github.com/taniwhy/ithub-backend/internal/app/domain/model"

// ITagRepository :
type ITagRepository interface {
	FindList() ([]*model.Tag, error)
	FindByName(name string) (*model.Tag, error)
	Insert(tag *model.Tag) error
	Update(tag *model.Tag) error
	Delete(id string) error
}
