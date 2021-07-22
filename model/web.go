package model

// UserRegistrationRequest 用户注册请求参数
type UserRegistrationRequest struct {
	User struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
}
