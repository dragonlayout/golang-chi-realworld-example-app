package main

import (
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/infrastructure/db"
	"github.com/dragonlayout/golang-chi-realworld-example-app/model"
	"github.com/dragonlayout/golang-chi-realworld-example-app/repository"
	"log"
)

func main() {
	sqlDB := db.NewSqlite()
	db.AutoMigrate(sqlDB)
	userRepo := repository.NewUserRepository(sqlDB)
	userRepo.Create(&model.User{Username: "chenlong", Email: "dragonlayoutt@gmail.com", Password: "test"})

	user, err := userRepo.GetById(1)
	if err != nil {
		log.Fatal(err)
	}
	if user != nil {
		fmt.Println(user.Username)
		fmt.Println(user.Email)
		fmt.Println(user.Password)
	}
}
