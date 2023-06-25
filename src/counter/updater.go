package counter

import (
	"fmt"
	"net/http"

	"views-counter/src/badge"
	"views-counter/src/db"
)

// IncrementCount HTTP handler for '/'
func IncrementCountHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Disable cache
	timestamp := "Mon, 01 Jan 2000 00:00:00 GMT"
	w.Header().Set("Expires", timestamp)
	w.Header().Set("Last-Modified", timestamp)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")

	// Set the content type to be an image
	w.Header().Set("Content-type", "image/svg+xml")

	// Increment the file and get the current count
	message := db.GetCurrentCountFromFile("views.txt", true)

	response := badge.Create(message)

	// Output the response (SVG image)
	fmt.Fprintf(w, response)
}
