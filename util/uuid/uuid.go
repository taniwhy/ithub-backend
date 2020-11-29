package uuid

import (
	"github.com/google/uuid"
)

// New : UUIDの生成
var New = func() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err.Error())
	}

	return uuid.String()
}
