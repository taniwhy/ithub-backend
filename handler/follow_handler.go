package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/handler/util"
	"github.com/taniwhy/ithub-backend/middleware/auth"
)

// IFollowHandler :
type IFollowHandler interface {
	GetFollows(c *gin.Context)
	GetFollowers(c *gin.Context)
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

func (h *followHandler) GetFollows(c *gin.Context) {
	name := c.Params.ByName("name")
	follows, err := h.followRepository.FindFollowsByName(name)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(c, follows)
}

func (h *followHandler) GetFollowers(c *gin.Context) {
	name := c.Params.ByName("name")
	followers, err := h.followRepository.FindFollowersByName(name)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(c, followers)
}

func (h *followHandler) Create(c *gin.Context) {
	target := c.Query("target")
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userName := claims["user_name"].(string)
	newFollow := model.NewFollow(userName, target)
	if err := h.followRepository.Insert(newFollow); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

func (h *followHandler) Delete(c *gin.Context) {
	target := c.Query("target")
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userName := claims["user_name"].(string)
	if err := h.followRepository.Delete(userName, target); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}
