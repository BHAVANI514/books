package handler

import (
	model "book-api/models"
	"book-api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReviewHandler(c *gin.Context) {
	var review model.Review

	// Bind JSON
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Add review and calculate average rating
	avg, err := service.AddReview(review.BookID, review.Rating, review.Review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Broadcast WebSocket event
	Broadcast(map[string]interface{}{
		"event":          "new_review",
		"book_id":        review.BookID,
		"rating":         review.Rating,
		"review":         review.Review,
		"average_rating": avg,
	})

	// ✅ Return success
	c.JSON(http.StatusOK, gin.H{
		"message":        "Review successfully submitted",
		"average_rating": avg,
	})
}

// GetReviewsHandler retrieves all reviews for a specific book
func GetReviewsHandler(c *gin.Context) {
	bookIDStr := c.Query("book_id")

	if bookIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book_id is required"})
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book_id must be a number"})
		return
	}
	fmt.Println("bookID----------->", bookID)
	reviews, err := service.GetReviews(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
