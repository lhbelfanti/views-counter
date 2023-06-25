package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

const FILENAME = "views.txt"

type FileDatabase struct {
	mutex *sync.Mutex
}

// NewFileDatabase creates a new *FileDatabase
// It also initializes the mutex
func NewFileDatabase() *FileDatabase {
	return &FileDatabase{
		mutex: &sync.Mutex{},
	}
}

// GetCurrentCount returns the current view count
func (fileDB *FileDatabase) GetCurrentCount() int {
	fileDB.mutex.Lock()
	defer fileDB.mutex.Unlock()

	file, err := openFile()
	if err == nil {
		defer file.Close()
		count := getCount(file)

		return count
	}

	// Create the file if it doesn't exist
	count := 1
	createFile(count)

	return count
}

// UpdateCurrentCount updates the current view count and returns it
func (fileDB *FileDatabase) UpdateCurrentCount() int {
	fileDB.mutex.Lock()
	defer fileDB.mutex.Unlock()

	file, err := openFile()
	if err == nil {
		defer file.Close()

		count := getCount(file)
		count = updateCount(count, file)

		return count
	}

	// Create the file if it doesn't exist
	count := 1
	createFile(count)

	return count
}

func openFile() (*os.File, error) {
	_, err := os.Stat(FILENAME)
	if err == nil {
		// Open the file for reading and writing
		file, err := os.OpenFile(FILENAME, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Failed to open the file.")
			os.Exit(1)
		}

		return file, nil
	}

	return nil, err
}

func getCount(file *os.File) int {
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

	return count
}

func updateCount(count int, file *os.File) int {
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

	return count
}

func createFile(initialCount int) {
	err := ioutil.WriteFile(FILENAME, []byte(strconv.Itoa(initialCount)), 0644)
	if err != nil {
		fmt.Println("Failed to create the file.")
		os.Exit(1)
	}
}
