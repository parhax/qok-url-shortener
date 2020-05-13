package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"qok.com/url_shortener/http/controller"
)

func Run(port string) {
	fmt.Print("QoK Url Shortener http router is running :) !")
	r := mux.NewRouter()

	r.HandleFunc("/shorten", controller.ShortenHandler).Methods("POST")
	// r.HandleFunc("/direct",controller.DirectHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
