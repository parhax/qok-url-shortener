package main

import (
	"qok.com/url_shortener/http/config"
	"qok.com/url_shortener/http/router"
)

func main() {
	router.Run(config.Load().Http_port)
}
