package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/internal/app/datastore/errors"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
)

type noteTagDatastore struct {
	db *gorm.DB
}

// NewnNoteTagDatastore : ユーザーデータストアの生成
func NewnNoteTagDatastore(db *gorm.DB) repository.INoteTagRepository {
	return &noteTagDatastore{db}
}

func (d *noteTagDatastore) FindByID(id string) ([]*model.NoteTag, error) {
	t := []*model.NoteTag{}

	err := d.db.Where("note_id = ?", id).Find(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return t, nil
}

func (d *noteTagDatastore) Insert(noteTag *model.NoteTag) error {
	return d.db.Create(noteTag).Error
}

func (d *noteTagDatastore) Delete(id string) error {
	return nil
}
