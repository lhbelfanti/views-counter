package counter

import (
	"fmt"
	"net/http"

	"views-counter/src/db"
)

// UpdateCurrentCountHTTPHandler HTTP handler for '/'
type UpdateCurrentCountHTTPHandler func(w http.ResponseWriter, r *http.Request)

// MakeUpdateCurrentCountHTTPHandler return a new UpdateCurrentCountHTTPHandler
func MakeUpdateCurrentCountHTTPHandler(database db.CounterPersistence) UpdateCurrentCountHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Increment the file and get the current count
		counter := database.UpdateCurrentCount()

		// Output the response (SVG image)
		fmt.Fprintf(w, "%d", counter)
	}
}
