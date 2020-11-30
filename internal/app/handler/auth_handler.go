package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/internal/app/domain/model"
	"github.com/taniwhy/ithub-backend/internal/app/domain/repository"
	"github.com/taniwhy/ithub-backend/internal/app/domain/service"
	"github.com/taniwhy/ithub-backend/internal/app/middleware/auth"
	"github.com/taniwhy/ithub-backend/internal/pkg/error"
	jsonmodel "github.com/taniwhy/ithub-backend/internal/pkg/json"
	"github.com/taniwhy/ithub-backend/internal/pkg/response"
	"gopkg.in/go-playground/validator.v9"
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
	validate       *validator.Validate
}

// NewAuthHandler : GoogleOAuth認証ハンドラの生成
func NewAuthHandler(uR repository.IUserRepository, uS service.IUserService) IAuthHandler {
	return &authHandler{
		userRepository: uR, userService: uS,
		validate: validator.New(),
	}
}

func (h *authHandler) Login(c *gin.Context) {
	reqBody := &jsonmodel.LoginRequest{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, error.GetMsg(error.ERROR))
		return
	}

	err := h.validate.Struct(reqBody)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	resp, err := http.Get(googleURL + "?access_token=" + reqBody.IDToken)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, error.GetMsg(error.ERROR_ACCESS_GOOGLE_API_FAIL))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response.Error(c, http.StatusBadRequest, error.ERROR, error.GetMsg(error.ERROR_AUTH_TOKEN_INVALID))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, error.GetMsg(error.ERROR_READ_GOOGLE_API_RESPONSE_FAIL))
		return
	}

	gU := model.GoogleUser{}
	if err := json.Unmarshal(body, &gU); err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	err = h.validate.Struct(gU)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	isExist, err := h.userService.IsExist(gU.ID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	if isExist {
		isDeleted, err := h.userService.IsDeleted(gU.ID)
		if err != nil {
			response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
			return
		}
		if isDeleted {
			// アカウント復旧
			if err := h.userRepository.Restore(gU.ID); err != nil {
				response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
				return
			}
		} else {
			// 新規登録
			newUser := model.NewUser(gU.ID, gU.Name, gU.Picture, gU.Email)
			if err := h.userRepository.Insert(newUser); err != nil {
				response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
				return
			}
		}
	}
	// ログイン
	u, err := h.userRepository.FindByID(gU.ID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, error.ERROR, err.Error())
		return
	}

	accessToken := auth.GenerateAccessToken(u.UserID, u.UserName.String, u.IsAdmin)

	session := sessions.Default(c)
	session.Set("_token", accessToken)
	session.Save()

	response.Success(c, nil)
}

func (h *authHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	response.Success(c, nil)
}
