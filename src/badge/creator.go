package badge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"views-counter/src/formatter"
)

func Create(message int) string {
	// Set parameters for the shields.io URL
	label := "Views"
	labelColor := "640464"
	logo := "eye"
	logoColor := "white"
	counter := formatter.ShortNumber(float64(message), 2)
	counterColor := "7c007c"
	style := "for-the-badge"

	// Build the URL with an SVG image of the view counter
	url := fmt.Sprintf("https://custom-icon-badges.demolab.com/badge/%s-%s-%s.svg?labelColor=%s&logo=%s&logoColor=%s&style=%s",
		label, counter, counterColor, labelColor, logo, logoColor, style)

	// Get the contents of the URL
	return curlGetContents(url)
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
