package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"hotel_cms/common"
	"hotel_cms/entity"
	"hotel_cms/vo"
	"time"
)

type LoginToken vo.LoginToken

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*entity.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, vo.NoLogin())
		c.Abort()
	}
}

// GenToken 生成JWT
func GenToken(info vo.UserSimpleInfo) (string, error) {
	// 创建一个我们自己的声明
	c := LoginToken{
		UserSimpleInfo: vo.UserSimpleInfo{
			Id:     info.Id,
			Name:   info.Name,
			Phone:  info.Phone,
			Access: info.Access,
		},
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(common.TokenExpireDuration).Unix(),
			// 签发人
			Issuer: common.ProjectName,
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(common.TokenSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*LoginToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &LoginToken{},
		func(token *jwt.Token) (i interface{}, err error) { return common.TokenSecret, nil })
	if err != nil {
		return nil, err
	}

	// 校验token
	if claims, ok := token.Claims.(*LoginToken); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
