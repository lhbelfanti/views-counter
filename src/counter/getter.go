package counter

import (
	"fmt"
	"net/http"

	"github.com/lhbelfanti/views-counter/src/badge"
	"github.com/lhbelfanti/views-counter/src/db"
)

// GetCurrentViewCountHTTPHandler HTTP handler for '/count'
func GetCurrentViewCountHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Disable cache
	timestamp := "Mon, 01 Jan 2000 00:00:00 GMT"
	w.Header().Set("Expires", timestamp)
	w.Header().Set("Last-Modified", timestamp)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")

	// Set the content type to be an image
	w.Header().Set("Content-type", "image/svg+xml")

	// Get current count
	message := db.GetCurrentCountFromFile("views.txt", true)

	response := badge.Create(message)

	// Output the response (SVG image)
	fmt.Fprintf(w, response)
}