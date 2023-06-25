package server

import (
	"fmt"
	"net/http"
	"os"

	"views-counter/src/badge"
	"views-counter/src/counter"
	"views-counter/src/db"
)

func Init() {

	/* Dependencies */
	fileDatabase := db.NewFileDatabase()
	createBadge := badge.MakeCreate()

	/* Create handlers functions */
	getCurrentCountHTTPHandler := counter.MakeGetCurrentCountHTTPHandler(fileDatabase, createBadge)
	updateCurrentCountHTTPHandler := counter.MakeUpdateCurrentCountHTTPHandler(fileDatabase, createBadge)

	/* Handlers */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/count", getCurrentCountHTTPHandler)
	http.HandleFunc("/increment", updateCurrentCountHTTPHandler)
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
