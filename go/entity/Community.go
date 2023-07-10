package entity

func (Post) TableName() string {
	return "community_main"
}

type Post struct {
	UserID     int    `json:"UserID"`
	PostID     int    `json:"PostID"`
	Title      string `json:"Title"`
	Tag1       string `json:"Tag1"`
	Tag2       string `json:"Tag2"`
	Contents   string `json:"Contents"`
	Photo      []byte `json:"photo"`
	Photo1     []byte `json:"photo1"`
	Photo2     []byte `json:"photo2"`
	Photo3     []byte `json:"photo3"`
	Photo4     []byte `json:"photo4"`
	Photo5     []byte `json:"photo5"`
	Photo6     []byte `json:"photo6"`
	Photo7     []byte `json:"photo7"`
	Photo8     []byte `json:"photo8"`
	Likenum    int    `json:"Likenum"`
	Collection int    `json:"Collection"`
	Commentnum int    `json:"Commentnum"`
	RTime      string `json:"RTime"`
}
