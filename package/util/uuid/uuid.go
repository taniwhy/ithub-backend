package uuid

import (
	"github.com/google/uuid"
)

// UuID : UUID
var UuID = func() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err.Error())
	}

	return uuid.String()
}
