package counter

import (
	"fmt"
	"net/http"

	"views-counter/src/badge"
	"views-counter/src/db"
)

type (
	// UpdateCurrentCountHTTPHandler http handler for '/api/increment'
	UpdateCurrentCountHTTPHandler func(w http.ResponseWriter, r *http.Request)
	// UpdateCurrentCountAndReturnBadgeHTTPHandler http handler for '/api/increment/badge'
	UpdateCurrentCountAndReturnBadgeHTTPHandler func(w http.ResponseWriter, r *http.Request)
)

// MakeUpdateCurrentCountHTTPHandler return a new UpdateCurrentCountHTTPHandler
func MakeUpdateCurrentCountHTTPHandler(database db.CounterPersistence) UpdateCurrentCountHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Increment the file and get the current count
		counter := database.UpdateCurrentCount()

		fmt.Fprintf(w, "%d", counter)
	}
}

// MakeUpdateCurrentCountAndReturnBadgeHTTPHandler return a new UpdateCurrentCountAndReturnBadgeHTTPHandler
func MakeUpdateCurrentCountAndReturnBadgeHTTPHandler(database db.CounterPersistence, createBadge badge.Create) UpdateCurrentCountAndReturnBadgeHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		timestamp := "Mon, 01 Jan 2000 00:00:00 GMT"
		w.Header().Set("Expires", timestamp)
		w.Header().Set("Last-Modified", timestamp)
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Cache-Control", "no-cache, must-revalidate")

		w.Header().Set("Content-type", "image/svg+xml")

		// Increment the file and get the current count
		counter := database.UpdateCurrentCount()
		viewsBadge := createBadge(counter)

		fmt.Fprintf(w, viewsBadge)
	}
}
