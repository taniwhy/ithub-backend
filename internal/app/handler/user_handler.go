package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/auth"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
	"github.com/taniwhy/ithub-backend/internal/pkg/json"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
	"gopkg.in/guregu/null.v3"
)

// IUserHandler : インターフェース
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
	token := session.Get("_token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	ID := claims["sub"].(string)
	user, err := h.userRepository.FindByID(ID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}
	isyou := token == ID

	response.Success(c, json.GetUserResJSON{
		ID:            user.ID,
		UserID:        null.NewString(user.UserID.String, user.UserID.Valid),
		Name:          user.Name,
		TwitterLink:   null.NewString(user.TwitterLink.String, user.TwitterLink.Valid),
		GithubLink:    null.NewString(user.TwitterLink.String, user.GithubLink.Valid),
		UserText:      null.NewString(user.UserText.String, user.UserText.Valid),
		UserIcon:      null.NewString(user.UserIcon.String, user.UserIcon.Valid),
		CreatedAt:     user.CreatedAt,
		FollowCount:   0,
		FollowerCount: 0,
		CommentCount:  0,
		IsYou:         isyou,
	})
}

// GetByName : GetByName関数はユーザー情報を取得しレスポンスを返却します
func (h *userHandler) GetByName(c *gin.Context) {
	name := c.Params.ByName("name")
	user, err := h.userRepository.FindByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, json.GetUserResJSON{
		ID:            user.ID,
		UserID:        null.NewString(user.UserID.String, user.UserID.Valid),
		Name:          user.Name,
		TwitterLink:   null.NewString(user.TwitterLink.String, user.TwitterLink.Valid),
		GithubLink:    null.NewString(user.TwitterLink.String, user.GithubLink.Valid),
		UserText:      null.NewString(user.UserText.String, user.UserText.Valid),
		UserIcon:      null.NewString(user.UserIcon.String, user.UserIcon.Valid),
		CreatedAt:     user.CreatedAt,
		FollowCount:   0,
		FollowerCount: 0,
		CommentCount:  0,
	})
}

// Update : Update関数はユーザー情報を更新しレスポンスを返却します
func (h *userHandler) Update(c *gin.Context) {
	body := json.UpdateUserReqJSON{}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	session := sessions.Default(c)
	token := session.Get("_token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	ID := claims["sub"].(string)

	err = h.userRepository.Update(
		&model.User{
			ID:          ID,
			UserID:      sql.NullString{String: body.UserID, Valid: body.UserID != ""},
			Name:        body.Name,
			TwitterLink: sql.NullString{String: body.TwitterLink.String, Valid: body.TwitterLink.String != ""},
			GithubLink:  sql.NullString{String: body.GithubLink.String, Valid: body.GithubLink.String != ""},
			UserText:    sql.NullString{String: body.UserText.String, Valid: body.UserText.String != ""},
			UserIcon:    sql.NullString{String: body.UserIcon.String, Valid: body.UserIcon.String != ""},
			UpdatedAt:   clock.Now(),
		},
	)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}

// Delete : Delete関数はユーザー情報を削除しレスポンスを返却します
func (h *userHandler) Delete(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	userID := claims["sub"].(string)
	if err := h.userRepository.Delete(userID); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}
