package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/internal/app/datastore/errors"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
)

type commentDatastore struct {
	db *gorm.DB
}

// NewCommentDatastore : ユーザーデータストアの生成
func NewCommentDatastore(db *gorm.DB) repository.ICommentRepository {
	return &commentDatastore{db}
}

func (d *commentDatastore) FindByNoteID(id string) ([]*model.Comment, error) {
	t := []*model.Comment{}

	err := d.db.Order("created_at asc").Where("note_id = ?", id).Find(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return t, nil
}

func (d *commentDatastore) Insert(comment *model.Comment) error {
	return d.db.Create(comment).Error
}
