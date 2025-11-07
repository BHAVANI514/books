package handler

import (
	"book-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchBooksHandler handles book search requests based on a query string
func SearchBooksHandler(c *gin.Context) {
	query := c.Query("search")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	books, err := service.FetchBooks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
