package counter

import (
	"fmt"
	"net/http"

	"views-counter/src/db"
)

// GetCurrentCounterHTTPHandler HTTP handler for '/count'
type GetCurrentCounterHTTPHandler func(w http.ResponseWriter, r *http.Request)

// MakeGetCurrentCountHTTPHandler Creates a new GetCurrentCounterHTTPHandler
func MakeGetCurrentCountHTTPHandler(database db.CounterPersistence) GetCurrentCounterHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get current count
		counter := database.GetCurrentCount()

		// Output the response (SVG image)
		fmt.Fprintf(w, "%d", counter)
	}
}
