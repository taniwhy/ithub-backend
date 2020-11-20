package errors

import (
	"fmt"
	"strings"

	"github.com/taniwhy/ithub-backend/handler/json"
)

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

// ErrUserUpdateReqBinding : TODO
type ErrUserUpdateReqBinding struct {
	Body json.UpdateUserReqJSON
}

func (e ErrUserUpdateReqBinding) Error() string {
	var errMsg []string
	if e.Body.UserName == "" {
		errMsg = append(errMsg, "user_name")
	}
	if e.Body.Name == "" {
		errMsg = append(errMsg, "name")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}
