package user

import (
	"hotel_cms/common"
	"hotel_cms/db_op"
	"hotel_cms/entity"
	"hotel_cms/util"
	"hotel_cms/vo"
	"hotel_cms/vo/req"
	"time"
)

type RegisterServiceReq req.RegisterServiceReq

// valid 验证表单
func (service *RegisterServiceReq) valid() *vo.Response {
	if service.PasswordConfirm != service.Password {
		return &vo.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	db_op.DB.Table(common.User).Model(&entity.User{}).Where("phone = ?", service.Phone).Count(&count)
	if count > 0 {
		return &vo.Response{
			Code: 40001,
			Msg:  "手机已被注册",
		}
	}

	count = 0
	db_op.DB.Table(common.User).Model(&entity.User{}).Where("identity_card = ?", service.IdentityCard).Count(&count)
	if count > 0 {
		return &vo.Response{
			Code: 40001,
			Msg:  "身份证已注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *RegisterServiceReq) Register() vo.Response {
	user := entity.User{
		Name:         service.UserName,
		IdentityCard: service.IdentityCard,
		Phone: service.Phone,
		Password: service.Password,
		Status: 1,
		Access: 9,
		EntryTime: time.Now(),
		LastPayday: time.Now(),
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return vo.Err(
			vo.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := db_op.DB.Table(common.User).Create(&user).Error; err != nil {
		return vo.DBErr(err, "注册失败")
	}

	var userSimple vo.UserSimpleInfo
	util.Convert(&user, &userSimple)
	return userSimple.BuildUserResponse()
}
