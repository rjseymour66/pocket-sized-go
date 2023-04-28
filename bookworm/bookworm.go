package main

import (
	"encoding/json"
	"os"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Initialize the type the file decodes into.
	var bookworms []Bookworm

	// Decode the file and store the contents in the value bookworms.
	// err = json.NewDecoder(f).Decode(&bookworms)
	// if err != nil {
	// 	return nil, err
	// }
	if err := json.NewDecoder(f).Decode(&bookworms); err != nil {
		return nil, err
	}

	return bookworms, nil
}
