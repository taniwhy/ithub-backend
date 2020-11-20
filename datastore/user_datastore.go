package datastore

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/datastore/errors"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/util/clock"
)

type userDatastore struct {
	db *gorm.DB
}

// NewUserDatastore : ユーザーデータストアの生成
func NewUserDatastore(db *gorm.DB) repository.IUserRepository {
	return &userDatastore{db}
}

func (d *userDatastore) FindByID(id string) (*model.User, error) {
	u := &model.User{}
	err := d.db.Where("user_id = ?", id).First(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}
	return u, nil
}

func (d *userDatastore) FindDeletedByID(id string) (*model.User, error) {
	u := &model.User{}
	err := d.db.Unscoped().Where("user_id = ?", id).First(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}
	return u, nil
}

func (d *userDatastore) FindByName(name string) (*model.User, error) {
	u := &model.User{}
	err := d.db.Where("user_name = ?", name).First(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.ErrRecordNotFound{}
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}
	return u, nil
}

func (d *userDatastore) Insert(user *model.User) error {
	return d.db.Create(user).Error
}

func (d *userDatastore) Update(user *model.User) error {
	return d.db.Model(&user).Where("user_id = ?", user.UserID).Updates(&user).Error
}

func (d *userDatastore) Delete(id string) error {
	user := model.User{}
	err := d.db.
		Model(&user).Where("user_id = ?", id).
		Update("deleted_at", sql.NullTime{Time: clock.Now(), Valid: true}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{}
	}
	if err != nil {
		return errors.ErrDatabase{Detail: err.Error()}
	}
	return nil
}

func (d *userDatastore) Restore(id string) error {
	user := model.User{}
	err := d.db.
		Model(&user).Unscoped().Where("user_id = ?", id).
		Update("deleted_at", sql.NullTime{Valid: false}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{}
	}
	if err != nil {
		return errors.ErrDatabase{Detail: err.Error()}
	}
	return nil
}
