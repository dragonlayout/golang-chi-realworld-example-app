package persistence

import (
	"gorm.io/gorm"
)

type Repositories struct {
	db *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	// TODO: db connection url
	// TODO: db connection pool
	db, err := gorm.Open()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		db: db,
	}, nil
}

func (s *Repositories) AutoMigrate() error {
	s.db.AutoMigrate()
}
