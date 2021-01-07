package common

import "time"

const (
	// CookieIdString cookie中存放user id的key
	CookieIdString   = "user_id"

	// CookieJwtString cookie中存放 jwt 的key
	CookieJwtString = "jwt-token"

	// RedisKeyPrefix redis中存放user的key prefix
	RedisKeyPrefix   = "user_id_"

	// PassWordCost 密码加密难度
	PassWordCost = 12

	// TokenExpireDuration jwt 过期时间
	TokenExpireDuration = time.Hour * 24

	// TokenSecret jwt token 密钥
	TokenSecret = "g0gogO"

	// ProjectName 项目名
	ProjectName = "hotel_cms"

)


