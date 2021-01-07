package vo

import (
	"gopkg.in/dgrijalva/jwt-go.v3"
)

// LoginToken 登陆信息
type LoginToken struct {
	UserSimpleInfo
	jwt.StandardClaims
}