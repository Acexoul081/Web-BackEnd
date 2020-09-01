package models

type Playlist struct {
	ID             string            `json:"id"`
	OwnerID        string             `json:"ownerId"`
	Title          string            `json:"title"`
	Privacy        bool              `json:"privacy"`
	Date           string            `json:"date"`
	Sort			string				`json:"sort"`
	Description		string 	`json:"description"`
}

type PlaylistDetail struct {
	PlaylistID string `json:"playlistId"`
	VideoID      string `json:"videoId"`
	TempID string `json: tempId`
}

type PlaylistSub struct {
	PlaylistID string `json:"playlistId"`
	UserID       string  `json:"userId"`
}

