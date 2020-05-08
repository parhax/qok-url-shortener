package model

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"

	"qok.com/url_shortener/http/db"
)

type Shortener struct {
	longUrl  string
	shortUrl string
}

func (shortener *Shortener) SetUrls(long_url string) {
	shortener.longUrl = long_url
	shortener.shortUrl = shortenLong(long_url)
}

func (sh *Shortener) ShortUrl() string {
	return sh.shortUrl
}

func (shortener *Shortener) StoreInDb() {
	// mysqlUrl := os.Getenv("MYSQL_URL")
	// mysqlDB := os.Getenv("MYSQL_DB")
	// dataSourceName := fmt.Sprintf("root:root@tcp(%s)/%s", mysqlUrl, mysqlDB)
	// db, err := sql.Open("mysql", dataSourceName)
	msql, err := db.GetMysqlConnection()
	if err != nil {
		panic(err.Error())
	}
	defer msql.Close()

	checkForTableExistance(msql)

	var query = "INSERT IGNORE INTO urls (`long_url`,`short_url`) VALUES (?, ?)"

	insert, err := msql.Query(query, shortener.longUrl, shortener.shortUrl)

	if err != nil {
		panic(err.Error())
	}

	insert.Close()

}

func shortenLong(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	shortedStr := base64.URLEncoding.EncodeToString(h.Sum(nil))
	shortedUrl := "q.ok/" + shortedStr[:8]
	fmt.Println(str, shortedStr)
	return shortedUrl
}

func checkForTableExistance(msql *sql.DB) {
	var tableQuery = `CREATE TABLE IF NOT EXISTS urls 
	(id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	 long_url VARCHAR(100) NOT NULL,
	 short_url VARCHAR(50),
	 reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP );`
	create, err := msql.Query(tableQuery)
	if err != nil {
		panic(err.Error())
	}
	create.Close()
}
