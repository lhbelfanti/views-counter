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
	//fileDatabase := db.NewFileDatabase()

	mongoDatabase := db.NewMongoDatabase()
	defer mongoDatabase.Close()

	createBadge := badge.MakeCreate()

	/* Create handlers functions */
	getCurrentCountHTTPHandler := counter.MakeGetCurrentCountHTTPHandler(mongoDatabase, createBadge)
	updateCurrentCountHTTPHandler := counter.MakeUpdateCurrentCountHTTPHandler(mongoDatabase, createBadge)

	/* Handlers */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/views", getCurrentCountHTTPHandler)
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
