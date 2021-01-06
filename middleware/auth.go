package middleware

import (
	"github.com/gin-gonic/gin"
	"hotel_cms/cache"
	"hotel_cms/common"
	"hotel_cms/entity"
	"hotel_cms/vo"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, _ := c.Cookie(common.CookieIdString)

		if uid != "" {
			var user entity.User
			user, _ = (entity.User) cache.RedisClient.Get(common.RedisKeyPrefix + uid)
			var err error
			err = nil
			if user == nil {
				// 获取用户信息
				user, err := entity.GetUser(uid)
			}

			if err == nil {
				c.SetCookie("user", &user, 86400, "/", "localhost")
			}
		}
		c.Next()
	}
}

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
