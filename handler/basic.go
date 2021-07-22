package handler

import "github.com/dragonlayout/golang-chi-realworld-example-app/model"

type UserHandler interface {
	Registration(r *model.UserRegistrationRequest)
}

type userHandler struct {
	// TODO: service
}
