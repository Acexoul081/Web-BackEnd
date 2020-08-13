package models

type Reply struct {
	ID        string             `json:"id"`
	CommentID string             `json:"commentId"`
	UserID      string              `json:"userID"`
	Reply     string             `json:"reply"`
	ReplyDate string             `json:"replyDate"`
}
