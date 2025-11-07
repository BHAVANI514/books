package handler

import (
	model "book-api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// --- In-memory storage (thread-safe) ---
var (
	reviewsMu sync.RWMutex
	// map bookID -> slice of ratings for average calculation and optional storage
	ratingsMap = make(map[int][]int)

	// optional: store full review texts per book (if you want)
	reviewsTextMap = make(map[int][]string)
)

// --- WebSocket hub (simple) ---
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	clientsMu sync.RWMutex
	clients   = make(map[*websocket.Conn]bool)

	// channel to broadcast JSON-serializable messages
	broadcast = make(chan model.WSMessage, 16)
)

// --- REST: submit a review ---
func reviewPostHandler(c *gin.Context) {
	var req model.ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if req.BookID == 0 || req.Rating < 0 || req.Rating > 5 || req.Review == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload (book_id, rating 0-5, review required)"})
		return
	}

	// Save in memory (thread-safe)
	reviewsMu.Lock()
	ratingsMap[req.BookID] = append(ratingsMap[req.BookID], req.Rating)
	reviewsTextMap[req.BookID] = append(reviewsTextMap[req.BookID], req.Review)
	// compute average
	sum := 0
	for _, r := range ratingsMap[req.BookID] {
		sum += r
	}
	avg := float32(sum) / float32(len(ratingsMap[req.BookID]))
	reviewsMu.Unlock()

	// Respond to HTTP client
	c.JSON(http.StatusOK, gin.H{
		"message":        "review added",
		"average_rating": avg,
	})

	// Broadcast WS event (non-blocking send)
	msg := model.WSMessage{
		Event:         "new_review",
		BookID:        req.BookID,
		Rating:        req.Rating,
		Review:        req.Review,
		AverageRating: avg,
	}
	select {
	case broadcast <- msg:
	default:
		// if channel full, drop to avoid blocking; optionally log
		log.Println("broadcast channel full, dropping message")
	}
}

// --- REST: get reviews (simple) ---
func reviewsGetHandler(c *gin.Context) {
	bookIDStr := c.DefaultQuery("bookId", "")
	if bookIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookId query parameter required"})
		return
	}
	var bookID int
	if _, err := fmt.Sscanf(bookIDStr, "%d", &bookID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bookId"})
		return
	}

	reviewsMu.RLock()
	texts := reviewsTextMap[bookID]
	rates := ratingsMap[bookID]
	reviewsMu.RUnlock()

	// simple response: list of reviews and average
	if len(rates) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no reviews for book"})
		return
	}
	sum := 0
	for _, r := range rates {
		sum += r
	}
	avg := float32(sum) / float32(len(rates))

	c.JSON(http.StatusOK, gin.H{
		"book_id":        bookID,
		"average_rating": avg,
		"ratings":        rates,
		"reviews":        texts,
	})
}

// --- WS handler: upgrade & register client ---
func WsHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws upgrade:", err)
		return
	}
	clientsMu.Lock()
	clients[ws] = true
	clientsMu.Unlock()
	log.Println("ws client connected")

	// Read loop keeps connection alive and detects disconnects
	for {
		if _, _, err := ws.NextReader(); err != nil {
			// client disconnected or error
			clientsMu.Lock()
			delete(clients, ws)
			clientsMu.Unlock()
			_ = ws.Close()
			log.Println("ws client disconnected")
			return
		}
	}
}

// --- broadcaster goroutine ---
func StartBroadcaster() {
	for msg := range broadcast {
		// marshal to JSON once
		payload, err := json.Marshal(msg)
		if err != nil {
			log.Println("marshal:", err)
			continue
		}

		clientsMu.RLock()
		for client := range clients {
			// write asynchronously per client to avoid one slow client blocking others
			go func(cl *websocket.Conn, p []byte) {
				if err := cl.WriteMessage(websocket.TextMessage, p); err != nil {
					// on error, remove client
					log.Println("write ws:", err)
					clientsMu.Lock()
					delete(clients, cl)
					clientsMu.Unlock()
					_ = cl.Close()
				}
			}(client, payload)
		}
		clientsMu.RUnlock()
	}
}

// var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

// AddClient adds a new connected WebSocket client
func AddClient(conn *websocket.Conn) {
	mu.Lock()
	clients[conn] = true
	mu.Unlock()
}

// RemoveClient removes a disconnected client
func RemoveClient(conn *websocket.Conn) {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
}

// Broadcast sends message to all connected clients
func Broadcast(message interface{}) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling message:", err)
		return
	}

	mu.Lock()
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println("WebSocket write error:", err)
			client.Close()
			delete(clients, client)
		}
	}
	mu.Unlock()
}

var BroadcastChan = make(chan model.ReviewBroadcast)

// This is called from ReviewHandler to push message to broadcaster
func BroadcastReview(msg model.ReviewBroadcast) {
	BroadcastChan <- msg
}
