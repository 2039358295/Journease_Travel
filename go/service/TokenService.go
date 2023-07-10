package service

import (
	"Gin_Project/go/entity"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 指定加密密钥
var jwtSecret = []byte("setting.JwtSecret")

// 根据用户的用户名和密码产生token
func GenerateToken(email, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := entity.Claims{
		Email:    email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证token
func JwtVerify(c *gin.Context) (token string) {
	//过滤是否验证token
	t := c.GetHeader("token")
	if t == "" {
		panic("token not exist !")
	}
	return t
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseTokenHs256(c *gin.Context, tokenString string) (*entity.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*entity.Claims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
