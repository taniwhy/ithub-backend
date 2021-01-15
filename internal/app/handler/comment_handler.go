package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/auth"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
	"gopkg.in/guregu/null.v3"
)

// IcommentHandler :
type ICommentHandler interface {
	GetByNoteID(c *gin.Context)
	Create(c *gin.Context)
}

type commentHandler struct {
	commentRepository repository.ICommentRepository
	userRepository    repository.IUserRepository
}

// NewCommentHandler : フォローハンドラの生成
func NewCommentHandler(fR repository.ICommentRepository, uR repository.IUserRepository) ICommentHandler {
	return &commentHandler{commentRepository: fR, userRepository: uR}
}

type getCommentResponse struct {
	ID        string          `json:"id" binding:"required"`
	User      getUserResponse `json:"user" binding:"required"`
	Commnet   string          `json:"comment" binding:"required"`
	CreatedAt time.Time       `json:"created_at" binding:"required"`
}

func (h *commentHandler) GetByNoteID(c *gin.Context) {
	id := c.Params.ByName("id")

	comments, err := h.commentRepository.FindByNoteID(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	commentResponse := []getCommentResponse{}

	for _, comment := range comments {
		user, err := h.userRepository.FindByName(comment.UserName)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}

		u := getUserResponse{
			UserID:        null.NewString(user.UserName.String, user.UserName.Valid),
			Name:          user.Name,
			IconLink:      null.NewString(user.UserIcon.String, user.UserIcon.Valid),
			GithubLink:    null.NewString(user.TwitterUsername.String, user.GithubUsername.Valid),
			TwitterLink:   null.NewString(user.TwitterUsername.String, user.TwitterUsername.Valid),
			UserText:      null.NewString(user.UserText.String, user.UserText.Valid),
			FollowCount:   0,
			FollowerCount: 0,
			IsYou:         true,
			CreatedAt:     user.CreatedAt,
		}

		commentResponse = append(commentResponse, getCommentResponse{
			ID:        comment.CommentID,
			User:      u,
			Commnet:   comment.Comment,
			CreatedAt: comment.CreatedAt})
	}

	response.Success(c, commentResponse)
}

type createCommentRequest struct {
	Comment string `json:"comment" binding:"required"`
}

func (h *commentHandler) Create(c *gin.Context) {
	id := c.Params.ByName("id")
	body := createCommentRequest{}

	if err := c.BindJSON(&body); err != nil {
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

	userName := claims["user_name"].(string)

	newComment := model.NewComment(userName, id, body.Comment)
	if err := h.commentRepository.Insert(newComment); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}
