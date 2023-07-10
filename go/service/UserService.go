package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
)

// 新建User信息
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

// 获取user集合
func GetAllUser() (userList []*entity.User, err error) {
	if err := dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// 根据id查询用户User
func GetUserById(id int) (user entity.User, err error) {
	if err := dao.SqlSession.Where("id=?", id).First(&user).Error; err != nil {
		return user, err
	}
	return
}

// 根据email查询用户User
func GetUserByEmail(email string) (user entity.User, err error) {
	if err := dao.SqlSession.Where("email=?", email).First(&user).Error; err != nil {
		return user, err
	}
	return
}

// 根据id删除user
func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

// 更新用户信息
func UpdateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}

func CheckUserName(ID int) (user entity.User) {
	err := dao.SqlSession.Where("id=?", ID).First(&user).Error
	if err != nil {
		return user
	} else {
		return user
	}
}

func GetUserAvatar(userID int) (user entity.User) {
	if err := dao.PostSession.Where("id = ?", userID).Find(&user).Error; err != nil {
		return
	}
	return user
}
