package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/handler/json"
	"github.com/taniwhy/ithub-backend/handler/util"
	"gopkg.in/guregu/null.v3"
)

// ITagHandler :
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

func (h *tagHandler) GetList(c *gin.Context) {
	tags, err := h.tagRepository.FindList()
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	res := []json.GetTagsResJSON{}
	for _, t := range tags {
		r := json.GetTagsResJSON{
			TagID:     t.TagID,
			TagName:   t.TagName,
			TagIcon:   null.NewString(t.TagIcon.String, t.TagIcon.Valid),
			CreatedAt: t.CreatedAt,
		}
		res = append(res, r)
	}
	util.SuccessDataResponser(c, res)
}

func (h *tagHandler) Create(c *gin.Context) {
	body := json.CreateTagReqJSON{}
	if err := c.BindJSON(&body); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	newTag := model.NewTag(body.TagName, body.TagIcon)
	fmt.Println(newTag)
	err := h.tagRepository.Insert(newTag)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

func (h *tagHandler) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *tagHandler) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
