package service

import (
	model "book-api/models"
	"fmt"
)

// In-memory storage for reviews (bookId -> list of reviews)
var reviews = make(map[int][]model.Review)

// AddReview adds a review for a specific book and returns new average rating
func AddReview(bookID int, rating int, reviewText string) (float64, error) {
	// Validate inputs
	if rating < 0 || rating > 5 {
		return 0, fmt.Errorf("invalid rating. Must be between 0 and 5")
	}
	if reviewText == "" {
		return 0, fmt.Errorf("review cannot be empty")
	}
	fmt.Println("review")
	// Create review object
	review := model.Review{
		BookID: bookID,
		Rating: rating,
		Review: reviewText,
	}
	fmt.Println("review----------->", review)
	// Store review in memory
	reviews[bookID] = append(reviews[bookID], review)

	// ✅ Calculate average
	newAvg := calculateAverage(reviews[bookID])

	return newAvg, nil
}

// GetReviews retrieves all reviews for a specific book
func GetReviews(bookID int) ([]model.Review, error) {
	// Retrieve the reviews for the specified book ID
	if reviewList, found := reviews[bookID]; found {
		fmt.Println("reviewList", found, reviewList)
		return reviewList, nil
	}
	fmt.Println("reviewList")
	return nil, fmt.Errorf("no reviews found for book ID %d", bookID)
}
func calculateAverage(list []model.Review) float64 {
	sum := 0
	for _, r := range list {
		sum += r.Rating
	}
	return float64(sum) / float64(len(list))
}
