package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/internal/app/datastore/errors"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
)

type followDatastore struct {
	db *gorm.DB
}

// NewfollowDatastore : ユーザーデータストアの生成
func NewFollowDatastore(db *gorm.DB) repository.IFollowRepository {
	return &followDatastore{db}
}

func (d *followDatastore) FindFollowsByName(name string) ([]*model.Follow, error) {
	t := []*model.Follow{}

	err := d.db.Order("created_at desc").Where("user_name = ?", name).Find(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return t, nil
}

func (d *followDatastore) FindFollowersByName(name string) ([]*model.Follow, error) {
	t := []*model.Follow{}

	err := d.db.Order("created_at desc").Where("follow_user_name = ?", name).Find(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDatabase{Detail: err.Error()}
	}

	return t, nil
}

func (d *followDatastore) Insert(follow *model.Follow) error {
	return d.db.Create(follow).Error
}

func (d *followDatastore) Delete(name, target string) error {
	f := model.Follow{}
	err := d.db.Where("user_name = ? AND follow_user_name = ?", name, target).Delete(&f).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		return errors.ErrDatabase{Detail: err.Error()}
	}
	return nil
}

func (d *followDatastore) FollowCount(name string) (int, int, error) {
	var followCount, followerCount int

	err := d.db.Model(&model.Follow{}).Where("user_name = ?", name).Count(&followCount).Error
	err = d.db.Model(&model.Follow{}).Where("follow_user_name = ?", name).Count(&followerCount).Error

	if err != nil {
		return 0, 0, errors.ErrDatabase{Detail: err.Error()}
	}

	return followCount, followerCount, nil
}
