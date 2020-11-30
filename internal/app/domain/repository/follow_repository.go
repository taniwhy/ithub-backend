package repository

import (
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/pkg/json"
)

// IFollowRepository :
type IFollowRepository interface {
	FindFollowsByName(name string) ([]*json.GetFollowsResJSON, error)
	FindFollowersByName(name string) ([]*json.GetFollowsResJSON, error)
	Insert(follow *model.Follow) error
	Delete(name, target string) error
}
