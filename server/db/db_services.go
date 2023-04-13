package db

import (
	"fmt"
	"gym/server/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func Transfer(connection *gorm.DB) {
	db = connection
}

func CreateRecord(data model.User) error {
	fmt.Println("DFJBDF", data)
	err := db.Create(&data).Error
	return err
}
