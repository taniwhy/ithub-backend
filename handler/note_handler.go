package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/repository"
)

// INoteHandler :
type INoteHandler interface {
	GetList(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type noteHandler struct {
	noteRepository repository.INoteRepository
}

// NewNoteHandler : ノートハンドラの生成
func NewNoteHandler(nR repository.INoteRepository) INoteHandler {
	return &noteHandler{noteRepository: nR}
}

func (h *noteHandler) GetList(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *noteHandler) GetByID(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *noteHandler) Create(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *noteHandler) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *noteHandler) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
