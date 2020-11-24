package repository

import "github.com/taniwhy/ithub-backend/domain/model"

// ITagRepository :
type ITagRepository interface {
	FindList() ([]*model.Tag, error)
	Insert(tag *model.Tag) error
	Update(tag *model.Tag) error
	Delete(id string) error
}