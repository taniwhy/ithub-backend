package handler

import (
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
	"github.com/taniwhy/ithub-backend/util/uuid"
)

// INoteHandler :
type INoteHandler interface {
	Get(c *gin.Context)
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

// Get : Get関数は自ユーザー情報を取得しレスポンスを返却します
func (h *noteHandler) Get(c *gin.Context) {
	//panic("not implemented") // TODO: Implement
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	noteID := claims["sub"].(string)
	note, err := h.noteRepository.FindByID(noteID)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(
		c, json.GetNoteResJSON{
			Note: json.NoteJSON{
				NoteID: uuid.UuID(),
				UserID: note.userID,
				//NoteTitle: note.noteTitle,
				//NoteText:  note.noteText,
				CreatedAt: note.CreatedAt,
			},
		})
}

func (h *noteHandler) GetByID(c *gin.Context) {
	//panic("not implemented") // TODO: Implement
	id := c.Params.ByID("id")
	note, err := h.noteRepository.FindByID(id)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessDataResponser(c, note)
	/*
		util.SuccessDataResponser(c,
			json.GetNoteResJSON{
				Note: json.NoteJSON{
					NoteID:    uuid.UuID(),
					UserID:    note.userID,
					NoteTitle: note.noteTitle,
					NoteText:  note.noteText,
					CreatedAt: note.CreatedAt,
				},
			},
		)
	*/
}

func (h *noteHandler) Create(c *gin.Context) {
	//panic("not implemented") // TODO: Implement
	target := c.Query("target")
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	userName := claims["user_name"].(string)
	newNote := model.NewNote(userName, target)
	if err := h.noteRepository.Insert(newNote); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

// Update : Update関数はユーザー情報を更新しレスポンスを返却します
func (h *noteHandler) Update(c *gin.Context) {
	//panic("not implemented") // TODO: Implement
	body := json.UpdateNoteReqJSON{}
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
	noteID := claims["sub"].(string)
	if err := h.noteRepository.Update(&model.Note{
		NoteID:    uuid.UuID(),
		UserID:    body.userID,
		NoteTitle: body.noteTitle,
		NoteText:  body.noteText,
		CreatedAt: clock.Now(),
		UpdatedAt: clock.Now(),
		//DeletedAt: sql.NullTime{Valid: body.DeletedAt},
	}); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}

// Delete : Delete関数はユーザー情報を削除しレスポンスを返却します
func (h *noteHandler) Delete(c *gin.Context) {
	//panic("not implemented") // TODO: Implement
	session := sessions.Default(c)
	token := session.Get("_token")
	claims, err := auth.GetTokenClaimsFromToken(token.(string))
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	noteID := claims["sub"].(string)
	if err := h.noteRepository.Delete(noteID); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	util.SuccessMessageResponser(c, "ok")
}
