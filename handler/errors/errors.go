package errors

import (
	"fmt"
	"strings"
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
