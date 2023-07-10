package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (Spot1) TableName() string {
	return "spot_collection"
}

type Spot1 struct {
	Id  int `json:"id"`
	Uid int `json:"uid"`
	Sid int `json:"sid"`
}
