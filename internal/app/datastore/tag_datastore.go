package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/internal/app/datastore/errors"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
)

type tagDatastore struct {
	db *gorm.DB
}

// NewTagDatastore : ユーザーデータストアの生成
func NewTagDatastore(db *gorm.DB) repository.ITagRepository {
	return &tagDatastore{db}
}

func (d *tagDatastore) FindList() ([]*model.Tag, error) {
	t := []*model.Tag{}

	err := d.db.Find(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return t, nil
}

func (d *tagDatastore) FindByName(name string) (*model.Tag, error) {
	t := model.Tag{}

	err := d.db.Where("tag_name = ?", name).First(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return &t, nil
}

func (d *tagDatastore) Insert(tag *model.Tag) error {
	return d.db.Create(tag).Error
}

func (d *tagDatastore) Update(tag *model.Tag) error {
	return nil
}

func (d *tagDatastore) Delete(id string) error {
	return nil
}
