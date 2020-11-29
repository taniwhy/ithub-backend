package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/config"
)

// TokenAuth :　トークンで認証を行う
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(config.SecretKey)
			return b, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
		}
	}
}

// AdminAuth :　管理者判定で認証を行う
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["is_admin"].(bool) == false {
			c.JSON(http.StatusForbidden, gin.H{"message": http.StatusText(http.StatusForbidden)})
			c.Abort()
		}
	}
}
