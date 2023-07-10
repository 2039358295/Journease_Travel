package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
	"fmt"
)

func GetFP(UserID int) (FPList *[]entity.UserFP, err error) {
	FPList = new([]entity.UserFP)
	if err := dao.UserFP.Where("user_id=?", UserID).Find(&FPList).Error; err != nil {
		return nil, err
	}
	return
}

func AddFP(c entity.UserFP) (isFind bool, err error) {
	var temp entity.UserFP
	if err := dao.UserFP.Where("user_id=? AND post_id=?", c.UserID, c.PostID).First(&temp).Error; err != nil {
		if err = dao.UserFP.Create(c).Error; err != nil {
			return false, err
		} else {
			var post = entity.Post{
				UserID:     0,
				PostID:     0,
				Title:      "",
				Tag1:       "",
				Tag2:       "",
				Contents:   "",
				Photo:      nil,
				Photo1:     nil,
				Photo2:     nil,
				Photo3:     nil,
				Photo4:     nil,
				Photo5:     nil,
				Photo6:     nil,
				Photo7:     nil,
				Photo8:     nil,
				Likenum:    0,
				Collection: 0,
				Commentnum: 0,
				RTime:      "",
			}
			err = dao.PostSession.Where("post_id=?", c.PostID).First(&post).Error
			if err != nil {
				panic(err)
			} else {
				post.Collection += 1
				err = dao.PostSession.Model(entity.Post{}).Where("post_id=?", c.PostID).Update(post).Error
				if err != nil {
					panic(err)
				}
			}
			return true, err
		}
	} else {
		return false, err
	}
}

func DeleteFP(c entity.UserFP) (isFind bool, err error) {
	var temp entity.UserFP
	if err := dao.UserFP.Where("user_id=? AND post_id=?", c.UserID, c.PostID).First(&temp).Error; err == nil {
		if err = dao.UserFP.Where("user_id=? AND post_id=?", c.UserID, c.PostID).Delete(&entity.UserFP{}).Error; err != nil {
			return false, err
		} else {
			var post = entity.Post{
				UserID:     0,
				PostID:     0,
				Title:      "",
				Tag1:       "",
				Tag2:       "",
				Contents:   "",
				Photo:      nil,
				Photo1:     nil,
				Photo2:     nil,
				Photo3:     nil,
				Photo4:     nil,
				Photo5:     nil,
				Photo6:     nil,
				Photo7:     nil,
				Photo8:     nil,
				Likenum:    0,
				Collection: 0,
				Commentnum: 0,
				RTime:      "",
			}
			err = dao.PostSession.Where("post_id = ?", c.PostID).First(&post).Error
			if err != nil {
				panic(err)
			} else {
				post.Collection -= 1
				stmt := fmt.Sprintf("UPDATE `community`.`community_main` SET `collection` = ? WHERE (`post_id` = ?)")
				dao.PostSession.Exec(stmt, post.Collection, c.PostID)
			}
			return true, err
		}
	} else {
		return false, err
	}
}

func GetMyFP(UserID int) (FPList []*entity.UserFP, err error) {
	if err := dao.UserFP.Where("user_id=?", UserID).Find(&FPList).Error; err != nil {
		return nil, err
	}
	return
}
