package models

type Video struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	View        int    `json:"view"`
	Like        int    `json:"like"`
	Dislike     int    `json:"dislike"`
	DateUpload  string `json:"dateUpload"`
	DatePublish string `json:"datePublish"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Category    int    `json:"category"`
	Label       bool `json:"label"`
	Privacy     bool   `json:"privacy"`
	Location    string `json:"location"`
	Premium bool `json:"premium"`
}
