package entity

import "github.com/dgrijalva/jwt-go"

// Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (User) TableName() string {
	return "sys_user"
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Avatar     []byte `json:"avatar"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	DelStatus  int    `json:"delStatus"`
	CreateTime int64  `json:"createTime"`
	Token      string `json:"token"`
}

type Res struct {
	Mid       int
	Name      string
	Title     string
	Path      string
	Component string
}

/* 用户表的创建

CREATE TABLE `sys_user` (
`id` int(50) unsigned NOT NULL AUTO_INCREMENT,
`name` varchar(50) NOT NULL COMMENT '用户名',
`nick_name` varchar(150) DEFAULT NULL COMMENT '昵称',
`avatar` varchar(150) DEFAULT NULL COMMENT '头像',
`password` varchar(100) DEFAULT NULL COMMENT '密码',
`email` varchar(100) DEFAULT NULL COMMENT '邮箱',
`mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
`create_time` bigint(50) DEFAULT NULL COMMENT '更新时间',
`del_status` tinyint(4) DEFAULT '0' COMMENT '是否删除 -1：已删除   0：正常',
PRIMARY KEY (`id`),
UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='用户表';

*/
