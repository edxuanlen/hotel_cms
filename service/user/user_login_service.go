package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hotel_cms/common"
	"hotel_cms/entity"
	"hotel_cms/middleware"
	"hotel_cms/util"
	"hotel_cms/vo"
	"hotel_cms/vo/req"
)

type LoginServiceReq req.LoginServiceReq

// setCookie 设置cookie
func (service *LoginServiceReq) setCookie (c *gin.Context, user entity.User) {
	uid, _ := c.Cookie(common.CookieIdString)

	s.Set("user_id", user.Id)
	err := s.Save()
	if err != nil {util.Log().Error("set Session error s%", err)}
}

// Login 用户登录函数
func (service *LoginServiceReq) Login(c *gin.Context) vo.Response {
	var user entity.User

	err := c.ShouldBind(&user)

	// 解析 user 失败
	if err != nil {
		return vo.ParamErr(err)
	}


	if err := entity.DB.Where("user_name = ?", service.UserName).
		First(&user).Error; err != nil ||
		user.CheckPassword(service.Password) == false {
		return vo.ParamErr(nil, "账号或密码错误")
	}

	if user.CheckPassword(service.Password) == false {
		return vo.ParamErr(nil, "账号或密码错误")
	}

	// 设置session
	service.setSession(c, user)

	return vo.BuildUserResponse(user)
}
