package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/handler/errors"
	"github.com/taniwhy/ithub-backend/handler/json"
	"github.com/taniwhy/ithub-backend/handler/util"
	"github.com/taniwhy/ithub-backend/middleware/auth"
	"github.com/taniwhy/ithub-backend/util/clock"
	"gopkg.in/guregu/null.v3"
)

// IUserHandler : ユーザーハンドラのインターフェース
type IUserHandler interface {
	GetMe(c *gin.Context)
	GetByName(c *gin.Context)
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

// GetMe : GetMe関数は自ユーザー情報を取得しレスポンスを返却します
func (h *userHandler) GetMe(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := claims["sub"].(string)
	user, err := h.userRepository.FindByID(userID)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(c,
		json.GetUserResJSON{
			UserID:          user.UserID,
			UserName:        null.NewString(user.UserName.String, user.UserName.Valid),
			Name:            user.Name,
			TwitterUsername: null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
			GithubUsername:  null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
			UserText:        null.NewString(user.UserText.String, user.UserText.Valid),
			UserIcon:        null.NewString(user.UserIcon.String, user.UserIcon.Valid),
			CreatedAt:       user.CreatedAt,
		},
	)
}

// GetByName : GetByName関数はユーザー情報を取得しレスポンスを返却します
func (h *userHandler) GetByName(c *gin.Context) {
	name := c.Params.ByName("name")
	user, err := h.userRepository.FindByName(name)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(c,
		json.GetUserResJSON{
			UserID:          user.UserID,
			UserName:        null.NewString(user.UserName.String, user.UserName.Valid),
			Name:            user.Name,
			TwitterUsername: null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
			GithubUsername:  null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
			UserText:        null.NewString(user.UserText.String, user.UserText.Valid),
			UserIcon:        null.NewString(user.UserIcon.String, user.UserIcon.Valid),
			CreatedAt:       user.CreatedAt,
		},
	)
}

// Update : Update関数はユーザー情報を更新しレスポンスを返却します
func (h *userHandler) Update(c *gin.Context) {
	body := json.UpdateUserReqJSON{}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, errors.ErrUserUpdateReqBinding{Body: body}.Error())
		return
	}
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := claims["sub"].(string)
	if err := h.userRepository.Update(&model.User{
		UserID:          userID,
		UserName:        sql.NullString{String: body.UserName, Valid: body.UserName != ""},
		Name:            body.Name,
		TwitterUsername: sql.NullString{String: body.TwitterUsername.String, Valid: body.TwitterUsername.String != ""},
		GithubUsername:  sql.NullString{String: body.GithubUsername.String, Valid: body.GithubUsername.String != ""},
		UserText:        sql.NullString{String: body.UserText.String, Valid: body.UserText.String != ""},
		UserIcon:        sql.NullString{String: body.UserIcon.String, Valid: body.UserIcon.String != ""},
		UpdatedAt:       clock.Now(),
	}); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

// Delete : Delete関数はユーザー情報を削除しレスポンスを返却します
func (h *userHandler) Delete(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := claims["sub"].(string)
	if err := h.userRepository.Delete(userID); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}
