package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lhbelfanti/views-counter/src/counter"
)

func Init() {
	/* Handlers */
	http.HandleFunc("/", counter.IncrementCountHTTPHandler)

	http.HandleFunc("/count", counter.GetCurrentViewCountHTTPHandler)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Failed to start the HTTP server.")
		os.Exit(1)
	}
}
