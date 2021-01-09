package middleware

import (
	"github.com/gin-gonic/gin"
	"hotel_cms/common"
	"hotel_cms/util"
	"hotel_cms/vo"
)

type LoginToken vo.LoginToken

// IsLogin 登录状态
func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, _ := c.Cookie(common.CookieJwtString); token != "" {
			c.Next()
			return
		}

		c.JSON(200, vo.NoLogin())
		c.Abort()
	}
}

// AuthRequired 登录权限校验
func AuthRequired(level byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, _ := c.Cookie(common.CookieJwtString); token != "" {
			// 解析token中的信息
			loginToken, err := util.ParseToken(token)
			if err == nil {
				// 需要的权限level
				if loginToken.Access == level ||
					// 拥有所有权限的root
					loginToken.Access == common.AllPermission {
					c.Next()
					return
				}
			}
		}

		c.JSON(200, vo.InsufficientAuthority())
		c.Abort()
	}
}
