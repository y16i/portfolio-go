package database

import (
	"database/sql"
	"fmt"
	"time"

	"bitbucket.org/y16i/backend-go/tool"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase(wpConfigPath string) *sql.DB {
	wpConfig := tool.ReadWpConfig(wpConfigPath)
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		wpConfig.DbUser, wpConfig.DbPassword, wpConfig.DbHost, wpConfig.DbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return db
}
