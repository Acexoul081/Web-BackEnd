package models

type ReplyLikeDetail struct {
	ReplyID string `json:"replyID"`
	UserID  string  `json:"userID"`
	Like  bool   `json:"like"`
}

type LikeDetail struct {
	VideoID string `json:"videoID"`
	UserID  string  `json:"userID"`
	Like  bool  `json:"like"`
}

type CommentLikeDetail struct {
	CommentID string `json:"commentID"`
	UserID    string    `json:"userID"`
	Like    bool     `json:"like"`
}
