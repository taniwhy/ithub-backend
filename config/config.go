package config

import (
	"os"

	"github.com/joho/godotenv"
)

// SignBytes : 秘密鍵
var SignBytes string

func init() {
	switch os.Getenv("GO_ENV") {
	case "dev":
		err := godotenv.Load("./config/env/.env.dev")
		if err != nil {
			panic(err)
		}
	default:
	}
	SignBytes = os.Getenv("AUTHORIZE_RSA")
}

// GetDatabaseConf :　データベースの接続情報の取得
func GetDatabaseConf() (dsn string) { return os.Getenv("DB_URL") }
