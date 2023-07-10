package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
)

// 新建spot信息
func CreateSpot1(spot *entity.Spot1) (err error) {
	if err = dao.UserSession.Create(spot).Error; err != nil {
		return err
	}
	return
}

// 获取spot集合
func GetAllSpot1() (spotList []*entity.Spot1, err error) {
	if err := dao.UserSession.Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 根据id查询景点spot
func GetSpotById1(uid int, sid int) (spot entity.Spot1, err error) {
	if err := dao.UserSession.Where("uid=? and sid=?", uid, sid).First(&spot).Error; err != nil {
		return spot, err
	}
	return
}

// 根据id删除spot
func DeleteSpotById1(uid int, sid int) (err error) {
	err = dao.UserSession.Where("uid=? and sid=?", uid, sid).Delete(&entity.Spot1{}).Error
	return
}

// 获取当前用户收藏列表
func GetUserSpot(uid int) (spotList []*entity.Spot1, err error) {
	if err := dao.UserSession.Where("uid=?", uid).Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 获取当前用户收藏列表
func GetC(sid int) (spotList []*entity.Spot1, err error) {
	if err := dao.UserSession.Where("sid=?", sid).Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}
