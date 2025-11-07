package model

// Author represents an author of the book
type Author struct {
	Name      string `json:"name"`
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
}

// Book represents a book and its associated details
type Book struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Authors       []Author `json:"authors"`
	Languages     []string `json:"languages"`
	DownloadCount int      `json:"download_count"`
}

// Review represents a review for a book
type Review struct {
	BookID int    `json:"book_Id"`
	Rating int    `json:"rating"`
	Review string `json:"review"`
}

// --- Models ---
type ReviewRequest struct {
	BookID int    `json:"book_id"` // input JSON uses snake_case
	Rating int    `json:"rating"`
	Review string `json:"review"`
}

type WSMessage struct {
	Event         string  `json:"event"`
	BookID        int     `json:"book_id"`
	Rating        int     `json:"rating"`
	Review        string  `json:"review"`
	AverageRating float32 `json:"average_rating"`
}

type ReviewBroadcast struct {
	BookID        int     `json:"book_id"`
	Rating        int     `json:"rating"`
	AverageRating float64 `json:"average_rating"`
}
