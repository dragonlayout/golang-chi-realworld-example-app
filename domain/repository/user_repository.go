package repository

import "github.com/dragonlayout/golang-chi-realworld-example-app/domain/entity"

type UserRepository interface {
	GetUser(uint64) (*entity.User, error)
}
