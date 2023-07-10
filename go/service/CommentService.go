package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
	"fmt"
)

// CreateComment 创建评论
func CreateComment(c *entity.Comment) (err error) {
	var temp entity.Comment
	var Post entity.Post = entity.Post{
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
	err = dao.PostSession.Where("post_id=?", c.PostID).Find(&Post).Error
	if err != nil {
		panic(err)
		return err
	} else {
		Post.Commentnum += 1
		err = dao.PostSession.Model(&entity.Post{}).Where("post_id=?", c.PostID).Update(Post).Error
		if err != nil {
			return err
		}
	}
	err = dao.SqlSession.Where("post_id=?", c.PostID).Order("comment_id DESC").First(&temp).Error
	if err != nil {
		c.CommentID = 10001
	} else {
		c.CommentID = temp.CommentID + 1
	}
	if err = dao.SqlSession.Create(c).Error; err != nil {
		return err
	}
	return
}

// GetPostComment 获取当前帖子的所有评论
func GetPostComment(PostID int) (cList []*entity.Comment, err error) {
	if err := dao.SqlSession.Where("post_id=?", PostID).Order("up_number desc").Find(&cList).Error; err != nil {
		return nil, err
	}
	return
}

// DeleteCommentById 删评论
func DeleteCommentById(PostID int, CommentID int) (err error) {
	err = dao.SqlSession.Where("comment_id=? AND post_id=?", CommentID, PostID).Delete(&entity.Comment{}).Error
	if err != nil {
		return
	} else {
		var Post entity.Post = entity.Post{
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
		err = dao.PostSession.Where("post_id=?", PostID).Find(&Post).Error
		if err != nil {
			return
		} else {
			Post.Commentnum -= 1
			err = dao.PostSession.Model(&entity.Post{}).Where("post_id=?", PostID).Update(Post).Error
			if err != nil {
				return
			}
		}
	}
	return
}

// LikeComment 点赞评论
func LikeComment(PostID int, CommentID int, comment *entity.Comment) (err error) {
	if err = dao.SqlSession.Where("comment_id=? AND post_id=?", CommentID, PostID).Find(&comment).Error; err != nil {
		return nil
	}
	comment.UpNumber = comment.UpNumber + 1
	err = dao.SqlSession.Model(&entity.Comment{}).Where("comment_id=? AND post_id=?", CommentID, PostID).Update(comment).Error
	return
}

// DisLikeComment 取消点赞评论
func DisLikeComment(PostID int, CommentID int, comment *entity.Comment) (err error) {
	if err = dao.SqlSession.Where("comment_id=? AND post_id=?", CommentID, PostID).Find(&comment).Error; err != nil {
		return nil
	}
	comment.UpNumber = comment.UpNumber - 1
	stmt := fmt.Sprintf("UPDATE `community`.`community_comment` SET `up_number` = ? WHERE (`comment_id` = ? AND `post_id` = ?);")
	dao.SqlSession.Exec(stmt, comment.UpNumber, CommentID, PostID)
	return
}

func TimeSort(PostID int) (cList []*entity.Comment, err error) {
	if err := dao.SqlSession.Where("post_id=?", PostID).Order("time desc").Find(&cList).Error; err != nil {
		return nil, err
	}
	return
}
