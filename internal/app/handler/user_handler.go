package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/auth"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
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
	userRepository   repository.IUserRepository
	followRepository repository.IFollowRepository
	noteRepository   repository.INoteRepository
}

// NewUserHandler : ユーザーハンドラの生成
func NewUserHandler(
	uR repository.IUserRepository,
	fR repository.IFollowRepository,
	nR repository.INoteRepository,
) IUserHandler {
	return &userHandler{
		userRepository:   uR,
		followRepository: fR,
		noteRepository:   nR,
	}
}

type followTags struct {
	TagID     string      `json:"id" binding:"required"`
	TagName   string      `json:"name" binding:"required"`
	TagIcon   null.String `json:"link" binding:"required"`
	CreatedAt time.Time   `json:"created_at" binding:"required"`
}

type getUserResponse struct {
	ID            string       `json:"id" binding:"required"`
	UserID        null.String  `json:"user_id" binding:"required"`
	Name          string       `json:"name" binding:"required"`
	IconLink      null.String  `json:"icon_link" binding:"required"`
	GithubLink    null.String  `json:"github_link" binding:"required"`
	TwitterLink   null.String  `json:"twitter_link" binding:"required"`
	UserText      null.String  `json:"user_text" binding:"required"`
	FollowCount   int          `json:"follow_count" binding:"required"`
	FollowerCount int          `json:"follower_count" binding:"required"`
	PostCount     int          `json:"post_count" binding:"required"`
	CommentCount  int          `json:"comment_count" binding:"required"`
	FollowTags    []followTags `json:"follow_tags" binding:"required"`
	IsYou         bool         `json:"is_you" binding:"required"`
	CreatedAt     time.Time    `json:"created_at" binding:"required"`
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

	userID := claims["sub"].(string)
	name := claims["user_name"].(string)
	user, err := h.userRepository.FindByID(userID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	FollowCount, FollowerCount, err := h.followRepository.FollowCount(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	postCount, err := h.noteRepository.PostCount(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, getUserResponse{
		ID:            user.UserID,
		UserID:        null.NewString(user.UserName.String, user.UserName.Valid),
		Name:          user.Name,
		IconLink:      null.NewString(user.UserIcon.String, user.UserIcon.Valid),
		GithubLink:    null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
		TwitterLink:   null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
		UserText:      null.NewString(user.UserText.String, user.UserText.Valid),
		FollowCount:   FollowCount,
		FollowerCount: FollowerCount,
		PostCount:     postCount,
		CommentCount:  0,
		IsYou:         true,
		CreatedAt:     user.CreatedAt,
	})
}

// GetByName : GetByName関数はユーザー情報を取得しレスポンスを返却します
func (h *userHandler) GetByName(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("_token")
	var userName string

	if token != nil {
		claims, err := auth.GetTokenClaimsFromToken(token.(string))
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}
		userName = claims["user_name"].(string)
	}

	name := c.Params.ByName("name")
	user, err := h.userRepository.FindByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	FollowCount, FollowerCount, err := h.followRepository.FollowCount(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	postCount, err := h.noteRepository.PostCount(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, getUserResponse{
		ID:            user.UserID,
		UserID:        null.NewString(user.UserName.String, user.UserName.Valid),
		Name:          user.Name,
		IconLink:      null.NewString(user.UserIcon.String, user.UserIcon.Valid),
		GithubLink:    null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
		TwitterLink:   null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
		UserText:      null.NewString(user.UserText.String, user.UserText.Valid),
		FollowCount:   FollowCount,
		FollowerCount: FollowerCount,
		PostCount:     postCount,
		CommentCount:  0,
		IsYou:         userName == user.UserID,
		CreatedAt:     user.CreatedAt,
	})
}

type updateUserRequest struct {
	UserID      string      `json:"user_id" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	IconLink    null.String `json:"icon_link" binding:"required"`
	TwitterLink null.String `json:"twitter_link" binding:"required"`
	GithubLink  null.String `json:"github_link" binding:"required"`
	UserText    null.String `json:"user_text" binding:"required"`
}

// Update : Update関数はユーザー情報を更新しレスポンスを返却します
func (h *userHandler) Update(c *gin.Context) {
	body := updateUserRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, "binding error")
		return
	}

	session := sessions.Default(c)
	token := session.Get("_token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, "token error")
		return
	}

	userID := claims["sub"].(string)

	updateUser := &model.User{
		UserID:          userID,
		UserName:        sql.NullString{String: body.UserID, Valid: body.UserID != ""},
		Name:            body.Name,
		TwitterUsername: sql.NullString{String: body.TwitterLink.String, Valid: body.TwitterLink.String != ""},
		GithubUsername:  sql.NullString{String: body.GithubLink.String, Valid: body.GithubLink.String != ""},
		UserText:        sql.NullString{String: body.UserText.String, Valid: body.UserText.String != ""},
		UserIcon:        sql.NullString{String: body.IconLink.String, Valid: body.IconLink.String != ""},
		UpdatedAt:       clock.Now(),
	}

	err = h.userRepository.Update(updateUser)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	FollowCount, FollowerCount, err := h.followRepository.FollowCount(claims["user_name"].(string))
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	accessToken := auth.GenerateAccessToken(updateUser.UserID, updateUser.UserName.String, updateUser.IsAdmin)

	session.Set("_token", accessToken)
	session.Save()

	response.Success(c, getUserResponse{
		UserID:        null.NewString(updateUser.UserName.String, updateUser.UserName.Valid),
		Name:          updateUser.Name,
		IconLink:      null.NewString(updateUser.UserIcon.String, updateUser.UserIcon.Valid),
		GithubLink:    null.NewString(updateUser.TwitterUsername.String, updateUser.GithubUsername.Valid),
		TwitterLink:   null.NewString(updateUser.TwitterUsername.String, updateUser.TwitterUsername.Valid),
		UserText:      null.NewString(updateUser.UserText.String, updateUser.UserText.Valid),
		FollowCount:   FollowCount,
		FollowerCount: FollowerCount,
		PostCount:     0,
		CommentCount:  0,
		IsYou:         true,
		CreatedAt:     updateUser.CreatedAt,
	})
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
