package database

import (
	"database/sql"
	"fmt"
	"time"

	"bitbucket.org/y16i/backend-go/tool"
	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func InitDatabase(wpConfigPath string) {
	wpConfig := tool.ReadWpConfig(wpConfigPath)
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		wpConfig.DbUser, wpConfig.DbPassword, wpConfig.DbHost, wpConfig.DbName)
	var err error
	conn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(5)
}
