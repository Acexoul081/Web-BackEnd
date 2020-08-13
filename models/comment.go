package models

type Comment struct {
	ID          string               `json:"id"`
	UserID      string               `json:"userId"`
	VideoID     string               `json:"videoId"`
	Comment     string               `json:"comment"`
	CommentDate string               `json:"commentDate"`
}
