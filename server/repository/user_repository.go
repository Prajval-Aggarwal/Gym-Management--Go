package repository

import (
	"gym/server/model"

	"gorm.io/gorm"
)

type UserRespository struct {
	DB *gorm.DB
}

func (repository UserRespository) Create(user *model.User) {
	repository.DB.Create(user)
}
