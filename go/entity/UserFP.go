package entity

func (UserFP) InitTable() string {
	return "PostFavorite"
}

type UserFP struct {
	UserID int `json:"userID"`
	PostID int `json:"postID"`
}
