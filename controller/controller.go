package controller

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/qok_url_shortener")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var query = "INSERT IGNORE INTO urls (`long_url`,`short_url`) VALUES (?, ?)"

	insert, err := db.Query(query, shortener.longUrl, shortener.shortUrl)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
