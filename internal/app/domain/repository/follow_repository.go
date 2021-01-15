package repository

import (
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
)

// IFollowRepository :
type IFollowRepository interface {
	FindFollowsByName(name string) ([]*model.Follow, error)
	FindFollowersByName(name string) ([]*model.Follow, error)
	Insert(follow *model.Follow) error
	Delete(name, target string) error
	FollowCount(name string) (int, int, error)
}
