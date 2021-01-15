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
	"github.com/taniwhy/ithub-backend/internal/pkg/json"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
	"github.com/taniwhy/ithub-backend/internal/pkg/util/clock"
	"gopkg.in/guregu/null.v3"
)

// INoteHandler : インターフェース
type INoteHandler interface {
	GetListByID(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type noteHandler struct {
	noteRepository    repository.INoteRepository
	userRepository    repository.IUserRepository
	tagRepository     repository.ITagRepository
	noteTagRepository repository.INoteTagRepository
}

// NewNoteHandler : ノートハンドラの生成
func NewNoteHandler(
	nR repository.INoteRepository,
	uR repository.IUserRepository,
	tR repository.ITagRepository,
	ntR repository.INoteTagRepository,
) INoteHandler {
	return &noteHandler{
		noteRepository:    nR,
		userRepository:    uR,
		tagRepository:     tR,
		noteTagRepository: ntR,
	}
}

type getNoteListResponse struct {
	ID            string            `json:"id" binding:"required"`
	User          getUserResponse   `json:"user" binding:"required"`
	Title         string            `json:"memo_title" binding:"required"`
	Tags          []getTagsResponse `json:"tags" binding:"required"`
	FavoriteCount int               `json:"favorite_count" binding:"required"`
	CommentCount  int               `json:"comment_count" binding:"required"`
	CreatedAt     time.Time         `json:"created_at" binding:"required"`
}

func (h *noteHandler) GetListByID(c *gin.Context) {
	name := c.Params.ByName("name")

	note, err := h.noteRepository.FindListByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	user, err := h.userRepository.FindByName(name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	notes := []getNoteListResponse{}

	for _, n := range note {
		noteTags, err := h.noteTagRepository.FindByID(n.NoteID)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}

		tagsResponse := []getTagsResponse{}

		for _, t := range noteTags {
			tag, err := h.tagRepository.FindByName(t.TagName)
			if err != nil {
				response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
				return
			}
			r := getTagsResponse{
				ID:   tag.TagID,
				Name: tag.TagName,
			}
			tagsResponse = append(tagsResponse, r)
		}

		n := getNoteListResponse{
			ID: n.NoteID,
			User: getUserResponse{
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
			},
			Title:         n.NoteTitle,
			Tags:          tagsResponse,
			FavoriteCount: 0,
			CommentCount:  0,
			CreatedAt:     n.CreatedAt,
		}
		notes = append(notes, n)
	}

	response.Success(c, notes)
}

type getNoteResponse struct {
	ID            string            `json:"id" binding:"required"`
	User          getUserResponse   `json:"user" binding:"required"`
	Title         string            `json:"memo_title" binding:"required"`
	Tags          []getTagsResponse `json:"tags" binding:"required"`
	FavoriteCount int               `json:"favorite_count" binding:"required"`
	CommentCount  int               `json:"comment_count" binding:"required"`
	CreatedAt     time.Time         `json:"created_at" binding:"required"`
	Markdown      string            `json:"markdown" binding:"required"`
}

func (h *noteHandler) GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	note, err := h.noteRepository.FindByID(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	user, err := h.userRepository.FindByName(note.UserName)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	noteTags, err := h.noteTagRepository.FindByID(note.NoteID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	tagsResponse := []getTagsResponse{}

	for _, t := range noteTags {
		tag, err := h.tagRepository.FindByName(t.TagName)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}
		r := getTagsResponse{
			ID:   tag.TagID,
			Name: tag.TagName,
		}
		tagsResponse = append(tagsResponse, r)
	}

	response.Success(c, getNoteResponse{
		ID: note.NoteID,
		User: getUserResponse{
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
		},
		Title:         note.NoteTitle,
		Tags:          tagsResponse,
		FavoriteCount: 0,
		CommentCount:  0,
		CreatedAt:     note.CreatedAt,
		Markdown:      note.NoteText,
	})
}

type createTags struct {
	Name string `json:"name"`
}

type createNoteRequest struct {
	Title    string       `json:"memo_title" binding:"required"`
	Tags     []createTags `json:"tags"`
	Markdown string       `json:"markdown" binding:"required"`
}

type createNoteResponse struct {
	ID string `json:"id" binding:"required"`
}

func (h *noteHandler) Create(c *gin.Context) {
	body := createNoteRequest{}

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

	UserName := claims["user_name"].(string)
	newNote := model.NewNote(UserName, body.Title, body.Markdown)

	err = h.noteRepository.Insert(newNote)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	for _, t := range body.Tags {
		tag, err := h.tagRepository.FindByName(t.Name)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}
		if tag == nil {
			newTag := model.NewTag(t.Name)
			h.tagRepository.Insert(newTag)
		}
		newNoteTag := model.NewNoteTag(newNote.NoteID, t.Name)
		h.noteTagRepository.Insert(newNoteTag)
	}

	response.Success(c, createNoteResponse{ID: newNote.NoteID})
}

func (h *noteHandler) Update(c *gin.Context) {
	body := json.UpdateNoteReqJSON{}

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

	noteID := c.Params.ByName("id")
	UserName := claims["sub"].(string)

	err = h.noteRepository.Update(
		&model.Note{
			NoteID:    noteID,
			UserName:  UserName,
			NoteTitle: body.NoteTitle,
			NoteText:  body.NoteText,
			UpdatedAt: clock.Now(),
		},
	)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *noteHandler) Delete(c *gin.Context) {

	noteID := c.Params.ByName("id")

	if err := h.noteRepository.Delete(noteID); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}
