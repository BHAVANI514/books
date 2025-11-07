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

2. Install Dependencies

Make sure you have Go installed. Then, fetch the required Go modules:

go mod tidy

3. Run the API Server

Start the API server by running:

go run main.go


This will start the server at http://localhost:8080.

4. WebSocket Notifications

The WebSocket server runs at ws://localhost:8080/ws. You can use a WebSocket client like Postman or wscat to receive real-time notifications whenever a new review is posted.

Testing the API with Postman
1. Search Books

To search for books, send a GET request with the search query parameter.

Request:
Method: GET
URL: http://localhost:8080/books?search=Sherlock

Response:

[
  {
    "id": 1342,
    "title": "The Adventures of Sherlock Holmes",
    "authors": ["Arthur Conan Doyle"],
    "languages": ["en"],
    "downloads": 321
  },
  {
    "id": 1343,
    "title": "A Study in Scarlet",
    "authors": ["Arthur Conan Doyle"],
    "languages": ["en"],
    "downloads": 250
  }
]

2. Submit a Review

To submit a review, use a POST request with the review data in the request body.

Request:

Method: POST

URL: http://localhost:8080/review

Headers:

Content-Type: application/json

Body (raw JSON):

{
  "book_id": 1342,
  "rating": 5,
  "review": "Amazing book!"
}


Response:

{
  "message": "Review submitted successfully",
  "average_rating": 5.0
}

3. Get Reviews for a Book

To get all reviews for a specific book, use a GET request with the book_id query parameter.

Request:

Method: GET

URL: http://localhost:8080/reviews?book_id=1342

Response:

{
  "book_id": 1342,
  "average_rating": 5.0,
  "reviews": [
    {
      "rating": 5,
      "review": "Amazing book!"
    }
  ]
}

4. Real-Time Notifications via WebSocket

To receive real-time notifications of new reviews, use a WebSocket client like Postman.

Steps for Testing WebSocket in Postman:

Open Postman and create a new request.

Set the Request Type to WebSocket.

In the URL field, enter the WebSocket URL:
ws://localhost:8080/ws

Click Connect to establish the WebSocket connection.

Once connected, you will receive live updates like the following when new reviews are posted:

WebSocket Message:

{
  "book_id": 1342,
  "rating": 5,
  "review": "Amazing book!",
  "average_rating": 5.0
}

