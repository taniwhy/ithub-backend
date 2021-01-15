package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/internal/app/datastore/errors"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
)

type noteDatastore struct {
	db *gorm.DB
}

// NewNoteDatastore : ユーザーデータストアの生成
func NewNoteDatastore(db *gorm.DB) repository.INoteRepository {
	return &noteDatastore{db}
}

func (d *noteDatastore) FindListByName(name string) ([]*model.Note, error) {
	t := []*model.Note{}

	err := d.db.Order("created_at desc").Where("user_name = ?", name).Find(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return t, nil
}

func (d *noteDatastore) FindByID(id string) (*model.Note, error) {
	n := &model.Note{}

	err := d.db.Where("note_id = ?", id).First(&n).Error
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (d *noteDatastore) Insert(note *model.Note) error {
	return d.db.Create(note).Error
}

func (d *noteDatastore) Update(note *model.Note) error {
	return nil
}

func (d *noteDatastore) Delete(id string) error {
	return nil
}

func (d *noteDatastore) PostCount(name string) (int, error) {
	var postCount int

	err := d.db.Model(&model.Note{}).Where("user_name = ?", name).Count(&postCount).Error

	if err != nil {
		return 0, errors.ErrDatabase{Detail: err.Error()}
	}

	return postCount, nil
}
