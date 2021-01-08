package db_op

import (
	"hotel_cms/entity"
)

type User entity.User

// GetUser 用ID获取用户
func GetUser(Id interface{}) (User, error) {
	var user User
	result := DB.First(&user, Id)
	return user, result.Error
}
