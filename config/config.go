package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	switch os.Getenv("GO_ENV") {
	case "dev":
		err := godotenv.Load("./config/env/.env.dev")
		if err != nil {
			panic(err)
		}
	default:
	}
}

// GetDatabaseConf :　データベースの接続情報の取得
func GetDatabaseConf() (dsn string) { return os.Getenv("DB_URL") }
