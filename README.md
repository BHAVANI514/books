# Book Rating API

A simple Golang web API to search books, submit reviews, and receive real-time notifications for new reviews. The application uses the [Gutendex API](https://gutendex.com/) for book details and implements an in-memory review system with optional WebSocket notifications.

---

## Features

1. **Search Books**
   - **Endpoint**: `GET /books?search=<query>`
   - Search for books by title using the Gutendex API.
   - **Returns**: A list of books with basic details:
     - Title
     - Authors
     - Languages
     - Download count

2. **Submit Review**
   - **Endpoint**: `POST /review`
   - **Payload example**:
     ```json
     {
       "book_id": 1342,
       "rating": 5,
       "review": "Amazing book!"
     }
     ```
   - **Validates**:
     - Rating (between 0 and 5).
     - Review text (cannot be empty).
   - **Stores**: Reviews in memory and calculates the average rating for the book.
   - **Broadcasts**: The review to WebSocket clients in real-time.

3. **Get Reviews**
   - **Endpoint**: `GET /reviews?book_id=<id>`
   - **Returns**: All reviews and the average rating for the given book ID.

4. **Real-Time Notifications**
   - **WebSocket endpoint**: `ws://localhost:8080/ws`
   - Receives a live update whenever a new review is submitted, including:
     - Book ID
     - Rating
     - Review text
     - Updated average rating

---

## Setup Instructions

Follow these steps to get the Book Rating API up and running locally.

### 1. Clone the Repository
Clone the repository to your local machine:
```bash
git clone <repository_url>
cd book-api
