package main

import (
	"book-api/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Define the route for searching books
	r.GET("/books", handler.SearchBooksHandler) //// Get reviews for a book
	r.POST("/review", handler.ReviewHandler)    // Submit a review
	r.GET("/reviews", handler.GetReviewsHandler)
	//r.POST("/review", handlers.AddReviewHandler)

	// WebSocket endpoint
	r.GET("/ws", handler.WsHandler)

	// Start broadcaster
	go handler.StartBroadcaster()
	// Start the server
	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

/*
1.http://localhost:8080/books?search=harry

2. http://localhost:8080/review
{
  "book_id": 1342,
  "rating": 5,
  "review": "amazing book!"
}

3.http://localhost:8080/reviews?book_id=84

WebSocket endpoint
4.ws://localhost:8080/ws
*/
