package service

import (
	"database/sql"
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	// Registration 用户注册
	Registration(r model.UserRegistrationRequest) error
}

type userService struct {
	database *sql.DB
}

func (u *userService) Registration(r model.UserRegistrationRequest) error {
	// 业务逻辑
	hash, err := bcrypt.GenerateFromPassword([]byte(r.User.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePW := string(hash)
	fmt.Println(encodePW)

	// 执行数据库操作
	return nil
}

func New(database *sql.DB) UserService {
	return &userService{
		database: database,
	}
}
