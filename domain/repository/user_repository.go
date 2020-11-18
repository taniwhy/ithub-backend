package repository

import "github.com/taniwhy/ithub-backend/domain/model"

// IUserRepository : ユーザーのリポジトリ
type IUserRepository interface {
	FindByID(id string) (*model.User, error)
	Insert(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
}
