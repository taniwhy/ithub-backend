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
	"gopkg.in/guregu/null.v3"
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
	userRepository   repository.IUserRepository
}

// NewFollowHandler : フォローハンドラの生成
func NewFollowHandler(fR repository.IFollowRepository, uR repository.IUserRepository) IFollowHandler {
	return &followHandler{followRepository: fR, userRepository: uR}
}

type getFollowResponse struct {
	ID     string      `json:"id" binding:"required"`
	UserID null.String `json:"user_id" binding:"required"`
	Name   string      `json:"name" binding:"required"`
	Link   null.String `json:"icon_link" binding:"required"`
}

func (h *followHandler) GetFollows(c *gin.Context) {
	name := c.Params.ByName("name")

	follows, err := h.followRepository.FindFollowsByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	followUsers := []getFollowResponse{}

	for _, f := range follows {

		user, err := h.userRepository.FindByName(f.FollowUserName)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}

		u := getFollowResponse{
			ID:     user.UserID,
			UserID: null.NewString(user.UserName.String, user.UserName.Valid),
			Name:   user.Name,
			Link:   null.NewString(user.UserIcon.String, user.UserIcon.Valid),
		}
		followUsers = append(followUsers, u)
	}

	response.Success(c, followUsers)
}

func (h *followHandler) GetFollowers(c *gin.Context) {
	name := c.Params.ByName("name")

	followers, err := h.followRepository.FindFollowersByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	followUsers := []getFollowResponse{}

	for _, f := range followers {

		user, err := h.userRepository.FindByName(f.UserName)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}

		u := getFollowResponse{
			ID:     user.UserID,
			UserID: null.NewString(user.UserName.String, user.UserName.Valid),
			Name:   user.Name,
			Link:   null.NewString(user.UserIcon.String, user.UserIcon.Valid),
		}
		followUsers = append(followUsers, u)
	}

	response.Success(c, followUsers)
}

func (h *followHandler) Create(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	userName := claims["user_name"].(string)
	target := c.Query("target")

	if target == "" {
		response.Error(c, http.StatusBadRequest, error.ERROR, "target is required")
		return
	}

	if userName == target {
		response.Error(c, http.StatusBadRequest, error.ERROR, "error")
		return
	}

	newFollow := model.NewFollow(userName, target)
	if err := h.followRepository.Insert(newFollow); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *followHandler) Delete(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token").(string)

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
