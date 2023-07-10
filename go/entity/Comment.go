package entity

func (Comment) TableName() string {
	return "community_comment"
}

type Comment struct {
	UserID    int    `json:"userID" `
	CommentID int    `json:"commentID"`
	PostID    int    `json:"postID"`
	Content   string `json:"content" `
	Time      string `json:"time"`
	UpNumber  int    `json:"upNumber"`
}
