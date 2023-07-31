package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/carlmjohnson/gateway"

	"views-counter/src/counter"
	"views-counter/src/db"
)

func Init() {
	/* Dependencies */
	//fileDatabase := db.NewFileDatabase()

	mongoDatabase := db.NewMongoDatabase()
	defer mongoDatabase.Close()

	/* Create handlers functions */
	getCurrentCountHTTPHandler := counter.MakeGetCurrentCountHTTPHandler()
	updateCurrentCountHTTPHandler := counter.MakeUpdateCurrentCountHTTPHandler(mongoDatabase)

	/* Handlers */
	http.Handle("/views", http.FileServer(http.Dir("./public/views")))
	http.Handle("/increment", http.FileServer(http.Dir("./public/increment")))

	http.Handle("/api/views", http.Handler(http.HandlerFunc(getCurrentCountHTTPHandler)))
	http.Handle("/api/increment", http.Handler(http.HandlerFunc(updateCurrentCountHTTPHandler)))

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
