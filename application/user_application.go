package application

import "github.com/dragonlayout/golang-chi-realworld-example-app/domain/repository"

type UserApplication struct {
	userRepository repository.UserRepository
}

var _ UserApplicationInterface = &UserApplication{}
