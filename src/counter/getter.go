package counter

import (
	"fmt"
	"net/http"

	"views-counter/src/db"
)

// GetCurrentCounterHTTPHandler HTTP handler for '/count'
type GetCurrentCounterHTTPHandler func(w http.ResponseWriter, r *http.Request)

// MakeGetCurrentCountHTTPHandler Creates a new GetCurrentCounterHTTPHandler
func MakeGetCurrentCountHTTPHandler() GetCurrentCounterHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		mongoDatabase := db.NewMongoDatabase()
		defer mongoDatabase.Close()

		// Get current count
		counter := mongoDatabase.GetCurrentCount()

		// Output the response (SVG image)
		fmt.Fprintf(w, "%d", counter)
	}
}
