package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// SecretKey : 秘密鍵
var SecretKey string

func init() {
	switch os.Getenv("GO_ENV") {
	case "dev":
		err := godotenv.Load("./configs/env/.env.dev")
		if err != nil {
			panic(err)
		}
	default:
	}

	SecretKey = os.Getenv("AUTHORIZE_RSA")
}

// GetDatabaseConf :　データベースの接続情報の取得
func GetDatabaseConf() (dsn string) { return os.Getenv("DB_URL") }
