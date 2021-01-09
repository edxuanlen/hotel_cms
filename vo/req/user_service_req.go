package req

// RegisterServiceReq 管理用户注册服务
type RegisterServiceReq struct {
	UserName        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	IdentityCard        string `form:"identityCard" json:"identityCard" binding:"required,min=2,max=30"`
	Phone string `form:"phone" json:"phone" binding:"required,min=11,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// UserLoginService 管理用户登录的服务
type LoginServiceReq struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}
