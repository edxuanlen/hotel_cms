package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
	"hotel_cms/conf"
	"hotel_cms/entity"
	"hotel_cms/vo"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, vo.Response{
		Code: 0,
		Msg:  "Ping Success",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *entity.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*entity.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) vo.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return vo.ParamErr(
				err,
				fmt.Sprintf("%s%s", field, tag),
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return vo.ParamErr(err, "JSON类型不匹配")
	}

	return vo.ParamErr(err)
}
