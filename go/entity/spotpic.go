package entity

func (Spotpic) spotpic() string {
	return ""
}

type Spotpic struct {
	ID      int    `json:"id"`
	Photo   string `json:"src1"`
	Comment string `json:"comment"`
}
