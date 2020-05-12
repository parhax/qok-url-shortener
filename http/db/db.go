package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"qok.com/url_shortener/http/config"
)

func GetMysqlConnection() (*sql.DB, error) {
	conf := config.Load()

	dataSourceName := fmt.Sprintf("root:root@tcp(%s)/%s", conf.Mysql_url, conf.Mysql_db)
	db, err := sql.Open("mysql", dataSourceName)
	return db, err
}
