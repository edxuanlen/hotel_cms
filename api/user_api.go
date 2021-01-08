package api

import (
	"hotel_cms/common"
	"hotel_cms/service/user"
	"hotel_cms/vo"

	"github.com/gin-gonic/gin"
)

var userLoginService user.LoginServiceReq
var userRegisterService user.RegisterServiceReq

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {

	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(res.Code, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	c.SetCookie(common.CookieJwtString, common.CookieJwtString,  -1, "/", "localhost", false, true)
	c.JSON(200, vo.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
