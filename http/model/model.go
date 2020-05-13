package model

import (
	"database/sql"

	"qok.com/url_shortener/http/db"
)

type Shortener struct {
	LongUrl  string
	ShortUrl string
}

//StoreInDB store shorted url in database
func (shortener *Shortener) StoreInDb() {
	msql, err := db.GetMysqlConnection()
	if err != nil {
		panic(err.Error())
	}
	defer msql.Close()

	checkForTableExistance(msql)

	var query = "INSERT IGNORE INTO urls (`long_url`,`short_url`) VALUES (?, ?)"

	insert, err := msql.Query(query, shortener.LongUrl, shortener.ShortUrl)

	if err != nil {
		panic(err.Error())
	}

	insert.Close()

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
