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

	body, _ := ioutil.ReadAll(req.Body)
	var longUrl string
	err := json.Unmarshal(body, &longUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", longUrl)

	var shortener model.Shortener

	shortener.SetUrls(longUrl)
	shortener.StoreInDb()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "the Shorted URL  is ", shortener.ShortUrl())
	return

}
