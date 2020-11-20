package repository

import "github.com/taniwhy/ithub-backend/domain/model"

// IUserRepository : ユーザーのリポジトリ
type IUserRepository interface {
	FindByID(id string) (*model.User, error)
	FindDeletedByID(id string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	Insert(user *model.User) error
	Update(user *model.User) error
	Delete(name string) error
	Restore(name string) error
}
