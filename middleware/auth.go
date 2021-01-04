package middleware

import (
	"hotel_cms/cache"
	"hotel_cms/model"
	"hotel_cms/serializer"

	"github.com/gin-gonic/gin"
)

const (
	RedisKeyPrefix = "user_id_"
	CookieName = "user_id"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, _ := c.Cookie(CookieName)

		if uid != "" {
			uid := cache.RedisClient.Get(RedisKeyPrefix + uid)

			// 获取用户信息
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.NoLogin())
		c.Abort()
	}
}
