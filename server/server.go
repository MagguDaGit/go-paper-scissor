package server

import (
	"log"
	"net/http"
)

func Serve() {
	log.Default().Println("Started webserver at: port 8080")
	http.HandleFunc("/random/", makeHandler(randomHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
