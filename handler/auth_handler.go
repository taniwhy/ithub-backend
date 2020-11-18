package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/domain/model"
	"github.com/taniwhy/ithub-backend/domain/service"
)

const googleURL string = "https://www.googleapis.com/oauth2/v1/userinfo"

// IAuthHandler : インターフェース
type IAuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	userService service.IUserService
}

// NewAuthHandler : GoogleOAuth認証ハンドラの生成
func NewAuthHandler(uS service.IUserService) IAuthHandler {
	return &authHandler{userService: uS}
}

type loginReqBody struct {
	IDToken string `json:"id_token" binding:"required"`
}

// ErrLoginReqBinding : TODO
type ErrLoginReqBinding struct {
	IDToken string
}

func (e ErrLoginReqBinding) Error() string {
	var errMsg []string
	if e.IDToken == "" {
		errMsg = append(errMsg, "id_token")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrInvalidToken : TODO
type ErrInvalidToken struct {
	IDToken string
}

func (e ErrInvalidToken) Error() string {
	return fmt.Sprintf("this token is invalid! - " + e.IDToken)
}

func (h *authHandler) Login(c *gin.Context) {
	reqBody := &loginReqBody{}
	if err := c.Bind(reqBody); err != nil {
		ErrorResponser(c, http.StatusBadRequest, ErrLoginReqBinding{IDToken: reqBody.IDToken}.Error())
	}
	resp, err := http.Get(googleURL + "?access_token=" + reqBody.IDToken)
	if err != nil {
		log.Fatal(err)
		ErrorResponser(c, http.StatusBadRequest, "googleAPI access error")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ErrorResponser(c, http.StatusBadRequest, ErrInvalidToken{IDToken: reqBody.IDToken}.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ErrorResponser(c, http.StatusBadRequest, "googleAPI response read error")
	}
	gU := model.GoogleUser{}
	if err := json.Unmarshal(body, &gU); err != nil {
		log.Fatal(err)
		ErrorResponser(c, http.StatusBadRequest, "googleAPI response binding error")
	}
	ok, err := h.userService.IsExist(gU.ID)
	if err != nil {
		ErrorResponser(c, http.StatusBadRequest, err.Error())
	}
	if !ok {
		// ログイン
	}
	// 新規登録
	// 前に削除したユーザーのアカウント復旧処理が必要
	newUser := model.NewUser(gU.ID, gU.Name, gU.Picture, gU.Email)
	SuccessResponser(c, newUser)
}
