package server

import (
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/random/", makeHandler(randomHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
