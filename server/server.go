package server

import (
	"log"
	"net/http"
	"time"
)

func Serve() {
	// ServeMux for handling routes
	mux := http.NewServeMux()

	// Handlers for routes
	mux.HandleFunc("/simulation/", makeHandler(randomSimulateHandler))
	mux.HandleFunc("/play/", makeHandler(playRandomHandler))
	// Create a custom server with the mux
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux, // Use the custom ServeMux
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // Keep-alive
	}

	log.Default().Println("Started webserver at: port 8080")
	log.Fatal(server.ListenAndServe())
}
