package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"qok.com/url_shortener/http/model"
)

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

	// shortener := Shortener{
	// 	longUrl:  longUrl,
	// 	shortUrl: shortenIt(longUrl),
	// }

	var shortener model.Shortener

	// shortener = shortener.SetLongurl(longUrl)
	shortener.SetUrls(longUrl)
	shortener.StoreInDb()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "the Shorted URL  is ", shortener.ShortUrl())
	return

}
