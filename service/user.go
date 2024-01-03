package service

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id       int    `gorm:"colum:id;primaryKey'"`
	Name     string `gorm:"colum:name"`
	Password string `gorm:"colum:password"`
}
