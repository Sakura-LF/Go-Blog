package service

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model

	Id         int       `gorm:"column:id;primaryKey"`
	UserId     int       `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	Article    string    `gorm:"column:article"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
