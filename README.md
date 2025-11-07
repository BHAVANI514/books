# Book Rating API

A simple Golang web API that allows users to search for books, submit reviews, and receive real-time notifications for new reviews. This application utilizes the [Gutendex API](https://gutendex.com/) for fetching book details and features an in-memory review system with optional WebSocket notifications.

---

## Features

1. **Search Books**
   - **Endpoint**: `GET /books?search=<query>`
   - Search for books by title using the Gutendex API.
   - **Returns**: A list of books with the following details:
     - Title
     - Authors
     - Languages
     - Download count

2. **Submit Review**
   - **Endpoint**: `POST /review`
   - **Payload**:
     ```json
     {
       "book_id": 1342,
       "rating": 5,
       "review": "Amazing book!"
     }
     ```
   - **Validation**: Ensures that the rating is between 0 and 5 and that the review text is provided.
   - **Stores**: Reviews in memory and calculates the average rating for the book.
   - **Real-time Notifications**: Broadcasts the review to all connected WebSocket clients.

3. **Get Reviews**
   - **Endpoint**: `GET /reviews?book_id=<id>`
   - **Returns**: All reviews for a specific book ID and the updated average rating.

4. **Real-Time Notifications**
   - **WebSocket endpoint**: `ws://localhost:8080/ws`
   - Receive live updates when a new review is submitted, including:
     - Book ID
     - Rating
     - Review text
     - Updated average rating

---

## Setup Instructions

Follow these steps to set up and run the Book Rating API locally.

### 1. Clone the Repository
First, clone the repository to your local machine:
```bash
git clone <repository_url>
cd book-api

