package vo

import (
	"time"
)

const (
	userSimpleInfo = iota
	userAllInfo
)

type User interface {
	BuildUser()
}

// UserSimpleInfo 用户简单信息序列化器
type UserSimpleInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Access byte   `json:"access"`
}

// UserAllInfo 用户详情序列化器
type UserAllInfo struct {
	Id           int       `json:"id"`
	IdentityCard string    `json:"identity_card"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Access       byte      `json:"access"`
	Salary       int       `json:"salary"`
	Status       byte      `json:"status"`
	EntryTime    time.Time `json:"entry_time"`
	LastPayday   time.Time `json:"last_payday"`
}


// (u *UserSimpleInfo) BuildUser 序列化用户简单信息
func (u *UserSimpleInfo) BuildUser() UserSimpleInfo {
	return UserSimpleInfo{
		Id:     u.Id,
		Name:   u.Name,
		Phone:  u.Phone,
		Access: u.Access,
	}
}

// (u *UserAllInfo) BuildUser 序列化用户所有信息
func (u *UserAllInfo) BuildUser() UserAllInfo {
	return UserAllInfo{
		Id:           u.Id,
		IdentityCard: u.IdentityCard,
		Name:         u.Name,
		Phone:        u.Phone,
		Access:       u.Access,
		Salary:       u.Salary,
		Status:       u.Status,
		EntryTime:    u.EntryTime,
		LastPayday:   u.LastPayday,
	}
}

// (u *UserSimpleInfo) BuildUserResponse 序列化用户简单信息响应
func (u *UserSimpleInfo) BuildUserResponse() Response {
	return Response{
		Data: u.BuildUser(),
	}
}

// (u *UserAllInfo) BuildUserResponse 序列化用户所有信息响应
func (u *UserAllInfo) BuildUserResponse() Response {
	return Response{
		Data: u.BuildUser(),
	}
}
