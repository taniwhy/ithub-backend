package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/ithub-backend/config"
)

// GenerateAccessToken : アクセストークンの生成
func GenerateAccessToken(userID string, isAdmin bool) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      userID,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 1000).Unix(),
		"is_admin": isAdmin,
	})
	accessToken, err := token.SignedString([]byte(config.SignBytes))
	if err != nil {
		panic(err)
	}
	return accessToken
}

// GetTokenClaimsFromToken : トークンからトークンClaimを取得
func GetTokenClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.SignBytes, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// GetTokenClaimsFromRequest : コンテキストからトークンClaimを取得
func GetTokenClaimsFromRequest(c *gin.Context) (jwt.MapClaims, error) {
	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return config.SignBytes, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}