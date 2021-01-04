package vo

import (
	"hotel_cms/entity"
	"time"
)

const (
	userSimpleInfo = iota
	userAllInfo
)

type User interface {
	BuildUser()
}

// User 用户简单信息序列化器
type UserSimpleInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Access byte   `json:"access"`
}

// User 用户详情序列化器
type UserAllInfo struct {
	ID           int       `json:"id"`
	IdentityCard string    `json:"identity_card"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Access       byte      `json:"access"`
	Salary       int       `json:"salary"`
	Status       byte      `json:"status"`
	EntryTime    time.Time `json:"entry_time"`
	LastPayday   time.Time `json:"last_payday"`
}


// Simple Info 序列化用户
func (u *UserSimpleInfo) BuildUser(user entity.User) UserSimpleInfo {
	return UserSimpleInfo{
		ID:     user.Id,
		Name:   user.Name,
		Phone:  user.Password,
		Access: user.Access,
	}
}

// All Info 序列化用户
func (u *UserAllInfo) BuildUser(user entity.User) UserAllInfo {
	return UserAllInfo{
		ID:           user.Id,
		IdentityCard: user.IdentityCard,
		Name:         user.Name,
		Phone:        user.Password,
		Access:       user.Access,
		Salary:       user.Salary,
		Status:       user.Status,
		EntryTime:    user.EntryTime,
		LastPayday:   user.LastPayday,
	}
}

// BuildUserResponse 序列化用户响应
func (u *UserSimpleInfo) BuildUserResponse(user entity.User) Response {
	return Response{
		Data: u.BuildUser(user),
	}
}

// BuildUserResponse 序列化用户响应
func (u *UserAllInfo) BuildUserResponse(user entity.User) Response {
	return Response{
		Data: u.BuildUser(user),
	}
}
