package user

import (
	"github.com/gin-gonic/gin"
	redis "hotel_cms/cache"
	"hotel_cms/common"
	"hotel_cms/db_op"
	"hotel_cms/entity"
	"hotel_cms/util"
	"hotel_cms/vo"
	"hotel_cms/vo/req"
	"strconv"
	)

type LoginServiceReq req.LoginServiceReq

// setCookie 设置cookie
func (service *LoginServiceReq) setCookie (c *gin.Context, token string) {
	c.SetCookie(common.CookieJwtString, token, 86400, "/",
		"localhost", false, true)
}

// Login 用户登录函数
func (service *LoginServiceReq) Login() vo.Response {

	var user entity.User

	// 用户为空或者密码错误
	if err := db_op.DB.Table(common.User).Where("name = ?", service.UserName).
		First(&user).Error; err != nil ||
		user.CheckPassword(service.Password) == false {
		return vo.ParamErr(nil, "账号或密码错误")
	}

	// 设置redis
	redis.RedisClient.Set(common.RedisKeyPrefix + strconv.Itoa(user.Id), user, common.TokenExpireDuration)

	userSimpleInfo := vo.UserSimpleInfo{}
	util.Convert(&user, &userSimpleInfo)

	token, err := util.GenToken(userSimpleInfo)

	if err != nil {
		return vo.Err(501, "生成token错误", err)
	}

	return vo.Response{
		Code:  200,
		Data:  token,
		Msg:   "成功",
	}
}
