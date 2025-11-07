# Book Rating API

A simple Golang web API to search books, submit reviews, and receive real-time notifications for new reviews. The application uses the [Gutendex API](https://gutendex.com/) for book details and implements an in-memory review system with optional WebSocket notifications.

---

## Features

1. **Search Books**
   - Endpoint: `GET /books?search=<query>`
   - Search for books by title using the Gutendex API.
   - Returns a list of books with basic details: title, authors, languages, download count.

2. **Submit Review**
   - Endpoint: `POST /review`
   - Payload example:
     ```json
     {
       "book_id": 1342,
       "rating": 5,
       "review": "Amazing book!"
     }
     ```
   - Validates rating (0-5) and review text.
   - Stores reviews in memory and calculates the average rating.
   - Broadcasts the review to WebSocket clients in real-time.

3. **Get Reviews**
   - Endpoint: `GET /reviews?book_id=<id>`
   - Returns all reviews and average rating for the given book ID.

4. **Real-Time Notifications**
   - WebSocket endpoint: `ws://localhost:8080/ws`
   - Receives a live update whenever a new review is submitted, including:
     - Book ID
     - Rating
     - Review text
     - Updated average rating

---

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone <repository_url>
   cd book-api
