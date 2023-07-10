package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
	"fmt"
)

var PostSort = "time"
var MySort = "time"

func CreatePost(post *entity.Post) (err error) {
	var temp entity.Post
	if err = dao.PostSession.Order("post_id DESC").First(&temp).Error; err != nil {
		post.PostID = 1
	} else {
		post.PostID = temp.PostID + 1
	}
	if err = dao.PostSession.Create(post).Error; err != nil {
		return err
	}
	return
}

func ChangePostSort() (postList []*entity.Post, err error) {
	if PostSort == "time" {
		if err := dao.PostSession.Order("likenum DESC").Find(&postList).Error; err != nil {
			return nil, err
		}
		PostSort = "likenum"
	} else {
		if err := dao.PostSession.Order("r_time DESC").Find(&postList).Error; err != nil {
			return nil, err
		}
		PostSort = "time"
	}
	return
}

func ChangeMySort(id int) (mypost []*entity.Post, err error) {
	if MySort == "time" {
		if err := dao.PostSession.Where("user_id=?", id).Order("likenum DESC").Find(&mypost).Error; err != nil {
			return nil, err
		}
		MySort = "likenum"
	} else {
		if err := dao.PostSession.Where("user_id=?", id).Order("r_time DESC").Find(&mypost).Error; err != nil {
			return nil, err
		}
		MySort = "time"
	}
	return
}

func GetAllPost() (postList []*entity.Post, err error) {
	if err := dao.PostSession.Order("likenum DESC").Find(&postList).Error; err != nil {
		return nil, err
	}
	return
}

func GetPosts() (postList []*entity.Post, err error) {
	if err := dao.PostSession.Order("r_time DESC").Find(&postList).Error; err != nil {
		return nil, err
	}
	return
}

func GetMyPost(id int) (mypost []*entity.Post, err error) {

	if err := dao.PostSession.Where("user_id=?", id).Order("likenum DESC").Find(&mypost).Error; err != nil {
		return nil, err
	}
	return
}

func DeletePostByID(id int) (err error) {
	if err = dao.PostSession.Where("post_id=?", id).Delete(&entity.Post{}).Error; err != nil {
		panic(err)
	}
	if err = dao.UserFP.Where("post_id=?", id).Delete(&entity.Post{}).Error; err != nil {
		panic(err)
	}
	return
}

func GetPostByTag(tag string) (post []*entity.Post, err error) {
	if err = dao.PostSession.Where("tag1=? OR tag2=?", tag, tag).Find(&post).Error; err != nil {
		return nil, err
	}
	return
}

func GetPostByID(id int) (post []*entity.Post, err error) {
	if err = dao.PostSession.Where("post_id=?", id).Find(&post).Error; err != nil {
		return nil, err
	}
	return
}

func UpdatePost(id int, post *entity.Post) (err error) {
	if err = dao.PostSession.Model(&entity.Post{}).Where("post_id=?", id).Update(post).Error; err != nil {
		return err
	}
	return
}

func LikePost(PostID int, post *entity.Post) (err error) {
	if err = dao.PostSession.Where("post_id=?", PostID).Find(&post).Error; err != nil {
		return nil
	}
	post.Likenum += 1
	err = dao.SqlSession.Model(&entity.Post{}).Where("post_id=?", PostID).Update(post).Error
	return
}

func DislikePost(PostID int, post *entity.Post) (err error) {
	if err = dao.PostSession.Where("post_id=?", PostID).Find(&post).Error; err != nil {
		return nil
	}
	post.Likenum -= 1
	stmt := fmt.Sprintf("UPDATE `community`.`community_main` SET `likenum` = ? WHERE (`post_id` = ?);")
	dao.PostSession.Exec(stmt, post.Likenum, PostID)
	return
}

func GetImages(PostID int) (post entity.Post) {
	if err := dao.PostSession.Where("post_id = ?", PostID).Find(&post).Error; err != nil {
		return
	}
	return post
}

func SearchPost(tag string) (post []*entity.Post, err error) {
	likeTag := "%" + tag + "%"
	if err = dao.PostSession.Where("tag1 LIKE ? OR tag2 LIKE ? OR title LIKE ? OR contents LIKE ?", likeTag, likeTag, likeTag, likeTag).Find(&post).Error; err != nil {
		return nil, err
	}
	return
}
