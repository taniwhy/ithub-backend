package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/ithub-backend/config"

	// Postgres ドライバ
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	conn *gorm.DB
	err  error
)

// NewDatabase :　データベースコネクションの確立
func NewDatabase() *gorm.DB {
	dsn := config.GetDatabaseConf()
	if err != nil {
		panic(err.Error())
	}

	conn, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	return conn
}
