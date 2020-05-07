package controller

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Shortener struct {
	longUrl  string
	shortUrl string
}

func ShortenHandler(w http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(req)
	// longUrl := vars["longUrl"]

	body, _ := ioutil.ReadAll(req.Body)
	var longUrl string
	err := json.Unmarshal(body, &longUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", longUrl)

	shortener := Shortener{
		longUrl:  longUrl,
		shortUrl: shortenIt(longUrl),
	}

	shortener.StoreInDb()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "the shorted url is ", shortener.shortUrl)
	return

}

func shortenIt(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	shortedStr := base64.URLEncoding.EncodeToString(h.Sum(nil))
	shortedUrl := "q.ok/" + shortedStr[:8]
	fmt.Println(str, shortedStr)
	return shortedUrl
}

func (shortener *Shortener) StoreInDb() {
	mysqlUrl := os.Getenv("MYSQL_URL")
	mysqlDB := os.Getenv("MYSQL_DB")
	dataSourceName := fmt.Sprintf("root:root@tcp(%s)/%s", mysqlUrl, mysqlDB)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var tableQuery = `CREATE TABLE IF NOT EXISTS urls 
	(id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	 long_url VARCHAR(100) NOT NULL,
	 short_url VARCHAR(50),
	 reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP );`
	create, err := db.Query(tableQuery)
	if err != nil {
		panic(err.Error())
	}
	create.Close()

	var query = "INSERT IGNORE INTO urls (`long_url`,`short_url`) VALUES (?, ?)"

	insert, err := db.Query(query, shortener.longUrl, shortener.shortUrl)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
