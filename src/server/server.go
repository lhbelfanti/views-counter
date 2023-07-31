package server

import (
	"fmt"
	"net/http"
	"os"
	"views-counter/src/badge"

	"github.com/carlmjohnson/gateway"

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
	getCurrentCountHTTPHandler := counter.MakeGetCurrentCountHTTPHandler(mongoDatabase)
	getCurrentCounterAndReturnBadgeHTTPHandler := counter.MakeGetCurrentCounterAndReturnBadgeHTTPHandler(mongoDatabase, createBadge)
	updateCurrentCountHTTPHandler := counter.MakeUpdateCurrentCountHTTPHandler(mongoDatabase)
	updateCurrentCountAndReturnBadgeHTTPHandler := counter.MakeUpdateCurrentCountAndReturnBadgeHTTPHandler(mongoDatabase, createBadge)

	/* Handlers */
	http.Handle("/views", http.FileServer(http.Dir("./public/views")))
	http.Handle("/increment", http.FileServer(http.Dir("./public/increment")))

	http.Handle("/api/views", http.Handler(http.HandlerFunc(getCurrentCountHTTPHandler)))
	http.Handle("/api/increment", http.Handler(http.HandlerFunc(updateCurrentCountHTTPHandler)))

	http.Handle("/api/views/badge", http.Handler(http.HandlerFunc(getCurrentCounterAndReturnBadgeHTTPHandler)))
	http.Handle("/api/increment/badge", http.Handler(http.HandlerFunc(updateCurrentCountAndReturnBadgeHTTPHandler)))

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	listener := gateway.ListenAndServe
	// Start the HTTP server
	err := listener(":"+port, nil)
	if err != nil {
		fmt.Println("Failed to start the HTTP server.")
		os.Exit(1)
	}
}
