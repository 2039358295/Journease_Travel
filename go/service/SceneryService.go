package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
)

func GetAllComments(id int) (commentsList []*entity.Comments, err error) {
	if err := dao.SqlSession.Where("sceneryid=?", id).Find(&commentsList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAllPhotos() (photosList []*entity.Spotpic, err error) {
	if err := dao.SqlSession.Find(&photosList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAllOrders(userid int) (oderlist []*entity.Userorders, err error) {
	if err := dao.SqlSession.Where("userid=?", userid).Find(&oderlist).Error; err != nil {
		return nil, err
	}
	return
}
