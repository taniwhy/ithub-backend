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
)

const googleURL string = "https://www.googleapis.com/oauth2/v1/userinfo"

// IAuthHandler : インターフェース
type IAuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct{}

// NewAuthHandler : GoogleOAuth認証ハンドラの生成
func NewAuthHandler() IAuthHandler {
	return &authHandler{}
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

func (aH *authHandler) Login(c *gin.Context) {
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
	SuccessResponser(c, gU)
}
