package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/repository"
)

// IFollowHandler :
type IFollowHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type followHandler struct {
	followRepository repository.IFollowRepository
}

// NewFollowHandler : フォローハンドラの生成
func NewFollowHandler(fR repository.IFollowRepository) IFollowHandler {
	return &followHandler{followRepository: fR}
}

func (h *followHandler) GetByID(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *followHandler) Create(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *followHandler) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
