package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/domain/repository"
)

// IUserService : インターフェース
type IUserService interface {
	IsAdmin(id string) (bool, error)
	IsExist(id string) (bool, error)
	IsDeleted(id string) (bool, error)
}

type userService struct {
	userRepository repository.IUserRepository
}

// NewUserService : UserServiceの生成
func NewUserService(uR repository.IUserRepository) IUserService {
	return &userService{
		userRepository: uR,
	}
}

// ユーザーが管理者であればtrueを返却
func (s *userService) IsAdmin(id string) (bool, error) {
	res, err := s.userRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}
	return res.IsAdmin, nil
}

// ユーザーが存在しなければtrueを返却
func (s *userService) IsExist(id string) (bool, error) {
	res, err := s.userRepository.FindByID(id)
	if gorm.IsRecordNotFoundError(err) {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	if res != nil && res.DeletedAt.Valid == false {
		return false, nil
	}
	return true, nil
}

// ユーザーが削除済みであればtrueを返却
func (s *userService) IsDeleted(id string) (bool, error) {
	res, err := s.userRepository.FindDeletedByID(id)
	if res != nil && res.DeletedAt.Valid == true {
		fmt.Println("test3")
		return true, nil
	}
	if err != nil {
		return false, err
	}
	fmt.Println("test4")
	return false, nil
}
