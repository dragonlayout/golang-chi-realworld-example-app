package persistence

import (
	"github.com/dragonlayout/golang-chi-realworld-example-app/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (userRepositoryImpl *UserRepositoryImpl) GetUser(id uint64) (*entity.User, error) {
	userRepositoryImpl.db.First()
	return nil, nil
}
