package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/auth"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
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
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, follows)
}

func (h *followHandler) GetFollowers(c *gin.Context) {
	name := c.Params.ByName("name")

	followers, err := h.followRepository.FindFollowersByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, followers)
}

func (h *followHandler) Create(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	userName := claims["user_name"].(string)
	target := c.Query("target")

	newFollow := model.NewFollow(userName, target)
	if err := h.followRepository.Insert(newFollow); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *followHandler) Delete(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	userName := claims["user_name"].(string)
	target := c.Query("target")

	if err := h.followRepository.Delete(userName, target); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}
