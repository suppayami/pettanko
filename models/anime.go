package models

// Anime model
type Anime struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
}
