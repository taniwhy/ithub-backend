package uuid

import (
	"github.com/google/uuid"
)

// UuID : UUIDのグローバル変数
var UuID = func() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err.Error())
	}
	return uuid.String()
}
