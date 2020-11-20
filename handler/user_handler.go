package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/handler/json"
	"github.com/taniwhy/ithub-backend/handler/util"
	"github.com/taniwhy/ithub-backend/middleware/auth"
	"gopkg.in/guregu/null.v3"
)

// IUserHandler : インターフェース
type IUserHandler interface {
	GetMe(c *gin.Context)
	GetByUserName(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct {
	userRepository repository.IUserRepository
}

// NewUserHandler : ユーザーハンドラの生成
func NewUserHandler(uR repository.IUserRepository) IUserHandler {
	return &userHandler{userRepository: uR}
}

func (h *userHandler) GetMe(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	userID := claims["sub"].(string)
	user, err := h.userRepository.FindByID(userID)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	util.SuccessResponser(
		c, json.GetUserResJSON{
			User: json.UserJSON{
				UserID:          user.UserID,
				UserName:        null.NewString(user.UserName.String, user.UserName.Valid),
				Name:            user.Name,
				TwitterUsername: null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
				GithubUsername:  null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
				UserText:        null.NewString(user.UserText.String, user.UserText.Valid),
				UserIcon:        null.NewString(user.UserIcon.String, user.UserIcon.Valid),
				CreatedAt:       user.CreatedAt,
			},
		})
}

func (h *userHandler) GetByUserName(c *gin.Context) {
	userName := c.Params.ByName("name")
	user, err := h.userRepository.FindByUserName(userName)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	util.SuccessResponser(
		c, json.GetUserResJSON{
			User: json.UserJSON{
				UserID:          user.UserID,
				UserName:        null.NewString(user.UserName.String, user.UserName.Valid),
				Name:            user.Name,
				TwitterUsername: null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
				GithubUsername:  null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
				UserText:        null.NewString(user.UserText.String, user.UserText.Valid),
				UserIcon:        null.NewString(user.UserIcon.String, user.UserIcon.Valid),
				CreatedAt:       user.CreatedAt,
			},
		})
}

func (h *userHandler) Update(c *gin.Context) {
	body := json.UpdateUserReqJSON{}
	if err := c.BindJSON(&body); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	userID := claims["sub"].(string)
	user, err := h.userRepository.FindByID(userID)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	if err := h.userRepository.Update(&model.User{
		UserID: user.UserID,
	}); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
}

func (h *userHandler) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
