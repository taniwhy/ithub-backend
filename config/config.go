package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	switch os.Getenv("GO_ENV") {
	case "dev":
		err := godotenv.Load(
			fmt.Sprintf("%s/src/github.com/taniwhy/ithub-backend/config/env/.env.dev", os.Getenv("GOPATH")))
		if err != nil {
			panic(err)
		}
	default:
		err := godotenv.Load("/workspace/config/env/.env.dev")
		if err != nil {
			panic(err)
		}
	}
}

// GetDatabaseConf :　データベースの接続情報の取得
func GetDatabaseConf() (dsn string) { return os.Getenv("DB_URL") }
