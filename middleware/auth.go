package middleware

import (
	"github.com/gin-gonic/gin"
	"hotel_cms/entity"
	"hotel_cms/vo"
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
