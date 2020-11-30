package errors

import (
	"fmt"
)

// ErrDatabase : TODO
type ErrDatabase struct {
	Detail string
}

func (e ErrDatabase) Error() string {
	return fmt.Sprintf("database error! - " + e.Detail)
}

// ErrRecordNotFound : TODO
type ErrRecordNotFound struct{}

func (e ErrRecordNotFound) Error() string {
	return fmt.Sprintf("record notfound error!")
}
