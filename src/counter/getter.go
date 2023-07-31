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
		w.Header().Set("Content-type", "image/svg+xml")

		// Get current count
		counter := database.GetCurrentCount()
		viewsBadge := createBadge(counter)

		fmt.Fprintf(w, viewsBadge)
	}
}
