package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
	"github.com/taniwhy/ithub-backend/internal/pkg/json"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
)

// ITagHandler : インターフェース
type ITagHandler interface {
	GetList(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type tagHandler struct {
	tagRepository repository.ITagRepository
}

// NewTagHandler : タグハンドラの生成
func NewTagHandler(tR repository.ITagRepository) ITagHandler {
	return &tagHandler{tagRepository: tR}
}

type getTagsResponse struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (h *tagHandler) GetList(c *gin.Context) {
	tags, err := h.tagRepository.FindList()
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	res := []getTagsResponse{}

	for _, t := range tags {
		r := getTagsResponse{
			ID:   t.TagID,
			Name: t.TagName,
		}
		res = append(res, r)
	}

	response.Success(c, res)
}

func (h *tagHandler) Create(c *gin.Context) {
	body := json.CreateTagReqJSON{}

	if err := c.BindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	newTag := model.NewTag(body.TagName)

	err := h.tagRepository.Insert(newTag)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *tagHandler) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *tagHandler) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
