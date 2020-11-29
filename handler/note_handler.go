package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/handler/json"
	"github.com/taniwhy/ithub-backend/middleware/auth"
	"github.com/taniwhy/ithub-backend/package/error"
	"github.com/taniwhy/ithub-backend/package/response"
	"github.com/taniwhy/ithub-backend/package/util/clock"
)

// INoteHandler : インターフェース
type INoteHandler interface {
	GetList(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type noteHandler struct {
	noteRepository repository.INoteRepository
}

// NewNoteHandler : ノートハンドラの生成
func NewNoteHandler(nR repository.INoteRepository) INoteHandler {
	return &noteHandler{noteRepository: nR}
}

func (h *noteHandler) GetList(c *gin.Context) {
	note, err := h.noteRepository.FindList()
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	notes := []json.GetNoteResJSON{}

	for _, n := range note {
		n := json.GetNoteResJSON{
			NoteID:    n.NoteID,
			UserName:  n.UserID,
			NoteTitle: n.NoteTitle,
			NoteText:  n.NoteText,
			CreatedAt: n.CreatedAt,
		}
		notes = append(notes, n)
	}

	response.Success(c, notes)
}

func (h *noteHandler) GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	note, err := h.noteRepository.FindByID(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, note)
}

func (h *noteHandler) Create(c *gin.Context) {
	body := json.CreateNoteReqJSON{}

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

	userID := claims["sub"].(string)
	newNote := model.NewNote(userID, body.NoteTitle, body.NoteText)

	err = h.noteRepository.Insert(newNote)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
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
	userID := claims["sub"].(string)

	err = h.noteRepository.Update(
		&model.Note{
			NoteID:    noteID,
			UserID:    userID,
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
	session := sessions.Default(c)
	token := session.Get("_token").(string)

	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	userID := claims["sub"].(string)
	noteID := c.Params.ByName("id")

	if err := h.noteRepository.Delete(userID, noteID); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	response.Success(c, nil)
}
