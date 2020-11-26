package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/repository"
	"github.com/taniwhy/ithub-backend/domain/service"
	"github.com/taniwhy/ithub-backend/handler/errors"
	"github.com/taniwhy/ithub-backend/handler/util"
	"github.com/taniwhy/ithub-backend/middleware/auth"
)

const googleURL string = "https://www.googleapis.com/oauth2/v1/userinfo"

// IAuthHandler : インターフェース
type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	userRepository repository.IUserRepository
	userService    service.IUserService
}

// NewAuthHandler : GoogleOAuth認証ハンドラの生成
func NewAuthHandler(uR repository.IUserRepository, uS service.IUserService) IAuthHandler {
	return &authHandler{userRepository: uR, userService: uS}
}

type loginReqBody struct {
	IDToken string `json:"id_token" binding:"required"`
}

func (h *authHandler) Login(c *gin.Context) {
	reqBody := &loginReqBody{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, errors.ErrLoginReqBinding{IDToken: reqBody.IDToken}.Error())
		return
	}
	resp, err := http.Get(googleURL + "?access_token=" + reqBody.IDToken)
	if err != nil {
		log.Fatal(err)
		util.ErrorResponser(c, http.StatusBadRequest, "googleAPI access error")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		util.ErrorResponser(c, http.StatusBadRequest, errors.ErrInvalidToken{IDToken: reqBody.IDToken}.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		util.ErrorResponser(c, http.StatusBadRequest, "googleAPI response read error")
		return
	}
	gU := model.GoogleUser{}
	if err := json.Unmarshal(body, &gU); err != nil {
		log.Fatal(err)
		util.ErrorResponser(c, http.StatusBadRequest, "googleAPI response binding error")
		return
	}
	ok, err := h.userService.IsExist(gU.ID)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	if ok {
		ok, err := h.userService.IsDeleted(gU.ID)
		if err != nil {
			util.ErrorResponser(c, http.StatusBadRequest, err.Error())
			return
		}
		if ok {
			// アカウント復旧
			if err := h.userRepository.Restore(gU.ID); err != nil {
				util.ErrorResponser(c, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			// 新規登録
			newUser := model.NewUser(gU.ID, gU.Name, gU.Picture, gU.Email)
			if err := h.userRepository.Insert(newUser); err != nil {
				util.ErrorResponser(c, http.StatusBadRequest, err.Error())
				return
			}
		}
	}
	// ログイン
	u, err := h.userRepository.FindByID(gU.ID)
	if err != nil {
		util.ErrorResponser(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken := auth.GenerateAccessToken(u.UserID, u.UserName.String, u.IsAdmin)
	session := sessions.Default(c)
	session.Set("_token", accessToken)
	session.Save()
	util.SuccessMessageResponser(c, "ok")
}

func (h *authHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	util.SuccessMessageResponser(c, "ok")
}
