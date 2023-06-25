package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

// GetCurrentCountFromFile returns the current view count
// If the parameter incrementByOne is true, it will increment the current count
func GetCurrentCountFromFile(filename string, incrementByOne bool) int {
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

		if incrementByOne {
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
