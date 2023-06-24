package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	. "github.com/tbxark/g4vercel"
)

// Disable cache so that the image will be fetched every time
func disableCache(w http.ResponseWriter) {
	timestamp := "Mon, 01 Jan 2000 00:00:00 GMT"
	w.Header().Set("Expires", timestamp)
	w.Header().Set("Last-Modified", timestamp)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
}

var mutex = &sync.Mutex{}

// Increment the file and return the current number
func incrementFile(filename string) int {
	// If the file exists
	if _, err := os.Stat(filename); err == nil {
		mutex.Lock()
		defer mutex.Unlock()

		// Open the file for reading and writing
		file, err := os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Failed to open the file.")
			os.Exit(1)
		}
		defer file.Close()

		// Read the file and add 1
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Failed to read the file.")
			os.Exit(1)
		}
		count, err := strconv.Atoi(string(content))
		if err != nil {
			fmt.Println("Failed to convert file content to integer.")
			os.Exit(1)
		}
		count++

		// Delete the contents
		if err := file.Truncate(0); err != nil {
			fmt.Println("Failed to truncate the file.")
			os.Exit(1)
		}

		// Go to the beginning of the file
		if _, err := file.Seek(0, 0); err != nil {
			fmt.Println("Failed to seek to the beginning of the file.")
			os.Exit(1)
		}

		// Write the new count
		if _, err := file.WriteString(strconv.Itoa(count)); err != nil {
			fmt.Println("Failed to write to the file.")
			os.Exit(1)
		}

		// Return the current file contents
		return count
	}

	// Create the file if it doesn't exist
	count := 1
	err := ioutil.WriteFile(filename, []byte(strconv.Itoa(count)), 0644)
	if err != nil {
		fmt.Println("Failed to create the file.")
		os.Exit(1)
	}

	// Return the current file contents
	return count
}

type LookupItem struct {
	Value  float64
	Symbol string
}

func shortNumber(num float64, digits int) string {
	lookup := []LookupItem{
		{Value: 1, Symbol: ""},
		{Value: 1e3, Symbol: "k"},
		{Value: 1e6, Symbol: "M"},
		{Value: 1e9, Symbol: "G"},
		{Value: 1e12, Symbol: "T"},
		{Value: 1e15, Symbol: "P"},
		{Value: 1e18, Symbol: "E"},
	}

	for i := len(lookup) - 1; i >= 0; i-- {
		item := lookup[i]
		if num >= item.Value {
			result := num / item.Value
			formatted := fmt.Sprintf("%.*f", digits, result)

			// Remove trailing zeros
			for formatted[len(formatted)-1] == '0' {
				formatted = formatted[:len(formatted)-1]
			}

			// Remove the decimal point if there are no decimal digits
			if formatted[len(formatted)-1] == '.' {
				formatted = formatted[:len(formatted)-1]
			}

			return formatted + item.Symbol
		}
	}

	return "0"
}

// Get contents of a URL with HTTP GET
func curlGetContents(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to make the HTTP GET request.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response body.")
		os.Exit(1)
	}

	return string(body)
}

// Handler http handler for Vercel hosting
func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))

	server.GET("/", func(context *Context) {
		// Disable cache
		disableCache(w)

		// Set the content type to be an image
		w.Header().Set("Content-type", "image/svg+xml")

		// Increment the file and get the current count
		message := incrementFile("views.txt")

		// Set parameters for the shields.io URL
		label := "Views"
		labelColor := "640464"
		logo := "eye"
		logoColor := "white"
		counter := shortNumber(float64(message), 2)
		counterColor := "7c007c"
		style := "for-the-badge"

		// Build the URL with an SVG image of the view counter
		url := fmt.Sprintf("https://custom-icon-badges.demolab.com/badge/%s-%s-%s.svg?labelColor=%s&logo=%s&logoColor=%s&style=%s",
			label, counter, counterColor, labelColor, logo, logoColor, style)

		// Get the contents of the URL
		response := curlGetContents(url)

		// Output the response (SVG image)
		fmt.Fprintf(w, response)
	})

	server.GET("/favicon.ico", func(context *Context) {})
}
