package model

type UserWithoutToken struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio" binding:"-"`
}

// UserRegistrationRequest 用户注册请求参数
type UserRegistrationRequest struct {
	User UserWithoutToken `json:"user"`
}
