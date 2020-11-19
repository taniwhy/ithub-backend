package handler

import (
	"github.com/gin-gonic/gin"
)

// IUserHandler : インターフェース
type IUserHandler interface {
	GetMe(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct{}

// NewUserHandler : ユーザーハンドラの生成
func NewUserHandler() IUserHandler {
	return &userHandler{}
}

func (h *userHandler) GetMe(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *userHandler) GetByID(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *userHandler) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *userHandler) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
