package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"qok.com/url_shortener/controller"
)

func main() {
	fmt.Print("QoK Url Shortener app is running !!")
	r := mux.NewRouter()

	r.HandleFunc("/shorten", controller.ShortenHandler).Methods("POST")
	// r.HandleFunc("/direct",controller.DirectHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8787", r))

}
