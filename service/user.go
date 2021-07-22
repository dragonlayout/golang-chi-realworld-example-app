package service

import (
	"database/sql"
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/model"
	"github.com/google/uuid"
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

	_, err = u.database.Exec("INSERT INTO user (id, username, password, email) VALUE(?,?,?,?)",uuid.New(), r.User.Username, encodePW, r.User.Email)
	if err != nil {
		return err
	}
	return nil
}

func New(database *sql.DB) UserService {
	return &userService{
		database: database,
	}
}
