package entity

func (Userorders) userorders() string {
	return "userorders"
}

type Userorders struct {
	OrderID     int    `json:"id"`
	UserID      int    `json:"userid"`
	Date        string `json:"date"`
	Sceneryname string `json:"sceneryname"`
	Type        string `json:"type"`
	Num         int    `json:"num"`
	Price       int    `json:"price"`
}
