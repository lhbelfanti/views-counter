package counter

import (
	"fmt"
	"net/http"
	"views-counter/src/badge"

	"views-counter/src/db"
)

type (
	// GetCurrentCounterHTTPHandler HTTP handler for '/api/views'
	GetCurrentCounterHTTPHandler func(w http.ResponseWriter, r *http.Request)
	// GetCurrentCounterAndReturnBadgeHTTPHandler HTTP handler for '/api/views/badge'
	GetCurrentCounterAndReturnBadgeHTTPHandler func(w http.ResponseWriter, r *http.Request)
)

// MakeGetCurrentCountHTTPHandler Creates a new GetCurrentCounterHTTPHandler
func MakeGetCurrentCountHTTPHandler(database db.CounterPersistence) GetCurrentCounterHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get current count
		counter := database.GetCurrentCount()

		fmt.Fprintf(w, "%d", counter)
	}
}

// MakeGetCurrentCounterAndReturnBadgeHTTPHandler Creates a new GetCurrentCounterAndReturnBadgeHTTPHandler
func MakeGetCurrentCounterAndReturnBadgeHTTPHandler(database db.CounterPersistence, createBadge badge.Create) GetCurrentCounterAndReturnBadgeHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		timestamp := "Mon, 01 Jan 2000 00:00:00 GMT"
		w.Header().Set("Expires", timestamp)
		w.Header().Set("Last-Modified", timestamp)
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Cache-Control", "no-cache, must-revalidate")

		w.Header().Set("Content-type", "image/svg+xml")

		// Get current count
		counter := database.GetCurrentCount()
		viewsBadge := createBadge(counter)

		fmt.Fprintf(w, viewsBadge)
	}
}
