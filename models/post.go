package models

type Post struct {
	ID       string            `json:"id"`
	UserID    string            `json:"userId"`
	Post     string            `json:"post"`
	PostDate string            `json:"postDate"`
	ChannelID string `json:"channelId"`
	Thumbnail string `json:"thumbnail"`
}

type PostLikeDetail struct {
	PostID string `json:"postId"`
	UserID string `json:"userId"`
	Like bool  `json:"like"`
}
