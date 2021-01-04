package serializer

import (
	"hotel_cms/model"
	"time"
)

// User 用户序列化器
type User struct {
	ID           uint      `json:"id"`
	IdentityCard string    `json:"identity_card"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Salary       int       `json:"salary"`
	Status       byte      `json:"status"`
	EntryTime    time.Time `json:"entry_time"`
	LastPayday   time.Time `json:"last_payday"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.Id,
		UserName:  user.Name,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
