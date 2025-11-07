package service

import (
	model "book-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FetchBooks makes a GET request to the Gutendex API to search for books by title
func FetchBooks(query string) ([]model.Book, error) {
	// Build the URL to search for books in Gutendex
	url := fmt.Sprintf("https://gutendex.com/books/?search=%s", query)

	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books from Gutendex: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response into a struct
	var response struct {
		Results []model.Book `json:"results"` // Access the 'results' field instead of 'books'
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Return the list of books from 'results'
	return response.Results, nil
}
