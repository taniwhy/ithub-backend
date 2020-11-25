package repository

import (
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/handler/json"
)

// IFollowRepository :
type IFollowRepository interface {
	FindFollowsByName(name string) ([]*json.GetFollowsResJSON, error)
	FindFollowersByName(name string) ([]*json.GetFollowsResJSON, error)
	Insert(follow *model.Follow) error
	Delete(	name, target string) error
}
