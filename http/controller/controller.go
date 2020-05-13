package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"qok.com/url_shortener/http/logwrapper"
	"qok.com/url_shortener/http/model"
)

func ShortenHandler(w http.ResponseWriter, req *http.Request) {
	logger := logwrapper.Load()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Fatal(err)
	}
	var longUrl string
	unmarshErr := json.Unmarshal(body, &longUrl)

	if unmarshErr != nil {
		logger.Fatal(unmarshErr)
	}

	var shortener model.Shortener

	shortener.SetUrls(longUrl)
	shortener.StoreInDb()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "the Shorted URL  is ", shortener.ShortUrl())
	return

}
