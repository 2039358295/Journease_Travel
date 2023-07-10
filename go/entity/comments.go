package entity

func (Comments) comments() string {
	return "comments"
}

type Comments struct {
	ID      int    `json:"id"`
	Userid  int    `json:"userid"`
	Avatar  []byte `json:"avatar"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Likes   int    `json:"likes"`
	Time    string `json:"time"`
	Rating  int    `json:"rating"`
}
