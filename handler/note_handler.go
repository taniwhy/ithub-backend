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
	"github.com/taniwhy/ithub-backend/util/clock"
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
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	res := []json.GetNoteResJSON{}
	for _, n := range note {
		r := json.GetNoteResJSON{
			NoteID:    n.NoteID,
			UserName:  n.UserID,
			NoteTitle: n.NoteTitle,
			NoteText:  n.NoteText,
			CreatedAt: n.CreatedAt,
		}
		res = append(res, r)
	}
	util.SuccessDataResponser(c, res)
}

func (h *noteHandler) GetByID(c *gin.Context) {
	id := c.Params.ByName("id")
	note, err := h.noteRepository.FindByID(id)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(c, note)
}

func (h *noteHandler) Create(c *gin.Context) {
	body := json.CreateNoteReqJSON{}
	if err := c.BindJSON(&body); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
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
	newNote := model.NewNote(userID, body.NoteTitle, body.NoteText)
	err = h.noteRepository.Insert(newNote)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

func (h *noteHandler) Update(c *gin.Context) {
	noteID := c.Params.ByName("id")
	body := json.UpdateNoteReqJSON{}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
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
	if err := h.noteRepository.Update(&model.Note{
		NoteID:    noteID,
		UserID:    userID,
		NoteTitle: body.NoteTitle,
		NoteText:  body.NoteText,
		UpdatedAt: clock.Now(),
	}); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

func (h *noteHandler) Delete(c *gin.Context) {
	noteID := c.Params.ByName("id")
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userID := claims["sub"].(string)
	if err := h.noteRepository.Delete(userID, noteID); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}
