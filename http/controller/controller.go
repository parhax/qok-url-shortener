package controller

import (
	"crypto/sha1"
	"encoding/base64"
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
	var longURL string
	unmarshErr := json.Unmarshal(body, &longURL)

	if unmarshErr != nil {
		logger.Fatal(unmarshErr)
	}

	shortener := model.Shortener{
		LongUrl:  longURL,
		ShortUrl: shortenLong(longURL),
	}

	shortener.StoreInDb()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "the Shorted URL  is ", shortener.ShortUrl)
	return

}

func shortenLong(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	shortedStr := base64.URLEncoding.EncodeToString(h.Sum(nil))
	shortedUrl := "q.ok/" + shortedStr[:8]
	fmt.Println(str, shortedStr)
	return shortedUrl
}
